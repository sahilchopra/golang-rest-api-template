package middleware

import (
	"github.com/gin-gonic/gin"
	xss "github.com/sahilchopra/gin-gonic-xss-middleware"
)

func XssMiddleware() gin.HandlerFunc {
	var xssMdlwr xss.XssMw
	return xssMdlwr.RemoveXss()
}
