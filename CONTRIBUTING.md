## Welcome!

We're so glad you're thinking about contributing to an 18F open source project! If you're unsure about anything, just ask -- or submit the issue or pull request anyway. The worst that can happen is you'll be politely asked to change something. We love all friendly contributions.

We want to ensure a welcoming environment for all of our projects. Our staff follow the [18F Code of Conduct](https://github.com/18F/code-of-conduct/blob/master/code-of-conduct.md) and all contributors should do the same.

We encourage you to read this project's CONTRIBUTING policy (you are here), its [LICENSE](LICENSE.md), and its [README](README.md).

If you have any questions or want to read more, check out the [18F Open Source Policy GitHub repository]( https://github.com/18f/open-source-policy), or just [shoot us an email](mailto:18f@gsa.gov).

## Development

This project uses [glide](https://github.com/Masterminds/glide) to manage vendored dependencies.

### Project setup

1. Install Go 1.6+, and ensure your `GOPATH` is set. Using [gvm](https://github.com/moovweb/gvm) is recommended.
1. Install the tool.

    ```sh
    go get github.com/opencontrol/compliance-masonry
    compliance-masonry
    ```

This should print out usage documentation. You can find the code in `$GOPATH/src/github.com/opencontrol/compliance-masonry/`.

### Running tests

```sh
# Get test dependencies
go get -t ./...
# Run tests
go test $(glide nv)
```

### Updating dependencies

Masonry uses [glide](https://github.com/Masterminds/glide) to manage dependencies.

```sh
go get github.com/Masterminds/glide
glide up --all-dependencies
```

### Creating Binaries

#### One Time Setup for Uploading Binaries

1. Install [goxc](go get github.com/laher/goxc)

    ```sh
    go get github.com/laher/goxc
    ```

1. [Get a GitHub API token](https://github.com/settings/tokens/new). The token should have write access to repos.
1. Add a .goxc.local.json file with a github api key

    ```sh
    goxc -wlc default publish-github -apikey=123456789012
    ```

#### Compiling and Uploading Binaries

1. Set version number in:
    * [`.goxc.json`](.goxc.json)
    * As `app.Version` in [`masonry-go.go`](masonry-go.go)
1. Run the release script

    ```sh
    ./release.sh
    ```

## Public domain

This project is in the public domain within the United States, and
copyright and related rights in the work worldwide are waived through
the [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).

All contributions to this project will be released under the CC0
dedication. By submitting a pull request, you are agreeing to comply
with this waiver of copyright interest.
