package shared

import "github.com/gin-gonic/gin"

func Failed(c *gin.Context, applies ...func(b *Base)) {
	baseResponse := Failure()
	for _, apply := range applies {
		apply(baseResponse)
	}

	respond(c, baseResponse)
}

func Success(c *gin.Context, applies ...func(b *Base)) {
	baseResponse := Successful()
	for _, apply := range applies {
		apply(baseResponse)
	}

	respond(c, baseResponse)
}

// ErrorMessage ...
func respond(c *gin.Context, baseResponse *Base) {
	c.JSON(baseResponse.StatusCode, baseResponse)
}
