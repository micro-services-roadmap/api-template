FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
#ENV GOPROXY https://goproxy.cn,direct
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories


RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ARG GH_PRIVATE_TOKEN=""
ARG GOPRIVATE="github.com/wordpress-plus"
RUN apk add git\
    && git config --global url.https://${GH_PRIVATE_TOKEN}@github.com/.insteadOf https://github.com/

ADD go.mod .
ADD go.sum .
COPY . .
COPY ./etc /app/etc
# RUN go build -ldflags="-s -w" -o /app/api-admin .

RUN go mod tidy \
    && go build -o /app/api-app .


FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/api-app /app/api-app
COPY --from=builder /app/etc /app/etc

CMD ["./api-app"]
