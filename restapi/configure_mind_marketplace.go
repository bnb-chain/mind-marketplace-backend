// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"github.com/bnb-chain/mind-marketplace-backend/restapi/handlers"
	"github.com/bnb-chain/mind-marketplace-backend/service"
	"github.com/bnb-chain/mind-marketplace-backend/service/twitter"
	"github.com/bnb-chain/mind-marketplace-backend/util"
	"github.com/go-openapi/swag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"github.com/bnb-chain/mind-marketplace-backend/restapi/operations"
	"github.com/bnb-chain/mind-marketplace-backend/restapi/operations/account"
	"github.com/bnb-chain/mind-marketplace-backend/restapi/operations/item"
	"github.com/bnb-chain/mind-marketplace-backend/restapi/operations/purchase"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ../../mind-marketplace-backend --name GreenfieldDataMarketplace --spec ../swagger.yaml --principal interface{}

var cliOpts = struct {
	ConfigFilePath string `short:"c" long:"config-path" description:"Config path" default:"config/server/dev.json"`
}{}

func configureFlags(api *operations.MindMarketplaceAPI) {
	param := swag.CommandLineOptionsGroup{
		ShortDescription: "config",
		Options:          &cliOpts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{param}
}

func configureAPI(api *operations.MindMarketplaceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		util.Logger.Fatal(http.ListenAndServe(":9292", nil))
	}()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.AccountGetAccountHandler = account.GetAccountHandlerFunc(handlers.HandleGetAccount())
	api.AccountUpdateAccountHandler = account.UpdateAccountHandlerFunc(handlers.HandleUpdateAccount())

	api.ItemGetCategoryHandler = item.GetCategoryHandlerFunc(handlers.HandleGetAllCategory())
	api.ItemGetItemHandler = item.GetItemHandlerFunc(handlers.HandleGetItem())
	api.ItemGetItemByGroupHandler = item.GetItemByGroupHandlerFunc(handlers.HandleGetItemByGroup())
	api.ItemSearchItemHandler = item.SearchItemHandlerFunc(handlers.HandleSearchItem())

	api.PurchaseGetPurchaseHandler = purchase.GetPurchaseHandlerFunc(handlers.HandleGetPurchase())
	api.PurchaseSearchPurchaseHandler = purchase.SearchPurchaseHandlerFunc(handlers.HandleSearchPurchase())

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	configFilePath := cliOpts.ConfigFilePath
	config := util.ParseServerConfigFromFile(configFilePath)

	util.InitLogger(config.LogConfig)

	db, err := database.ConnectDBWithConfig(config.DBConfig)
	if err != nil {
		panic(err)
	}

	accountDao := dao.NewDbAccountDao(db)
	itemDao := dao.NewDbItemDao(db)
	categoryDao := dao.NewDbCategoryDao(db)
	purchaseDao := dao.NewDbPurchaseDao(db)

	service.AccountSvc = service.NewAccountService(accountDao)
	service.ItemSvc = service.NewItemService(itemDao)
	service.CategorySvc = service.NewCategoryService(categoryDao)
	service.PurchaseSvc = service.NewPurchaseService(purchaseDao)

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	config := util.ParseServerConfigFromFile(cliOpts.ConfigFilePath)
	return mountTwitterOauth(handlers.SetupHandler(handler, "mind-marketplace", config))
}

func mountTwitterOauth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/twitter_verify" {
			twitter.RedirectToTwitter(w, r)
			return
		} else if r.URL.Path == "/v1/twitter_token" {
			twitter.GetTwitterToken(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}
}
