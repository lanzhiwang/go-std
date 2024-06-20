```bash
$ go help mod
```

Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.
请注意，所有 go 命令都内置了对模块的支持，而不仅仅是“go mod”。
例如，日常添加、删除、升级和降级依赖项应使用“go get”完成。
有关模块功能的概述，请参阅“go help modules”。

Usage:

	go mod <command> [arguments]

The commands are:

	download    download modules to local cache
	edit        edit go.mod from tools or scripts
	graph       print module requirement graph
	init        initialize new module in current directory
	tidy        add missing and remove unused modules
	vendor      make vendored copy of dependencies
	verify      verify dependencies have expected content
	why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.

```bash
$ go help mod init
```

usage: go mod init [module-path]

Init initializes and writes a new go.mod file in the current directory, in
effect creating a new module rooted at the current directory. The go.mod file
must not already exist.
Init 初始化并在当前目录中写入一个新的 go.mod 文件，实际上创建一个以当前目录为根的新模块。 go.mod 文件必须不存在。

Init accepts one optional argument, the module path for the new module. If the
module path argument is omitted, init will attempt to infer the module path
using import comments in .go files, vendoring tool configuration files (like
Gopkg.lock), and the current directory (if in GOPATH).
Init 接受一个可选参数，即新模块的模块路径。
如果省略模块路径参数，init 将尝试使用 .go 文件中的导入注释、供应商工具配置文件（如 Gopkg.lock）和当前目录（如果在 GOPATH 中）推断模块路径。

See https://golang.org/ref/mod#go-mod-init for more about 'go mod init'.

```bash
$ go help mod why
```

usage: go mod why [-m] [-vendor] packages...

Why shows a shortest path in the import graph from the main module to
each of the listed packages. If the -m flag is given, why treats the
arguments as a list of modules and finds a path to any package in each
of the modules.

By default, why queries the graph of packages matched by "go list all",
which includes tests for reachable packages. The -vendor flag causes why
to exclude tests of dependencies.

The output is a sequence of stanzas, one for each package or module
name on the command line, separated by blank lines. Each stanza begins
with a comment line "# package" or "# module" giving the target
package or module. Subsequent lines give a path through the import
graph, one package per line. If the package or module is not
referenced from the main module, the stanza will display a single
parenthesized note indicating that fact.

For example:

	$ go mod why golang.org/x/text/language golang.org/x/text/encoding
	# golang.org/x/text/language
	rsc.io/quote
	rsc.io/sampler
	golang.org/x/text/language

	# golang.org/x/text/encoding
	(main module does not need package golang.org/x/text/encoding)
	$

See https://golang.org/ref/mod#go-mod-why for more about 'go mod why'.

