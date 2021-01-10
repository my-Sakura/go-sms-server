# sms-server 

<a href = https://www.github.com/my-sakura/sms-server><img src = "https://img.shields.io/badge/readme%20style-standard-green"></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/my-sakura/sms-server)](https://goreportcard.com/report/github.com/my-sakura/sms-server)

## Table of Contents

- [smsProvider](https://github.com/my-Sakura/sms-server#smsprovider)
  - [tianyan](https://github.com/my-sakura/sms-server#tianyan)
  - [dingxin](https://github.com/my-sakura/sms-server#dingxin)
- [Install](https://github.com/my-Sakura/sms-server#install)
- [Usage](https://github.com/my-Sakura/sms-server#usage)
  - [Generator](https://github.com/my-Sakura/sms-server#generator)
- [API](https://github.com/my-Sakura/sms-server#api)
- [Examples](https://github.com/my-Sakura/sms-server#example)
- [License](https://github.com/my-Sakura/sms-server#license)

## smsProvider
### tianYan

  https://market.aliyun.com/products/57000002/cmapi00039249.html
### dingXin

  https://market.aliyun.com/products/56928004/cmapi023305.html
  
## Install

## Usage

1.select [smsProvider](https://github.com/my-Sakura/sms-server#smsprovider)
   选择短信 API 服务商，目前支持 [tianyan](https://market.aliyun.com/products/57000002/cmapi00039249.html) 和 [dingxin](https://market.aliyun.com/products/56928004/cmapi023305.html)，记录供应商名称(全部小写)，之后会用到
   
2.进入云市场购买产品，购买完之后申请模板，记录 templateCode 和 appCode, eg:
  ![](https://github.com/my-Sakura/sms-server/blob/main/pictures/first.png)
   
  ![](https://github.com/my-Sakura/sms-server/blob/main/pictures/second.png)
   
  联系客服申请模板，获得 templateCode
  ![](https://github.com/my-Sakura/sms-server/blob/main/pictures/third.png)
   
  ![](https://github.com/my-Sakura/sms-server/blob/main/pictures/fourth.png)
   
  记录 appCode
  ![](https://github.com/my-Sakura/sms-server/blob/main/pictures/fifth.png)
  
3.修改 [config.yaml](https://github.com/my-Sakura/sms-server/blob/main/config/config.yaml) 文件
  将记录下的 smsProvider、templateCode、appCode 写入 config.yaml 文件中
  
4.调用 API

  ```
  client := api.NewClient()
  err := client.Send(phone_number, verification_code_length)    
  if err != nil {
     log.Println(err)
     }
  ```

## API

To see how the specification has been applied, see the [API](https://github.com/my-Sakura/sms-server/tree/main/api) directory.

## Example

To see how the specification has been applied, see the [Examples](https://github.com/my-Sakura/sms-server/tree/main/examples) directory.

  ```
  package main
   
  import (
      "github.com/my-Sakura/sms-server/api"
  )
   
  func main() {
  client := api.NewClient()
  err := client.Send(phone_number, 6)    
  if err != nil {
     log.Println(err)
     }
  }
  ```

## License

[MIT](https://github.com/my-Sakura/sms-server/blob/main/LICENSE) © my-Sakura :heart: :star: :sparkles: :dizzy:
