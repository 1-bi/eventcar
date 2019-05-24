# eventcar
基于SIDECAR模式事件调用的模型

该模型建立基于eventbus的基制进行分布式任务协调处理，以提供多节点的并发逻辑处理能力。 eventcar，基意思是专门处理的event基制交通小汽车




初始化Schema

进入到模型代码库
cd schema

执行命令
protoc --gofast_out=. eventmsg.proto