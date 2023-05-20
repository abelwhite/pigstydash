// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave, noSurf)
	//we wrap

	router.Handler(http.MethodGet, "/signup", dynamicMiddleware.ThenFunc(app.signup))
	router.Handler(http.MethodPost, "/signup", dynamicMiddleware.ThenFunc(app.signupSubmit))
	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.login))
	router.Handler(http.MethodPost, "/login", dynamicMiddleware.ThenFunc(app.loginSubmit))
	router.Handler(http.MethodGet, "/logout", dynamicMiddleware.ThenFunc(app.logoutSubmit))

	protected := dynamicMiddleware.Append(app.requireAuthenticationMiddleware)

	router.Handler(http.MethodGet, "/dashboard", protected.ThenFunc(app.dashboard))
	router.Handler(http.MethodGet, "/setting", protected.ThenFunc(app.settings))
	router.Handler(http.MethodGet, "/profile", protected.ThenFunc(app.profile))

	router.Handler(http.MethodGet, "/pig/create", protected.ThenFunc(app.pigCreateShow))
	router.Handler(http.MethodPost, "/pig/create", protected.ThenFunc(app.pigCreateSubmit))
	router.Handler(http.MethodGet, "/pig/show", protected.ThenFunc(app.pigShow))
	router.Handler(http.MethodGet, "/pig/delete", protected.ThenFunc(app.pigDelete))
	router.Handler(http.MethodGet, "/pig/update", protected.ThenFunc(app.pigUpdate))
	router.Handler(http.MethodPost, "/pig/update", protected.ThenFunc(app.pigUpdateQuery))

	router.Handler(http.MethodGet, "/room/create", protected.ThenFunc(app.roomCreateShow))
	router.Handler(http.MethodPost, "/room/create", protected.ThenFunc(app.roomCreateSubmit))
	router.Handler(http.MethodGet, "/room/show", protected.ThenFunc(app.roomShow))
	router.Handler(http.MethodGet, "/room/delete", protected.ThenFunc(app.roomDelete))
	router.Handler(http.MethodGet, "/room/update", protected.ThenFunc(app.roomUpdate))
	router.Handler(http.MethodPost, "/room/update", protected.ThenFunc(app.roomUpdateQuery))

	router.Handler(http.MethodGet, "/pigsty/create", protected.ThenFunc(app.pigstyCreateShow))
	router.Handler(http.MethodPost, "/pigsty/create", protected.ThenFunc(app.pigstyCreateSubmit))
	router.Handler(http.MethodGet, "/pigsty/show", protected.ThenFunc(app.pigstyShow))
	router.Handler(http.MethodGet, "/pigsty/delete", protected.ThenFunc(app.pigstyDelete))
	router.Handler(http.MethodGet, "/pigsty/update", protected.ThenFunc(app.pigstyUpdate))
	router.Handler(http.MethodPost, "/pigsty/update", protected.ThenFunc(app.pigstyUpdateQuery))

	router.Handler(http.MethodGet, "/temperature", protected.ThenFunc(app.temperature))
	router.Handler(http.MethodGet, "/humidity", protected.ThenFunc(app.humidity))
	router.Handler(http.MethodGet, "/feedbin", protected.ThenFunc(app.feedbin))
	router.Handler(http.MethodGet, "/waterbin", protected.ThenFunc(app.waterbin))

	//returns to the router to our middleware interesting in between go server and mux
	//Client -> Goserver ->Middleware -> Router -> Handler

	//tidy up the middle wear
	standardMiddleware := alice.New(app.RecoverPanicMiddleware, app.logRequestMiddleware, securityHeadersMiddleware)

	return standardMiddleware.Then(router)
}
