.PHONY : test format

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./modules/membership/model ./modules/membership/repository
	go test -race \
					./modules/membership/model \
					./modules/membership/repository
