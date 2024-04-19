package handler

import (
	"net/http"

	"github.com/broemp/growBro/view/errorPages"
	"go.uber.org/zap"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) error {
	zap.L().Debug("not found", zap.String("path", r.URL.Path), zap.String("ip", r.RemoteAddr))
	return render(w, r, errorPages.NotFoundIndex())
}
