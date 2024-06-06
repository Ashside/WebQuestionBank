# API 文档

## 概述 

文档中主要介绍了后端提供的 API，以及调用方法和返回值。

## API 目录

### 1. 用户管理 API

#### 1.1 验证登录用户名密码是否合法

提交方式：`POST`

请求路径：`/api/usr/loginCheck`

请求参数：

| 参数名称 | 参数类型 | 描述   |
| :------- | -------- | ------ |
| username | String   | 用户名 |
| password | String   | 密码   |

返回参数：

| 参数名称 | 参数类型 | 描述                                   |
| :------- | -------- | -------------------------------------- |
| success  | Boolean  | 是否合法                               |
| reason   | String   | 不合法原因，若前一项为 true，则为 null |
| type     | String   | 用户类型，分为两种，teacher 和 admin   |

#### 1.2 验证注册用户是否合法

提交方式：`POST`

请求路径：`/api/usr/registerCheck`

请求参数：

| 参数名称 | 参数类型 | 描述                                     |
| :------- | -------- | ---------------------------------------- |
| username | String   | 用户名                                   |
| password | String   | 密码                                     |
| type     | String   | 注册用户类型，分为两种，teacher 和 admin |

返回参数：

| 参数名称 | 参数类型 | 描述                                   |
| :------- | -------- | -------------------------------------- |
| success  | Boolean  | 是否合法                               |
| reason   | String   | 不合法原因，若前一项为 true，则为 null |

### 2. 题库 API

#### 2.1 获取一个教师的所有题目

#### 2.2 增加一个简答题

提交方式：`POST`

请求路径：`/api/questionBank/addQuestion/simpleAnswer`

请求参数：

| 参数名称   | 参数类型 | 描述                          |
| :--------- | -------- | ----------------------------- |
| question   | String   | Markdown 形式的题目           |
| answer     | String   | 答案                          |
| difficulty | Int      | 题目难度值：1-3，对应从易到难 |
| subject    | String   | 科目类型                      |
| username   | String   | 录入题目用户的 username       |

返回参数：

| 参数名称 | 参数类型 | 描述                                   |
| :------- | -------- | -------------------------------------- |
| success  | Boolean  | 是否成功                               |
| reason   | String   | 不合法原因，若前一项为 true，则为 null |

### 2.3 增加一个选择题

提交方式：`POST`

请求路径：`/api/questionBank/addQuestion/multipleChoice`

请求参数：

| 参数名称   | 参数类型 | 描述                          |
| :--------- | -------- | ----------------------------- |
| question   | String   | Markdown 形式的题目           |
| answer     | String   | 答案，选项。可能是一个或多个  |
| option     | Json     | 所有的选项，默认有4个         |
| difficulty | Int      | 题目难度值：1-3，对应从易到难 |
| subject    | String   | 科目类型                      |
| username   | String   | 录入题目用户的 username       |

返回参数：

| 参数名称 | 参数类型 | 描述                                   |
| :------- | -------- | -------------------------------------- |
| success  | Boolean  | 是否成功                               |
| reason   | String   | 不合法原因，若前一项为 true，则为 null |

示例：

请求数据：

```json
{
  'question': '以下是数学家的有', 
  'answer': 'option1, option2, option3', 
  'option': {
    'option1': '高斯',
    'option2': '欧拉',
    'option3': '莱布尼茨', 
    'option4': '张学友'
  }, 
  'difficulty': 2,
  'subject': 'history', 
  'username': 'example@example.com'
}
```

返回数据：

```json
{
  'success': False,
  'reason': '题目重复'
}
```

