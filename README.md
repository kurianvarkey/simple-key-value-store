# Simple Key-Value Store

## Running the Application

```bash
$ go run cmd/main.go
```

## Running with Docker

```bash
$ docker build -t kvs .
$ docker run -it --rm kvs
```

## About the Application

The application is a simple key-value store that allows you to set, get, delete, and list key-value pairs.
On exit, the application saves the key-value pairs to a json file. The store is structured as an interface and is implemented as a file store and can be extended to other stores in the future.
