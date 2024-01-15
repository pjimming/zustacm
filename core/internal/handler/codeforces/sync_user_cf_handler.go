package codeforces

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/codeforces"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncUserCfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SyncUserCfReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := codeforces.NewSyncUserCfLogic(r.Context(), svcCtx)
		err := l.SyncUserCf(&req)

		httpresp.HTTP(w, r, nil, err)

	}
}
