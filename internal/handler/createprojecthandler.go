package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"projects/api/internal/logic"
	"projects/api/internal/svc"
	"projects/api/internal/types"
)

func CreateProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateProjectLogic(r.Context(), svcCtx)
		resp, err := l.CreateProject(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
