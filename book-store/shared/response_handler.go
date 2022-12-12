package shared

import (
	"github.com/geb/aweproj/book-store/shared/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetMandatoryRequestV2(ec echo.Context) dto.MandatoryRequestV2Dto {
	var (
		header = ec.Request().Header
	)

	return dto.MandatoryRequestV2Dto{

		Username: GetHeaderV2OrDefault(header, "X-Username", ""),
		Language: GetHeaderV2OrDefault(header, "Accept-Language", ""),
	}
}

func Response(ec echo.Context, name string, data interface{}, err error) error {
	if err == nil {
		successResponse := New(SUCCESS, nil)
		return ec.Render(successResponse.GetHTTPStatus(), name, data)
	}
	if s, ok := err.(ErrorStandard); ok {
		return ec.JSON(s.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       s.GetCode(),
			Message:    s.GetMessage(),
			Errors:     s.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	} else {
		errResponse := New(SYSTEM_ERROR, err)
		return ec.JSON(errResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       errResponse.GetCode(),
			Message:    errResponse.GetMessage(),
			Errors:     errResponse.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
}

func ResponseJson(ec echo.Context, data interface{}, err error) error {
	if err == nil {
		successResponse := New(SUCCESS, nil)
		return ec.JSON(successResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       successResponse.GetCode(),
			Message:    successResponse.GetMessage(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
	if s, ok := err.(ErrorStandard); ok {
		return ec.JSON(s.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       s.GetCode(),
			Message:    s.GetMessage(),
			Errors:     s.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	} else {
		errResponse := New(SYSTEM_ERROR, err)
		return ec.JSON(errResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       errResponse.GetCode(),
			Message:    errResponse.GetMessage(),
			Errors:     errResponse.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
}

func GetHeaderV2OrDefault(header http.Header, key, value string) string {
	var (
		val = header.Get(key)
	)

	if len(val) == 0 {
		return value
	}

	return val
}
