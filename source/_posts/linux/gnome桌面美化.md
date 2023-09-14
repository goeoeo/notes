---
title: gnome桌面美化
categories:
- linux

tags:
- 桌面美化
- gnome
---


## 安装以下软件
```shell
sudo apt install gnome-tweak-tool gnome-shell-extensions chrome-gnome-shell
```


## gnome 插件
地址 https://extensions.gnome.org

### User Themes ：改变Gnome主题必要插件！
https://extensions.gnome.org/extension/19/user-themes/

### Dock from Dash ：类似与mac底部dash的插件
https://extensions.gnome.org/extension/4703/dock-from-dash/


### bing-wallpaper
桌面壁纸切换  
注意桌面切换在ubuntu20.04 上面需要通过tweak 将工作区关闭调，否则在有两个显示器的情况，壁纸切换会使得主显示器dock出现重影

### TopIcons Plus
右上角展示微信图标
https://extensions.gnome.org/extension/1031/topicons/


## 锁屏背景
```shell
wget github.com/thiggy01/change-gdm-background/raw/master/change-gdm-background
chmod +x change-gdm-background

sudo ./change-gdm-background /home/yu/图片/BingWallpaper/20220404-Godafoss_ZH-CN9460037606_UHD.jpg
```