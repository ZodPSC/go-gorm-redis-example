FROM golang:1.23.4 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
#RUN go mod tidy
#RUN go get github.com/mattn/go-sqlite3

COPY *.go ./
COPY storage storage

#RUN go env -w CGO_ENABLED=1
#RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o /golang-app
RUN CGO_ENABLED=1 GOOS=linux go build -o /golang-app

# run test
FROM build-stage AS run-test-stage
RUN go test -v ./..

# deploy
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /golang-app /golang-app

EXPOSE 8088

USER nonroot:nonroot

ENTRYPOINT ["/golang-app"]