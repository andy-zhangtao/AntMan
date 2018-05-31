
.PHONY: build
name = dns-antman

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)
	mv $(name) bin/$(name)

run: build
	./bin/$(name)

release: *.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o $(name)
	mv $(name) bin/$(name)
	docker build -t ccr.ccs.tencentyun.com/eqxiu/caas-$(name) .
	docker push ccr.ccs.tencentyun.com/eqxiu/caas-$(name)
	rm bin/$(name)
