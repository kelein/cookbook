package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/swaggest/swgui/v5emb"

	"github.com/kelein/cookbook/govoyage/assets"
)

var (
	apiDocAddr = "0.0.0.0:8090"
	apiDocPath = "/api/v1/docs/"
	apiDocName = "Govoyage"
)

func init() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}

	logger := slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replace,
		},
	))
	slog.SetDefault(logger)
}

func main() {
	// * Register static OpenAPI yaml file
	http.Handle(assets.OpenAPIFilePath, http.FileServer(http.FS(assets.OpenAPIFile)))
	slog.Info("server openAPI file", "path", assets.OpenAPIFilePath)

	// * Register swagger UI and docs
	http.Handle(apiDocPath, v5emb.New(apiDocName, assets.OpenAPIFilePath, apiDocPath))
	slog.Info("server openAPI docs", "path", apiDocPath)

	// * Register index handler
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(index))
	})

	slog.Info("server start listen on", "addr", apiDocAddr)
	http.ListenAndServe(apiDocAddr, http.DefaultServeMux)
}

var index = `<!DOCTYPE html>
<html lang="en"><body>
<h3>Govoyage</h3>
<li><a href="/openapi.yaml">OpenAPI Yaml</a></li>
<li><a href="/api/v1/docs/">OpenAPI Docs<a></li>
</body></html>`
