package logic

import (
	"context"

	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/svc"
	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
	return &types.Response{Message: req.Name}, nil
}
