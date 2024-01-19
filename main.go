package main

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Config holds the configuration settings
type Config struct {
	WebsocketRPCURL   string   `json:"websocket_rpc_url"`
	ChainID           int      `json:"chain_id"`
	ContractAddress   string   `json:"contract_address"`
	EventSignatures   []string `json:"event_signatures"`
	DiscordWebhookURL string   `json:"discord_webhook_url"`
}

func main() {
	// Load configuration
	var config Config
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// Initialize Ethereum client
	client, err := ethclient.Dial(config.WebsocketRPCURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	log.Println("Successfully connected to Ethereum client:", config.WebsocketRPCURL)

	// Load the smart contract ABI from abi.json
	ABIFile, err := os.ReadFile("abi.json")
	if err != nil {
		log.Fatalf("Failed to read abi.json: %v", err)
	}
	var ABI abi.ABI
	err = json.Unmarshal(ABIFile, &ABI)
	if err != nil {
		log.Fatalf("Failed to parse abi.json: %v", err)
	}

	// Compute the Keccak-256 hashes of the event signatures
	eventSignatures := getEventSignatureHashes(config.EventSignatures)

	httpClient := &http.Client{Timeout: time.Second * 30}

	// Listening to each event signature in a separate goroutine
	for _, sigHash := range eventSignatures {

		// Start a new goroutine for each event signature
		go func(sigHash common.Hash, config Config, client *ethclient.Client, ABI *abi.ABI) {
			query := ethereum.FilterQuery{
				Addresses: []common.Address{common.HexToAddress(config.ContractAddress)},
				Topics:    [][]common.Hash{{sigHash}},
			}

			// Subscribe to the logs
			logs := make(chan types.Log)
			sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				log.Fatalf("Failed to subscribe to logs: %v", err)
			} else {
				log.Println("Subscription successful for signature hash:", sigHash.Hex())
			}

			// Listen for new events
			for {
				select {

				// Handle errors
				case err := <-sub.Err():
					log.Printf("Error in subscription: %v", err)
					log.Fatal(err)

				// Process the received event
				case eventLog := <-logs:
					// Process the event data as needed
					processEvent(eventLog, config.DiscordWebhookURL, httpClient, ABI)
				}
			}
		}(sigHash, config, client, &ABI)
	}

	// Keep the main goroutine alive while other goroutines run
	select {}
}

func getEventSignatureHashes(eventSignatures []string) []common.Hash {
	var sigHashes []common.Hash
	for _, sig := range eventSignatures {
		sigHash := crypto.Keccak256Hash([]byte(sig))
		sigHashes = append(sigHashes, sigHash)
	}
	return sigHashes
}

// processEvent processes the received event and sends a message to Discord
func processEvent(eventLog types.Log, webhookURL string, httpClient *http.Client, ABI *abi.ABI) {

	eventData := make(map[string]interface{})

	event, err := ABI.EventByID(eventLog.Topics[0])
	if err != nil {
		log.Fatalf("Failed to get event by ID: %v", err)
	}

	// Log the event name
	log.Printf("Event: %s", event.Name)

	// Process Topics - assuming they are addresses
	for i, topic := range eventLog.Topics {
		if i == 0 {
			// Skip the first topic as it's the event signature
			continue
		}
		address := common.HexToAddress(topic.Hex()).Hex()
		eventData["address"+strconv.Itoa(i)] = address
	}

	// Process Data - assuming it's a big.Int
	if len(eventLog.Data) > 0 {
		value := new(big.Int)
		value.SetBytes(eventLog.Data)
		eventData["value"] = value
	}

	// Print the eventData map
	for key, value := range eventData {
		log.Printf("%s: %v", key, value)
	}

	// TODO: Send a message to Discord
}
