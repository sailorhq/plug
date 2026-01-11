package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SAILOR_MAGIC_COOKIE",
	MagicCookieValue: "jkda2981dASA219",
}

var PluginMap = map[string]plugin.Plugin{
	"hook": &SailorPlugin{},
}

type SailorPlugin struct {
	plugin.Plugin

	Impl SailorHook
}

// Server should return the RPC server compatible struct to serve
// the methods that the Client calls over net/rpc.
func (sp *SailorPlugin) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &HookRPCServer{Impl: sp.Impl}, nil
}

// Client returns an interface implementation for the plugin you're
// serving that communicates to the server end of the plugin.
func (sp *SailorPlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return HookRPCClient{client: c}, nil
}
