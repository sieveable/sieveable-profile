all: build
check-env:
ifndef DB
	$(error Environment variable DB is undefined)
endif
ifndef USER
	$(error Environemnt variable USER is undefined)
endif
deps:
	go get -v ./...
test-deps:
	go get -t -v ./...
test: check-env test-deps
	go test -v -covermode=count -coverprofile=profile_dbwriter.cov github.com/sieveable/sieveable-profile/dbwriter
	go test -v -covermode=count -coverprofile=profile_dbretrieval.cov github.com/sieveable/sieveable-profile/dbretrieval

build: deps test
	go build -o sieveable_profile_writer.out sieveable_profile_writer.go

