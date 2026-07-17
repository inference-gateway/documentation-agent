package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	resty "github.com/go-resty/resty/v2"
	server "github.com/inference-gateway/adk/server"
	zap "go.uber.org/zap"
)

// context7APIBase is the documented REST v2 base URL.
// See https://context7.com/docs/howto/api-keys for auth.
const context7APIBase = "https://context7.com/api/v2"

// ResolveLibraryIDTool struct holds the tool with services
type ResolveLibraryIDTool struct {
	logger *zap.Logger
}

// NewResolveLibraryIDTool creates a new resolve_library_id tool
func NewResolveLibraryIDTool(logger *zap.Logger) server.Tool {
	tool := &ResolveLibraryIDTool{
		logger: logger,
	}
	return server.NewBasicTool(
		"resolve_library_id",
		"Resolves library name to Context7-compatible library ID and returns matching libraries",
		map[string]any{
			"type": "object",
			"properties": map[string]any{
				"libraryName": map[string]any{
					"description": "Official library name with proper punctuation (e.g., 'Next.js', 'Three.js', 'Customer.io'). Used by Context7 to rank candidate libraries.",
					"type":        "string",
				},
				"query": map[string]any{
					"description": "The question or task the caller is trying to accomplish. Used by Context7 to rank library results by relevance. Do not include secrets.",
					"type":        "string",
				},
			},
			"required": []string{"libraryName", "query"},
		},
		tool.ResolveLibraryIDHandler,
	)
}

// context7SearchResult mirrors a single entry in /api/v2/libs/search results.
type context7SearchResult struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Branch         string   `json:"branch,omitempty"`
	LastUpdateDate string   `json:"lastUpdateDate,omitempty"`
	State          string   `json:"state,omitempty"`
	TotalTokens    int      `json:"totalTokens,omitempty"`
	TotalSnippets  int      `json:"totalSnippets,omitempty"`
	Stars          int      `json:"stars,omitempty"`
	TrustScore     float64  `json:"trustScore,omitempty"`
	BenchmarkScore float64  `json:"benchmarkScore,omitempty"`
	Versions       []string `json:"versions,omitempty"`
}

// ResolveLibraryIDHandler handles the resolve_library_id tool execution
func (t *ResolveLibraryIDTool) ResolveLibraryIDHandler(ctx context.Context, args map[string]any) (string, error) {
	span := startToolSpan(ctx, "resolve_library_id")
	defer span.End()

	t.logger.Debug("resolve_library_id handler called", zap.Any("args", args))

	libraryName, ok := args["libraryName"].(string)
	if !ok || strings.TrimSpace(libraryName) == "" {
		return "", fmt.Errorf("libraryName parameter is required and must be a non-empty string")
	}
	query, ok := args["query"].(string)
	if !ok || strings.TrimSpace(query) == "" {
		return "", fmt.Errorf("query parameter is required and must be a non-empty string")
	}

	t.logger.Info("searching for library", zap.String("libraryName", libraryName), zap.String("query", query))

	params := url.Values{}
	params.Set("libraryName", libraryName)
	params.Set("query", query)
	endpoint := context7APIBase + "/libs/search?" + params.Encode()

	body, err := context7Get(ctx, t.logger, endpoint)
	if err != nil {
		return "", err
	}

	var parsed struct {
		Results []context7SearchResult `json:"results"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		t.logger.Error("failed to parse Context7 search response", zap.Error(err), zap.String("body", truncateString(string(body), 500)))
		return "", fmt.Errorf("failed to parse Context7 search response: %w", err)
	}

	if len(parsed.Results) == 0 {
		t.logger.Warn("no libraries matched", zap.String("libraryName", libraryName))
		return fmt.Sprintf(`{"error": "No libraries matched %q", "totalMatches": 0}`, libraryName), nil
	}

	selected := parsed.Results[0]
	response := map[string]any{
		"selectedLibraryID": selected.ID,
		"selectedLibrary":   selected,
		"allMatches":        parsed.Results,
		"totalMatches":      len(parsed.Results),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		t.logger.Error("failed to marshal response", zap.Error(err))
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}

	t.logger.Info("successfully resolved library ID",
		zap.String("libraryName", libraryName),
		zap.String("selectedID", selected.ID),
		zap.Int("totalMatches", len(parsed.Results)))
	return string(responseJSON), nil
}

// context7Get performs a GET against Context7's REST v2 API with
// `Authorization: Bearer` auth per
// https://context7.com/docs/howto/api-keys. Returns the raw response
// body on 2xx, or a wrapped error otherwise.
func context7Get(ctx context.Context, logger *zap.Logger, endpoint string) ([]byte, error) {
	apiKey := os.Getenv("CONTEXT7_API_KEY")
	if apiKey == "" {
		logger.Warn("CONTEXT7_API_KEY not set, proceeding without authentication")
	} else {
		logger.Debug("using Context7 API key", zap.String("keyPrefix", apiKey[:min(8, len(apiKey))]+"..."))
	}

	client := resty.New()
	if logger.Core().Enabled(zap.DebugLevel) {
		client.SetDebug(true)
	}

	req := client.R().
		SetContext(ctx).
		SetHeader("User-Agent", "documentation-agent/0.1.0").
		SetHeader("Accept", "application/json, text/plain;q=0.9, */*;q=0.5")
	if apiKey != "" {
		req.SetHeader("Authorization", "Bearer "+apiKey)
	}

	logger.Debug("making Context7 API request", zap.String("url", endpoint))

	resp, err := req.Get(endpoint)
	if err != nil {
		logger.Error("request to Context7 API failed", zap.Error(err))
		return nil, fmt.Errorf("failed to make request to Context7 API: %w", err)
	}

	body := resp.Body()
	logger.Debug("received response from Context7",
		zap.Int("statusCode", resp.StatusCode()),
		zap.String("body", truncateString(string(body), 1000)))

	switch resp.StatusCode() {
	case http.StatusOK:
		return body, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("invalid Context7 API key (status 401): check CONTEXT7_API_KEY")
	case http.StatusNotFound:
		return nil, fmt.Errorf("Context7 resource not found (status 404): %s", truncateString(string(body), 300))
	default:
		return nil, fmt.Errorf("Context7 API returned status %d: %s", resp.StatusCode(), truncateString(string(body), 500))
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
