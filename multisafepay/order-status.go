package multisafepay

// The constants below correspond to all possible transaction statuses.
// See: https://docs.multisafepay.com/api/#transaction-statuses
const (
	StatusCompleted       = "completed"
	StatusInitialized     = "initialized"
	StatusDeclined        = "declined"
	StatusCancelled       = "cancelled"
	StatusVoid            = "void"
	StatusUncleared       = "uncleared"
	StatusRefunded        = "refunded"
	StatusPartialRefunded = "partial_refunded"
	StatusReserved        = "reserved"
	StatusChargedBack     = "chargedback"
	StatusShipped         = "shipped"
)

// IsValidStatus returns true if the status string is one of the possible responses as mentioned in the MultiSafePay documentation
// See: https://docs.multisafepay.com/api/#transaction-statuses
func IsValidStatus(status string) bool {
	switch status {
	case StatusCompleted, StatusInitialized, StatusDeclined, StatusCancelled, StatusVoid, StatusUncleared, StatusRefunded, StatusPartialRefunded, StatusReserved, StatusChargedBack, StatusShipped:
		return true
	default:
		return false
	}
}
