FROM golang:1.18.1-alpine AS build
COPY . /opt
WORKDIR /opt
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk add git
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth


FROM alpine
COPY --from=build /opt/auth /opt/auth
ENTRYPOINT ["/opt/auth"]
