FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/zzayne/go-blog

COPY . .
RUN ls

RUN GOARCH=amd64  GOOS=linux go build 



FROM alpine:lastest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/zzayne/go-blog/goblog .


CMD ["./goblog"]
