run:
	docker build -t red-scare:latest .
	docker run -it red-scare:latest

test:
	docker build --target=base -t red-scare-tests .
	docker run red-scare-tests go test -cover ./...