package util

import (
	"github.com/geb/aweproj/book-store/shared/dto"
	"strings"
)

const (
	ApplicationJson = "application/json"
	AcceptEnCoding  = "accept-Encoding"
	GZIP            = "gzip"
	ContentType     = "content-Type"
	Authorization   = "Authorization"

	// ANDROID Platform
	ANDROID = "android"
	// IOS Platform
	IOS = "ios"
	// MOBILEWEB Platform
	MOBILEWEB = "mobile"
	// DESKTOP Platform
	DESKTOP = "desktop"
	// WEB - Swagger Default Value for ChannelID
	WEB = "web"
)


func IsLogin(mandatory dto.MandatoryRequestV2Dto) bool {
	if strings.ToLower(mandatory.Username) != "" && strings.ToLower(mandatory.Username) != "guest" {
		return true
	}

	return false

}
