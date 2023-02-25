# Resource Blocker Backend ![version](https://img.shields.io/badge/v0.1.3-blue.svg)

[![Build Develop](https://github.com/cloudsftp/ResourceBlockerBackend/actions/workflows/build.yml/badge.svg?branch=develop)](https://github.com/cloudsftp/ResourceBlockerBackend/actions/workflows/build.yml)

## Usage

```
go run main/main.go
```

### Config

Copy `config.example.json` to `config.json`.
The example has the following content:

```
{
    "port": 5000,
    "resources": {
        "test": {
            "name": "Test Resource",
            "min": 0,
            "max": 3
        }
    }
}
```

The server will listen on port `5000`.

There is only one resource configured.
It has the id `test`.
When requesting the status over the API, use this id.
It's name is `Test Resource` and it's status has to be in the range between `0` and `3`.

### For Developers

```
gow run main/main.go
```

Adds automatic reloads.
If `gow` is not installed:
```
go install github.com/mitranim/gow@latest
```

## API

### `/`

#### GET

Returns all registered resources.
For example:

```
{
    "resources": [
        "bikebox1",
        "garage1"
    ]
}
```

### `/{resource}/`

#### GET

Returns the status of the resource.
For example:

```
{
    "name": "Test Resource",
    "num": 1
}
```

#### POST

To request an update of the resource status.
The request must be in the following format:

```
{
    "delta": -1
}
```

It then returns the updated status of the resource (see above)
