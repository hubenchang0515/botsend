# botsend

botsend 是一个命令行的企业微信机器人消息发送工具

## 安装

```bash
go get -v github.com/hubenchang0515/botsend/cmd/botsend
```

## 使用说明

## 企业微信添加群机器人
![367578229-afe4ed99-0a10-459a-b8fb-517353a60332](https://github.com/user-attachments/assets/9031ce1d-5b80-4f9c-bb92-c2f86a33086e)
![367578237-d03c958c-6532-4b4f-a3a9-0f1ec3d0a217](https://github.com/user-attachments/assets/9f27b62c-6631-4bdd-a1be-8cb102e00a18)
![367578455-3a538b90-dd0b-49a9-9c9e-4ce3bad7cce8](https://github.com/user-attachments/assets/54379fc1-8dfd-4279-9318-1a98af9ffe9b)
![367578457-a53cd3a3-7576-4b9a-9fab-7d8bcca91e23](https://github.com/user-attachments/assets/7b5fb28a-b1e4-4d26-9ed6-2752eec6af81)




### 设置KEY
* `botsend -key xxxxxxxxxxxxx`

### 普通消息
* 发送普通消息 `botsend 消息内容`
* 发送普通消息并@群成员 `botsend -at=成员的手机号(逗号分隔) 消息内容`
* 发送普通消息并@全体成员 `botsend -at=@all 消息内容`

### 图片消息
* 发送图片 `botsend -type=picture 图片文件`

### Markdown消息
* 通过文件发送 `botsend -type=markdown xxx.md`
* 通过字符串发送 `botsend -type=markdown -markdown=xxxxxxx`

### 链接消息
* 发送链接 `botsend -type=url -url=链接 -title=标题 -description=说明(可选) -cover=封面图的URL(可选)`

### 文件消息
* 发送文件 `botsend -type=file 文件`
