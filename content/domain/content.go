package content

import "time"

type Content struct {
	ID         string
	Writer     string
	WriteScope Scope
	ViewScope  Scope
	CreatedAt  time.Time
}

type Scope string

const (
	Public  Scope = "public"
	Friend  Scope = "friend"
	Private Scope = "private"
)
