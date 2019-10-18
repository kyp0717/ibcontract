package client

const (
	BaseUrl = "https://localhost:5000/v1/portal"

	// market data
	MarketHistory  = "/iserver/marketdata/history"
	MarketSnapshot = "/iserver/marketdata/snapshot"
	// session
	SesseionValidate = "/sso/validate"
	SessionStatus    = "/iserver/auth/status"
	SessionReauth    = "/iserver/reauthenticate"
	SessionLogout    = "/logout"
	SessionTickle    = "/tickle"
	//ib cust
	CustInfo = "/ibcust/entity/info"
	// portfolio analysis
	Portfolio        = "/pa/performance"
	PortfolioSummary = "/pa/summary"
	// trades
	Trades = "/iserver/account/trades"
	//scanner
	ScanParams = "/iserver/scanner/params"
	ScanRun    = "/iserver/scanner/run"
	// profit and loss
	Pnl = "/iserver/account/pnl/partitioned"
	// portfolio
	PfolioAcctAlloc = "/portfolio/{accountId}/allocation"
	PfolioAcctPos   = "/portfolio/{accountId}/position/{conid}"
	PfolioAcctPage  = "/portfolio/{accountId}/positions/{pageId}"
	PfolioPos       = "/portfolio/positions/{conid}"
	/*"/portfolio/{accountId}/positions/invalidate" */
	//"/portfolio/allocation"
	//// order
	//"/iserver/account/{accountId}/order/{origCustomerOrderId}"
	//"/iserver/account/orders"
	//"/iserver/account/{accountId}/order"
	//"/iserver/account/{accountId}/order/{origCustomerOrderId}"
	//"/iserver/account/{accountId}/order/whatif"
	//"/iserver/account/{accountId}/orders"
	/*"/iserver/reply/{replyid}"*/
	// account endpoints
	// list accts
	Accts         = "/iserver/accounts"
	AcctLedger    = "/portfolio/{accountId}/ledger"
	AcctMeta      = "/portfolio/{accountId}/meta"
	AcctSumm      = "/portfolio/{accountId}/summary"
	PfolioAcct    = "/portfolio/accounts"
	PfolioSubAcct = "/portfolio/subaccounts"
	// post acct (set the acct for analysis)
	Acct = "/iserver/account"

// contract
/*"/iserver/contract/{conid}/info" */
//"/trsrv/futures"
//"/iserver/secdef/search"
//"/iserver/secdef/search"
//"/trsrv/secdef"
//"/iserver/contract/{conid}/info"
//"/trsrv/futures"
//"/iserver/secdef/search"
//"/trsrv/secdef"
//// fyi
//"/fyi/deliveryoptions/{deviceId}"
//"/fyi/deliveryoptions"
//"/fyi/disclaimer/{typecode}"
//"/fyi/notifications"
//"/fyi/notifications/more"
//"/fyi/settings"
//"/fyi/unreadnumber"
//"/fyi/deliveryoptions/device"
//"/fyi/settings/{typecode}"
//"/fyi/deliveryoptions/email"
//"/fyi/disclaimer/{typecode}"
/*"/fyi/notifications/{notificationId}"*/
)
