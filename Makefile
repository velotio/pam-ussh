MODULE := pam_authz
NEED_SYMLINK := $(shell if ! stat -q .go/src/pam-ussh 2>&1 > /dev/null ; then echo "yes" ; fi)

module: test
	GOPATH=${PWD}/.go go build -buildmode=c-shared -o ${MODULE}.so

install: module
	sudo cp ${MODULE}.so /lib/x86_64-linux-gnu/security/${MODULE}.so
	docker cp ${MODULE}.so ssh-ca-cert-run-1:/lib/x86_64-linux-gnu/security/${MODULE}.so

test: *.go .go/src
	GOPATH=${PWD}/.go go test -cover

.go/src:
	-mkdir -p ${PWD}/.go/src
ifeq ($(NEED_SYMLINK),yes)
	ln -s ${PWD} ${PWD}/.go/src/pam-ussh
endif
	GOPATH=${PWD}/.go go get golang.org/x/crypto/ssh
	GOPATH=${PWD}/.go go get golang.org/x/crypto/ssh/agent
	GOPATH=${PWD}/.go go get github.com/google/logger
	GOPATH=${PWD}/.go go get github.com/stretchr/testify/require

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h
	-rm -rf .go/

docker_build:
	docker build --rm -f "Dockerfile" -t ssh-ca-cert:latest .

docker_run: docker_build
	docker stop ssh-ca-cert-run-1 || true
	docker rm ssh-ca-cert-run-1 || true
	docker run --rm -d -p 2201:22/tcp --name ssh-ca-cert-run-1 ssh-ca-cert:latest

.PHONY: test module download_deps clean
