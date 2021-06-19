package filters

import (
	"telepher/types"
)

func All(message *types.Message) bool {
	return true
}
func Photo(message *types.Message) bool {
	return message.Photo != nil
}
func Forwarded(message *types.Message) bool {
	return message.ForwardFrom != nil
}

func Create(f func(*types.Message) bool) func(*types.Message) bool {
	return f

}
