FROM golang:1.20.5

# Set the Current Working Directory inside the container
WORKDIR /app/go-simple-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-simple-app .

# Run the binary program produced by `go install`
CMD ["./out/go-simple-app"]