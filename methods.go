package telepher

import (
	"encoding/json"
	"fmt"
	"github.com/goTelegramBot/telepher/types"
	"net/url"
	"strconv"
)

func (telepher Bot) bot() *types.User {
	return telepher.Self
}
func (telepher Bot) ID() int {
	return telepher.Self.ID
}
func (telepher Bot) IsBot() bool {
	return telepher.Self.IsBot
}
func (telepher Bot) FirstName() string {
	return telepher.Self.FirstName
}
func (telepher Bot) LastName() string {
	return telepher.Self.FirstName
}
func (telepher Bot) UserName() string {
	return telepher.Self.UserName
}
func (telepher Bot) LanguageCode() string {
	return telepher.Self.LanguageCode
}
func (telepher Bot) CanJoinGroups() bool {
	return telepher.Self.CanJoinGroups
}
func (telepher Bot) CanReadAllGroupMessages() bool {
	return telepher.Self.CanReadAllGroupMessages
}
func (telepher Bot) SupportsInlineQueries() bool {
	return telepher.Self.SupportsInlineQueries
}

func (telepher Bot) GetMe() (*types.User, error) {
	data, err := telepher.post("getMe", nil)
	if err != nil {
		return nil, err
	}
	var user types.User

	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (telepher Bot) GetUpdates(offset int, limit int) ([]types.Update, error) {

	values := url.Values{}
	values.Add("offset", strconv.Itoa(offset))
	values.Add("limit", strconv.Itoa(limit))

	data, err := telepher.post("getUpdates", values)
	if err != nil {
		return nil, err
	}

	var update []types.Update
	err = json.Unmarshal(data, &update)
	if err != nil {
		return nil, err
	}
	return update, nil

}

func (telepher Bot) SendMessage(chat_id int, text string, option interface{}) (*types.Message, error) {

	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("text", text)

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}

	response, err := telepher.post("sendMessage", values)
	if err != nil {
		return nil, err
	}
	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil

}

func (telepher Bot) CopyMessage(chat_id int, from_chat_id int, message_id int, option interface{}) (*types.MessageId, error) {

	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("from_chat_id", strconv.Itoa(from_chat_id))
	values.Add("message_id", strconv.Itoa(message_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}

	response, err := telepher.post("copyMessage", values)
	if err != nil {
		return nil, err
	}
	var messageID *types.MessageId
	err = json.Unmarshal(response, &messageID)

	if err != nil {
		return nil, err
	}
	return messageID, nil

}

func (telepher Bot) ForwardMessage(chat_id int, from_chat_id int, message_id int, option interface{}) (*types.Message, error) {

	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("from_chat_id", strconv.Itoa(from_chat_id))
	values.Add("message_id", strconv.Itoa(message_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}

	response, err := telepher.post("forwardMessage", values)
	if err != nil {
		return nil, err
	}
	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil

}

func (telepher Bot) SendPhoto(chat_id int, photo types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}

	message, err := telepher.sendFile("sendPhoto", "photo", photo, values)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) SendAudio(chat_id int, photo types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendAudio", "audio", photo, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) SendDocument(chat_id int, document types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendDocument", "document", document, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) SendVideo(chat_id int, video types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendVideo", "video", video, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) SendAnimation(chat_id int, animation types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendAnimation", "animation", animation, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) SendVoice(chat_id int, voice types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendVoice", "voice", voice, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) SendVideoNote(chat_id int, video_note types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendVideoNote", "video_note", video_note, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

/*
func (telepher Bot) SendMediaGroup(chat_id int, media []types.InputFile,option interface{})(*types.Message, error){
  values := url.Values{}
  values.Add("chat_id",strconv.Itoa(chat_id))

err := telepher.addOptions(values,option)
if err != nil {
      return nil,err
}
  message, err := telepher.sendFile("sendMediaGroup","media", media, values)

  if err != nil {
      return nil, err
  }
  return message, nil
}


*/

func (telepher Bot) SendLocation(chat_id int, latitude float64, longitude float64, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	values.Add("latitude", fmt.Sprint(latitude))

	values.Add("longitude", fmt.Sprint(longitude))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("sendLocation", values)
	if err != nil {
		return nil, err
	}
	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) EditMessageLiveLocation(latitude float64, longitude float64, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("latitude", fmt.Sprint(latitude))
	values.Add("longitude", fmt.Sprint(longitude))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("editMessageLiveLocation", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) StopMessageLiveLocation(option interface{}) (*types.Message, error) {
	values := url.Values{}

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("stopMessageLiveLocation", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) SendVenue(chat_id int, latitude float64, longitude float64, title string, address string, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("latitude", fmt.Sprint(latitude))
	values.Add("longitude", fmt.Sprint(longitude))
	values.Add("title", title)
	values.Add("address", address)

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("sendVenue", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) SendContact(chat_id int, phone_number string, first_name string, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("phone_number", phone_number)
	values.Add("first_name", first_name)
	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("sendContact", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) SendPoll(chat_id int, question string, poll_options []string, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("question", question)

	options, err := json.Marshal(poll_options)
	if err != nil {
		return nil, err
	}

	values.Add("options", string(options))

	err = telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("sendPoll", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) SendDice(chat_id int, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("sendDice", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) SendChatAction(chat_id int, action string) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("action", action)

	response, err := telepher.post("sendChatAction", values)

	if err != nil {
		return nil, err
	}

	var message *types.Message
	err = json.Unmarshal(response, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) GetUserProfilePhotos(user_id int, option interface{}) (*types.UserProfilePhotos, error) {
	values := url.Values{}
	values.Add("user_id", strconv.Itoa(user_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	response, err := telepher.post("getUserProfilePhotos", values)

	if err != nil {
		return nil, err
	}

	var user_profile_photos *types.UserProfilePhotos
	err = json.Unmarshal(response, &user_profile_photos)

	if err != nil {
		return nil, err
	}

	return user_profile_photos, nil
}

func (telepher Bot) GetFile(file_id string) (*types.File, error) {
	values := url.Values{}
	values.Add("file_id", file_id)

	response, err := telepher.post("getFile", values)

	if err != nil {
		return nil, err
	}

	var file *types.File
	err = json.Unmarshal(response, &file)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (telepher Bot) KickChatMember(chat_id int, user_id int, option interface{}) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("user_id", strconv.Itoa(user_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return false, err
	}
	response, err := telepher.post("kickChatMember", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) UnbanChatMember(chat_id int, user_id int, option interface{}) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("user_id", strconv.Itoa(user_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return false, err
	}
	response, err := telepher.post("unbanChatMember", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) RestrictChatMember(chat_id int, user_id int, permissions *types.ChatPermissions, option interface{}) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("user_id", strconv.Itoa(user_id))

	permission, err := json.Marshal(permissions)
	if err != nil {
		return false, err
	}

	values.Add("permissions", string(permission))

	err = telepher.addOptions(values, option)
	if err != nil {
		return false, err
	}
	response, err := telepher.post("restrictChatMember", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) PromoteChatMember(chat_id int, user_id int, option interface{}) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("user_id", strconv.Itoa(user_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return false, err
	}
	response, err := telepher.post("promoteChatMember", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) SetChatAdministratorCustomTitle(chat_id int, user_id int, custom_title string) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("user_id", strconv.Itoa(user_id))
	values.Add("custom_title", custom_title)

	response, err := telepher.post("setChatAdministratorCustomTitle", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) SetChatPermissions(chat_id int, permissions *types.ChatPermissions) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	permission, err := json.Marshal(permissions)
	if err != nil {
		return false, err
	}

	values.Add("permissions", string(permission))

	response, err := telepher.post("setChatPermissions", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) ExportChatInviteLink(chat_id int) (string, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("exportChatInviteLink", values)

	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (telepher Bot) CreateChatInviteLink(chat_id int, option interface{}) (*types.ChatInviteLink, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}

	response, err := telepher.post("createChatInviteLink", values)

	if err != nil {
		return nil, err
	}

	var result *types.ChatInviteLink
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) EditChatInviteLink(chat_id int, invite_link string, option interface{}) (*types.ChatInviteLink, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("invite_link", invite_link)

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}

	response, err := telepher.post("editChatInviteLink", values)

	if err != nil {
		return nil, err
	}

	var result *types.ChatInviteLink
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) RevokeChatInviteLink(chat_id int, invite_link string) (*types.ChatInviteLink, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("invite_link", invite_link)

	response, err := telepher.post("revokeChatInviteLink", values)

	if err != nil {
		return nil, err
	}

	var result *types.ChatInviteLink
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) SetChatPhoto(chat_id int, photo types.InputFile) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	message, err := telepher.sendFile("setChatPhoto", "photo", photo, values)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (telepher Bot) DeleteChatPhoto(chat_id int) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("deleteChatPhoto", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) SetChatTitle(chat_id int, title string) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("title", title)

	response, err := telepher.post("setChatTitle", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) SetChatDescription(chat_id int, description string) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("description", description)

	response, err := telepher.post("setChatDescription", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) PinChatMessage(chat_id int, message_id int) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("message_id", strconv.Itoa(message_id))

	response, err := telepher.post("pinChatMessage", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) UnpinChatMessage(chat_id int, message_id int) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("message_id", strconv.Itoa(message_id))

	response, err := telepher.post("unpinChatMessage", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) UnpinAllChatMessages(chat_id int) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("unpinAllChatMessages", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) LeaveChat(chat_id int) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("leaveChat", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) GetChat(chat_id int) (*types.Chat, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("getChat", values)

	if err != nil {
		return nil, err
	}
	var result *types.Chat
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) GetChatAdministrators(chat_id int) (*[]types.ChatMember, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("getChatAdministrators", values)

	if err != nil {
		return nil, err
	}
	var result *[]types.ChatMember
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) GetChatMembersCount(chat_id int) (int, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("getChatMembersCount", values)

	if err != nil {
		return 0, err
	}
	var result int
	err = json.Unmarshal(response, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (telepher Bot) GetChatMember(chat_id int, user_id int) (*types.ChatMember, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("user_id", strconv.Itoa(user_id))

	response, err := telepher.post("getChatMember", values)

	if err != nil {
		return nil, err
	}
	var result *types.ChatMember
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) SetChatStickerSet(chat_id int, stickersetname string) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("sticker_set_name", stickersetname)

	response, err := telepher.post("setChatStickerSet", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) DeleteChatStickerSet(chat_id int) (bool, error) {
	values := url.Values{}

	values.Add("chat_id", strconv.Itoa(chat_id))

	response, err := telepher.post("deleteChatStickerSet", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) AnswerCallbackQuery(callback_query_id string) (bool, error) {
	values := url.Values{}

	values.Add("callback_query_id", callback_query_id)
	

	response, err := telepher.post("answerCallbackQuery", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) SetMyCommands(command []types.BotCommand) (bool, error) {
	values := url.Values{}

	commands, err := json.Marshal(command)
	if err != nil {
		return false, err
	}

	values.Add("commands", string(commands))

	response, err := telepher.post("setMyCommands", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) GetMyCommands() ([]*types.BotCommand, error) {
	values := url.Values{}

	response, err := telepher.post("getMyCommands", values)

	if err != nil {
		return nil, err
	}
	var result []*types.BotCommand
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) EditMessageText(text string) (*types.Message, error) {
	values := url.Values{}
	values.Add("text", text)

	response, err := telepher.post("editMessageText", values)

	if err != nil {
		return nil, err
	}
	var result *types.Message
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) EditMessageCaption() (*types.Message, error) {
	values := url.Values{}

	response, err := telepher.post("editMessageCaption", values)

	if err != nil {
		return nil, err
	}
	var result *types.Message
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
func (telepher Bot) EditMessageMedia(media types.InputFile)(*types.Message, error){
  values := url.Values{}
   values.Add("chat_id","1520625615")
    values.Add("message_id","7000")


 message, err := telepher.sendFile("editMessageMedia","media", media, values)
  if err != nil {
      return nil,err
  }
  return message, nil
}
*/

func (telepher Bot) EditMessageReplyMarkup() (*types.Message, error) {
	values := url.Values{}

	response, err := telepher.post("editMessageReplyMarkup", values)

	if err != nil {
		return nil, err
	}
	var result *types.Message
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) StopPoll(chat_id int, message_id int) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("message_id", strconv.Itoa(message_id))
	response, err := telepher.post("stopPoll", values)

	if err != nil {
		return nil, err
	}
	var result *types.Message
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) DeleteMessage(chat_id int, message_id int) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))
	values.Add("message_id", strconv.Itoa(message_id))

	response, err := telepher.post("deleteMessage", values)

	if err != nil {
		return nil, err
	}
	var result *types.Message
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) SendSticker(chat_id int, video_note types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chat_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("sendSticker", "sticker", video_note, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) GetStickerSet(name string) (*types.StickerSet, error) {
	values := url.Values{}
	values.Add("name", name)

	response, err := telepher.post("getStickerSet", values)

	if err != nil {
		return nil, err
	}
	var result *types.StickerSet
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (telepher Bot) UploadStickerFile(user_id int, png_sticker types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("user_id", strconv.Itoa(user_id))

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("uploadStickerFile", "png_sticker", png_sticker, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) CreateNewStickerSet(user_id int, name string, title string, emojis string, png_sticker types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("user_id", strconv.Itoa(user_id))
	values.Add("name", name)
	values.Add("title", title)
	values.Add("emojis", emojis)

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("createNewStickerSet", "png_sticker", png_sticker, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) AddStickerToSet(user_id int, name string, emojis string, png_sticker types.InputFile, option interface{}) (*types.Message, error) {
	values := url.Values{}
	values.Add("user_id", strconv.Itoa(user_id))
	values.Add("name", name)
	values.Add("emojis", emojis)

	err := telepher.addOptions(values, option)
	if err != nil {
		return nil, err
	}
	message, err := telepher.sendFile("addStickerToSet", "png_sticker", png_sticker, values)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (telepher Bot) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	values := url.Values{}

	values.Add("sticker", sticker)
	values.Add("position", strconv.Itoa(position))

	response, err := telepher.post("setStickerPositionInSet", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (telepher Bot) DeleteStickerFromSet(sticker string) (bool, error) {
	values := url.Values{}

	values.Add("sticker", sticker)

	response, err := telepher.post("deleteStickerFromSet", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}
func (telepher Bot) SetStickerSetThumb(name string, user_id int) (bool, error) {
	values := url.Values{}

	values.Add("name", name)
	values.Add("user_id", strconv.Itoa(user_id))

	response, err := telepher.post("deleteStickerFromSet", values)

	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(response, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

