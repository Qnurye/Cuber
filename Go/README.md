# Rubik's Cube Solver

## Introduction

This Go program is designed to send commands to an Arduino microcontroller via serial communication.
It utilizes a
Rubik's Cube formula parser to convert a sequence of cube rotations into corresponding commands sent over the serial
port.

## Directory Structure

The project directory structure is organized as follows:

```plaintext
├── build
├── config
│   ├── cmd.json
│   ├── delay.json
│   └── formula.json
├── go.mod
├── go.sum
├── main.go
├── pkg
│   ├── config
│   │   ├── cmd.go
│   │   ├── delay.go
│   │   └── formula.go
│   └── rubiksCube
│       ├── parser.go
│       └── README.md
├── README.md
├── LICENSE
└── test
    ├── cmd_test.go
    ├── delay_test.go
    ├── formula_test.go
    └── parser_test.go
```

- `build`: Contains build artifacts.
- `config`: Stores configuration files.
- `go.mod` and `go.sum`: Go module files.
- `main.go`: The main entry point of the program.
- `pkg`: Package directory containing reusable code.
    - `config`: Configuration-related functionality.
    - `rubiksCube`: Rubik's Cube formula parser and related documentation.
- `test`: Unit tests for the project.

## Getting Started

1. Install Go on your system.
2. Clone the repository:

   ```bash
   git clone https://github.com/Qnurye/Cuber.git
   ```

3. Navigate to the project directory:

   ```bash
   cd Cuber/Go
   ```

4. Install dependencies:

   ```bash
   go mod download
   ```

5. Build the project:

   ```bash
   go build
   ```

6. Run the executable:

   ```bash
   ./main
   ```

## Configuration

The `config` directory contains three JSON files:

- `cmd.json`: Command configuration.
- `delay.json`: Delay configuration.
- `formula.json`: Formula configuration.

Edit these files to customize the behavior of the Rubik's Cube solver.

## Usage

The program reads a Rubik's Cube formula from the `formula.json` file and sends corresponding commands to the Arduino
over serial communication.

## Testing

Unit tests for the project are located in the `test` directory. Run tests using the following command:

```bash
go test ./...
```

## Contributing

If you would like to contribute to this project, please follow the standard GitHub fork and pull request workflow.

## License

This project is licensed under the [MIT License](LICENSE).
