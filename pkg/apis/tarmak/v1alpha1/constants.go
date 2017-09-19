package v1alpha1

import (
	"time"
)

const (
	StackNameState      = "state"
	StackNameNetwork    = "network"
	StackNameTools      = "tools"
	StackNameVault      = "vault"
	StackNameKubernetes = "kubernetes"
)

const (
	ImageTagEnvironment   = "tarmak_environment"
	ImageTagBaseImageName = "tarmak_base_image_name"
)

var KubernetesEpoch time.Time = time.Unix(1437436800, 0)