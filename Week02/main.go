package main

import (
	"Week02/service"
	"database/sql"
	"errors"
	"fmt"
	pkg_errors "github.com/pkg/errors"
)

func main() {
	res, err := service.Query()
	if err != nil {
		if errors.Is(pkg_errors.Cause(err), sql.ErrNoRows) {
			fmt.Printf("打印错误堆栈信息\n")
			fmt.Printf("%+v\n", pkg_errors.Unwrap(err))
			return
		}
	}
	fmt.Println(res)
}
