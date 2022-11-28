### Overview

* https://pkg.go.dev/testing@go1.19.3

Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form  包测试为Go包的自动化测试提供了支持。它旨在与“go test”命令一起使用，“go test”命令自动执行表单的任何功能

```go
func TestXxx(*testing.T)
```

where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.  Xxx不是以小写字母开头的。函数名用于标识测试例程。

Within these functions, use the Error, Fail or related methods to signal failure.  在这些函数中，使用Error、Fail或相关方法来表示故障。

To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the "go test" command is run. For more detail, run "go help test" and "go help testflag".  要编写一个新的测试套件，创建一个文件名以_test结尾的文件。它包含TestXxx函数，如这里所述。将该文件放在与被测试的文件相同的包中。该文件将从常规的包构建中排除，但在运行“go test”命令时将包含其中。更多详细信息，运行"go help test"和"go help testflag"。

A simple test function looks like this:  一个简单的测试函数是这样的:

```go
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

#### Benchmarks

Functions of the form

```go
func BenchmarkXxx(*testing.B)
```

are considered benchmarks, and are executed by the "go test" command when its `-bench` flag is provided. Benchmarks are run sequentially.

For a description of the testing flags, see https://golang.org/cmd/go/#hdr-Testing_flags.

A sample benchmark function looks like this:

```go
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably. The output

```
BenchmarkRandInt-8       68453040            17.8 ns/op
```

means that the loop ran 68453040 times at a speed of 17.8 ns per loop.

If a benchmark needs some expensive setup before running, the timer may be reset:

```go
func BenchmarkBigLen(b *testing.B) {
    big := NewBig()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        big.Len()
    }
}
```

If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function; such benchmarks are intended to be used with the `go test -cpu` flag:

```go
func BenchmarkTemplateParallel(b *testing.B) {
    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
    b.RunParallel(func(pb *testing.PB) {
        var buf bytes.Buffer
        for pb.Next() {
            buf.Reset()
            templ.Execute(&buf, "World")
        }
    })
}
```

A detailed specification of the benchmark results format is given in https://golang.org/design/14313-benchmark-format.

There are standard tools for working with benchmark results at https://golang.org/x/perf/cmd. In particular, https://golang.org/x/perf/cmd/benchstat performs statistically robust A/B comparisons.

#### Examples

The package also runs and verifies example code. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run. (The comparison ignores leading and trailing space.) These are examples of an example:  该包还运行和验证示例代码。示例函数可能包含以“Output:”开头的结束行注释，并在运行测试时与函数的标准输出进行比较。(这个比较忽略了前导和后导空格。)以下是一个例子中的例子:

```go
func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}

func ExampleSalutations() {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}
```

The comment prefix "Unordered output:" is like "Output:", but matches any line order:

```go
func ExamplePerm() {
    for _, value := range Perm(5) {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}
```

Example functions without output comments are compiled but not executed.

The naming convention to declare examples for the package, a function F, a type T and method M on type T are:

```go
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

Multiple example functions for a package/type/function/method may be provided by appending a distinct suffix to the name. The suffix must start with a lower-case letter.

```go
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

The entire test file is presented as the example when it contains a single example function, at least one other function, type, variable, or constant declaration, and no test or benchmark functions.

#### Fuzzing

'go test' and the testing package support fuzzing, a testing technique where a function is called with randomly generated inputs to find bugs not anticipated by unit tests.

Functions of the form

```go
func FuzzXxx(*testing.F)
```

are considered fuzz tests.

For example:

```go
func FuzzHex(f *testing.F) {
  for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
    f.Add(seed)
  }
  f.Fuzz(func(t *testing.T, in []byte) {
    enc := hex.EncodeToString(in)
    out, err := hex.DecodeString(enc)
    if err != nil {
      t.Fatalf("%v: decode: %v", in, err)
    }
    if !bytes.Equal(in, out) {
      t.Fatalf("%v: not equal after round trip: %v", in, out)
    }
  })
}
```

A fuzz test maintains a seed corpus, or a set of inputs which are run by default, and can seed input generation. Seed inputs may be registered by calling `(*F).Add` or by storing files in the directory `testdata/fuzz/<Name>` (where `<Name>` is the name of the fuzz test) within the package containing the fuzz test. Seed inputs are optional, but the fuzzing engine may find bugs more efficiently when provided with a set of small seed inputs with good code coverage. These seed inputs can also serve as regression tests for bugs identified through fuzzing.  模糊测试维护一个种子语料库，或一组默认运行的输入，并可以种子输入生成。种子输入可以通过调用`(*F)`来注册。在包含模糊测试的包中添加或存储文件到`testdata/fuzz/<Name>`(其中`<Name>`是模糊测试的名称)目录下。种子输入是可选的，但是当提供一组具有良好代码覆盖率的小种子输入时，模糊引擎可能更有效地发现错误。这些种子输入还可以作为通过模糊识别的错误的回归测试。

The function passed to `(*F).Fuzz` within the fuzz test is considered the fuzz target. A fuzz target must accept a `*T parameter`, followed by one or more parameters for random inputs. The types of arguments passed to `(*F).Add` must be identical to the types of these parameters. The fuzz target may signal that it's found a problem the same way tests do: by calling `T.Fail` (or any method that calls it like `T.Error` or `T.Fatal`) or by panicking.  传递给`(*F)`的函数。模糊测试中的模糊被认为是模糊目标。模糊目标必须接受一个`*T`参数，后面跟着一个或多个随机输入参数。传递给`(*F)`的参数类型。Add必须与这些参数的类型相同。模糊目标可能会以与测试相同的方式发出发现问题的信号:通过调用`T.Fail`(或调用它的任何方法，如`T.Error`或`T.Fatal`)或惊慌失措。

When fuzzing is enabled (by setting the `-fuzz` flag to a regular expression that matches a specific fuzz test), the fuzz target is called with arguments generated by repeatedly making random changes to the seed inputs. On supported platforms, `'go test'` compiles the test executable with fuzzing coverage instrumentation. The fuzzing engine uses that instrumentation to find and cache inputs that expand coverage, increasing the likelihood of finding bugs. If the fuzz target fails for a given input, the fuzzing engine writes the inputs that caused the failure to a file in the directory `testdata/fuzz/<Name>` within the package directory. This file later serves as a seed input. If the file can't be written at that location (for example, because the directory is read-only), the fuzzing engine writes the file to the fuzz cache directory within the build cache instead.  当启用模糊处理(通过将`-fuzz`标志设置为匹配特定模糊测试的正则表达式)时，通过反复随机更改种子输入生成的参数调用模糊目标。在受支持的平台上，`go test`使用模糊覆盖工具编译测试可执行文件。模糊引擎使用该工具来查找和缓存扩展覆盖范围的输入，增加发现错误的可能性。如果模糊目标对于给定的输入失败，模糊引擎将导致失败的输入写入包目录中`testdata/fuzz/<Name>`目录下的文件。该文件稍后用作种子输入。如果不能在该位置写入文件(例如，因为该目录是只读的)，则模糊引擎将该文件写入构建缓存中的模糊缓存目录。

When fuzzing is disabled, the fuzz target is called with the seed inputs registered with `F.Add` and seed inputs from `testdata/fuzz/<Name>`. In this mode, the fuzz test acts much like a regular test, with subtests started with `F.Fuzz` instead of `T.Run`.  当fuzzing被禁用时，fuzzing目标被调用，使用`F.Add`注册的种子输入和来自`testdata/fuzz/<Name>`的种子输入。在此模式下，模糊测试的行为与常规测试非常相似，其子测试以`F.Fuzz`而不是`T.Run`开始。

See [Go Fuzzing - The Go Programming Language](https://go.dev/doc/fuzz) for documentation about fuzzing.

#### Skipping

Tests or benchmarks may be skipped at run time with a call to the Skip method of `*T` or `*B`:

```go
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

The Skip method of `*T` can be used in a fuzz target if the input is invalid, but should not be considered a failing input. For example:

```go
func FuzzJSONMarshaling(f *testing.F) {
    f.Fuzz(func(t *testing.T, b []byte) {
        var v interface{}
        if err := json.Unmarshal(b, &v); err != nil {
            t.Skip()
        }
        if _, err := json.Marshal(v); err != nil {
            t.Error("Marshal: %v", err)
        }
    })
}
```

#### Subtests and Sub-benchmarks

The Run methods of T and B allow defining subtests and sub-benchmarks, without having to define separate functions for each. This enables uses like table-driven benchmarks and creating hierarchical tests. It also provides a way to share common setup and tear-down code:

```go
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // <tear-down code>
}
```

Each subtest and sub-benchmark has a unique name: the combination of the name of the top-level test and the sequence of names passed to Run, separated by slashes, with an optional trailing sequence number for disambiguation.

The argument to the `-run`, `-bench`, and `-fuzz` command-line flags is an unanchored regular expression that matches the test's name. For tests with multiple slash-separated elements, such as subtests, the argument is itself slash-separated, with expressions matching each name element in turn. Because it is unanchored, an empty expression matches any string. For example, using "matching" to mean "whose name contains":

```bash
go test -run ''        # Run all tests.
go test -run Foo       # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=    # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1      # For all top-level tests, run subtests matching "A=1".
go test -fuzz FuzzFoo  # Fuzz the target matching "FuzzFoo"
```

The `-run` argument can also be used to run a specific value in the seed corpus, for debugging. For example:

```bash
go test -run=FuzzFoo/9ddb952d9814
```

The `-fuzz` and `-run` flags can both be set, in order to fuzz a target but skip the execution of all other tests.

Subtests can also be used to control parallelism. A parent test will only complete once all of its subtests complete. In this example, all tests are run in parallel with each other, and only with each other, regardless of other top-level tests that may be defined:

```go
func TestGroupedParallel(t *testing.T) {
    for _, tc := range tests {
        tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            ...
        })
    }
}
```

Run does not return until parallel subtests have completed, providing a way to clean up after a group of parallel tests:

```go
func TestTeardownParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

#### Main

It is sometimes necessary for a test or benchmark program to do extra setup or teardown before or after it executes. It is also sometimes necessary to control which code runs on the main thread. To support these and other cases, if a test file contains a function:  有时，测试或基准测试程序在执行之前或之后需要进行额外的设置或销毁。有时还需要控制哪些代码运行在主线程上。为了支持这些情况和其他情况，如果一个测试文件包含一个函数:

```go
func TestMain(m *testing.M)
```

then the generated test will call TestMain(m) instead of running the tests or benchmarks directly. TestMain runs in the main goroutine and can do whatever setup and teardown is necessary around a call to `m.Run`. `m.Run` will return an exit code that may be passed to os.Exit. If TestMain returns, the test wrapper will pass the result of `m.Run` to `os.Exit` itself.  那么生成的测试将调用TestMain(m)，而不是直接运行测试或基准测试。TestMain在主gor例程中运行，并且可以在调用`m.Run`时进行任何必要的设置和拆卸。`m.Run`将返回一个退出码，该退出码可能被传递给`os.Exit`。如果TestMain返回，测试包装器将把`m.f` rrun的结果传递给os。退出本身。

When TestMain is called, flag.Parse has not been run. If TestMain depends on command-line flags, including those of the testing package, it should call flag.Parse explicitly. Command line flags are always parsed by the time test or benchmark functions run.  当TestMain被调用时，标记。解析未运行。如果TestMain依赖于命令行标志，包括测试包的标志，那么它应该调用flag。显式解析。命令行标志总是根据运行测试或基准测试函数的时间进行解析。

A simple implementation of TestMain is:

```go
func TestMain(m *testing.M) {
    // call flag.Parse() here if TestMain uses flags
    os.Exit(m.Run())
}
```

TestMain is a low-level primitive and should not be necessary for casual testing needs, where ordinary test functions suffice.


```go
func AllocsPerRun(runs int, f func()) (avg float64)
func CoverMode() string
func Coverage() float64
func Init()
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, ...)
func RegisterCover(c Cover)
func RunBenchmarks(matchString func(pat, str string) (bool, error), ...)
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
func Short() bool
func Verbose() bool


// Benchmarks
type B struct {
	N int
	// contains filtered or unexported fields
}
func (c *B) Cleanup(f func())
func (c *B) Error(args ...any)
func (c *B) Errorf(format string, args ...any)
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...any)
func (c *B) Fatalf(format string, args ...any)
func (c *B) Helper()
func (c *B) Log(args ...any)
func (c *B) Logf(format string, args ...any)
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ReportMetric(n float64, unit string)
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Setenv(key, value string)
func (c *B) Skip(args ...any)
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...any)
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()
func (c *B) TempDir() string


type BenchmarkResult struct {
	N         int           // The number of iterations.
	T         time.Duration // The total time taken.
	Bytes     int64         // Bytes processed in one iteration.
	MemAllocs uint64        // The total number of memory allocations.
	MemBytes  uint64        // The total number of bytes allocated.

	// Extra records additional metrics reported by ReportMetric.
	Extra map[string]float64
}
func Benchmark(f func(b *B)) BenchmarkResult
func (r BenchmarkResult) AllocedBytesPerOp() int64
func (r BenchmarkResult) AllocsPerOp() int64
func (r BenchmarkResult) MemString() string
func (r BenchmarkResult) NsPerOp() int64
func (r BenchmarkResult) String() string


type Cover struct {
	Mode            string
	Counters        map[string][]uint32
	Blocks          map[string][]CoverBlock
	CoveredPackages string
}


type CoverBlock struct {
	Line0 uint32 // Line number for block start.
	Col0  uint16 // Column number for block start.
	Line1 uint32 // Line number for block end.
	Col1  uint16 // Column number for block end.
	Stmts uint16 // Number of statements included in this block.
}


// Fuzzing
type F struct {
	// contains filtered or unexported fields
}
func (f *F) Add(args ...any)
func (c *F) Cleanup(f func())
func (c *F) Error(args ...any)
func (c *F) Errorf(format string, args ...any)
func (f *F) Fail()
func (c *F) FailNow()
func (c *F) Failed() bool
func (c *F) Fatal(args ...any)
func (c *F) Fatalf(format string, args ...any)
func (f *F) Fuzz(ff any)
func (f *F) Helper()
func (c *F) Log(args ...any)
func (c *F) Logf(format string, args ...any)
func (c *F) Name() string
func (c *F) Setenv(key, value string)
func (c *F) Skip(args ...any)
func (c *F) SkipNow()
func (c *F) Skipf(format string, args ...any)
func (f *F) Skipped() bool
func (c *F) TempDir() string


type InternalBenchmark struct {
	Name string
	F    func(b *B)
}


type InternalExample struct {
	Name      string
	F         func()
	Output    string
	Unordered bool
}


type InternalFuzzTarget struct {
	Name string
	Fn   func(f *F)
}


type InternalTest struct {
	Name string
	F    func(*T)
}


// Main
type M struct {
	// contains filtered or unexported fields
}
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, ...) *M
func (m *M) Run() (code int)


//
type PB struct {
	// contains filtered or unexported fields
}
func (pb *PB) Next() bool


// 普通测试
type T struct {
	// contains filtered or unexported fields
}
func (c *T) Cleanup(f func())
func (t *T) Deadline() (deadline time.Time, ok bool)
func (c *T) Error(args ...any)
func (c *T) Errorf(format string, args ...any)
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...any)
func (c *T) Fatalf(format string, args ...any)
func (c *T) Helper()
func (c *T) Log(args ...any)
func (c *T) Logf(format string, args ...any)
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (t *T) Setenv(key, value string)
func (c *T) Skip(args ...any)
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...any)
func (c *T) Skipped() bool
func (c *T) TempDir() string


type TB interface {
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
	// contains filtered or unexported methods
}

```

