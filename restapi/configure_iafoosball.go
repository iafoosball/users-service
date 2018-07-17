// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/tylerb/graceful"

	"github.com/iafoosball/users-service/restapi/operations"
	"github.com/iafoosball/users-service/iaf/users"
	"github.com/iafoosball/users-service/iaf/friends"

)

//go:generate swagger generate server --target .. --name iafoosball-test --spec ../swagger.yml

func configureFlags(api *operations.IafoosballAPI) {

}

func configureAPI(api *operations.IafoosballAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//api.XMLProducer = runtime.XMLProducer()

	//api.TxtProducer = runtime.TextProducer()

	api.PostUsersHandler = operations.PostUsersHandlerFunc(users.CreateUser())
	api.GetUsersUserIDHandler = operations.GetUsersUserIDHandlerFunc(users.GetUserByID())

	api.GetFriendsUserIDHandler = operations.GetFriendsUserIDHandlerFunc(friends.GetFriends())
	api.PostFriendsUserIDFriendIDHandler = operations.PostFriendsUserIDFriendIDHandlerFunc(friends.MakeFriendRequest())
	api.PatchFriendsUserIDFriendIDHandler = operations.PatchFriendsUserIDFriendIDHandlerFunc(friends.AcceptFriendRequest())
	api.DeleteFriendsFriendshipIDHandler = operations.DeleteFriendsFriendshipIDHandlerFunc(friends.DeleteFriend())

	//){
	//api.GetUsersUserIDHandler = operations.GetUsersUserIDHandlerFunc(func(params operations.GetUsersUserIDParams) middleware.Responder {
	//	return middleware.NotImplemented("operation .GetUsersUserID has not yet been implemented")
	//})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
	addr = "0.0.0.0:4444"
}


// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
