package global

const (
	FIT_VERSION string = "v0.2.1"

	DEFAULT_FORTINET_ADDR string = "192.168.10.1:1003"
	DEFAULT_REQ_TIMEOUT   int    = 15
	DEFAULT_REFRESH_TIME  int    = 10800
	DEFAULT_MAX_RETRIES   int    = 10
	DEFAULT_TERM_TIME     string = "4:30:0" // Terminate old session at 04:30:0 every day

	TEST_TARGET string = "https://web.skype.com"
	F_AUTH      string = "fgtauth"
	F_ALIVE     string = "keepalive"
	F_LOGOUT    string = "logout"
	WAIT_TIME   int    = 3
)
