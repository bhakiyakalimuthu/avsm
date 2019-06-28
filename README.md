# Abstract Vehicle State Machine
```
GO implementaion of State transition for a abstract vehicle.
State transitions are performed based on the permission of user roles.

```

# Technical requirements
- If the state transition is not valid, avsm return a descriptive error.
- If the state transition is valid, avsm return a nil error
- The library needs to have a reasonable performance to be used in a soft real­time API solution.
- The solution should include the git history.
- The solution must be stateless. Assume that any required state will be provided to the library.

# User roles
- End­users ­ Regular app­users / riders.
- Hunters ­ End users who have signed up to be chargers of vehicles and are responsible for
- picking up low battery vehicles.
- 3. Admins ­ Super users who can do everything

# Valid states
> Operational statutes
- 1. Ready ­ The vehicle is operational and can be claimed by an end­user
- 2. Battery_low ­ The vehicle is low on battery but otherwise operational. The vehicle cannot be
- claimed by an end­user but can be claimed by a hunter.
- 3. Bounty ­ Only available for “Hunters” to be picked up for charging.
- 4. Riding ­ An end user is currently using this vehicle; it can not be claimed by another user or hunter.
- 5. Collected ­ A hunter has picked up a vehicle for charging.
- 6. Dropped ­ A hunter has returned a vehicle after being charged.
> Not commissioned for service , not claimable by either end­users nor hunters.
- 7. Service_mode
- 8. Terminated
- 9. Unknown

# Note
```
 This is a drafted and working version.
 This can be improved by implementing commandline functionalities.
 avsm pacakge can be divided into multiple modules.
 Hardcoded transition map can be modified via logic.
```
# Installation
```
go get github.com/bhakiyakalimuthu/avsm

```
# Usage
```go
package main 

import (
	"github.com/bhakiyakalimuthu/avsm"
	"fmt"
	)

type T struct {
	V *avsm.Vehicle
}

func main() {
	fmt.Println("testing")
	t := &T{V: &avsm.Vehicle{}}

	// add initial rule 
	t.V.SetStateTransitionRules() // Ready state initiated by default
	// if err != nil {
	// 	// handle
	// }

	// Happy Journey 
	// All Happy state transition returns nill
	t.V.StateTransition("Riding","Hunter") // State is transited from Ready to Riding

	t.V.StateTransition("Ready","EndUser") // State is transited from Riding to Ready

	t.V.StateTransition("Bounty","Automatic") // State is transited from Ready to Bounty

	t.V.StateTransition("Collected","Hunter") // State is transited from Bounty to Collected

	t.V.StateTransition("Dropped","Hunter") // State is transited from Collected to Dropped

	t.V.StateTransition("Ready","Hunter") // State is transited from Dropped to Ready


	// Panic scenario 
	t.V.StateTransition("Riding","EndUser")
	errTransitionNotPermitted := t.V.StateTransition("Collected","Hunter") // transition from state Ready to collected  is not permitted.ErrorTransitionNotPermitted  
	fmt.Println(errTransitionNotPermitted)
	
	errRolePermissionDenied := t.V.StateTransition("BatteryLow","EndUser") // Invalid permission transition from state Riding to BatteryLow for a role EndUser.ErrorRolePermissionDenied
	fmt.Println(errRolePermissionDenied)

	// get the current state
	state := t.V.CurrentState() // "BatteryLow"
	fmt.Println(state)
}

```