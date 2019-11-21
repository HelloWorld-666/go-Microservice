# go-Microservice
go微服务

Go Micro是可插拔的微服务开发框架。  
 1.consul : 服务注册及发现(通过http/DNS)，健康检查，key/value动态存储配置文件，提供web管理界面（除了基础和路由服务，还要额外启动consul服务器）  
 2.grpc   : router <-> base通过http/2协议通信，通过protobuf和json进行编解码  
 3.iris   : router  
 4.xorm   : base（需要连接Mysql数据库）  
 5.seelog : logger  
 6.mysql  ：数据库（部署在腾讯云上，创建student表的.sql文件在当前目录下）  
 7.包管理工具：govendor  

 
 **启动服务步骤：**  
 1.分别运行router和base下的run.bat脚本（运行路由、基础服务）  
 &emsp因为github限制不能上传超过100M的文件，所以windows上用的consul.exe无法上传，继而压缩，使用时需要解压到压缩文件所在目录下  
 2.运行common/consul目录下的run.bat脚本	（运行consul服务）  
 3.Postman请求"http://127.0.0.1:8080/v1/student/set_student_info"  
 
 	Body参数如下：  
	{  
		"name": "pmy",  
		"age": "12",  
		"sex": "male",  
		"mathScore": "199"  
	}  
	Headers参数中 添加Content-Type为application/json  
	
	
	http://127.0.0.1:8500/ui/dc1/services  
	consul服务可以通过网页查看服务状态，包括健康检查...  
	
	
	另外，test\common\consul 目录下的run.bat脚本如下：  
	::consul agent -dev -config-dir . -data-dir=./tmp  
	consul agent -server -ui -bootstrap -data-dir=./tmp -node=consul-1 -client=0.0.0.0 -bind=127.0.0.1 -config-dir .  
	
	参数说明：  
	-dev 单节点consul模式，不能数据持久化，因此不能用于生产环境中  
	-server server模式，可以数据持久化  
	
	-data-dir：保存consul中key-value的数据，实现持久化.  
	-config-dir： 要加载的配置文件的目录，Consul将加载后缀为“.json”的所有文件.  
	-node：集群中此节点的名称  
	-client：Consul将绑定客户端接口的地址  
	-bind：应为内部集群通信绑定的地址  
	更多见：https://www.cnblogs.com/sunsky303/p/9209024.html  
	
	注：连续两个冒号::是命令行中的注释符  
  
  
  
**附录：**  
consul日志相关配置：    
创建consul web ui上面的配置文件（config目录下的log_configuration配置）：config/log_configuration   

```
	{  
	  "log": {  
		  "isDevelopmentEnv":"1",  
		  "format":"[%Date(2006-01-02 15:04:05.000)] [%LEVEL] [%RelFile:%Line] %Msg%n",  
		  "infoFileName":"./logs/info/info",  
		  "debugFileName":"./logs/debug/debug",  
		  "errorFileName":"./logs/error/error",  
		  "criticalFileName":"./logs/critical/critical",  
		  "saveOldInfoLogPath":"./logs/info/",  
		  "saveOldDebugLogPath":"./logs/debug/",  
		  "saveOldErrorLogPath":"./logs/error/",  
		  "saveOldCriticalLogPath":"./logs/critical/",  
		  "timePattern":"20060102_15.log",  
		  "maxRolls":720  
		}  
	}  
```


报错：有时由于重新启动时创建的服务节点数过多，导致常常访问接口报internal server错误（负载均衡，将请求转向其他错误服务地址）：  
如果consul中出现节点数量过多，那就是-data-dir=./tmp出了问题，清空./tmp下的所有文件，重新启动consul并注册base和router服务.
 
