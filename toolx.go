package mcp_golang

import (
	"github.com/invopop/jsonschema"
)

type ToolDef struct {
	Name        string             `json:"name"`
	Desc        string             `json:"desc"`
	InputSchema *jsonschema.Schema `json:"params"`
}

type PromptDef struct {
	Definition *PromptSchema `json:"definition"`
	Response   []PromptRes   `json:"response"`
}

type PromptRes struct {
	Content string `json:"content"`
	Role    string `json:"role`
}
