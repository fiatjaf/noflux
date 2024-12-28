FROM golang:1 AS build
ENV CGO_ENABLED=0
ADD . /go/src/app
WORKDIR /go/src/app
RUN make noflux

FROM rockylinux:9
RUN dnf install --setopt=install_weak_deps=False -y rpm-build systemd-rpm-macros
RUN mkdir -p /root/rpmbuild/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
RUN echo "%_topdir /root/rpmbuild" >> .rpmmacros
COPY --from=build /go/src/app/noflux /root/rpmbuild/SOURCES/noflux
COPY --from=build /go/src/app/LICENSE /root/rpmbuild/SOURCES/
COPY --from=build /go/src/app/noflux.1 /root/rpmbuild/SOURCES/
COPY --from=build /go/src/app/packaging/systemd/noflux.service /root/rpmbuild/SOURCES/
COPY --from=build /go/src/app/packaging/noflux.conf /root/rpmbuild/SOURCES/
COPY --from=build /go/src/app/packaging/rpm/noflux.spec /root/rpmbuild/SPECS/noflux.spec
