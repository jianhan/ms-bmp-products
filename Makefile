build:
	GOOS=linux GOARCH=amd64 go build -o products-services .
	docker build -t ms-bmp-products .

run:
	docker run -e MICRO_REGISTRY=mdns ms-bmp-products