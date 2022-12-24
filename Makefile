pubg:
	go build ${GCFLAGS} -ldflags "${LDFLAGS}" -o pubg ./cmd

fmt:
	find . -type f -name '*.go' -exec goimports -l -w {} \;

vet:
	go vet ./...

clean:
	rm -vf ./pubg
