package tools

import (
	"google.golang.org/genai"
)

type ServerTool struct {
	t       *genai.Tool
	name    string
	handler func(args map[string]any)

	// Will the tool only perform reading / GET operations
	readOnly bool
}

func CreateTool(name string, fd *genai.FunctionDeclaration, handler func(args map[string]any), readOnly bool) *ServerTool {
	tool := &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{fd},
	}

	serverTool := &ServerTool{
		t:        tool,
		name:     name,
		handler:  handler,
		readOnly: readOnly,
	}
	return serverTool
}
