package ai

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type AIClient struct {
	C             *genai.Client
	Conf          *GeminiConfig
	Ctx           context.Context
	cachedContent *genai.CachedContent
}

// TODO: Add tools
func (client *AIClient) createCacheContent() error {
	cfg := &genai.CreateCachedContentConfig{
		TTL:               3600,
		DisplayName:       "geminicode_cache",
		SystemInstruction: CachedSystemPrompt(),
	}

	cachedContent, err := client.C.Caches.Create(client.Ctx, client.Conf.Model, cfg)
	if err != nil {
		return err
	}
	client.cachedContent = cachedContent
	return nil
}

func (client *AIClient) Generate(textContent string) (*genai.GenerateContentResponse, error) {
	var generateConfig *genai.GenerateContentConfig

	if client.cachedContent == nil {
		err := client.createCacheContent()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(client.cachedContent.Name)
	generateConfig = client.Conf.CreateGenerateContentConfig(client.cachedContent.Name)

	content := &genai.Content{
		Parts: []*genai.Part{
			genai.NewPartFromText(textContent),
		},
		Role: "model",
	}

	result, err := client.C.Models.GenerateContent(client.Ctx, client.Conf.Model, []*genai.Content{content}, generateConfig)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateClient(ctx context.Context, conf *GeminiConfig) (*AIClient, error) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  conf.APIKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, fmt.Errorf("Error creating client: %w", err)
	}

	aiClient := &AIClient{
		C:    client,
		Conf: conf,
		Ctx:  ctx,
	}
	return aiClient, nil

}
