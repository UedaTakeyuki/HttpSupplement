package HttpSupplement

func getJsonFromReadCloserWithLimitSize (body io.ReadCloser, limit int) (jsonObj map[string]interface{}, err error){
  buff, err := ioutil.ReadAll(io.LimitReader(resp.Body, limit))
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
}
