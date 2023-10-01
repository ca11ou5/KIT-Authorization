proto:
	protoc ./internal/pb/auth.proto --go_out=plugins=grpc:.

build:
	docker build -t authorization .

compose:
	docker compose -f docker-compose.yml up