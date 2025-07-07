package promocode

import (
	"context"
	"time"

	db "github.com/ArtemSankov/golang-promo-code/internal/db/sqlc"
	domain "github.com/ArtemSankov/golang-promo-code/internal/domain/promocode"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateInput struct {
	Code           string
	DiscountType   string
	DiscountValue  int32
	MaxActivations int32
	ExpiresAt      time.Time
}


type Repository interface {
	Create(ctx context.Context, input CreateInput ) (uuid.UUID, error)
	GetByCode(ctx context.Context, code string) (domain.Promocode, error)
	IncrementActivations(ctx context.Context, id uuid.UUID) error
}

type repo struct {
	q *db.Queries
}

func New(pool *pgxpool.Pool) Repository {
	return &repo{
		q: db.New(pool),
	}
}

func (r *repo) Create(ctx context.Context, input CreateInput) (uuid.UUID, error) {
	return r.q.CreatePromoCode(ctx, db.CreatePromoCodeParams{
		Code:           input.Code,
		DiscountType:   input.DiscountType,
		DiscountValue:  input.DiscountValue,
		MaxActivations: input.MaxActivations,
		ExpiresAt:      input.ExpiresAt,
	})
}

func (r *repo) GetByCode(ctx context.Context, code string) (domain.Promocode, error) {
	pm, err := r.q.GetPromoCodeByCode(ctx, code)
	if err != nil {
		return domain.Promocode{}, nil
	}

	return fromDBModel(pm), nil
}

func (r *repo) IncrementActivations(ctx context.Context, id uuid.UUID) error {
	return r.q.IncrementActivationsCount(ctx, id)
}