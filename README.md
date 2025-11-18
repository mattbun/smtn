# smtn

smtn is a [SMTP](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol) server that forwards any messages it receives to one or more notification services.

It uses [shoutrrr](https://github.com/containrrr/shoutrrr) under the hood so it supports any of [these services](https://containrrr.dev/shoutrrr/services/overview/), including webhooks!

## Usage

smtn can be configured via command line flag or by environment variable:

```
NAME:
   smtn - Run a SMTP server that forwards any messages it receives to one or more notification services

USAGE:
   smtn [global options]

GLOBAL OPTIONS:
   --listen-addr string, -l string                                                Address that the SMTP server should listen on (default: "127.0.0.1") [$LISTEN_ADDR]
   --notification-url string, -n string [ --notification-url string, -n string ]  Shoutrrr notification url(s) [$NOTIFICATION_URL]
   --port int, -p int                                                             Port that the SMTP server should listen on (default: 25) [$PORT]
   --verbose, -v                                                                  Enable verbose logging [$VERBOSE]
   --help, -h                                                                     show help
```

## Getting Started

First off, figure out your notification URL based on [shoutrrr's docs](https://containrrr.dev/shoutrrr/services/overview/).

### Binaries

Binaries can be downloaded from [Releases](https://github.com/mattbun/smtn/releases).

```shell
smtn -n pushover://shoutrrr:{api-token}@{user-token}/
```

### Docker

Docker images are published to Github Container Registry as `ghcr.io/mattbun/smtn`. You can find a list of images [here](https://github.com/mattbun/smtn/pkgs/container/smtn).

Example:

```shell
docker run --rm -it \
    -e "NOTIFICATION_URL=pushover://shoutrrr:{api-token}@{user-token}" \
    -e "LISTEN_ADDR=0.0.0.0" \
    -p "25:25" \
    ghcr.io/mattbun/smtn:latest
```
