package httpserver

import (
	"fmt"
	"net/http"

	go_spec_greet "github.com/Xenokrat/go-specs-greet/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, go_spec_greet.Greet(name))
}
