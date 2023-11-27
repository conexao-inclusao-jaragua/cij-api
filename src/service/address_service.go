package service

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
)

type addressService struct {
	addressRepo domain.AddressRepo
}

func NewAddressService(addressRepo domain.AddressRepo) domain.AddressService {
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
