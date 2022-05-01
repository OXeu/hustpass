FROM golang:1.18.1-alpine AS build
COPY . /opt
WORKDIR /opt
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth
COPY /opt .

FROM alpine
COPY --from=build /opt/auth /opt/auth
ENTRYPOINT ["/opt/auth"]
