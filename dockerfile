FROM golang:1.15 as builder
WORKDIR /go/src/app
ADD . .
COPY startapp.sh /startapp.sh
# #BUILD APP
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./cmd/ .
# ENV PORT 8080

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/app ./
RUN chmod 775 /startapp.sh
ENTRYPOINT /startapp.sh

# CMD ["go", "run", "cmd/main.go"]