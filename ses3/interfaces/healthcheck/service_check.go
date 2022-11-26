package healthcheck

import (
	"context"
	"github.com/pkg/errors"

	healthcheck_dto "github.com/geb/aweproj/ses3/shared/dto/healthcheck"
)

func (s *service) Check(ctx context.Context, request healthcheck_dto.CheckRequestDto) (*healthcheck_dto.CheckResponseDto, error) {

	res, err := s.ApplicationHolder.HealthcheckService.Check(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed Check")
	}

	return res, nil

}
