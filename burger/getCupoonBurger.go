package burger

import (
	"encoding/json"
	"freeEda/burger/model"
	"io/ioutil"
	"net/http"
)

func GetBurgerCupoons(url string) (model.BurgerCupoons, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.BurgerCupoons{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.BurgerCupoons{}, err
	}
	decodedJson := model.BurgerCupoons{}

	err = json.Unmarshal(body, &decodedJson)
	if err != nil {
		return model.BurgerCupoons{}, err
	}

	return decodedJson, nil
}
