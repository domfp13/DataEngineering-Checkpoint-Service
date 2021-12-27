###########
# BUILDER #
###########

FROM golang:1.16-buster as build
LABEL maintainer="Luis Enrique Fuentes Plata"

ENV APP_HOME /usr/src/app

WORKDIR $APP_HOME

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /checkpoint-service

#########
# FINAL #
#########

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /checkpoint-service /checkpoint-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/checkpoint-service"]
