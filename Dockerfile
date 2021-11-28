FROM golang:1.16
WORKDIR /go/src/github.com/stevemurr/cbot/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/stevemurr/cbot/cmd/config.json ./
COPY --from=0 /go/src/github.com/stevemurr/cbot/app ./
ENTRYPOINT ["/root/app"]  