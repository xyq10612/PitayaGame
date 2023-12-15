package router

import (
	"common/constants"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
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
