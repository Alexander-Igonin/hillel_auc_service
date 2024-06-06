build:
	docker build -t hillel-auc-app .

up:
	docker run hillel-auc-app

migration-up:
	migrate -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' -path ./migrations up

.PHONY: migration-up
