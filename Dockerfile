FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY pkg .
COPY cmd/server/* .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kajo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/kajo .
EXPOSE 5050
RUN chmod +x kajo
CMD ["./kajo"]