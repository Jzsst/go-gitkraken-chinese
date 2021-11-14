# gitkraken-chinese
GitKraken的中文汉化补丁 - by K-Skye

## compare.html使用
我总结了一下步骤：

1、下载vscode

2、安装 live server 插件

3、创建一个文件夹放入compare.html、当前需要翻译的string.json、之前翻译过的string.json(重命名为zh.json) 如下
![image](https://user-images.githubusercontent.com/44743391/137101937-91f24d35-0047-447a-8a5e-50fb09a9b599.png)

4、用vscode打开所创建的文件，然后找到compare.html，再右键点击选择“Open with  live server ” （确定是否启动为5500端口，不是就将你当前浏览器打开的端口，去替换compare.html中的5500）

5、进行第4步后，会进入compare.html，这时候会自动加载string.json，这时候F12 然后选择console。

6、先点击对比，然后再点击生成json，此时就会生成需要添加到zh.json中的数据（该数据为新版本新增的需要翻译的数据）

7、该数据为json格式，需要使用[格式化工具](https://www.bejson.com/)进行格式化，然后将格式化的数据添加到其中
![image](https://user-images.githubusercontent.com/44743391/137103672-f98446d3-feca-4eb0-a9ce-ca0e2c447f3b.png)

**需要注意：** 需要添加的数据是写入到原有的"strings"。

这是我基于8.0.1处理出来的样例（后面没有翻译的就是添加的）
[zh.txt](https://github.com/k-skye/gitkraken-chinese/files/7336732/zh.txt)

## 说明
自从用上了GitKraken就爱上了，卸载了其他相关git的gui，它的界面非常合我的胃口，但是苦于官方没有中文简体，于是便有了汉化的想法.  
  
更新于2021.03.18 新增对比新旧版本区别，自动生成新版本的json文件的工具compare.html（感谢@DreamSaddle）  
更新于2020.08.18 在windows 2.7.0版本 测试通过（感谢@Black-Spree）  
更新于2019.10.01 在MacOS 10.14 GitKraken 6.2.0测试通过  

## 原理

通过修改软件目录下english语言对应的一个json文件内容来完成汉化目的

## 操作步骤

1. 将项目中的 `strings.json` 替换到 GitKraken 语言目录下的 `strings.json`.  
(实际目录可能会不一样,但文件名一定是strings.json)
  
   - Windows: `%程序安装目录%\gitkraken\app-x.x.x\resources\app\src\strings.json` (x.x.x 是你的GitKraken版本)
   - Mac: `/Applications/GitKraken.app/Contents/Resources/app/src/strings.json`
   - Linux: `/usr/share/gitkraken/resources/app.asar.unpacked/src` (感谢@lyydhy 10.31补充 Gitkraken是deepin 通过deb 安装的)
     
2. 重启GitKraken.

## issue

GitKraken旧版本目录不一样，应该是以下目录：
   - Windows: `%程序安装目录%\gitkraken\app-x.x.x\resources\app.asar.unpacked\src\strings.json` (x.x.x 是你的GitKraken版本)
   - Mac: `/Applications/GitKraken.app/Contents/Resources/app.asar.unpacked/src/strings.json`
