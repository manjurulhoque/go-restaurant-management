package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `json:"name" validate:"required"`
	Category  string             `json:"category" validate:"required"`
	StartDate *time.Time         `bson:"start_date" json:"start_date"`
	EndDate   *time.Time         `bson:"end_date" json:"end_date"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	MenuId    string             `bson:"menu_id" json:"menu_id"`
}
