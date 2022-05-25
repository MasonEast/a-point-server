package logic

import (
	"context"
	"fmt"

	"a-point-server/core/internal/svc"
	"a-point-server/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.SendRequest) (resp *types.SendReponse, err error) {
	fmt.Println(req.Data, "------")
	return
}
