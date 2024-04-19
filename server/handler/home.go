package handler

import (
	"net/http"

	"github.com/broemp/growBro/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, home.Index())
}
