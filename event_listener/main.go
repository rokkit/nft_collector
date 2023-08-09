package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	collection "event_listener/collection"
	collection_factory "event_listener/collection_factory"
)

type TokenMinted struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

type Event struct {
	Tx   string `json:"tx"`
	Name string `json:"name"`
}

var events []Event

const COLLECTION_FACTORY_ADDRESS = "0xB4223503Ac3285C73FE5B34d30E2bB8CA5BdC066"
const NFT_COLLECTION_ADDRESS = "0x8915Acdb4c024B16b42f542908c34f0D388fa8e5"

func getHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(events)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func serveEvents() {
	http.HandleFunc("/events", getHello)
	log.Println("Request to http://localhost:3333/events")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// go serveEvents()

	WSS_ETH_ENDPOINT := "wss://polygon-mumbai.g.alchemy.com/v2/MN3dkqJ6G_Y6sD90x6F-hD93sfsfZLRy"

	fmt.Println("Listening to contract events...")

	wss_client, err := ethclient.Dial(WSS_ETH_ENDPOINT)
	if err != nil {
		log.Fatal(err)
	}

	go listenTokenMintedEvents(wss_client)
	go listenCollectionCreatedEvents(wss_client)

	serveEvents()
}

func listenTokenMintedEvents(wss_client *ethclient.Client) {
	address := common.HexToAddress(NFT_COLLECTION_ADDRESS)
	tc, err := collection.NewCollection(address, wss_client)
	if err != nil {
		log.Fatal(err)
	}
	eventChannel := make(chan *collection.CollectionTokenMinted)
	sub, err := tc.WatchTokenMinted(nil, eventChannel, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case t := <-eventChannel:
			fmt.Println(t)
			events = append(events, Event{t.Raw.TxHash.Hex(), "TokenMintend"})
		}
	}
}

func listenCollectionCreatedEvents(wss_client *ethclient.Client) {
	address := common.HexToAddress(COLLECTION_FACTORY_ADDRESS)
	eventChannel := make(chan *collection_factory.CollectionFactoryCollectionCreated)
	tc, err := collection_factory.NewCollectionFactory(address, wss_client)
	if err != nil {
		log.Fatal(err)
	}
	sub, err := tc.WatchCollectionCreated(nil, eventChannel, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case t := <-eventChannel:
			fmt.Println(t)
			events = append(events, Event{t.Raw.TxHash.Hex(), "CollectionCreated"})
		}
	}
}
