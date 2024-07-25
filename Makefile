.PHONY: up down logs run

# Up containers of docker file
up:
	docker-compose up -d

# Down containers
down:
	docker-compose down

# Logs of containers
logs:
	docker-compose logs

# Run main.go
run:
	go run main.go
