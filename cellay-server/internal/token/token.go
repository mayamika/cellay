package token

import "github.com/google/uuid"

const TokenLength = 8

func New() string {
	return uuid.New().String()[:TokenLength]
}
