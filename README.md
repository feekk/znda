# booking_car
处于好奇心及自我无聊时间打发的网约车项目

## 系统描述 
整个系统主要为服务端的相关功能研发及设计，不包含客户端、小程序、H5等的页面设计（后端狗不太懂，学习是不可能的，一辈子都不可能回去学的，这里人又多，说话又好听，超喜欢和他们在一起的）

## 模块描述

base-service-auth 鉴权服务，用于生成token及验证

base-service-connection 长连接服务

base-service-discover 服务发现

base-service-gateway 网关服务

base-service-lbs 位置服务

base-service-push 触达服务。短信、端内外推送等

bussiness-client-driver 司机端API及逻辑服务

bussiness-client-mis 内部管理服务

bussiness-client-passenger 乘客端API及逻辑服务

bussiness-service-bill 账单服务

bussiness-service-order 订单服务

bussiness-service-pay 支付服务

bussiness-service-user 用户司机车辆管理服务

common 公共资源包
