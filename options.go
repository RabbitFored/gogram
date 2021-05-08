package telepher

import(
"strconv"
  "net/url"
)
type Options struct{
ParseMode string
entities MessageEntity
DisableWebPagePreview bool
DisableNotification bool
ReplyToMessageID int
AllowSendingWithoutReply bool
ReplyMarkup ReplyMarkup 
}

func (telepher Telepher) addOptions(optional []string,values url.Values,opt *Options){
  if opt != nil{
if optionExist(optional,"ParseMode"){
    if opt.ParseMode == ""{
    opt.ParseMode = telepher.ParseMode
      }
    values.Add("parse_mode",opt.ParseMode)
}
if optionExist(optional,"DisableWebPagePreview"){
      if opt.DisableWebPagePreview == true{
           values.Add("disable_web_page_preview","true")
}
}
if optionExist(optional,"ReplyToMessageID"){
    if opt.ReplyToMessageID != 0 {
values.Add("reply_to_message_id",strconv.Itoa(opt.ReplyToMessageID))
      }
}
}
}
func optionExist(slice []string, val string) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}
