# 模型推理适配
## 模型适配步骤
### 创建推理代理
- 创建并启动 pg数据通道和图片通道，推理代理将保存推理数据以及图片分别存到这两个通道，如果已经存在通道则忽略该步骤

```
    创建数据通道的接口:
    https://192.168.2.75/bff-admin/api/v1/adhub/data-connectors/schema-based-connectors
    POST请求:
    Header:
    Authorization:
    content-type: application/json; charset=UTF-8
    end-org:apulis
```
body: Schema.json 内容
注意修改name字段，这是数据通道名称, dataPoolId每个平台不一定相同，在各个平台自行查看

- init-job-tool 适配数据回传代码,并制作成推理代理镜像
  - harbor.apulis.cn:8443/apulis-iqi/infra/init-job-tool:v3.7.2-cxcc
- 创建推理代理，填入上面步骤创建好的数据通道
![img.png](images/infer_agent.png)

### 上传算法模型
- 检查算法包，manifest.yaml里面根节点下，添加两个参数，如果已经有则忽略
```
vendor: apulis.infer
model_version: cxcc.v2
```
![img.png](images/model_upload.png)
### 创建项目
![img.png](images/project_create.png)
### 训练，评估，模型发布
- 数据集均由算法提供
![img.png](images/train.png)
### 部署发布-创建部署包
![img.png](images/deploys.png)
![img.png](images/deploy_create.png)
### 创建推理服务
- 选择 创建好的项目
- 选择 创建的部署包
- 选择 创建好的推理代理

### 推理测试
- 调用推理接口成功
- 检查数据通道查看数据及图片是否正常存储
- 以上检查均没问题，则推理适配成功