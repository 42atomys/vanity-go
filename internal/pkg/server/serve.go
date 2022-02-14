//go:build !skip

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
func Serve(port int) error {
	fsys, err := fs.Sub(embededFiles, "templates/index.html")
	if err != nil {
		panic(err)
	}

	if !validPort(port) {
		return fmt.Errorf("invalid port")
	}

	cachedTemplate = template.Must(template.ParseFS(fsys))
	http.HandleFunc("/", proxyHandler)
	log.Info().Msgf("Listening on port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// validPort returns true if the port is valid
// following the RFC https://datatracker.ietf.org/doc/html/rfc6056#section-2.1
func validPort(port int) bool {
	return port > 0 && port < 65535
}
