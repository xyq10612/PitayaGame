package router

import (
	"common/constants"
	"context"
	"errors"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/route"
	"math/rand"
)

func GetRandomLobby() *cluster.Server {
	m, err := pitaya.GetServersByType(constants.LobbyServer)
	if err != nil || len(m) == 0 {
		return nil
	}

	v := rand.Intn(len(m))
	i := 0
	for _, s := range m {
		if i == v {
			return s
		}
		i++
	}
	return nil
}

func IsLobbyAlive(lobbyId string) bool {
	m, err := pitaya.GetServersByType(constants.LobbyServer)
	if err != nil || len(m) == 0 {
		return false
	}
	_, ok := m[lobbyId]
	return ok
}

func LobbyRouterFunc(
	ctx context.Context,
	route *route.Route,
	payload []byte,
	servers map[string]*cluster.Server,
) (*cluster.Server, error) {
	// 转发到绑定的 lobby
	session := pitaya.GetSessionFromCtx(ctx)
	if session == nil || !session.HasKey(constants.SessionLobbyIdKey) {
		return nil, errors.New("not find binding lobby in session")
	}

	lobbyId := session.Get(constants.SessionLobbyIdKey).(string)
	s, ok := servers[lobbyId]
	if !ok {
		return nil, errors.New("bind lobby is not exist")
	}

	return s, nil
}
