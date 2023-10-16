package handlers

import (
	"encoding/json"
	"shortener-service/pubsub/kafka"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

// createComment handler
func CreateComment(c *gin.Context) {

	// Instantiate new Message struct
	cmt := new(Comment)

	// //  Parse body into comment struct
	// if err := c.BodyParser(cmt); err != nil {
	// 	log.Println(err)
	// 	c.Status(400).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": err,
	// 	})
	// 	return err
	// }

	if err := c.BindJSON(&cmt); err != nil {
		logrus.Error(err)
		return
	}

	// convert body into bytes and send it to kafka
	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		logrus.Error(err)
	}

	err = kafka.PushCommentToQueue("comments", cmtInBytes)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error creating product",
		})
		return
	}

	// Return Comment in JSON format
	c.JSON(200, gin.H{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})

}
