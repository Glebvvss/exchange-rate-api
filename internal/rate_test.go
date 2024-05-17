package internal

import (
	"github.com/Glebvvss/exchange-rate-api/internal"
	"testing"
)

// For getting rate we depends on external
// service and must be sure that it is available
func TestGetUsdRate(t *testing.T) {
	_, err := internal.GetUsdRate()
	if err != nil {
		t.Fatal(err)
	}
}
