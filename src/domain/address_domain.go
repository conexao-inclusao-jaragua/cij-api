package domain

import "cij_api/src/model"

type AddressRepo interface {
	GetAddressById(id int) (model.Address, error)
	UpsertAddress(address model.Address) (int, error)
	DeleteAddress(id int) error
}

type AddressService interface {
	GetAddressById(id int) (model.Address, error)
}
