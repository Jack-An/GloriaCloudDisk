## 功能
1. 用户登录
2. 用户上传文件 暂时支持文本、视频、音频、图片
3. 用户预览文件（音视频图片转码)
4. 用户下载（可以多选打包）


### Timeline
    -- 项目开始   2022.5.8


### Before Dev

```bash
    CREATE DATABASE net_disk_dev CHARACTER SET UTF8;  # for dev 
    CREATE DATABASE net_disk_test CHARACTER SET UTF8;  # for unit test (not yet)

```


## 后端模块（暂时想到的） 

### user rpc
    1. 注册
    2. 登录
    3. 修改个人信息

### storage rpc
    1. 文件的上传

### transcode rpc
    2. 转码服务

### zip rpc
    3. 文件打包


### 数据库模型设计
user:  用户基本信息表

identity:  用户认证信息表

upload:  上传原始文件信息表

----------------------------------

document:

image:

video:

audio:

user_log:
