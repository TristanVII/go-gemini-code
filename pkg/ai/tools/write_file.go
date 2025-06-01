package tools

import (
	"bufio"
	"os"

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
	fp, ok := args["absolute_file_path"]
	if !ok {
		return "Args did not contain absolute_file_path"
	}
	filePath, ok := fp.(string)
	if !ok {
		return "absolute_file_path is not string"
	}

	content, ok := args["content"]
	if !ok {
		return "Args did not contain content"
	}

	contents, ok := content.(string)
	if !ok {
		return "absolute_file_path is not string"
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "Error reading file. Are you sure that is the correct FULL ABSOLUTE PATH?"
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString(contents)
	if err != nil {
		return "Error writing to file"
	}

	return "Succesfully written to file " + filePath
}

// func ReadFileTool() *ServerTool {
// 	return CreateTool(Name, createFd())
// }
