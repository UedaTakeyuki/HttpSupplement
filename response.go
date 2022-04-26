package HttpSupplement

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func GetJsonFromReadCloserWithLimitSize (body io.ReadCloser, limit int64) (jsonObj map[string]interface{}, jsonByteArray []byte, err error){
  jsonByteArray, err = ioutil.ReadAll(io.LimitReader(body, limit))
  if err != nil {
    log.Println(err)
		return
	} else {
    jsonObj = make(map[string]interface{})
    json.Unmarshal(jsonByteArray, &jsonObj)
    return
  }
  return
}

func GetJsonFromReadCloser (body io.ReadCloser) (jsonObj map[string]interface{}, jsonByteArray []byte, err error){
  jsonObj, jsonByteArray, err = GetJsonFromReadCloserWithLimitSize(body, 1024)
	return
}
