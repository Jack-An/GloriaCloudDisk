FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY user/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/user user/api/user.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/UTC /usr/share/zoneinfo/UTC
ENV TZ UTC

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc

EXPOSE 6600

CMD ["./user", "-f", "etc/user-api.yaml"]
