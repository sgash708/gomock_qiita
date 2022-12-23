.PHONY: add-network
add-network:
	docker network create gomock_backend_network

.PHONY: up
up:
	docker-compose up

.PHONY: upd
upd:
	docker-compose up -d

.PHONY: upnd
upnd:
	make add-network
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: exec
exec:
	docker-compose exec gomock_backend sh

.PHONY: build
build:
	docker-compose build --no-cache

.PHONY: build-up
build-up:
	docker-compose up -d --remove-orphans

.PHONY: test
test:
	docker-compose exec -T gomock_backend sh ./scripts/go_test.sh

.PHONY: test-local
test-local:
	docker-compose exec gomock_backend go test -v server/...

.PHONY: migrate-up
migrate-up:
	docker-compose exec -T gomock_backend sh ./scripts/migrate-up.sh

.PHONY: migrate-down
migrate-down:
	docker-compose exec -T gomock_backend sh ./scripts/migrate-down.sh

.PHONY: mockgen
mockgen:
	docker-compose exec gomock_backend mockgen -source=api/$(Path)/$(FileName) -destination=./mock/$(Path)/mock_$(FileName)

.PHONY: create-mig-local
create-mig-local:
	docker-compose exec -T gomock_backend sh -c "cd ../db/migrations && goose mysql 'mock-user:password@tcp(mock-mysql:3306)/mock-db?charset=utf8mb4&parseTime=true' create $(FileName) sql"

.PHONY: img-prune
img-prune:
	docker volume ls | xargs docker volume rm | 2>/dev/null
	docker image prune -a
	docker builder prune
