package server

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"

	"atomys.codes/vanity-go/pkg/repository"
)

// Page is the data structure that is passed to the template
type Page struct {
	Repositories []*repository.Repository
}

var cachedTemplate *template.Template

// proxyHandler is the handler for the proxy
// It is the main entry point for the proxy
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	var page = Page{
		Repositories: RepositoriesForNamespace(r.Host),
	}

	if err := cachedTemplate.ExecuteTemplate(w, "index.html", page); err != nil {
		log.Error().Err(err).Msg("error executing template")
	}
}
