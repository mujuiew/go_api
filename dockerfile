FROM golang:1.11.2-alpine3.8
WORKDIR /app

COPY go.mod .
COPY go.sum .
# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...
# RUN go build
COPY . .
RUN go test ./.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
CMD ["./app"]