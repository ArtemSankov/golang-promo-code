package promocodes

import (
	"net/http"

	promoCodeService "github.com/ArtemSankov/golang-promo-code/internal/service/promocode"
	"github.com/labstack/echo/v4"
)

type createPromoCodeRequest struct {
	Code           string `json:"code" validate:"required"`
	DiscountType   string `json:"discount_type" validate:"required,oneof=percent fixed_amount"`
	DiscountValue  int    `json:"discount_value" validate:"required,min=1"`
	MaxActivations int    `json:"max_activations" validate:"required,min=1"`
	ExpiresAt      string `json:"expires_at" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type createPromoCodeResponse struct {
	ID string `json:"id"`
}

type Handler struct {
	service promoCodeService.Service
}

func NewHandler(service promoCodeService.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreatePromoCodeHandler(c echo.Context) error {
	var req createPromoCodeRequest
	
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

 	id, err :=	h.service.Create(c.Request().Context(), promoCodeService.CreatePromoCodeInput{
		Code:           req.Code,
		DiscountType:   req.DiscountType,
		DiscountValue:  req.DiscountValue,
		MaxActivations: req.MaxActivations,
		ExpiresAt:      req.ExpiresAt,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusCreated, createPromoCodeResponse{ID: id})
}

func (h *Handler) GetPromoCodeByCode(c echo.Context) error {
	code := c.Param("code")

	pm, err := h.service.GetByCode(c.Request().Context(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, pm)
}
