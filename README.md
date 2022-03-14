# go-deconstruct

Takes a non-stripped go binary that is built with go module support and creates a go.mod from it.

## Usage

```
./go-deconstruct generate ${BINARY} [-o /path/to/go.mod]
```

## Example
```
./go-deconstruct generate ./go-deconstruct
```
Returns:
```
module github.com/mrueg/go-deconstruct

go 1.17

require (
        github.com/rsc/goversion v1.2.0
        github.com/spf13/cobra v1.4.0
        github.com/spf13/pflag v1.0.5
)
```

## Caveats

* Does not work on go binaries built without module support
* Does not work on stripped binaries
* Does include additional dependencies, that are dependencies from other dependencies (e.g. spf13/cobra depends on spf13/pflag)
* Does not include exclude directives (as they are not included in the go binary)
