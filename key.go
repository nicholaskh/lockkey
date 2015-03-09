package lockkey

import (
	"fmt"
)

func User(uid int64) string {
	return fmt.Sprintf("lk:u:%d", uid)
}

func Attackee(k, x, y int16) string {
	return fmt.Sprintf("lk:m:%d,%d,%d", k, x, y)
}
