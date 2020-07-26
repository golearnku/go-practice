# 登录相关接口

## 登录账户

### Path:

> /v1/login

### Method:
 > **POST**

## 参数:

| 名称 | 方法 | 类型 | 必须 | 说明 |
| :-: | :-: | :-: | :-: | :-: |
| device_id | POST | string | 是 |  客户端设备号 |
| mobile | POST | string | 是 |  手机号 |
| captcha | POST | string | 是 |  验证码 |

## 返回值：
```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTU3NjM4MTcsImlhdCI6MTU5NTc1NjYxNywibmJmIjoxNTk1NzU2NjE3LCJzdWIiOiI0ZWJkNDNhZjYxOWQ0NTY0YjUzMjhhNTI3MzM0MTNiZiJ9.044s5P9yxWYHoQeKvnSc7kIq4kgYI1XGzeRpC21uxJ9q-DNMlvHfABSzXBLQytt-3v2k7OisnYCFvKGRQaRv1w",
    "token_type": "Bearer",
    "expires_at": 1595763817
}
```

------ 

## 退出登录

### Path:

> /v1/logout

### Method:
 > **DELETE**

## 返回值：
```json
{
    "status": "OK"
}
```

------ 

## 刷新令牌

### Path:

> /v1/current/refresh-token

### Method:
 > **POST**

## 返回值：
```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTU3NjM4MTcsImlhdCI6MTU5NTc1NjYxNywibmJmIjoxNTk1NzU2NjE3LCJzdWIiOiI0ZWJkNDNhZjYxOWQ0NTY0YjUzMjhhNTI3MzM0MTNiZiJ9.044s5P9yxWYHoQeKvnSc7kIq4kgYI1XGzeRpC21uxJ9q-DNMlvHfABSzXBLQytt-3v2k7OisnYCFvKGRQaRv1w",
    "token_type": "Bearer",
    "expires_at": 1595763817
}
```