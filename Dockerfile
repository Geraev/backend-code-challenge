FROM golang:1.16-alpine
RUN mkdir -p /app
COPY . /app
WORKDIR /app/cmd/backend-code-challenge
RUN go build -o app .
CMD ["/app/cmd/backend-code-challenge/app"]

EXPOSE 3331