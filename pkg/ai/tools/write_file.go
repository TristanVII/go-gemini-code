package tools

import (
	"os"

	"bufio"
	"gemini-cli/pkg/utils"
	"google.golang.org/genai"
)

const WriteFileName = "read_file"
const WriteFileDescription = `
	Write content to the file path provided.
	Returns: Success or error
	`

func createWriteFileSchema() *genai.FunctionDeclaration {
	params := &genai.Schema{Type: genai.TypeObject, Properties: make(map[string]*genai.Schema), Required: []string{"absolute_file_path", "content"}}

	params.Properties["absolute_file_path"] = &genai.Schema{Type: genai.TypeString, Description: "The absolute file path on the users computer that will be written to."}
	params.Properties["content"] = &genai.Schema{Type: genai.TypeString, Description: "File content to be written to the file. Appropriate formatting is expected"}

	fd := &genai.FunctionDeclaration{
		Name:        WriteFileName,
		Description: WriteFileDescription,
		Parameters:  params,
	}
	return fd
}

func WriteFileHandler(args map[string]any) string {
	fp, err := utils.GetArg[string](args, "absolute_file_path")
	if err != nil {
		return err.Error()
	}

	content, err := utils.GetArg[string](args, "content")
	if err != nil {
		return err.Error()
	}

	file, err := os.OpenFile(fp, 0x1, 0644)
	if err != nil {
		return "Error reading file. Are you sure that is the correct FULL ABSOLUTE PATH?"
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(content)
	if err != nil {
		return "Error writing to file"
	}
	err = w.Flush()
	if err != nil {
		return err.Error()
	}

	file.Close()

	return "Succesfully wrote to file " + fp
}

// func ReadFileTool() *ServerTool {
// 	return CreateTool(Name, createFd())
// }
