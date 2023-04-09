package query

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// GetString 从query中读取一个string
func GetString(c *gin.Context, k string) string {
	return c.Query(k)
}

// GetInt 从query中读取一个int值
func GetInt(c *gin.Context, k string) int64 {
	v := c.Query(k)
	iv, _ := strconv.ParseInt(v, 10, 64)

	return iv
}

// GetFloat 从query中读取一个float值
func GetFloat(c *gin.Context, k string) float64 {
	v := c.Query(k)
	fv, _ := strconv.ParseFloat(v, 64)

	return fv
}

// GetBool 从query中读取一个bool值
func GetBool(c *gin.Context, k string) bool {
	v := c.Query(k)
	if strings.ToUpper(v) == "TRUE" || v == "1" {
		return true
	}

	return false
}
