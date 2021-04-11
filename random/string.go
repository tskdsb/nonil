package random

import (
	"github.com/google/uuid"
)

func randomString() string {
	return uuid.NewString()
}
