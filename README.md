# orderid
orderid 生成


#### 使用发放

##### 在入口文件初始化
```azure
orderid.Init(2, 1680995200) // 初始化机器ID和起始时间戳
```
#### 然后在自己业务中使用
```azure
##### 生成订单id
```azure
orderid.GenerateOrderId()
```