package assets

import (
	"embed"
)

// OpenAPIFile embed openapi.yaml
//
//go:embed openapi.yaml
var OpenAPIFile embed.FS

// OpenAPIFileName openAPI file name
var OpenAPIFileName = "openapi.yaml"

// OpenAPIFilePath openAPI file path
var OpenAPIFilePath = "/openapi.yaml"
