# git-view-work-item

Git add-on for opening work item details in your browser based on the current branch.
It uses a regular expression to extract a work item identifier from the current branch name and inserts it into a URL via Go text templates.

You can configure the add-on's behavior through Git properties. See the following example:

```gitconfig
[git-view-work-item]
    url = "https://dev.azure.com/org/project/_workitems/edit/{{ .Identifier }}"
    regex = "[0-9]+"
```

The following properties are available for you to insert into the URL:

- `Directory`: the current directory name (but not the full path).
- `Identifier`: the identifier extracted from the current branch name.
- `Repository`: the repository name, taken from `remote.origin.url`.
- `Url`: the URL of the repository, taken from `remote.origin.url`.

You can learn more about Go text templates in the [package documentation][go-text-templates].

## Installation

### Homebrew

```
brew install hugginsio/tap/git-vwi
```

### Manual installation

1. Navigate to the [releases page][github-releases] and download the appropriate binary for your system.
2. Copy the `git-vwi` binary to somewhere on your PATH.
3. Run `git vwi -h` in your terminal to validate.

### Building from source

1. Install the latest version of [Go][go-install].
2. Install the latest version of [Task][task-install].
3. Clone the repository.
4. Run `task install`.

<!-- References -->
[go-text-templates]: https://pkg.go.dev/text/template
[github-releases]: https://github.com/hugginsio/git-view-work-item/releases
[go-install]: https://go.dev/dl/
[task-install]: https://taskfile.dev/installation/
