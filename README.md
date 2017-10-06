web-payments
============

This image runs a web service in 8080 port, used for show micorservice concept. It comes from rawmind/alpine-base.

## Build

```
docker build -t rawmind/web-payments:<version> .
```

## Versions

- `0.1-1` [(Dockerfile)](https://github.com/rawmind0/web-payments/blob/0.1-1/Dockerfile)


## Usage

```
docker run rawmind/web-payments:<version> 
```

## Example

See [rancher-example][rancher-example], rancher catalog package that runs web-monolith in a cattle environment.

## Microservice uri

- / web-payments app

[rancher-example]: https://github.com/rawmind0/web-payments/tree/master/rancher