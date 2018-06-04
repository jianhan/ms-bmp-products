build:
	GOOS=linux GOARCH=amd64 go build -o products-service .
	docker build -t ms-bmp-products .

run:
	docker run ms-bmp-products