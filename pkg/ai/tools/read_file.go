package tools

import (
	"fmt"
	"os"

	"gemini-cli/pkg/utils"
	"google.golang.org/genai"
)

const ReadFileName = "read_file"
const ReadFileDescription = `
	Read the contents of a file.
	Returns: Contents of the file or error
	`

func createReadFileSchema() *genai.FunctionDeclaration {

	params := &genai.Schema{Type: genai.TypeObject, Properties: make(map[string]*genai.Schema), Required: []string{"absolute_file_path"}}

	params.Properties["absolute_file_path"] = &genai.Schema{Type: genai.TypeString, Description: "The absolute file path on the users computer that will be read."}

	fd := &genai.FunctionDeclaration{
		Name:        ReadFileName,
		Description: ReadFileDescription,
		Parameters:  params,
	}
	return fd
}

func ReadFileHandler(args map[string]any) string {
	fp, err := utils.GetArg[string](args, "absolute_file_path")
	if err != nil {
		return err.Error()
	}
	content, err := os.ReadFile(fp)
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}
	return string(content)
}

// func ReadFileTool() *ServerTool {
// 	return CreateTool(Name, createReadFileSchema())
// }
