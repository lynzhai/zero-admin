package handler

import (
	"net/http"

	"go-zero-admin/front-api/internal/logic/subject/subject"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func FindSubjectsByStudentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindSubjectsByStudentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFindSubjectsByStudentLogic(r.Context(), ctx)
		resp, err := l.FindSubjectsByStudent(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
