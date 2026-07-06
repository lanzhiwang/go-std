
--------------------------------------------------------------------------------------

```bash
$ docker run -ti --rm --name golang golang:1.26.4 bash
root@82157ef140f0:/go#
root@82157ef140f0:/go# go version
go version go1.26.4 linux/amd64
root@82157ef140f0:/go#
root@82157ef140f0:/go# go help list
usage: go list [-f format] [-json] [-m] [list flags] [build flags] [packages]

List lists the named packages, one per line.
The most commonly-used flags are -f and -json, which control the form
of the output printed for each package. Other list flags, documented below,
control more specific details.

The default output shows the package import path:

    bytes
    encoding/json
    github.com/gorilla/mux
    golang.org/x/net/html

The -f flag specifies an alternate format for the list, using the
syntax of package template. The default output is equivalent
to -f '{{.ImportPath}}'. The struct being passed to the template is:

    type Package struct {
        Dir            string   // directory containing package sources
        ImportPath     string   // import path of package in dir
        ImportComment  string   // path in import comment on package statement
        Name           string   // package name
        Doc            string   // package documentation string
        Target         string   // install path
        Shlib          string   // the shared library that contains this package (only set when -linkshared)
        Goroot         bool     // is this package in the Go root?
        Standard       bool     // is this package part of the standard Go library?
        Stale          bool     // would 'go install' do anything for this package?
        StaleReason    string   // explanation for Stale==true
        Root           string   // Go root or Go path dir containing this package
        ConflictDir    string   // this directory shadows Dir in $GOPATH
        BinaryOnly     bool     // binary-only package (no longer supported)
        ForTest        string   // package is only for use in named test
        Export         string   // file containing export data (when using -export)
        BuildID        string   // build ID of the compiled package (when using -export)
        Module         *Module  // info about package's containing module, if any (can be nil)
        Match          []string // command-line patterns matching this package
        DepOnly        bool     // package is only a dependency, not explicitly listed
        DefaultGODEBUG string  // default GODEBUG setting, for main packages

        // Source files
        GoFiles           []string   // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
        CgoFiles          []string   // .go source files that import "C"
        CompiledGoFiles   []string   // .go files presented to compiler (when using -compiled)
        IgnoredGoFiles    []string   // .go source files ignored due to build constraints
        IgnoredOtherFiles []string // non-.go source files ignored due to build constraints
        CFiles            []string   // .c source files
        CXXFiles          []string   // .cc, .cxx and .cpp source files
        MFiles            []string   // .m source files
        HFiles            []string   // .h, .hh, .hpp and .hxx source files
        FFiles            []string   // .f, .F, .for and .f90 Fortran source files
        SFiles            []string   // .s source files
        SwigFiles         []string   // .swig files
        SwigCXXFiles      []string   // .swigcxx files
        SysoFiles         []string   // .syso object files to add to archive
        TestGoFiles       []string   // _test.go files in package
        XTestGoFiles      []string   // _test.go files outside package

        // Embedded files
        EmbedPatterns      []string // //go:embed patterns
        EmbedFiles         []string // files matched by EmbedPatterns
        TestEmbedPatterns  []string // //go:embed patterns in TestGoFiles
        TestEmbedFiles     []string // files matched by TestEmbedPatterns
        XTestEmbedPatterns []string // //go:embed patterns in XTestGoFiles
        XTestEmbedFiles    []string // files matched by XTestEmbedPatterns

        // Cgo directives
        CgoCFLAGS    []string // cgo: flags for C compiler
        CgoCPPFLAGS  []string // cgo: flags for C preprocessor
        CgoCXXFLAGS  []string // cgo: flags for C++ compiler
        CgoFFLAGS    []string // cgo: flags for Fortran compiler
        CgoLDFLAGS   []string // cgo: flags for linker
        CgoPkgConfig []string // cgo: pkg-config names

        // Dependency information
        Imports      []string          // import paths used by this package
        ImportMap    map[string]string // map from source import to ImportPath (identity entries omitted)
        Deps         []string          // all (recursively) imported dependencies
        TestImports  []string          // imports from TestGoFiles
        XTestImports []string          // imports from XTestGoFiles

        // Error information
        Incomplete bool            // this package or a dependency has an error
        Error      *PackageError   // error loading package
        DepsErrors []*PackageError // errors loading dependencies
    }

Packages stored in vendor directories report an ImportPath that includes the
path to the vendor directory (for example, "d/vendor/p" instead of "p"),
so that the ImportPath uniquely identifies a given copy of a package.
The Imports, Deps, TestImports, and XTestImports lists also contain these
expanded import paths. See golang.org/s/go15vendor for more about vendoring.

The error information, if any, is

    type PackageError struct {
        ImportStack   []string // shortest path from package named on command line to this one
        Pos           string   // position of error (if present, file:line:col)
        Err           string   // the error itself
    }

The module information is a Module struct, defined in the discussion
of list -m below.

The template function "join" calls strings.Join.

The template function "context" returns the build context, defined as:

    type Context struct {
        GOARCH        string   // target architecture
        GOOS          string   // target operating system
        GOROOT        string   // Go root
        GOPATH        string   // Go path
        CgoEnabled    bool     // whether cgo can be used
        UseAllFiles   bool     // use files regardless of //go:build lines, file names
        Compiler      string   // compiler to assume when computing target paths
        BuildTags     []string // build constraints to match in //go:build lines
        ToolTags      []string // toolchain-specific build constraints
        ReleaseTags   []string // releases the current release is compatible with
        InstallSuffix string   // suffix to use in the name of the install dir
    }

For more information about the meaning of these fields see the documentation
for the go/build package's Context type.

The -json flag causes the package data to be printed in JSON format
instead of using the template format. The JSON flag can optionally be
provided with a set of comma-separated required field names to be output.
If so, those required fields will always appear in JSON output, but
others may be omitted to save work in computing the JSON struct.

The -compiled flag causes list to set CompiledGoFiles to the Go source
files presented to the compiler. Typically this means that it repeats
the files listed in GoFiles and then also adds the Go code generated
by processing CgoFiles and SwigFiles. The Imports list contains the
union of all imports from both GoFiles and CompiledGoFiles.

The -deps flag causes list to iterate over not just the named packages
but also all their dependencies. It visits them in a depth-first post-order
traversal, so that a package is listed only after all its dependencies.
Packages not explicitly listed on the command line will have the DepOnly
field set to true.

The -e flag changes the handling of erroneous packages, those that
cannot be found or are malformed. By default, the list command
prints an error to standard error for each erroneous package and
omits the packages from consideration during the usual printing.
With the -e flag, the list command never prints errors to standard
error and instead processes the erroneous packages with the usual
printing. Erroneous packages will have a non-empty ImportPath and
a non-nil Error field; other information may or may not be missing
(zeroed).

The -export flag causes list to set the Export field to the name of a
file containing up-to-date export information for the given package,
and the BuildID field to the build ID of the compiled package.

The -find flag causes list to identify the named packages but not
resolve their dependencies: the Imports and Deps lists will be empty.
With the -find flag, the -deps, -test and -export commands cannot be
used.

The -test flag causes list to report not only the named packages
but also their test binaries (for packages with tests), to convey to
source code analysis tools exactly how test binaries are constructed.
The reported import path for a test binary is the import path of
the package followed by a ".test" suffix, as in "math/rand.test".
When building a test, it is sometimes necessary to rebuild certain
dependencies specially for that test (most commonly the tested
package itself). The reported import path of a package recompiled
for a particular test binary is followed by a space and the name of
the test binary in brackets, as in "math/rand [math/rand.test]"
or "regexp [sort.test]". The ForTest field is also set to the name
of the package being tested ("math/rand" or "sort" in the previous
examples).

The Dir, Target, Shlib, Root, ConflictDir, and Export file paths
are all absolute paths.

By default, the lists GoFiles, CgoFiles, and so on hold names of files in Dir
(that is, paths relative to Dir, not absolute paths).
The generated files added when using the -compiled and -test flags
are absolute paths referring to cached copies of generated Go source files.
Although they are Go source files, the paths may not end in ".go".

The -m flag causes list to list modules instead of packages.

When listing modules, the -f flag still specifies a format template
applied to a Go struct, but now a Module struct:

    type Module struct {
        Path       string        // module path
        Query      string        // version query corresponding to this version
        Version    string        // module version
        Versions   []string      // available module versions
        Replace    *Module       // replaced by this module
        Time       *time.Time    // time version was created
        Update     *Module       // available update (with -u)
        Main       bool          // is this the main module?
        Indirect   bool          // module is only indirectly needed by main module
        Dir        string        // directory holding local copy of files, if any
        GoMod      string        // path to go.mod file describing module, if any
        GoVersion  string        // go version used in module
        Retracted  []string      // retraction information, if any (with -retracted or -u)
        Deprecated string        // deprecation message, if any (with -u)
        Error      *ModuleError  // error loading module
        Sum        string        // checksum for path, version (as in go.sum)
        GoModSum   string        // checksum for go.mod (as in go.sum)
        Origin     any           // provenance of module
        Reuse      bool          // reuse of old module info is safe
    }

    type ModuleError struct {
        Err string // the error itself
    }

The file GoMod refers to may be outside the module directory if the
module is in the module cache or if the -modfile flag is used.

The default output is to print the module path and then
information about the version and replacement if any.
For example, 'go list -m all' might print:

    my/main/module
    golang.org/x/text v0.3.0 => /tmp/text
    rsc.io/pdf v0.1.1

The Module struct has a String method that formats this
line of output, so that the default format is equivalent
to -f '{{.String}}'.

Note that when a module has been replaced, its Replace field
describes the replacement module, and its Dir field is set to
the replacement's source code, if present. (That is, if Replace
is non-nil, then Dir is set to Replace.Dir, with no access to
the replaced source code.)

The -u flag adds information about available upgrades.
When the latest version of a given module is newer than
the current one, list -u sets the Module's Update field
to information about the newer module. list -u will also set
the module's Retracted field if the current version is retracted.
The Module's String method indicates an available upgrade by
formatting the newer version in brackets after the current version.
If a version is retracted, the string "(retracted)" will follow it.
For example, 'go list -m -u all' might print:

    my/main/module
    golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
    rsc.io/pdf v0.1.1 (retracted) [v0.1.2]

(For tools, 'go list -m -u -json all' may be more convenient to parse.)

The -versions flag causes list to set the Module's Versions field
to a list of all known versions of that module, ordered according
to semantic versioning, earliest to latest. The flag also changes
the default output format to display the module path followed by the
space-separated version list.

The -retracted flag causes list to report information about retracted
module versions. When -retracted is used with -f or -json, the Retracted
field explains why the version was retracted.
The strings are taken from comments on the retract directive in the
module's go.mod file. When -retracted is used with -versions, retracted
versions are listed together with unretracted versions. The -retracted
flag may be used with or without -m.

The arguments to list -m are interpreted as a list of modules, not packages.
The main module is the module containing the current directory.
The active modules are the main module and its dependencies.
With no arguments, list -m shows the main module.
With arguments, list -m shows the modules specified by the arguments.
Any of the active modules can be specified by its module path.
The special pattern "all" specifies all the active modules, first the main
module and then dependencies sorted by module path.
A pattern containing "..." specifies the active modules whose
module paths match the pattern.
A query of the form path@version specifies the result of that query,
which is not limited to active modules.
See 'go help modules' for more about module queries.

The template function "module" takes a single string argument
that must be a module path or query and returns the specified
module as a Module struct. If an error occurs, the result will
be a Module struct with a non-nil Error field.

When using -m, the -reuse=old.json flag accepts the name of file containing
the JSON output of a previous 'go list -m -json' invocation with the
same set of modifier flags (such as -u, -retracted, and -versions).
The go command may use this file to determine that a module is unchanged
since the previous invocation and avoid redownloading information about it.
Modules that are not redownloaded will be marked in the new output by
setting the Reuse field to true. Normally the module cache provides this
kind of reuse automatically; the -reuse flag can be useful on systems that
do not preserve the module cache.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.

For more about modules, see https://golang.org/ref/mod.
root@82157ef140f0:/go#
```

--------------------------------------------------------------------------------------

Q: 假设你是一位精通 golang 语言的高级开发专家, 我使用命令 `go help list` 查看了 go list 的帮助文档, 内容如下
```
```
请帮我详细总结上述帮助文档的所有内容, 并且根据不同的场景举例说明 go list 的用法

`go list` 是 Go 语言工具链中一个非常强大且灵活的辅助工具, 主要用于检索、查询和分析 Go 包(packages)或模块(modules)的元数据. 它不仅可以输出包的导入路径, 还可以配合模板(`-f`)或 JSON(`-json`)输出包的内部结构、依赖关系、源文件列表以及模块版本等丰富的信息.

以下是对您提供的 `go help list` 帮助文档内容的详细总结, 以及在实际开发中常见场景下的使用示例.

---

### 一、`go list` 核心内容总结

#### 1. 两种运行模式

* 包模式(默认): 对 Go 包进行检索. 元数据映射为 `Package` 结构体, 默认输出包的导入路径(`ImportPath`).

* 模块模式(使用 `-m` 参数): 对 Go 模块(Go Modules)进行检索. 元数据映射为 `Module` 结构体, 默认输出模块路径、版本及替换(replace)信息.

#### 2. 输出格式化工具

* `-f`(模板格式化): 基于 Go 的 `text/template` 语法.

    * 在包模式下, 上下文为 `Package` 结构体. 例如默认格式等价于 `-f '{{.ImportPath}}'`.

    * 在模块模式下, 上下文为 `Module` 结构体. 例如默认格式等价于 `-f '{{.String}}'`.

    * 内置模板函数:

        * `join`: 等同于 `strings.Join`, 用于拼接切片.

        * `context`: 返回当前构建上下文(`Context` 结构体, 含 `GOOS`、`GOARCH`、`GOPATH` 等).

        * `module`(仅限 `-m`): 传入模块路径或查询条件, 返回 `Module` 结构体.

* `-json`(JSON 格式化): 将数据以 JSON 格式输出.

    * 支持按需计算(如 `-json=ImportPath,Dir`, 只输出指定字段), 从而减少工具在收集未请求字段时(如依赖解析)的计算开销.

#### 3. 关键命令行参数(Flags)

* `-deps`: 不仅列出指定的包, 还会以深度优先后序遍历(depth-first post-order)的方式递归列出其所有依赖. 非显式指定的依赖包其 `DepOnly` 字段会被标记为 `true`.

* `-e`: 容错模式. 默认情况下, 遇到无法加载或错误的包时, `go list` 会向标准错误输出错误并跳过. 使用 `-e` 后, 错误不会直接打印到标准错误, 而是填充在 `Package.Error` 字段中正常输出.

* `-find`: 快速查找模式. 仅识别包本身, 不解析其依赖(`Imports` 和 `Deps` 字段为空), 无法与 `-deps`、`-test`、`-export` 共用.

* `-test`: 额外报告测试二进制文件(如 `math/rand.test`)以及为了测试而重新编译的依赖包.

* `-export`: 额外输出包的导出数据文件路径(`Export` 字段, 指向本地缓存的 `.a` 文件)和构建 ID(`BuildID`).

* `-u`(需配合 `-m`): 查询并展示模块的可用更新版本(`Update` 字段)以及撤回(Retracted)信息.

* `-versions`(需配合 `-m`): 列出模块所有可用的版本(按语义化版本排序).

* `-retracted`: 报告模块版本撤回信息(可配合或不配合 `-m` 使用).

* `-reuse=old.json`(需配合 `-m`): 重用先前的 JSON 输出, 避免重新下载相同的模块信息.

---

### 二、实际开发场景与使用示例

#### 场景 1: 获取本地包的基本元数据(路径、源文件等)

在日常开发中, 我们常常需要知道某个包对应的本地目录在哪里, 或者它包含哪些 Go 源文件.

* 示例 1.1: 查询当前目录包的绝对路径和名称

    ```bash
    go list -f 'Path: {{.ImportPath}} | Dir: {{.Dir}} | Name: {{.Name}}' .
    ```

    输出类似于:

    ```text
    Path: myproject/utils | Dir: /home/user/myproject/utils | Name: utils
    ```

* 示例 1.2: 列出某个包中所有的 Go 源文件(不含测试文件)

    由于 `GoFiles` 是个字符串切片, 我们可以使用内置的 `join` 函数将其连接.

    ```bash
    go list -f '{{join .GoFiles ", "}}' fmt
    ```

    输出类似于:

    ```text
    doc.go, format.go, print.go, scan.go
    ```

* 示例 1.3: 将当前包的完整元数据以 JSON 格式输出

    ```bash
    go list -json .
    ```

---

#### 场景 2: 依赖关系分析

`go list` 是分析项目依赖非常有用的工具.

* 示例 2.1: 递归列出当前包依赖的所有包(去重且按拓扑顺序)

    ```bash
    go list -deps .
    ```

* 示例 2.2: 找出当前包依赖了哪些第三方包(排除标准库)

    我们可以利用 `Standard` 字段来过滤掉 Go 标准库.

    ```bash
    go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}' -deps .
    ```

* 示例 2.3: 只查看当前包直接导入(非递归)的包

    直接读取 `Imports` 切片.

    ```bash
    go list -f '{{join .Imports "\n"}}' .
    ```

---

#### 场景 3: 模块管理与版本升级(Go Modules 场景)

在模块化开发中, `-m` 参数可以帮我们快速查询模块仓库.

* 示例 3.1: 列出当前项目的所有活跃模块(主模块及其所有依赖)

    ```bash
    go list -m all
    ```

* 示例 3.2: 检查项目依赖的模块是否有可用的新版本

    使用 `-u` 参数, 如果有更新, 会在括号内 `[...]` 显示最新版本.

    ```bash
    go list -m -u all
    ```

    输出类似于:

    ```text
    my/main/module
    github.com/gin-gonic/gin v1.9.0 [v1.10.0]
    golang.org/x/text v0.14.0
    ```

* 示例 3.3: 查询某个开源库(例如 `rsc.io/pdf`)在远程仓库中所有的可用版本

    ```bash
    go list -m -versions rsc.io/pdf
    ```

    输出类似于:

    ```text
    rsc.io/pdf v0.1.1 v0.1.2
    ```

---

#### 场景 4: 测试与测试依赖分析

编写自动化测试或 CI/CD 脚本时, 有时需要了解测试包的结构.

* 示例 4.1: 列出某个包的测试源文件和外部测试源文件

    `TestGoFiles` 对应包内测试(`package xxx`), `XTestGoFiles` 对应外部测试(`package xxx_test`).

    ```bash
    go list -f 'Test files: {{join .TestGoFiles ", "}} | XTest files: {{join .XTestGoFiles ", "}}' math/rand
    ```

* 示例 4.2: 查看测试二进制文件的依赖关系

    使用 `-test` 会额外列出测试套件本身(`.test`)和其对应的测试依赖.

    ```bash
    go list -test -deps fmt
    ```

---

#### 场景 5: 排查包或依赖的错误

在自动化工具链(如 IDE 插件、静态分析工具)中, 我们可能需要静默读取包信息, 即使有些包存在编译错误.

* 示例 5.1: 容错模式下输出包导入路径和其错误信息

    如果不加 `-e`, 遇到无法导入的包命令会直接报错中断. 加上 `-e` 后, 可以通过模板读取错误:

    ```bash
    go list -e -f '{{if .Error}}Error in {{.ImportPath}}: {{.Error.Err}}{{else}}OK: {{.ImportPath}}{{end}}' ./...
    ```

---

#### 场景 6: 获取构建上下文(Build Context)

* 示例 6.1: 打印当前的构建操作系统(GOOS)和架构(GOARCH)

    通过 `context` 函数可以获取当前的构建环境元数据.

    ```bash
    go list -f '{{with context}}{{.GOOS}}/{{.GOARCH}}{{end}}' .
    ```

    输出类似于:

    ```text
    linux/amd64
    ```

--------------------------------------------------------------------------------------

```bash
$ docker run -ti --rm --name cluster-operator \
-v /Users/huzhi/work/code/go_code/rabbitmq/cluster-operator:/cluster-operator \
-w /cluster-operator \
golang:1.26.4 bash
root@f593cfc7d958:/cluster-operator# git status
Refresh index: 100% (374/374), done.
On branch learn-v2.22.0
nothing to commit, working tree clean
root@f593cfc7d958:/cluster-operator# pwd
/cluster-operator
root@f593cfc7d958:/cluster-operator#

# 将当前包的完整元数据以 JSON 格式输出
root@f593cfc7d958:/cluster-operator# go list -json ./api/v1beta1/
{
	"Dir": "/cluster-operator/api/v1beta1",
	"ImportPath": "github.com/rabbitmq/cluster-operator/v2/api/v1beta1",
	"Name": "v1beta1",
	"Doc": "Package v1beta1 contains API Schema definitions for the rabbitmq v1beta1 API group +kubebuilder:object:generate=true +groupName=rabbitmq.com",
	"Root": "/cluster-operator",
	"Module": {
		"Path": "github.com/rabbitmq/cluster-operator/v2",
		"Main": true,
		"Dir": "/cluster-operator",
		"GoMod": "/cluster-operator/go.mod",
		"GoVersion": "1.26.4"
	},
	"Match": [
		"./api/v1beta1"
	],
	"Stale": true,
	"StaleReason": "stale dependency: internal/goarch",
	"GoFiles": [
		"groupversion_info.go",
		"rabbitmqcluster_status.go",
		"rabbitmqcluster_types.go",
		"zz_generated.deepcopy.go"
	],
	"Imports": [
		"fmt",
		"github.com/rabbitmq/cluster-operator/v2/internal/status",
		"k8s.io/api/apps/v1",
		"k8s.io/api/core/v1",
		"k8s.io/apimachinery/pkg/api/resource",
		"k8s.io/apimachinery/pkg/apis/meta/v1",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"sigs.k8s.io/controller-runtime/pkg/scheme",
		"slices",
		"strconv",
		"strings"
	],
	"Deps": [
		"bufio",
		"bytes",
		"cmp",
		"compress/flate",
		"compress/gzip",
		"container/list",
		"context",
		"crypto",
		"crypto/aes",
		"crypto/cipher",
		"crypto/des",
		"crypto/dsa",
		"crypto/ecdh",
		"crypto/ecdsa",
		"crypto/ed25519",
		"crypto/elliptic",
		"crypto/fips140",
		"crypto/hkdf",
		"crypto/hmac",
		"crypto/hpke",
		"crypto/internal/boring",
		"crypto/internal/boring/bbig",
		"crypto/internal/boring/sig",
		"crypto/internal/constanttime",
		"crypto/internal/entropy/v1.0.0",
		"crypto/internal/fips140",
		"crypto/internal/fips140/aes",
		"crypto/internal/fips140/aes/gcm",
		"crypto/internal/fips140/alias",
		"crypto/internal/fips140/bigmod",
		"crypto/internal/fips140/check",
		"crypto/internal/fips140/drbg",
		"crypto/internal/fips140/ecdh",
		"crypto/internal/fips140/ecdsa",
		"crypto/internal/fips140/ed25519",
		"crypto/internal/fips140/edwards25519",
		"crypto/internal/fips140/edwards25519/field",
		"crypto/internal/fips140/hkdf",
		"crypto/internal/fips140/hmac",
		"crypto/internal/fips140/mlkem",
		"crypto/internal/fips140/nistec",
		"crypto/internal/fips140/nistec/fiat",
		"crypto/internal/fips140/rsa",
		"crypto/internal/fips140/sha256",
		"crypto/internal/fips140/sha3",
		"crypto/internal/fips140/sha512",
		"crypto/internal/fips140/subtle",
		"crypto/internal/fips140/tls12",
		"crypto/internal/fips140/tls13",
		"crypto/internal/fips140cache",
		"crypto/internal/fips140deps/byteorder",
		"crypto/internal/fips140deps/cpu",
		"crypto/internal/fips140deps/godebug",
		"crypto/internal/fips140deps/time",
		"crypto/internal/fips140hash",
		"crypto/internal/fips140only",
		"crypto/internal/impl",
		"crypto/internal/rand",
		"crypto/internal/randutil",
		"crypto/internal/sysrand",
		"crypto/md5",
		"crypto/mlkem",
		"crypto/rand",
		"crypto/rc4",
		"crypto/rsa",
		"crypto/sha1",
		"crypto/sha256",
		"crypto/sha3",
		"crypto/sha512",
		"crypto/subtle",
		"crypto/tls",
		"crypto/tls/internal/fips140tls",
		"crypto/x509",
		"crypto/x509/pkix",
		"encoding",
		"encoding/asn1",
		"encoding/base32",
		"encoding/base64",
		"encoding/binary",
		"encoding/hex",
		"encoding/json",
		"encoding/pem",
		"errors",
		"flag",
		"fmt",
		"github.com/fxamacker/cbor/v2",
		"github.com/go-logr/logr",
		"github.com/json-iterator/go",
		"github.com/modern-go/concurrent",
		"github.com/modern-go/reflect2",
		"github.com/rabbitmq/cluster-operator/v2/internal/status",
		"github.com/x448/float16",
		"go.yaml.in/yaml/v2",
		"go/ast",
		"go/build/constraint",
		"go/doc",
		"go/doc/comment",
		"go/internal/scannerhooks",
		"go/parser",
		"go/scanner",
		"go/token",
		"golang.org/x/net/http/httpguts",
		"golang.org/x/net/http2",
		"golang.org/x/net/http2/hpack",
		"golang.org/x/net/idna",
		"golang.org/x/net/internal/httpcommon",
		"golang.org/x/net/internal/httpsfv",
		"golang.org/x/text/secure/bidirule",
		"golang.org/x/text/transform",
		"golang.org/x/text/unicode/bidi",
		"golang.org/x/text/unicode/norm",
		"gopkg.in/inf.v0",
		"hash",
		"hash/crc32",
		"internal/abi",
		"internal/asan",
		"internal/bisect",
		"internal/bytealg",
		"internal/byteorder",
		"internal/chacha8rand",
		"internal/coverage/rtcov",
		"internal/cpu",
		"internal/filepathlite",
		"internal/fmtsort",
		"internal/goarch",
		"internal/godebug",
		"internal/godebugs",
		"internal/goexperiment",
		"internal/goos",
		"internal/lazyregexp",
		"internal/msan",
		"internal/nettrace",
		"internal/oserror",
		"internal/poll",
		"internal/profilerecord",
		"internal/race",
		"internal/reflectlite",
		"internal/runtime/atomic",
		"internal/runtime/cgroup",
		"internal/runtime/exithook",
		"internal/runtime/gc",
		"internal/runtime/gc/scan",
		"internal/runtime/maps",
		"internal/runtime/math",
		"internal/runtime/pprof/label",
		"internal/runtime/sys",
		"internal/runtime/syscall/linux",
		"internal/saferio",
		"internal/singleflight",
		"internal/strconv",
		"internal/stringslite",
		"internal/sync",
		"internal/synctest",
		"internal/syscall/execenv",
		"internal/syscall/unix",
		"internal/testlog",
		"internal/trace/tracev2",
		"internal/unsafeheader",
		"io",
		"io/fs",
		"io/ioutil",
		"iter",
		"k8s.io/api/apps/v1",
		"k8s.io/api/core/v1",
		"k8s.io/api/discovery/v1",
		"k8s.io/apimachinery/pkg/api/equality",
		"k8s.io/apimachinery/pkg/api/operation",
		"k8s.io/apimachinery/pkg/api/resource",
		"k8s.io/apimachinery/pkg/api/validate/constraints",
		"k8s.io/apimachinery/pkg/api/validate/content",
		"k8s.io/apimachinery/pkg/apis/meta/v1",
		"k8s.io/apimachinery/pkg/conversion",
		"k8s.io/apimachinery/pkg/conversion/queryparams",
		"k8s.io/apimachinery/pkg/fields",
		"k8s.io/apimachinery/pkg/labels",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"k8s.io/apimachinery/pkg/runtime/serializer/cbor/direct",
		"k8s.io/apimachinery/pkg/runtime/serializer/cbor/internal/modes",
		"k8s.io/apimachinery/pkg/selection",
		"k8s.io/apimachinery/pkg/types",
		"k8s.io/apimachinery/pkg/util/errors",
		"k8s.io/apimachinery/pkg/util/intstr",
		"k8s.io/apimachinery/pkg/util/json",
		"k8s.io/apimachinery/pkg/util/naming",
		"k8s.io/apimachinery/pkg/util/net",
		"k8s.io/apimachinery/pkg/util/runtime",
		"k8s.io/apimachinery/pkg/util/sets",
		"k8s.io/apimachinery/pkg/util/validation",
		"k8s.io/apimachinery/pkg/util/validation/field",
		"k8s.io/apimachinery/pkg/watch",
		"k8s.io/apimachinery/third_party/forked/golang/reflect",
		"k8s.io/klog/v2",
		"k8s.io/klog/v2/internal/buffer",
		"k8s.io/klog/v2/internal/clock",
		"k8s.io/klog/v2/internal/dbg",
		"k8s.io/klog/v2/internal/serialize",
		"k8s.io/klog/v2/internal/severity",
		"k8s.io/klog/v2/internal/sloghandler",
		"k8s.io/klog/v2/internal/verbosity",
		"k8s.io/klog/v2/textlogger",
		"k8s.io/kube-openapi/pkg/util",
		"k8s.io/utils/internal/third_party/forked/golang/net",
		"k8s.io/utils/net",
		"k8s.io/utils/ptr",
		"log",
		"log/internal",
		"log/slog",
		"log/slog/internal",
		"log/slog/internal/buffer",
		"maps",
		"math",
		"math/big",
		"math/bits",
		"math/rand",
		"math/rand/v2",
		"mime",
		"mime/multipart",
		"mime/quotedprintable",
		"net",
		"net/http",
		"net/http/httptrace",
		"net/http/internal",
		"net/http/internal/ascii",
		"net/http/internal/httpcommon",
		"net/netip",
		"net/textproto",
		"net/url",
		"os",
		"os/user",
		"path",
		"path/filepath",
		"reflect",
		"regexp",
		"regexp/syntax",
		"runtime",
		"runtime/cgo",
		"runtime/debug",
		"sigs.k8s.io/controller-runtime/pkg/scheme",
		"sigs.k8s.io/json",
		"sigs.k8s.io/json/internal/golang/encoding/json",
		"sigs.k8s.io/randfill",
		"sigs.k8s.io/randfill/bytesource",
		"sigs.k8s.io/structured-merge-diff/v6/value",
		"slices",
		"sort",
		"strconv",
		"strings",
		"sync",
		"sync/atomic",
		"syscall",
		"time",
		"unicode",
		"unicode/utf16",
		"unicode/utf8",
		"unique",
		"unsafe",
		"vendor/golang.org/x/crypto/chacha20",
		"vendor/golang.org/x/crypto/chacha20poly1305",
		"vendor/golang.org/x/crypto/cryptobyte",
		"vendor/golang.org/x/crypto/cryptobyte/asn1",
		"vendor/golang.org/x/crypto/internal/alias",
		"vendor/golang.org/x/crypto/internal/poly1305",
		"vendor/golang.org/x/net/dns/dnsmessage",
		"vendor/golang.org/x/net/http/httpguts",
		"vendor/golang.org/x/net/http/httpproxy",
		"vendor/golang.org/x/net/http2/hpack",
		"vendor/golang.org/x/net/idna",
		"vendor/golang.org/x/sys/cpu",
		"vendor/golang.org/x/text/secure/bidirule",
		"vendor/golang.org/x/text/transform",
		"vendor/golang.org/x/text/unicode/bidi",
		"vendor/golang.org/x/text/unicode/norm",
		"weak"
	],
	"TestGoFiles": [
		"rabbitmqcluster_status_test.go",
		"rabbitmqcluster_types_test.go",
		"suite_test.go"
	],
	"TestImports": [
		"context",
		"fmt",
		"github.com/onsi/ginkgo/v2",
		"github.com/onsi/gomega",
		"github.com/rabbitmq/cluster-operator/v2/internal/status",
		"k8s.io/api/apps/v1",
		"k8s.io/api/core/v1",
		"k8s.io/api/discovery/v1",
		"k8s.io/apimachinery/pkg/api/errors",
		"k8s.io/apimachinery/pkg/api/resource",
		"k8s.io/apimachinery/pkg/apis/meta/v1",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/types",
		"k8s.io/client-go/kubernetes/scheme",
		"k8s.io/client-go/rest",
		"k8s.io/utils/ptr",
		"path/filepath",
		"sigs.k8s.io/controller-runtime/pkg/client",
		"sigs.k8s.io/controller-runtime/pkg/envtest",
		"sigs.k8s.io/controller-runtime/pkg/log",
		"sigs.k8s.io/controller-runtime/pkg/log/zap",
		"testing"
	]
}
root@f593cfc7d958:/cluster-operator#

# 查询当前目录包的绝对路径和名称
root@f593cfc7d958:/cluster-operator# go list -f 'Path: {{.ImportPath}} | Dir: {{.Dir}} | Name: {{.Name}}' ./api/v1beta1/
Path: github.com/rabbitmq/cluster-operator/v2/api/v1beta1 | Dir: /cluster-operator/api/v1beta1 | Name: v1beta1
root@f593cfc7d958:/cluster-operator#

# 列出某个包中所有的 Go 源文件(不含测试文件)
# 由于 `GoFiles` 是个字符串切片, 我们可以使用内置的 `join` 函数将其连接.
root@f593cfc7d958:/cluster-operator# go list -f '{{join .GoFiles ", "}}' ./api/v1beta1/
groupversion_info.go, rabbitmqcluster_status.go, rabbitmqcluster_types.go, zz_generated.deepcopy.go
root@f593cfc7d958:/cluster-operator#

# 列出某个包中所有的 Go 源文件(包含测试文件)
root@f593cfc7d958:/cluster-operator# go list --test -f '{{join .GoFiles ", "}}' ./api/v1beta1/
groupversion_info.go, rabbitmqcluster_status.go, rabbitmqcluster_types.go, zz_generated.deepcopy.go
/root/.cache/go-build/66/66a411c52a57ecb266cb253c70b8922f490f6bee1fe2c3be462111f8d6d8944e-d
groupversion_info.go, rabbitmqcluster_status.go, rabbitmqcluster_types.go, zz_generated.deepcopy.go, rabbitmqcluster_status_test.go, rabbitmqcluster_types_test.go, suite_test.go
root@f593cfc7d958:/cluster-operator#

# 递归列出当前包依赖的所有包(去重且按拓扑顺序)
root@f593cfc7d958:/cluster-operator# go list -deps ./api/v1beta1/
internal/goarch
unsafe
internal/abi
internal/unsafeheader
internal/cpu
internal/bytealg
internal/byteorder
internal/chacha8rand
internal/coverage/rtcov
internal/godebugs
internal/goexperiment
internal/goos
internal/profilerecord
internal/runtime/atomic
internal/runtime/syscall/linux
math/bits
internal/strconv
internal/runtime/cgroup
internal/runtime/exithook
internal/runtime/gc
internal/runtime/sys
internal/runtime/gc/scan
internal/asan
internal/msan
internal/race
internal/runtime/math
internal/runtime/maps
internal/runtime/pprof/label
internal/stringslite
internal/trace/tracev2
runtime
internal/reflectlite
errors
cmp
iter
math
unicode/utf8
strconv
sync/atomic
internal/sync
internal/synctest
sync
unicode
reflect
slices
internal/fmtsort
io
internal/oserror
path
internal/bisect
internal/godebug
syscall
time
io/fs
internal/filepathlite
internal/syscall/unix
internal/poll
internal/syscall/execenv
internal/testlog
os
fmt
bytes
math/rand
strings
math/big
gopkg.in/inf.v0
encoding
encoding/base64
unicode/utf16
encoding/json
encoding/base32
encoding/binary
encoding/hex
github.com/x448/float16
sort
github.com/fxamacker/cbor/v2
sigs.k8s.io/json/internal/golang/encoding/json
sigs.k8s.io/json
k8s.io/apimachinery/pkg/runtime/serializer/cbor/internal/modes
k8s.io/apimachinery/pkg/runtime/serializer/cbor/direct
k8s.io/apimachinery/pkg/api/resource
k8s.io/apimachinery/pkg/api/validate/constraints
regexp/syntax
regexp
k8s.io/apimachinery/pkg/api/validate/content
k8s.io/apimachinery/third_party/forked/golang/reflect
k8s.io/apimachinery/pkg/conversion
k8s.io/apimachinery/pkg/selection
k8s.io/apimachinery/pkg/fields
k8s.io/apimachinery/pkg/util/sets
k8s.io/apimachinery/pkg/util/errors
k8s.io/apimachinery/pkg/util/validation/field
bufio
context
flag
log/internal
log
log/slog/internal
log/slog/internal/buffer
log/slog
github.com/go-logr/logr
k8s.io/klog/v2/internal/severity
k8s.io/klog/v2/internal/buffer
k8s.io/klog/v2/internal/clock
k8s.io/klog/v2/internal/dbg
k8s.io/klog/v2/internal/serialize
k8s.io/klog/v2/internal/sloghandler
runtime/cgo
os/user
path/filepath
k8s.io/klog/v2
vendor/golang.org/x/net/dns/dnsmessage
internal/nettrace
internal/singleflight
weak
unique
net/netip
net
k8s.io/utils/internal/third_party/forked/golang/net
k8s.io/utils/net
k8s.io/apimachinery/pkg/util/validation
k8s.io/apimachinery/pkg/labels
go/token
go/internal/scannerhooks
go/scanner
go/ast
go/doc/comment
internal/lazyregexp
go/doc
go/build/constraint
go/parser
k8s.io/apimachinery/pkg/api/operation
net/url
k8s.io/apimachinery/pkg/conversion/queryparams
k8s.io/apimachinery/pkg/runtime/schema
k8s.io/apimachinery/pkg/util/json
runtime/debug
k8s.io/apimachinery/pkg/util/naming
k8s.io/klog/v2/internal/verbosity
k8s.io/klog/v2/textlogger
compress/flate
hash
hash/crc32
compress/gzip
container/list
crypto
crypto/internal/fips140deps/godebug
crypto/internal/fips140
crypto/internal/fips140/alias
crypto/internal/fips140deps/byteorder
crypto/internal/fips140deps/cpu
crypto/internal/impl
crypto/internal/fips140/sha256
crypto/internal/constanttime
crypto/internal/fips140/subtle
crypto/internal/fips140/sha3
crypto/internal/fips140/sha512
crypto/internal/fips140/hmac
crypto/internal/fips140/check
crypto/internal/fips140/aes
crypto/internal/fips140deps/time
crypto/internal/entropy/v1.0.0
crypto/internal/sysrand
crypto/internal/fips140/drbg
crypto/internal/fips140/aes/gcm
crypto/fips140
crypto/internal/fips140only
crypto/subtle
crypto/cipher
crypto/internal/boring/sig
crypto/internal/boring
math/rand/v2
crypto/internal/randutil
crypto/internal/rand
crypto/rand
crypto/aes
crypto/des
crypto/internal/fips140/nistec/fiat
crypto/internal/fips140/nistec
crypto/internal/fips140/ecdh
crypto/internal/fips140/edwards25519/field
crypto/ecdh
crypto/elliptic
crypto/internal/boring/bbig
crypto/internal/fips140/bigmod
crypto/internal/fips140/ecdsa
crypto/internal/fips140cache
crypto/sha3
crypto/internal/fips140hash
crypto/sha512
internal/saferio
encoding/asn1
vendor/golang.org/x/crypto/cryptobyte/asn1
vendor/golang.org/x/crypto/cryptobyte
crypto/ecdsa
crypto/internal/fips140/edwards25519
crypto/internal/fips140/ed25519
crypto/ed25519
crypto/internal/fips140/hkdf
crypto/hkdf
crypto/hmac
crypto/internal/fips140/mlkem
crypto/mlkem
crypto/sha256
vendor/golang.org/x/crypto/internal/alias
vendor/golang.org/x/crypto/chacha20
vendor/golang.org/x/crypto/internal/poly1305
vendor/golang.org/x/sys/cpu
vendor/golang.org/x/crypto/chacha20poly1305
crypto/hpke
crypto/internal/fips140/tls12
crypto/internal/fips140/tls13
crypto/md5
crypto/rc4
crypto/internal/fips140/rsa
crypto/rsa
crypto/sha1
crypto/tls/internal/fips140tls
crypto/dsa
crypto/x509/pkix
encoding/pem
maps
crypto/x509
crypto/tls
vendor/golang.org/x/text/transform
vendor/golang.org/x/text/unicode/bidi
vendor/golang.org/x/text/secure/bidirule
vendor/golang.org/x/text/unicode/norm
vendor/golang.org/x/net/idna
net/textproto
vendor/golang.org/x/net/http/httpguts
vendor/golang.org/x/net/http/httpproxy
vendor/golang.org/x/net/http2/hpack
mime
mime/quotedprintable
mime/multipart
net/http/httptrace
net/http/internal
net/http/internal/ascii
net/http/internal/httpcommon
net/http
k8s.io/apimachinery/pkg/util/runtime
k8s.io/kube-openapi/pkg/util
io/ioutil
github.com/modern-go/concurrent
github.com/modern-go/reflect2
github.com/json-iterator/go
go.yaml.in/yaml/v2
sigs.k8s.io/structured-merge-diff/v6/value
k8s.io/apimachinery/pkg/runtime
k8s.io/apimachinery/pkg/types
sigs.k8s.io/randfill/bytesource
sigs.k8s.io/randfill
k8s.io/apimachinery/pkg/util/intstr
golang.org/x/text/transform
golang.org/x/text/unicode/bidi
golang.org/x/text/secure/bidirule
golang.org/x/text/unicode/norm
golang.org/x/net/idna
golang.org/x/net/http/httpguts
golang.org/x/net/http2/hpack
golang.org/x/net/internal/httpcommon
golang.org/x/net/internal/httpsfv
golang.org/x/net/http2
k8s.io/apimachinery/pkg/util/net
k8s.io/utils/ptr
k8s.io/apimachinery/pkg/watch
k8s.io/apimachinery/pkg/apis/meta/v1
k8s.io/api/core/v1
k8s.io/api/apps/v1
k8s.io/api/discovery/v1
k8s.io/apimachinery/pkg/api/equality
github.com/rabbitmq/cluster-operator/v2/internal/status
sigs.k8s.io/controller-runtime/pkg/scheme
github.com/rabbitmq/cluster-operator/v2/api/v1beta1
root@f593cfc7d958:/cluster-operator#

# 找出当前包依赖了哪些第三方包(排除标准库)
# 我们可以利用 `Standard` 字段来过滤掉 Go 标准库.
root@f593cfc7d958:/cluster-operator# go list -json -deps ./api/v1beta1/
root@f593cfc7d958:/cluster-operator# go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}' -deps ./api/v1beta1/
gopkg.in/inf.v0
github.com/x448/float16
github.com/fxamacker/cbor/v2
sigs.k8s.io/json/internal/golang/encoding/json
sigs.k8s.io/json
k8s.io/apimachinery/pkg/runtime/serializer/cbor/internal/modes
k8s.io/apimachinery/pkg/runtime/serializer/cbor/direct
k8s.io/apimachinery/pkg/api/resource
k8s.io/apimachinery/pkg/api/validate/constraints
k8s.io/apimachinery/pkg/api/validate/content
k8s.io/apimachinery/third_party/forked/golang/reflect
k8s.io/apimachinery/pkg/conversion
k8s.io/apimachinery/pkg/selection
k8s.io/apimachinery/pkg/fields
k8s.io/apimachinery/pkg/util/sets
k8s.io/apimachinery/pkg/util/errors
k8s.io/apimachinery/pkg/util/validation/field
github.com/go-logr/logr
k8s.io/klog/v2/internal/severity
k8s.io/klog/v2/internal/buffer
k8s.io/klog/v2/internal/clock
k8s.io/klog/v2/internal/dbg
k8s.io/klog/v2/internal/serialize
k8s.io/klog/v2/internal/sloghandler
k8s.io/klog/v2
k8s.io/utils/internal/third_party/forked/golang/net
k8s.io/utils/net
k8s.io/apimachinery/pkg/util/validation
k8s.io/apimachinery/pkg/labels
k8s.io/apimachinery/pkg/api/operation
k8s.io/apimachinery/pkg/conversion/queryparams
k8s.io/apimachinery/pkg/runtime/schema
k8s.io/apimachinery/pkg/util/json
k8s.io/apimachinery/pkg/util/naming
k8s.io/klog/v2/internal/verbosity
k8s.io/klog/v2/textlogger
k8s.io/apimachinery/pkg/util/runtime
k8s.io/kube-openapi/pkg/util
github.com/modern-go/concurrent
github.com/modern-go/reflect2
github.com/json-iterator/go
go.yaml.in/yaml/v2
sigs.k8s.io/structured-merge-diff/v6/value
k8s.io/apimachinery/pkg/runtime
k8s.io/apimachinery/pkg/types
sigs.k8s.io/randfill/bytesource
sigs.k8s.io/randfill
k8s.io/apimachinery/pkg/util/intstr
golang.org/x/text/transform
golang.org/x/text/unicode/bidi
golang.org/x/text/secure/bidirule
golang.org/x/text/unicode/norm
golang.org/x/net/idna
golang.org/x/net/http/httpguts
golang.org/x/net/http2/hpack
golang.org/x/net/internal/httpcommon
golang.org/x/net/internal/httpsfv
golang.org/x/net/http2
k8s.io/apimachinery/pkg/util/net
k8s.io/utils/ptr
k8s.io/apimachinery/pkg/watch
k8s.io/apimachinery/pkg/apis/meta/v1
k8s.io/api/core/v1
k8s.io/api/apps/v1
k8s.io/api/discovery/v1
k8s.io/apimachinery/pkg/api/equality
github.com/rabbitmq/cluster-operator/v2/internal/status
sigs.k8s.io/controller-runtime/pkg/scheme
github.com/rabbitmq/cluster-operator/v2/api/v1beta1
root@f593cfc7d958:/cluster-operator#

# 只查看当前包直接导入(非递归)的包
root@f593cfc7d958:/cluster-operator# go list -f '{{join .Imports "\n"}}' ./api/v1beta1/
fmt
github.com/rabbitmq/cluster-operator/v2/internal/status
k8s.io/api/apps/v1
k8s.io/api/core/v1
k8s.io/apimachinery/pkg/api/resource
k8s.io/apimachinery/pkg/apis/meta/v1
k8s.io/apimachinery/pkg/runtime
k8s.io/apimachinery/pkg/runtime/schema
sigs.k8s.io/controller-runtime/pkg/scheme
slices
strconv
strings
root@f593cfc7d958:/cluster-operator#

# 列出当前项目的所有活跃模块(主模块及其所有依赖)
root@f593cfc7d958:/cluster-operator# go list -m all
github.com/rabbitmq/cluster-operator/v2
bitbucket.org/creachadair/shell v0.0.7
bitbucket.org/liamstask/goose v0.0.0-20150115234039-8488cc47d90c
cel.dev/expr v0.25.2
cloud.google.com/go/compute v1.23.1
cloud.google.com/go/compute/metadata v0.9.0
cloud.google.com/go/monitoring v1.16.1
cloud.google.com/go/trace v1.10.2
contrib.go.opencensus.io/exporter/stackdriver v0.13.14
filippo.io/edwards25519 v1.1.0
github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161
github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.31.0
github.com/Masterminds/semver/v3 v3.5.0
github.com/NYTimes/gziphandler v1.1.1
github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8
github.com/alecthomas/kingpin/v2 v2.4.0
github.com/alecthomas/units v0.0.0-20240927000941-0f3dac36c52b
github.com/antihax/optional v1.0.0
github.com/antlr4-go/antlr/v4 v4.13.1
github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
github.com/aws/aws-sdk-go v1.46.4
github.com/beorn7/perks v1.0.1
github.com/bgentry/speakeasy v0.1.0
github.com/blang/semver/v4 v4.0.0
github.com/bufbuild/protocompile v0.6.0
github.com/bwesterb/go-ristretto v1.2.0
github.com/cenkalti/backoff/v4 v4.2.1
github.com/cenkalti/backoff/v5 v5.0.3
github.com/census-instrumentation/opencensus-proto v0.4.1
github.com/cespare/xxhash/v2 v2.3.0
github.com/chzyer/readline v1.5.1
github.com/cloudflare/backoff v0.0.0-20161212185259-647f3cdfc87a
github.com/cloudflare/cfssl v1.6.5
github.com/cloudflare/circl v1.1.0
github.com/cloudflare/redoctober v0.0.0-20211013234631-6a74ccc611f6
github.com/cncf/udpa/go v0.0.0-20220112060539-c52dc94e7fbe
github.com/cncf/xds/go v0.0.0-20260202195803-dba9d589def2
github.com/coreos/go-oidc v2.5.0+incompatible
github.com/coreos/go-semver v0.3.1
github.com/coreos/go-systemd/v22 v22.7.0
github.com/cpuguy83/go-md2man/v2 v2.0.6
github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
github.com/dustin/go-humanize v1.0.1
github.com/eclipse/paho.mqtt.golang v1.5.1
github.com/emicklei/go-restful/v3 v3.13.0
github.com/envoyproxy/go-control-plane v0.14.0
github.com/envoyproxy/go-control-plane/envoy v1.37.0
github.com/envoyproxy/go-control-plane/ratelimit v0.1.0
github.com/envoyproxy/protoc-gen-validate v1.3.3
github.com/evanphx/json-patch v5.9.11+incompatible
github.com/evanphx/json-patch/v5 v5.9.11
github.com/felixge/httpsnoop v1.0.4
github.com/frankban/quicktest v1.14.6
github.com/fsnotify/fsnotify v1.10.1
github.com/fullstorydev/grpcurl v1.8.9
github.com/fxamacker/cbor/v2 v2.9.2
github.com/getsentry/sentry-go v0.11.0
github.com/gkampitakis/ciinfo v0.3.2
github.com/gkampitakis/go-diff v1.3.2
github.com/gkampitakis/go-snaps v0.5.15
github.com/go-jose/go-jose/v4 v4.1.4
github.com/go-logr/logr v1.4.3
github.com/go-logr/stdr v1.2.2
github.com/go-logr/zapr v1.3.0
github.com/go-openapi/jsonpointer v0.23.1
github.com/go-openapi/jsonreference v0.21.6
github.com/go-openapi/swag v0.26.1
github.com/go-openapi/swag/cmdutils v0.26.1
github.com/go-openapi/swag/conv v0.26.1
github.com/go-openapi/swag/fileutils v0.26.1
github.com/go-openapi/swag/jsonname v0.26.1
github.com/go-openapi/swag/jsonutils v0.26.1
github.com/go-openapi/swag/jsonutils/fixtures_test v0.26.1
github.com/go-openapi/swag/loading v0.26.1
github.com/go-openapi/swag/mangling v0.26.1
github.com/go-openapi/swag/netutils v0.26.1
github.com/go-openapi/swag/stringutils v0.26.1
github.com/go-openapi/swag/typeutils v0.26.1
github.com/go-openapi/swag/yamlutils v0.26.1
github.com/go-openapi/testify/enable/yaml/v2 v2.5.1
github.com/go-openapi/testify/v2 v2.5.1
github.com/go-sql-driver/mysql v1.8.1
github.com/go-stomp/stomp v2.1.4+incompatible
github.com/go-task/slim-sprig/v3 v3.0.0
github.com/goccy/go-yaml v1.18.0
github.com/gogo/protobuf v1.3.2
github.com/golang-jwt/jwt/v4 v4.5.0
github.com/golang-jwt/jwt/v5 v5.3.1
github.com/golang/glog v1.2.5
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
github.com/golang/mock v1.6.0
github.com/golang/protobuf v1.5.4
github.com/golang/snappy v1.0.0
github.com/google/btree v1.1.3
github.com/google/cel-go v0.28.1
github.com/google/certificate-transparency-go v1.1.7
github.com/google/gnostic-models v0.7.1
github.com/google/go-cmp v0.7.0
github.com/google/go-github/v50 v50.2.0
github.com/google/go-querystring v1.1.0
github.com/google/gofuzz v1.2.0
github.com/google/pprof v0.0.0-20260604005048-7023385849c0
github.com/google/s2a-go v0.1.7
github.com/google/trillian v1.5.3
github.com/google/uuid v1.6.0
github.com/googleapis/enterprise-certificate-proxy v0.3.2
github.com/googleapis/gax-go/v2 v2.12.0
github.com/gorilla/mux v1.8.0
github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674
github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.1.0
github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.3
github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
github.com/grpc-ecosystem/grpc-gateway v1.16.0
github.com/grpc-ecosystem/grpc-gateway/v2 v2.29.0
github.com/ianlancetaylor/demangle v0.0.0-20250417193237-f615e6bd150b
github.com/inconshreveable/mousetrap v1.1.0
github.com/jessevdk/go-flags v1.6.1
github.com/jhump/protoreflect v1.15.3
github.com/jmespath/go-jmespath v0.4.0
github.com/jmhodges/clock v1.2.0
github.com/jmoiron/sqlx v1.4.0
github.com/jonboulle/clockwork v0.5.0
github.com/josharian/intern v1.0.0
github.com/joshdk/go-junit v1.0.0
github.com/jpillora/backoff v1.0.0
github.com/json-iterator/go v1.1.12
github.com/julienschmidt/httprouter v1.3.0
github.com/kisielk/sqlstruct v0.0.0-20201105191214-5f3e10d3ab46
github.com/klauspost/compress v1.18.6
github.com/konsorten/go-windows-terminal-sequences v1.0.1
github.com/kr/pretty v0.3.1
github.com/kr/pty v1.1.1
github.com/kr/text v0.2.0
github.com/kylelemons/go-gypsy v1.0.0
github.com/kylelemons/godebug v1.1.0
github.com/letsencrypt/pkcs11key/v4 v4.0.0
github.com/lib/pq v1.10.9
github.com/mailru/easyjson v0.7.7
github.com/maruel/natural v1.1.1
github.com/mattn/go-runewidth v0.0.13
github.com/mattn/go-sqlite3 v1.14.22
github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0
github.com/mfridman/tparse v0.18.0
github.com/michaelklishin/rabbit-hole/v3 v3.5.0
github.com/miekg/pkcs11 v1.1.1
github.com/moby/spdystream v0.5.1
github.com/moby/term v0.5.0
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee
github.com/mreiferson/go-httpclient v0.0.0-20201222173833-5e475fde3a4d
github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f
github.com/olekukonko/tablewriter v0.0.5
github.com/onsi/ginkgo/v2 v2.32.0
github.com/onsi/gomega v1.42.1
github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
github.com/pelletier/go-toml v1.9.5
github.com/peterbourgon/diskv v2.0.1+incompatible
github.com/pierrec/lz4 v2.6.1+incompatible
github.com/pkg/errors v0.9.1
github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2
github.com/pquerna/cachecontrol v0.1.0
github.com/prometheus/client_golang v1.23.2
github.com/prometheus/client_model v0.6.2
github.com/prometheus/common v0.68.1
github.com/prometheus/procfs v0.20.1
github.com/prometheus/prometheus v0.47.2
github.com/rabbitmq/amqp091-go v1.12.0
github.com/rabbitmq/rabbitmq-stream-go-client v1.8.1
github.com/rivo/uniseg v0.4.4
github.com/rogpeppe/fastuuid v1.2.0
github.com/rogpeppe/go-internal v1.14.1
github.com/rs/cors v1.10.1
github.com/russross/blackfriday/v2 v2.1.0
github.com/sergi/go-diff v1.3.1
github.com/sirupsen/logrus v1.9.3
github.com/soheilhy/cmux v0.1.5
github.com/spaolacci/murmur3 v1.1.0
github.com/spf13/cobra v1.10.2
github.com/spf13/pflag v1.0.10
github.com/spiffe/go-spiffe/v2 v2.6.0
github.com/stoewer/go-strcase v1.3.0
github.com/stretchr/objx v0.5.2
github.com/stretchr/testify v1.11.1
github.com/tidwall/gjson v1.18.0
github.com/tidwall/match v1.1.1
github.com/tidwall/pretty v1.2.1
github.com/tidwall/sjson v1.2.5
github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75
github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
github.com/transparency-dev/merkle v0.0.2
github.com/urfave/cli v1.22.14
github.com/weppos/publicsuffix-go v0.30.2-0.20230730094716-a20f9abcc222
github.com/x448/float16 v0.8.4
github.com/xhit/go-str2duration/v2 v2.1.0
github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510
github.com/yuin/goldmark v1.4.13
github.com/ziutek/mymysql v1.5.4
github.com/zmap/rc2 v0.0.0-20190804163417-abaa70531248
github.com/zmap/zcertificate v0.0.1
github.com/zmap/zcrypto v0.0.0-20231219022726-a1f61fb1661c
github.com/zmap/zlint/v3 v3.6.0
go.etcd.io/bbolt v1.4.3
go.etcd.io/etcd/api/v3 v3.6.8
go.etcd.io/etcd/client/pkg/v3 v3.6.8
go.etcd.io/etcd/client/v2 v2.305.10
go.etcd.io/etcd/client/v3 v3.6.8
go.etcd.io/etcd/etcdctl/v3 v3.5.10
go.etcd.io/etcd/etcdutl/v3 v3.5.10
go.etcd.io/etcd/pkg/v3 v3.6.8
go.etcd.io/etcd/raft/v3 v3.5.10
go.etcd.io/etcd/server/v3 v3.6.8
go.etcd.io/etcd/tests/v3 v3.5.10
go.etcd.io/etcd/v3 v3.5.10
go.etcd.io/raft/v3 v3.6.0
go.opencensus.io v0.24.0
go.opentelemetry.io/auto/sdk v1.2.1
go.opentelemetry.io/contrib/detectors/gcp v1.42.0
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.65.0
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.69.0
go.opentelemetry.io/otel v1.44.0
go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.44.0
go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.44.0
go.opentelemetry.io/otel/metric v1.44.0
go.opentelemetry.io/otel/sdk v1.44.0
go.opentelemetry.io/otel/sdk/metric v1.44.0
go.opentelemetry.io/otel/trace v1.44.0
go.opentelemetry.io/proto/otlp v1.10.0
go.uber.org/automaxprocs v1.6.0
go.uber.org/goleak v1.3.0
go.uber.org/multierr v1.11.0
go.uber.org/zap v1.28.0
go.yaml.in/yaml/v2 v2.4.4
go.yaml.in/yaml/v3 v3.0.4
golang.org/x/crypto v0.53.0
golang.org/x/exp v0.0.0-20260603202125-055de637280b
golang.org/x/mod v0.37.0
golang.org/x/net v0.56.0
golang.org/x/oauth2 v0.36.0
golang.org/x/sync v0.21.0
golang.org/x/sys v0.46.0
golang.org/x/telemetry v0.0.0-20260508192327-42602be52be6
golang.org/x/term v0.44.0
golang.org/x/text v0.38.0
golang.org/x/time v0.15.0
golang.org/x/tools v0.45.0
golang.org/x/tools/go/expect v0.1.1-deprecated
golang.org/x/tools/go/packages/packagestest v0.1.1-deprecated
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
gomodules.xyz/jsonpatch/v2 v2.5.0
gonum.org/v1/gonum v0.17.0
google.golang.org/api v0.149.0
google.golang.org/appengine v1.6.8
google.golang.org/genproto v0.0.0-20231016165738-49dd2c1f3d0b
google.golang.org/genproto/googleapis/api v0.0.0-20260526163538-3dc84a4a5aaa
google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa
google.golang.org/grpc v1.81.1
google.golang.org/protobuf v1.36.12-0.20260120151049-f2248ac996af
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
gopkg.in/cheggaaa/pb.v1 v1.0.28
gopkg.in/evanphx/json-patch.v4 v4.13.0
gopkg.in/go-jose/go-jose.v2 v2.6.3
gopkg.in/inf.v0 v0.9.1
gopkg.in/ini.v1 v1.67.3
gopkg.in/natefinch/lumberjack.v2 v2.2.1
gopkg.in/yaml.v2 v2.2.2
gopkg.in/yaml.v3 v3.0.1
k8s.io/api v0.36.2
k8s.io/apiextensions-apiserver v0.36.1
k8s.io/apimachinery v0.36.2
k8s.io/apiserver v0.36.1
k8s.io/client-go v0.36.2
k8s.io/code-generator v0.36.1
k8s.io/component-base v0.36.1
k8s.io/gengo/v2 v2.0.0-20250922181213-ec3ebc5fd46b
k8s.io/klog/v2 v2.140.0
k8s.io/kms v0.36.1
k8s.io/kube-openapi v0.0.0-20260603220949-865597e52e25
k8s.io/streaming v0.36.2
k8s.io/utils v0.0.0-20260626114624-be93311217bd
sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.36.0
sigs.k8s.io/controller-runtime v0.24.1
sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730
sigs.k8s.io/randfill v1.0.0
sigs.k8s.io/structured-merge-diff/v6 v6.4.0
sigs.k8s.io/yaml v1.6.0
root@f593cfc7d958:/cluster-operator#

root@f593cfc7d958:/cluster-operator# go list -json -m all

# 检查项目依赖的模块是否有可用的新版本
root@f593cfc7d958:/cluster-operator# go list -m -u sigs.k8s.io/yaml
sigs.k8s.io/yaml v1.6.0
root@f593cfc7d958:/cluster-operator#
root@f593cfc7d958:/cluster-operator# go list -m -u cloud.google.com/go/compute
cloud.google.com/go/compute v1.23.1 [v1.64.0]
root@f593cfc7d958:/cluster-operator#
root@f593cfc7d958:/cluster-operator# go list -m -u all

# 查询某个开源库(例如 `cloud.google.com/go/compute`)在远程仓库中所有的可用版本
root@f593cfc7d958:/cluster-operator# go list -m -versions cloud.google.com/go/compute
cloud.google.com/go/compute v0.1.0 v1.0.0 v1.1.0 v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.7.0 v1.8.0 v1.9.0 v1.10.0 v1.11.0 v1.12.0 v1.12.1 v1.13.0 v1.14.0 v1.15.0 v1.15.1 v1.18.0 v1.19.0 v1.19.1 v1.19.2 v1.19.3 v1.20.0 v1.20.1 v1.21.0 v1.22.0 v1.23.0 v1.23.1 v1.23.2 v1.23.3 v1.23.4 v1.24.0 v1.25.0 v1.25.1 v1.26.0 v1.27.0 v1.27.1 v1.27.2 v1.27.3 v1.27.4 v1.27.5 v1.28.0 v1.28.1 v1.28.2 v1.28.3 v1.29.0 v1.30.0 v1.31.0 v1.31.1 v1.32.0 v1.33.0 v1.34.0 v1.34.1 v1.35.0 v1.36.0 v1.36.1 v1.37.0 v1.38.0 v1.39.0 v1.40.0 v1.41.0 v1.42.0 v1.43.0 v1.44.0 v1.45.0 v1.46.0 v1.47.0 v1.48.0 v1.49.0 v1.49.1 v1.50.0 v1.51.0 v1.52.0 v1.53.0 v1.54.0 v1.55.0 v1.56.0 v1.57.0 v1.58.0 v1.59.0 v1.60.0 v1.61.0 v1.62.0 v1.63.0 v1.64.0
root@f593cfc7d958:/cluster-operator#

# 列出某个包的测试源文件和外部测试源文件
root@f593cfc7d958:/cluster-operator# go list ./api/v1beta1/
github.com/rabbitmq/cluster-operator/v2/api/v1beta1
root@f593cfc7d958:/cluster-operator# go list -json -test ./api/v1beta1/
root@f593cfc7d958:/cluster-operator# go list -f 'Test files: {{join .TestGoFiles ", "}} | XTest files: {{join .XTestGoFiles ", "}}' ./api/v1beta1/
Test files: rabbitmqcluster_status_test.go, rabbitmqcluster_types_test.go, suite_test.go | XTest files:
root@f593cfc7d958:/cluster-operator#

# 打印当前的构建操作系统(GOOS)和架构(GOARCH)
root@f593cfc7d958:/cluster-operator# go list ./api/v1beta1/
github.com/rabbitmq/cluster-operator/v2/api/v1beta1
root@f593cfc7d958:/cluster-operator#
root@f593cfc7d958:/cluster-operator# go list -f '{{with context}}{{.GOOS}}/{{.GOARCH}}{{end}}' ./api/v1beta1/
linux/amd64
root@f593cfc7d958:/cluster-operator#

```

--------------------------------------------------------------------------------------

Q: 假设你是一位精通 golang 语言的高级开发专家, 我使用命令 `go help list` 查看了 go list 的帮助文档, 内容如下
```
root@82157ef140f0:/go# go help list
usage: go list [-f format] [-json] [-m] [list flags] [build flags] [packages]

List lists the named packages, one per line.
The most commonly-used flags are -f and -json, which control the form
of the output printed for each package. Other list flags, documented below,
control more specific details.

The default output shows the package import path:

    bytes
    encoding/json
    github.com/gorilla/mux
    golang.org/x/net/html

The -f flag specifies an alternate format for the list, using the
syntax of package template. The default output is equivalent
to -f '{{.ImportPath}}'. The struct being passed to the template is:

    type Package struct {
        Dir            string   // directory containing package sources
        ImportPath     string   // import path of package in dir
        ImportComment  string   // path in import comment on package statement
        Name           string   // package name
        Doc            string   // package documentation string
        Target         string   // install path
        Shlib          string   // the shared library that contains this package (only set when -linkshared)
        Goroot         bool     // is this package in the Go root?
        Standard       bool     // is this package part of the standard Go library?
        Stale          bool     // would 'go install' do anything for this package?
        StaleReason    string   // explanation for Stale==true
        Root           string   // Go root or Go path dir containing this package
        ConflictDir    string   // this directory shadows Dir in $GOPATH
        BinaryOnly     bool     // binary-only package (no longer supported)
        ForTest        string   // package is only for use in named test
        Export         string   // file containing export data (when using -export)
        BuildID        string   // build ID of the compiled package (when using -export)
        Module         *Module  // info about package's containing module, if any (can be nil)
        Match          []string // command-line patterns matching this package
        DepOnly        bool     // package is only a dependency, not explicitly listed
        DefaultGODEBUG string  // default GODEBUG setting, for main packages

        // Source files
        GoFiles           []string   // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
        CgoFiles          []string   // .go source files that import "C"
        CompiledGoFiles   []string   // .go files presented to compiler (when using -compiled)
        IgnoredGoFiles    []string   // .go source files ignored due to build constraints
        IgnoredOtherFiles []string // non-.go source files ignored due to build constraints
        CFiles            []string   // .c source files
        CXXFiles          []string   // .cc, .cxx and .cpp source files
        MFiles            []string   // .m source files
        HFiles            []string   // .h, .hh, .hpp and .hxx source files
        FFiles            []string   // .f, .F, .for and .f90 Fortran source files
        SFiles            []string   // .s source files
        SwigFiles         []string   // .swig files
        SwigCXXFiles      []string   // .swigcxx files
        SysoFiles         []string   // .syso object files to add to archive
        TestGoFiles       []string   // _test.go files in package
        XTestGoFiles      []string   // _test.go files outside package

        // Embedded files
        EmbedPatterns      []string // //go:embed patterns
        EmbedFiles         []string // files matched by EmbedPatterns
        TestEmbedPatterns  []string // //go:embed patterns in TestGoFiles
        TestEmbedFiles     []string // files matched by TestEmbedPatterns
        XTestEmbedPatterns []string // //go:embed patterns in XTestGoFiles
        XTestEmbedFiles    []string // files matched by XTestEmbedPatterns

        // Cgo directives
        CgoCFLAGS    []string // cgo: flags for C compiler
        CgoCPPFLAGS  []string // cgo: flags for C preprocessor
        CgoCXXFLAGS  []string // cgo: flags for C++ compiler
        CgoFFLAGS    []string // cgo: flags for Fortran compiler
        CgoLDFLAGS   []string // cgo: flags for linker
        CgoPkgConfig []string // cgo: pkg-config names

        // Dependency information
        Imports      []string          // import paths used by this package
        ImportMap    map[string]string // map from source import to ImportPath (identity entries omitted)
        Deps         []string          // all (recursively) imported dependencies
        TestImports  []string          // imports from TestGoFiles
        XTestImports []string          // imports from XTestGoFiles

        // Error information
        Incomplete bool            // this package or a dependency has an error
        Error      *PackageError   // error loading package
        DepsErrors []*PackageError // errors loading dependencies
    }

Packages stored in vendor directories report an ImportPath that includes the
path to the vendor directory (for example, "d/vendor/p" instead of "p"),
so that the ImportPath uniquely identifies a given copy of a package.
The Imports, Deps, TestImports, and XTestImports lists also contain these
expanded import paths. See golang.org/s/go15vendor for more about vendoring.

The error information, if any, is

    type PackageError struct {
        ImportStack   []string // shortest path from package named on command line to this one
        Pos           string   // position of error (if present, file:line:col)
        Err           string   // the error itself
    }

The module information is a Module struct, defined in the discussion
of list -m below.

The template function "join" calls strings.Join.

The template function "context" returns the build context, defined as:

    type Context struct {
        GOARCH        string   // target architecture
        GOOS          string   // target operating system
        GOROOT        string   // Go root
        GOPATH        string   // Go path
        CgoEnabled    bool     // whether cgo can be used
        UseAllFiles   bool     // use files regardless of //go:build lines, file names
        Compiler      string   // compiler to assume when computing target paths
        BuildTags     []string // build constraints to match in //go:build lines
        ToolTags      []string // toolchain-specific build constraints
        ReleaseTags   []string // releases the current release is compatible with
        InstallSuffix string   // suffix to use in the name of the install dir
    }

For more information about the meaning of these fields see the documentation
for the go/build package's Context type.

The -json flag causes the package data to be printed in JSON format
instead of using the template format. The JSON flag can optionally be
provided with a set of comma-separated required field names to be output.
If so, those required fields will always appear in JSON output, but
others may be omitted to save work in computing the JSON struct.

The -compiled flag causes list to set CompiledGoFiles to the Go source
files presented to the compiler. Typically this means that it repeats
the files listed in GoFiles and then also adds the Go code generated
by processing CgoFiles and SwigFiles. The Imports list contains the
union of all imports from both GoFiles and CompiledGoFiles.

The -deps flag causes list to iterate over not just the named packages
but also all their dependencies. It visits them in a depth-first post-order
traversal, so that a package is listed only after all its dependencies.
Packages not explicitly listed on the command line will have the DepOnly
field set to true.

The -e flag changes the handling of erroneous packages, those that
cannot be found or are malformed. By default, the list command
prints an error to standard error for each erroneous package and
omits the packages from consideration during the usual printing.
With the -e flag, the list command never prints errors to standard
error and instead processes the erroneous packages with the usual
printing. Erroneous packages will have a non-empty ImportPath and
a non-nil Error field; other information may or may not be missing
(zeroed).

The -export flag causes list to set the Export field to the name of a
file containing up-to-date export information for the given package,
and the BuildID field to the build ID of the compiled package.

The -find flag causes list to identify the named packages but not
resolve their dependencies: the Imports and Deps lists will be empty.
With the -find flag, the -deps, -test and -export commands cannot be
used.

The -test flag causes list to report not only the named packages
but also their test binaries (for packages with tests), to convey to
source code analysis tools exactly how test binaries are constructed.
The reported import path for a test binary is the import path of
the package followed by a ".test" suffix, as in "math/rand.test".
When building a test, it is sometimes necessary to rebuild certain
dependencies specially for that test (most commonly the tested
package itself). The reported import path of a package recompiled
for a particular test binary is followed by a space and the name of
the test binary in brackets, as in "math/rand [math/rand.test]"
or "regexp [sort.test]". The ForTest field is also set to the name
of the package being tested ("math/rand" or "sort" in the previous
examples).

The Dir, Target, Shlib, Root, ConflictDir, and Export file paths
are all absolute paths.

By default, the lists GoFiles, CgoFiles, and so on hold names of files in Dir
(that is, paths relative to Dir, not absolute paths).
The generated files added when using the -compiled and -test flags
are absolute paths referring to cached copies of generated Go source files.
Although they are Go source files, the paths may not end in ".go".

The -m flag causes list to list modules instead of packages.

When listing modules, the -f flag still specifies a format template
applied to a Go struct, but now a Module struct:

    type Module struct {
        Path       string        // module path
        Query      string        // version query corresponding to this version
        Version    string        // module version
        Versions   []string      // available module versions
        Replace    *Module       // replaced by this module
        Time       *time.Time    // time version was created
        Update     *Module       // available update (with -u)
        Main       bool          // is this the main module?
        Indirect   bool          // module is only indirectly needed by main module
        Dir        string        // directory holding local copy of files, if any
        GoMod      string        // path to go.mod file describing module, if any
        GoVersion  string        // go version used in module
        Retracted  []string      // retraction information, if any (with -retracted or -u)
        Deprecated string        // deprecation message, if any (with -u)
        Error      *ModuleError  // error loading module
        Sum        string        // checksum for path, version (as in go.sum)
        GoModSum   string        // checksum for go.mod (as in go.sum)
        Origin     any           // provenance of module
        Reuse      bool          // reuse of old module info is safe
    }

    type ModuleError struct {
        Err string // the error itself
    }

The file GoMod refers to may be outside the module directory if the
module is in the module cache or if the -modfile flag is used.

The default output is to print the module path and then
information about the version and replacement if any.
For example, 'go list -m all' might print:

    my/main/module
    golang.org/x/text v0.3.0 => /tmp/text
    rsc.io/pdf v0.1.1

The Module struct has a String method that formats this
line of output, so that the default format is equivalent
to -f '{{.String}}'.

Note that when a module has been replaced, its Replace field
describes the replacement module, and its Dir field is set to
the replacement's source code, if present. (That is, if Replace
is non-nil, then Dir is set to Replace.Dir, with no access to
the replaced source code.)

The -u flag adds information about available upgrades.
When the latest version of a given module is newer than
the current one, list -u sets the Module's Update field
to information about the newer module. list -u will also set
the module's Retracted field if the current version is retracted.
The Module's String method indicates an available upgrade by
formatting the newer version in brackets after the current version.
If a version is retracted, the string "(retracted)" will follow it.
For example, 'go list -m -u all' might print:

    my/main/module
    golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
    rsc.io/pdf v0.1.1 (retracted) [v0.1.2]

(For tools, 'go list -m -u -json all' may be more convenient to parse.)

The -versions flag causes list to set the Module's Versions field
to a list of all known versions of that module, ordered according
to semantic versioning, earliest to latest. The flag also changes
the default output format to display the module path followed by the
space-separated version list.

The -retracted flag causes list to report information about retracted
module versions. When -retracted is used with -f or -json, the Retracted
field explains why the version was retracted.
The strings are taken from comments on the retract directive in the
module's go.mod file. When -retracted is used with -versions, retracted
versions are listed together with unretracted versions. The -retracted
flag may be used with or without -m.

The arguments to list -m are interpreted as a list of modules, not packages.
The main module is the module containing the current directory.
The active modules are the main module and its dependencies.
With no arguments, list -m shows the main module.
With arguments, list -m shows the modules specified by the arguments.
Any of the active modules can be specified by its module path.
The special pattern "all" specifies all the active modules, first the main
module and then dependencies sorted by module path.
A pattern containing "..." specifies the active modules whose
module paths match the pattern.
A query of the form path@version specifies the result of that query,
which is not limited to active modules.
See 'go help modules' for more about module queries.

The template function "module" takes a single string argument
that must be a module path or query and returns the specified
module as a Module struct. If an error occurs, the result will
be a Module struct with a non-nil Error field.

When using -m, the -reuse=old.json flag accepts the name of file containing
the JSON output of a previous 'go list -m -json' invocation with the
same set of modifier flags (such as -u, -retracted, and -versions).
The go command may use this file to determine that a module is unchanged
since the previous invocation and avoid redownloading information about it.
Modules that are not redownloaded will be marked in the new output by
setting the Reuse field to true. Normally the module cache provides this
kind of reuse automatically; the -reuse flag can be useful on systems that
do not preserve the module cache.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.

For more about modules, see https://golang.org/ref/mod.
root@82157ef140f0:/go#
```
请帮我将上述帮助文档翻译为中文

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------

--------------------------------------------------------------------------------------
