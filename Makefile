default: bin/ag7if

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
SRC = $(basename $(wildcard */*.go))
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

bin/ag7if: main.go $(foreach f, $(SRC), $(f).go)
	go build ${LDFLAGS} -o bin/ag7if

.PHONY: install
install: bin/ag7if
	-@rm ${GOPATH}/bin/ag7if
	cp bin/ag7if ${GOPATH}/bin/

.PHONY: debug
debug: bin/ag7if
	go get github.com/cosmtrek/air
	air -d -c .air.conf

.image: bin/ag7if Dockerfile
	docker build .

.PHONY: run
run: .image
	go get github.com/mikefarah/yq


.PHONY: clean
clean:
	-@rm ag7if
