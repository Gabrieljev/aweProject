package member

import (
	"bytes"
	"encoding/json"
	"github.com/geb/aweproj/book-store/shared/dto"
	"github.com/geb/aweproj/book-store/shared/dto/member"
	"io"
	"io/ioutil"
	"net/http"
)

func (s *service) Login(username, password string) (str *string, err error) {

	body, _ := json.Marshal(member.UserReq{
		Username: username,
		Password: password,
	})

	httpResponse, err := s.SharedHolder.WebClient.Outbound.GetToken().Post(s.SharedHolder.WebClient.Outbound.Path().Token,
		bytes.NewBuffer(body), http.Header{
			"Content-Type": []string{"application/json"},
		})

	if err != nil && httpResponse == nil {
		return nil, err
	}

	defer func() {
		if httpResponse != nil && httpResponse.Body != nil {
			_, _ = io.Copy(ioutil.Discard, httpResponse.Body)

			_ = httpResponse.Body.Close()
		}
	}()
	var res = dto.BaseResponseDto{}
	bodyRes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyRes, &res)
	if err != nil {
		return nil, err
	}
	tkn := res.Data.(string)
	str = &tkn
	return
}
