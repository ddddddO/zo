# syntax=docker/dockerfile:1

FROM golang:1.13-alpine as build
WORKDIR /go/src/github.com/ddddddO/zo
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
COPY cmd/ui/main.go ./cmd/ui/main.go
COPY storage ./storage
RUN go build -o /ui cmd/ui/main.go

FROM alpine
WORKDIR /
COPY --from=build /ui /ui
EXPOSE 8080
ENV PORT=8080
ENV BUCKET_NAME=testddddddo
CMD [ "/ui" ]
