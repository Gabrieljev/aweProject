package healthcheck

import (
	"context"
	"github.com/pkg/errors"

	healthcheck_dto "github.com/geb/aweproj/book-store/shared/dto/healthcheck"
)

func (s *service) Check(ctx context.Context, request healthcheck_dto.CheckRequestDto) (*healthcheck_dto.CheckResponseDto, error) {

	res, err := s.ApplicationHolder.HealthCheckService.Check(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed Check")
	}

	return res, nil

}
