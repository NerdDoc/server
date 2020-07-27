package server

import (
	"encoding/json"
	"fmt"
	"mod/request"
	"net/http"
	"time"

	"github.com/NerdDoc/server/training"

	"github.com/NerdDoc/server/dashboard"

	"github.com/NerdDoc/server/network"
	"github.com/gookit/color"
	"github.com/julienschmidt/httprouter"
	gocache "github.com/patrickmn/go-cache"
)

var (
	// Create the neural network variable to use it everywhere
	neuralNetworks map[string]network.Network
	// Initializes the cache with a 5 minute lifetime
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

// Serve serves the server in the given port
func Serve(_neuralNetworks map[string]network.Network, port string) {
	// Set the current global network as a global variable
	neuralNetworks = _neuralNetworks

	// Initializes the router
	//router := mux.NewRouter()
	//router.HandleFunc("/callback", spotify.CompleteAuth)
	//// Serve the websocket
	//router.HandleFunc("/websocket", SocketHandle)
	//// Handle local connection
	//router.HandleFunc("/api/v1/request", request.Req).Methods("POST")
	// Serve the API
	//router.HandleFunc("/api/{locale}/dashboard", GetDashboardData).Methods("GET")
	//router.HandleFunc("/api/{locale}/intent", dashboard.CreateIntent).Methods("POST")
	//router.HandleFunc("/api/{locale}/intent", dashboard.DeleteIntent).Methods("DELETE", "OPTIONS")
	//router.HandleFunc("/api/{locale}/train", Train).Methods("POST")
	//router.HandleFunc("/api/{locale}/intents", dashboard.GetIntents).Methods("GET")
	//router.HandleFunc("/api/coverage", analysis.GetCoverage).Methods("GET")

	//magenta := color.FgMagenta.Render
	//fmt.Printf("\nServer listening on the port %s...\n", magenta(port))
	//
	// Serves the chat
	//err := http.ListenAndServe("0.0.0.0:"+port, router)
	//if err != nil {
	//	panic(err)
	//}
	router := httprouter.New()
	router.POST("/api/v1/request", request.Req)
	//router.POST("/api/v1/hello/:name", Hello)
	fmt.Println("starting server")
	http.ListenAndServe("0.0.0.0:8888", router)
}

// Train is the route to re-train the neural network
func Train(w http.ResponseWriter, r *http.Request) {
	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Olivia-Token")
	if !dashboard.ChecksToken(token) {
		json.NewEncoder(w).Encode(dashboard.Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	magenta := color.FgMagenta.Render
	fmt.Printf("\nRe-training the %s..\n", magenta("neural network"))

	for locale := range neuralNetworks {
		neuralNetworks[locale] = training.CreateNeuralNetwork(locale, true)
	}
}
