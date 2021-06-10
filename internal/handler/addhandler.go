package handler

import (
	"net/http"

	"github.com/lokichoggio/gateway/internal/common/errorx"
	"github.com/lokichoggio/gateway/internal/logic"
	"github.com/lokichoggio/gateway/internal/svc"
	"github.com/lokichoggio/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewCodeError(r.Context(), errorx.ParamErrorCode, err))
			return
		}

		l := logic.NewAddLogic(r.Context(), ctx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
