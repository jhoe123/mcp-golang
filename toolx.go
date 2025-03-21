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
	Name        string        `json:"name"`
	Desc        string        `json:"desc"`
	InputSchema *PromptSchema `json:"params"`
}
