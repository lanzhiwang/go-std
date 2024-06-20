# Go 1.16 中关于 go get 和 go install 你需要注意的地方

Go (golang) 已于 18 日发布了 1.16 beta1 版本，至此其主体功能已经基本确定。我看大多数人都在关注 Go 在苹果(Apple) M1 上的支持，甚至 Go 官方博客中也有一篇专门的说明 Go on ARM and Beyond，来介绍 Go 在此方面的支持。

我就不凑热闹了，我来聊聊 Go 1.16 中关于 go get 和 go install 你需要注意的地方。

目前 Docker 官方镜像尚未发布，我是本地构建了个镜像来使用。

```bash
(MoeLove) ➜  go version
go version go1.16beta1 linux/amd64
```

## 概览

Go 1.16 中包含着大量的 Modules 相关的更新，详细内容可直接查看其 Release Note。整体而言，包含以下要点：

- GO111MODULE 默认为 on ，如果要恢复到之前的行为，则需要将 GO111MODULE 设置为 auto ，这样差不多意味着 GOPATH 模式要逐步淡出人们的视野了；

- go install 命令可以接受一个版本后缀了，（例如，go install sigs.k8s.io/kind@v0.9.0），并且它是在模块感知的模式下运行，可忽略当前目录或上层目录的 go.mod 文件。这对于在不影响主模块依赖的情况下，安装二进制很方便；

- 在将来，go install 被设计为“用于构建和安装二进制文件”， go get 则被设计为 “用于编辑 go.mod 变更依赖”，并且使用时，应该与 -d 参数共用，在将来版本中 -d 可能会默认启用；

- go build 和 go test 默认情况下不再修改 go.mod 和 go.sum。可通过 go mod tidy ，go get 或者手动完成；

总结而言，关于 go install 和 go get 必须要注意的是：

- 基本上 `go install <package>@<version>` 是用于命令的全局安装：

  例如：go install sigs.k8s.io/kind@v0.9.0;

- go get 安装二进制的功能，后续版本将会删除；

- go get 主要被设计为修改 go.mod 追加依赖之类的，但还存在类似 go mod tidy 之类的命令，所以使用频率可能不会很高；

## Go 1.16 中已解决的工具安装问题

到目前为止，Go 一直使用 go get 命令，将我们需要的工具安装到 $GOPATH/bin 目录下，但这种方式存在一个很严重的问题。go get 由于具备更改 go.mod 文件的能力，因此我们 必须要避免执行 go get 命令时，让它接触到我们的 go.mod 文件 ，否则它会将我们安装的工具作为一个依赖。

目前的解决方案通常是：

```bash
(MoeLove) ➜  cd $(mktemp -d); GO111MODULE=on go get sigs.k8s.io/kind@v0.9.0
```

自 1.16 开始，我们可以直接使用下面的方式：

```bash
(MoeLove) ➜  go install sigs.k8s.io/kind@v0.9.0
```

非常的简单直观。需要注意的是 `go install <package>@<version>` 是从 1.16 开始增加的，无论你当前是否在一个模块下，此命令都会在 $GOPATH/bin 下安装指定版本的工具。

此外由于 Go 1.16 中 GO111MODULE 默认是打开的，go install 不会修改 go.mod 之类的文件，不会造成任何意外。

> 注意：
> @version 只能安装主软件包。非主程序包不受此格式约束。

## 关于不带 @version 的 go install

在模块外，不带 @version 是无法安装的，会有如下错误:

```bash
(MoeLove) ➜  go install  -v sigs.k8s.io/kind
go install: version is required when current directory is not in a module
        Try 'go install sigs.k8s.io/kind@latest' to install the latest version
```

如果你在模块目录中，并且你不带 @version 执行安装的话，只能安装 go.mod 中已经包含的版本。并且不能安装未出现在 go.mod 中的包。

```bash
(MoeLove) ➜  mkdir -p /go/src/github.com/moelove/iris

(MoeLove) ➜  cd /go/src/github.com/moelove/iris

# 初始化模块
(MoeLove) ➜  /go/src/github.com/moelove/iris go mod init
go: creating new go.mod: module github.com/moelove/iris

(MoeLove) ➜  /go/src/github.com/moelove/iris cat go.mod
module github.com/moelove/iris

go 1.16


# 不带 @version 无法安装
(MoeLove) ➜  /go/src/github.com/moelove/iris go install -v sigs.k8s.io/kind
no required module provides package sigs.k8s.io/kind; try 'go get -d sigs.k8s.io/kind' to add it

# 用 go get -d 下载
(MoeLove) ➜  /go/src/github.com/moelove/iris go get -d sigs.k8s.io/kind
go get: added sigs.k8s.io/kind v0.9.0

# 可以看到已经被添加到了模块依赖中
(MoeLove) ➜  /go/src/github.com/moelove/iris cat go.mod
module github.com/moelove/iris

go 1.16

require sigs.k8s.io/kind v0.9.0 // indirect

# 删除本地的 kind 工具
(MoeLove) ➜  /go/src/github.com/moelove/iris which kind
/go/bin/kind

(MoeLove) ➜  /go/src/github.com/moelove/iris rm /go/bin/kind

(MoeLove) ➜  /go/src/github.com/moelove/iris which kind

# 不带 @version 进行安装
(MoeLove) ➜  /go/src/github.com/moelove/iris go install -v sigs.k8s.io/kind
(MoeLove) ➜  /go/src/github.com/moelove/iris which kind
/go/bin/kind
(MoeLove) ➜  /go/src/github.com/moelove/iris kind version
kind v0.9.0 go1.16beta1 linux/amd64
```

## 关于 go get 和 go.mod

go get 将二进制安装相关的功能都转移到了 go install, 仅作为用于编辑 go.mod 文件的命令存在。在后续版本（计划是 Go 1.17）中删掉 go get 安装二进制的功能，接下来 go get 的行为就等同于我们现在执行 go get -d 命令了，仅需下载源码，并将依赖添加至 go.mod 即可。

## go.mod 如何编辑

在 Go 1.16 中，另一个行为变更是 go build 和 go test 不会自动编辑 go.mod 了，基于以上信息，Go 1.16 中将进行如下处理：

- 通过在代码中修改 import 语句，来修改 go.mod：

  - go get 可用于添加新模块；

  - go mod tidy 删除掉无用的模块；

- 将未导入的模块写入 go.mod:

  - go get <package>[@<version>];

  - go mod tidy 也可以；

  - 手动编辑；

## 从 1.15 升级需要注意什么？

由于 go build 和 go test 不会自动编辑 go.mod 了，所以可以将原本的行为通过 go mod tidy 共同处理。

## 总结

Go 1.16 中 go install 和 go get 方面有些不兼容的变更，但是 1.16 中模块更加简洁，减少了使用时的心智负担，我还是很期待这个版本的。
