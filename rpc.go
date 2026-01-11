package plug

import "net/rpc"

type HookRPCServer struct{ Impl SailorHook }

func (hs *HookRPCServer) OnDeploy(args DeployRequest, resp *DeployResponse) error {
	*resp = hs.Impl.OnDeploy(args)
	return nil
}

type HookRPCClient struct{ client *rpc.Client }

func (hc *HookRPCClient) OnDeploy(req DeployRequest) DeployResponse {
	var resp DeployResponse
	_ = hc.client.Call("Plugin.OnDeploy", req, &resp)
	return resp
}
