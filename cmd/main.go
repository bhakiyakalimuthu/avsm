package main

import "github.com/bhakiyakalimuthu/avsm"

func main()  {
	v := avsm.Vehicle{}

	v.SetStateTransitionRules()

	println(v.CurrentState())

}