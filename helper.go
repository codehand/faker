package faker

import (
	"math"
	"math/rand"
)

const hashtag = '#'
const questionmark = '?'

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	if len(dataVal) == 2 {
		_, checkOk = Data[dataVal[0]]
		if checkOk {
			_, checkOk = Data[dataVal[0]][dataVal[1]]
		}
	}

	return checkOk
}

// Get Random Value
func getRandValue(dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return Data[dataVal[0]][dataVal[1]][rand.Intn(len(Data[dataVal[0]][dataVal[1]]))]
}

// Replace # with numbers
func replaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = byte(randDigit())
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(rand.Intn(8)+1) + '0'
	}

	return string(bytestr)
}

// Replace ? with ASCII lowercase letters
func replaceWithLetters(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(randLetter())
		}
	}

	return string(bytestr)
}

// Generate random lowercase ASCII letter
func randLetter() rune {
	return rune(byte(rand.Intn(26)) + 'a')
}

// Generate random ASCII digit
func randDigit() rune {
	return rune(byte(rand.Intn(10)) + '0')
}

// Generate random integer between min and max
func randIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

func randFloat32Range(min, max float32) float32 {
	if min == max {
		return min
	}
	return rand.Float32()*(max-min) + min
}

func randFloat64Range(min, max float64) float64 {
	if min == max {
		return min
	}
	return rand.Float64()*(max-min) + min
}

// Letter ...
func Letter() string {
	return string(randLetter())
}

// Digit ...
func Digit() string {
	return string(randDigit())
}

// Lexify ...
func Lexify(str string) string {
	return replaceWithLetters(str)
}

// ShuffleStrings ...
func ShuffleStrings(a []string) {
	swap := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	//to avoid upgrading to 1.10 I copied the algorithm
	n := len(a)
	if n <= 1 {
		return
	}

	//if size is > int32 probably it will never finish, or ran out of entropy
	i := n - 1
	for ; i > 0; i-- {
		j := int(rand.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

// RandString ...
func RandString(a []string) string {
	size := len(a)
	if size == 0 {
		return ""
	}
	return a[rand.Intn(size)]
}

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Uint8 will generate a random uint8 value
func Uint8() uint8 {
	return uint8(randIntRange(0, math.MaxUint8))
}

// Uint16 will generate a random uint16 value
func Uint16() uint16 {
	return uint16(randIntRange(0, math.MaxUint16))
}

// Uint32 will generate a random uint32 value
func Uint32() uint32 {
	return uint32(randIntRange(0, math.MaxInt32))
}

// Uint64 will generate a random uint64 value
func Uint64() uint64 {
	return uint64(rand.Int63n(math.MaxInt64))
}

// Int8 will generate a random Int8 value
func Int8() int8 {
	return int8(randIntRange(math.MinInt8, math.MaxInt8))
}

// Int16 will generate a random int16 value
func Int16() int16 {
	return int16(randIntRange(math.MinInt16, math.MaxInt16))
}

// Int32 will generate a random int32 value
func Int32() int32 {
	return int32(randIntRange(math.MinInt32, math.MaxInt32))
}

// Int64 will generate a random int64 value
func Int64() int64 {
	return rand.Int63n(math.MaxInt64) + math.MinInt64
}

// Float32 will generate a random float32 value
func Float32() float32 {
	return randFloat32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float32Range will generate a random float32 value between min and max
func Float32Range(min, max float32) float32 {
	return randFloat32Range(min, max)
}

// Float64 will generate a random float64 value
func Float64() float64 {
	return randFloat64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Float64Range will generate a random float64 value between min and max
func Float64Range(min, max float64) float64 {
	return randFloat64Range(min, max)
}

// Numerify will replace # with random numerical values
func Numerify(str string) string {
	return replaceWithNumbers(str)
}

// ShuffleInts will randomize a slice of ints
func ShuffleInts(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
