// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/iafoosball/users-service/users"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/restapi/operations"
)

//go:generate swagger generate server --target .. --name users --spec ../users.yml

func configureFlags(api *operations.UsersAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UsersAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	//[Start: Users end points]
	api.PostUsersHandler = operations.PostUsersHandlerFunc(users.CreateUser())
	api.GetUsersUserIDHandler = operations.GetUsersUserIDHandlerFunc(users.GetUserByID())
	//[End: Users end points]

	//[Start: Friends end points}
	//api.GetFriendsUserIDHandler = operations.GetFriendsUserIDHandlerFunc(usersApi.GetFriends())
	//api.PostFriendsUserIDFriendIDHandler = operations.PostFriendsUserIDFriendIDHandlerFunc(usersApi.MakeFriendRequest())
	//api.PatchFriendsUserIDFriendIDHandler = operations.PatchFriendsUserIDFriendIDHandlerFunc(usersApi.AcceptFriendRequest())
	//api.DeleteFriendsFriendshipIDHandler = operations.DeleteFriendsFriendshipIDHandlerFunc(usersApi.DeleteFriend())
	//[End: Friends end points}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DeleteFriendsFriendshipIDHandler = operations.DeleteFriendsFriendshipIDHandlerFunc(func(params operations.DeleteFriendsFriendshipIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteFriendsFriendshipID has not yet been implemented")
	})
	api.DeleteUsersUserIDHandler = operations.DeleteUsersUserIDHandlerFunc(func(params operations.DeleteUsersUserIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteUsersUserID has not yet been implemented")
	})
	api.GetFriendsUserIDHandler = operations.GetFriendsUserIDHandlerFunc(func(params operations.GetFriendsUserIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetFriendsUserID has not yet been implemented")
	})
	api.GetUsersUserIDHandler = operations.GetUsersUserIDHandlerFunc(func(params operations.GetUsersUserIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetUsersUserID has not yet been implemented")
	})
	api.PatchFriendsUserIDFriendIDHandler = operations.PatchFriendsUserIDFriendIDHandlerFunc(func(params operations.PatchFriendsUserIDFriendIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PatchFriendsUserIDFriendID has not yet been implemented")
	})
	api.PostFriendsUserIDFriendIDHandler = operations.PostFriendsUserIDFriendIDHandlerFunc(func(params operations.PostFriendsUserIDFriendIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostFriendsUserIDFriendID has not yet been implemented")
	})
	api.PostUsersHandler = operations.PostUsersHandlerFunc(func(params operations.PostUsersParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostUsers has not yet been implemented")
	})
	api.PutUsersUserIDHandler = operations.PutUsersUserIDHandlerFunc(func(params operations.PutUsersUserIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutUsersUserID has not yet been implemented")
	})

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
func configureServer(s *http.Server, scheme, addr string) {
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
