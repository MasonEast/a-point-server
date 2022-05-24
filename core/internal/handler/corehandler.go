package handler

import (
	"net/http"

	"a-point-server/core/internal/logic"
	"a-point-server/core/internal/svc"
	"a-point-server/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCoreLogic(r.Context(), svcCtx)
		resp, err := l.Core(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
