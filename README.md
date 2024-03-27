# Web sink

A simple program used to listen to port 3001 (hardcoded, change it in the code if needed, I might update it in the future). It receives http requests and prints them to stdout. It is useful for debugging of software making requests to some external endpoints (in local environment).

## Usage

You have to have Go installed on your machine. If you don't have it, you can download it from [here](https://golang.org/dl/).

### Start the server

```bash
go run main.go
```

Optionally pass the initial response code (used for all requests' responses) as an argument (200 is the default):

```bash
go run main.go 403
```

### While the server is running

To clear the console type in `clear`

To change the response code for subsequent requests type in `setres <response_code>`, ex: `setres 403`
