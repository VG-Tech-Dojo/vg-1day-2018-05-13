.DEFAULT_GOAL := help

GOOS = darwin
GOARCH = amd64
VERSION = $(shell git rev-parse HEAD)
ENV = development
host=localhost:8080

help:
	@cat Makefile

deps:
	which dep || go get -u github.com/golang/dep/...
	which sql-migrate || go get -u github.com/rubenv/sql-migrate/...
	dep ensure

run: fmt vet env/env.go dev.db
	go run server.go

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X=main.version=$(VERSION)" server.go

fmt:
	go fmt $$(go list ./... | grep -v vendor)

vet:
	go vet $$(go list ./... | grep -v vendor)

clean:
	rm -rf vendor
	rm dev.db

test: test_init test.db
	GIN_MODE=test go test -v

test_init:
	@rm -f test.db

test.db:
	@cp -i _etc/seed.db $@

## Migrate db schema
migrate_up:
	sql-migrate up -env=$(ENV)

## Migrate db schema(dryrun)
migrate_dryrun:
	sql-migrate up -env=$(ENV) -dryrun

## Show migration status
migrate_status:
	sql-migrate status -env=$(ENV)

curl_messages_ping:
	curl -i $(host)/api/ping

curl_messages_get_all:
	curl -i $(host)/api/messages

curl_messages_get:
	curl -i $(host)/api/messages/$(id)

curl_message_post:
	curl -i -X POST $(host)/api/messages -d '{"body": "$(body)"}'

curl_message_put:
	curl -i -X PUT $(host)/api/messages/$(id) -d '{"body": "$(body)"}'

curl_message_delete:
	curl -i -X DELETE $(host)/api/messages/$(id)

dev.db:
	$(MAKE) db_init

## Initialize database
db_init:
	cp -i _etc/seed.db dev.db

## Initialize environment configuration file
env/env.go:
	cp env/env.go.tmpl env/env.go

## Build godoc server(http://localhost:6060)
godoc_server:
	godoc -http=:6060
