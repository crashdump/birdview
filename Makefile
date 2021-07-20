build:
	GOOS=darwin GOARCH=arm64 go build -o dist/birdview-darwin-arm  cmd/*.go

run:
	go run ./cmd/

test:
	go fmt ./...
	go test -v ./... -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

release-test:
	goreleaser --snapshot --skip-publish --rm-dist
