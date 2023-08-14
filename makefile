.PHONY: build run

build:
	docker build -t dealls .
run:
	docker run -p 8080:8080 dealls
compose-up:
	docker-compose up
compose-down:
	docker-compose down