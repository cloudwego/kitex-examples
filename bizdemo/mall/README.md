# Mall Demo Using KiteX
## 一个电商系统的组成
### 角色
* 商家：商品发布、编辑，素材管理，库存管理，评价管理，履约发货等
* 用户：订单创建、退单、收货、浏览商列商详、评价
* 运营：系统元数据配置（如类目维护、销售属性维护、SPU维护、资质维护、标品运营），商家判罚
* 审核：商品审核、召回、解封，资质鉴真
### 前台服务划分
* 商家系统前台
* 消费者系统前台
* 运营平台前台
* 审核系统前台
### 中后台服务划分
* 用户服务
    * 登录
    * 注册
    * 用户详情
    * 会员注册
    * 权限增加、删除、校验
* 商品服务
    * product:
        * 2B: 商品创建、发布、编辑、详情、列表、提审...
        * 2C: 搜索、商详、商列、评价...
    * category: 类目开通、类目检索、类目展示、类目校验...
    * brand: 品牌入驻、品牌校验...
    * stock: 更新库存、计算库存（区域库存、现货库存、活动库存...）、锁定库存、秒杀服务
* 订单服务
    * 订单创建、更新、取消、删除、详情、列表...
    * 锁定库存、超时取消订单返还库存
    * 支付
* 交易服务
    * 创建、详情...
* 商家服务
    * 商家入驻
* 治理服务
    * 商品审核(机审、人审）、商家判罚、商品封禁...
* 营销服务
    * 发券、核销
* 履约服务
* 售后服务
## 运行依赖
* MySQL 存储组件
* ElasticSearch 搜索组件
* ETCD 服务注册发现组件
* Kibana ES数据可视化组件
* RocketMQ 消息组件
## 业务模型设计
### 用户服务
* t_user 账号表
```go
type User struct {
gorm.Model
UserName string `json:"user_name"`
Password string `json:"password"`
}
```
* t_user_role 用户角色表
```go
type UserRole struct {
	gorm.Model
	UserName string `json:"user_name"`
	Roles    string `json:"roles"`
}
```
### 商家服务
* t_shop 商家表
```go
type ShopDO struct {
	gorm.Model
	ShopId   int64  `json:"shop_id"`
	ShopName string `json:"shop_name"`
	UserId   uint   `json:"user_id"`
}
```
### 商品服务
* t_brand 品牌表
```go
type BrandDO struct {
	gorm.Model
	ShopId     int64  `json:"shop_id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	BrandStory string `json:"brand_story"`
}
```
## api接口文档
使用swagger进行接口文档管理&接口测试：`http://localhost:8080/swagger/index.html#/`
## 启动服务
* 构建&运行容器
```shell
$ docker-compose up -d
```