# Name
BINARY=Cli
IMPORTPATH=`go list`/

# Variables
VERSION=0.2.0
COMMIT=`git rev-parse --short HEAD`

FLAG_RELEASE=Release
FLAG_DEBUG=Debug
FLAG_TEST=Test
LDFLAG=-ldflags
LDBASEFLAGS=-X ${IMPORTPATH}cli.Version=${VERSION} -X ${IMPORTPATH}cli.Commit=${COMMIT} -X ${IMPORTPATH}cli.Build=

.DEFAULT_GOAL: ${BINARY}

${BINARY}:
	go build ${LDFLAG} "${LDBASEFLAGS}${FLAG_DEBUG}" -o ${BINARY}

#TODO make static

run:
	go run ${LDFLAG} "${LDBASEFLAGS}${FLAG_DEBUG}" *.go

test:
	go test ${LDFLAG} "${LDBASEFLAGS}${FLAG_TEST}" ./...

testv:
	go test -v ${LDFLAG} "${LDBASEFLAGS}${FLAG_TEST}" ./...

testvr:
	go test -v -race ${LDFLAG} "${LDBASEFLAGS}${FLAG_TEST}" ./...

install:
	# Installing the executable
	go install ${LDFLAG} "${LDBASEFLAGS}${FLAG_RELEASE}" 

installdbg:
	# Installing the debug executable
	go install ${LDFLAG} "${LDBASEFLAGS}${FLAG_DEBUG}" 

release:
	go build ${LDFLAG} "${LDBASEFLAGS}${FLAG_RELEASE}" -o ${BINARY}

installdeps:
	# Installing dependencies to embed assets
	go get github.com/GeertJohan/go.rice/...

updatedeps:
	# Updating dependencies to embed assets
	go get -u github.com/GeertJohan/go.rice/...

clean:
	-rm **/rice-box.go
	go clean
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
