FROM docker.io/library/golang:bookworm AS build
ADD . /go/src/app
WORKDIR /go/src/app
RUN make noflux

FROM gcr.io/distroless/base-debian12:nonroot

LABEL org.opencontainers.image.title=Noflux
LABEL org.opencontainers.image.description="Noflux is a minimalist and opinionated feed reader"
LABEL org.opencontainers.image.vendor="Frédéric Guillot"
LABEL org.opencontainers.image.licenses=Apache-2.0
LABEL org.opencontainers.image.url=https://noflux.nostr.technology
LABEL org.opencontainers.image.source=https://github.com/fiatjaf/noflux
LABEL org.opencontainers.image.documentation=https://noflux.nostr.technology/docs/

EXPOSE 8064
ENV LISTEN_ADDR=0.0.0.0:8064
COPY --from=build /go/src/app/noflux /usr/bin/noflux
CMD ["/usr/bin/noflux"]
