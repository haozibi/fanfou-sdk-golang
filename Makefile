
.PHONY: .govet
govet:
	@ go vet . && go fmt ./... && \
	(if [[ "$(gofmt -d $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./tests/*" -not -path "./assets/*"))" == "" ]]; then echo "Good format"; else echo "Bad format"; exit 33; fi);