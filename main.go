package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

// Agent represents the agent configuration
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
			PushNotifications       bool `yaml:"pushNotifications"`
			StateTransitionHistory  bool `yaml:"stateTransitionHistory"`
		} `yaml:"capabilities"`
		Agent struct {
			Provider     string  `yaml:"provider"`
			Model        string  `yaml:"model"`
			SystemPrompt string  `yaml:"systemPrompt"`
			MaxTokens    int     `yaml:"maxTokens"`
			Temperature  float64 `yaml:"temperature"`
		} `yaml:"agent"`
		Tools []struct {
			Name        string      `yaml:"name"`
			Description string      `yaml:"description"`
			Schema      interface{} `yaml:"schema"`
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
	} `yaml:"spec"`
}

// loadAgentConfig loads the agent configuration from agent.yaml
func loadAgentConfig() (*Agent, error) {
	file, err := os.ReadFile("agent.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read agent.yaml: %w", err)
	}

	var agent Agent
	if err := yaml.Unmarshal(file, &agent); err != nil {
		return nil, fmt.Errorf("failed to parse agent.yaml: %w", err)
	}

	return &agent, nil
}

// healthHandler handles health check requests
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(response)
}

// toolsHandler handles tool discovery requests
func toolsHandler(agent *Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(agent.Spec.Tools)
	}
}

// executeHandler handles tool execution requests
func executeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toolName := vars["tool"]

	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Mock response for now - in a real implementation, this would execute the actual tool
	response := map[string]interface{}{
		"tool":    toolName,
		"request": request,
		"result":  fmt.Sprintf("Executed %s with parameters", toolName),
		"status":  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// agentCardHandler serves the agent metadata
func agentCardHandler(agent *Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		agentCard := map[string]interface{}{
			"schemaVersion": "0.1.0",
			"name":          agent.Metadata.Name,
			"version":       agent.Metadata.Version,
			"description":   agent.Metadata.Description,
			"capabilities":  agent.Spec.Capabilities,
			"tools":         agent.Spec.Tools,
		}

		json.NewEncoder(w).Encode(agentCard)
	}
}

func main() {
	// Load agent configuration
	agent, err := loadAgentConfig()
	if err != nil {
		log.Fatalf("Failed to load agent configuration: %v", err)
	}

	// Create HTTP router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/tools", toolsHandler(agent)).Methods("GET")
	router.HandleFunc("/execute/{tool}", executeHandler).Methods("POST")
	router.HandleFunc("/.well-known/agent.json", agentCardHandler(agent)).Methods("GET")

	// Get port from environment or use default
	port := agent.Spec.Server.Port
	if port == 0 {
		port = 8080
	}

	log.Printf("Starting %s agent on port %d", agent.Metadata.Name, port)
	log.Printf("Description: %s", agent.Metadata.Description)
	log.Printf("Version: %s", agent.Metadata.Version)

	// Start HTTP server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}