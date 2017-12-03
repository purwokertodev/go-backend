.PHONY : test format

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

test: ./modules/membership/model ./modules/membership/repository
	go test -race \
					./modules/membership/model \
					./modules/membership/repository

cover: cover.txt

cover.txt: coverages/membership-model.txt coverages/membership-repo.txt
	gocovmerge $^ > $@

coverages/membership-model.txt:  $(shell find ./modules/membership/model -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/model

coverages/membership-repo.txt:  $(shell find ./modules/membership/repository -type f)
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/membership/repository
