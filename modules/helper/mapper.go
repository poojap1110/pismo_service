package helper

import (
	"fmt"
	"strconv"
	"strings"
)

// MapToIntSlice ...
func MapToIntSlice(uS []string) (sI []int) {
	for _, item := range uS {
		aI, _ := strconv.Atoi(item)
		sI = append(sI, aI)
	}

	return
}

// ArrayToString ...
func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
