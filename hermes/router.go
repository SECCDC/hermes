package hermes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}
