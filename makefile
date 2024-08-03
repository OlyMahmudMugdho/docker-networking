build-server:
	@docker build -t go-server ./go-server

build-client:
	@docker build -t go-client ./go-client

create-network:
	@docker network create mynetwork

run-all:
	@docker compose up

run-all-detached:
	@docker compose up -d

run-server:
	@docker run -p 8080:8080 --name go-server go-server

run-server-detached:
	@docker run -p 8080:8080 --name go-server go-server -d

run-client:
	@docker run -p 8080:8080 --name go-client go-client

run-client-detached:
	@docker run -p 8080:8080 --name go-client go-client -d

stop-all:
	@docker stop go-server go-client postgres
