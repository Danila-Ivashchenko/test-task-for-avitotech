FROM golang:1.20 as builder
WORKDIR /build
COPY go.mod . 
COPY go.sum . 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM alpine
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]