package container

import (
	containerd "github.com/containerd/containerd/v2/client"
)

type HitherContainer struct {
	ID        string
	Container containerd.Container
	Task      containerd.Task
}

func NewHitherContainer(id string, container containerd.Container, task containerd.Task) *HitherContainer {
	return &HitherContainer{
		ID:        id,
		Container: container,
		Task:      task,
	}
}
