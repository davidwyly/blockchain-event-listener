# Ethereum Event Listener (Proof-of-Concept)

## Description
This project is an Ethereum Event Listener written in Go and serves as a proof of concept. It connects to an Ethereum client using RPC, subscribes to specific contract event logs, and processes these events as they occur. As a proof of concept, it will not cover all edge cases or be suited for production use without further development and testing.

## Example Output

```
2024/01/19 13:45:55 Successfully connected to Ethereum client: wss://polygon.api.onfinality.io/public-ws
2024/01/19 13:45:55 Subscription successful for signature hash: 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
2024/01/19 13:45:55 Subscription successful for signature hash: 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925
2024/01/19 13:46:35 Event: Transfer
2024/01/19 13:46:35 address1: 0xdBcdBD5e1cb36f2c7497B0EB55a847354dEbc51B
2024/01/19 13:46:35 address2: 0x74dD45dd579caD749f9381D6227e7e02277C944B
2024/01/19 13:46:35 value: 4612499812901264
2024/01/19 13:46:35 Event: Transfer
2024/01/19 13:46:35 address1: 0xdBcdBD5e1cb36f2c7497B0EB55a847354dEbc51B
2024/01/19 13:46:35 address2: 0xC742516Bbfd161640c8c0521D8d20b7eE33aF82A
2024/01/19 13:46:35 value: 149137493950474208
```

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

## Disclaimer

This project is provided as a proof of concept and is intended for educational and experimental purposes only. The author(s)/contributor(s) of this project make no guarantees and bear no responsibility for any damage, loss of data, financial losses, or any other types of losses resulting from the direct or indirect use of this software.

This project is not intended for use in a production environment and has not been tested for such use. It may contain bugs, vulnerabilities, or other issues that could be harmful if used in real-world applications. Users are advised to use this software with caution and at their own risk.

By using this software, you acknowledge and agree to this disclaimer. If you do not agree, do not use the software.
