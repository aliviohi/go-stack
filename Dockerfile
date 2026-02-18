FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-stack ./main.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /go-stack /go-stack

ENTRYPOINT ["/go-stack"]
