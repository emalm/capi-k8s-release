package capi_model

import (
	"encoding/json"

	kpack_build "github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

const BuildStagedState = "STAGED"
const BuildFailedState = "FAILED"
const KpackLifecycleType = "kpack"

// Build represents the payload that will be sent to CAPI when a kpack
// Build has been updated.
type Build struct {
	State     string    `json:"state"`
	Error     string    `json:"error"`
	Lifecycle Lifecycle `json:"lifecycle"`
}

type Lifecycle struct {
	Type string        `json:"type"`
	Data LifecycleData `json:"data"`
}
type LifecycleData struct {
	Image string `json:"image"`
}

func (b *Build) ToJSON() []byte {
	j, _ := json.Marshal(b)
	return j
}

func NewBuild(build *kpack_build.Build) Build {
	return Build{
		State: BuildStagedState,
		Lifecycle: Lifecycle{
			Type: KpackLifecycleType,
			Data: LifecycleData{
				Image: build.Status.LatestImage,
			},
		},
	}
}
