# XETH
⭐An open source platform about ethereum data intelligent analysis and management.

## 后端开发指南
### 模型搭建
+ `config`：相关配置文件
  + API返回格式的封装`resultMsg.go`，可自行添加
  + 数据库等变量的配置`config.go`，这里读取环境变量`.env`
  + 返回码的枚举`code.go`，可自行添加
+ `sql`：数据库相关配置
  + `xeth.sql`可以直接运行导入本地数据库`xeth`中
  + `block.png`相关可视化的图，上面附有数据库变量名以及数据类型，可能有不准/有些值没有设置成`NOT NULL`，后续修改，读者可以依赖此图完成`mode`层相关配置的编写
  具体可参考[MySQL和Go数据类型互相转换](https://segmentfault.com/a/1190000011016366)
+ `dbDriver`：数据库相关配置，由于初步只是建立一个数据库，所以配置了全局变量，同时保留了之前的函数方便不同数据库的配置
+ `server`：相关路由配置，跨域已初步配置好了，现在只需要参照已配置好的相关路由组配置其他`model`对应的路由组
+ `model`：直接与数据库相连，为不同结构体，对应数据库的不同表
+ `controller`：控制层，借助在`server.go`中配置的路由，获取前端传来的参数，并调用`service`层相关函数进行实际处理
+ `service`：真正的逻辑处理层，上面有相关查询的不同数据，根据不同的查询需求定义不同的函数
+ `DTO`：由于推荐传参方式为POST传JSON，故定义不同DTO结构体，作为`service`层和`controller`层相关函数的参数，设计目的是便于修改以及便于传参等，需要自行定义
+ `utils`：相关工具层
  + `constants.go`定义相关常量
  + `encrypt.go`定义MD5加密，不知道有没有用
  + `log`个性化配置，非必要，待完善
+ `.env`：相关环境变量，与其相配的读取方法位于`config`层
+ `pkg`：相关外部模块，比如`download.py`，我有做修改，比如`py`脚本下载的相关数据

  P.S. 此为不确定功能，因为压缩包还要解压存入数据库，极其恶心，为图简便建议提前下载好数据，写代码存入数据库（推荐使用带有可视化界面的软件进行输入），同时更新`xeth.sql`文件 
  
  标准化方法：调用下载脚本下载，然后解压，解析csv文件并存入数据库（以上种种函数可以放到`config`层或`utils`层当中，新建go程序）