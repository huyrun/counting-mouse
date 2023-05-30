up:
	@docker compose -f docker/docker-compose.yml up --build -d

test:
	@echo "==> Running tests..."
	@go clean -testcache ./...
	@go test -vet=off -tags="musl $(TAGS)" `go list ./... | grep -v cmd` -p 1 -race --cover -coverprofile=output/cover.out

test-up:
	@COMPOSE_HTTP_TIMEOUT=180 docker-compose \
		-f docker/docker-compose.test.yml up \
		--force-recreate \
		--abort-on-container-exit \
		--exit-code-from app \
		--build

test-down:
	@COMPOSE_HTTP_TIMEOUT=180 docker-compose \
		-f docker/docker-compose.test.yml down \
 		-v --rmi local