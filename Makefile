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

.PHONY: run
run: bin/ag7if
	air -d -c .air.conf

.PHONY: clean
clean:
	rm ag7if
