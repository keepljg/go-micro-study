# Search Service

This is the Search service

Generated with

```
micro new micro-shop/srv/search --namespace=go.tutu --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.tutu.srv.search
- Type: srv
- Alias: search

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
./search-srv
```

Build a docker image
```
make docker
```