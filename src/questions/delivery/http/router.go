package questiondelivery

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func mw(ct http.HandlerFunc) *negroni.Negroni {
	//jwtMiddleware := negroni.HandlerFunc(middlewares.ValidateToken)
	return negroni.New(negroni.WrapFunc(ct))
}

func (t QuestionDelivery) Start(router *mux.Router) {

	questionRoute := router.PathPrefix("/v1/questions").Subrouter()
	questionRoute.Handle("/", mw(t.CreateQuestion)).Methods("POST")
	questionRoute.Handle("/", mw(t.FetchQuestions)).Methods("GET")
	questionRoute.Handle("/{questionID}", mw(t.GetQuestion)).Methods("GET")
	questionRoute.Handle("/{questionID}", mw(t.UpdateQuestion)).Methods("PUT")
questionRoute.Handle("/{questionID}", mw(t.DeleteQuestion)).Methods("DELETE")

}

