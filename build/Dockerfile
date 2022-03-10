FROM golang:1.17-alpine AS build

WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o vanity-go

FROM alpine

LABEL maintener "Atomys<contact@atomys.fr>"
LABEL repository "https://gitlab.com/Atomys/vanity-go"

COPY --from=build /build/vanity-go /vanity-go

EXPOSE 8080

ENTRYPOINT [ "/vanity-go" ]

CMD [ "serve" ]