package ai

import "google.golang.org/genai"

type GeminiConfig struct {
	APIKey string
	Model  string
}

func (c *GeminiConfig) CreateGenerateContentConfig(cachedContentName string) *genai.GenerateContentConfig {

	budget := int32(4096)
	temperature := float32(0.2)

	return &genai.GenerateContentConfig{
		Temperature:    &temperature,
		CachedContent:  cachedContentName,
		ThinkingConfig: &genai.ThinkingConfig{IncludeThoughts: true, ThinkingBudget: &budget},
	}

}
