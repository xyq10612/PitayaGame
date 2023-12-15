package service

import (
	"common/constants"
	"common/proto/proto"
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"proxyServer/model/loginModel"
	"proxyServer/router"
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
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	session := s.app.GetSessionFromCtx(ctx)

	rsp := &proto.CommonResponse{Err: proto.ErrCode_ERR}

	lobby := router.GetRandomLobby()
	if lobby == nil {
		logger.Errorf("cannot find random lobby!")
		return rsp, nil
	}

	err := s.app.RPCTo(ctx, lobby.ID, "lobby.account.register", rsp, req)
	if err != nil {
		return nil, err
	}

	_ = session.Set(constants.SessionLobbyIdKey, lobby.ID)

	rsp.Err = proto.ErrCode_OK
	return rsp, nil
}

func (s *AccountService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	session := s.app.GetSessionFromCtx(ctx)

	rsp := &proto.LoginResponse{Ret: proto.ErrCode_ERR}

	var lobbyId string

	// 1. 从session获取上次的lobby, 找不到就从redis缓存获取上次登录的lobby
	if session.HasKey(constants.SessionLobbyIdKey) {
		lobbyId = session.Get(constants.SessionLobbyIdKey).(string)
		logger.Infof("get lobby from session, account: %s", req.Account)
	}
	if lobbyId == "" {
		loginModel, err := loginModel.Get(req.Account)
		if err != nil {
			return rsp, err
		}
		lobbyId = loginModel.LobbyId
		logger.Infof("get lobby from cache, account: %s", req.Account)
	}

	// 2. 没有上次登录的, 或者上次登录的lobby已经不存在了, 随机分配一个lobby
	if lobbyId == "" || !router.IsLobbyAlive(lobbyId) {
		lobby := router.GetRandomLobby()
		if lobby == nil {
			logger.Errorf("cannot find random lobby!")
			return rsp, nil
		}
		lobbyId = lobby.ID
		logger.Infof("get lobby from random, account: %s", req.Account)
	}

	// 3. 转发登录到lobby, 由lobby处理登录逻辑, 初始化玩家数据等
	err := s.app.RPCTo(ctx, lobbyId, "lobby.account.login", rsp, req)
	if err != nil {
		logger.Errorf("rpc to lobby err: %v", err.Error())
		rsp.Ret = proto.ErrCode_ERR
		return rsp, nil
	}

	logger = logger.WithField("userId", rsp.Uid).WithField("lobby", lobbyId)

	// 5. 更新redis缓存, 记录登录的lobby
	loginModel.Save(req.Account, lobbyId)

	// 6. 绑定session-lobby, session-uid
	session.Bind(ctx, rsp.Uid)
	session.Set(constants.SessionLobbyIdKey, lobbyId)

	logger.Infof("login success account: %s", req.Account)

	return rsp, nil
}
