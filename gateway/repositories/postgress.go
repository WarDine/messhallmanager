package repositories

import "domain"

type PostgressManager struct {
}

// Enforce interface
var _ domain.PostgressManagerInterface = (*PostgressManager)(nil)
