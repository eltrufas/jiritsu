FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.sum .
COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine
COPY --from=builder /app/jiritsu /app/
EXPOSE 8080
ENTRYPOINT ["/app/jiritsu"]