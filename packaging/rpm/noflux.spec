%undefine _disable_source_fetch

Name:    noflux
Version: %{_noflux_version}
Release: 1.0
Summary: Nostr and RSS feed reader
URL: https://noflux.app/
License: ASL 2.0
Source0: noflux
Source1: noflux.service
Source2: noflux.conf
Source3: noflux.1
Source4: LICENSE
BuildRoot: %{_topdir}/BUILD/%{name}-%{version}-%{release}
BuildArch: x86_64
Requires(pre): shadow-utils

%{?systemd_ordering}

AutoReqProv: no

%define __strip /bin/true
%define __os_install_post %{nil}

%description
%{summary}

%install
mkdir -p %{buildroot}%{_bindir}
install -p -m 755 %{SOURCE0} %{buildroot}%{_bindir}/noflux
install -D -m 644 %{SOURCE1} %{buildroot}%{_unitdir}/noflux.service
install -D -m 600 %{SOURCE2} %{buildroot}%{_sysconfdir}/noflux.conf
install -D -m 644 %{SOURCE3} %{buildroot}%{_mandir}/man1/noflux.1
install -D -m 644 %{SOURCE4} %{buildroot}%{_docdir}/noflux/LICENSE

%files
%defattr(755,root,root)
%{_bindir}/noflux
%{_docdir}/noflux
%defattr(644,root,root)
%{_unitdir}/noflux.service
%{_mandir}/man1/noflux.1*
%{_docdir}/noflux/*
%defattr(600,root,root)
%config(noreplace) %{_sysconfdir}/noflux.conf

%pre
getent group noflux >/dev/null || groupadd -r noflux
getent passwd noflux >/dev/null || \
    useradd -r -g noflux -d /dev/null -s /sbin/nologin \
    -c "Noflux Daemon" noflux
exit 0

%post
%systemd_post noflux.service

%preun
%systemd_preun noflux.service

%postun
%systemd_postun_with_restart noflux.service

%changelog
