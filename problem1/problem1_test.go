package problem1

import "testing"

func TestStrToMath(t *testing.T) {

	if strToMath("1+2+3") != 6 {
		t.Errorf("SALAH! harusnya 6")
    }
    
	if strToMath("1+2-3") != 0 {
		t.Errorf("SALAH! harusnya 0")
    }
    
	if strToMath("1+2-4") != -1 {
		t.Errorf("SALAH! harusnya -1")
	}
}
