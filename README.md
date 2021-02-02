# SFQ

CLI tool to query different structured files

## Supported structured files

- JSON
- YAML

## Operation

- get
- set

## Query language

### Query value from object by key

Query the same way you would access filed on an object.

#### Example

File:

```JSON
{
    "key": "value"
}
```

Command:

`sfq -o get -q 'key' -f file.json -s json`

Result:

`value`

### Arrays

For array use same syntax as in most programming languages.

- query every value of an array `[]`
- query value from array by index (lets use 0 as index placeholder) `[0]`

#### Example

File:

```JSON
[
    "value 1",
    "value 2",
    "value 3"
]
```

Command:

`sfq -o get -q '[0]' -f file.json -s json`

Result:

`value 1`

#### Example 2

File:

```JSON
[
    "value 1",
    "value 2",
    "value 3"
]
```

Command:

`sfq -o get -q '[]' -f file.json -s json`

Result:

```
value 1
value 2
value 3
```

### More complicated example

File:

```JSON
{
    "key": [
        {
            "arrayKey": "val1"
        },
        {
            "arrayKey": "val2"
        },
        {
            "arrayKey": "val3"
        }
    ]
}
```

Command:

`sfq -o get -q 'key.[].arrayKey' -f file.json -s json`

Result:

```
val1
val2
val3
```

### Any value could be returned from

File:

```JSON
{
    "key": {
        "innerKey1":
    }
}
```

Command:

`sfq -o get -q 'key.[].arrayKey' -f file.json -s json`

Result:

```
val1
val2
val3
```

## Set

Set operation is using the same query syntax, but ends with `=value`.
The whole updated file is returned.
