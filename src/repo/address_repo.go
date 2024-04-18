package repo

import (
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type AddressRepo interface {
	GetAddressById(id int) (model.Address, error)
	UpsertAddress(address model.Address) (int, error)
	DeleteAddress(id int) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) AddressRepo {
	return &addressRepo{
		db: db,
	}
}

func (n *addressRepo) GetAddressById(id int) (model.Address, error) {
	var address model.Address

	err := n.db.Model(model.Address{}).Where("id = ?", id).Find(&address).Error
	if err != nil {
		return address, errors.New("failed to get the address")
	}

	return address, nil
}

func (n *addressRepo) UpsertAddress(address model.Address) (int, error) {
	if err := n.db.Save(&address).Error; err != nil {
		return 0, errors.New("failed to upsert address")
	}

	return address.Id, nil
}

func (n *addressRepo) DeleteAddress(id int) error {
	if err := n.db.Model(model.Address{}).Where("id = ?", id).Delete(&model.Address{}).Error; err != nil {
		return errors.New("failed to delete address")
	}

	return nil
}
