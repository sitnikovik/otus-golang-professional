package event

import (
	"errors"
	"time"
)

const (
	// DefaultNotifyBefore is the default time before the event to notify
	DefaultNotifyBefore = 5 * time.Minute
)

var (
	// ErrNotFound is used when the requested event is not found
	ErrNotFound = errors.New("event not found")
)

// Event defines the event entity
type Event struct {
	// ID Unique identifier
	ID uint64 `json:"id" db:"id"`
	// Title Event title
	Title string `json:"title" db:"title"`
	// CreatedAt Creation time
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	// FinishedAt Event end time
	FinishedAt *time.Time `json:"finishedAt" db:"finished_at"`
	// Description Event description
	Description string `json:"description" db:"description"`
	// OwnerID Event owner ID
	OwnerID uint64 `json:"ownerId" db:"owner_id"`
	// NotifyBefore Notify before time
	NotifyBefore time.Duration `json:"notifyBefore" db:"notify_before"`
}
