package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

const (
	commandStart = "start"
	commandList  = "list"

	startMessage = "Привет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandList:
		return b.handleListCommand()
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты уже авторизовался с помощью своего Pocket аккаунта. Присылай ссылку чтобы сохранить.")
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleListCommand() error {
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	if err := b.saveLink(message, accessToken); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Ссылка успешно сохранена!")
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) saveLink(message *tgbotapi.Message, accessToken string) error {
	if err := b.validateURL(message.Text); err != nil {
		return invalidUrlError
	}

	if err := b.client.Add(context.Background(), pocket.AddInput{
		URL:         message.Text,
		AccessToken: accessToken,
	}); err != nil {
		return unableToSaveError
	}

	return nil
}

func (b *Bot) validateURL(text string) error {
	_, err := url.ParseRequestURI(text)
	return err
}
