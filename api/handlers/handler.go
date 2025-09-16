package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"qrcode-api/models"

	"github.com/gin-gonic/gin"
)

// CreateCollectionEvent calls the smart contract to create a new collection event
func CreateCollectionEvent(c *gin.Context) {
	var collectionEvent models.CollectionEvent

	if err := c.ShouldBindJSON(&collectionEvent); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	// Generate a unique batch ID from the data
	batchID := collectionEvent.BatchID

	// Convert the data to JSON string for the smart contract
	eventJSON, err := json.Marshal(collectionEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to marshal event data",
			Error:   err.Error(),
		})
		return
	}

	// Execute the smart contract transaction using network.sh
	invokeArgs := fmt.Sprintf(`{"Args":["CreateCollectionEvent","%s","%s"]}`, batchID, strings.ReplaceAll(string(eventJSON), `"`, `\"`))
	cmd := exec.Command("./network.sh", "cc", "invoke",
		"-c", "herbtrace",
		"-ccn", "herbtrace",
		"-ccic", invokeArgs)
	cmd.Dir = "/home/rishi/Desktop/herbtrace/blockchain/test-network"

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to execute smart contract transaction",
			Error:   err.Error() + ": " + string(output),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Collection event created successfully on blockchain",
		Data:    collectionEvent,
	})
}

// GetAllBlockchainEvents queries all transaction data from the blockchain
func GetAllBlockchainEvents(c *gin.Context) {
	// Execute the smart contract query using network.sh
	queryArgs := `{"Args":["GetAllBlockchainEvents"]}`
	cmd := exec.Command("./network.sh", "cc", "query",
		"-c", "herbtrace",
		"-ccn", "herbtrace",
		"-ccqc", queryArgs)
	cmd.Dir = "/home/rishi/Desktop/herbtrace/blockchain/test-network"

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to query blockchain",
			Error:   err.Error() + ": " + string(output),
		})
		return
	}

	// Parse the JSON response from the blockchain
	var events []models.CollectionEvent
	outputStr := strings.TrimSpace(string(output))
	if err := json.Unmarshal([]byte(outputStr), &events); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to parse blockchain response",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Retrieved all blockchain events successfully",
		Data:    events,
	})
}
