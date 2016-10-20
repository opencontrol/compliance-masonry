# GitBook

## Creating

1. Update dependencies

    ```bash
    compliance-masonry get
    ```

1. Run the gitbook command

    ```bash
    compliance-masonry docs gitbook LATO
    # Or
    compliance-masonry docs gitbook FedRAMP-low
    ```

The `gitbook` command by default will create a folder called `exports` that contains the files needed to create a gitbook. Visit the [gitbook documentation](https://github.com/GitbookIO/gitbook-cli) for more information on creating gitbooks via the CLI.

## Adding additional markdown content (optional)

Security documentation usually requires unstructured information that is not captured in the control documentation. The `markdowns` directory can be used to add this supplemental information.

1. Create a `markdowns` folder in the same directory as the `opencontrol.yaml`.

    ```bash
    mkdir markdowns
    ```

2. Create the `markdowns/SUMMARY.md` and `markdowns/README.md` documents.
    ```bash
    touch markdowns/SUMMARY.md
    touch markdowns/README.md
    ```

The content of the `markdowns/SUMMARY.md` and `markdowns/README.md` files and the files they reference is prepended to the Gitbook documentation.

For more information on using the `SUMMARY.md` and `README.md` files visit the [Gitbook documentation.](http://toolchain.gitbook.com/structure.html) For an example `markdowns` directory visit the [cloud.gov `markdowns`](https://github.com/18F/cg-compliance/tree/master/markdowns).

## View locally

Requires [NodeJS](https://nodejs.org/). After running the steps above,

1. Install the gitbook CLI

    ```bash
    npm install -g gitbook-cli
    ```

1. Navigate to the `exports` directory

    ```bash
    cd exports
    ```

1. Serve the gitbook site locally

    ```bash
    gitbook serve
    ```

1. Open the site: http://localhost:4000

After making any edits, view the changes by running

```bash
compliance-masonry get && compliance-masonry docs gitbook <certification>
```

## Export as a PDF

1. Following [the steps above](#creating-gitbook-documentation)
1. Navigate to the `exports` directory

    ```bash
    cd exports
    ```

1. Follow [these instructions](http://toolchain.gitbook.com/ebook.html)
