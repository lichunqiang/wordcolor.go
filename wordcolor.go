/*
This package help you convert word to the same rgb

Tanks to javascript version which original created by afc163
*/
package wordcolor

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	MAGIC_NUMBER = 5
	COLOR_LIMIT  = 242
)

//calculate the word rgb color
func WordColor(word string) string {
	rgb := GetRGB(word)

	return fmt.Sprintf("rgb(%s)", strings.Join([]string{rgb[0], rgb[1], rgb[2]}, ", "))
}

//Just get the original rgb slice
func GetRGB(word string) [3]string {
	word = strings.Trim(word, " ")
	rgb := [3]int{0, 0, 0}
	bs := []rune(word)
	length := len(bs)

	for i := 0; i < length; i++ {
		level := i / len(rgb)

		rgb[i%3] += int(getHashNum(string(bs[i:i+1])) / getRatio(float64(level)))
	}

	res := [3]string{}
	for key, val := range rgb {
		if val > 255 {
			res[key] = "255"
		} else {
			res[key] = strconv.Itoa(val)
		}
	}
	return res
}

//Get the ratio
func getRatio(level float64) float64 {
	return math.Pow(MAGIC_NUMBER, level)
}

//Get hash number.
func getHashNum(c string) float64 {
	ca := getCharCodeAt(c, 0)
	return float64((ca << MAGIC_NUMBER) % COLOR_LIMIT)
}

// get char code at, just as javascript.
// should know the `strings` pacakge
// 一个汉字在UTF-8中占3个字节
func getCharCodeAt(str string, index int) int {
	//need know the `rune`
	a := []rune(str)

	return int(a[index])
}
