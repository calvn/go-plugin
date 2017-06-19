// The plugin package exposes functions and helpers for communicating to
// plugins which are implemented as standalone binary applications.
//
// plugin.Client fully manages the lifecycle of executing the application,
// connecting to it, and returning the RPC client for dispensing plugins.
//
// plugin.Serve fully manages listeners to expose an RPC server from a binary
// that plugin.Client can connect to.
package plugin

import (
	"net/rpc"

	"google.golang.org/grpc"
)

// Plugin is the interface that is implemented to serve/connect to an
// inteface implementation.
type Plugin interface {
	// Server should return the RPC server compatible struct to serve
	// the methods that the Client calls over net/rpc.
	Server(*MuxBroker) (interface{}, error)

	// Client returns an interface implementation for the plugin you're
	// serving that communicates to the server end of the plugin.
	Client(*MuxBroker, *rpc.Client) (interface{}, error)
}

// GRPCPlugin is the interface that is implemented to serve/connect to
// a plugin over gRPC.
type GRPCPlugin interface {
	// GRPCServer should register this plugin for serving with the
	// given GRPCServer. Unlike Plugin.Server, this is only called once
	// since gRPC plugins serve singletons.
	GRPCServer(*grpc.Server) error

	// GRPCClient should return the interface implementation for the plugin
	// you're serving via gRPC.
	GRPCClient(*grpc.ClientConn) (interface{}, error)
}
