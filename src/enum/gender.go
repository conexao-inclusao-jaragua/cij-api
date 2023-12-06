package enum

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
	Other  GenderEnum = "other"
)

func (g GenderEnum) IsValid() bool {
	return g == Male || g == Female || g == Other
}
