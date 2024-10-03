package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/coding-challenge/api-searching/helpers/respond"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var (
	topic *pubsub.Topic

	// token is used to verify push requests.
	projectId = "brilliant-tower-398410"
	topicName = "deployment_golang_topic"
)

type pushRequest struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

type PubController struct {
}

const maxMessages = 10

func (pubCtrl *PubController) PublishTopic(c *gin.Context) {
	ctx := context.Background()

	pubKey := os.Getenv("PUB_SUB_KEY")
	msgBytes, err := json.Marshal(pubKey)

	client, err := pubsub.NewClient(ctx, projectId, option.WithCredentialsJSON(msgBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	topic = client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !exists {
		log.Printf("Topic %v doesn't exist - creating it", topicName)
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx = context.Background()

	msg := &pushRequest{
		Message: struct {
			Attributes map[string]string
			Data       []byte
			ID         string `json:"message_id"`
		}{
			Attributes: make(map[string]string),
			Data:       []byte{},
			ID:         "data publish",
		},
		Subscription: "gcp-subscribe",
	}

	// Convert your custom structure `msgData` to JSON (if necessary)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal message"})
		return
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: msgBytes,
		Attributes: map[string]string{
			"customAttribute": "attributeValue",
		},
	})
	// Get the result of the Publish call
	serverID, err := result.Get(ctx)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, respond.InternalServerError())
		return
	}
	c.JSON(http.StatusOK, respond.Success(serverID, "Message published."))
}
