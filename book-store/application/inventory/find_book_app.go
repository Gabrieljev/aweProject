package inventory

import (
	"encoding/json"
	"fmt"
	"github.com/geb/aweproj/book-store/shared/dto"
	"github.com/geb/aweproj/book-store/shared/dto/inventory"
	"io"
	"io/ioutil"
	"net/http"
)

func (s *service) FindBookByPubId(pubId int, token string) (bookDto []inventory.BookDto, err error) {
	url := fmt.Sprintf("%s%d", s.SharedHolder.WebClient.Outbound.Path().FindBookByPubId, pubId)
	httpResponse, err := s.SharedHolder.WebClient.Outbound.FindBookByPubId().Get(url,
		http.Header{
			"Content-Type": []string{"application/json"},
			"X-Token":      []string{token},
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

	resByte, _ := json.Marshal(res.Data)

	err = json.Unmarshal(resByte, &bookDto)
	if err != nil {
		return nil, err
	}

	return bookDto, nil
}
