package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type WebhookPayload struct {
	Event          string    `json:"event"`
	Type           string    `json:"type"`
	AccountName    string    `json:"account_name"`
	AccountNumber  string    `json:"account_number"`
	BankID         int       `json:"bank_id"`
	BankName       string    `json:"bank_name"`
	Amount         string    `json:"amount"`
	Charge         string    `json:"charge"`
	Currency       string    `json:"currency"`
	Status         string    `json:"status"`
	Reference      string    `json:"reference"`
	ChapaReference string    `json:"chapa_reference"`
	BankReference  string    `json:"bank_reference"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type WebhookHandler interface {
	HandleWebHook(payload WebhookPayload) error
}

type WebhookModule struct {
	router           *gin.Engine
	handler          WebhookHandler
	handlerParameter string
	port             string
	secretKey        string
}

func NewChapaWebHook(handler WebhookHandler, port, handlerParameter, secretKey string) *WebhookModule {
	r := gin.Default()

	return &WebhookModule{
		router:           r,
		handler:          handler,
		port:             port,
		handlerParameter: handlerParameter,
		secretKey:        secretKey,
	}
}

func (wm *WebhookModule) verifySignature(c *gin.Context, body []byte) bool {
	signature := c.GetHeader("Chapa-Signature")
	if signature == "" {
		return false
	}

	h := hmac.New(sha256.New, []byte(wm.secretKey))
	h.Write(body)
	computedSignature := hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(computedSignature), []byte(signature))
}

func (wm *WebhookModule) setupRoutes() {
	wm.router.POST(fmt.Sprintf("/%s", wm.handlerParameter), func(c *gin.Context) {

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "Failed to read request body"})
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(body))

		if !wm.verifySignature(c, body) {
			c.JSON(401, gin.H{"error": "Invalid signature"})
			return
		}

		var payload WebhookPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": "Invalid payload"})
			return
		}

		if err := wm.handler.HandleWebHook(payload); err != nil {
			c.JSON(500, gin.H{"error": "Failed to process webhook"})
			return
		}

		c.JSON(200, gin.H{"status": "success"})
	})
}

func (wm *WebhookModule) Run() error {
	wm.setupRoutes()
	return wm.router.Run(wm.port)
}
