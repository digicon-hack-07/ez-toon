.PHONY: init
init:
	cd client && npm ci
	cd server && go mod download

.PHONY: build-client
build-client:
	cd client && npm run build

.PHONY: build-server
build-server:
	cd server && go build -o ../ez-toon *.go

.PHONY: build
build: build-client build-server
	rm -rf public
	cp -r client/dist public

.PHONY: run
run: build-client
	rm -rf server/public
	cp -r client/dist server/public
	cd server && go run *.go
