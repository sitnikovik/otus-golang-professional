package event

import (
	"errors"
	"time"
)

const (
	// DefaultNotifyBefore is the default time before the event to notify.
	DefaultNotifyBefore = 5 * time.Minute
)

// ErrNotFound is used when the requested event is not found.
var ErrNotFound = errors.New("event not found")

// Event defines the event entity.
type Event struct {
	// ID Unique identifier.
	ID uint64 `json:"id" db:"id"`
	// Title Event title.
	Title string `json:"title" db:"title"`
	// CreatedAt Creation time.
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	// FinishedAt Event end time.
	FinishedAt *time.Time `json:"finishedAt" db:"finished_at"`
	// Description Event description.
	Description string `json:"description" db:"description"`
	// OwnerID Event owner ID.
	OwnerID uint64 `json:"ownerId" db:"owner_id"`
	// NotifyBefore Time before the event to notify.
	NotifyBefore time.Duration `json:"notifyBefore" db:"notify_before"`
	// IsNotified Is the event notified.
	IsNotified bool `json:"isNotified" db:"is_notified"`
}

// IsToNotify checks if the event is to notify.
func (e *Event) IsToNotify() bool {
	if e.IsNotified || e.FinishedAt == nil {
		return false
	}

	now := time.Now()
	finishedAt := *e.FinishedAt
	notifyTime := finishedAt.Add(-e.NotifyBefore)

	// Check if the event is within the notification window
	if now.After(notifyTime) && now.Before(finishedAt) {
		return true
	}

	// Check if the event has passed and not yet notified
	if now.After(finishedAt) {
		return true
	}

	return false
}
