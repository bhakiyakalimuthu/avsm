package avsm

import (
	"fmt"
	"sync"
	
)

// State denotes different states of the vehicle. 
type State string
const (
	Ready       = "Ready"
	Riding      = "Riding"
	BatteryLow  = "BatteryLow"
	Bounty      = "Bounty"
	Collected   = "Collected"
	Dropped     = "Dropped"
	ServiceMode = "ServiceMode"
	Terminated  = "Terminated"
	Unknown     = "Unknown"
)
// Role denotes user roles who operates the vehicle
type Role string
const (
     Automatic =  "Automatic"
     Admin =     "Admin"
     EndUser=   "EndUser"
     Hunter=    "Hunter"

)
// Abstract Vehicle State Machine 
type Vehicle struct {
	state State
	mu sync.RWMutex
	transitions map[State]map[State][]Role 

}

// Initial state and Transition rules are created.Then set to vehicle struct.
func (v *Vehicle)SetStateTransitionRules(){

	v.mu.RLock()
	defer v.mu.RUnlock()
	// Entire valid transitions are created as State Map contains State Map of List of Roles
	transitions := map[State]map[State][]Role{
		Ready: map[State][]Role{
			Bounty:{Automatic},Riding:{Admin,EndUser,Hunter},Unknown:{Automatic}},
		BatteryLow:map[State][]Role{
			Bounty:{Automatic},},
		Bounty:map[State][]Role{
			Collected:{Hunter},}, 
		Riding:map[State][]Role{
			Ready:{Admin,EndUser,Hunter},BatteryLow:{Automatic},}, 
		Collected:map[State][]Role{
			Dropped:{Hunter},}, 
		Dropped:map[State][]Role{
			Ready:{Hunter},},
		Unknown:{},
		Terminated:{},
		ServiceMode:{},				 	
		
	}
	//Initial state is set to Ready and hardcoded
	v.state = Ready
	// Valid transitions are set to Vehicles transition
	v.transitions = transitions
	
}
// CurrentState returns the machine's current state. If the State returned is
// "", then the machine has not been given an initial state.
func (v *Vehicle) CurrentState() State {
	v.mu.RLock()
	defer v.mu.RUnlock()

	return v.state
}
//Main state transition logic 
func (v *Vehicle) StateTransition(toState State, role Role) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	// if this is nil we cannot assume any state.Though currently initial state is hardcoded
	// Below check is not called,will be used for future purpose
	if v.transitions == nil {
		return newErrorStruct("no states added to the Vehicle", ErrorVehicleNotInitialized)
	}

	// if the state is nothing, this is probably the initial state
	// Below check is not called,will be used for future purpose
	if v.state == "" {
		// if the state is not defined, it's invalid
		if _, ok := v.transitions[toState]; !ok {
			return newErrorStruct("the initial state has not been defined within the machine", ErrorStateUndefined)
		}

		// set the state
		v.state = toState
		return nil
	}

	// if the destination state was not defined...
	// Below check is not called,will be used for future purpose
	if _, ok := v.transitions[toState]; !ok {
		return newErrorStruct(fmt.Sprintf("state %s has not been registered", toState), ErrorStateUndefined)
	}

	// if we are not permitted to transition to this state...
	// Role Admin can perform any transition state
	if role != Admin {
		roles, ok := v.transitions[v.state][toState]; 
		if !ok {
			return newErrorStruct(fmt.Sprintf("transition from state %s to %s is not permitted", v.state, toState), ErrorTransitionNotPermitted)
		}

		for _, ok := range roles {
			if role == ok {
				// set the state
				v.state = toState
				// Return nil when transitions are valid
				return nil
			}
		}
		return newErrorStruct(fmt.Sprintf("Invalid permission transition from state %s to %s for a role %s", v.state, toState, role), ErrorRolePermissionDenied)
	}
	// set the state
	v.state = toState

	return nil
}
