package ecs

import (
	"testing"

	"github.com/compose-spec/compose-go/types"
	"gotest.tools/v3/assert"
)

func TestCapacityProvider(t *testing.T) {
	ami, machine := getUserDefinedMachine(types.ServiceConfig{
		Deploy: &types.DeployConfig{
			Placement: types.Placement{
				Constraints: []string{
					"node.ami == ami-12345",
					"node.machine == t2.micro",
				},
			},
		},
	})
	assert.Equal(t, ami, "ami-12345")
	assert.Equal(t, machine, "t2.micro")
}
