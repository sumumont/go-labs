## 命令使用说明

### 登录获取token

```shell
aiarts login -u username -g usergroup
请输入密码:
```

### 默认的环境变量
```shell
#首次初始化默认环境变量
aiarts init env 
    dataset=xxx #默认使用的数据集
    image=xxx #默认使用的镜像
    project=xxx #默认使用此项目
    quota-cpu=xxx #默认的资源规格 cpu核数
    quota-mem=xxx #默认的资源规格 内存
    quota-gpu=xxx #默认的资源规格 显卡
    quota-num=xxx #默认的资源规格 显卡数量
    quota-node=xxx #默认启动的节点数
#展示全部环境变量
aiarts list env
#修改单个环境变量
aiarts set env image=xxx
```



### 训练

```shell
# 创建训练任务
aiarts submit --image="harbor.apulic.cn/ubuntu1.20:v1" --dataset="xxx数据集"  -p=项目名 "code1||remote:code1" -c "command....." 
返回远端代码路径
remote:xxxxx
如果代码路径是remote: 则不会上传代码，直接运行远端的代码

# 训练任务列表
aiarts list run -p=项目1 -n=1
-n 页码
-p 项目名称

# 获取任务详情
aiarts describe run  -p=项目1 "train-xxxxxxxx"

# 停止或删除任务
aiarts delete run -p=项目1 "train-xxxxxxxx"
#停止任务
aiarts stop run -p=项目1 "train-xxxxxxxx"
# 导出任务结果到本地
aiarts export run -p=项目1 "train-xxxxxxxx" "./outputs/train-xxxxx"
# 查看任务日志
aiarts logs run -p=项目1 "train-xxxxxxxx"

```

### 评估

```shell
# 创建评估任务
aiarts submit --image="harbor.apulic.cn/ubuntu1.20:v1" --dataset="xxx数据集"  -p=项目名 --train='train-xxxxxxxx' "code1||remote:code1" -c "command....." 
返回远端代码路径
remote:xxxxx
如果代码路径是remote: 则不会上传代码，直接运行远端的代码

# 任务列表
aiarts list run -p=项目1 -n=1 --train='train-xxxxxxxx'

-n 页码
-p 项目名称

# 获取任务详情
aiarts describe run -p=项目1 "eval-xxxxxxxx"

# 停止或删除任务
aiarts delete run -p=项目1 "eval-xxxxxxxx"
#停止任务
aiarts stop run -p=项目1 "eval-xxxxxxxx"
# 导出任务结果到本地
aiarts export run -p=项目1 "train-xxxxxxxx"
# 查看任务日志
aiarts logs run -f -p=项目1 "eval-xxxxxxxx"
```