package utils

import "github.com/gin-gonic/gin"

func GetDefaultWebErrorTags(c *gin.Context) map[string]string {
	tags := make(map[string]string, 3)
	tags["endpoint"] = c.Request.Host + c.Request.RequestURI
	tags["URL"] = c.Request.RequestURI
	tags["clientIP"] = c.ClientIP()
	tags["method"] = c.Request.Method

	return tags
}
