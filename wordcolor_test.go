package wordcolor

import (
	"testing"
)

//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh%2F11.3.md
func TestConvert(t *testing.T) {
	txt := "Hello world"
	_fixture := "rgb(255, 56, 13)"
	rgb := WordColor(txt)
	if _fixture != rgb {
		t.Error(rgb)
	}
}

func TestGetRGB(t *testing.T) {
	txt := "Hello world"
	_fixture := [3]string{"255", "56", "13"}
	rgb := GetRGB(txt)
	if _fixture != rgb {
		t.Error(rgb)
	}
}

func TestUnicodeConvert(t *testing.T) {
	txt := "李牧"
	_fixture := "rgb(255, 0, 0)"
	if _fixture != WordColor(txt) {
		t.Error("Convert unicode char error")
	}
}
