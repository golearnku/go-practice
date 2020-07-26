# 账户相关接口

## 账户绑定的手机发送验证码

### Path:

> /v1/account/sms/captcha

### Method:
 > **POST**

## 参数:

| 名称 | 方法 | 类型 | 必须 | 说明 |
| :-: | :-: | :-: | :-: | :-: |
| mobile | POST | string | 是 |  手机号 |

## 返回值：
```json
{
    "status": "OK"
}
```

------ 

## 换绑手机号，获取旧手机号token

### Path:

> /v1/account/change_mobile_token

### Method:
 > **POST**

## 参数:

| 名称 | 方法 | 类型 | 必须 | 说明 |
| :-: | :-: | :-: | :-: | :-: |
| mobile | POST | string | 是 |  旧手机号 |
| captcha | POST | string | 是 |  验证码 |

## 返回值：
```json
{
    "token": "bdc6656883c44c8695dd927658a09188"
}
```

--------

## 换绑手机号

### Path:

> /v1/account/change_mobile

### Method:
 > **POST**

## 参数:

| 名称 | 方法 | 类型 | 必须 | 说明 |
| :-: | :-: | :-: | :-: | :-: |
| mobile | POST | string | 是 |  新手机号 |
| captcha | POST | string | 是 |  验证码 |
| token | POST | string | 是 |  换绑token |

## 返回值：
```json
{
    "status": "OK"
}
```