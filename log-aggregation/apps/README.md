# apps

This section contains code written in Golang, Java and Python. Each codebase represents a web application which exposes a couple of endpoints as described below.

Each app logs to a specified directory. The logs generated are the subject of the log aggregation strategy demonstrated by this SRE module.

## Application Endpoints

|Endpoint|Feature|
| :----: | :---- |
|/|Return HTTP 200 with application version.<br><br>e.g. `{"version": "0.1.0"}`|
|/http/<http_status_code>|Return the HTTP status code designated by `<http_status_code>`. Returns HTTP 400 - Bad Request when `<http_status_code>` is invalid.<br><br>e.g. `/http/503` returns HTTP 503 with payload `{"code": 503, "message": "Service Unavailable"}`|
|/random|Return the HTTP status code designated by `<http_status_code>`.<br><br>e.g. `/http/503` returns HTTP 503 with payload `{"code": 503, "message": "Service Unavailable"}`<br><br> e.g. `/http/foobar` returns HTTP 400 with payload `{"code": 400, "message": "Bad Request"}`|
|/exception|Raises an unhandled exception. May HTTP 503 - Internal Server Error along with web server specific payload.|

All endpoints are HTTP method agnostic except for requests to the javaapp.

## Before you start

1. Install [golang](https://golang.org/doc/install)
1. Install [mvn](https://maven.apache.org/install.html)

## References

1. HTTP response codes: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status