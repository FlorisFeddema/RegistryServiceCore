FROM golang:1.15.6-alpine as build

WORKDIR /build

LABEL maintainer="Floris Feddema <floris1996@hotmail.com>"

ARG VCS_REF
ARG BUILD_DATE

LABEL org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/RegistryProxy/CoreService" \
      org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.docker.dockerfile="/Dockerfile"

ENV GIN_MODE=release

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./
COPY ./internal/imports ./internal/imports
RUN go build ./internal/imports
COPY . .
RUN mkdir bin
RUN go build -a -installsuffix cgo -o bin ./...

FROM build as test
CMD go test -v ./...

FROM scratch as runtime
COPY --from=build /build/bin/CoreService /bin/service
ENTRYPOINT ["/bin/service"]