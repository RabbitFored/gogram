package telepher

import (
  "net/http"
  "strings"
  "unicode"
)

func Bot(token string,settings *Settings) (*Telepher,error) {
    telepher := &Telepher{
    Token: token,
    Update: make(chan Update),
  
  }
  
  if settings != nil{
    if settings.BaseURL == ""{
      settings.BaseURL = DefaultBaseURL
    }
    if settings.BaseFileURL == ""{
      settings.BaseFileURL =  DefaultBaseFileURL
    }
    if settings.Client == nil{
      settings.Client = DefaultClient
    }
     if settings.ParseMode == ""{
      settings.ParseMode = DefaultParseMode
     if settings.DisableTokenCheck == false{
       
       isValidToken := CheckToken(token)
      if isValidToken == false{
        return nil,inValidTokenErr
      }
     }
    
    }
    telepher.BaseURL = settings.BaseURL
    telepher.BaseFileURL = settings.BaseFileURL
    telepher.Client = settings.Client
    telepher.ParseMode = settings.ParseMode
    telepher.DisableTokenCheck = settings.DisableTokenCheck


    

  } else {

  telepher.BaseURL = DefaultBaseURL
  telepher.BaseFileURL =  DefaultBaseFileURL
  telepher.Client =  DefaultClient
  telepher.ParseMode = DefaultParseMode
  if telepher.DisableTokenCheck == false{
    isValidToken := CheckToken(token)
      if isValidToken == false{
        return nil,inValidTokenErr
      }
  }
  }
  
  me, err := telepher.GetMe()
	if err != nil {
		return nil, err
	}
  telepher.Self = me
  return telepher , nil
  
  }
 


type Telepher struct{
  Token 		string
  BaseURL 		string
  BaseFileURL 	string
  Client      	*http.Client
  ParseMode 	string
  Update 		chan Update
  DisableTokenCheck   bool
  Self          *User
}

type Settings struct{
  BaseURL 		string
  BaseFileURL 	string
  Client 		*http.Client
  ParseMode 	string
  DisableTokenCheck	  bool
}


func CheckToken(token string)(bool){
	if len(token) > 40 {  
		for _, char := range token {
			if unicode.IsSpace(char){  
						return false
        			}
    	 }
    	if strings.Contains(token, ":") {
      		split := strings.Split(token, ":")
      		if len(split) == 2{
          		for _, c := range split[0] {
                	if !unicode.IsDigit(c) {
            				return false   }
				 }
    			return true
			    } else {
				       return false
			          }
    
    	 } else {
				return false
    		   }
    	 return true
    } else {
    	return false
  			}
}
