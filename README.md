# orderid
orderid 生成


#### 使用发放

```azure
go get github.com/dvliwei/orderid
```

##### 在入口文件初始化
```azure
import(
    "github.com/dvliwei/orderid"
)
orderid.Init(2, 1680995200) // 初始化机器ID和起始时间戳
```
#### 然后在自己业务中使用
```azure
##### 生成订单id
```azure
import(
    "github.com/dvliwei/orderid"
)
orderid.GenerateOrderId()
```


