FROM golang:1.17-alpine

COPY . /src

WORKDIR /src

RUN go build cmd/main/main.go

FROM alpine:3.15

COPY --from=0 /src/main /

CMD /main