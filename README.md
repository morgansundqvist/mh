# mh

`mh` is a command-line tool for interacting with APIs. It supports basic HTTP methods like GET, POST, PUT, and DELETE, and allows you to configure the root URL for your API.

## Features

- Initialize configuration with a root URL
- Perform GET, POST, PUT, and DELETE requests
- Automatically format JSON responses
- Color-coded HTTP status codes

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/morgansundqvist/mh.git
    cd mh
    ```

2. Build the project:
    ```sh
    go build -o mh ./cmd/cli
    ```

## Usage

### Initialize Configuration

Before using the tool, you need to initialize the configuration with the root URL of your API:

```sh
./mh init
```

This will create a `.mh` file in your home directory with the following content:

```json
{
    "root": ""
}
```

You can then set the root URL by editing the file:

```json
{
    "root": "https://api.example.com"
}
```

### Perform Requests

You can now perform requests to your API using the following syntax:

```sh
./mh <method> <path> [body_args...]
```

For example, to perform a GET request to `/users`:

```sh
./mh GET /users
```

Or to perform a POST request to `/users` with a JSON body:

```sh
./mh POST /users name=John age=30 is_active=true
```

