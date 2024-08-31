package auth_test

import (
    "context"
    "fmt"
    "net/http"
    "strings"
    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "github.com/gin-gonic/gin"
    "google.golang.org/api/option"
	"time"
)

func InitializeFirebase() (*auth.Client, error) {
    // Use your Firebase Admin SDK credentials JSON file
    opt := option.WithCredentialsFile("api-gateway/firebase-adminsdk.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return nil, fmt.Errorf("error initializing app: %v", err)
    }

    client, err := app.Auth(context.Background())
    if err != nil {
        return nil, fmt.Errorf("error getting Auth client: %v", err)
    }

    return client, nil
}

func FirebaseAuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            c.Abort()
            return
        }

        token, err := authClient.VerifyIDToken(context.Background(), tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

		if token.Expires < time.Now().Unix()*1000 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
            c.Abort()
            return
        }
        // Add the user info to the request context
        c.Set("firebaseUid", token.UID)
		c.Set("firebaseEmail", token.Claims["email"])
        c.Next()
    }
}
