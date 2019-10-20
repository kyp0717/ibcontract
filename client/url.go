package client

const (
	BaseUrl = "https://localhost:5000/v1/portal"
	// contract
	ContractInfo    = "/iserver/contract/{conid}/info"
	ContractFutures = "/trsrv/futures"
	ContractSymbol  = "/iserver/secdef/search"
	ContractConId   = "/trsrv/secdef"
	// session
	SesseionValidate = "/sso/validate"
	SessionStatus    = "/iserver/auth/status"
	SessionReauth    = "/iserver/reauthenticate"
	SessionLogout    = "/logout"
	SessionTickle    = "/tickle"
)
