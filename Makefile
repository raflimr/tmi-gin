migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://postgres:12345678@localhost:5432/tmi-gin?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:12345678@localhost:5432/tmi-gin?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:12345678@localhost:5432/tmi-gin?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:12345678@localhost:5432/tmi-gin?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v F:\!RyuEXP\Golang\project\tmi-gin:/src -w /src kjconroy/sqlc generate

migratecreate:
	migrate create -ext sql -dir db/migration -seq #<naming for your migration>

