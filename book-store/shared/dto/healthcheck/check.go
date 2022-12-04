package healthcheck

import (
	"github.com/geb/aweproj/book-store/shared/dto"
)

type (
	CheckRequestDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
	}

	CheckResponseDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
	}
)
