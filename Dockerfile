# hadolint ignore=DL3026
FROM golang:1.13 as base

WORKDIR /usr/src/app/core-paas-sumologic-provider

COPY go.mod go.sum ./
RUN go mod download
COPY . /usr/src/app/core-paas-sumologic-provider

FROM base as test
RUN make test
RUN mkdir /sonar && mv .coverage /sonar/.coverage

FROM base as builder
RUN make build

FROM artifacts.msap.io/mulesoft/core-paas-base-image-ubuntu:v3.0.199 AS prod

COPY --from=builder /go/bin/sumologic-terraform-provider /terraform-provider-sumologic

USER 2020
CMD ["/terraform-provider-sumologic"]
