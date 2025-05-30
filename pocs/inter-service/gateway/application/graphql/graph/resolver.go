package graph

import "github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/core"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	MessageService core.MessageService
}
