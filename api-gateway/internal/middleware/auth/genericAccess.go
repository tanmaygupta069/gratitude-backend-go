package auth_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
    "bytes"
	"github.com/gin-gonic/gin"
)

var fetchedUserId = "abcd"

func GenericAccess() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Read the request body
        body, err := io.ReadAll(c.Request.Body)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
            c.Abort()
            return
        }

        c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

        // Parse the request body to extract the user ID
        var requestData map[string]interface{}
        if err := json.Unmarshal(body, &requestData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid request format :",err.Error())})
            c.Abort()
            return
        }

        incomingUserID := requestData["userId"]

        // Fetch user ID from the database
        if fetchedUserId != incomingUserID {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
            c.Abort()
            return
        }

        // User ID matches, continue to the next handler
        c.Set("userID", fetchedUserId) // Optional: Set userID in context for later use
        c.Next()
    }
}