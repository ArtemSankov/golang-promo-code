include .env
export

# Путь к директории миграций
MIGRATIONS_DIR=./migrations

MIGRATE_BIN := $(shell which migrate)

ifndef MIGRATE_BIN
$(error "migrate not found in PATH.")
endif

# DSN к PostgreSQL (можно вынести в .env и считывать через shell)
DB_URL=$(DATABASE_URL)

# Путь к бинарнику migrate (если не в $PATH — пропиши вручную)
MIGRATE_BIN=migrate

# Создать новую миграцию: make migrate-create name=create_table
migrate-create:
	@read -p "Migration name: " name && \
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_DIR) $$name

# Применить все миграции
migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Откатить на один шаг
migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Откатить все миграции
migrate-down-all:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Посмотреть текущую версию миграции
migrate-version:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

# Повторно применить последнюю миграцию (down + up)
migrate-redo:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up 1
