package HttpSupplement

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func getJsonFromReadCloserWithLimitSize (body io.ReadCloser, limit int64) (jsonObj map[string]interface{}, err error){
  buff, err := ioutil.ReadAll(io.LimitReader(body, limit))
  if err != nil {
    log.Println(err)
		return
	} else {
    jsonObj = make(map[string]interface{})
    json.Unmarshal(buff, &jsonObj)
    return
  }
  return
}

func getJsonFromReadCloser (body io.ReadCloser) (jsonObj map[string]interface{}, err error){
  jsonObj, err = getJsonFromReadCloserWithLimitSize(body, 1024)
	return
}
