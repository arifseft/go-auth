package flag

type LoginFlag struct {
	Message string
	Flag    string
	Error
}

var (
	LoginSuccess LoginFlag = LoginFlag{
		Message: "Successfully logged in",
	}

	LogoutSuccess LoginFlag = LoginFlag{
		Message: "Successfully logged out",
	}

	LogoutUnauthorized LoginFlag = LoginFlag{
		Message: "You need to be authorized to access this route",
	}
)
