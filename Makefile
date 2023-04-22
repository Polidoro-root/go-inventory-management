.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: restart
restart:
	docker-compose restart

.PHONY: test
test:
	go test -cover -timeout 30s ./...