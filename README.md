# git-view-work-item

Git add-on for opening work item details in your browser based on the current branch.
It uses a regular expression to extract a work item identifier from the current branch name and inserts it into a URL via Go text templates.

You can configure the add-on's behavior through Git properties. See the following example:

```
[git-view-work-item]
    url = "https://dev.azure.com/org/project/_workitems/edit/{{ .Identifier }}"
    regex = "[0-9]+"
```

The following properties are available for you to insert into the URL:

- `Identifier`: the identifier extracted from the current branch name.

You can learn more about Go text templates in the [package documentation][go-text-templates].

<!-- References -->
[go-text-templates]: https://pkg.go.dev/text/template
