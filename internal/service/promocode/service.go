package promocode

import (
	"context"
	"time"

	domain "github.com/ArtemSankov/golang-promo-code/internal/domain/promocode"
	"github.com/ArtemSankov/golang-promo-code/internal/repository/promocode"
)

type Service interface {
	Create(ctx context.Context, input CreatePromoCodeInput) (string, error)
	GetByCode(ctx context.Context, code string) (domain.Promocode, error)
}


type CreatePromoCodeInput struct {
	Code           string
	DiscountType   string
	DiscountValue  int
	MaxActivations int
	ExpiresAt      string
}

type service struct {
	repo promocode.Repository
}

func NewService(repo promocode.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, input CreatePromoCodeInput) (string, error) {

	expiresAt, err := time.Parse(time.RFC3339, input.ExpiresAt)
	if err != nil {
		return "", err
	}

	id, err := s.repo.Create(ctx, promocode.CreateInput{
		Code:           input.Code,
		DiscountType:   input.DiscountType,
		DiscountValue:  int32(input.DiscountValue),
		MaxActivations: int32(input.MaxActivations),
		ExpiresAt:      expiresAt,
	})

	// TOOD: handle error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *service) GetByCode(ctx context.Context, code string) (domain.Promocode, error) {
	pm, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return domain.Promocode{}, nil
	}

	return  pm, nil
}