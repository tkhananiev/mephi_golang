# Go Service Finance
## REST-сервис: банковское приложение на Go, позволяющее управлять пользователями, счетами, картами, кредитами и операциями. 
### Предварительные требования

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [PostgreSQL](https://www.postgresql.org/download/)

### 🚀 Быстрый старт
Для запуска приложения достаточно выполнить команду ``docker-compose up --build -d``

## 🌟 Логика работы

Приложение выполняет следующие основные функции:
1) При старте подключается к БД и выполняет скрипт миграции, который адаптирован для многократного выполнения
2) Запускается планировщик и выполняет задачу `CreditManager.PaymentForCredit()` по крону каждый час, для выполнения списаний по кредитам
3) Для взаимодействия с приложением сначала нужно зарегистрировать пользователя и получить JWT-токен, который нужно положить в хидер каждого запроса
4) для многоих операций выполняется валидация данных
5) Если JWT-токен протух нужно авторизоваться заново и получить новый
6) Зарегистрированному пользователю доступны операции со счетами, картами, кредитами и платежами

### 📌 Список endpoint
#### Авторизация
 * `/health` - проверка, JWT-токен не нужен
 * `/register` - регистрация, JWT-токен не нужен
 * `/login` - аутентификация, JWT-токен не нужен
#### Счета
 * `/analytics` - получить аналитику
 * `/accounts/add` - создать счет
 * `/accounts/{id}/get` - получить счет но id
 * `/accounts/all` - получить список счетов
 * `/accounts/{id}/predict` - прогноз баланса
#### Карты
 * `/cards/add` - выпустить карту (📧 будет отправлено письмо)
 * `/cards/{id}/get` - получить карту
 * `/cards/all` - получить список карт
#### Операции
 * `/operation/debet` - выполнить операцию дебета (📧 будет отправлено письмо)
 * `/operation/credit` - выполнить операцию кредита (📧 будет отправлено письмо)
 * `/operation/transfer` - выполнить перевод (📧 будет отправлено письмо)
 * `/operation/{id}/all` - список всех операций пользователя по счету
 * `/operation/all` - список всех операций пользователя
#### Кредиты
 * `/credits/add` - выдать кредит
 * `/credits/{id}/get` - получить информацию о кредите
 * `/credits/all` - получить список кредитов пользователя
 * `/credits/{id}/schedule` - график платежей по кредиту

 #### Проверка работы приложения:
 Проверить работоспобность можно перейдя по адресу http://localhost:8080/health обязан вернуть {"status":"UP"}

