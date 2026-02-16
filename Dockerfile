FROM golang:1.26-alpine AS builder
WORKDIR /
COPY ./ ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /myapp ./cmd

# Stage 2: Run the application
FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder /myapp /myapp

EXPOSE 8080
ENV PORT=8080

CMD ["/myapp"]