package run

import (
	"log"
)

func Run(rpcUrl, faucetPrivateKey string, senderCount, txCount int) {
	generator, err := NewGenerator(rpcUrl, faucetPrivateKey, senderCount, txCount)
	if err != nil {
		log.Fatalf("Failed to create generator: %v", err)
	}

	txsMap, err := generator.GenerateSimple()

	transmitter, err := NewTransmitter(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to create transmitter: %v", err)
	}

	err = transmitter.Broadcast(txsMap)
	if err != nil {
		log.Fatalf("Failed to broadcast transactions: %v", err)
	}
}
