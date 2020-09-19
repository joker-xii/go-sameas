package sameas

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
)

func (self *QueryClient) getRawData(queryType, srcType, src string) ([]byte, error) {
	reqUrl := self.ServiceURL + queryType + srcType + html.EscapeString(src)

	resp, err := self.Client.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (self *QueryClient) getResult(srcType, src string) (*Result, error) {
	data, err := self.getRawData(string(Json), srcType, src)
	if err != nil {
		return nil, err
	}
	var result []Result
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	if result == nil || len(result) < 1 {
		return nil, errors.New("parse response error: data " + string(data))
	}
	return &result[0], nil
}

func (self *QueryClient) GetRawByWord(queryType RawQueryType, source string) ([]byte, error) {
	return self.getRawData(string(queryType), queryWord, source)
}

func (self *QueryClient) GetRaw(queryType RawQueryType, source string) ([]byte, error) {
	return self.getRawData(string(queryType), queryURL, source)
}

func (self *QueryClient) GetResult(source string) (*Result, error) {
	return self.getResult(queryURL, source)
}

func (self *QueryClient) GetResultByWord(source string) (*Result, error) {
	return self.getResult(queryWord, source)
}
