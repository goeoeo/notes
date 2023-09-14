---
title: ubuntu安装AndroidSDK和NDK
categories: 
- android
tags:
- android
---


# 安装java 
参考:  
https://medium.com/@barcelos.ds/install-openjdk-11-lts-in-the-ubuntu-20-04-lts-2c06f17c990

## 下载
https://adoptopenjdk.net
```
sudo mkdir ~/DevTools
sudo mkdir ~/DevTools/JDK
sudo tar -zxvf ~/下载/OpenJDK11U-jdk***tar.gz -C ~/DevTools/JDK/
```

<!--more-->

## 设置JAVA_HOME环境变量
```
sudo  vim ~/.profile

DEV_TOOLS="/home/$USER/DevTools"
JAVA_HOME="$DEV_TOOLS/JDK/jdk-11.0.7+10"
export JAVA_HOME
PATH="$JAVA_HOME/bin:$PATH"

source ~/.profile
```
## 测试
```
#JRE
java --version

#JDK
javac --version
```


# 安装Android SDK
参考:  
https://medium.com/@barcelos.ds/install-android-sdk-in-the-ubuntu-20-04-lts-without-android-studio-1b629924d6c5

## 下载命令行工具
网址：https://developer.android.com/studio
![](ubuntu安装AndroidSDK和NDK/command-tools.png)

## 目录
```
sudo mkdir ~/DevTools
sudo mkdir -P ~/DevTools/Android/cmdline-tools/latest
```

## 解压放至于  ~/DevTools/Android/cmdline-tools/latest
```

```

## 设置ANDROID_HOME环境变量
```
sudo  vim ~/.profile

DEV_TOOLS="/home/$USER/DevTools"
JAVA_HOME="$DEV_TOOLS/JDK/jdk-11.0.7+10"
ANDROID_HOME="$DEV_TOOLS/Android"
export JAVA_HOME
export ANDROID_HOME
PATH="$JAVA_HOME/bin:$ANDROID_HOME/cmdline-tools/latest/bin:$ANDROID_HOME/platform-tools:$PATH"

source ~/.profile

```

## Install Platform for Android 29
```
sdkmanager "platform-tools" "platforms;android-29"
```

## Install Build Tools for Android 29
```
sdkmanager "build-tools" "build-tools;29.0.3"
```

## Accept Android Licenses
```
sdkmanager --licenses
```

## Update Android Packages when necessary:
```
sdkmanager --update
```

## 测试
```
sdkmanager --version
sdkmanager --list
```

# 安装NDK

## 官网下载
https://developer.android.com/ndk/downloads/index.html

## 解压到 ~/DevTools/Android

## 设置ANDROID_NDK
```
sudo  vim ~/.profile

DEV_TOOLS="/home/$USER/DevTools"
JAVA_HOME="$DEV_TOOLS/JDK/jdk-11.0.9.1+1"
ANDROID_HOME="$DEV_TOOLS/Android"

export JAVA_HOME
export ANDROID_HOME
export ANDROID_NDK=$ANDROID_HOME/android-ndk-r21b

export GOPATH=/home/yu/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin:$JAVA_HOME/bin:$ANDROID_HOME/cmdline-tools/latest/bin:$ANDROID_HOME/platform-tools:$ANDROID_NDK
export GOROOT=/usr/local/go
export GOPROXY=https://goproxy.cn

source ~/.profile
```

## 测试
```
ndk-build -v
```