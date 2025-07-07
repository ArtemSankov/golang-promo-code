package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ArtemSankov/golang-promo-code/internal/api/http/promocodes"
	"github.com/ArtemSankov/golang-promo-code/internal/api/http/validator"
	promoCodeRepo "github.com/ArtemSankov/golang-promo-code/internal/repository/promocode"
	promoCodeService "github.com/ArtemSankov/golang-promo-code/internal/service/promocode"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("env DATABASE_URL is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer pool.Close()

	e := echo.New()
	e.Validator = validator.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	promoCodeRepo := promoCodeRepo.New(pool)
	promoCodeService := promoCodeService.NewService(promoCodeRepo)
	promoCodeHandler := promocodes.NewHandler(promoCodeService)

	promoCodes := e.Group("/promocodes")
	promoCodes.POST("/", promoCodeHandler.CreatePromoCodeHandler)
	promoCodes.GET("/:code", promoCodeHandler.GetPromoCodeByCode)

	addr := ":3000"

	log.Printf("Starting server on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}

}
