TEST_FILE?=c.out

default: alltest


.PHONY: .govet
govet:
	@ go vet . && go fmt ./... && \
	(if [[ "$(gofmt -d $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./tests/*" -not -path "./assets/*"))" == "" ]]; then echo "Good format"; else echo "Bad format"; exit 33; fi);

.PHONY: clean
clean:
	@ find . -name ${TEST_FILE} -print0 | xargs -0 rm -rf

.PHONY: alltest
alltest: clean
	go test -v -coverprofile=${TEST_FILE} ./...