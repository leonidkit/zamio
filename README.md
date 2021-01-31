## Тестовое задание

* Создать публичный репозиторий, например Github/Gitlab.
* Создать таблицу Postgresql для хранения информации об аккаунте и балансе, например с полями account, balance.
* Написать GO сервис реализующий метод перевода средств с аккаунта А на аккаунт Б. Важно придерживаться Чистой архитектуры. 
* Метод сервиса должен быть доступен через API или CLI.
* Написать тесты.

В качестве примера можно взять https://github.com/eminetto/clean-architecture-go-v2
В качестве примера был взят https://github.com/zhashkevych/courses-backend

### Дополнительно на ваше усмотрение:
* Написать Dockerfile.
* Реализовать метод получения всех балансов.

## Usage
```
make build
./zamio -h
./zamio -email-first another@mail.ru -email-second some@mail.ru -sum 100
```