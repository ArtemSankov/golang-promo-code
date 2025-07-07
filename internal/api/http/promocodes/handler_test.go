package promocodes

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArtemSankov/golang-promo-code/internal/api/http/validator"
	domain "github.com/ArtemSankov/golang-promo-code/internal/domain/promocode"
	repo "github.com/ArtemSankov/golang-promo-code/internal/repository/promocode"
	service "github.com/ArtemSankov/golang-promo-code/internal/service/promocode"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {}

func (m *mockRepo) Create(ctx context.Context, input repo.CreateInput) (uuid.UUID, error) {
    return uuid.New(), nil
}

func (m *mockRepo) GetByCode(ctx context.Context, code string) (domain.Promocode, error) {
    return domain.Promocode{}, nil
}

func (m *mockRepo) IncrementActivations(ctx context.Context, id uuid.UUID) error {
    return nil
}

func TestCreatePromoCodeHandler_InvalidRequest(t *testing.T) {
	e := echo.New()
	e.Validator = validator.New()

	reqBody := `{ "code": "", "discount_type": "invalid_type" }`

	req := httptest.NewRequest(http.MethodPost, "/promocodes", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := service.NewService(&mockRepo{})
	handler := NewHandler(service)

	err :=  handler.CreatePromoCodeHandler(c)
	if assert.Error(t, err) {
		httpErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
	}
}


func TestCreatePromoCodeHandler(t *testing.T) {
	e := echo.New()
	e.Validator = validator.New()	 

	reqBody := `{
		"code": "SUMMER2025",
		"discount_type": "percent",
		"discount_value": 20,
		"max_activations": 100,
		"expires_at": "2025-12-31T23:59:59Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/promocodes", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	service := service.NewService(&mockRepo{})
	handler := NewHandler(service)

	err :=  handler.CreatePromoCodeHandler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `"id"`)
	}
}

func TestGetPromoCodeByCodeHandler(t *testing.T) {
	e := echo.New()
	e.Validator = validator.New()	 

	code := "SUMMER2025"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/promocodes/%s", code), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/promocodes/:code")
	c.SetParamNames("code")
	c.SetParamValues(code)
	
	service := service.NewService(&mockRepo{})
	handler := NewHandler(service)

	err := handler.GetPromoCodeByCode(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
