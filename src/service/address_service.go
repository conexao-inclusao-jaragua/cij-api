package service

import (
	"cij_api/src/model"
	"cij_api/src/repo"
	"errors"
)

type AddressService interface {
	GetAddressById(id int) (model.Address, error)
}

type addressService struct {
	addressRepo repo.AddressRepo
}

func NewAddressService(addressRepo repo.AddressRepo) AddressService {
	return &addressService{
		addressRepo: addressRepo,
	}
}

func (n *addressService) GetAddressById(id int) (model.Address, error) {
	address, err := n.addressRepo.GetAddressById(id)
	if err != nil {
		return address, errors.New("failed to get address")
	}

	return address, nil
}
