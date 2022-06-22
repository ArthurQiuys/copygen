// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/_tests/cyclic/domain"
	"github.com/switchupcb/copygen/examples/_tests/cyclic/duplicate"
	"github.com/switchupcb/copygen/examples/_tests/cyclic/models"
)

// ModelsToDomain copies a *models.Account, *models.User to a *domain.Account.
func ModelsToDomain(tA *domain.Account, fA *models.Account, fU *models.User) {
	// *domain.Account fields
	tA.ID = fA.ID
	tA.Name = fA.Name
	tA.Info.UserID = fA.Info.UserID
	tA.Info.Username = fA.Info.Username
}

// DuplicateCyclic copies a *duplicate.Account to a *domain.Account.
func DuplicateCyclic(tA *domain.Account, fA *duplicate.Account) {
	// *domain.Account fields
	tA.ID = fA.ID
	tA.Name = fA.Name
	tA.Email = fA.Email
	tA.Info.UserID = fA.Info.UserID
	tA.Info.Username = fA.Info.Username
	tA.Info.Account = fA.Info.Account
	tA.Owner = fA.Owner
}

// SuperCyclic copies a domain.CyclicInterface to a *domain.CyclicInterface.
func SuperCyclic(tC *domain.CyclicInterface, fC domain.CyclicInterface) {
	// *domain.CyclicInterface fields
	tC = &fC
}
