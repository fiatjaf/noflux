FROM docker.io/library/golang:alpine3.20 AS build
RUN apk add --no-cache build-base git make
ADD . /go/src/app
WORKDIR /go/src/app
RUN make noflux

FROM docker.io/library/alpine:3.21

LABEL org.opencontainers.image.title=Noflux
LABEL org.opencontainers.image.description="Noflux is a minimalist and opinionated feed reader"
LABEL org.opencontainers.image.vendor="fiatjaf"
LABEL org.opencontainers.image.licenses=Apache-2.0
LABEL org.opencontainers.image.url=https://noflux.nostr.technology
LABEL org.opencontainers.image.source=https://github.com/fiatjaf/noflux
LABEL org.opencontainers.image.documentation=https://noflux.nostr.technology/docs/

EXPOSE 8064
ENV LISTEN_ADDR=0.0.0.0:8064
RUN apk --no-cache add ca-certificates tzdata
COPY --from=build /go/src/app/noflux /usr/bin/noflux
USER 65534
CMD ["/usr/bin/noflux"]
