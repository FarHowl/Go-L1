package main

import "fmt"

// We have two different objects OldAPI and NewAPI.
//Adapter contains an instance of the OldAPI, so we can initialize our NewAPI object and call getAllUsers()
type OldAPI struct{}

func (o *OldAPI) getAllUsersOld() {
	fmt.Println("All users - old version")
}

type NewAPI interface {
	getAllUsers()
}

type Adapter struct {
	oldAPI *OldAPI
}

func (a *Adapter) getAllUsers() {
	a.oldAPI.getAllUsersOld()
}

func main() {
	oldAPI := OldAPI{}
	adapter := &Adapter{oldAPI: &oldAPI}

	newAPI := NewAPI(adapter)

	newAPI.getAllUsers()
}
