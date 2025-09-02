package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

// Agent represents the ADL agent configuration
type Agent struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Version     string `yaml:"version"`
	} `yaml:"metadata"`
	Spec struct {
		Capabilities struct {
			Streaming               bool `yaml:"streaming"`
			PushNotifications      bool `yaml:"pushNotifications"`
			StateTransitionHistory bool `yaml:"stateTransitionHistory"`
		} `yaml:"capabilities"`
		Agent struct {
			Provider     string  `yaml:"provider"`
			Model        string  `yaml:"model"`
			SystemPrompt string  `yaml:"systemPrompt"`
			MaxTokens    int     `yaml:"maxTokens"`
			Temperature  float64 `yaml:"temperature"`
		} `yaml:"agent"`
		Tools []struct {
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Schema      struct {
				Type       string `yaml:"type"`
				Properties map[string]interface{} `yaml:"properties"`
				Required   []string `yaml:"required"`
			} `yaml:"schema"`
		} `yaml:"tools"`
		Server struct {
			Port  int  `yaml:"port"`
			Debug bool `yaml:"debug"`
		} `yaml:"server"`
		Language struct {
			Go struct {
				Module  string `yaml:"module"`
				Version string `yaml:"version"`
			} `yaml:"go"`
		} `yaml:"language"`
		Sandbox struct {
			Flox struct {
				Enabled bool `yaml:"enabled"`
			} `yaml:"flox"`
		} `yaml:"sandbox"`
	} `yaml:"spec"`
}

// Tool represents an agent tool
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Schema      map[string]interface{} `json:"schema"`
}

// Server represents the A2A agent server
type Server struct {
	agent  Agent
	router *mux.Router
}

// NewServer creates a new agent server
func NewServer(configPath string) (*Server, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var agent Agent
	if err := yaml.Unmarshal(data, &agent); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	s := &Server{
		agent:  agent,
		router: mux.NewRouter(),
	}

	s.setupRoutes()
	return s, nil
}

// setupRoutes configures the HTTP routes
func (s *Server) setupRoutes() {
	s.router.HandleFunc("/health", s.healthHandler).Methods("GET")
	s.router.HandleFunc("/tools", s.toolsHandler).Methods("GET")
	s.router.HandleFunc("/execute", s.executeHandler).Methods("POST")
}

// healthHandler handles health check requests
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "healthy",
		"agent":  s.agent.Metadata.Name,
		"version": s.agent.Metadata.Version,
	})
}

// toolsHandler returns available tools
func (s *Server) toolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var tools []Tool
	for _, tool := range s.agent.Spec.Tools {
		tools = append(tools, Tool{
			Name:        tool.Name,
			Description: tool.Description,
			Schema:      map[string]interface{}(tool.Schema.Properties),
		})
	}
	
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tools": tools,
	})
}

// executeHandler handles tool execution requests
func (s *Server) executeHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Tool   string                 `json:"tool"`
		Params map[string]interface{} `json:"params"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// For now, return a mock response
	// In a real implementation, this would call the actual Context7 MCP server
	response := map[string]interface{}{
		"tool":   request.Tool,
		"result": "Mock result - implement actual Context7 integration here",
		"status": "success",
	}

	json.NewEncoder(w).Encode(response)
}

// Start starts the HTTP server
func (s *Server) Start() error {
	port := s.agent.Spec.Server.Port
	if port == 0 {
		port = 8080
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.router,
	}

	// Channel to listen for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting agent server on port %d", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func main() {
	configPath := "agent.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	server, err := NewServer(configPath)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}