package loginModel

import (
	"common/helper"
	"common/util"
	"context"
	"time"
)

const keyPrefix = "login"

func getRedisKey(username string) string {
	return util.JoinKey(keyPrefix, username)
}

type LoginModel struct {
	Account   string `redis:"account"`
	LobbyId   string `redis:"lobbyId"`
	Timestamp int64  `redis:"timestamp"`
}

func (m *LoginModel) GetRedisKey() string {
	return getRedisKey(m.Account)
}

func (m *LoginModel) Save() error {
	return helper.GetRedis().HSet(context.TODO(), m.GetRedisKey(), m).Err()
}

func Get(account string) (LoginModel, error) {
	var model LoginModel

	cmd := helper.GetRedis().HGetAll(context.TODO(), getRedisKey(account))
	if cmd.Err() != nil {
		return model, cmd.Err()
	}

	err := cmd.Scan(&model)
	if err != nil {
		return model, err
	}
	return model, nil
}

func Save(account, lobbyId string) error {
	m := LoginModel{Account: account, LobbyId: lobbyId, Timestamp: time.Now().Unix()}
	return m.Save()
}
