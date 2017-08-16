## Welcome!

We're so glad you're thinking about contributing to an 18F open source project! If you're unsure about anything, just ask -- or submit the issue or pull request anyway. The worst that can happen is you'll be politely asked to change something. We love all friendly contributions.

We want to ensure a welcoming environment for all of our projects. Our staff follow the [18F Code of Conduct](https://github.com/18F/code-of-conduct/blob/master/code-of-conduct.md) and all contributors should do the same.

We encourage you to read this project's CONTRIBUTING policy (you are here), its [LICENSE](LICENSE.md), and its [README](README.md).

If you have any questions or want to read more, check out the [18F Open Source Policy GitHub repository]( https://github.com/18f/open-source-policy), or just [shoot us an email](mailto:18f@gsa.gov).

See [this page](docs/development.md) for developer documentation.

## Public domain

This project is in the public domain within the United States, and
copyright and related rights in the work worldwide are waived through
the [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).

All contributions to this project will be released under the CC0
dedication. By submitting a pull request, you are agreeing to comply
with this waiver of copyright interest.

## Development

### Using the supplied `Makefile`

The `Makefile` is adapted from [a standard Go Makefile](https://github.com/vincentbernat/hellogopher/tree/feature/glide). Want to use it? Then follow these steps:

1. _Directory Structure._ The `Makefile` assumes you have a ["standard" structure](https://github.com/golang/go/wiki/GithubCodeLayout) when you pull down the software. For example:

    ```bash
    [top-level-dir]
      -> git
         -> src
            -> github.com
               -> opencontrol
                  -> compliance-masonry
    ```

    The advantage to this structure is that it works very well with any Git project, from any workspace.
1. _Install Go._ Strongly recommend *not* to install Go directly but instead to use [Go Version Manager (gvm)](https://github.com/vincentbernat/hellogopher/tree/feature/glide). As an example, selecting a specific Go version is as easy as:

    ```bash
    gvm use [installed go-version from 'gvm list']
    ```
1. _Install <code>make</code>_. See [GNU Make](https://www.gnu.org/software/make/) for details on what `make` is and can do.

    `yum`-based systems (RHEL / CentOS / etc.):

    ```bash
    sudo yum install make
    ```

    `apt`-based systems (Ubuntu / etc.). The below actually installs lots more than just `make`:

    ```bash
    sudo apt-get install build-essential
    ```

    How about Mac? First, install [Homebrew](https://brew.sh). Then:

    ```bash
    brew install make
    ```
1. _OPTIONAL: Install Debugger._ Unfortunately, software does at times behave oddly. The debugger used by this `Makefile` is [Delve (dlv)](https://github.com/derekparker/delve). It can take a little effort to get the debugger installed and working, but the results are far worth it!
1. _One-Time Setup._ Just like in the Manual Instructions above, the code dependencies must be installed. There is a simple target for this:

    ```bash
    make depend
    ```

    The above will pull down all the required dependencies and will not need to be run again.
1. _Common Targets_. Invoke a target by using:

    ```bash
    make [target-name]
    ```

    The common targets include:
    * `all` - Performs a `build`
    * `build` - Builds the source code and places `compliance-masonry` binary into the `./bin` folder (excluded via `.gitignore`)
    * `debug` - Invoke `dlv` _(did you install it above?)_ and invoke it with any `DEBUG_OPTIONS` defined on the command line (example below).
    * `clean` - Simply removes the `compliance-masonry` binary from the `./bin` folder
    * `rebuild` - Invokes `clean` and `build`.
    * `test` - Runs the tests.
    * `lint` - Checks to see if the Go code is properly formatted. (If you want to contribute to the project, use this target; you will need to make sure your code follows accepted standards.)
1. _Examples_.

    * Build the CLI:

        ```bash
        make build
        ```

    * Debug the CLI using `dlv`, running the `export` command. For this to work, let's assume a simple layout as shown below:

        ```bash
        ~/myproj/myapp
          -> opencontrols (output from 'compliance-masonry get')
          -> output
             -> (empty folder; will be filled by 'compliance-masonry export')
        ```

        (Note we do not reference an SSP Template; the `compliance-masonry export` command simply processes the output from `compliance-masonry get`.)

        Given the above, invoke the debugger by using something like the following (the numerous variables are used simply to keep the text formatted nicely in the browser).

        ```bash
        MY_DEBUG_OPTION='--debug'
        MY_DIR="$HOME/myproj/myapp"
        MY_OPENCONTROLS="$MY_DIR/opencontrols"
        MY_LONG_OPTS='--destination - --format json --flatten --infer-keys --docxtemplater'
        MY_OPTIONS="$MY_DEBUG_OPTION export --opencontrols '$MY_OPENCONTROLS' $MY_LONG_OPTS FedRAMP-moderate"
        make debug DEBUG_OPTIONS="$MY_OPTIONS"
        ```
 
        _(Use `compliance-masonry export --help` for more details on what the specific options perform.)_

        If all goes well, you should get a debugger prompt and be able to execute debugger commands:
        
        ```
        Type 'help' for list of commands.
        (dlv) b export.exportJSON
        Breakpoint 1 set at 0x138fcfb for github.com/opencontrol/compliance-masonry/commands/export.exportJSON() ./commands/export/export.go:22
        (dlv) c
        2017/08/16 12:07:01 Running with debug
        > github.com/opencontrol/compliance-masonry/commands/export.exportJSON() ./commands/export/export.go:22 (hits goroutine(1):1 total:1) (PC: 0x138fcfb)
            17:
            18:  ////////////////////////////////////////////////////////////////////////
            19:  // Package functions
            20:
            21:  // exportJSON - JSON output
        =>  22:  func exportJSON(config *Config, workspace common.Workspace, output *exportOutput, writer io.Writer) []error {
            23:    // result
            24:    var errors []error
            25:
            26:    // work vars
            27:    var byteSlice []byte
        (dlv) n
        > github.com/opencontrol/compliance-masonry/commands/export.exportJSON() ./commands/export/export.go:24 (PC: 0x138fd36)
            19:  // Package functions
            20:
            21:  // exportJSON - JSON output
            22:  func exportJSON(config *Config, workspace common.Workspace, output *exportOutput, writer io.Writer) []error {
            23:    // result
        =>  24:    var errors []error
            25:
            26:    // work vars
            27:    var byteSlice []byte
            28:    var err error
            29:
        (dlv) p config
        *github.com/opencontrol/compliance-masonry/commands/export.Config {
          Debug: true,
          Certification: "FedRAMP-moderate",
          OpencontrolDir: "/Users/l.abruce/myproj/myapp/opencontrols",
          DestinationFile: "-",
          OutputFormat: 0,
          Flatten: true,
          InferKeys: true,
          Docxtemplater: true,
          KeySeparator: ":",}
        (dlv) q
        ```
Enjoy working with the debugger!

## Resources

### Tools

* [XML Tree Chrome extension](https://chrome.google.com/webstore/detail/xml-tree/gbammbheopgpmaagmckhpjbfgdfkpadb)
* [XML Viewer Chrome extension](https://chrome.google.com/webstore/detail/xv-%E2%80%94-xml-viewer/eeocglpgjdpaefaedpblffpeebgmgddk?hl=en)

### Reference

* [WordprocessingML information](http://officeopenxml.com/anatomyofOOXML.php)
* [Structure of a WordprocessingML document](https://msdn.microsoft.com/en-us/library/office/gg278308.aspx)
