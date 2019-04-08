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

1. Get the [`dep`](https://golang.github.io/dep/) tool.

    https://golang.github.io/dep/docs/installation.html

1. When needed, update dependencies by running the `dep` tool in the project.

   ```sh
   dep ensure -update github.com/foo/bar
   ```

1. You can add dependencies using `dep`.

   ```sh
   dep ensure -add github.com/foo/bar
   ```

1. If dependencies are not in the `vendor/` folder, you can use the `dep` tool to make everything proper.

   ```sh
   dep ensure
   ```

## Running tests

```sh
make test
```

## Tagging a New Release

1. Checkout the master branch
NOTE: Make sure that the master branch is clean and has the latest commits from GitHub.

    ```sh
    git checkout master
    ```

1. Using `v.1.1.1` as an example, tag the new release using the convention in the example below:

    ```sh
    git tag -m "Bump to v1.1.1" v1.1.1
    ```

1. Using `v1.1.1` as an example, push the tag back to GitHub

    ```sh
    git push origin v1.1.1
    ```

1. CircleCI will then run through the tests. Since there is a new tag, CircleCI will also install and run
[GoReleaser](https://github.com/goreleaser/goreleaser) which will build and upload the binaries for release.

