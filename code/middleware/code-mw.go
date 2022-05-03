package middleware

// package name = test-service
var middlewareHead = `package middleware

import (
	"%s/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)`
