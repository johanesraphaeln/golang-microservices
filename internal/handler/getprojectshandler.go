package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"projects/api/internal/logic"
	"projects/api/internal/svc"
	"projects/api/internal/types"
)

func GetProjectsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetProjectsLogic(r.Context(), svcCtx)
		resp, err := l.GetProjects(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
