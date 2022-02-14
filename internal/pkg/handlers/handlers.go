package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"atomys.codes/go-proxy/pkg/repository"
	"github.com/rs/zerolog/log"
)

var tmp = template.Must(template.ParseFiles("web/template/index.html"))

type Page struct {
	Repositories []repository.Repository
}

func GoProxyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, world!")

	var page = Page{
		Repositories: []repository.Repository{},
	}

	rep, err := repository.New("atomys.codes/go-proxy", "https://github.com/42atomys/go-proxy.git")
	if err != nil {
		log.Error().Err(err).Msg("error creating repository")
		return
	}

	page.Repositories = append(page.Repositories, *rep)

	if err := tmp.ExecuteTemplate(w, "index.html", page); err != nil {
		log.Error().Err(err).Msg("error executing template")
	}
}
