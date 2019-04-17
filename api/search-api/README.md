# Search-Api Service

This is the Search-Api service

Generated with

```
micro new micro-shop/api/search-api --namespace=go.tutu --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.tutu.api.search-api
- Type: api
- Alias: search-api

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./search-api-api
```

Build a docker image
```
make docker
```