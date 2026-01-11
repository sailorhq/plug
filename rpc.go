package plug

import "net/rpc"

type HookRPCServer struct{ Impl SailorHook }

func (hs *HookRPCServer) OnDeploy(args DeployRequest, resp *DeployResponse) error {
	*resp = hs.Impl.OnDeploy(args)
	return nil
}

func (hs *HookRPCServer) Health(args any, resp *DoneResponse) error {
	*resp = hs.Impl.Health()
	return nil
}

type HookRPCClient struct{ client *rpc.Client }

func (hc *HookRPCClient) OnDeploy(req DeployRequest) DeployResponse {
	var resp DeployResponse
	_ = hc.client.Call("Plugin.OnDeploy", req, &resp)
	return resp
}

func (hc *HookRPCClient) Health() DoneResponse {
	var resp DoneResponse
	_ = hc.client.Call("Plugin.OnDeploy", nil, &resp)
	return resp
}
