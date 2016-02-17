# Set

Set is a set implementation for Golang. It includes four implementations described below:

## Implementations

* SimpleSet - Unordered set where the Item itself is used as the unique value
* OrderedSimpleSet - Ordered set where the Item itself is used as the unique value
* HashSet - Unordered set where the Item's Hash method dictates the unique value
* OrderedHashSet - Ordered set where the Item's Hash method dictates the unique value

## Code Generation

Due to the lack of generics in Go a generator tool is provided that can output a wrapper around this library that will accept and return concrete values rather than requiring the user to perform type assertions. The reason for this is that interfaces are used internally to allow arbitrary types to be stored. This tool can be used in conjunction with the Go tool's "go generate" functionality to allow the code to be generated at build time.
