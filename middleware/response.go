package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return w.ResponseWriter.Write(data)
}

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer file.Close()

		if _, err := file.WriteString("Nova linha no arquivo de log\n"); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		customWriter := &responseWriter{ResponseWriter: c.Writer, body: make([]byte, 0)}
		c.Writer = customWriter

		c.Next()

		responseData := gin.H{}

		if response, ok := c.Get("Response"); ok {
			responseData["Response"] = response
		}

		startTime := time.Now()
		route := c.FullPath()
		clientIP := c.ClientIP()

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)

		httpStatusCode := c.Writer.Status()

		logData := gin.H{
			"ExecutionTime":  executionTime.String(),
			"Route":          route,
			"HttpStatusCode": httpStatusCode,
			"IP":             clientIP,
		}

		combinedData := gin.H{
			"LogData":      logData,
			"ResponseData": responseData,
		}

		responseJSON, err := json.Marshal(combinedData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Header("Content-Type", "application/json")
		c.Writer.WriteHeader(httpStatusCode)
		c.Writer.Write(responseJSON)
	}
}
