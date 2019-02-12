# Installation

Compliance Masonry is packaged into a downloadable executable program for those who want to use Compliance Masonry without the need to install any external dependencies or programs.

## Mac OS X

In your terminal, run the following:

```sh
cd ~/Downloads
curl -L https://github.com/opencontrol/compliance-masonry/releases/download/v1.1.5/compliance-masonry_1.1.5_darwin_amd64.zip -o compliance-masonry.zip
unzip compliance-masonry.zip
cp compliance-masonry_1.1.5_darwin_amd64/compliance-masonry /usr/local/bin
```

## Windows

1. Go to [the Github Release](https://github.com/opencontrol/compliance-masonry/releases/latest).
1. Download the package that corresponds to your machine and operating system.
    - For 32 Bit Windows, you'll want the file ending `_windows_386.zip`
    - For 64 Bit Windows, you'll want the file ending `_windows_amd64.zip`
1. Double-click on the downloaded package to unzip the archive. The resulting folder should contain a file called `compliance-masonry.exe`.
1. Create a folder, e.g. `C:\Masonry\bin`.
1. Drag `compliance-masonry.exe` into the new folder.
1. Open PowerShell.
    * Search your Start menu / Cortana for it.
1. [Add `C:\Masonry\bin` to your `PATH`.](https://www.java.com/en/download/help/path.xml)

## Linux

The instructions below are for 64-bit architectures. See the [releases](https://github.com/opencontrol/compliance-masonry/releases) page for others.

In your terminal, run the following:

```sh
curl -L https://github.com/opencontrol/compliance-masonry/releases/download/v1.1.5/compliance-masonry_1.1.5_linux_amd64.tar.gz -o compliance-masonry.tar.gz
tar -xf compliance-masonry.tar.gz
sudo cp compliance-masonry_1.1.5_linux_amd64/compliance-masonry /usr/local/bin
```

---

See the [usage](usage.md) guide next.
