package User

import (
	"github.com/satori/go.uuid"
	"time"
)

type Session struct {
	SessionID string

	UserID uuid.UUID

	LoginTime time.Time

	LastSeenTime time.Time
}
