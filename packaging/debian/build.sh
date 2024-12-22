#!/bin/sh

PKG_ARCH=$(dpkg --print-architecture)
PKG_DATE=$(date -R)
PKG_VERSION=$(cd /src && git describe --tags --abbrev=0 | sed 's/^v//')

echo "PKG_VERSION=$PKG_VERSION"
echo "PKG_ARCH=$PKG_ARCH"
echo "PKG_DATE=$PKG_DATE"

cd /src

if [ "$PKG_ARCH" = "armhf" ]; then
    make noflux-no-pie
else
    CGO_ENABLED=0 make noflux
fi

mkdir -p /build/debian && \
cd /build && \
cp /src/noflux /build/ && \
cp /src/noflux.1 /build/ && \
cp /src/LICENSE /build/ && \
cp /src/packaging/noflux.conf /build/ && \
cp /src/packaging/systemd/noflux.service /build/debian/ && \
cp /src/packaging/debian/compat /build/debian/compat && \
cp /src/packaging/debian/copyright /build/debian/copyright && \
cp /src/packaging/debian/noflux.manpages /build/debian/noflux.manpages && \
cp /src/packaging/debian/noflux.postinst /build/debian/noflux.postinst && \
cp /src/packaging/debian/rules /build/debian/rules && \
cp /src/packaging/debian/noflux.dirs /build/debian/noflux.dirs && \
echo "noflux ($PKG_VERSION) experimental; urgency=low" > /build/debian/changelog && \
echo "  * Noflux version $PKG_VERSION" >> /build/debian/changelog && \
echo " -- fiatjaf <fiatjaf@gmail.com>  $PKG_DATE" >> /build/debian/changelog && \
sed "s/__PKG_ARCH__/${PKG_ARCH}/g" /src/packaging/debian/control > /build/debian/control && \
dpkg-buildpackage -us -uc -b && \
lintian --check --color always ../*.deb && \
cp ../*.deb /pkg/
