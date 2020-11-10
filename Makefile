build:
	docker build -t golang-api .
run:
	docker run -p 8080:8000 golang-api

up:
	build run
	
