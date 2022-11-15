FROM alpine:latest

EXPOSE 9117

WORKDIR /app

COPY . ./

RUN apk add go

RUN go build

CMD ["./nsq-metrics"]
