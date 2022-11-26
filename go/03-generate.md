* https://go.dev/blog/generate

```bash
$ go help generate
usage: go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

Generate runs commands described by directives within existing
files. Those commands can run any process but the intent is to
create or update Go source files.
Generate运行由现有文件中的指令描述的命令。这些命令可以运行任何进程，但目的是创建或更新Go源文件。

Go generate is never run automatically by go build, go test,
and so on. It must be run explicitly.
Go generate从来不会通过Go构建、Go测试等自动运行。它必须显式地运行。

Go generate scans the file for directives, which are lines of
the form,

	//go:generate command argument...

(note: no leading spaces and no space in "//go") where command
is the generator to be run, corresponding to an executable file
that can be run locally. It must either be in the shell path
(gofmt), a fully qualified path (/usr/you/bin/mytool), or a
command alias, described below.
(注意:前导空格和"//go"中没有空格)where command是要运行的生成器，对
应于可在本地运行的可执行文件。它必须位于shell路径(gofmt)、完全限定的
路径(/usr/you/bin/mytool)或命令别名中，如下所述。

Note that go generate does not parse the file, so lines that look
like directives in comments or multiline strings will be treated
as directives.
注意，go generate不会解析文件，因此在注释或多行字符串中看起来像指令的行将被视为指令。

The arguments to the directive are space-separated tokens or
double-quoted strings passed to the generator as individual
arguments when it is run.
指令的参数是空格分隔的标记或双引号字符串，在运行指令时作为单独的参数传递给生成器。

Quoted strings use Go syntax and are evaluated before execution; a
quoted string appears as a single argument to the generator.
加引号的字符串使用Go语法，并在执行前求值;加引号的字符串作为生成器的单个参数出现。

To convey to humans and machine tools that code is generated,
generated source should have a line that matches the following
regular expression (in Go syntax):
为了向人类和机器工具传递生成的代码，生成的源代码应该有一行匹配以下正则表达式(在Go语法中):

	^// Code generated .* DO NOT EDIT\.$

This line must appear before the first non-comment, non-blank
text in the file.

Go generate sets several variables when it runs the generator:

	$GOARCH
		The execution architecture (arm, amd64, etc.)
	$GOOS
		The execution operating system (linux, windows, etc.)
	$GOFILE
		The base name of the file.
	$GOLINE
		The line number of the directive in the source file.
	$GOPACKAGE
		The name of the package of the file containing the directive.
	$DOLLAR
		A dollar sign.

Other than variable substitution and quoted-string evaluation, no
special processing such as "globbing" is performed on the command
line.
除了变量替换和引用字符串求值外，没有在命令行上执行诸如“globbing”之类的特殊处理。

As a last step before running the command, any invocations of any
environment variables with alphanumeric names, such as $GOFILE or
$HOME, are expanded throughout the command line. The syntax for
variable expansion is $NAME on all operating systems. Due to the
order of evaluation, variables are expanded even inside quoted
strings. If the variable NAME is not set, $NAME expands to the
empty string.
作为运行该命令之前的最后一步，对具有字母数字名称的任何环境变量(如$GOFILE或
$HOME)的任何调用都将在整个命令行中展开。在所有操作系统上，变量展开的语法都是
$NAME。由于求值的顺序，变量甚至在带引号的字符串中也会展开。如果没有设置变量
NAME，则$NAME展开为空字符串。

A directive of the form,

	//go:generate -command xxx args...

specifies, for the remainder of this source file only, that the
string xxx represents the command identified by the arguments. This
can be used to create aliases or to handle multiword generators.
For example,
仅对于此源文件的其余部分，指定字符串XXX表示由参数标识的命令。这可用于创建别名或
处理多字生成器。例如,

	//go:generate -command foo go tool foo

specifies that the command "foo" represents the generator
"go tool foo".

Generate processes packages in the order given on the command line,
one at a time. If the command line lists .go files from a single directory,
they are treated as a single package. Within a package, generate processes the
source files in a package in file name order, one at a time. Within
a source file, generate runs generators in the order they appear
in the file, one at a time. The go generate tool also sets the build
tag "generate" so that files may be examined by go generate but ignored
during build.

For packages with invalid code, generate processes only source files with a
valid package clause.

If any generator returns an error exit status, "go generate" skips
all further processing for that package.

The generator is run in the package\'s source directory.

Go generate accepts one specific flag:

	-run=""
		if non-empty, specifies a regular expression to select
		directives whose full original source text (excluding
		any trailing spaces and final newline) matches the
		expression.

It also accepts the standard build flags including -v, -n, and -x.
The -v flag prints the names of packages and files as they are
processed.
The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.

```
