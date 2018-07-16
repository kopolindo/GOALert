BINARY=goalert

VERSION=0.1.1
COMMIT=`git rev-parse --short HEAD`

FLAG_RELEASE=Release
FLAG_DEBUG=Debug
FLAG_TEST=Test
LDFLAG=-ldflags
LDBASEFLAGS=-X main.Version=${VERSION} -X main.Commit=${COMMIT} -X main.Build=

.DEFAULT_GOAL: ${BINARY}

${BINARY}:
	go build -race ${LDFLAG} "${LDBASEFLAGS}${FLAG_TEST}" -o ${BINARY}

