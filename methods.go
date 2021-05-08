package telepher
import (
"strconv"
  "net/url"
 "encoding/json"
)


func (telepher Telepher)GetMe()(*User,error){
data ,err :=telepher.post("getMe",nil)

	if err != nil {
    return nil, err
	} 
	var user struct {
		Result User
	}
 	if err := json.Unmarshal(data,&user); err != nil {
   return nil, err
	}  
  return &user.Result , nil
}
func (telepher Telepher)GetUpdates(offset int,limit int)([]Update,error) {

v := url.Values{}
v.Add("offset",strconv.Itoa(offset))
v.Add("limit", strconv.Itoa(limit))

data ,_ :=telepher.post("getUpdates", v)

	var update struct {
		Result []Update
	}

  
	if err := json.Unmarshal(data,&update); err != nil {
   return nil, err
	}  
return update.Result , nil

}
func (telepher Telepher) SendMessage(chat_id string ,text string,option *Options)(*Message,error){
values := url.Values{}
values.Add("chat_id",chat_id)
values.Add("text", text)

if option != nil{
optional := []string{"ParseMode","entities" ,"DisableWebPagePreview","DisableNotification","ReplyToMessageID","AllowSendingWithoutReply","ReplyMarkup" }
telepher.addOptions(optional,values,option)
}
data ,err :=telepher.post("sendMessage", values)

var res struct{ 
Result *Message
}
 err = json.Unmarshal(data, &res) 
    if err != nil {
return nil, err }
return res.Result,nil

}


func (telepher Telepher) SendPhoto(chat_id string,path string){
    
param := map[string]string{
		"chat_id":   chat_id,
	}


telepher.postWithFile("photo",path,param)


}


