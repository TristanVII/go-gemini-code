package ai

type GeminiConfig struct {
	APIKey string
	// Backend is the GenAI backend to use for the client.
	Backend int
	Model   string
}

// Add tools later
func (p *GeminiConfig) createCachedConfig() {

}
