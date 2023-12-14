package service

import (
	"common/proto/proto"
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"lobbyServer/helper"
	"lobbyServer/model/accountModel"
	"regexp"
)

type AccountService struct {
	component.Base
	app pitaya.Pitaya
}

func NewAccountService(app pitaya.Pitaya) *AccountService {
	return &AccountService{
		app: app,
	}
}

// 长度限制 4 - 10
// 合法字符限制 a-z A-Z 0-9
func checkNameValid(name string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]{4,10}$")
	return re.MatchString(name)
}

func (s *AccountService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.CommonResponse, error) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)

	// check params
	if req.Account == "" || req.Password == "" {
		logger.Error("Account or password is empty!")
		return &proto.CommonResponse{Err: proto.ErrCode_UpParam}, nil
	}

	// 合法性
	if !checkNameValid(req.Account) {
		return &proto.CommonResponse{Err: proto.ErrCode_AccountRegister_NameInvalid}, nil
	}

	// 重复性
	if accountModel.Exist(req.Account) {
		return &proto.CommonResponse{Err: proto.ErrCode_AccountRegister_NameExist}, nil
	}

	uid := helper.GenerateUid()
	if uid == "" {
		logger.Errorf("Failed to generate uid!")
		return &proto.CommonResponse{Err: proto.ErrCode_ERR}, nil
	}

	// 注册
	model := &accountModel.AccountModel{
		Name:     req.Account,
		Password: req.Password,
		Uid:      uid,
	}

	err := model.New()
	if err != nil {
		logger.Errorf("Failed to create account! name: %s", req.Account)
		return &proto.CommonResponse{Err: proto.ErrCode_ERR}, nil
	}

	return &proto.CommonResponse{Err: proto.ErrCode_OK}, nil
}
