package model

type PersonDisability struct {
	Acquired     bool `gorm:"type:boolean;not null"`
	PersonId     int  `gorm:"type:int;not null"`
	DisabilityId int  `gorm:"type:int;not null"`
	Person       *Person
	Disability   *Disability
}

type PersonDisabilityResponse struct {
	Acquired bool `json:"acquired"`
	DisabilityResponse
}

func (pd *PersonDisability) ToResponse() PersonDisabilityResponse {
	return PersonDisabilityResponse{
		Acquired:           pd.Acquired,
		DisabilityResponse: pd.Disability.ToResponse(),
	}
}
