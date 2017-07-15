# Developer setup

Interested in contributing to Masonry? Awesome! Take a look at our [contribution guidelines](../CONTRIBUTING.md) first.

## Project setup

1. Install Go 1.6+, and ensure your `GOPATH` is set. Using [gvm](https://github.com/moovweb/gvm) is recommended.
1. Install the tool.

    ```sh
    go get github.com/opencontrol/compliance-masonry
    compliance-masonry
    ```

This should print out usage documentation. You can find the code in `$GOPATH/src/github.com/opencontrol/compliance-masonry/`.

## Running tests

```sh
# Get test dependencies
go get -t ./...
# Run tests
go test $(go list ./... | grep -v vendor)
```

## Updating dependencies

1. Install [Dep](https://github.com/golang/dep).
1. Run

    ```sh
    dep ensure -update
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
