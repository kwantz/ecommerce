package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/kwantz/ecommerce/services/product/internal/controller"
	"github.com/kwantz/ecommerce/services/product/internal/repository"
	"github.com/kwantz/ecommerce/services/product/internal/usecase"
)

func main() {
	db := setupDatabase()
	defer db.Close()

	accountHostname := "http://account-service:8080"

	cartRepository := repository.NewCartRepository(db)
	productRepository := repository.NewProductRepository(db)
	accountRepository := repository.NewAccountRepository(accountHostname)

	accountUsecase := usecase.NewAccountUsecase(accountRepository)
	productUsecase := usecase.NewProductUsecase(productRepository)
	cartUsecase := usecase.NewCartUsecase(cartRepository, productRepository)

	pingController := controller.NewPingController()
	cartController := controller.NewCartController(cartUsecase, accountUsecase)
	productController := controller.NewProductController(productUsecase)

	router := chi.NewRouter()

	setupRouting(router, Controller{
		Ping:    pingController,
		Cart:    cartController,
		Product: productController,
	})

	log.Println("server starting at :8080")
	http.ListenAndServe(":8080", router)
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user_product:password_product@tcp(mysql-product:3306)/ecommerce_product")
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

	router.Post("/cart", controller.Cart.AddProductToCartHandler)
	router.Delete("/cart/{cartID}", controller.Cart.RemoveProductFromCartHandler)

	router.Post("/cart/order", controller.Cart.GetCartProductsOrderHandler)
	router.Delete("/cart/order", controller.Cart.DeleteCartProductsOrderHandler)

	router.Post("/", controller.Product.CreateProductHandler)
	router.Get("/", controller.Product.GetAllProductHandler)
	router.Get("/{productID}", controller.Product.GetProductHandler)
}

type Controller struct {
	Ping    *controller.PingController
	Cart    *controller.CartController
	Product *controller.ProductController
}
