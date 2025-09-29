# stringcase

stringcase is a go package for converting string case to various cases (e.g. sbake case or camel case)

## Example

| Function                                        | Result    |
|-------------------------------------------------|-----------|
| ToSnakeCase("testCase")                         | test_case |
| ToSnakeCaseWithIgnore("test-case", "-")         | test-case |
| ToScreamingSnakeCase("testCase")                | TEST_CASE |
| ToKebabCase("testCase")                         | test-case |
| ToScreamingKebabCase("testCase")                | TEST-CASE |
| ToDelimitedCase("testCase", "-")                | test-case |
| ToScreamingDelimitedCase("testCase", "-", true) | TEST-CASE |

## Install

```Shell
go get -u github.com/xogas/stringcase
```

## Custom Acronyms for ToCamelCase and ToSnakeCase

```Go
import "github.com/xogas/stringcase"

func init() {
    // results in "Api" using ToCamelCase("API")
    // results in "api" using ToPascalCase("API")
    stringcase.ConfigureAcronyms("API", "api")

    // results in "PostgreSQL" using ToCamelCase("PostgreSQL")
    // results in "postgreSQL" using ToPascalCase("PostgreSQL")
    strcase.ConfigureAcronym("PostgreSQL", "PostgreSQL")
}
```
