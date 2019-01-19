package fs_tools_number

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GetRndNumber(size int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(size)+"v", rnd().Int31n(int32(math.Pow(10, float64(size)))))
}

func rnd() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+{}|:?><~"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
