package telepher
import (
  "telepher/types"
  "net/http"

  )
  
  
func NewBot(token string,settings *Settings) (*Gobot,error) {

  telepher := &Bot{
    Token: token,
    Update: make(chan types.Update),
    Handlers: make(map[int][]Handler),
    Commands: make(map[string][]Commands),
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
     }
    
    
    telepher.BaseURL = settings.BaseURL
    telepher.BaseFileURL = settings.BaseFileURL
    telepher.Client = settings.Client
    telepher.ParseMode = settings.ParseMode   

  }else{

  telepher.BaseURL = DefaultBaseURL
  telepher.BaseFileURL =  DefaultBaseFileURL
  telepher.Client =  DefaultClient
  telepher.ParseMode = DefaultParseMode
  }
  
  me, err := telepher.GetMe()
	if err != nil {
		return nil, err
	}

  telepher.Self = me
  return telepher , nil
  
  }
 


type Bot struct{
  Token string
  BaseURL string
  BaseFileURL string
  Handlers map[int][]Handler
  Commands map[string][]Commands
  Client      *http.Client
  ParseMode string
  Update chan types.Update
  Self          *types.User
}

type Settings struct{
  BaseURL string
  BaseFileURL string
  Client *http.Client
  ParseMode string
}

/*
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
*/
