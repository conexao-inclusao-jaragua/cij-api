package utils

type Error struct {
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Fields  []string `json:"fields,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) GetCode() string {
	return e.Code
}

func (e Error) GetFields() []string {
	return e.Fields
}

func NewError(message string, code string) Error {
	return Error{
		Message: message,
		Code:    code,
	}
}

func NewErrorWithFields(message string, code string, fields []string) Error {
	return Error{
		Message: message,
		Code:    code,
		Fields:  fields,
	}
}

// User errors
var (
	FailedToCreateUser      = Error{Message: "failed to create user", Code: "ERR-1001"}
	FailedToUpdateUser      = Error{Message: "failed to update user", Code: "ERR-1002"}
	FailedToDeleteUser      = Error{Message: "failed to delete user", Code: "ERR-1003"}
	FailedToGetUser         = Error{Message: "failed to get user", Code: "ERR-1004"}
	FailedToEncryptPassword = Error{Message: "failed to encrypt password", Code: "ERR-1005"}
)

// Person errors
var (
	FailedToCreatePerson = Error{Message: "failed to create person", Code: "ERR-2001"}
	FailedToGetPerson    = Error{Message: "failed to get person", Code: "ERR-2002"}
	FailedToUpdatePerson = Error{Message: "failed to update person", Code: "ERR-2003"}
	FailedToDeletePerson = Error{Message: "failed to delete person", Code: "ERR-2004"}
	FailedToListPeople   = Error{Message: "failed to list people", Code: "ERR-2005"}
)

// Address errors
var (
	FailedToUpsertAddress = Error{Message: "failed to upsert address", Code: "ERR-3001"}
	FailedToGetAddress    = Error{Message: "failed to get address", Code: "ERR-3002"}
	FailedToDeleteAddress = Error{Message: "failed to delete address", Code: "ERR-3003"}
)

// Disability errors
var (
	FailedToUpsertDisability = Error{Message: "failed to upsert disability", Code: "ERR-4001"}
	FailedToGetDisability    = Error{Message: "failed to get disability", Code: "ERR-4002"}
	FailedToClearDisability  = Error{Message: "failed to clear disabilities", Code: "ERR-4003"}
)
