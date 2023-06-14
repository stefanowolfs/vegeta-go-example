# Vegeta Go Example

## Description

This is a very simple project intended to serve as an example on how to test an endpoint using [vegeta](https://github.com/tsenart/vegeta) as a lib in golang.

## Parameters

There are some constants you can change inside main file:
#### `-frequency`
Specifies the number of requests per `n` times.
#### `-frequencyPer`
Specifies the request frequency `n` time-box.
#### `-duration`
Specifies the amount of time the test will be executing.
#### `-targetUrl`
The URL used to make the test calls
#### `-httpMethod`
The HTTP verb used to make the test calls

## How to run

To execute simply run:

```bash
go run .main.go
```