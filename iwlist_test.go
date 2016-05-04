package iwlist

import (
	"reflect"
	"testing"
)

func TestBestQuality(t *testing.T) {
	ap := aps.BestQuality()
	if reflect.DeepEqual(ap5, *ap) {
		t.Logf("Expected %v, but %v", ap5, *ap)
	}
}

func TestBestSignalLevel(t *testing.T) {
	ap := aps.BestSignalLevel()
	if reflect.DeepEqual(ap3, *ap) {
		t.Logf("Expected %v, but %v", ap3, *ap)
	}
}
