# telepher

[![Support Chat](https://img.shields.io/badge/Support%20chat-grey?style=flat-square&logo=telegram)](https://t.me/ostrichdiscussion)
[![Off-topic](https://img.shields.io/badge/English%20chat-grey?style=flat-square&logo=telegram)](https://t.me/unlaidchat)

### Features

- Full [Telegram Bot API 5.2](https://core.telegram.org/bots/api) support ( NOT UPDATED TO BOT API 5.3 )
- Lightweight
- Easy to write
- [Support](https://t.me/ostrichdiscussion)


### Usage
    package main

    import (
        tl "github.com/goTelegramBot/telepher"
           "github.com/goTelegramBot/telepher/types"
     )
  

    func main() {
    
           b,err := tl.NewBot(os.Getenv("TOKEN"),nil)
    
           if err != nil{
                log.Println(err)
                return
           }

           b.Command("start",start)
           b.Start()
    }
    func start(bot tl.Bot,message *types.Message) {
            text := "Hi there!"
    
            markup := tl.InlineKeyboardMarkup()
            but1 := types.InlineKeyboardButton{Text:"Channel",Url: "https://t.me/theostrich"}
            but2 := types.InlineKeyboardButton{Text:"Support",Url: "https://t.me/ostrichdiscussion"}
            
            row1 := markup.Row(but1,but2)
            keyboard := markup.Parse(row1)
     
            bot.SendMessage(message.Chat.Id, text,&tl.Options{ReplyMarkup:&keyboard,ParseMode:"Markdown"})
     }
