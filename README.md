# WhatsMyIP

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/jere-mie/whatsmyip)


## Overview

This is a simple Go web application that runs an HTTP server, responding to the root path ("/") with the client's IP address. The server listens on a customizable port, allowing you to set the port using the `WHATSMYIP_PORT` environment variable or defaulting to port 8080 if the variable is not set.

## How to Use

### Prerequisites

- Go installed on your machine (you can download it from [golang.org](https://golang.org/dl/))

### Running the Application

1. Clone the repository:

   ```bash
   git clone https://github.com/jere-mie/whatsmyip
   cd whatsmyip
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

   The server will start, and you can access it at [http://localhost:8080](http://localhost:8080) or the specified port.

### Customizing the Port

You can customize the port by setting the `WHATSMYIP_PORT` environment variable. For example:

```bash
WHATSMYIP_PORT=9000 go run main.go
```

This will make the server run on [http://localhost:9000](http://localhost:9000).

## How It Works

The application uses the standard Go `net/http` package to create an HTTP server. When a request is made to the root path ("/"), the server extracts the client's IP address from the request headers and responds with a message containing the IP address.

## Why I Made It

This project was created as a simple example of a Go web application for learning purposes. It demonstrates how to create a basic HTTP server, handle requests, and use environment variables for configuration.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
