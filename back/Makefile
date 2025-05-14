.PHONY: migrate-create migrate-up migrate-down migrate-force sqlboiler-generate

# 環境変数の読み込み
include .env
export

# マイグレーションコマンド
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./migrations -seq $$name; \
	echo "Migration created: $$name"

migrate-up:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations up

migrate-down:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations down

migrate-force:
	@read -p "Enter version number: " version; \
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./migrations force $$version

sqlboiler-generate:
	export MYSQL_HOST=localhost && \
	export MYSQL_USER=$(MYSQL_USER) && \
	export MYSQL_PASSWORD=$(MYSQL_PASSWORD) && \
	export MYSQL_PORT=$(MYSQL_PORT) && \
	export MYSQL_DATABASE=$(MYSQL_DATABASE) && \
	sqlboiler mysql -c ./infra/database/sqlboiler/config.toml -o ./infra/database/sqlboiler/models -p models --no-tests --wipe

# ヘルプコマンド
help:
	@echo "Available commands:"
	@echo "  make migrate-create    - 新しいマイグレーションファイルを作成"
	@echo "  make migrate-up        - 保留中のマイグレーションを実行"
	@echo "  make migrate-down      - 最後に実行したマイグレーションをロールバック"
	@echo "  make migrate-force     - 特定のバージョンに強制的にマイグレーション"
	@echo "  make sqlboiler-generate - sqlboilerを生成"