<p align='center'>
  <pre style="float:left;">
 .-') _                                               (`-.                ('-.     (`\ .-') /`
(  OO) )                                            _(OO  )_            _(  OO)     `.( OO ),'
/     '._   .-'),-----.  .-'),-----.  ,--.      ,--(_/   ,. \  ,-.-')  (,------. ,--./  .--.  
|'--...__) ( OO'  .-.  '( OO'  .-.  ' |  |.-')  \   \   /(__/  |  |OO)  |  .---' |      |  |  
'--.  .--' /   |  | |  |/   |  | |  | |  | OO )  \   \ /   /   |  |  \  |  |     |  |   |  |, 
   |  |    \_) |  |\|  |\_) |  |\|  | |  |`-' |   \   '   /,   |  |(_/ (|  '--.  |  |.'.|  |_)
   |  |      \ |  | |  |  \ |  | |  |(|  '---.'    \     /__) ,|  |_.'  |  .--'  |         |  
   |  |       `'  '-'  '   `'  '-'  ' |      |      \   /    (_|  |     |  `---. |   ,'.   |  
   `--'         `-----'      `-----'  `------'       `-'       `--'     `------' '--'   '--'  
  </pre>
</p>

<p align='center'>
方便地<sup><em>ToolView</em></sup> 开发组件可视化工具
<br> 
</p>

<br>

## 背景


封装各个组件的操作

为可视化工具提供接口


## 特性

- 🗂 etcd
- 📦 ...

## 使用手册
### 安装
<br>

*克隆到本地*

```bash
git clone https://github.com/wwqdrh/toolview.git

// GOOS指定目标操作系统 GOARCH指定cpu指令集
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=web .
```
<br>

*直接安装*

```bash
go install -tags=web github.com/wwqdrh/toolview/@latest
```

如果只需要api，不需要界面可以使用(release中同样也可以选择版本)

```bash
go install -tags=api github.com/wwqdrh/toolview/@latest
```

*release*

进入[release](https://github.com/wwqdrh/toolview/releases)界面直接下载

### 使用
<br>

```bash
toolview -port 8080
```
<br>

访问`localhost:8080`
<br>

