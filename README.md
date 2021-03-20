##  创建 go module 项目流程及步骤
首先介绍创建go module 项目工程

#### 1、创建工程
1、 在github 创建一个仓库地址， 然后将其拉下来
2、 进入对应目录， 执行 : go mod init go mod init github.com/duxinglangzi/go-utils
会提示:  `go: creating new go.mod: module github.com/duxinglangzi/go-utils`  
3、 编写自己的代码文件 ~  并提交代码
4、创建 git tag 标签， 命令:   **` git tag v0.0.1 `**
5、 此时推送git tag 到远程仓库 : git push origin v0.0.1


#### 2、在其他工程引入时可能出现的错误
1、提示错误: github.com/duxinglangzi/go-utils: module github.com/duxinglangzi/go-utils@latest found (v0.0.3), but does not contain package github.com/duxinglangzi/go-utils
此类型错误请尝试: **`go clean -modcache  `**    
然后再重新下载: **`go mod download  `**


#### 3、Go module 基本知识理解
1、go module 版本必须遵循 `X.Y.Z` 规则， 即  Y 和 Z 属于小版本更新。 需注意的是 如果变更 X 则需要修改 go.mod 第一行，在module那行最后加上/v2。module github.com/duxinglangzi/go-utils/v2 .  
请注意:  **`v 所有版本号都是 v 开头。`**
|名称| 解释  |
|--|--|
| MAJOR | 主版本号，如果有大的版本更新，导致 API 和之前版本不兼容。我们遇到的就是这个问题 |
|MINOR|次版本号，当你做了向下兼容的新 feature。|
|PATCH|修订版本号，当你做了向下兼容的修复 bug fix。|

![如图](https://img-blog.csdnimg.cn/20210320115657601.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTAwMjUyOTQ=,size_16,color_FFFFFF,t_70#pic_center)

#### 4、go mod 常用命令 

|  命令 | 解释  |
|--|--|
| download <br> `(常用)` | 使用此命令来下载指定的模块 `go mod download` 附带参数 [-dir] [-json] [modules]。<br> 模块的格式可以根据主模块依赖的形式或者path@version形式指定。如果没有指定参数，此命令会将主模块下的所有依赖下载下来。<br> 默认情况下，下载错误会输出到标准输出，正常情况下没有任何输出。**`-json`** 参数会以JSON的格式打印下载的模块对象，对应的Go对象结构。<br> 如果该文件夹的内容太大或者未能下载最新版，可以通过命令 go clean -cache 清理一下缓存 |
| edit | 从工具或脚本中编辑go.mod文件 |
| graph | 打印模块需求图 |
| init | 在当前目录下初始化新的模块 |
| tidy  <br> `(常用)`| 添加缺失的模块以及移除无用的模块 |
| verify | 验证依赖项是否达到预期的目的 |
| why | 解释为什么需要包或模块 |
|vendor <br> `(常用)` | 常规用法: `go mod vendor ` , 如果需要指定版本 `go mod vendor [-v] ` 命令会将build阶段需要的所有依赖包放到主模块所在的vendor目录中, 并且测试所有主模块的包。<br>同理 go mod vendor -v 会将添加到vendor中的模块打印到标准输出。 |

