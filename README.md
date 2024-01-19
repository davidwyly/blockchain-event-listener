# Ethereum Event Listener

## Description
This project is an Ethereum Event Listener written in Go and serves as a proof of concept. It connects to an Ethereum client using RPC, subscribes to specific contract event logs, and processes these events as they occur. As a proof of concept, it may not cover all edge cases or be suited for production use without further development and testing.

## Installation

To set up this project locally, follow these steps:

1. Clone the repository:

```
git clone https://github.com/davidwyly/blockchain-event-listener
```

2. Navigate to the project directory

```
cd [project directory]
```

3. Install dependencies

```
go get ./...
```

## Usage

To run the Ethereum Event Listener, follow these steps:

1. Ensure you have a valid `config.json` and `abi.json` file in the project directory.

2. Build the project:

```
go build
```

3. Run the compiled binary:

```
go run main.go
```
