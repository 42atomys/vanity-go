FROM golang:1.17-alpine AS build

WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o go-proxy

FROM alpine

LABEL maintener "Atomys<contact@atomys.fr>"
LABEL repository "https://gitlab.com/Atomys/go-proxy"

COPY --from=build /build/go-proxy /go-proxy
COPY config/ config/

CMD [ "/go-proxy" ]