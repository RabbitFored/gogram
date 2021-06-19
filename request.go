package telepher

import (
  "net/http"
  "net/url"
  "mime/multipart"
  "io/ioutil"
  "io"
  "bytes"
  "fmt"
  "os"
  "main/types"
  "encoding/json"

  
  )
 

type apiResponse struct {
	OK          bool                `json:"ok"`
	Result     json.RawMessage      `json:"result"`
	Description string              `json:"description"`
	ErrorCode   int                 `json:"error_code"`
  Parameters  *types.ResponseParameters  `json:"parameters"`
}


func (telepher Bot) post(endpoint string, params url.Values )(json.RawMessage ,error){
  var resp *http.Response
  var err error

  URL := telepher.BaseURL + telepher.Token + "/" + endpoint

if params == nil {
		resp, err = http.Post(URL, "application/x-www-form-urlencoded", nil)
	} else {
	resp, err = http.PostForm(URL,params)
	}
	

    if err != nil {
return nil,err
}
    
  data, err := ioutil.ReadAll(resp.Body)


  if err != nil {
   return nil , err
 }
var response apiResponse
	if err := json.Unmarshal(data,&response); err != nil {
   return nil, err
	} 
  



if !response.OK{
fmt.Println(response)
return nil , fmt.Errorf("telepher : api error %s : %d", response.Description, response.ErrorCode)
}

return response.Result ,nil
}


func (telepher Bot) postWithFile(endpoint string ,filetype string,path string, params url.Values)(json.RawMessage ,error){

  

  URL := telepher.BaseURL + telepher.Token + "/" + endpoint

  
  body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	File, err := os.Open(path)

    if err != nil {
return nil, err
	}
    
    part, err := writer.CreateFormFile(filetype, path)
    if err != nil {
return nil, err
	}
   	_, err = io.Copy(part, File)
	for key:= range params {
		_ = writer.WriteField(key,params.Get(key))
	}	
  
  err = writer.Close()
    if err != nil {
return nil, err
	}

  
  req, err := http.NewRequest("POST", URL, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
    if err != nil {
return nil,err
	}
  resp, err := telepher.Client.Do(req)
  data, err := ioutil.ReadAll(resp.Body)
  
  fmt.Println(string(data))



if err != nil {
   return nil , err
 }
var response apiResponse
	if err := json.Unmarshal(data,&response); err != nil {
   return nil, err
	} 
if !response.OK{
return nil , fmt.Errorf("telepher : api error %d : %s", response.ErrorCode, response.Description)
}

return response.Result ,nil
}



