package telepher

import (
	_ "fmt"
	"log"
	"main/types"
	"sort"
	"strings"
)

func (telepher Bot) Start() {

	go telepher.Poll(telepher.Update)

	for update := range telepher.Update {

		if update.Message != nil {
			telepher.handle("message", update.Message)
			telepher.handleMessage(update.Message)

			if update.Message.Text != "" {
				telepher.handle("text", update.Message)

				if len(update.Message.Entities) != 0 && update.Message.Entities[0].Type == "bot_command" && update.Message.Entities[0].Offset == 0 {

					command := update.Message.Text[1:update.Message.Entities[0].Length]
					command_parts := strings.Split(command, "@")
					command_parts = append(command_parts, telepher.Self.UserName)

					if command_parts[1] == telepher.Self.UserName {
						cmd := strings.ToLower(command_parts[0])
						telepher.CommandHandler(cmd, update.Message)
					}

				}
			}

			for _, entity := range update.Message.Entities {
				if entity.Type == "url" {
					telepher.handle("url", update.Message)
				}
				if entity.Type == "mention" {
					telepher.handle("mention", update.Message)
				}
				if entity.Type == "hashtag" {
					telepher.handle("hashtag", update.Message)
				}
				if entity.Type == "cashtag" {
					telepher.handle("cashtag", update.Message)
				}
				if entity.Type == "bot_command" {
					telepher.handle("bot_command", update.Message)

				}
				if entity.Type == "email" {
					telepher.handle("email", update.Message)
				}
				if entity.Type == "phone_number" {
					telepher.handle("phone_number", update.Message)
				}
				if entity.Type == "bold" {
					telepher.handle("bold", update.Message)
				}
				if entity.Type == "italic" {
					telepher.handle("italic", update.Message)
				}
				if entity.Type == "underline" {
					telepher.handle("underline", update.Message)
				}
				if entity.Type == "strikethrough" {
					telepher.handle("strikethrough", update.Message)
				}
				if entity.Type == "code" {
					telepher.handle("code", update.Message)
				}
				if entity.Type == "pre" {
					telepher.handle("pre", update.Message)
				}
				if entity.Type == "text_link" {
					telepher.handle("text_link", update.Message)
				}
				if entity.Type == "text_mention" {
					telepher.handle("text_mention", update.Message)
				}
			}

			if update.Message.Animation != nil {
				telepher.handle("animation", update.Message)
			}
			if update.Message.Audio != nil {
				telepher.handle("audio", update.Message)
			}

			if update.Message.Document != nil {
				telepher.handle("document", update.Message)
			}
			if update.Message.Photo != nil {
				telepher.handle("photo", update.Message)
			}
			if update.Message.Sticker != nil {
				telepher.handle("sticker", update.Message)
			}
			if update.Message.Video != nil {
				telepher.handle("video", update.Message)
			}
			if update.Message.VideoNote != nil {
				telepher.handle("VideoNote", update.Message)
			}
			if update.Message.Voice != nil {
				telepher.handle("Voice", update.Message)
			}
			for _, entity := range update.Message.CaptionEntities {
				if entity.Type == "url" {
					telepher.handle("caption_url", update.Message)
				}
				if entity.Type == "mention" {
					telepher.handle("caption_mention", update.Message)
				}
				if entity.Type == "hashtag" {
					telepher.handle("caption_hashtag", update.Message)
				}
				if entity.Type == "cashtag" {
					telepher.handle("caption_cashtag", update.Message)
				}
				if entity.Type == "bot_command" {
					telepher.handle("caption_bot_command", update.Message)

				}
				if entity.Type == "email" {
					telepher.handle("caption_email", update.Message)
				}
				if entity.Type == "phone_number" {
					telepher.handle("caption_phone_number", update.Message)
				}
				if entity.Type == "bold" {
					telepher.handle("caption_bold", update.Message)
				}
				if entity.Type == "italic" {
					telepher.handle("caption_italic", update.Message)
				}
				if entity.Type == "underline" {
					telepher.handle("caption_underline", update.Message)
				}
				if entity.Type == "strikethrough" {
					telepher.handle("caption_strikethrough", update.Message)
				}
				if entity.Type == "code" {
					telepher.handle("caption_code", update.Message)
				}
				if entity.Type == "pre" {
					telepher.handle("caption_pre", update.Message)
				}
				if entity.Type == "text_link" {
					telepher.handle("caption_text_link", update.Message)
				}
				if entity.Type == "text_mention" {
					telepher.handle("caption_text_mention", update.Message)
				}
			}
			if update.Message.Contact != nil {
				telepher.handle("contact", update.Message)
			}
			if update.Message.Dice != nil {
				telepher.handle("Dice", update.Message)
			}
			if update.Message.Game != nil {
				telepher.handle("Game", update.Message)
			}
			if update.Message.Poll != nil {
				telepher.handle("Poll", update.Message)
			}
			if update.Message.Venue != nil {
				telepher.handle("Venue", update.Message)
			}
			if update.Message.Location != nil {
				telepher.handle("Location", update.Message)
			}
			if update.Message.NewChatMembers != nil {
				telepher.handle("NewChatMembers", update.Message)
			}
			if update.Message.LeftChatMember != nil {
				telepher.handle("LeftChatMember", update.Message)
			}
			if update.Message.NewChatTitle != "" {
				telepher.handle("NewChatTitle", update.Message)
			}
			if update.Message.NewChatPhoto != nil {
				telepher.handle("NewChatPhoto", update.Message)
			}
			if update.Message.MessageAutoDeleteTimerChanged != nil {
				telepher.handle("MessageAutoDeleteTimerChanged", update.Message)
			}

			if update.Message.PinnedMessage != nil {
				telepher.handle("PinnedMessage", update.Message)
			}
			if update.Message.Invoice != nil {
				telepher.handle("Invoice", update.Message)
			}
			if update.Message.SuccessfulPayment != nil {
				telepher.handle("SuccessfulPayment", update.Message)
			}

			/////////
		}
		if update.CallbackQuery != nil {

			var groupIDs []int

			for i := range telepher.Handlers {
				groupIDs = append(groupIDs, i)
			}
			sort.Ints(groupIDs)

			for _, groupID := range groupIDs {
				for _, handler := range telepher.Handlers[groupID] {

					if handler.Type == "CallbackQuery" {
						handler.function.(func(Bot, *types.CallbackQuery))(telepher, update.CallbackQuery)

					}
				}
			}
		}

	}
}

func (telepher Bot) Poll(ch chan types.Update) {
	LastUpdateID := 10
	for {
		resp, err := telepher.GetUpdates(LastUpdateID, 1)
		if err != nil {
			log.Fatal(err)

		}

		var key = len(resp) - 1
		if key != -1 {
			LastUpdateID = (resp[key].UpdateID) + 1

			ch <- resp[key]
		}
	}
}
