
--------------------------------------------------------------------------------------

```bash
$ docker run -ti --rm --name golang golang:1.26.4 bash
root@8a9985d3bbfb:/go#
root@8a9985d3bbfb:/go# go version
go version go1.26.4 linux/amd64
root@8a9985d3bbfb:/go#
root@8a9985d3bbfb:/go# go help get
usage: go get [-t] [-u] [-tool] [build flags] [packages]

Get resolves its command-line arguments to packages at specific module versions,
updates go.mod to require those versions, and downloads source code into the
module cache.

To add a dependency for a package or upgrade it to its latest version:

	go get example.com/pkg

To upgrade or downgrade a package to a specific version:

	go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:

	go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:

	go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:

	go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages.
Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install'
may be used to build and install commands instead. When a version is specified,
'go install' runs in module-aware mode and ignores the go.mod file in the
current directory. For example:

	go install example.com/pkg@v1.2.3
	go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.

When the -t and -u flags are used together, get will update
test dependencies as well.

The -tool flag instructs go to add a matching tool line to go.mod for each
listed package. If -tool is used with @none, the line will be removed.

The -x flag prints commands as they are executed. This is useful for
debugging version control commands when a module is downloaded directly
from a repository.

For more about build flags, see 'go help build'.

For more about modules, see https://golang.org/ref/mod.

For more about using 'go get' to update the minimum Go version and
suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

See also: go build, go install, go clean, go mod.
root@8a9985d3bbfb:/go#
```

--------------------------------------------------------------------------------------

Q: 假设你是一位精通 golang 语言的高级开发专家, 我使用命令 `go help get` 查看了 go get 的帮助文档, 内容如下
```
root@8a9985d3bbfb:/go# go help get
usage: go get [-t] [-u] [-tool] [build flags] [packages]

Get resolves its command-line arguments to packages at specific module versions,
updates go.mod to require those versions, and downloads source code into the
module cache.

To add a dependency for a package or upgrade it to its latest version:

	go get example.com/pkg

To upgrade or downgrade a package to a specific version:

	go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:

	go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:

	go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:

	go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages.
Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install'
may be used to build and install commands instead. When a version is specified,
'go install' runs in module-aware mode and ignores the go.mod file in the
current directory. For example:

	go install example.com/pkg@v1.2.3
	go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.

When the -t and -u flags are used together, get will update
test dependencies as well.

The -tool flag instructs go to add a matching tool line to go.mod for each
listed package. If -tool is used with @none, the line will be removed.

The -x flag prints commands as they are executed. This is useful for
debugging version control commands when a module is downloaded directly
from a repository.

For more about build flags, see 'go help build'.

For more about modules, see https://golang.org/ref/mod.

For more about using 'go get' to update the minimum Go version and
suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

See also: go build, go install, go clean, go mod.
root@8a9985d3bbfb:/go#
```
请帮我将上述帮助文档翻译为中文, 详细总结上述帮助文档的所有内容, 并且根据不同的场景举例说明 go get 的用法

您好! 作为一名长期使用 Go 语言的开发者, 我非常乐意为您详细解析 `go help get` 的官方文档.

随着 Go 1.16+ 尤其是后续版本(如 Go 1.24 引入的内置工具链管理和工具依赖项机制)的演进, `go get` 的定位已经发生了根本性的变化.

以下为您提供中文翻译、详细总结以及常见场景的实战示例.

---

### 一、`go help get` 帮助文档中文翻译

```bash
root@8a9985d3bbfb:/go# go help get
usage: go get [-t] [-u] [-tool] [build flags] [packages]

Get resolves its command-line arguments to packages at specific module versions,
updates go.mod to require those versions, and downloads source code into the
module cache.
Get 用于将命令行参数解析为特定模块版本的包, 更新 go.mod 文件以要求(require)这些版本, 并向模块缓存(module cache)中下载源代码.

To add a dependency for a package or upgrade it to its latest version:
要添加某个包作为依赖项, 或将其升级到最新版本:

	go get example.com/pkg

To upgrade or downgrade a package to a specific version:
要将某个包升级或降级到特定版本:

	go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:
要移除对某个模块的依赖, 并降级其他需要它的模块:

	go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:
要将项目所需的最低 Go 版本升级到最新发布的 Go 版本:

	go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:
要将 Go 工具链(Go toolchain)升级到当前 Go 工具链的最新补丁版本(patch release):

	go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages.
Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install'
may be used to build and install commands instead. When a version is specified,
'go install' runs in module-aware mode and ignores the go.mod file in the
current directory. For example:
在 Go 的早期版本中, "go get" 用于构建和安装包(译者注: 即生成可执行二进制文件).
现在, "go get" 专用于调整 go.mod 中的依赖关系.
若要构建和安装命令(译者注: 指 CLI 工具等可执行文件), 应改用 "go install".
当指定了版本时, "go install" 会在模块感知模式(module-aware mode)下运行, 并忽略当前目录中的 go.mod 文件. 例如:

	go install example.com/pkg@v1.2.3
	go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.
"go get" 接受以下标志(flags):

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.
-t 标志指示 get 同时考虑构建命令行中指定的包的测试(test)所需要的模块.

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.
-u 标志指示 get 升级那些提供命令行中所指定包的依赖项的模块, 以便在有较新的次要版本(minor)或补丁版本(patch)可用时使用它们.

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.
-u=patch 标志(注意: 不是 -u patch)同样指示 get 升级依赖项, 但将默认行为改为选择补丁版本(patch releases).

When the -t and -u flags are used together, get will update
test dependencies as well.
当 -t 和 -u 标志一起使用时, get 也会同时升级测试依赖项.

The -tool flag instructs go to add a matching tool line to go.mod for each
listed package. If -tool is used with @none, the line will be removed.
-tool 标志指示 go 为列出的每个包在 go.mod 中添加一行匹配的 tool 声明. 如果 -tool 与 @none 一起使用, 该行将被移除.

The -x flag prints commands as they are executed. This is useful for
debugging version control commands when a module is downloaded directly
from a repository.
-x 标志会在执行命令时将其打印出来. 这在直接从版本控制系统(VCS)仓库下载模块时, 对于调试版本控制命令非常有用.

For more about build flags, see 'go help build'.

For more about modules, see https://golang.org/ref/mod.

For more about using 'go get' to update the minimum Go version and
suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

See also: go build, go install, go clean, go mod.
root@8a9985d3bbfb:/go#

```

---

### 二、核心内容详细总结

通过这份文档, 我们可以清晰地看到现代 Go 语言中对依赖管理的设计思路:

1. 定位的变化(彻底剥离"安装"职责)

   * 以前的 `go get` 既负责下载依赖, 又负责编译并安装二进制到 `$GOPATH/bin`.
   * 现在的 `go get` 只负责修改依赖关系(修改 `go.mod` / `go.sum`)并将源码下载到模块缓存中.
   * 任何可执行工具的安装, 均已被分流给 `go install`. `go install` 在指定版本(如 `@latest`)时会忽略当前项目的 `go.mod`, 确保不会污染你当前项目的依赖树.

2. 多维度的版本控制

   * 支持通过 `@latest` 引入最新版.
   * 支持通过 `@v1.2.3` 锁死特定版本.
   * 支持通过 `@none` 物理删除依赖.

3. 现代 Go 工具链与 Go 版本的就地升级

   * Go 1.21+ 引入了更强大的前向兼容和工具链管理机制. 现在可以直接通过 `go get go@latest` 和 `go get toolchain@patch` 在 `go.mod` 中声明和升级项目期望的 Go 运行时环境和工具链版本, 无需手动修改 `go.mod` 文件.

4. 对传递依赖(Transitive Dependencies)的升级策略

   * 默认情况下, `go get` 仅会下载当前包需要的依赖.
   * `-u` 标志允许递归升级当前包的子依赖(次要版本或补丁版本).
   * `-u=patch` 提供了更安全的升级方式, 只允许子依赖升级补丁版本(修 bug, 而不引入新特性).

5. 项目本地"开发工具"的管理(Go 1.24+ 新增)

   * 新增的 `-tool` 标志解决了一个历史痛点: 在此之前, 管理项目所需的开发工具(如代码生成器 `stringer`、格式化工具 `goimports` 等)通常需要写一个不参与编译的 `tools.go` 文件.

   * 现在, 通过 `go get -tool <path>` 可以直接在 `go.mod` 中注册该工具(会生成 `tool <path>` 声明), 随后即可通过 `go tool <command>` 直接运行, 实现了项目开发工具的标准化和版本锁定.

---

### 三、场景实战与用法示例

在日常的企业级或开源项目开发中, 以下是 `go get` 最常用的应用场景:

#### 场景 1: 在项目中引入一个全新的第三方库, 或升级到其最新的 Minor/Patch 版本

当你需要使用某个第三方库(例如 `gin` 框架)时, 在项目根目录下执行:

```bash
go get github.com/gin-gonic/gin
```

效果: Go 会自动寻找该库最新的语义化版本(比如 `v1.10.0`), 将其写入 `go.mod` 的 `require` 指令中, 同时下载源码至缓存.

#### 场景 2: 将某个依赖锁定到特定的 Bug 修复版本或指定版本

如果最新版存在兼容性问题, 或者出于安全考量需要指定精确的版本:

```bash
go get github.com/gin-gonic/gin@v1.9.1
```

效果: 强行将项目的 `go.mod` 中该依赖的版本指定为 `v1.9.1`. 无论本地缓存是否有该版本, Go 都会下载并应用它.

#### 场景 3: 删除某个不再使用的依赖, 并清理其子依赖

当你想在代码中废弃某个库(例如 `github.com/foo/bar`)并彻底移除其在 `go.mod` 中的声明时:

```bash
go get github.com/foo/bar@none
```

效果: 从 `go.mod` 中移出该依赖, 并自动降级或删除仅由该依赖引入的其他子依赖.
专家建议: 执行完 `@none` 后, 通常需要配合运行一次 `go mod tidy`, 以确保整个依赖树的干净和一致.

#### 场景 4: 安全地升级项目中已有依赖的子依赖(仅限补丁升级)

如果你想升级项目内 `github.com/my/pkg` 的所有依赖, 但担心引入次要版本(Minor Version)带来的潜在不兼容问题, 只希望升级安全的 Bug 修复版(Patch Version):

```bash
go get -u=patch github.com/my/pkg
```

效果: 它会寻找并升级 `github.com/my/pkg` 所引用的所有底层子依赖, 但限制其只升级到最新的 Patch 版本(例如从 `v1.2.1` 升级到 `v1.2.5`, 而绝不会升级到 `v1.3.0`).

#### 场景 5: 升级项目所使用的最低 Go 版本

当你想要在项目中使用 Go 的新特性(如新增的标准库或语法), 需要修改 `go.mod` 里的 `go` 版本声明:

```bash
go get go@latest
```

效果: 会将 `go.mod` 中的 `go 1.xx` 自动修改为您当前安装的 Go 编译器所支持的最新的 Go 语义版本.

#### 场景 6: [现代特性]管理项目本地的代码生成或 Lint 工具 (Go 1.24+)

如果你想在项目中使用代码生成工具 `stringer` 或 Linter 工具, 并让团队中的所有人在运行该工具时都保持版本一致:

```bash
go get -tool golang.org/x/tools/cmd/stringer@v0.30.0
```

效果:
  1. 下载该工具并锁定其版本.
  2. 自动在项目的 `go.mod` 中添加一行 `tool golang.org/x/tools/cmd/stringer`.
  3. 以后, 团队成员无需在本地全局安装 `stringer`, 只需在项目目录下执行 `go tool stringer` 即可直接调用该指定版本的工具.

如果要移除该工具:

  ```bash
  go get -tool golang.org/x/tools/cmd/stringer@none
  ```

#### 场景 7: 误区纠正 - 安装全局可执行文件时(请不要使用 go get)

如果你在网上看到旧教程让你通过 `go get github.com/some/cli-tool` 来安装一个终端命令行工具, 在现代 Go 中请使用:

```bash
go install github.com/some/cli-tool@latest
```
效果: 它会把工具编译并安装到 `$GOPATH/bin`(或整个系统的 PATH 中), 并且完全不会修改你当前所在项目的 `go.mod` 依赖.

--------------------------------------------------------------------------------------

Q: 上述 go get 帮助文档中有关于 `-u` 和 `-u=patch` 的说明, 请通过实际的使用场景举例说明他们的作用

在 Go 的模块(Go Modules)管理中, `-u` 和 `-u=patch` 标志主要用于控制在引入或升级目标包时, 其底层的"传递依赖"(即依赖的依赖, Transitive Dependencies)该如何同步升级.

为了让您直观地理解它们的作用, 我们先构建一个非常贴近实际开发的场景.

---

### 场景准备

假设您的项目 `my-project` 依赖了一个名为 `SDK` 的库, 而 `SDK` 又依赖了一个名为 `Logger` 的日志库.

目前, 您项目中的 `go.mod` 状态如下:
* `SDK` 的版本锁定在: `v1.0.0`
* `Logger` 的版本锁定在: `v1.2.0`(这是 `SDK v1.0.0` 声明的最低兼容版本)

此时, 这两个库的开发者在远程仓库发布了新版本:
* `SDK` 发布了新版本: `v1.1.0`(其内部声明对 `Logger` 的要求仍然是 `>= v1.2.0`)
* `Logger` 陆续发布了两个新版本:
  * `v1.2.1`(这是一个 Patch/补丁 版本, 只修复了 Bug)
  * `v1.3.0`(这是一个 Minor/次要 版本, 引入了新特性)

现在, 您决定在项目根目录下通过 `go get` 升级 `SDK` 库.

---

### 场景 A: 使用默认的 `go get`(不带任何 `-u` 标志)

您在终端执行:

```bash
go get github.com/user/sdk
```

* 执行结果:
  * `SDK` 被升级到: `v1.1.0`
  * `Logger` 保持不变, 依然是: `v1.2.0`

* 原理解析:
  Go 模块采用的是最小版本选择(MVS, Minimal Version Selection)算法.
  因为新版 `SDK v1.1.0` 声明它只需要 `Logger >= v1.2.0` 即可工作, 而您项目中已经锁定了 `Logger v1.2.0`, 完全满足它的要求. 为了保证项目的极致稳定性, Go 默认不会去主动帮您下载更新的 `Logger v1.2.1` 或 `v1.3.0`.

---

### 场景 B: 使用 `go get -u`(全量升级依赖项)

您在终端执行:

```bash
go get -u github.com/user/sdk
```

* 执行结果:
  * `SDK` 被升级到: `v1.1.0`
  * `Logger` 被强制升级到: `v1.3.0`(最新的 Minor 版本)

* 原理解析:
  `-u` 标志(即 `upgrade`)会明确指示 Go 编译器: "在升级 `SDK` 的同时, 把 `SDK` 依赖的所有子包, 都顺便升级到它们各自最新的 Minor 或 Patch 版本."
  因此, Go 发现了 `Logger` 有更新的 `v1.3.0`, 便会直接将其一并拉取并写入您的 `go.mod`.

* 优缺点及适用场景:
  * 优点: 能够让您的项目享受到子依赖库的最全功能和最新优化.
  * 缺点/风险: 尽管语义化版本(SemVer)规定 Minor 版本(如 `v1.3.0`)应该保持向后兼容, 但有时第三方库的开发者可能会不小心引入破坏性变更(Breaking Changes), 这会导致您的项目在编译或运行时报错.

---

### 场景 C: 使用 `go get -u=patch`(安全、受控的补丁升级)

您在终端执行:

```bash
go get -u=patch github.com/user/sdk
```
(注意: 语法是 `-u=patch`, 中间是等号, 而不是 `-u patch`)

* 执行结果:
  * `SDK` 被升级到: `v1.1.0`
  * `Logger` 被升级到: `v1.2.1`(而不是 `v1.3.0`)

* 原理解析:
  `-u=patch` 会指示 Go 编译器: "在升级 `SDK` 的同时, 允许对子依赖项进行升级, 但严格限制在相同的 Minor 版本内, 仅获取最新的 Patch(补丁)版本."
  因为您项目原先锁定的 `Logger` 处于 `v1.2.x` 分支, Go 只会在 `v1.2.x` 范围内寻找最新的补丁, 即 `v1.2.1`. 它会主动忽略代表新特性的 `v1.3.0`.

* 优缺点及适用场景:
  * 优点: 极高安全性. Patch 版本在语义化版本中只允许用于"修复安全漏洞和 Bug", 不包含新功能, 因此几乎不会带来任何代码不兼容的风险.
  * 适用场景: 非常适合在生产环境进行例行的安全漏洞修补(Security Patching)或例行维护.

---

### 总结对比表

| 命令                  | `SDK` 结果 | `Logger` 结果 | 核心设计目的                                           |
| :-------------------- | :--------- | :------------ | :----------------------------------------------------- |
| `go get sdk`          | `v1.1.0`   | `v1.2.0`      | 保持依赖绝对稳定, 能不动的子依赖绝不动(默认策略).      |
| `go get -u sdk`       | `v1.1.0`   | `v1.3.0`      | 获取最新特性, 适合积极跟进社区最新演进的项目.          |
| `go get -u=patch sdk` | `v1.1.0`   | `v1.2.1`      | 获取子依赖的安全和 bug 修复, 同时将不兼容风险降到最低. |

--------------------------------------------------------------------------------------

Q: 上述 go get 帮助文档中有关于 `-tool` 的说明, 请通过实际的使用场景举例说明他们的作用

在 Go 1.24 之前, 管理项目级别的"开发工具"(例如代码生成器 `stringer`、Linter 工具 `golangci-lint`、Mock 生成器 `mockgen` 等)一直是 Go 社区的一个痛点.

为了说明 `go get` 中的 `-tool` 标志的作用, 我们可以从历史痛点、现代解决方案以及实战演练三个维度来剖析.

---

### 一、痛点背景: 过去我们是如何管理工具的?

在 Go 1.24 之前, 如果你希望团队里的所有人使用完全相同版本的开发工具(如 `stringer`), 你无法直接在 `go.mod` 中声明它, 因为 `go.mod` 只用于记录代码中 `import` 导入的运行时依赖包.

社区过去普遍采用一个被称为 `tools.go` 避孕套/占位符的 Hack 方案:
1. 在项目里建一个名为 `tools.go` 的文件.
2. 使用 `//go:build tools` 条件编译指令, 防止它被编入最终的二进制文件.
3. 在文件里硬写一些未使用的 `import`, 例如 `import _ "golang.org/x/tools/cmd/stringer"`.
4. 只有这样, `go mod tidy` 才不会把这个工具依赖当作"无用依赖"自动删掉.

这种方式虽然行之有效, 但略显臃肿和非官方化.

---

### 二、现代方案: `go get -tool` 做了什么?

Go 1.24 引入了原生的工具依赖管理机制.

* 运行 `go get -tool <path>` 时, Go 会直接在 `go.mod` 中添加一行 `tool <path>` 声明.
* 该工具及其依赖会被下载到本地的模块缓存(Module Cache)中.
* 不需要创建任何 dummy 文件, `go mod tidy` 会自动识别并保留这些工具.
* 可以通过 `go tool <tool_name>` 直接调用它们, 无需事先在操作系统全局安装.

---

### 三、场景实战: 使用 `stringer` 自动生成枚举的 `String()` 方法

假设您正在编写一个状态机的微服务, 其中定义了一个 `State` 枚举类型. 您希望使用官方的 `stringer` 工具来自动为它生成 `String()` 方法.

#### 步骤 1: 添加工具到项目中

在项目根目录下, 使用 `-tool` 标志将 `stringer` 声明为该项目的开发工具:

```bash
go get -tool golang.org/x/tools/cmd/stringer@v0.30.0
```

(注: 建议加上精确版本号以确保团队一致. 如果不加, 默认获取最新版 `@latest`)

此时 `go.mod` 发生的变化:
打开 `go.mod` 可能会发现多了类似的指令:

```go
module my-project

go 1.24.0

// 自动生成这行, 用于记录该项目绑定的工具
tool golang.org/x/tools/cmd/stringer

require (
    // 同时也会在 require 中记录该工具所处的包版本, 从而精确锁定其依赖
    golang.org/x/tools v0.30.0
)
```

#### 步骤 2: 在项目中使用该工具

现在, 您可以在 Go 代码文件中安全地写上 `//go:generate` 声明. 例如在 `state.go` 中:

```go
package state

//go:generate go tool stringer -type=State

type State int

const (
    Start State = iota
    Running
    Stop
)
```

注意这里的写法: 以前写的是 `//go:generate stringer ...`(这要求每个开发者的电脑全局 PATH 里都必须提前装好 `stringer`, 否则会报错).
现在直接写: `go tool stringer`.

#### 步骤 3: 一键生成

在项目根目录下执行 Go 标准的代码生成命令:

```bash
go generate ./...
```

* 效果: Go 会查找 `go.mod` 里的 `tool` 声明, 编译缓存中对应 `v0.30.0` 版本的 `stringer` 并临时运行它, 完美生成 `state_string.go` 文件.
* 优势: 哪怕团队里的新同事刚入职、电脑里什么全局工具都没装, 只要他们本地安装了 Go 编译器, 运行该命令就会自动下载并运行该工具, 真正做到了"开箱即用、版本一致".

#### 步骤 4: 移除不需要的工具

如果后续项目重构, 不再需要 `stringer` 工具, 可以使用帮助文档里提到的 `@none` 方式:

```bash
go get -tool golang.org/x/tools/cmd/stringer@none
```

效果: Go 会自动在 `go.mod` 中删除对应的 `tool golang.org/x/tools/cmd/stringer` 声明, 并且在运行 `go mod tidy` 后, 该工具引入的依赖也会被彻底清理.

---

### 四、核心总结

`-tool` 标志极大地规范了 Go 团队协作的工程化体验:
1. 团队版本对齐: 工具的版本直接和项目代码库绑定并提交到 Git. 不再会出现"因为我本地的 `stringer` 比你新, 所以我们生成的代码格式不一样"的尴尬情况.
2. 零安装负担: 新开发者拉下代码后, 不需要去额外 `go install` 任何 CLI 工具, Go 依靠 `go.mod` 就能在执行 `go tool` 时自动在后台下载和管理它们.

--------------------------------------------------------------------------------------
