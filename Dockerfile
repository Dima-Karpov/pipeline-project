ARG GOLANG_VERSION=1.22.0

FROM golang:${GOLANG_VERSION}-alpine AS build-stage

WORKDIR /app

COPY . /app
ENV CGO_ENABLED=0
RUN go build -o /app/pipeline-app /app/main.go

FROM scratch
WORKDIR /
COPY --from=build-stage /app/pipeline-app .
EXPOSE 8989
ENTRYPOINT [ "/pipeline-app" ]


