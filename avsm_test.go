package avsm

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func TestHappyJourney(t *testing.T) {
	v:=Vehicle{}
	v.SetStateTransitionRules()
	
	assert.Equal(t,string(v.CurrentState()), "Ready")

	v.StateTransition(Riding,Hunter)
	assert.Equal(t,string(v.CurrentState()), "Riding")

	v.StateTransition(Ready,EndUser)
	assert.Equal(t,string(v.CurrentState()), "Ready")

	v.StateTransition(Bounty,Automatic)
	assert.Equal(t,string(v.CurrentState()), "Bounty")

	v.StateTransition(Collected,Hunter)
	assert.Equal(t,string(v.CurrentState()), "Collected")

	v.StateTransition(Dropped,Hunter)
	assert.Equal(t,string(v.CurrentState()), "Dropped")

	v.StateTransition(Ready,Hunter)
	assert.Equal(t,string(v.CurrentState()), "Ready")

	v.StateTransition(Unknown,Automatic)
	assert.Equal(t,string(v.CurrentState()), "Unknown")

	v.StateTransition(Terminated,Admin)
	assert.Equal(t,string(v.CurrentState()), "Terminated")

	v.StateTransition(ServiceMode,Admin)
	assert.Equal(t,string(v.CurrentState()), "ServiceMode")
}

func TestPanicStateTransitionValidation(t *testing.T) {
	v:=Vehicle{}
	v.SetStateTransitionRules()
	

	err := v.StateTransition(Collected,Hunter)
	assert.Equal(t,err, newErrorStruct(fmt.Sprintf("transition from state %s to %s is not permitted", Ready, Collected),ErrorTransitionNotPermitted))


	v.StateTransition(Riding,Hunter)

	err1 := v.StateTransition(Dropped,Hunter)
	assert.Equal(t,err1, newErrorStruct(fmt.Sprintf("transition from state %s to %s is not permitted", Riding, Dropped),ErrorTransitionNotPermitted))


}

func TestPanicStateRolePermission(t *testing.T) {
	
	v:=Vehicle{}
	v.SetStateTransitionRules()
	
	v.StateTransition(Riding,EndUser)
	err := v.StateTransition(BatteryLow,Hunter)
	assert.Equal(t,err, newErrorStruct(fmt.Sprintf("Invalid permission transition from state %s to %s for a role %s", Riding, BatteryLow,Hunter),ErrorRolePermissionDenied))


	v.StateTransition(BatteryLow,Automatic)

	err1 := v.StateTransition(Bounty,EndUser)
	assert.Equal(t,err1, newErrorStruct(fmt.Sprintf("Invalid permission transition from state %s to %s for a role %s", BatteryLow, Bounty,EndUser),ErrorRolePermissionDenied))

}

