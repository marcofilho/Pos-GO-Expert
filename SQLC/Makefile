create-migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose up

migrate-down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose down

.PHONY: migrate migrate-down create-migration