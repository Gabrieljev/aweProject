package dto

const (
	ServiceName     = "aweProj"
)

type (

	MandatoryRequestV2Dto struct {
		Username      string `json:"X-Username" validate:"required"`
		Language      string `json:"Accept-Language" validate:"required"`

	}

	BaseResponseDto struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Errors     []string    `json:"errors"`
		ServerTime int64       `json:"serverTime"`
	}

)
