package app_status

const (
	Success                 = "000"
	NoRecordFound           = "001"
	StatusNotYetImplemented = "X00"
	GeneralError            = "X06"
	ErrorNotMatchValidation = "X01"
	ErrorLackInfo           = "X02"
	ErrorUnauthorized       = "X04"
)

var statusMessage = map[string]string{
	Success:                 "Success",
	NoRecordFound:           "No Record Found",
	StatusNotYetImplemented: "Not yet implemented",
	GeneralError:            "General error",
	ErrorNotMatchValidation: "Unsatisfied validation",
	ErrorLackInfo:           "Please fill required %s",
	ErrorUnauthorized:       "User unauthorized",
}

func StatusText(code string) string {
	return statusMessage[code]
}
