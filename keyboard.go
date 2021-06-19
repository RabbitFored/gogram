package telepher

import (
"github.com/goTelegramBot/telepher/types"
)
func InlineKeyboardMarkup() (*InlineKeyboard){
InlineKeyboard := &InlineKeyboard{}
return InlineKeyboard
}
func ReplyKeyboardMarkup(keyboardOptions *KeyboardOptions) (*ReplyKeyboard){
ReplyKeyboard := &ReplyKeyboard{}

if keyboardOptions != nil {
ReplyKeyboard.ResizeKeyboard = keyboardOptions.ResizeKeyboard
ReplyKeyboard.OneTimeKeyboard = keyboardOptions.OneTimeKeyboard
ReplyKeyboard.Selective = keyboardOptions.Selective
}
return ReplyKeyboard
}

type InlineKeyboard struct{
}

type ReplyKeyboard struct{
    ResizeKeyboard bool 
	  OneTimeKeyboard bool 
	  Selective bool 
}

type KeyboardOptions struct {
  	ResizeKeyboard bool 
	  OneTimeKeyboard bool 
	  Selective bool 
}



func (inline InlineKeyboard) Row(InlineButton ...types.InlineKeyboardButton)([]types.InlineKeyboardButton){
  
  var row []types.InlineKeyboardButton

  for _, data := range InlineButton {
     row = append(row,types.InlineKeyboardButton{

       Text: data.Text,
       Url: data.Url,
       LoginUrl: data.LoginUrl,
       CallbackData: data.CallbackData,
       SwitchInlineQuery: data.SwitchInlineQuery,
       SwitchInlineQueryCurrentChat: data.SwitchInlineQueryCurrentChat,
       CallbackGame: data.CallbackGame,
       Pay: data.Pay,

       })
  }

return row

}

func (keyboard ReplyKeyboard) Row(keyboardButton ...types.KeyboardButton)([]types.KeyboardButton){
    var row []types.KeyboardButton

  for _, data := range keyboardButton {
     row = append(row,types.KeyboardButton{

	Text: data.Text,
	RequestContact: data.RequestContact,
	RequestLocation: data.RequestLocation,
	RequestPoll: data.RequestPoll,
       })
  }

return row

}


func (inlineKeyboard InlineKeyboard) Parse(buttons ...[]types.InlineKeyboardButton)(types.ReplyMarkup){

reply_markup  := types.ReplyMarkup{InlineKeyboardMarkup: &types.InlineKeyboardReplyMarkup{InlineKeyboardMarkup:buttons}}
return reply_markup
}

func (replyKeyboard ReplyKeyboard) Parse(buttons ...[]types.KeyboardButton)(types.ReplyMarkup){


reply_markup  := types.ReplyMarkup{ReplyKeyboardMarkup: &types.KeyboardMarkup{Keyboard:buttons,ResizeKeyboard:replyKeyboard.ResizeKeyboard,OneTimeKeyboard:replyKeyboard.OneTimeKeyboard,Selective:replyKeyboard.Selective}}

return reply_markup
}
