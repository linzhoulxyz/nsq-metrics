FROM alpine:latest AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

WORKDIR /app

COPY . ./

RUN apk add go

RUN go build

FROM alpine:latest

EXPOSE 9117

WORKDIR /app

COPY --from=builder /app/nsq-metrics ./

CMD ["./nsq-metrics"]
