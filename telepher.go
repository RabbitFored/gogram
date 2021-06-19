package telepher

import (
  "net/http"
  "time"
 
  )
const (
  DefaultBaseURL = "https://api.telegram.org/bot"
  DefaultBaseFileURL = "https://api.telegram.org/file/bot"
  DefaultParseMode = ""
  DefaultCheckBool = true
 )

const DefaultHandlerGroup = 0

var DefaultClient = &http.Client{
  Timeout: time.Second * 10,
}

type ParseMode = string

var PARSE_MODES = []string{"markdownv2", "html", "markdown",""}

const (
	Default    ParseMode = ""
  Markdown   ParseMode = "Markdown"
	MarkdownV2 ParseMode = "MarkdownV2"
	HTML       ParseMode = "HTML"
)


// events
const (
    Message                    =      "message"
    EditedMessage              =      "edited_message"
    ChannelPost                =      "channel_post"
    EditedChannelPost          =      "edited_channel_post"
    InlineQuery                =      "inline_query"
    ChosenInlineResult         =      "chosen_inline_result"
    CallbackQuery              =      "callback_query"
    ShippingQuery              =      "shipping_query"
    PreCheckoutQuery           =      "pre_checkout_query"
    Poll                       =      "poll"
    PollAnswer                 =      "inline_query"
    MyChatMember               =      "my_chat_member"
    ChatMember                 =      "chat_member"
    Text                        =      "text"

)
