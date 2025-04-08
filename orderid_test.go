/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  orderid_test
 * @Version: 1.0.0
 * @Date: 2025/04/08 14:09
 * @Update liwei 2025/4/8 14:09
 */

package orderid

import (
	"fmt"
	"testing"
)

func TestGetOrderID(t *testing.T) {
	Init(2, 1680995200) // 初始化机器ID和起始时间戳
	for i := 0; i < 10000000; i++ {
		go func() {
			order := GenerateOrderID("OD")
			fmt.Println(order)
		}()
	}
}
