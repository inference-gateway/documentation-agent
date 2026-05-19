package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	server "github.com/inference-gateway/adk/server"
	zap "go.uber.org/zap"
)

// GetLibraryDocsTool struct holds the tool with services
type GetLibraryDocsTool struct {
	logger *zap.Logger
}

// NewGetLibraryDocsTool creates a new get_library_docs tool
func NewGetLibraryDocsTool(logger *zap.Logger) server.Tool {
	tool := &GetLibraryDocsTool{
		logger: logger,
	}
	return server.NewBasicTool(
		"get_library_docs",
		"Fetches up-to-date documentation for a library using Context7-compatible library ID",
		map[string]any{
			"type": "object",
			"properties": map[string]any{
				"libraryId": map[string]any{
					"description": "Exact Context7-compatible library ID (e.g., '/mongodb/docs', '/vercel/next.js', '/supabase/supabase', '/vercel/next.js/v14.3.0-canary.87') retrieved from resolve_library_id or supplied by the caller in the format '/org/project' or '/org/project/version'.",
					"type":        "string",
				},
				"query": map[string]any{
					"description": "Specific question or task to answer from the documentation. Be specific - 'How to set up JWT auth in Express.js' or 'React useEffect cleanup examples', not 'auth' or 'hooks'. Do not include secrets.",
					"type":        "string",
				},
			},
			"required": []string{"libraryId", "query"},
		},
		tool.GetLibraryDocsHandler,
	)
}

// GetLibraryDocsHandler handles the get_library_docs tool execution
func (t *GetLibraryDocsTool) GetLibraryDocsHandler(ctx context.Context, args map[string]any) (string, error) {
	t.logger.Debug("get_library_docs handler called", zap.Any("args", args))

	libraryID, ok := args["libraryId"].(string)
	if !ok || strings.TrimSpace(libraryID) == "" {
		return "", fmt.Errorf("libraryId parameter is required and must be a non-empty string")
	}
	if !strings.HasPrefix(libraryID, "/") {
		return "", fmt.Errorf("libraryId must be in format '/org/project' or '/org/project/version', got: %s", libraryID)
	}

	query, ok := args["query"].(string)
	if !ok || strings.TrimSpace(query) == "" {
		return "", fmt.Errorf("query parameter is required and must be a non-empty string")
	}

	t.logger.Info("fetching documentation", zap.String("libraryID", libraryID), zap.String("query", query))

	params := url.Values{}
	params.Set("libraryId", libraryID)
	params.Set("query", query)
	endpoint := context7APIBase + "/context?" + params.Encode()

	body, err := context7Get(ctx, t.logger, endpoint)
	if err != nil {
		return "", err
	}

	documentation := strings.TrimSpace(string(body))
	if documentation == "" {
		t.logger.Warn("no documentation returned", zap.String("libraryID", libraryID))
		return fmt.Sprintf(`{"error": "No documentation found for library: %s"}`, libraryID), nil
	}

	response := map[string]any{
		"libraryID":     libraryID,
		"query":         query,
		"documentation": documentation,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		t.logger.Error("failed to marshal response", zap.Error(err))
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}

	t.logger.Info("successfully retrieved documentation",
		zap.String("libraryID", libraryID),
		zap.Int("chars", len(documentation)))
	return string(responseJSON), nil
}
