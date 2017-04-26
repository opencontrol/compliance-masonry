Name:		compliance-masonry
Version:	1.1.2
Release:	1%{?dist}
Summary:	Compliance Masonry is a command-line interface (CLI) that allows users to construct
		certification documentation using the OpenControl Schema.

Group:		
License:	CC0 1.0 Universal Public Domain 
URL:		https://github.com/opencontrol/compliance-masonry
Source0:	https://github.com/opencontrol/compliance-masonry/archive/v%{version}.tar.gz

BuildRequires:	gcc
BuildRequires:	golang >= 1.2-7
#BuildRequires:	golang(github.com/example/thing) >= 0-0.13
Requires:	

%description
Compliance Masonry is a command-line interface (CLI) that allows users to construct certification documentation using the OpenControl Schema.


%prep
%setup -qa -n compliance-masonry-%{version}

# Many golang binaries are "vendoring" (bundling) sources, which is a terrible
# practice. We must remove them and build those dependency packages independently.
rm -rf vendor/

%build
# Setup temporary build gopath, and put compliance-masonry there
mkdir -p ./_build/src/github.com/opencontrol
ln -s $(pwd) ./_build/src/github.com/opencontrol/compliance-masonry

export GOPATH=$(pwd)/_build:%{gopath}
go build -o compliance-masonry .

#%configure
#make %{?_smp_mflags}

%install
install -d %{buildroot}${_bindir}
install -p -m 0755 ./compliance-masonry %{buildroot}%{_bindir}/compliance-masonry
make install DESTDIR=%{buildroot}

%files
%defattr(-,root,root,-)

%doc AUTHORS CHANGELOG.md CONTRIBUTING.md LICENSE.MD docs/README.md docs/development.md docs/gitbook.md docs/install.md docs/masonry-for-the-compliance-literate.md docs/usage.md


%changelog
* Thurs Apr 27 2017 Shawn Wells <shawn@redhat.com> - 1.1.2-1
- Initial RPM for compliance-masonry
