package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Variables of custom error messages
var (
	errInvaildURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

//handleError send in telegram an error message depending on the error itself
func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Default)

	switch err {
	case errInvaildURL:
		msg.Text = "This is an invalid link!"
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "You are not logged in! Use the command /start"
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Oops, the link could not be saved. Please try again later."
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}

}
