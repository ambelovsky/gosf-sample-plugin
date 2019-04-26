package gosfsampleplugin

import (
	f "github.com/ambelovsky/gosf"
)

func init() {}

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

// Connect is an aspect-oriented modular plugin requirement
func (p Plugin) Connect(request *f.Request) {
}

// Disconnect is an aspect-oriented modular plugin requirement
func (p Plugin) Disconnect(request *f.Request) {
}

// PreReceive is an aspect-oriented modular plugin requirement
func (p Plugin) PreReceive(request *f.Request, clientMessage *f.Message) {
}

// PostReceive is an aspect-oriented modular plugin requirement
func (p Plugin) PostReceive(request *f.Request, clientMessage *f.Message) {
}

// PreRespond is an aspect-oriented modular plugin requirement
func (p Plugin) PreRespond(request *f.Request, clientMessage *f.Message, serverMessage *f.Message) {
}

// PostRespond is an aspect-oriented modular plugin requirement
func (p Plugin) PostRespond(request *f.Request, clientMessage *f.Message, serverMessage *f.Message) {
}
