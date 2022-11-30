package healthcheck

import (
	"context"
	healthcheck_dto "github.com/geb/aweproj/ses3/shared/dto/healthcheck"
)

func (s *service) Check(ctx context.Context, request healthcheck_dto.CheckRequestDto) (*healthcheck_dto.CheckResponseDto, error) {


	return &healthcheck_dto.CheckResponseDto{}, nil
}
