migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=rating password=postgres sslmode=disable" up

last-migrate-down:
	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=rating password=postgres sslmode=disable" down

migrate-create:
	goose -dir db/migrations create $(F) sql

lint:
	golangci-lint run --fix ./...

gen:
	protoc --go_out=. --go-grpc_out=. .\api\client.proto
	protoc --go_out=. --go-grpc_out=. .\api\feedback.proto
	protoc --go_out=. --go-grpc_out=. .\api\product.proto
