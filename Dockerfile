FROM golang:1.15-alpine AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/ldegaetano/agileengine-images 
ADD . /go/src/github.com/ldegaetano/agileengine-images
WORKDIR /go/src/github.com/ldegaetano/agileengine-images

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -tags musl -a -installsuffix cgo -o app .

FROM alpine
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /go/src/github.com/ldegaetano/agileengine-images .
ENTRYPOINT ["/root/app"]
