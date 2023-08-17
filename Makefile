.PHONY: run run-email run-kafka

run-email:
	cd ./email && docker-compose up --build -d

run-kafka:
	docker-compose up -d

run: run-kafka run-email
