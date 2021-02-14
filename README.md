# Pocket Бот

## <a href="https://www.youtube.com/playlist?list=PLbTTxxr-hMmzSGTsO5mdYLrvKY-RZFanp">Курс на YouTube</a>

Pocket - это Telegram бот, который позволяет сохранять ссылки в приложении <a href="https://app.getpocket.com/">Pocket</a>. Можно сказать что он является небольшим клиентом для данного сервиса. 

Для работы с Pocket API используется самописное SDK - <a href="https://github.com/zhashkevych/go-pocket-sdk">go-pocket-sdk</a>.

В качестве хранилища используется <a href="https://github.com/boltdb/bolt">Bolt DB</a>.

Чтобы реализовать авторизацию пользователей, вместе с ботом запускается HTTP сервер на порту 80, на который происходит редирект от Pocket при успешной авторизации пользователя. 

Когда сервер принимает запрос, он генерирует Access Token через Pocket API для пользователя и сохраняет его в хранилище.

### Стек:
- Go 1.15
- BoltDB
- Docker (для развертывания)
