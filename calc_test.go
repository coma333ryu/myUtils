package utils

import "testing"

func Test_AddNum (t *testing.T) {
	testVal := AddNum(1,2,3)

	if testVal != 7 {
		t.Error("Sum Error")
		t.Fatal("1+2+3 = 6 but ", testVal)
	}
}
