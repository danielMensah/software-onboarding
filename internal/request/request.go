package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string, out interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return err
	}

	unmarshalErr := json.Unmarshal(body, &out)

	if unmarshalErr != nil {
		return err
	}

	return nil
}
