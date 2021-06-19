package recipe

import (
	"testing"
)

func TestDeliveryTimeUnmarshalValidJSON(t *testing.T) {
	var dt DeliveryTime
	target := []byte(`"Monday 12AM - 11PM"`)
	err := dt.UnmarshalJSON(target)
	if err != nil {
		t.Fatalf("should not return error: %v\n", err)
	}
	if dt.From != 0 || dt.To != 11 {
		t.Fatal("unmarshal result incorrect")
	}
}

func TestDeliveryTimeUnmarshalWrongTimePattern(t *testing.T) {
	var dt DeliveryTime
	target := []byte(`"Monster 3AM - BPM"`)
	err := dt.UnmarshalJSON(target)
	if err == nil {
		t.Fatal("should return unmarshal error")
	}
	if err.Error() != "unmatched time group" {
		t.Fatal("wrong error message")
	}
}
