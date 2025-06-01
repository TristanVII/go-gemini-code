package tools

import (
	"fmt"
	"github.com/denormal/go-gitignore"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/genai"
)

const ListFilesName = "read_file"
const ListFilesDescription = `
	Write content to the file path provided.
	Returns: Success or error
	`

func createListFilesSchema() *genai.FunctionDeclaration {
	params := &genai.Schema{Type: genai.TypeObject, Properties: make(map[string]*genai.Schema)}

	fd := &genai.FunctionDeclaration{
		Name:        ListFilesName,
		Description: ListFilesDescription,
		Parameters:  params,
	}
	return fd
}

func isIgnored(file string, ignore gitignore.GitIgnore) bool {
	match := ignore.Match(file)
	if match != nil {
		if match.Ignore() {
			return true
		}
	}
	return false
}

func ListFilesHandler(args map[string]any) string {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting current working directory: %w", err).Error()
	}
	var gitItnoreFilePath = pwd + "/.gitignore"

	cmd := exec.Command("find", pwd, "-path", pwd+"/.git", "-prune", "-o", "-type", "f", "-print")
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error running command: %w", err).Error()
	}

	ignore, err := gitignore.NewFromFile(gitItnoreFilePath)
	if err != nil {
		return fmt.Errorf("Error listing files: %w", err).Error()
	}

	paths := strings.Split(string(out), "\n")

	results := ""
	for _, path := range paths {
		trimmedPath := strings.TrimSpace(path)
		if trimmedPath == "" {
			continue
		}
		if !isIgnored(trimmedPath, ignore) {
			results += path + "\n"
		}
	}
	return results
}

// func ReadFileTool() *ServerTool {
// 	return CreateTool(Name, createFd())
// }
