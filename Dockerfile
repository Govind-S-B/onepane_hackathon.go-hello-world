FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
EXPOSE 9009
CMD ["go", "run", "main.go"]