package telepher

import (
	"encoding/json"
	"fmt"
	"github.com/goTelegramBot/telepher/types"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Options struct {
	ParseMode                string
	Entities                 []types.MessageEntity
	DisableWebPagePreview    bool
	DisableNotification      bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
	ReplyMarkup              *types.ReplyMarkup
	Caption                  string
	CaptionEntities          []types.MessageEntity
	Params                   url.Values
}

func (telepher Bot) addOptions(values url.Values, option interface{}) error {

	switch option.(type) {
	case url.Values:
		for key := range option.(url.Values) {
			values.Add(key, option.(url.Values).Get(key))
		}
		return nil
	case *Options:
		err := telepher.extractOptions(values, option.(*Options))
		if err != nil {
			return err
		}
		return nil
	case map[string]string:
		for key, value := range option.(map[string]string) {
			values.Add(key, value)
		}
		return nil
	case nil:
		return nil
	default:
		return fmt.Errorf("unsupported options type %s", reflect.TypeOf(option))
	}
}

func (telepher Bot) extractOptions(values url.Values, option *Options) error {
	if option != nil {
		if option.ParseMode == "" {
			option.ParseMode = telepher.ParseMode
		}
		validParseMode := validParseMode(PARSE_MODES, strings.ToLower(option.ParseMode))

		if validParseMode {
			values.Add("parse_mode", option.ParseMode)
		} else {
			return fmt.Errorf("unsupported parse_mode %s", option.ParseMode)
		}

		values.Add("disable_web_page_preview", strconv.FormatBool(option.DisableWebPagePreview))
		values.Add("disable_notification", strconv.FormatBool(option.DisableNotification))

		if option.ReplyToMessageID != 0 {
			values.Add("reply_to_message_id", strconv.Itoa(option.ReplyToMessageID))
		}

		if option.ReplyMarkup != nil {
			err := addKeyboard(values, option.ReplyMarkup)
			if err != nil {
				return err
			}
		}

		values.Add("caption", option.Caption)

		/*	if option.Thumb != nil {
			switch option.Thumb.(type) {
			case string:
				values.Add("thumb", m)

			case NamedReader:
				values.Add("thumb", "attach://thumb")
				data["thumb"] = m

			case io.Reader:
				values.Add("thumb", "attach://thumb")
				data["thumb"] = NamedFile{File: m}

			case []byte:
				values.Add("thumb", "attach://thumb")
				data["thumb"] = NamedFile{File: bytes.NewReader(m)}

			default:
				return nil, fmt.Errorf("unknown type for InputFile: %T", opts.Thumb)
			}
		}*/

		if option.Entities != nil {
			entities, err := json.Marshal(option.CaptionEntities)
			if err != nil {
				return fmt.Errorf("failed to marshal field caption_entities: %w", err)
			}
			values.Add("entities", string(entities))
		}
		if option.CaptionEntities != nil {
			caption_entities, err := json.Marshal(option.CaptionEntities)
			if err != nil {
				return fmt.Errorf("failed to marshal field caption_entities: %w", err)
			}
			values.Add("caption_entities", string(caption_entities))
		}
		if option.Params != nil {
			for key := range option.Params {
				values.Add(key, option.Params.Get(key))
			}
		}
	}

	return nil

}

func validParseMode(array []string, item string) bool {

	for i := range array {
		if array[i] == item {
			return true
		}
	}
	return false
}

func addKeyboard(values url.Values, replyMarkup *types.ReplyMarkup) error {

	if replyMarkup.InlineKeyboardMarkup != nil {
		inline_keyboard_markup, err := json.Marshal(replyMarkup.InlineKeyboardMarkup)

		if err != nil {
			return err
		}
		values.Add("reply_markup", string(inline_keyboard_markup))
	}
	if replyMarkup.ReplyKeyboardMarkup != nil {
		reply_keyboard_markup, err := json.Marshal(replyMarkup.ReplyKeyboardMarkup)

		if err != nil {
			return err
		}
		values.Add("reply_markup", string(reply_keyboard_markup))
	}

	return nil
}
