build:
	docker build -t hillel-auc-app .

up:
	docker run hillel-auc-app

migration-up:
	migrate -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' -path ./migrations up

migration-down:
	migrate -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' -path ./migrations down

.PHONY: migration-up migration-down
