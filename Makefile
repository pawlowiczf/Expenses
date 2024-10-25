migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/expenses?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/expenses?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v ./... 
go:
	go run main.go

