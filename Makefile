test:
	go test -v ./...
proxy-test:
	docker run \
	-it \
	-p 0.0.0.0:80:80 \
	-v $(shell pwd)/test.conf:/etc/nginx/nginx.conf \
	nginx:1.13.12-alpine
