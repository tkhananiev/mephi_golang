package main

import (
	"net/http"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	. "rest_module/repository"
	. "rest_module/rest"
	. "rest_module/service"
)

func main() {
	var dbManager = NewDBManager()
	defer dbManager.CloseConnection()
	err := dbManager.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	var mailSender, _ = InitMailSender()

	var userRepository = InitUserRepository(dbManager)
	var userManager = UserManagerNewInstance(userRepository)
	var usersController = UsersControllerNewInstance(userManager)

	var accountRepository = InitAccountRepository(dbManager)
	var accountManager = AccountManagerNewInstance(userRepository, accountRepository)
	var accountController = AccountControllerNewInstance(accountManager)

	var cardRepository = InitCardRepository(dbManager)
	var cardManager = CardManagerNewInstance(mailSender, userRepository, cardRepository)
	var cardController = CardControllerNewInstance(cardManager)

	var operRepository = InitOperationRepository(dbManager)
	var operManager = OperationManagerNewInstance(mailSender, userRepository, accountRepository, operRepository)
	var operController = OperationControllerNewInstance(operManager)

	var paymentRepository = InitPaymentRepository(dbManager)
	var creditRepository = InitCreditRepository(dbManager)
	var creditManager = CreditManagerNewInstance(mailSender, userRepository, accountRepository, creditRepository, paymentRepository)
	var creditController = CreditControllerNewInstance(creditManager)

	// Планировщик
	c := cron.New()
	c.AddFunc("* 0 * * *", func() {
		log.Println("Старт задания списания платежей")
		err = creditManager.PaymentForCredit()
		if err != nil {
			log.Panicln(err.Error())
		}
	})
	c.Start()

	// Главный контроллер
	api := ApiNewInstance(usersController, accountController, cardController, operController, creditController)

	// ⬇️ Вызов Router(api)
	err = http.ListenAndServe(":8080", api.Router())
	if err != nil {
		log.Fatal(err)
	}

	c.Stop()
}
