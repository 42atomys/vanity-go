package server

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/rs/zerolog/log"
)

//go:embed templates/index.html
var embededFiles embed.FS

// Serve the proxy server
func Serve(port *int) error {
	if err := loadConfig(); err != nil {
		return err
	}

	fsys, err := fs.Sub(embededFiles, "templates/index.html")
	if err != nil {
		panic(err)
	}

	cachedTemplate = template.Must(template.ParseFS(fsys))

	http.HandleFunc("/", proxyHandler)

	log.Info().Msgf("Listening on port %d", *port)
	return http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
