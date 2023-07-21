package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "ABCDEFGHIJKLNMOPQRSTUVWXYZabcdefghijklnmopqrstuvwxyz"
const fullKeyword = "ABCDEFGHIJKLNMOPQRSTUVWXYZabcdefghijklnmopqrstuvwxyz0123456789"
const full = "ABCDEFGHIJKLNMOPQRSTUVWXYZabcdefghijklnmopqrstuvwxyz0123456789!#$%&+-=?/._"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(maxInt int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxInt)
}

// RandomInt generates a random integer between min and max
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomScreenName id generates a random string of length n
func RandomScreenNameID(n int) string {
	var sb strings.Builder
	k := len(fullKeyword)

	for i := 0; i < n; i++ {
		c := fullKeyword[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomToken id generates a random string of length n
func RandomToken(n int) string {
	var sb strings.Builder
	k := len(full)

	for i := 0; i < n; i++ {
		c := full[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomOwner generates a random owner name
func RandomScreenName() string {
	return RandomScreenNameID(15)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomPinCode() string {
	var ints strings.Builder
	for i := 0; i < 6; i++ {
		s := RandomInt(9)
		ints.WriteString(strconv.Itoa(s))
	}
	return ints.String()
}
