package healthcheck

import (
	"github.com/geb/aweproj/ses3/shared/dto"
)

type (
	CheckRequestDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
	}

	CheckResponseDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
	}
)
