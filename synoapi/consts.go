package synoapi

import (
	"fmt"
)

type EncryptionStatus int8

const (
	NOT_ENCRYPTED      = 0
	ENCRYPTED_LOCKED   = 1
	ENCRYPTED_UNLOCKED = 2
)

func (e EncryptionStatus) String() string {
	switch e {
	default:
		return fmt.Sprintf("Unknown: %v", e)
	case NOT_ENCRYPTED:
		return "Not encrypted"
	case ENCRYPTED_LOCKED:
		return "Encrypted, locked"
	case ENCRYPTED_UNLOCKED:
		return "Encrypted, unlocked"
	}
}
