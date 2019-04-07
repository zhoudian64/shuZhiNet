FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/shuZhiNet
WORKDIR /go/src/shuZhiNet
RUN go get && go build

FROM redis
COPY --from=builder /go/src/shuZhiNet/shuZhiNet /shuZhiNet
WORKDIR /
CMD ls
CMD ./shuZhiNet
EXPOSE 8000
