package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing farming events
type SmartContract struct {
	contractapi.Contract
}

// Domain structs
type LatLong struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Address string  `json:"address,omitempty"`
}

type EnvironmentalConditions struct {
	SoilQuality       string  `json:"soil_quality,omitempty"`
	Moisture          float64 `json:"moisture,omitempty"`
	Temperature       float64 `json:"temperature,omitempty"`
	Humidity          float64 `json:"humidity,omitempty"`
	WeatherConditions string  `json:"weather_conditions,omitempty"`
	IrrigationMethod  string  `json:"irrigation_method,omitempty"`
}

type FarmingInputs struct {
	Fertilizers      string `json:"fertilizers,omitempty"`
	PesticidesUsed   string `json:"pesticides_used,omitempty"`
	OrganicCertified bool   `json:"organic_certified"`
}

type PermitCompliance struct {
	PermitID   string `json:"permit_id"`
	PermitType string `json:"permit_type"`
	Issuer     string `json:"issuer"`
	ValidUntil string `json:"valid_until,omitempty"` // Using string instead of *time.Time
}

type CollectionEvent struct {
	BatchID     string                  `json:"batch_id"`
	ActorID     string                  `json:"actor_id"`
	CropID      string                  `json:"crop_id"`
	Location    LatLong                 `json:"location"`
	StartDate   time.Time               `json:"start_date"`
	HarvestDate time.Time               `json:"harvest_date"`
	Environment EnvironmentalConditions `json:"environment,omitempty"`
	Inputs      FarmingInputs           `json:"inputs,omitempty"`
	Permits     []PermitCompliance      `json:"permits,omitempty"`
}

// CreateCollectionEvent stores a new collection event on the ledger
func (s *SmartContract) CreateCollectionEvent(ctx contractapi.TransactionContextInterface, batchID string, eventJSON string) error {
	exists, err := s.EventExists(ctx, batchID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("collection event with batchID %s already exists", batchID)
	}

	// Frontend will send JSON payload, we just validate + save
	var event CollectionEvent
	err = json.Unmarshal([]byte(eventJSON), &event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event JSON: %v", err)
	}

	// Ensure the key matches the event batchID
	if event.BatchID != batchID {
		return fmt.Errorf("batchID mismatch between arg and JSON payload")
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(batchID, bytes)
}

// ReadCollectionEvent retrieves a collection event by batchID
func (s *SmartContract) ReadCollectionEvent(ctx contractapi.TransactionContextInterface, batchID string) (CollectionEvent, error) {
	eventJSON, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return CollectionEvent{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if eventJSON == nil {
		return CollectionEvent{}, fmt.Errorf("collection event %s does not exist", batchID)
	}

	var event CollectionEvent
	err = json.Unmarshal(eventJSON, &event)
	if err != nil {
		return CollectionEvent{}, err
	}
	return event, nil
}

// EventExists checks if a collection event exists by batchID
func (s *SmartContract) EventExists(ctx contractapi.TransactionContextInterface, batchID string) (bool, error) {
	data, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to read world state: %v", err)
	}
	return data != nil, nil
}

// GetAllBlockchainEvents returns all events stored on the blockchain
// This data is immutable and represents the complete state of all farming events
func (s *SmartContract) GetAllBlockchainEvents(ctx contractapi.TransactionContextInterface) ([]CollectionEvent, error) {
	// Get all data from the blockchain ledger
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get state from blockchain: %v", err)
	}
	defer resultsIterator.Close()

	var events []CollectionEvent
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error iterating through blockchain data: %v", err)
		}

		var event CollectionEvent
		err = json.Unmarshal(queryResponse.Value, &event)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal event data: %v", err)
		}

		events = append(events, event)
	}

	return events, nil
}
