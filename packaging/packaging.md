= Compliance Masonry Installation Packages

== Creating Packages

**Note:** RPM and DPKG packages are for personal installation and testing only and have not gone through distribution-specific review and approvals
          to be included as part of the distribution-specific repositories.

=== RPM

1. Install `rpmbuild`.
    `yum`-based systems (RHEL / CentOS / etc.):

    ```sh
    yum -y install rpmbuild
    ```
    `dnf`-based systems (Fedora / etc.):
    ```sh
    dnf -y install rpmbuild
    ```

1. Build RPM package.

   ```sh
   rpmbuild -ba compliance-masonry.spec
   ```

=== DPKG

**_[Help Wanted]_** some initial `dpkg` files are under the `debian` directory

=== Windows

TBD

=== MacOS

TBD
