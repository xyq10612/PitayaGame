package service

import (
	"common/proto/proto"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
)

// AccountService 账号服务
type AccountService struct {
	component.Base
	app pitaya.Pitaya
}

func NewAccountService(app pitaya.Pitaya) *AccountService {
	return &AccountService{app: app}
}

func (s *AccountService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.CommonResponse, error) {
	logrus.Debugf("register request: %v", req)

	logrus.Debugf("do something in proxy")

	rsp := &proto.CommonResponse{}

	err := s.app.RPC(ctx, "lobby.account.register", rsp, req)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
