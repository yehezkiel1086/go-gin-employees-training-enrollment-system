composeup:
	docker compose up -d

composedown:
	docker compose down

postgres:
	docker exec -it postgres psql -U yehezkiel1086 -b employees_training

build:
	go build -o cmd/main.go bin/main

dev:
	go run cmd/main.go

run:
	./bin/main

.PHONY: composeup composedown postgres build dev run
