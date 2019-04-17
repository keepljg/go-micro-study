# Fileinfo-Srv Service

This is the Fileinfo-Srv service

Generated with

```
micro new micro-shop/srv/fileinfo-srv --namespace=go.tutu --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.tutu.srv.fileinfo-srv
- Type: srv
- Alias: fileinfo-srv

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
./fileinfo-srv-srv
```

Build a docker image
```
make docker
```