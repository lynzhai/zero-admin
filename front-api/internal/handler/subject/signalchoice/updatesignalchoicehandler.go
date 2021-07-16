package handler

import (
	"net/http"

	"go-zero-admin/front-api/internal/logic/subject/signalchoice"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateSignalChoiceHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSignalChoiceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateSignalChoiceLogic(r.Context(), ctx)
		resp, err := l.UpdateSignalChoice(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
