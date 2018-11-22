# FROM golang:alpine as builder
# RUN apk update; apk add git
# RUN go get -u github.com/golang/dep/cmd/dep
# RUN mkdir -p $GOPATH/src/github.com/abits/puzzlr 
# RUN mkdir /build
# ADD . $GOPATH/src/github.com/abits/puzzlr
# WORKDIR $GOPATH/src/github.com/abits/puzzlr
# RUN dep ensure
# RUN go build -o /build/puzzlr .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY puzzlr /app/
WORKDIR /app
EXPOSE 80
CMD ["./puzzlr"]
