package global

type GlobalConst struct{}

const (
	NumOne = iota + 1
	NumTwo
	NumThree
	NumFour
	NumFive
	NumSix
	NumSeven
	NumEight
	NumNine
	NumTen
)

const (
	SuccessCode  = 10000
	ErrorCode    = 9999
	InvalidToken = iota + 9000
	InvalidParams
	ExpireToken
)

const (
	TokenPrefix        = "BLOG_USER_TOKEN_"
	TokenExpireSeconds = 3600
)
