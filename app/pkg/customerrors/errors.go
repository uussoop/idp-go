package customerrors

type Customerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	Success       Customerror = Customerror{Status: 200, Message: "operation done successfully"}
	SuccessVerify Customerror = Customerror{Status: 200, Message: "verification done successfully"}
)

var (
	BadBodyRequest  Customerror = Customerror{Status: 1000, Message: "request body is invalid"}
	InvalidAddress  Customerror = Customerror{Status: 1001, Message: "invalid public address"}
	NonceGeneration Customerror = Customerror{Status: 1002, Message: "error in generating nonce"}
	Usercreation    Customerror = Customerror{Status: 1003, Message: "error in creating a user"}
	NoncCheck       Customerror = Customerror{Status: 1004, Message: "error in checking nonce"}
	UserNotFound    Customerror = Customerror{Status: 1005, Message: "user not found"}
	SigMissing      Customerror = Customerror{Status: 1006, Message: "signature not found"}
	SigToPub        Customerror = Customerror{Status: 1007, Message: "signature failure"}
	AuthFailure     Customerror = Customerror{Status: 1008, Message: "authentication error address and recoveraddress dont match"}
	Unauthorized    Customerror = Customerror{Status: 1009, Message: "unauthorized access"}
)

func CreateCustomError(status int, message string) Customerror {
	return Customerror{Status: status, Message: message}
}
