# botsend

botsend 是一个命令行的企业微信机器人消息发送工具

## 安装

```bash
go get -v github.com/hubenchang0515/botsend/cmd/botsend
```

## 使用说明

## 企业微信添加群机器人
![6260133c7e895e94b027acf5b57b7cad_720](https://github.com/user-attachments/assets/afe4ed99-0a10-459a-b8fb-517353a60332)
![96a039f61c81ae868dc66b484691614d](https://github.com/user-attachments/assets/d03c958c-6532-4b4f-a3a9-0f1ec3d0a217)


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
