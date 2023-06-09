

# shop，mini demo


API 网关：处理 HTTP 请求
Auth 服务：提供注册、登录和生成 JWT 令牌
Product 服务：添加商品、扣减库存、查找商品
Order 服务：创建订单


mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| auth_svc           |
| order_svc          |
| product_svc        |
+--------------------+


# 架构图

<img src="doc/images/arch.jpg" width="500" height="400">

```text

proto 文件要保持一致


```