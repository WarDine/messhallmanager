package usecases

import "domain"

type MessHall struct {
}

// Enforce interface
var _ domain.MessHallInterface = (*MessHall)(nil)
