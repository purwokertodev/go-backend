.PHONY : test format cover

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./modules/membership/model ./modules/membership/repository ./modules/membership/query ./modules/membership/usecase ./modules/membership/presenter ./modules/auth/model \
	 ./modules/auth/query ./modules/auth/token ./modules/auth/usecase ./middleware ./config
	go test -race \
					./modules/membership/model \
					./modules/membership/repository \
					./modules/membership/query \
					./modules/membership/usecase \
					./modules/membership/presenter \
					./modules/auth/model \
					./modules/auth/query \
					./modules/auth/token \
					./modules/auth/usecase \
					./middleware \
					./config

cover: coverage.txt

coverage.txt: coverages/membership-model.txt coverages/membership-repo.txt  coverages/membership-quert.txt \
	coverages/membership-usecase.txt coverages/membership-presenter.txt coverages/auth-model.txt coverages/auth-query.txt \
	coverages/auth-token.txt coverages/auth-usecase.txt coverages/middleware.txt coverages/config.txt
	gocovmerge $^ > $@

coverages/membership-model.txt:  $(shell find ./modules/membership/model -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/model

coverages/membership-repo.txt:  $(shell find ./modules/membership/repository -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/repository

coverages/membership-quert.txt:  $(shell find ./modules/membership/query -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/query

coverages/membership-usecase.txt:  $(shell find ./modules/membership/usecase -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/usecase

coverages/membership-presenter.txt:  $(shell find ./modules/membership/presenter -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/presenter

coverages/auth-model.txt:  $(shell find ./modules/auth/model -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/auth/model

coverages/auth-query.txt:  $(shell find ./modules/auth/query -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/auth/query

coverages/auth-token.txt:  $(shell find ./modules/auth/token -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/auth/token

coverages/auth-usecase.txt:  $(shell find ./modules/auth/usecase -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/auth/usecase

coverages/middleware.txt:  $(shell find ./middleware -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./middleware

coverages/config.txt:  $(shell find ./config -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./config
