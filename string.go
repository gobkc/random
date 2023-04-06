package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
)

// AlphaNumber mix of numbers and letters
func AlphaNumber(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	s := fmt.Sprintf(`%X`, b)
	return s[:length]
}

func Number[T int | int32 | int64 | string](length int) (t T) {
	//maxNum := int64(math.Pow10(length) - 1)
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if reflect.TypeOf(t).Kind() == reflect.String {
		reflect.ValueOf(&t).Elem().SetString(result.String())
	} else {
		return T(result.Int64())
	}
	return
}
