# Developer setup

Interested in contributing to Masonry? Awesome! Take a look at our [contribution guidelines](../CONTRIBUTING.md) first.

## Project setup

1. Install Go 1.9+, and ensure your `GOPATH` is set. Using [gvm](https://github.com/moovweb/gvm) is recommended.
1. Get the code.

    ```sh
    go get github.com/opencontrol/compliance-masonry
    ```

1. Run the tool.

    ```sh
    compliance-masonry
    ```

This should print out usage documentation.

1. Install `make`. See [GNU Make](https://www.gnu.org/software/make/) for details on what `make` is and can do.

    `yum`-based systems (RHEL / CentOS / etc.):

    ```sh
    sudo yum install make
    ```

    `apt`-based systems (Ubuntu / etc.):

    ```sh
    sudo apt-get install build-essential
    ```

    MacOS systems: 
    First, install [Homebrew](https://brew.sh). Then, install `make`:

    ```sh
    brew install make

1. `Makefile` targets:

    ```sh
    make [target-name]
    ```

    The common targets include:
    * `all` - Performs a `build`
    * `build` - Builds the source code and places `compliance-masonry` binary into the `./build` folder.
    * `clean` - Simply removes the `./build` folder
    * `test` - Runs the tests.
    * `lint` - Checks to see if the Go code is properly formatted. (If you want to contribute to the project, use this target; you will need to make sure your code follows accepted standards.)

## Updating dependencies

As the dependencies now exist in the `git` tree under the `vendor/` folder,
dependencies should only have to be updated when they are out-of-date, need
to stick to a specific version, or need to add a new dependency.

1. Get the `vndr` handling tool.

    ```sh
    go get github.com/LK4D4/vndr
    ```

1. When needed, update dependencies by running the `vndr` tool in the project.

   ```sh
   vndr
   ```

1. If any dependencies do not exist in the `vendor/` folder, add them to `vendor.conf` and re-run the `vndr` tool.

## Running tests

```sh
make test
```

## Creating binaries

This will only be relevant for maintainers.

### One-time setup for uploading binaries

1. Install [goxc](https://github.com/laher/goxc)

    ```sh
    go get github.com/laher/goxc
    ```

1. [Get a GitHub API token](https://github.com/settings/tokens/new). The token should have write access to repos.
1. Add a .goxc.local.json file with a github api key

    ```sh
    goxc -wlc default publish-github -apikey=123456789012
    ```

### Compiling and uploading binaries

1. Set version number in:
    * [`.goxc.json`](.goxc.json)
    * As `app.Version` in [`masonry-go.go`](masonry-go.go)
1. Run the release script

    ```sh
    ./release.sh
    ```
