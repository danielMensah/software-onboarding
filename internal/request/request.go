package request

import (
	"io/ioutil"
	"net/http"
)

type Service struct {

}

func (s *Service) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return nil, err
	}

	return body, nil
}
