run: stop up

mod:
	# This make rule requires Go 1.11+
	GO111MODULE=on go mod tidy

up:
	docker-compose -f deployments/docker-compose.yml up -d

stop:
	docker-compose -f deployments/docker-compose.yml stop

down:
	docker-compose -f deployments/docker-compose.yml down
