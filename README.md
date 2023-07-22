# md-service

Websocket server in Go for receiving FIX market data updates

## Description

`md-service` is an example Go-based service that demonstrates how users can subscribe
to Market Data from reMarkets (https://remarkets.primary.ventures) and receive real-time
updates via WebSocket. Please note that this project is intended for illustrative purposes
only and should not be used in a production environment. It was developed as part of a
GitHub Copilot exploration by the I+D Tech team at Primary.

## Build

To compile the project, use the following command:

```shell
go build ./cmd/...
```

This will create an executable binary for the service.

## Usage

Once the binary is built, you can execute the service using the following command:

```shell
FIX_USERNAME="..." FIX_PASSWORD="..." ./md-service
```

Replace the values of the environment variables `FIX_USERNAME` and `FIX_PASSWORD` 
with your actual reMarkets credentials.

Please be aware that this project is not suitable for production usage. It serves
as an educational and experimental example of utilizing GitHub Copilot for development
purposes.
