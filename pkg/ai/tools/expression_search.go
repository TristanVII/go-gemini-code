package tools

import (
	"fmt"
	"gemini-cli/pkg/utils"
	"os"
	"os/exec"

	"google.golang.org/genai"
)

const ExpressionSearchsName = "expression_search"
const ExpressionSearchsDescription = `
	Returns absolute file paths of files where the expression was found.
	`

func createExpressionSearchsSchema() *genai.FunctionDeclaration {
	params := &genai.Schema{Type: genai.TypeObject, Properties: make(map[string]*genai.Schema)}

	fd := &genai.FunctionDeclaration{
		Name:        ExpressionSearchsName,
		Description: ExpressionSearchsDescription,
		Parameters:  params,
	}
	return fd
}

func ExpressionSearchsHandler(args map[string]any) string {
	pwd, err := os.Getwd()
	expression, err := utils.GetArg[string](args, "expression")
	if err != nil {
		return "Missing required argument: expression"
	}
	isRegex, err := utils.GetArg[bool](args, "is_regex")

	if err != nil {
		return "Missing required argument: is_regex"
	}

	cmdArgs := []string{
		"--color=never",
		"--files-with-matches",
	}
	if !isRegex {
		cmdArgs = append(cmdArgs, "--fixed-strings")
	}

	cmdArgs = append(cmdArgs, expression, pwd)
	cmd := exec.Command("rg", cmdArgs...)

	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error running command: %w", err).Error()
	}

	return string(out)
}

// func ReadFileTool() *ServerTool {
// 	return CreateTool(Name, createFd())
// }
