# Beego-Web Service

This is the Beego-Web service

Generated with

```
micro new micro-shop/web/beego-web --namespace=go.tutu --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.tutu.web.beego-web
- Type: web
- Alias: beego-web

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
./beego-web-web
```

Build a docker image
```
make docker
```