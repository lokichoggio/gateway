package errorx

import (
	"context"

	"github.com/lokichoggio/gateway/internal/types"

	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
)

const (
	SuccessCode = 0

	// 需要监控的业务状态码
	ServerErrorCode = 10000

	// 不需要监控的状态码
	ParamErrorCode = 10100
)

var errMap = map[int]string{
	SuccessCode:     "成功",
	ServerErrorCode: "服务器内部错误",
	ParamErrorCode:  "参数错误",
}

type CodeError struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type CodeErrorResponse struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func NewCodeError(ctx context.Context, code int, err error) error {
	err = errors.WithStack(err)

	logx.WithContext(ctx).Errorf("business_code: %d, err_detail: %+v", code, err)

	return &CodeError{ErrCode: code, ErrMsg: errMap[code]}
}

func (e *CodeError) Error() string {
	return e.ErrMsg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		ErrCode: e.ErrCode,
		ErrMsg:  e.ErrMsg,
	}
}

func SuccessResp() types.BaseResp {
	return types.BaseResp{
		ErrCode: SuccessCode,
		ErrMsg:  errMap[SuccessCode],
	}
}
