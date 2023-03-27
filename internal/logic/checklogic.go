package logic

import (
	"context"

	"github.com/jarcn/bookcheck/check"
	"github.com/jarcn/bookcheck/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)
	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
