package logic

import (
	"context"

	"github.com/lokichoggio/add/adder"
	"github.com/lokichoggio/gateway/internal/common/errorx"
	"github.com/lokichoggio/gateway/internal/svc"
	"github.com/lokichoggio/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})

	if err != nil {
		// l.Logger.Error(err)
		return nil, errorx.NewCodeError(l.ctx, errorx.ServerErrorCode, err)
	}

	return &types.AddResp{
		BaseResp: errorx.SuccessResp(),
		Data: types.AddData{
			Ok: resp.Ok,
		},
	}, nil
	// 手动代码结束
}
