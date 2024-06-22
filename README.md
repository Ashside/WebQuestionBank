# 2024年春季实践项目

## 第二轮开发进度

### 前端

- [x] 老师录入题目界面
- [x] 相似卷查找界面
- [ ] 相似卷添加界面功能
- [x] 题目检索功能
- [ ] 学生作答功能
- [x] 老师批改功能

#### maybe:
- [ ] 老师查看学生总体数据功能
- [ ] 题目修改功能

## Requirements

### 后端

后端使用 Go 语言编写，版本为 1.20，适用于 Windows / amd64 架构。Web 框架采用 gin-gonic / gin v1.7.4。你需要安装 Go 和相关的库来保证程序的正常运行。

### 前端

前端使用 Vue3 进行编写，你需要安装 Node.js 和 Vue 来保证程序的正常运行。

### 数据库

数据库使用 MySQL 数据库。

### AI 模块

AI 模块需要安装 Python 3.7 或更高版本，并安装 numpy, gensim 以及 flask 库以运行。

## 实现功能

## 如何部署和配置

### 前端部署

#### `npm` 部署

前端你可以直接使用 `npm` 进行部署。要确保你已经安装了 Node.js，如果没有安装，请前往 https://nodejs.org/ 进行安装。安装完成后，使用：

```
npm install
```

`npm` 会自动完成依赖配置，之后使用命令：

```
npm run serve
```

运行前端服务，并在浏览器中查看。

#### Docker 部署

使用 Docker 运行前端界面，切换到 `\frontend` 文件夹下，执行如下指令：

```
docker build -t frontend .
docker run -p 8080:3000 frontend
```

之后在本机的 `8080` 端口即可完成访问。

#### 运行 IP 配置

对于后端不同的服务器 IP 地址，可以在 `src/frontend/.env` 中进行配置。后端 API 访问的 IP 地址需修改一下字段：

```
VUE_APP_API_URL = http://localhost:8081
```

重新启动并运行。

### AI 模块说明

AI 生成词向量模块需要安装版本大于等于 3.7 的 Python。安装完成后运行以下指令安装依赖：

```
pip install Flask gensim numpy jieba
```

#### AI 提取题目关键词

目前使用[textrank](https://github.com/abner-wong/textrank)来进行题目的关键词提取。关键词提取的实现在 `src/AI/findOutKeyword.py` 中实现，以 Flask 框架运行在服务器的 `5050` 端口。要获取题目关键词，你可以用以下方式请求 `5050` 端口的 `/extract` 路径，请求的 json 文件格式示例如下：

```json
{
    "text": "欧亚经济委员会执委会一体化与宏观经济委员格拉济耶夫日前接受新华社记者采访时高度评价中国抗击新冠疫情工作，并表示期待欧亚经济联盟与中国加强抗疫合作，共同推动地区发展。格拉济耶夫说，中国依靠治理体系与全国人民协同努力，在抗疫工作上取得极大成效。中国采取的措施符合全球利益。格拉济耶夫认为，中国经济将会快速恢复，欧亚经济联盟许多企业与中国市场联系紧密，应与中国加强合作，采取协调措施降低此次疫情带来的消极影响。格拉济耶夫建议，面对疫情，欧亚经济联盟与中国扩大信息技术应用，推进商品清关程序自动化，更广泛地利用相关机制，为对外经济活动参与者建立绿色通道。谈及双方在医学卫生领域的合作时，格拉济耶夫说：“我们应从当前考验中汲取经验，在生物安全领域制定共同规划并联合开展生物工程研究。”格拉济耶夫还表示，俄罗斯与其他欧亚经济联盟国家金融市场更易受国际投机行为影响。欧亚经济联盟应借鉴中国的人民币国际化经验，加强与中国银行体系和金融市场对接。欧亚经济联盟成立于2015年，成员国包括俄罗斯、哈萨克斯坦、白俄罗斯、吉尔吉斯斯坦和亚美尼亚。欧亚经济委员会执委会是欧亚经济联盟最高权力机构。"
}
```

请求完成后会返回对应的关键词和权重，示例如下：

```json
{
    "keywords": [
        {"keyword": "中国", "weight": 1},
        {"keyword": "疫情", "weight": 2},
      	...
    ]
}
```

#### AI 生成词向量用于题型相似度匹配（测试）

使用 AI 进行生成词向量在 `src/AI/findOutKeyword.py` 中实现，并通过 Flask 框架运行在 `8083` 端口上。示例请求方式如下：

```json
{
    "word": "中国"
}
```
你将获得对应词的词向量，示例如下：

```json
{
    "word": "中国",
    "vector": [0.123, -0.234, 0.456, ...]  
}
```
我们使用腾讯开源词向量进行分析，你可以在[这里](https://ai.tencent.com/ailab/nlp/en/embedding.html)获取。你也可以替换成其他的词向量源，在以下位置完成替换：

```python
# 替换成你需要的源
file_path = '/home/user/tencent-embedding/tencent-ailab-embedding-zh-d100-v0.2.0-s.txt'
```

## OpenAPI

我们使用 Apifox 组织 API，你可以在[此网站](https://apifox.com/apidoc/shared-6bd451e3-8d10-40a4-bb52-5ce49f6262de)中查到目前开放的API接口。

## Acknowledgement

1. 感谢 OpenAI [ChatGPT](https://chatgpt.com/#) 的大力支持。

## 策划案

### 项目名称

题库管理与在线考试

### 基本需求
- 题库管理：能够录入题目、答案、难易度值等
- 能够处理多种题型：选择题（单选、多选）、判断题、问答题
- 组卷：选择题目构成试卷、重复判断、题目构成分析等
- 权限管理：老师之能维护各自题目，管理员有权组卷等
- 附加功能：在线答题、自动评判（选择题、判断题）

### 前期规划

- 前端
  - 录入题目（编辑器）
    - markdown文本
    - 图片文件（包括markdown中引用的）
    - 设定难易度（数值）
    - 设定题目类型（选择题、判断题、问答题）
    - 设定答案（文本及图片，与题目类型对应）
    - 上传题目
    
  - 组卷
    - 从后端获取题目（按难易度、类型等筛选）
    - 展示筛选后的题目
    - 添加可变数量的题目（按链表处理？）
    - 生成试卷（pdf格式）
    - 题目查重
    - 试卷分析（题目类型、难易度分布，或者其他tag，需要定义）
    - 上传试卷

  - 判卷
    - 自动判卷（选择题、判断题）
  - 在线答题
    - 从后端获取试卷
    - 自动判卷（选择题、判断题）
- 后端
  - 题库管理
    - 题目增删改查
      - 统一管理上传的图片文件和markdown文本
      - 需要做到能够定位图片文件
    - 题目类型增删改查
    - 题目难易度增删改查
    - （题目标签增删改查）
  - 组卷
    - 题目筛选
    - 试卷生成
    - 试卷上传
  - 判卷
    - 自动判卷
  - 在线答题
    - 试卷获取
    - 自动判卷
