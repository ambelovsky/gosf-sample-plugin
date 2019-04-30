package gosfsampleplugin

import (
	"log"

	"github.com/ambelovsky/gosf"
)

func init() {
	gosf.OnConnect(func(client *gosf.Client, request *gosf.Request) {
		log.Println("Client connected.")
	})
	gosf.OnDisconnect(func(client *gosf.Client, request *gosf.Request) {
		log.Println("Client disconnected.")
	})
	gosf.OnBeforeRequest(func(client *gosf.Client, request *gosf.Request) {
		log.Println("Request received for " + request.Endpoint + " endpoint.")
	})
	gosf.OnAfterRequest(func(client *gosf.Client, request *gosf.Request, response *gosf.Message) {
		log.Println("Request for " + request.Endpoint + " endpoint was processed by the controller.")
	})
	gosf.OnBeforeResponse(func(client *gosf.Client, request *gosf.Request, response *gosf.Message) {
		log.Println("Response for " + request.Endpoint + " endpoint is being prepared.")
	})
	gosf.OnAfterResponse(func(client *gosf.Client, request *gosf.Request, response *gosf.Message) {
		log.Println("Response for " + request.Endpoint + " endpoint was sent.")
	})
	gosf.OnBeforeBroadcast(func(endpoint string, room string, response *gosf.Message) {
		log.Println("Broadcast for " + endpoint + " endpoint is preparing to send to " + getRoom(room) + ".")
	})
	gosf.OnAfterBroadcast(func(endpoint string, room string, response *gosf.Message) {
		log.Println("Broadcast for " + endpoint + " endpoint was sent to " + getRoom(room) + ".")
	})
	gosf.OnBeforeClientBroadcast(func(client *gosf.Client, endpoint string, room string, response *gosf.Message) {
		log.Println("Broadcast for " + endpoint + " endpoint is preparing to send to " + getRoom(room) + ".")
	})
	gosf.OnAfterClientBroadcast(func(client *gosf.Client, endpoint string, room string, response *gosf.Message) {
		log.Println("Broadcast for " + endpoint + " endpoint was sent to " + getRoom(room) + ".")
	})
}

func getRoom(room string) string {
	if room == "" {
		return "all"
	} else {
		return room
	}
}

/** ACCESSIBLE METHODS **/

// AppMethods is a struct optionally exposed to clients after the plugin has been registered
type AppMethods struct{}

// Echo returns the message it was given as a parameter
func (a AppMethods) Echo(message string) string {
	return message
}

/** ASPECT **/

// Plugin is the aspect oriented element required by the modular plugin framework
type Plugin struct{}

// Activate is an aspect-oriented modular plugin requirement
func (p Plugin) Activate(app *map[string]interface{}, config *map[string]interface{}) {}

// Deactivate is an aspect-oriented modular plugin requirement
func (p Plugin) Deactivate(app *map[string]interface{}, config *map[string]interface{}) {}
