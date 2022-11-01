package wkusno

import (
	"encoding/json"
	"freeEda/wkusno/model"
	"io/ioutil"
	"net/http"
)

func GetWkusnoCupoons(url string) (model.WkusnoCupoons, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.WkusnoCupoons{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.WkusnoCupoons{}, err
	}
	decodedJson := model.WkusnoCupoons{}

	err = json.Unmarshal(body, &decodedJson)
	if err != nil {
		return model.WkusnoCupoons{}, err
	}

	return decodedJson, nil
}
