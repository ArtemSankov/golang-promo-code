package promocode

import (
	"time"

	"github.com/google/uuid"
)

type Promocode struct {
	ID               uuid.UUID `json:"id"`
	Code             string    `json:"code"`
	DiscountType     string    `json:"discount_type"`
	DiscountValue    int32     `json:"discount_value"`
	MaxActivations   int32     `json:"max_activations"`
	ActivationsCount int32     `json:"activations_count"`
	ExpiresAt        time.Time `json:"expires_at"`
	CreatedAt        time.Time `json:"created_at"`
}