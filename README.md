
# GODO

A versatile and user-friendly command-line tool designed to streamline your most frequently used actions, ranging from simple system commands to executing custom scripts and running applications. With GODO, you can effortlessly automate repetitive tasks, improve productivity, and simplify your workflow.

## Usage

```shell
godo [macro name or action] [--godo-<option>] [extras arguments]
```

**Note:**
`[extras arguments]` will be appended to the end of the string arguments, defined in the configuration file.

### Examples

```shell
# Pings google DNS.
godo ping-google

# Calls a script to clear a database.
godo clear-database --db-url "jdbc:oracle:thin:@//localhost:1521/test_db" --db-pass 12345

# List all configuration in the specified file.
godo --godo-config-file /opt/godo/config.yaml list

# Opens config file with the specified editor (nano).
godo --godo-config-editor nano edit
```

### Actions

- `list`: List all the macros.
- `edit`: Opens the configuration file with the default OS text editor (can be overwritten with options).

### Options

- `--godo-config-file`: Sets the path for the configuration file.
- `--godo-config-editor`: Overrides the default text editor.
- `--godo-config-editor-args`: Passes the arguments for the text editor. Must be split by commas (`--arg1,value1,arg2`).

### Environment Variables

- `GODO_CONFIGURATION_FILE`  Sets the path for the configuration file.
- `GODO_CONFIGURATION_EDITOR` Overrides the default text editor.
- `GODO_CONFIGURATION_EDITOR_ARGS` Passes the arguments for the text editor. Must be split by commas (`"--arg1,value1,arg2"`).

### Configuration File

```yaml
# Add all of your most used actions here
macros:

  - name: ping-google
    executable: ping
    arguments: 8.8.8.8 -c 4
    description: Ping Google DNS Shell

  - name: ping-google-cmd
    executable: cmd.exe
    arguments: /c ping 8.8.8.8
    description: Ping Google DNS Command Prompt

  - name: ping-google-powershell
    executable: powershell.exe
    arguments: ping 8.8.8.8
    description: Ping Google DNS Powershell

  # Scripts execution
  - name: hello-sh
    executable: bash
    arguments: ./examples/hello.sh godo
    description: Hello from Shell

  - name: hello-cmd
    executable: cmd.exe
    arguments: /c .\examples\hello.bat godo
    description: Hello from Command Prompt

  - name: hello-powershell
    executable: powershell.exe
    arguments: -f ./examples/hello.ps1 godo
    description: Hello from Powershell

  - name: hello-git-bash
    executable: C:/Program Files/Git/bin/bash.exe
    arguments: -c C:/Users/elonmusk/Downloads/examples/hello.sh
    description: Hello from Git Bash

  - name: interactive-sh
    executable: bash
    arguments: ./examples/interactive.sh
    description: Interactive hello from Shell

  # Other apps execution
  - name: run-spring-app
    executable: java
    arguments: -jar ./examples/demo-0.0.1-SNAPSHOT.jar
    description: Run my Spring Boot Application
```

## How to Build From Source

### Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/doc/install)

### Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/GabrielJuliao/godo.git
    ```

### Build the Go Application

#### Windows

To build the Go application on Windows, run the following command in Command Prompt or PowerShell:

```bash
go build -o godo.exe
```

This will create an executable binary with a `.exe` extension.

#### macOS and Linux

To build the Go application on macOS or Linux, run the following command in the terminal:

```bash
go build -o godo
```

This will create an executable binary without an extension.

### Execute the Go Application

#### Windows

After building the application, you can execute it on Windows by running:

```bash
godo.exe
```

#### macOS and Linux

After building the application, you can execute it on macOS or Linux by running:

```bash
./godo
```

### Clean Up

If you want to clean up the generated binary, you can run:

```bash
go clean
```

This removes the compiled binary.

### Additional Notes

- Refer to the [official Go documentation](https://golang.org/doc/) for more details on Go programming.