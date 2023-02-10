# Go 项目生成版本号

## 版本号的作用

版本信息管理，是项目开发中需要考虑的问题。尤其在各类开源软件中，重要的功能特性一定需要版本号绑定。通过版本号，用户才能知道该程序提供了哪些功能。

那么，如何为项目添加版本号呢？很多人应该都使用过硬编码方式，即将版本号直接写入源码或者配置文件，每次功能升级就修改版本号。这种方式，显然是可行的，但是也容易出错。一是发版时，容易忘记更新版本号，二是多个分支代码合并时，也可能搞混。

下文，给大家带来一个不一样的管理方案。

## ldflags -X 变量传递

go 链接器 Linker 是组装二进制文件的工具，我们在执行 go build 命令时，可以通过 --ldflags 标志设定链接器参数，使用以下语句可查看链接器可选参数。

```bash
$ go version
go version go1.18.1 darwin/amd64

$ go build --ldflags="--help"
usage: link [options] main.o
  -B note
    	add an ELF NT_GNU_BUILD_ID note when using ELF
  -E entry
    	set entry symbol name
  -H type
    	set header type
  -I linker
    	use linker as ELF dynamic linker
  -L directory
    	add specified directory to library path
  -R quantum
    	set address rounding quantum (default -1)
  -T address
    	set text segment address (default -1)
  -V	print version and exit
  -X definition
    	add string value definition of the form importpath.name=value
  -a	no-op (deprecated)
  -asan
    	enable ASan interface
  -aslr
    	enable ASLR for buildmode=c-shared on windows (default true)
  -benchmark string
    	set to 'mem' or 'cpu' to enable phase benchmarking
  -benchmarkprofile base
    	emit phase profiles to base_phase.{cpu,mem}prof
  -buildid id
    	record id as Go toolchain build id
  -buildmode mode
    	set build mode
  -c	dump call graph
  -compressdwarf
    	compress DWARF if possible (default true)
  -cpuprofile file
    	write cpu profile to file
  -d	disable dynamic executable
  -debugtextsize int
    	debug text section max size
  -debugtramp int
    	debug trampolines
  -dumpdep
    	dump symbol dependency graph
  -extar string
    	archive program for buildmode=c-archive
  -extld linker
    	use linker when linking in external mode
  -extldflags flags
    	pass flags to external linker
  -f	ignore version mismatch
  -g	disable go package data checks
  -h	halt on error
  -importcfg file
    	read import configuration from file
  -installsuffix suffix
    	set package directory suffix
  -k symbol
    	set field tracking symbol
  -libgcc string
    	compiler support lib for internal linking; use "none" to disable
  -linkmode mode
    	set link mode
  -linkshared
    	link against installed Go shared libraries
  -memprofile file
    	write memory profile to file
  -memprofilerate rate
    	set runtime.MemProfileRate to rate
  -msan
    	enable MSan interface
  -n	dump symbol table
  -o file
    	write output to file
  -pluginpath string
    	full path name for plugin
  -r path
    	set the ELF dynamic linker search path to dir1:dir2:...
  -race
    	enable race detector
  -s	disable symbol table
  -strictdups int
    	sanity check duplicate symbol contents during object file reading (1=warn 2=err).
  -tmpdir directory
    	use directory for temporary files
  -v	print link trace
  -w	disable DWARF generation
```

参数很多，但我们感兴趣的是 -X

```bash
$ go build --ldflags="--help"
  -X definition
    	add string value definition of the form importpath.name=value

```

-X 参数，指定 importpath.name=value，用于修改变量值。其中 importpath 表示包导入路径，name 是程序中的变量名，value 代表我们想要设定的变量值。

下面，我们通过示例项目来具体感受一下。

```bash
$ mkdir versionDemo

$ cd versionDemo

$ go mod init versiondemo
go: creating new go.mod: module versiondemo

$ touch main.go
```

在 main 函数中，我们打印 version 值。

```go
package main

import "fmt"

var version = "0.0.1"

func main() {
    fmt.Println("version: ", version)
}
```

如果正常编译执行程序，将得到以下结果

```bash
$ go build -o main
$ ./main
version:  0.0.1
```

此时，我们指定 --ldflags 的 -X 参数重新编译执行

```bash
$ go build -o main --ldflags="-X 'main.version=client-0.0.2'"
$ ./main
version: client-0.0.2
```

可以看到 version 参数值已经被改变。

## 添加 git 信息

开发中需要使用 git 工具，本文讨论的版本管理，也经常与 git tag 挂钩。那其实有更酷的操作：我们可以在构建期间，通过 git commit 信息自动填充版本号。

我们基于上文项目目录，添加 git commit 信息。

```bash
$ git init
$ git add .
$ git commit -m "initial commit"
```
通过以下命令，可拿到 git commit 的 hash 值

```bash
$ git rev-parse HEAD
46dab0ddb6ba20445c2c1f047575e25d3aad1a27
```

该值较长，我们可以添加 --short 选项获取短 hash 值。

```bash
$ git rev-parse --short HEAD
46dab0d
```

此时，通过指定 --ldflags 的 -X 参数，将 version 值替换成 git commit 的哈希值。这样，我们成功地将项目版本与 git 信息绑定在了一起。

```bash
$ go build -o main --ldflags="-X 'main.version=$(git rev-parse --short HEAD)'"
$ ./main
version:  46dab0d
```

## 总结

本文介绍了一种如何通过 ldflags -X 变量传递的方式。使用这种方式我们可以在构建时轻松设定一些元信息，例如本文示例的程序版本信息。而这种构建的动作不应该手动去执行，而是放入到 CI/CD 流程中，让整个过程变得更加丝滑。
