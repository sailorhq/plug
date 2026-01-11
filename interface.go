package plug

type HookStatus struct {
	Ok      bool
	Message string
}

type DeployRequest struct {
	Ns      string
	App     string
	Content []byte
}

type DeployResponse struct {
	Status HookStatus
}

type SailorHook interface {
	OnDeploy(req DeployRequest) DeployResponse
}
