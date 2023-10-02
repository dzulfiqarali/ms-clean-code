package shared

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Failed(c *gin.Context, applies ...func(b *Base)) {
	baseResponse := Failure()
	for _, apply := range applies {
		apply(baseResponse)
	}

	respondFailde(c, baseResponse)
}

func Success(c *gin.Context, data interface{}, applies ...func(b *Base)) {
	baseResponse := Successful()
	for _, apply := range applies {
		apply(baseResponse)
	}

	respondSuccess(c, data, baseResponse)
}

// ErrorMessage ...
func respondFailde(c *gin.Context, baseResponse *Base) {
	c.JSON(baseResponse.StatusCode, baseResponse)
}

func respondSuccess(c *gin.Context, data interface{}, baseResponse *Base) {
	marshal, err := json.Marshal(data)

	if err != nil {
		return
	}
	result := toMapResponse(marshal, baseResponse)
	c.JSON(baseResponse.StatusCode, result)
}

func toMapResponse(data []byte, baseResponse *Base) (res map[string]interface{}) {
	var respMap map[string]interface{}
	json.Unmarshal(data, &respMap)

	var baseResp map[string]interface{}
	baseRespByte, err := json.Marshal(baseResponse)
	if err != nil {
		return
	}
	json.Unmarshal(baseRespByte, &baseResp)

	for key, value := range baseResp {
		respMap[key] = value
	}

	res = respMap

	return
}
