package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/kwantz/ecommerce/services/order/internal/controller"
	"github.com/kwantz/ecommerce/services/order/internal/repository"
	"github.com/kwantz/ecommerce/services/order/internal/usecase"
)

func main() {
	db := setupDatabase()
	defer db.Close()

	cartHostname := "http://product-service:8080"
	accountHostname := "http://account-service:8080"

	orderRepository := repository.NewOrderRepository(db)
	cartRepository := repository.NewCartRepository(cartHostname)
	accountRepository := repository.NewAccountRepository(accountHostname)
	orderProductRepository := repository.NewOrderProductRepository(db)

	accountUsecase := usecase.NewAccountUsecase(accountRepository)
	orderUsecase := usecase.NewOrderUsecase(usecase.OrderUsecaseOption{
		CartRepository:         cartRepository,
		OrderRepository:        orderRepository,
		OrderProductRepository: orderProductRepository,
	})

	pingController := controller.NewPingController()
	orderController := controller.NewOrderController(orderUsecase, accountUsecase)

	router := chi.NewRouter()

	setupRouting(router, Controller{
		Ping:  pingController,
		Order: orderController,
	})

	log.Println("server starting at :8080")
	http.ListenAndServe(":8080", router)
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user_order:password_order@tcp(mysql-order:3306)/ecommerce_order")
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func setupRouting(router *chi.Mux, controller Controller) {
	router.Get("/ping", controller.Ping.PingHandler)

	router.Post("/", controller.Order.CreateOrderHandler)
}

type Controller struct {
	Ping  *controller.PingController
	Order *controller.OrderController
}
