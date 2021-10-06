package handler

import (
	"net/http"

	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/logic"
	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/svc"
	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
