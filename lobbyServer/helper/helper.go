package helper

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUid() string {
	id := uuid.New()
	uuid := strings.Replace(id.String(), "-", "", -1)
	return uuid
}
