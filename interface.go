package plug

type DoneResponse struct {
	Ok      bool
	Message string
}

type DeployRequest struct {
	Ns      string
	App     string
	Content []byte
}

type DeployResponse struct {
	Status DoneResponse
}

type SailorHook interface {
	Health() DoneResponse
	OnDeploy(req DeployRequest) DeployResponse
}
