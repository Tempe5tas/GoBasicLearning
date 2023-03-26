package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "Hello world"
	t.Log(str, len(str))
	//str[6] = '_'	//compile error, can't assign
	str = "逸一时，误一世"
	t.Log("Sentence length:", len(str)) //length is 21, 3 bytes for a Chinese character
	str = "\xE4\xB8\x96"                //UTF-8 code
	t.Log("UTF8 word and length:", str, len(str))

	rn := []rune(str)
	t.Log("UTF8 to rune length:", len(rn))
	t.Logf("Unicode: %X UTF8: %X", rn[0], str)
}

func TestCodingFormat(t *testing.T) {
	str := "逸一时，误一世"
	for _, e := range str {
		t.Logf("%[1]c\t%[1]X", e)
	}
}

func TestStringFunctions(t *testing.T) {
	str := "逸一时，误一世，逸久逸久罢已龄"
	strParts := strings.Split(str, "，")
	for _, e := range strParts {
		t.Log(e)
	}
	t.Log(strings.Join(strParts, "_"))

	str1 := strconv.Itoa(10) //Convert integer to string
	t.Log(str1)
	ret, err := strconv.Atoi(str1) //Convert string to integer, return integer and error code (if exist)
	if err != nil {
		t.Log("Convert failed")
	} else {
		t.Log(ret - 10)
	}
}
