package common

import "github.com/gin-gonic/gin"

func ErrorResponse(errcode, error string) gin.H {
	return gin.H{
		"errcode": errcode,
		"error":   error,
	}
}
