# GoURL

## Introduction

This is a simple URL shortener program written in Go using the Gin web framework. This program allows you to create shortened URLs for long website addresses and then redirect users to the original URLs when they visit the shortened ones.

## Getting Started

### Running the Program with Docker

You have the option to run this program within a Docker container for easy setup. Please refer to [the GitHub page](https://github.com/nekogravitycat/GoURL-Server) for detailed instructions on how to set up and run GoURL using Docker.

### Running this program locally

To run this program, you need to have Go installed on your system. You can download and install Go from the official website: https://golang.org/dl/

Once you have Go installed, you can follow these steps to get the URL shortener up and running:

1. Clone the repository:

   ```shell
   git clone https://github.com/nekogravitycat/GoURL.git
   cd GoURL
   ```

2. Install the necessary dependencies:

   ```shell
   go mod download
   ```

3. Start the application:

   ```shell
   go run .
   ```

The application should now be running locally on port 8080.

## Usage

### Shortening URLs

To shorten a URL, visit the admin interface by navigating to `http://0.0.0.0:8080/admin` in your web browser. You will be presented with a form where you can enter the long URL that you want to shorten. Submit the form, and you will receive a shortened URL that you can share.

### Accessing Shortened URLs

To access a shortened URL, simply enter it in your web browser's address bar using the following format:

```
http://0.0.0.0:8080/:shortened
```

Replace `:shortened` with the actual shortened code generated by the application. The program will then redirect you to the original long URL.

## Configuration

You can configure some aspects of the application by modifying the `main.go` file. Here are some of the configuration options:

- Port: You can change the port on which the application listens by modifying the `err := router.Run(":8080")` line. Replace `":8080"` with your desired port number.

## Dependencies

This program relies on the following external libraries:

- [Gin](https://github.com/gin-gonic/gin): A web framework for Go.

You can install these dependencies using the `go mod download` command, as mentioned in the "Getting Started" section.

## Troubleshooting

If you encounter any issues while running the application, please check the following:

- Ensure that Go is properly installed on your system.
- Make sure you have all the required dependencies installed..
