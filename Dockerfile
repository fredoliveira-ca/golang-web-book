FROM golang:1.13-buster as build

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
    
WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
CMD ["/app"]