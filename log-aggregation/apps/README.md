# apps

## Application Endpoints

|Endpoint|Feature|
| :----: | :---- |
|GET /|Return HTTP 200 with application version.<br><br>e.g. `{"version": "0.1.0"}`|
|GET /http/<http_status_code>|Return the HTTP status code designated by `<http_status_code>`. Returns HTTP 400 - Bad Request when `<http_status_code>` is invalid.<br><br>e.g. `GET /http/503` returns HTTP 503 with payload `{"status_code": 503, "description": "Service Unavailable"}`|
|GET  /random|Return the HTTP status code designated by `<http_status_code>`.<br><br>e.g. `/http/503` returns HTTP 503 with payload `{"status_code": 503, "description": "Service Unavailable"}`<br><br> e.g. `/http/foobar` returns HTTP 400 with payload `{"status_code": 400, "description": "Bad Request"}`|
|GET /exception|Raises an unhandled exception. Returns HTTP 503 - Internal Server Error along with web server specific payload.|