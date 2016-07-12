all: build
check-env:
ifndef DB
	$(error Environment variable DB is undefined)
endif
ifndef USER
	$(error Environemnt variable USER is undefined)
endif
test: check-env
	go test -v -covermode=count -coverprofile=profile_dbwriter.cov github.com/sieveable/sieveable-profile/dbwriter
	go test -v -covermode=count -coverprofile=profile_dbretrieval.cov github.com/sieveable/sieveable-profile/dbretrieval

build: test
	go build -o sieveable_profile_writer sieveable_profile_writer.go
