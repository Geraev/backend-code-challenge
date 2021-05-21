FROM golang:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
EXPOSE 3331:3031
WORKDIR /go/src/app/cmd/backend-code-challenge
RUN go build -o main .
CMD ["/go/src/app/cmd/backend-code-challenge/main"]
