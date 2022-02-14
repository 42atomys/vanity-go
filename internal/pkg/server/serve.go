package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Serve the proxy server
func Serve(port *int) error {
	if err := loadConfig(); err != nil {
		return err
	}

	cachedTemplate = template.Must(template.ParseFiles("web/template/index.html"))

	http.HandleFunc("/", proxyHandler)

	log.Info().Msgf("Listening on port %d", *port)
	return http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
