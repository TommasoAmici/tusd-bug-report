package web

// MaxMemory in bytes for use in ParseMultipartForm
const MaxMemory = 10000000 // 10MB

// forms input names
const (
	FormNameGuestEmail    = "guest-email"
	FormNameGuestCheckout = "guest"
	FormNameService       = "service"
)

// order summary tabs names
const (
	OrderSummaryTabPayment = iota
	OrderSummaryTabUpload
	OrderSummaryTabFirstRevision
	OrderSummaryTabSecondRevision
	OrderSummaryTabReview
)
