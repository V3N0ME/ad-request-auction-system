package errors

var (
	//DelayBidError is the error returned when BID_DELAY env variable is not set
	DelayBidError = New("BID_DELAY env variable not set or not an integer", 101)
)
