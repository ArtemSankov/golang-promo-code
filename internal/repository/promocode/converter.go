package promocode

import (
	db "github.com/ArtemSankov/golang-promo-code/internal/db/sqlc"
	domain "github.com/ArtemSankov/golang-promo-code/internal/domain/promocode"
)

func fromDBModel(pm db.Promocode) domain.Promocode {
	return domain.Promocode(pm)
}