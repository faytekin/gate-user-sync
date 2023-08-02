# Build Stage
FROM golang:1.20-alpine as BUILDER
LABEL authors="faytekin"

# install ca-certificates
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production Stage
FROM scratch

# Copy the certificates
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy the binary file
WORKDIR /app/
COPY --from=BUILDER /app/main .

CMD ["./main"]
