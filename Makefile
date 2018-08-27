BINARY=goalert

VERSION="0.1.1"
COMMIT=`git rev-parse --short HEAD`
BUILD="dev"

FLAG_RELEASE=Release
FLAG_DEBUG=Debug
FLAG_TEST=Test
LDFLAG=-ldflags
LDBASEFLAGS=-X main.Version=${VERSION} -X main.Commit=${COMMIT} -X main.Build=${BUILD}

.DEFAULT_GOAL: ${BINARY}

${BINARY}:
	go build -race ${LDFLAG} "${LDBASEFLAGS}" -o ${BINARY}

win:
	GOOS=windows GOARCH=amd64 go build ${LDFLAG} "${LDBASEFLAGS}" -o ${BINARY}.exe

clean:
	go clean
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f ${BINARY}.exe ] ; then rm ${BINARY}.exe ; fi
	echo gopath is ${GOPATH}

