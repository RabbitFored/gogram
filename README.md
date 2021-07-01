# telepher

[![Support Chat](https://img.shields.io/badge/English%20chat-grey?style=flat-square&logo=telegram)](https://t.me/ostrichdiscussion)
[![Off-topic](https://img.shields.io/badge/English%20chat-grey?style=flat-square&logo=telegram)](https://t.me/unlaidchat)

### Features

- Full [Telegram Bot API 5.2](https://core.telegram.org/bots/api) support ( NOT UPDATED TO BOT API V3 
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
