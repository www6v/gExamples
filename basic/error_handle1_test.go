package basic

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

// https://blog.csdn.net/wanglei19891210/article/details/128092331
// 打印堆栈信息
func Test_Main2(t *testing.T) {
	err := first()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("Data error, %+v\n", err)
		// fmt.Printf("%+v\n", err)
		return
	}
}

func second() error {
	return errors.Wrap(sql.ErrNoRows, "second failed")

}
func first() error {
	return errors.WithMessage(second(), "first failed")
}
