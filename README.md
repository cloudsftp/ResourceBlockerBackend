# Resource Blocker Backend

## Usage

```
go run main/main.go
```

### Config

Copy `config.example.json` to `config.json`.

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
