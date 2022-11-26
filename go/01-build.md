```bash
$ go help build
usage: go build [-o output] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.
Build编译按导入路径命名的包，
以及它们的依赖项，但它不安装结果。

$ go help packages
$ go build -o bin/manager github.com/apache/rocketmq-operator
$ go build github.com/apache/rocketmq-operator/api/v1alpha1

If the arguments to build are a list of .go files from a single directory,
build treats them as a list of source files specifying a single package.
如果要构建的参数是来自单个目录的.go文件列表，
Build将它们视为指定单个包的源文件列表。

$ go build main.go
$ go build main.go assets.go

When compiling packages, build ignores files that end in '_test.go'.

When compiling a single main package, build writes
the resulting executable to an output file named after
the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe').
The '.exe' suffix is added when writing a Windows executable.

When compiling multiple packages or a single non-main package,
build compiles the packages but discards the resulting object,
serving only as a check that the packages can be built.
当编译多个包或单个非主包时，
Build编译包，但丢弃结果对象，
仅作为可以构建包的检查。

The -o flag forces build to write the resulting executable or object
to the named output file or directory, instead of the default behavior described
in the last two paragraphs. If the named output is an existing directory or
ends with a slash or backslash, then any resulting executables
will be written to that directory.
-o标志强制build将生成的可执行文件或对象写入指定的输出文件或目录，
而不是前两段中描述的默认行为。如果指定的输出是一个现有的目录，
或者以斜杠或反斜杠结尾，那么任何结果的可执行程序都将写入该目录。

The -i flag installs the packages that are dependencies of the target.
The -i flag is deprecated. Compiled packages are cached automatically.
-i标志安装目标的依赖包。-i标志已弃用。编译后的包会自动缓存。

The build flags are shared by the build, clean, get, install, list, run,
and test commands:

	-a
		force rebuilding of packages that are already up-to-date.  强制重新构建已经是最新的包。

	---------------------------------------------------------------
	-n
		print the commands but do not run them.

	-v
		print the names of packages as they are compiled.

	-work
		print the name of the temporary work directory and
		do not delete it when exiting.

	-x
		print the commands.

	---------------------------------------------------------------

	-p n
		the number of programs, such as build commands or
		test binaries, that can be run in parallel.
		The default is GOMAXPROCS, normally the number of CPUs available.  可以并行运行的程序的数量，例如构建命令或测试二进制文件。默认值是GOMAXPROCS，通常是可用的cpu数量。

	---------------------------------------------------------------

	-race
		enable data race detection.
		Supported only on linux/amd64, freebsd/amd64, darwin/amd64, darwin/arm64, windows/amd64,
		linux/ppc64le and linux/arm64 (only for 48-bit VMA).
		启用数据竞态检测。
	$ https://zhuanlan.zhihu.com/p/78655582

	---------------------------------------------------------------

	-tags tag,list
		a comma-separated list of build tags to consider satisfied during the
		build. For more information about build tags, see the description of
		build constraints in the documentation for the go/build package.
		(Earlier versions of Go used a space-separated list, and that form
		is deprecated but still recognized.)  以逗号分隔的构建标记列
		表，以在构建期间考虑是否满足。有关构建标记的更多信息，请参见go/build
		包文档中对构建约束的描述。(早期版本的围棋使用空格分隔的列表，这种形式
		已被弃用，但仍可识别。)
	$ https://zhuanlan.zhihu.com/p/269746831
	$ https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags
	$ go doc go/build
	$ go help buildconstraint

	---------------------------------------------------------------

	-buildmode mode
		build mode to use. See 'go help buildmode' for more.

	-linkshared
		build code that will be linked against shared libraries previously
		created with -buildmode=shared.
	$ https://www.cnblogs.com/bergus/articles/go-plugin.html
	$ go build -buildmode=c-archive add.go
	$ go build -buildmode=c-shared -o add.so add.go
	$ go build -buildmode=shared std && go build -linkshared hello.go
	$ go build -buildmode=plugin -o greeter.so greeter.go

	---------------------------------------------------------------

	-asmflags '[pattern=]arg list'
		arguments to pass on each go tool asm invocation.
		传入每个go工具asm调用的参数。
		go tool asm [flags] file
	$ go doc cmd/asm

	-gccgoflags '[pattern=]arg list'
		arguments to pass on each gccgo compiler/linker invocation.
	$ https://blog.csdn.net/fengshenyun/article/details/97372589

	-gcflags '[pattern=]arg list'
		arguments to pass on each go tool compile invocation.
	$ https://www.bwangel.me/2022/01/12/go_gcflags/
	$ go doc cmd/compile

	-ldflags '[pattern=]arg list'
		arguments to pass on each go tool link invocation.
	$ https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
	$ go doc cmd/link

	---------------------------------------------------------------

	-msan
		enable interoperation with memory sanitizer.
		Supported only on linux/amd64, linux/arm64
		and only with Clang/LLVM as the host C compiler.
		On linux/arm64, pie build mode will be used.
		启用与内存杀毒器的互操作。

	-asan
		enable interoperation with address sanitizer.
		Supported only on linux/arm64, linux/amd64.
		启用与地址杀毒程序的互操作。

	-buildvcs
		Whether to stamp binaries with version control information. By default,
		version control information is stamped into a binary if the main package
		and the main module containing it are in the repository containing the
		current directory (if there is a repository). Use -buildvcs=false to
		omit version control information.
		是否用版本控制信息戳记二进制文件。默认情况下，如果主包和包含它的主模块
		位于包含当前目录的存储库中(如果存在存储库)，则版本控制信息将被戳到二进
		制文件中。使用-buildvcs=false可以省略版本控制信息。

	-compiler name
		name of compiler to use, as in runtime.Compiler (gccgo or gc).

	-installsuffix suffix
		a suffix to use in the name of the package installation directory,
		in order to keep output separate from default builds.
		If using the -race flag, the install suffix is automatically set to race
		or, if set explicitly, has _race appended to it. Likewise for the -msan
		and -asan flags. Using a -buildmode option that requires non-default compile
		flags has a similar effect.
		在包安装目录的名称中使用的后缀，以便将输出与默认构建分开。如果使用
		-race标志，则安装后缀将自动设置为race，如果显式设置，则添加_race。
		-msan和-asan标志也是如此。使用需要非默认编译标志的-buildmode选项也
		有类似的效果。

	-mod mode
		module download mode to use: readonly, vendor, or mod.
		By default, if a vendor directory is present and the go version in go.mod
		is 1.14 or higher, the go command acts as if -mod=vendor were set.
		Otherwise, the go command acts as if -mod=readonly were set.
		See https://golang.org/ref/mod#build-commands for details.
		默认情况下，如果存在供应商目录且go版本在go. modules中，则使用
		readonly、vendor或mod。Mod为1.14或更高，go命令就像设置了
		-mod=vendor一样。否则，go命令就像设置了-mod=readonly一样。详情请
		参见https://golang.org/ref/mod#build-commands。

	-modcacherw
		leave newly-created directories in the module cache read-write
		instead of making them read-only.
		在模块缓存中保持新创建的目录为可读写，而不是使其为只读。

	-modfile file
		in module aware mode, read (and possibly write) an alternate go.mod
		file instead of the one in the module root directory. A file named
		"go.mod" must still be present in order to determine the module root
		directory, but it is not accessed. When -modfile is specified, an
		alternate go.sum file is also used: its path is derived from the
		-modfile flag by trimming the ".mod" extension and appending ".sum".

	-overlay file
		read a JSON config file that provides an overlay for build operations.
		The file is a JSON struct with a single field, named 'Replace', that
		maps each disk file path (a string) to its backing file path, so that
		a build will run as if the disk file path exists with the contents
		given by the backing file paths, or as if the disk file path does not
		exist if its backing file path is empty. Support for the -overlay flag
		has some limitations: importantly, cgo files included from outside the
		include path must be in the same directory as the Go package they are
		included from, and overlays will not appear when binaries and tests are
		run through go run and go test respectively.

	-pkgdir dir
		install and load all packages from dir instead of the usual locations.
		For example, when building with a non-standard configuration,
		use -pkgdir to keep generated packages in a separate location.

	-trimpath
		remove all file system paths from the resulting executable.
		Instead of absolute file system paths, the recorded file names
		will begin either a module path@version (when using modules),
		or a plain import path (when using the standard library, or GOPATH).

	-toolexec 'cmd args'
		a program to use to invoke toolchain programs like vet and asm.
		For example, instead of running asm, the go command will run
		'cmd args /path/to/asm <arguments for asm>'.
		The TOOLEXEC_IMPORTPATH environment variable will be set,
		matching 'go list -f {{.ImportPath}}' for the package being built.

The -asmflags, -gccgoflags, -gcflags, and -ldflags flags accept a
space-separated list of arguments to pass to an underlying tool
during the build. To embed spaces in an element in the list, surround
it with either single or double quotes. The argument list may be
preceded by a package pattern and an equal sign, which restricts
the use of that argument list to the building of packages matching
that pattern (see 'go help packages' for a description of package
patterns). Without a pattern, the argument list applies only to the
packages named on the command line. The flags may be repeated
with different patterns in order to specify different arguments for
different sets of packages. If a package matches patterns given in
multiple flags, the latest match on the command line wins.
For example, 'go build -gcflags=-S fmt' prints the disassembly
only for package fmt, while 'go build -gcflags=all=-S fmt'
prints the disassembly for fmt and all its dependencies.
-asmflags、-gccgoflags、-gcflags和-ldflags标志接受一个用空格分隔的参数列
表，在构建过程中传递给底层工具。若要在列表中的元素中嵌入空格，请用单引号或双引号
将其括起来。实参列表前面可以有一个包模式和一个等号，这将限制使用该实参列表来构建
与该模式匹配的包(有关包模式的描述，请参阅“go help packages”)。如果没有模式，
参数列表只应用于命令行上命名的包。这些标志可以用不同的模式重复，以便为不同的包集
指定不同的参数。如果一个包匹配多个标志中给出的模式，则命令行上最新的匹配胜出。例
如，'go build -gcflags=-S fmt'只打印fmt包的反汇编，而'go build 
-gcflags=all=-S fmt'打印fmt及其所有依赖项的反汇编。

For more about specifying packages, see 'go help packages'.
For more about where packages and binaries are installed,
run 'go help gopath'.
For more about calling between Go and C/C++, run 'go help c'.

Note: Build adheres to certain conventions such as those described
by 'go help gopath'. Not all projects can follow these conventions,
however. Installations that have their own conventions or that use
a separate software build system may choose to use lower-level
invocations such as 'go tool compile' and 'go tool link' to avoid
some of the overheads and design decisions of the build tool.
构建遵循某些约定，例如“go help gopath”所描述的那些约定。然而，并不是所有的项
目都可以遵循这些约定。有自己的约定或使用单独软件构建系统的安装可能会选择使用较低
级别的调用，例如“go tool compile”和“go tool link”，以避免构建工具的一些开
销和设计决策。

See also: go install, go get, go clean.

```

