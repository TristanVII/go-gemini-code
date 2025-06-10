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
func (client *AIClient) createCacheContentConfig() error {
	// Default TTL 1hour
	cfg := &genai.CreateCachedContentConfig{
		SystemInstruction: genai.NewContentFromText(SYSTEM_PROMPT, genai.RoleUser),
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
		err := client.createCacheContentConfig()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(client.cachedContent.Name)
	generateConfig = client.Conf.CreateGenerateContentConfig(client.cachedContent.Name)

	result, err := client.C.Models.GenerateContent(client.Ctx, client.Conf.Model, genai.Text(textContent), generateConfig)
	if err != nil {
		return nil, err
	}

	var thinking = false
	for _, part := range result.Candidates[0].Content.Parts {
		if part.Thought {
			thinking = true
		} else {
			thinking = false
		}
		// text part
		if part.Text != "" {
			prefix := ""
			if !thinking {
				prefix = "Thinking: "
			}
			fmt.Printf("%s: %s\n", prefix, part.Text)
		}

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
