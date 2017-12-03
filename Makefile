.PHONY : test format cover

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./modules/membership/model ./modules/membership/repository ./modules/membership/query ./modules/auth/token ./middleware
	go test -race \
					./modules/membership/model \
					./modules/membership/repository \
					./modules/membership/query \
					./modules/auth/token \
					./middleware

cover: coverage.txt

coverage.txt: coverages/membership-model.txt coverages/membership-repo.txt  coverages/membership-quert.txt coverages/auth-token.txt coverages/middleware.txt
	gocovmerge $^ > $@

coverages/membership-model.txt:  $(shell find ./modules/membership/model -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/model

coverages/membership-repo.txt:  $(shell find ./modules/membership/repository -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/repository

coverages/membership-quert.txt:  $(shell find ./modules/membership/query -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/query

coverages/auth-token.txt:  $(shell find ./modules/auth/token -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/auth/token

coverages/middleware.txt:  $(shell find ./middleware -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./middleware
