＃ *** 项目

＃＃ 介绍

- 使用[Kitex](https://github.com/cloudwego/kitex/)框架
- 生成单元测试的基本代码。
- 提供基本的配置功能
- 提供最基本的MVC代码层次结构。

## 目录结构

| 目录          | 介绍          |
|-------------|-------------|
| conf        | 配置文件        |
| main.go     | 启动文件        |
| handler.go  | 用于请求处理返回响应。 |
| kitex_gen   | kitex 生成的代码 |
| biz/service | 实际的业务逻辑。    |
| biz/dal     | 存储层操作逻辑     |

## 如何运行

````外壳
sh build.sh
sh output/bootstrap.sh
````