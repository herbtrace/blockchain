package handlers

import (
	"net/http"

	"qrcode-api/models"

	"github.com/gin-gonic/gin"
)

// HandleQrCodeData binds incoming JSON to QrCodeData and returns it
func HandleQrCodeData(c *gin.Context) {
	var qrData models.QrCodeData

	if err := c.ShouldBindJSON(&qrData); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "QrCodeData received successfully",
		Data:    qrData,
	})
}
