package models

import (
	"context"
	"time"
)

type BaseModel struct {
	CreatedAt *time.Time `json:"created_at,omitempty" sql:"default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" sql:"default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// BeforeInsert - update createdAt and updatedAt
func (m *BaseModel) BeforeInsert(ctx context.Context) error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = &now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = &now
	}
	return nil
}

// BeforeUpdate - update updatedAt
func (m *BaseModel) BeforeUpdate(c context.Context) (context.Context, error) {
	now := time.Now()
	m.UpdatedAt = &now
	return c, nil
}
