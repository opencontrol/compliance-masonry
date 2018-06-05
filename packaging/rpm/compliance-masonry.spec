# Build with debug info rpm
%global with_debug 0
# Run tests in check section
%global with_check 1

%if 0%{?with_debug}
%global _dwz_low_mem_die_limit 0
%else
%global debug_package   %{nil}
%endif

%if ! 0%{?gobuild:1}
%define gobuild(o:) go build -ldflags "${LDFLAGS:-} -B 0x$(head -c20 /dev/urandom|od -An -tx1|tr -d ' \\n')" -a -v -x %{?**};
%endif

%global provider        github
%global provider_tld    com
%global project         opencontrol
%global repo            compliance-masonry
# https://github.com/opencontrol/compliance-masonry
%global provider_prefix %{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path     %{provider_prefix}

Name:           %{repo}
Version:        1.1.3
Release:        1%{?dist}
Summary:        Security Documentation Builder
License:        CC0 1.0 Universal Public Domain
URL:            https://%{provider_prefix}
Source0:        https://%{provider_prefix}/archive/%{repo}-%{version}.tar.gz

# e.g. el6 has ppc64 arch without gcc-go, so EA tag is required
ExclusiveArch:  %{?go_arches:%{go_arches}}%{!?go_arches:%{ix86} x86_64 aarch64 %{arm}}
# If go_compiler is not set to 1, there is no virtual provide. Use golang instead.
BuildRequires:  %{?go_compiler:compiler(go-compiler)}%{!?go_compiler:golang}

%description
Compliance Masonry is a command-line interface (CLI) that allows users to construct
certification documentation using the OpenControl Schema

%prep
%setup -q -n %{repo}-%{version}

%build
mkdir -p src/%{provider}.%{provider_tld}/%{project}
ln -s ../../../ src/%{import_path}
ln -s vendor src

export GOPATH=$(pwd):%{gopath}
%gobuild -o bin/%{name} %{import_path}/

%install
install -d -p %{buildroot}%{_bindir}
install -p -m 0755 bin/%{name} %{buildroot}%{_bindir}

%check

%files
%license LICENSE.md
%doc README.md CONTRIBUTING.md
%{_bindir}/%{name}

%changelog
* Fri May 25 2018 Gabe <redhatrises@gmail.com> - 1.1.3-1
- First Initial RPM package

