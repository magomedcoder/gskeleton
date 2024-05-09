package strutil

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"path"
	"strings"
	"time"
)

func GenValidateCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[newRand.Intn(10)])
	}

	return sb.String()
}

func Random(length int) string {
	var result []byte
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func GenImageName(ext string, width, height int) string {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), Random(10))

	return fmt.Sprintf("%x_%dx%d.%s", md5.Sum([]byte(str)), width, height, ext)
}

func GenFileName(ext string) string {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), Random(10))

	return fmt.Sprintf("%x.%s", md5.Sum([]byte(str)), ext)
}

func MtSubstr(value string, start, end int) string {
	if start > end {
		return ""
	}
	str := []rune(value)
	if length := len(str); end > length {
		end = length
	}

	return string(str[start:end])
}

func BoolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}

func FileSuffix(filename string) string {
	return strings.TrimPrefix(path.Ext(filename), ".")
}

func NewMsgId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func NewUuid() string {
	return uuid.New().String()
}
