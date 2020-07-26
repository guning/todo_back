package errno

var (
	//common error
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validate failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Token error."}

	//user error
	ErrUserNotFound      = &Errno{Code: 20102, Message: "User not found."}
	ErrEncrypt           = &Errno{Code: 20101, Message: "Encrypt error."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "Token invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "Password incorrect"}

	//task error
	ErrTaskCreate      = &Errno{Code: 20201, Message: "Task create failed."}
	ErrTaskUpdate      = &Errno{Code: 20202, Message: "Task update failed."}
	ErrTaskNotFound      = &Errno{Code: 20203, Message: "Task not found."}
)