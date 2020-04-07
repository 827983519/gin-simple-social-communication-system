sbs project framwork
```
/

├── cmd                    // 命令行入口，程序入口
│    └─ http_server             // 按类型分目录和文件
│          └── main.go
│   
├── apps                     //业务代码，按业务分目录      
│    ├─ common             // 基础服务
│    │    ├── controllers               // view层
│    │    ├── rooters            // 路由配置
│    │    ├── model            // 数据模型
│    │    ├── repository            // repository 数据库操作
│    │    └── services       // 业务逻辑代码
│    └─ example               //业务
│          ├── controllers               // view层
│          ├── rooters            // 路由配置
│          ├── model            // 数据模型
│          ├── repository            // repository 数据库操作
│          └── services       // 业务逻辑代码
│
├── agent
│    ├── db             // db agent
│    └── cache           // cache agent
├── middleware           //中间件 
├── deploy 
│	  └─ sql            //mysql建表语句
├── config              //配置文件
├── types             // 数据结构定义，不依赖内部package  
├── libs              //util、lib等独立package，不依赖项目内部package
├── docs             // 对外文档
├── readme.md
└── go.mod
```