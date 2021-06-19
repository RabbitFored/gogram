package telepher

import (
  "net/url"
  "fmt"
  "main/types"
  "encoding/json"
  "os"


)
func (telepher Bot)sendFile(endpoint string,filetype string,file types.InputFile, values url.Values)(*types.Message, error){

switch {
		case file.FileID != "":
      message, err := telepher.sendFileInCloud(endpoint ,filetype, file.FileID,values)
      if err != nil {
            return nil, err 
      }
      return message, nil
			
		case file.FileURL != "":
      message, err := telepher.sendFileInCloud(endpoint ,filetype, file.FileURL,values)
      if err != nil {
            return nil, err 
      }
      return message, nil

		case file.FilePath != "":
		  message, err := telepher.sendFileFromPath(endpoint , filetype, file.FilePath,values)
     if err != nil {
            return nil, err 
      }
return message, nil
		default:
			return nil, fmt.Errorf("telebot: File for field  doesn't exist")
		}


}

func (telepher Bot)sendFileInCloud(endpoint string,filetype string, file string,values url.Values)(*types.Message, error){

  values.Add(filetype,file)


  response ,err := telepher.post(endpoint, values)
    if err != nil {
     return nil, err 
  }
  var message *types.Message
  err = json.Unmarshal(response, &message) 

  if err != nil {
     return nil, err 
  }
  return message, nil
}

func (telepher Bot)sendFileFromPath(endpoint string,filetype string, path string,values url.Values)(*types.Message, error){

  _, err := os.Stat(path)
  if os.IsNotExist(err) {
       return nil, fmt.Errorf("file does not exist")
    }
  
  
 
	response, err := telepher.postWithFile(endpoint ,filetype, path, values)
      if err != nil {
     return nil, err 
  }


  var message *types.Message
  err = json.Unmarshal(response, &message)

  if err != nil {
     return nil, err 
  }
 return message, nil
}
