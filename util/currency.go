package util

const (
	USD = "USD"
	EUR = "EUR"
	CNY = "CNY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CNY:
		return true
	}
	return false
}
