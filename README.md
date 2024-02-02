# GODO
## A simple command line tool to execute macros

This README provides instructions on how to build and execute a Go application on multiple platforms.

## Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/doc/install)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/GabrielJuliao/godo.git
    
    cd godo
    ```

2. Navigate to the directory containing the Go source code:

    ```bash
    cd cmd/godo
    ```

## Build the Go Application

### Windows

To build the Go application on Windows, run the following command in Command Prompt or PowerShell:

```bash
go build -o godo.exe
```

This will create an executable binary with a `.exe` extension.

### macOS and Linux

To build the Go application on macOS or Linux, run the following command in the terminal:

```bash
go build -o godo
```

This will create an executable binary without an extension.

## Execute the Go Application

### Windows

After building the application, you can execute it on Windows by running:

```bash
godo.exe
```

### macOS and Linux

After building the application, you can execute it on macOS or Linux by running:

```bash
./godo
```

## Clean Up

If you want to clean up the generated binary, you can run:

```bash
go clean
```

This removes the compiled binary.

## Additional Notes

- Refer to the [official Go documentation](https://golang.org/doc/) for more details on Go programming.