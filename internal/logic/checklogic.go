package logic

import (
	"context"

	"github.com/lokichoggio/check/checker"
	"github.com/lokichoggio/gateway/internal/common/errorx"
	"github.com/lokichoggio/gateway/internal/svc"
	"github.com/lokichoggio/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: req.Book,
	})
	if err != nil {
		return nil, errorx.NewCodeError(l.ctx, errorx.ServerErrorCode, err)
	}

	return &types.CheckResp{
		BaseResp: errorx.SuccessResp(),
		Data: types.CheckData{
			Found: resp.Found,
			Price: resp.Price,
		},
	}, nil
	// 手动代码结束
}
