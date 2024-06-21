# The Go Programming Language Specification
Go 编程语言规范

* https://go.dev/ref/spec

## Language version go1.22 (Feb 6, 2024)
语言版本go1.22（2024年2月6日）

Table of Contents

- Introduction  介绍

- Notation  符号

- Source code representation  源代码表示
  - Characters  字符
  - Letters and digits  字母和数字

- Lexical elements  词汇元素
  - Comments  注释
  - Tokens
  - Semicolons  分号
  - Identifiers  标识
  - Keywords  关键词
  - Operators and punctuation  运算符和标点符号
  - Integer literals  整数文字
  - Floating-point literals  浮点文字
  - Imaginary literals  虚构的文字
  - Rune literals  符文文字
  - String literals  字符串文字

- Constants  常数

- Variables  变量

- Types  类型
  - Boolean types  布尔类型
  - Numeric types  数字类型
  - String types  字符串类型
  - Array types  数组类型
  - Slice types  切片类型
  - Struct types  结构类型
  - Pointer types  指针类型
  - Function types  功能类型
  - Interface types  接口类型
  - Map types  地图类型
  - Channel types  渠道类型

- Properties of types and values  类型和值的属性
  - Underlying types  底层类型
  - Core types 核心类型
  - Type identity  类型标识
  - Assignability  可分配性
  - Representability  代表性
  - Method sets  方法集

- Blocks  积木

- Declarations and scope  声明和范围
  - Label scopes  标签范围
  - Blank identifier  空白标识符
  - Predeclared identifiers  预声明标识符
  - Exported identifiers  导出的标识符
  - Uniqueness of identifiers  标识符的唯一性
  - Constant declarations  常量声明
  - Iota  伊奥塔
  - Type declarations  类型声明
  - Type parameter declarations  类型参数声明
  - Variable declarations  变量声明
  - Short variable declarations  简短的变量声明
  - Function declarations  函数声明
  - Method declarations  方法声明
  - Expressions  表达式
  - Operands  操作数
  - Qualified identifiers  合格标识符
  - Composite literals  复合文字
  - Function literals  函数字面量
  - Primary expressions  主要表达方式
  - Selectors  选择器
  - Method expressions  方法表达式
  - Method values  方法值
  - Index expressions  索引表达式
  - Slice expressions  切片表达式
  - Type assertions  类型断言
  - Calls  通话
  - Passing arguments to ... parameters  将参数传递给...参数
  - Instantiations  实例化
  - Type inference  类型推断
  - Operators  运营商
  - Arithmetic operators  算术运算符
  - Comparison operators  比较运算符
  - Logical operators  逻辑运算符
  - Address operators  地址运算符
  - Receive operator  接收操作员
  - Conversions  转换
  - Constant expressions  常量表达式
  - Order of evaluation  评估顺序

- Statements  声明
  - Terminating statements  终止语句
  - Empty statements  空语句
  - Labeled statements  标签语句
  - Expression statements  表达式语句
  - Send statements 发送报表
  - IncDec statements  IncDec 报表
  - Assignment statements  赋值语句
  - If statements  如果语句
  - Switch statements  Switch 语句
  - For statements  对于报表
  - Go statements  Go 语句
  - Select statements  选择语句
  - Return statements  退货报表
  - Break statements  中断语句
  - Continue statements  继续陈述
  - Goto statements  转到语句
  - Fallthrough statements  失败声明
  - Defer statements  推迟陈述

- Built-in functions  内置功能
  - Appending to and copying slices  追加和复制切片
  - Clear  清除
  - Close  关闭
  - Manipulating complex numbers  操作复数
  - Deletion of map elements  删除地图元素
  - Length and capacity  长度和容量
  - Making slices, maps and channels  制作切片、地图和通道
  - Min and max  最小值和最大值
  - Allocation  分配
  - Handling panics  处理恐慌
  - Bootstrapping  自举

- Packages  套餐
  - vSource file organization  源文件组织
  - Package clause  套餐条款
  - Import declarations  进口报关
  - An example package  示例包

- Program initialization and execution  程序初始化和执行
  - The zero value  零值
  - Package initialization  包初始化
  - Program initialization  程序初始化
  - Program execution  程序执行

- Errors  错误

- Run-time panics  运行时恐慌

- System considerations  系统注意事项
  - Package unsafe  包裹不安全
  - Size and alignment guarantees  尺寸和对齐保证

- Appendix  附录
  - Language versions  语言版本
  - Type unification rules  类型统一规则


## Introduction
介绍

This is the reference manual for the Go programming language. The pre-Go1.18 version, without generics, can be found [here](https://go.dev/doc/go1.17_spec.html). For more information and other documents, see [go.dev](https://go.dev/).
这是 Go 编程语言的参考手册。没有泛型的 Go1.18 之前的版本可以在这里找到。有关更多信息和其他文档，请参阅 go.dev。

Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from *packages*, whose properties allow efficient management of dependencies.
Go 是一种通用语言，其设计时考虑到了系统编程。它是强类型和垃圾收集的，并且明确支持并发编程。程序是由包构建的，包的属性允许有效管理依赖关系。

The syntax is compact and simple to parse, allowing for easy analysis by automatic tools such as integrated development environments.
语法紧凑且易于解析，可以通过集成开发环境等自动化工具轻松分析。

## Notation
符号

The syntax is specified using a [variant](https://en.wikipedia.org/wiki/Wirth_syntax_notation) of Extended Backus-Naur Form (EBNF):
使用扩展巴科斯范式 (EBNF) 的变体指定语法：

```
Syntax      = { Production } .
Production  = production_name "=" [ Expression ] "." .
Expression  = Term { "|" Term } .
Term        = Factor { Factor } .
Factor      = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:
产生式是由术语和以下运算符构造的表达式，优先级递增：

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lowercase production names are used to identify lexical (terminal) tokens. Non-terminals are in CamelCase. Lexical tokens are enclosed in double quotes `""` or back quotes ````.
小写的产生式名称用于标识词汇（终端）标记。非终结符采用驼峰命名法。词汇标记用双引号 `""` 或反引号 ` `` ` 括起来。

The form `a … b` represents the set of characters from `a` through `b` as alternatives. The horizontal ellipsis `…` is also used elsewhere in the spec to informally denote various enumerations or code snippets that are not further specified. The character `…` (as opposed to the three characters `...`) is not a token of the Go language.
`a … b` 形式表示从 `a` 到 `b` 的替代字符集。水平省略号 `…` 也在规范中的其他地方使用，以非正式地表示未进一步指定的各种枚举或代码片段。字符 `…` （与三个字符 `...` 相对）不是 Go 语言的标记。

A link of the form [[Go 1.xx](https://go.dev/ref/spec#Language_versions)] indicates that a described language feature (or some aspect of it) was changed or added with language version 1.xx and thus requires at minimum that language version to build. For details, see the [linked section](https://go.dev/ref/spec#Language_versions) in the [appendix](https://go.dev/ref/spec#Appendix).
[Go 1.xx] 形式的链接表示所描述的语言功能（或其某些方面）已随语言版本 1.xx 更改或添加，因此至少需要该语言版本才能构建。有关详细信息，请参阅附录中的链接部分。

## Source code representation
源代码表示

Source code is Unicode text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8). The text is not canonicalized, so a single accented code point is distinct from the same character constructed from combining an accent and a letter; those are treated as two code points. For simplicity, this document will use the unqualified term *character* to refer to a Unicode code point in the source text.
源代码是采用 UTF-8 编码的 Unicode 文本。文本未规范化，因此单个重音代码点不同于由重音和字母组合而成的同一字符；这些被视为两个代码点。为简单起见，本文档将使用非限定术语字符来引用源文本中的 Unicode 代码点。

Each code point is distinct; for instance, uppercase and lowercase letters are different characters.
每个代码点都是不同的；例如，大写字母和小写字母是不同的字符。

Implementation restriction: For compatibility with other tools, a compiler may disallow the NUL character (U+0000) in the source text.
实现限制：为了与其他工具兼容，编译器可能不允许在源文本中使用 NUL 字符 (U+0000)。

Implementation restriction: For compatibility with other tools, a compiler may ignore a UTF-8-encoded byte order mark (U+FEFF) if it is the first Unicode code point in the source text. A byte order mark may be disallowed anywhere else in the source.
实现限制：为了与其他工具兼容，如果 UTF-8 编码的字节顺序标记 (U+FEFF) 是源文本中的第一个 Unicode 代码点，则编译器可能会忽略它。源中的其他任何地方都可能不允许使用字节顺序标记。

### Characters
字符

The following terms are used to denote specific Unicode character categories:
以下术语用于表示特定的 Unicode 字符类别：

```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

In [The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/), Section 4.5 "General Category" defines a set of character categories. Go treats all characters in any of the Letter categories Lu, Ll, Lt, Lm, or Lo as Unicode letters, and those in the Number category Nd as Unicode digits.
在 Unicode 标准 8.0 中，第 4.5 节“常规类别”定义了一组字符类别。 Go 将任何字母类别 Lu、Ll、Lt、Lm 或 Lo 中的所有字符视为 Unicode 字母，将数字类别 Nd 中的所有字符视为 Unicode 数字。

### Letters and digits
字母和数字

The underscore character `_` (U+005F) is considered a lowercase letter.
下划线字符 `_` (U+005F) 被视为小写字母。

```
letter        = [unicode_letter](https://go.dev/ref/spec#unicode_letter) | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

## Lexical elements
词汇元素

### Comments
注释

Comments serve as program documentation. There are two forms:
注释作为程序文档。有两种形式：

1. *Line comments* start with the character sequence `//` and stop at the end of the line.
   行注释以字符序列 `//` 开始，到行尾结束。

2. *General comments* start with the character sequence `/*` and stop with the first subsequent character sequence `*/`.
   一般注释以字符序列 `/*` 开始，以第一个后续字符序列 `*/` 结束。

A comment cannot start inside a [rune](https://go.dev/ref/spec#Rune_literals) or [string literal](https://go.dev/ref/spec#String_literals), or inside a comment. A general comment containing no newlines acts like a space. Any other comment acts like a newline.
注释不能在符文或字符串文字内或注释内开始。不包含换行符的一般注释就像一个空格。任何其他注释都像换行符一样。

### Tokens

Tokens form the vocabulary of the Go language. There are four classes: *identifiers*, *keywords*, *operators and punctuation*, and *literals*. *White space*, formed from spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A), is ignored except as it separates tokens that would otherwise combine into a single token. Also, a newline or end of file may trigger the insertion of a [semicolon](https://go.dev/ref/spec#Semicolons). While breaking the input into tokens, the next token is the longest sequence of characters that form a valid token.
令牌构成了 Go 语言的词汇表。有四类：标识符、关键字、运算符和标点符号以及文字。由空格 (U+0020)、水平制表符 (U+0009)、回车符 (U+000D) 和换行符 (U+000A) 组成的空白将被忽略，除非它分隔了否则会组合成单个的标记。令牌。此外，换行符或文件结尾可能会触发分号的插入。将输入分解为标记时，下一个标记是形成有效标记的最长字符序列。

### Semicolons
分号

The formal syntax uses semicolons `";"` as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:
正式语法在许多产生式中使用分号 `";"` 作为终止符。 Go 程序可以使用以下两个规则省略大部分分号：

1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
   当输入被分解为标记时，如果该标记是，则分号会立即自动插入到标记流中，位于该行的最终标记之后

   - an [identifier](https://go.dev/ref/spec#Identifiers)
     一个标识符

   - an [integer](https://go.dev/ref/spec#Integer_literals), [floating-point](https://go.dev/ref/spec#Floating-point_literals), [imaginary](https://go.dev/ref/spec#Imaginary_literals), [rune](https://go.dev/ref/spec#Rune_literals), or [string](https://go.dev/ref/spec#String_literals) literal
     整数、浮点、虚数、符文或字符串文字

   - one of the [keywords](https://go.dev/ref/spec#Keywords) `break`, `continue`, `fallthrough`, or `return`
     关键字 `break` 、 `continue` 、 `fallthrough` 或 `return` 之一

   - one of the [operators and punctuation](https://go.dev/ref/spec#Operators_and_punctuation) `++`, `--`, `)`, `]`, or `}`
     运算符和标点符号之一 `++` 、 `--` 、 `)` 、 `]` 或 `}`

2. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing `")"` or `"}"`.
   为了允许复杂的语句占据一行，可以在结束 `")"` 或 `"}"` 之前省略分号。

To reflect idiomatic use, code examples in this document elide semicolons using these rules.
为了反映惯用用法，本文档中的代码示例使用这些规则省略了分号。

### Identifiers
标识

Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.
标识符命名程序实体，例如变量和类型。标识符是一个或多个字母和数字的序列。标识符中的第一个字符必须是字母。

```
identifier = letter { letter | unicode_digit } .
```

```
a
_x9
ThisVariableIsExported
αβ
```

Some identifiers are [predeclared](https://go.dev/ref/spec#Predeclared_identifiers).
一些标识符是预先声明的。

### Keywords
关键词

The following keywords are reserved and may not be used as identifiers.
以下关键字是保留的，不能用作标识符。

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operators and punctuation
运算符和标点符号

The following character sequences represent [operators](https://go.dev/ref/spec#Operators) (including [assignment operators](https://go.dev/ref/spec#Assignment_statements)) and punctuation [[Go 1.18](https://go.dev/ref/spec#Go_1.18)]:
以下字符序列表示运算符（包括赋值运算符）和标点符号[Go 1.18]：

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

### Integer literals
整数文字

An integer literal is a sequence of digits representing an [integer constant](https://go.dev/ref/spec#Constants). An optional prefix sets a non-decimal base: `0b` or `0B` for binary, `0`, `0o`, or `0O` for octal, and `0x` or `0X` for hexadecimal [[Go 1.13](https://go.dev/ref/spec#Go_1.13)]. A single `0` is considered a decimal zero. In hexadecimal literals, letters `a` through `f` and `A` through `F` represent values 10 through 15.
整数文字是表示整数常量的数字序列。可选前缀设置非十进制基数： `0b` 或 `0B` （二进制）、 `0` 、 `0o` 或 `0O` 或 `0X` 表示十六进制 [Go 1.13]。单个 `0` 被视为十进制零。在十六进制文字中，字母 `a` 到 `f` 和 `A` 到 `F` 表示值 10 到 15。

For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal's value.
为了便于阅读，下划线字符 `_` 可能出现在基本前缀之后或连续数字之间；这样的下划线不会改变文字的值。

```
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

```
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727
_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
```

### Floating-point literals
浮点文字

A floating-point literal is a decimal or hexadecimal representation of a [floating-point constant](https://go.dev/ref/spec#Constants).
浮点文字是浮点常量的十进制或十六进制表示形式。

A decimal floating-point literal consists of an integer part (decimal digits), a decimal point, a fractional part (decimal digits), and an exponent part (`e` or `E` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; one of the decimal point or the exponent part may be elided. An exponent value exp scales the mantissa (integer and fractional part) by 10exp.
十进制浮点文字由整数部分（小数位）、小数点、小数部分（小数位）和指数部分（ `e` 或 `E` 后跟组成可选的符号和小数位）。整数部分或小数部分之一可以省略；小数点或指数部分之一可以被省略。指数值 exp 将尾数（整数和小数部分）缩放 10 exp 。

A hexadecimal floating-point literal consists of a `0x` or `0X` prefix, an integer part (hexadecimal digits), a radix point, a fractional part (hexadecimal digits), and an exponent part (`p` or `P` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; the radix point may be elided as well, but the exponent part is required. (This syntax matches the one given in IEEE 754-2008 §5.12.3.) An exponent value exp scales the mantissa (integer and fractional part) by 2exp [[Go 1.13](https://go.dev/ref/spec#Go_1.13)].
十六进制浮点文字由 `0x` 或 `0X` 前缀、整数部分（十六进制数字）、小数点、小数部分（十六进制数字）和指数部分组成（ `p` 或 `P` 后跟可选的符号和小数位）。整数部分或小数部分之一可以省略；小数点也可以省略，但指数部分是必需的。 （此语法与 IEEE 754-2008 §5.12.3 中给出的语法匹配。）指数值 exp 将尾数（整数和小数部分）缩放 2 exp [Go 1.13]。

For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal value.
为了便于阅读，下划线字符 `_` 可能出现在基本前缀之后或连续数字之间；这样的下划线不会改变文字值。

```
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
```

```
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0
0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)
0x.p1        // invalid: mantissa has no digits
1p-2         // invalid: p exponent requires hexadecimal mantissa
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
1_.5         // invalid: _ must separate successive digits
1._5         // invalid: _ must separate successive digits
1.5_e1       // invalid: _ must separate successive digits
1.5e_1       // invalid: _ must separate successive digits
1.5e1_       // invalid: _ must separate successive digits
```

### Imaginary literals
虚构的文字

An imaginary literal represents the imaginary part of a [complex constant](https://go.dev/ref/spec#Constants). It consists of an [integer](https://go.dev/ref/spec#Integer_literals) or [floating-point](https://go.dev/ref/spec#Floating-point_literals) literal followed by the lowercase letter `i`. The value of an imaginary literal is the value of the respective integer or floating-point literal multiplied by the imaginary unit *i* [[Go 1.13](https://go.dev/ref/spec#Go_1.13)]
虚数文字表示复数常量的虚部。它由一个整数或浮点文字后跟小写字母 `i` 组成。虚数文字的值是相应整数或浮点文字乘以虚数单位的值[Go 1.13]

```
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

For backward compatibility, an imaginary literal's integer part consisting entirely of decimal digits (and possibly underscores) is considered a decimal integer, even if it starts with a leading `0`.
为了向后兼容，完全由十进制数字（可能还有下划线）组成的虚数文字的整数部分被视为十进制整数，即使它以前导 `0` 开头。

```
0i
0123i         // == 123i for backward-compatibility
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

### Rune literals
符文文字

A rune literal represents a [rune constant](https://go.dev/ref/spec#Constants), an integer value identifying a Unicode code point. **A rune literal is expressed as one or more characters enclosed in single quotes**, as in `'x'` or `'\n'`. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats.  
rune 字面量表示 rune 常量，即标识 Unicode 代码点的整数值。符文文字表示为用单引号括起来的一个或多个字符，如 `'x'` 或 `'\n'` 。引号内可以出现除换行符和未转义单引号之外的任何字符。单引号字符表示字符本身的 Unicode 值，而以反斜杠开头的多字符序列则以各种格式编码值。

The simplest form represents the single character within the quotes; since Go source text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent a single integer value. For instance, the literal `'a'` holds a single byte representing a literal `a`, Unicode U+0061, value `0x61`, while `'ä'` holds two bytes (`0xc3` `0xa4`) representing a literal `a`-dieresis, U+00E4, value `0xe4`.
最简单的形式表示引号内的单个字符；由于 Go 源文本是以 UTF-8 编码的 Unicode 字符，因此多个 UTF-8 编码的字节可以表示单个整数值。例如，文字 `'a'` 保存表示文字 `a` 的单个字节，Unicode U+0061，值 `0x61` ，而 `'ä'` 保存两个字节 ( `0xc3` `0xa4` ) 表示文字 `a` -dieresis, U+00E4, 值 `0xe4` 。

Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: `\x` followed by exactly two hexadecimal digits; `\u` followed by exactly four hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a plain backslash `\` followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.
多个反斜杠转义允许将任意值编码为 ASCII 文本。有四种方法可以将整数值表示为数值常量： `\x` 后跟正好两个十六进制数字； `\u` 后跟正好四个十六进制数字； `\U` 后跟正好八个十六进制数字，以及一个简单的反斜杠 `\` 后跟正好三个八进制数字。在每种情况下，文字的值都是由相应基数中的数字表示的值。

Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and `\U` represent Unicode code points so within them some values are illegal, in particular those above `0x10FFFF` and surrogate halves.
尽管这些表示都产生一个整数，但它们具有不同的有效范围。八进制转义符必须表示 0 到 255 之间的值（含 0 和 255）。十六进制转义符通过构造满足这个条件。转义符 `\u` 和 `\U` 表示 Unicode 代码点，因此其中的某些值是非法的，特别是 `0x10FFFF` 和代理项之上的值。

After a backslash, certain single-character escapes represent special values:
在反斜杠之后，某些单字符转义符代表特殊值：

```
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000B vertical tab
\\   U+005C backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
```

An unrecognized character following a backslash in a rune literal is illegal.
符文文字中反斜杠后面的无法识别的字符是非法的。

```
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```

```
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\k'         // illegal: k is not recognized after a backslash
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\400'       // illegal: octal value over 255
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

### String literals
字符串文字

A string literal represents a [string constant](https://go.dev/ref/spec#Constants) obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.
字符串文字表示通过连接字符序列获得的字符串常量。有两种形式：原始字符串文字和解释字符串文字。

Raw string literals are character sequences between back quotes, as in `` `foo` ``. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value.
原始字符串文字是反引号之间的字符序列，如 `foo` 中。引号内可以出现除反引号之外的任何字符。原始字符串文字的值是由引号之间的未解释（隐式 UTF-8 编码）字符组成的字符串；特别是，反斜杠没有特殊含义，并且字符串可能包含换行符。原始字符串文字内的回车符 ('\r') 将从原始字符串值中丢弃。

**Interpreted string literals are character sequences between double quotes**, as in `"bar"`. Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in [rune literals](https://go.dev/ref/spec#Rune_literals) (except that `\'` is illegal and `\"` is legal), with the same restrictions. The three-digit octal (`\`*nnn*) and two-digit hexadecimal (`\x`*nn*) escapes represent individual *bytes* of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual *characters*. Thus inside a string literal `\377` and `\xFF` represent a single byte of value `0xFF`=255, while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes `0xc3` `0xbf` of the UTF-8 encoding of character U+00FF.
解释的字符串文字是双引号之间的字符序列，如 `"bar"` 中。引号内可以出现除换行符和未转义双引号之外的任何字符。引号之间的文本形成文字的值，反斜杠转义被解释为符文文字中的内容（除了 `\'` 是非法的， `\"` 是合法的），具有相同的限制。三位数八进制 ( `\` nnn) 和两位十六进制 ( `\x` nn) 转义符表示结果字符串的各个字节；所有其他转义符表示单个字符的（可能是多字节）UTF-8 编码。因此，在字符串文字中 `\377` 和 `\xFF` 表示值 `0xFF` =255 的单个字节，而 `ÿ` 、 `\u00FF` 和 `\xc3\xbf` 表示字符 U+00FF 的 UTF-8 编码的两个字节 `0xc3` `0xbf` 。

```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

```
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:
这些示例都代表相同的字符串：

```
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

If the source code represents a character as two code points, such as a combining form involving an accent and a letter, the result will be an error if placed in a rune literal (it is not a single code point), and will appear as two code points if placed in a string literal.
如果源代码将一个字符表示为两个代码点，例如涉及重音符号和字母的组合形式，则如果放入符文文字（它不是单个代码点）中，结果将是错误，并且将显示为如果放置在字符串文字中，则有两个代码点。

## Constants
常数

There are *boolean constants*, *rune constants*, *integer constants*, *floating-point constants*, *complex constants*, and *string constants*. Rune, integer, floating-point, and complex constants are collectively called *numeric constants*.
有布尔常量、符文常量、整型常量、浮点常量、复数常量和字符串常量。符文常数、整数常数、浮点常数和复数常数统称为数值常数。

A constant value is represented by a [rune](https://go.dev/ref/spec#Rune_literals), [integer](https://go.dev/ref/spec#Integer_literals), [floating-point](https://go.dev/ref/spec#Floating-point_literals), [imaginary](https://go.dev/ref/spec#Imaginary_literals), or [string](https://go.dev/ref/spec#String_literals) literal, an identifier denoting a constant, a [constant expression](https://go.dev/ref/spec#Constant_expressions), a [conversion](https://go.dev/ref/spec#Conversions) with a result that is a constant, or the result value of some built-in functions such as `min` or `max` applied to constant arguments, `unsafe.Sizeof` applied to [certain values](https://go.dev/ref/spec#Package_unsafe), `cap` or `len` applied to [some expressions](https://go.dev/ref/spec#Length_and_capacity), `real` and `imag` applied to a complex constant and `complex` applied to numeric constants. The boolean truth values are represented by the predeclared constants `true` and `false`. The predeclared identifier [iota](https://go.dev/ref/spec#Iota) denotes an integer constant.  
常量值由符文、整数、浮点、虚数或字符串文字、表示常量的标识符、常量表达式、结果为常量的转换或某些内置函数的结果值表示。函数，例如应用于常量参数的 `min` 或 `max` 、应用于某些值的 `unsafe.Sizeof` 、 `cap` 或 `len` 应用于某些表达式， `real` 和 `imag` 应用于复杂常量， `complex` 应用于数值常量。布尔真值由预先声明的常量 `true` 和 `false` 表示。预先声明的标识符iota表示整型常量。

In general, complex constants are a form of [constant expression](https://go.dev/ref/spec#Constant_expressions) and are discussed in that section.  
一般来说，复数常量是常量表达式的一种形式，将在该部分中讨论。

Numeric constants represent exact values of arbitrary precision and do not overflow. Consequently, there are no constants denoting the IEEE-754 negative zero, infinity, and not-a-number values.  
数字常量表示任意精度的精确值并且不会溢出。因此，没有常量表示 IEEE-754 负零、无穷大和非数字值。

Constants may be [typed](https://go.dev/ref/spec#Types) or *untyped*. Literal constants, `true`, `false`, `iota`, and certain [constant expressions](https://go.dev/ref/spec#Constant_expressions) containing only untyped constant operands are untyped.  
常量可以是类型化的，也可以是非类型化的。文字常量、 `true` 、 `false` 、 `iota` 以及仅包含无类型常量操作数的某些常量表达式是无类型的。

A constant may be given a type explicitly by a [constant declaration](https://go.dev/ref/spec#Constant_declarations) or [conversion](https://go.dev/ref/spec#Conversions), or implicitly when used in a [variable declaration](https://go.dev/ref/spec#Variable_declarations) or an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or as an operand in an [expression](https://go.dev/ref/spec#Expressions). It is an error if the constant value cannot be [represented](https://go.dev/ref/spec#Representability) as a value of the respective type. If the type is a type parameter, the constant is converted into a non-constant value of the type parameter.  
常量可以通过常量声明或转换显式地指定类型，或者在变量声明或赋值语句中使用或作为表达式中的操作数时隐式地指定类型。如果常量值不能表示为相应类型的值，则这是一个错误。如果类型是类型参数，则常量将转换为类型参数的非常量值。

An untyped constant has a *default type* which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations) such as `i := 0` where there is no explicit type. The default type of an untyped constant is `bool`, `rune`, `int`, `float64`, `complex128`, or `string` respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.  
无类型常量具有默认类型，该类型是常量在需要类型化值的上下文中隐式转换为的类型，例如，在没有显式类型的短变量声明中，例如 `i := 0` 。无类型常量的默认类型为 `bool` 、 `rune` 、 `int` 、 `float64` 、 `complex128` 或 < b6> 分别取决于它是布尔值、符文值、整数、浮点值、复数还是字符串常量。

Implementation restriction: Although numeric constants have arbitrary precision in the language, a compiler may implement them using an internal representation with limited precision. That said, every implementation must:  
实现限制：尽管数字常量在语言中具有任意精度，但编译器可以使用精度有限的内部表示来实现它们。也就是说，每个实施都必须：

- Represent integer constants with at least 256 bits.  
  表示至少 256 位的整数常量。
- Represent floating-point constants, including the parts of a complex constant, with a mantissa of at least 256 bits and a signed binary exponent of at least 16 bits.  
  表示浮点常量，包括复数常量的部分，尾数至少为 256 位，带符号二进制指数至少为 16 位。
- Give an error if unable to represent an integer constant precisely.  
  如果无法精确表示整数常量，则给出错误。
- Give an error if unable to represent a floating-point or complex constant due to overflow.  
  如果由于溢出而无法表示浮点或复数常量，则给出错误。
- Round to the nearest representable constant if unable to represent a floating-point or complex constant due to limits on precision.  
  如果由于精度限制而无法表示浮点或复数常量，则舍入到最接近的可表示常量。

These requirements apply both to literal constants and to the result of evaluating [constant expressions](https://go.dev/ref/spec#Constant_expressions).  
这些要求既适用于文字常量，也适用于计算常量表达式的结果。

## Variables 变量

A variable is a storage location for holding a *value*. The set of permissible values is determined by the variable's *[type](https://go.dev/ref/spec#Types)*.  
变量是保存值的存储位置。允许值的集合由变量的类型确定。

A [variable declaration](https://go.dev/ref/spec#Variable_declarations) or, for function parameters and results, the signature of a [function declaration](https://go.dev/ref/spec#Function_declarations) or [function literal](https://go.dev/ref/spec#Function_literals) reserves storage for a named variable. Calling the built-in function [`new`](https://go.dev/ref/spec#Allocation) or taking the address of a [composite literal](https://go.dev/ref/spec#Composite_literals) allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) [pointer indirection](https://go.dev/ref/spec#Address_operators).  
变量声明，或者对于函数参数和结果，函数声明或函数文字的签名为命名变量保留存储。调用内置函数 `new` 或获取复合文字的地址会在运行时为变量分配存储空间。这样的匿名变量是通过（可能是隐式的）指针间接引用的。

*Structured* variables of [array](https://go.dev/ref/spec#Array_types), [slice](https://go.dev/ref/spec#Slice_types), and [struct](https://go.dev/ref/spec#Struct_types) types have elements and fields that may be [addressed](https://go.dev/ref/spec#Address_operators) individually. Each such element acts like a variable.  
数组、切片和结构类型的结构化变量具有可以单独寻址的元素和字段。每个这样的元素就像一个变量。

The *static type* (or just *type*) of a variable is the type given in its declaration, the type provided in the `new` call or composite literal, or the type of an element of a structured variable. Variables of interface type also have a distinct *dynamic type*, which is the (non-interface) type of the value assigned to the variable at run time (unless the value is the predeclared identifier `nil`, which has no type). The dynamic type may vary during execution but values stored in interface variables are always [assignable](https://go.dev/ref/spec#Assignability) to the static type of the variable.  
变量的静态类型（或仅类型）是其声明中给出的类型、 `new` 调用或复合文字中提供的类型，或者结构化变量的元素的类型。接口类型的变量也有一个独特的动态类型，它是运行时分配给变量的值的（非接口）类型（除非该值是预先声明的标识符 `nil` ，它没有类型）。动态类型在执行期间可能会发生变化，但存储在接口变量中的值始终可分配给变量的静态类型。

var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T

A variable's value is retrieved by referring to the variable in an [expression](https://go.dev/ref/spec#Expressions); it is the most recent value [assigned](https://go.dev/ref/spec#Assignment_statements) to the variable. If a variable has not yet been assigned a value, its value is the [zero value](https://go.dev/ref/spec#The_zero_value) for its type.  
通过引用表达式中的变量来检索变量的值；它是分配给变量的最新值。如果尚未为变量赋值，则其值是其类型的零值。

## Types 类型

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a *type name*, if it has one, which must be followed by [type arguments](https://go.dev/ref/spec#Instantiations) if the type is generic. A type may also be specified using a *type literal*, which composes a type from existing types.  
类型确定一组值以及特定于这些值的操作和方法。如果类型有类型名称，则可以用类型名称来表示；如果该类型是泛型，则类型名称后面必须跟有类型参数。类型也可以使用类型文字来指定，它由现有类型组成一个类型。

Type      = [TypeName](https://go.dev/ref/spec#TypeName) [ [TypeArgs](https://go.dev/ref/spec#TypeArgs) ] | [TypeLit](https://go.dev/ref/spec#TypeLit) | "(" [Type](https://go.dev/ref/spec#Type) ")" .
TypeName  = [identifier](https://go.dev/ref/spec#identifier) | [QualifiedIdent](https://go.dev/ref/spec#QualifiedIdent) .
TypeArgs  = "[" [TypeList](https://go.dev/ref/spec#TypeList) [ "," ] "]" .
TypeList  = [Type](https://go.dev/ref/spec#Type) { "," [Type](https://go.dev/ref/spec#Type) } .
TypeLit   = [ArrayType](https://go.dev/ref/spec#ArrayType) | [StructType](https://go.dev/ref/spec#StructType) | [PointerType](https://go.dev/ref/spec#PointerType) | [FunctionType](https://go.dev/ref/spec#FunctionType) | [InterfaceType](https://go.dev/ref/spec#InterfaceType) |
            [SliceType](https://go.dev/ref/spec#SliceType) | [MapType](https://go.dev/ref/spec#MapType) | [ChannelType](https://go.dev/ref/spec#ChannelType) .

The language [predeclares](https://go.dev/ref/spec#Predeclared_identifiers) certain type names. Others are introduced with [type declarations](https://go.dev/ref/spec#Type_declarations) or [type parameter lists](https://go.dev/ref/spec#Type_parameter_declarations). *Composite types*—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.  
该语言预先声明了某些类型名称。其他的则通过类型声明或类型参数列表来引入。复合类型（数组、结构体、指针、函数、接口、切片、映射和通道类型）可以使用类型文字来构造。

Predeclared types, defined types, and type parameters are called *named types*. An alias denotes a named type if the type given in the alias declaration is a named type.  
预声明类型、定义类型和类型参数称为命名类型。如果别名声明中给定的类型是命名类型，则别名表示命名类型。

### Boolean types 布尔类型

A *boolean type* represents the set of Boolean truth values denoted by the predeclared constants `true` and `false`. The predeclared boolean type is `bool`; it is a [defined type](https://go.dev/ref/spec#Type_definitions).  
布尔类型表示由预先声明的常量 `true` 和 `false` 表示的布尔真值集。预声明的布尔类型是 `bool` ；它是一个定义的类型。

### Numeric types 数字类型

An *integer*, *floating-point*, or *complex* type represents the set of integer, floating-point, or complex values, respectively. They are collectively called *numeric types*. The predeclared architecture-independent numeric types are:  
整数、浮点或复数类型分别表示整数、浮点或复数值的集合。它们统称为数字类型。预先声明的独立于体系结构的数字类型是：

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers
complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts
byte        alias for uint8
rune        alias for int32

The value of an *n*-bit integer is *n* bits wide and represented using [two's complement arithmetic](https://en.wikipedia.org/wiki/Two's_complement).  
n 位整数的值是 n 位宽，并使用二进制补码算术表示。

There is also a set of predeclared integer types with implementation-specific sizes:

uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

To avoid portability issues all numeric types are [defined types](https://go.dev/ref/spec#Type_definitions) and thus distinct except `byte`, which is an [alias](https://go.dev/ref/spec#Alias_declarations) for `uint8`, and `rune`, which is an alias for `int32`. Explicit conversions are required when different numeric types are mixed in an expression or assignment. For instance, `int32` and `int` are not the same type even though they may have the same size on a particular architecture.

### String types

A *string type* represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is `string`; it is a [defined type](https://go.dev/ref/spec#Type_definitions).

The length of a string `s` can be discovered using the built-in function [`len`](https://go.dev/ref/spec#Length_and_capacity). The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer [indices](https://go.dev/ref/spec#Index_expressions) 0 through `len(s)-1`. It is illegal to take the address of such an element; if `s[i]` is the `i`'th byte of a string, `&s[i]` is invalid.

### Array types 数组类型

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.  
数组是单一类型元素的编号序列，称为元素类型。元素的数量称为数组的长度，并且永远不会是负数。

ArrayType   = "[" [ArrayLength](https://go.dev/ref/spec#ArrayLength) "]" [ElementType](https://go.dev/ref/spec#ElementType) .
ArrayLength = [Expression](https://go.dev/ref/spec#Expression) .
ElementType = [Type](https://go.dev/ref/spec#Type) .

The length is part of the array's type; it must evaluate to a non-negative [constant](https://go.dev/ref/spec#Constants) [representable](https://go.dev/ref/spec#Representability) by a value of type `int`. The length of array `a` can be discovered using the built-in function [`len`](https://go.dev/ref/spec#Length_and_capacity). The elements can be addressed by integer [indices](https://go.dev/ref/spec#Index_expressions) 0 through `len(a)-1`. Array types are always one-dimensional but may be composed to form multi-dimensional types.

[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))

An array type `T` may not have an element of type `T`, or of a type containing `T` as a component, directly or indirectly, if those containing types are only array or struct types.

// invalid array types
type (
    T1 [10]T1                 // element type of T1 is T1
    T2 [10]struct{ f T2 }     // T2 contains T2 as component of a struct
    T3 [10]T4                 // T3 contains T3 as component of a struct in T4
    T4 struct{ f T3 }         // T4 contains T4 as component of array T3 in a struct
)
// valid array types
type (
    T5 [10]*T5                // T5 contains T5 as component of a pointer
    T6 [10]func() T6          // T6 contains T6 as component of a function type
    T7 [10]struct{ f []T7 }   // T7 contains T7 as component of a slice in a struct
)

### Slice types

A slice is a descriptor for a contiguous segment of an *underlying array* and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is `nil`.

SliceType = "[" "]" [ElementType](https://go.dev/ref/spec#ElementType) .

The length of a slice `s` can be discovered by the built-in function [`len`](https://go.dev/ref/spec#Length_and_capacity); unlike with arrays it may change during execution. The elements can be addressed by integer [indices](https://go.dev/ref/spec#Index_expressions) 0 through `len(s)-1`. The slice index of a given element may be less than the index of the same element in the underlying array.

A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

The array underlying a slice may extend past the end of the slice. The *capacity* is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by [*slicing*](https://go.dev/ref/spec#Slice_expressions) a new one from the original slice. The capacity of a slice `a` can be discovered using the built-in function [`cap(a)`](https://go.dev/ref/spec#Length_and_capacity).

A new, initialized slice value for a given element type `T` may be made using the built-in function [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels), which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with `make` always allocates a new, hidden array to which the returned slice value refers. That is, executing

make([]T, length, capacity)

produces the same slice as allocating an array and [slicing](https://go.dev/ref/spec#Slice_expressions) it, so these two expressions are equivalent:  
生成与分配数组并对其进行切片相同的切片，因此这两个表达式是等效的：

make([]int, 50, 100)
new([100]int)[0:50]

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.  
与数组一样，切片始终是一维的，但可以组合起来构造更高维的对象。对于数组的数组，内部数组在结构上总是相同的长度；然而，对于切片切片（或切片数组），内部长度可能会动态变化。此外，内部切片必须单独初始化。

### Struct types 结构类型

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-[blank](https://go.dev/ref/spec#Blank_identifier) field names must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers).  
结构体是命名元素（称为字段）的序列，每个元素都有名称和类型。字段名称可以显式指定（IdentifierList）或隐式指定（EmbeddedField）。在结构体中，非空白字段名称必须是唯一的。

StructType    = "struct" "{" { [FieldDecl](https://go.dev/ref/spec#FieldDecl) ";" } "}" .
FieldDecl     = ([IdentifierList](https://go.dev/ref/spec#IdentifierList) [Type](https://go.dev/ref/spec#Type) | [EmbeddedField](https://go.dev/ref/spec#EmbeddedField)) [ [Tag](https://go.dev/ref/spec#Tag) ] .
EmbeddedField = [ "*" ] [TypeName](https://go.dev/ref/spec#TypeName) [ [TypeArgs](https://go.dev/ref/spec#TypeArgs) ] .
Tag           = [string_lit](https://go.dev/ref/spec#string_lit) .

// An empty struct.
struct {}
// A struct with 6 fields.
struct {
    x, y int
    u float32
    _ float32  // padding
    A *[]int
    F func()
}

A field declared with a type but no explicit field name is called an *embedded field*. An embedded field must be specified as a type name `T` or as a pointer to a non-interface type name `*T`, and `T` itself may not be a pointer type. The unqualified type name acts as the field name.  
使用类型声明但没有显式字段名称的字段称为嵌入字段。嵌入字段必须指定为类型名称 `T` 或指向非接口类型名称 `*T` 的指针，并且 `T` 本身可能不是指针类型。非限定类型名称充当字段名称。

// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
    T1        // field name is T1
    *T2       // field name is T2
    P.T3      // field name is T3
    *P.T4     // field name is T4
    x, y int  // field names are x and y
}

The following declaration is illegal because field names must be unique in a struct type:  
以下声明是非法的，因为字段名称在结构类型中必须是唯一的：

struct {
    T     // conflicts with embedded field *T and *P.T
    *T    // conflicts with embedded field T and *P.T
    *P.T  // conflicts with embedded field T and *T
}

A field or [method](https://go.dev/ref/spec#Method_declarations) `f` of an embedded field in a struct `x` is called *promoted* if `x.f` is a legal [selector](https://go.dev/ref/spec#Selectors) that denotes that field or method `f`.  
如果 `x.f` 是表示字段或方法 `f` 中嵌入字段的字段或方法 `f` 被称为提升。 /b3> .

Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in [composite literals](https://go.dev/ref/spec#Composite_literals) of the struct.  
升级字段的作用与结构的普通字段类似，只是它们不能用作结构的复合文字中的字段名称。

Given a struct type `S` and a [named type](https://go.dev/ref/spec#Types) `T`, promoted methods are included in the method set of the struct as follows:  
给定一个结构体类型 `S` 和一个命名类型 `T` ，提升的方法将包含在该结构体的方法集中，如下所示：

- If `S` contains an embedded field `T`, the [method sets](https://go.dev/ref/spec#Method_sets) of `S` and `*S` both include promoted methods with receiver `T`. The method set of `*S` also includes promoted methods with receiver `*T`.  
  如果 `S` 包含嵌入字段 `T` ，则 `S` 和 `*S` 的方法集都包含接收者 `T` 的方法集还包括带有接收者 `*T` 的提升方法。
- If `S` contains an embedded field `*T`, the method sets of `S` and `*S` both include promoted methods with receiver `T` or `*T`.  
  如果 `S` 包含嵌入字段 `*T` ，则 `S` 和 `*S` 的方法集都包含接收者 `T` 。

A field declaration may be followed by an optional string literal *tag*, which becomes an attribute for all the fields in the corresponding field declaration. An empty tag string is equivalent to an absent tag. The tags are made visible through a [reflection interface](https://go.dev/pkg/reflect/#StructTag) and take part in [type identity](https://go.dev/ref/spec#Type_identity) for structs but are otherwise ignored.  
字段声明后面可以跟一个可选的字符串文字标签，该标签成为相应字段声明中所有字段的属性。空标签字符串相当于不存在的标签。这些标签通过反射接口变得可见，并参与结构的类型标识，但在其他情况下会被忽略。

struct {
    x, y float64 ""  // an empty tag string is like an absent tag
    name string  "any string is permitted as a tag"
    _    [4]byte "ceci n'est pas un champ de structure"
}
// A struct corresponding to a TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers;
// they follow the convention outlined by the reflect package.
struct {
    microsec  uint64 `protobuf:"1"`
    serverIP6 uint64 `protobuf:"2"`
}

A struct type `T` may not contain a field of type `T`, or of a type containing `T` as a component, directly or indirectly, if those containing types are only array or struct types.

// invalid struct types
type (
    T1 struct{ T1 }            // T1 contains a field of T1
    T2 struct{ f [10]T2 }      // T2 contains T2 as component of an array
    T3 struct{ T4 }            // T3 contains T3 as component of an array in struct T4
    T4 struct{ f [10]T3 }      // T4 contains T4 as component of struct T3 in an array
)
// valid struct types
type (
    T5 struct{ f *T5 }         // T5 contains T5 as component of a pointer
    T6 struct{ f func() T6 }   // T6 contains T6 as component of a function type
    T7 struct{ f [10][]T7 }    // T7 contains T7 as component of a slice in an array
)

### Pointer types 指针类型

A pointer type denotes the set of all pointers to [variables](https://go.dev/ref/spec#Variables) of a given type, called the *base type* of the pointer. The value of an uninitialized pointer is `nil`.

PointerType = "*" [BaseType](https://go.dev/ref/spec#BaseType) .
BaseType    = [Type](https://go.dev/ref/spec#Type) .

*Point
*[4]int

### Function types

A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is `nil`.

FunctionType   = "func" [Signature](https://go.dev/ref/spec#Signature) .
Signature      = [Parameters](https://go.dev/ref/spec#Parameters) [ [Result](https://go.dev/ref/spec#Result) ] .
Result         = [Parameters](https://go.dev/ref/spec#Parameters) | [Type](https://go.dev/ref/spec#Type) .
Parameters     = "(" [ [ParameterList](https://go.dev/ref/spec#ParameterList) [ "," ] ] ")" .
ParameterList  = [ParameterDecl](https://go.dev/ref/spec#ParameterDecl) { "," [ParameterDecl](https://go.dev/ref/spec#ParameterDecl) } .
ParameterDecl  = [ [IdentifierList](https://go.dev/ref/spec#IdentifierList) ] [ "..." ] [Type](https://go.dev/ref/spec#Type) .

Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-[blank](https://go.dev/ref/spec#Blank_identifier) names in the signature must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers). If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.  
在参数或结果列表中，名称 (IdentifierList) 必须全部存在或全部不存在。如果存在，则每个名称代表指定类型的一项（参数或结果），并且签名中的所有非空白名称必须是唯一的。如果不存在，则每种类型代表该类型的一项。参数和结果列表总是带括号的，除非只有一个未命名的结果，它可以写为不带括号的类型。

The final incoming parameter in a function signature may have a type prefixed with `...`. A function with such a parameter is called *variadic* and may be invoked with zero or more arguments for that parameter.  
函数签名中的最终传入参数可能具有前缀为 `...` 的类型。具有此类参数的函数称为可变参数，并且可以使用该参数的零个或多个参数来调用。

func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)

### Interface types 接口类型

An interface type defines a *type set*. A variable of interface type can store a value of any type that is in the type set of the interface. Such a type is said to [implement the interface](https://go.dev/ref/spec#Implementing_an_interface). The value of an uninitialized variable of interface type is `nil`.  
接口类型定义类型集。接口类型的变量可以存储接口类型集中的任何类型的值。这样的类型被称为实现接口。接口类型的未初始化变量的值为 `nil` 。

InterfaceType  = "interface" "{" { [InterfaceElem](https://go.dev/ref/spec#InterfaceElem) ";" } "}" .
InterfaceElem  = [MethodElem](https://go.dev/ref/spec#MethodElem) | [TypeElem](https://go.dev/ref/spec#TypeElem) .
MethodElem     = [MethodName](https://go.dev/ref/spec#MethodName) [Signature](https://go.dev/ref/spec#Signature) .
MethodName     = [identifier](https://go.dev/ref/spec#identifier) .
TypeElem       = [TypeTerm](https://go.dev/ref/spec#TypeTerm) { "|" [TypeTerm](https://go.dev/ref/spec#TypeTerm) } .
TypeTerm       = [Type](https://go.dev/ref/spec#Type) | [UnderlyingType](https://go.dev/ref/spec#UnderlyingType) .
UnderlyingType = "~" [Type](https://go.dev/ref/spec#Type) .

An interface type is specified by a list of *interface elements*. An interface element is either a *method* or a *type element*, where a type element is a union of one or more *type terms*. A type term is either a single type or a single underlying type.  
接口类型由接口元素列表指定。接口元素可以是方法，也可以是类型元素，其中类型元素是一个或多个类型术语的联合。类型术语可以是单一类型，也可以是单一基础类型。

#### Basic interfaces 基本接口

In its most basic form an interface specifies a (possibly empty) list of methods. The type set defined by such an interface is the set of types which implement all of those methods, and the corresponding [method set](https://go.dev/ref/spec#Method_sets) consists exactly of the methods specified by the interface. Interfaces whose type sets can be defined entirely by a list of methods are called *basic interfaces.*  
接口最基本的形式指定了一个（可能是空的）方法列表。这种接口定义的类型集是实现所有这些方法的类型集，并且相应的方法集恰好由该接口指定的方法组成。其类型集可以完全由方法列表定义的接口称为基本接口。

// A simple File interface.
interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
    Close() error
}

The name of each explicitly specified method must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers) and not [blank](https://go.dev/ref/spec#Blank_identifier).  
每个显式指定的方法的名称必须是唯一的且不能为空。

interface {
    String() string
    String() string  // illegal: String not unique
    _(x int)         // illegal: method must have non-blank name
}

More than one type may implement an interface. For instance, if two types `S1` and `S2` have the method set  
不止一种类型可以实现一个接口。例如，如果两个类型 `S1` 和 `S2` 设置了方法

func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error

(where `T` stands for either `S1` or `S2`) then the `File` interface is implemented by both `S1` and `S2`, regardless of what other methods `S1` and `S2` may have or share.  
（其中 `T` 代表 `S1` 或 `S2` ），则 `File` 接口由 `S1` 和 `S2` ，无论 `S1` 和 `S2` 可能拥有或共享什么其他方法。

Every type that is a member of the type set of an interface implements that interface. Any given type may implement several distinct interfaces. For instance, all types implement the *empty interface* which stands for the set of all (non-interface) types:  
作为接口类型集成员的每个类型都实现该接口。任何给定类型都可以实现多个不同的接口。例如，所有类型都实现空接口，它代表所有（非接口）类型的集合：

interface{}

For convenience, the predeclared type `any` is an alias for the empty interface. [[Go 1.18](https://go.dev/ref/spec#Go_1.18)]  
为了方便起见，预先声明的类型 `any` 是空接口的别名。 [转到1.18]

Similarly, consider this interface specification, which appears within a [type declaration](https://go.dev/ref/spec#Type_declarations) to define an interface called `Locker`:  
同样，考虑此接口规范，它出现在类型声明中以定义名为 `Locker` 的接口：

type Locker interface {
    Lock()
    Unlock()
}

If `S1` and `S2` also implement  
如果 `S1` 和 `S2` 也实现

func (p T) Lock() { … }
func (p T) Unlock() { … }

they implement the `Locker` interface as well as the `File` interface.  
它们实现了 `Locker` 接口以及 `File` 接口。

#### Embedded interfaces 嵌入式接口

In a slightly more general form an interface `T` may use a (possibly qualified) interface type name `E` as an interface element. This is called *embedding* interface `E` in `T` [[Go 1.14](https://go.dev/ref/spec#Go_1.14)]. The type set of `T` is the *intersection* of the type sets defined by `T`'s explicitly declared methods and the type sets of `T`’s embedded interfaces. In other words, the type set of `T` is the set of all types that implement all the explicitly declared methods of `T` and also all the methods of `E` [[Go 1.18](https://go.dev/ref/spec#Go_1.18)].  
在稍微更一般的形式中，接口 `T` 可以使用（可能限定的）接口类型名称 `E` 作为接口元素。这称为在 `T` 中嵌入接口 `E` [Go 1.14]。 `T` 的类型集是 `T` 显式声明的方法定义的类型集与 `T` 嵌入接口的类型集的交集。换句话说， `T` 的类型集是实现 `T` 的所有显式声明的方法以及 `E` 的所有方法的所有类型的集合[转到1.18]。

type Reader interface {
    Read(p []byte) (n int, err error)
    Close() error
}
type Writer interface {
    Write(p []byte) (n int, err error)
    Close() error
}
// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
    Reader  // includes methods of Reader in ReadWriter's method set
    Writer  // includes methods of Writer in ReadWriter's method set
}

When embedding interfaces, methods with the [same](https://go.dev/ref/spec#Uniqueness_of_identifiers) names must have [identical](https://go.dev/ref/spec#Type_identity) signatures.

type ReadCloser interface {
    Reader   // includes methods of Reader in ReadCloser's method set
    Close()  // illegal: signatures of Reader.Close and Close are different
}

#### General interfaces

In their most general form, an interface element may also be an arbitrary type term `T`, or a term of the form `~T` specifying the underlying type `T`, or a union of terms `t1|t2|…|tn` [[Go 1.18](https://go.dev/ref/spec#Go_1.18)]. Together with method specifications, these elements enable the precise definition of an interface's type set as follows:

- The type set of the empty interface is the set of all non-interface types.
- The type set of a non-empty interface is the intersection of the type sets of its interface elements.
- The type set of a method specification is the set of all non-interface types whose method sets include that method.
- The type set of a non-interface type term is the set consisting of just that type.
- The type set of a term of the form `~T` is the set of all types whose underlying type is `T`.
- The type set of a *union* of terms `t1|t2|…|tn` is the union of the type sets of the terms.

The quantification "the set of all non-interface types" refers not just to all (non-interface) types declared in the program at hand, but all possible types in all possible programs, and hence is infinite. Similarly, given the set of all non-interface types that implement a particular method, the intersection of the method sets of those types will contain exactly that method, even if all types in the program at hand always pair that method with another method.

By construction, an interface's type set never contains an interface type.

// An interface representing only the type int.
interface {
    int
}
// An interface representing all types with underlying type int.
interface {
    ~int
}
// An interface representing all types with underlying type int that implement the String method.
interface {
    ~int
    String() string
}
// An interface representing an empty type set: there is no type that is both an int and a string.
interface {
    int
    string
}

In a term of the form `~T`, the underlying type of `T` must be itself, and `T` cannot be an interface.  
在 `~T` 形式的术语中， `T` 的基础类型必须是其自身，并且 `T` 不能是接口。

type MyInt int
interface {
    ~[]byte  // the underlying type of []byte is itself
    ~MyInt   // illegal: the underlying type of MyInt is not MyInt
    ~error   // illegal: error is an interface
}

Union elements denote unions of type sets:  
联合元素表示类型集的联合：

// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either float32 or float64).
type Float interface {
    ~float32 | ~float64
}

The type `T` in a term of the form `T` or `~T` cannot be a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), and the type sets of all non-interface terms must be pairwise disjoint (the pairwise intersection of the type sets must be empty). Given a type parameter `P`:  
`T` 或 `~T` 形式的项中的类型 `T` 不能是类型参数，并且所有非接口项的类型集必须是成对不相交的（类型集的成对交集必须为空）。给定类型参数 `P` ：

interface {
    P                // illegal: P is a type parameter
    int | ~P         // illegal: P is a type parameter
    ~int | MyInt     // illegal: the type sets for ~int and MyInt are not disjoint (~int includes MyInt)
    float32 | Float  // overlapping type sets but Float is an interface
}

Implementation restriction: A union (with more than one term) cannot contain the [predeclared identifier](https://go.dev/ref/spec#Predeclared_identifiers) `comparable` or interfaces that specify methods, or embed `comparable` or interfaces that specify methods.  
实现限制：联合（具有多个术语）不能包含预声明的标识符 `comparable` 或指定方法的接口，也不能嵌入 `comparable` 或指定方法的接口。

Interfaces that are not [basic](https://go.dev/ref/spec#Basic_interfaces) may only be used as type constraints, or as elements of other interfaces used as constraints. They cannot be the types of values or variables, or components of other, non-interface types.  
非基本接口只能用作类型约束，或者用作用作约束的其他接口的元素。它们不能是值或变量的类型，也不能是其他非接口类型的组件。

var x Float                     // illegal: Float is not a basic interface
var x interface{} = Float(nil)  // illegal
type Floatish struct {
    f Float                 // illegal
}

An interface type `T` may not embed a type element that is, contains, or embeds `T`, directly or indirectly.  
接口类型 `T` 不得直接或间接嵌入、包含或嵌入 `T` 的类型元素。

// illegal: Bad may not embed itself
type Bad interface {
    Bad
}
// illegal: Bad1 may not embed itself using Bad2
type Bad1 interface {
    Bad2
}
type Bad2 interface {
    Bad1
}
// illegal: Bad3 may not embed a union containing Bad3
type Bad3 interface {
    ~int | ~string | Bad3
}
// illegal: Bad4 may not embed an array containing Bad4 as element type
type Bad4 interface {
    [10]Bad4
}

#### Implementing an interface

实现接口

A type `T` implements an interface `I` if

- `T` is not an interface and is an element of the type set of `I`; or
- `T` is an interface and the type set of `T` is a subset of the type set of `I`.

A value of type `T` implements an interface if `T` implements the interface.

### Map types

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique *keys* of another type, called the key type. The value of an uninitialized map is `nil`.

MapType     = "map" "[" [KeyType](https://go.dev/ref/spec#KeyType) "]" [ElementType](https://go.dev/ref/spec#ElementType) .
KeyType     = [Type](https://go.dev/ref/spec#Type) .

The [comparison operators](https://go.dev/ref/spec#Comparison_operators) `==` and `!=` must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics).

map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}

The number of map elements is called its length. For a map `m`, it can be discovered using the built-in function [`len`](https://go.dev/ref/spec#Length_and_capacity) and may change during execution. Elements may be added during execution using [assignments](https://go.dev/ref/spec#Assignment_statements) and retrieved with [index expressions](https://go.dev/ref/spec#Index_expressions); they may be removed with the [`delete`](https://go.dev/ref/spec#Deletion_of_map_elements) and [`clear`](https://go.dev/ref/spec#Clear) built-in function.  
地图元素的数量称为其长度。对于映射 `m` ，可以使用内置函数 `len` 发现它，并且可能在执行过程中发生变化。可以在执行期间使用赋值添加元素，并使用索引表达式检索元素；它们可以使用 `delete` 和 `clear` 内置函数删除。

A new, empty map value is made using the built-in function [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels), which takes the map type and an optional capacity hint as arguments:  
使用内置函数 `make` 创建一个新的空地图值，该函数将地图类型和可选的容量提示作为参数：

make(map[string]int)
make(map[string]int, 100)

The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of `nil` maps. A `nil` map is equivalent to an empty map except that no elements may be added.  
初始容量不限制其大小：地图会增长以容纳其中存储的项目数量， `nil` 地图除​​外。 `nil` 映射相当于空映射，只不过不能添加任何元素。

### Channel types  渠道类型

A channel provides a mechanism for [concurrently executing functions](https://go.dev/ref/spec#Go_statements) to communicate by [sending](https://go.dev/ref/spec#Send_statements) and [receiving](https://go.dev/ref/spec#Receive_operator) values of a specified element type. The value of an uninitialized channel is `nil`.  
通道提供了一种并发执行函数的机制，通过发送和接收指定元素类型的值来进行通信。未初始化通道的值为 `nil` 。

ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) [ElementType](https://go.dev/ref/spec#ElementType) .

The optional `<-` operator specifies the channel *direction*, *send* or *receive*. If a direction is given, the channel is *directional*, otherwise it is *bidirectional*. A channel may be constrained only to send or only to receive by [assignment](https://go.dev/ref/spec#Assignment_statements) or explicit [conversion](https://go.dev/ref/spec#Conversions).  
可选的 `<-` 运算符指定通道方向、发送或接收。如果给定方向，则通道是有方向的，否则是双向的。通过分配或显式转换，通道可以被限制为仅发送或仅接收。

chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints

The `<-` operator associates with the leftmost `chan` possible:  
`<-` 运算符与最左边的 `chan` 可能关联：

chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)

A new, initialized channel value can be made using the built-in function [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels), which takes the channel type and an optional *capacity* as arguments:  
可以使用内置函数 `make` 创建新的初始化通道值，该函数将通道类型和可选容量作为参数：

make(chan int, 100)

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A `nil` channel is never ready for communication.  
容量（以元素数量表示）设置通道中缓冲区的大小。如果容量为零或不存在，则通道无缓冲，并且仅当发送方和接收方都准备好时通信才会成功。否则，如果缓冲区未满（发送）或不为空（接收），则通道将被缓冲，并且通信会成功而不会阻塞。 `nil` 通道从未准备好进行通信。

A channel may be closed with the built-in function [`close`](https://go.dev/ref/spec#Close). The multi-valued assignment form of the [receive operator](https://go.dev/ref/spec#Receive_operator) reports whether a received value was sent before the channel was closed.  
可以使用内置函数 `close` 关闭通道。接收操作符的多值分配形式报告在通道关闭之前是否发送了接收到的值。

A single channel may be used in [send statements](https://go.dev/ref/spec#Send_statements), [receive operations](https://go.dev/ref/spec#Receive_operator), and calls to the built-in functions [`cap`](https://go.dev/ref/spec#Length_and_capacity) and [`len`](https://go.dev/ref/spec#Length_and_capacity) by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. For example, if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent.  
单个通道可以用于任意数量的 goroutine 的发送语句、接收操作以及对内置函数 `cap` 和 `len` 的调用，而无需进一步同步。通道充当先进先出队列。例如，如果一个 Goroutine 在通道上发送值，而第二个 Goroutine 接收它们，则这些值将按照发送的顺序接收。

## Properties of types and values

类型和值的属性

### Underlying types 底层类型

Each type `T` has an *underlying type*: If `T` is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is `T` itself. Otherwise, `T`'s underlying type is the underlying type of the type to which `T` refers in its declaration. For a type parameter that is the underlying type of its [type constraint](https://go.dev/ref/spec#Type_constraints), which is always an interface.  
每个类型 `T` 都有一个基础类型：如果 `T` 是预声明的布尔、数字或字符串类型之一，或者类型文字，则相应的基础类型为 `T` 的基础类型是 `T` 在其声明中引用的类型的基础类型。对于作为其类型约束的基础类型的类型参数，它始终是一个接口。

type (
    A1 = string
    A2 = A1
)
type (
    B1 string
    B2 B1
    B3 []B1
    B4 B3
)
func f[P any](x P) { … }

The underlying type of `string`, `A1`, `A2`, `B1`, and `B2` is `string`. The underlying type of `[]B1`, `B3`, and `B4` is `[]B1`. The underlying type of `P` is `interface{}`.  
`string` 、 `A1` 、 `A2` 、 `B1` 和 `B2` 的基础类型是 `string` 、 `B3` 和 `B4` 的基础类型是 `[]B1` 。 `P` 的基础类型是 `interface{}` 。

### Core types 核心类型

Each non-interface type `T` has a *core type*, which is the same as the [underlying type](https://go.dev/ref/spec#Underlying_types) of `T`.  
每个非接口类型 `T` 都有一个核心类型，它与 `T` 的底层类型相同。

An interface `T` has a core type if one of the following conditions is satisfied:  
如果满足以下条件之一，则接口 `T` 具有核心类型：

1. There is a single type `U` which is the [underlying type](https://go.dev/ref/spec#Underlying_types) of all types in the [type set](https://go.dev/ref/spec#Interface_types) of `T`; or  
   有一个类型 `U` ，它是 `T` 类型集中所有类型的基础类型；或者
2. the type set of `T` contains only [channel types](https://go.dev/ref/spec#Channel_types) with identical element type `E`, and all directional channels have the same direction.  
   `T` 的类型集仅包含具有相同元素类型 `E` 的通道类型，并且所有定向通道具有相同的方向。

No other interfaces have a core type.  
没有其他接口具有核心类型。

The core type of an interface is, depending on the condition that is satisfied, either:  
根据满足的条件，接口的核心类型是：

1. the type `U`; or  
   类型 `U` ；或者
2. the type `chan E` if `T` contains only bidirectional channels, or the type `chan<- E` or `<-chan E` depending on the direction of the directional channels present.  
   如果 `T` 仅包含双向通道，则为 `chan E` 类型，或者根据存在的定向通道的方向，为 `chan<- E` 或 `<-chan E` 类型。

By definition, a core type is never a [defined type](https://go.dev/ref/spec#Type_definitions), [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), or [interface type](https://go.dev/ref/spec#Interface_types).  
根据定义，核心类型绝不是已定义的类型、类型参数或接口类型。

Examples of interfaces with core types:  
具有核心类型的接口示例：

type Celsius float32
type Kelvin  float32
interface{ int }                          // int
interface{ Celsius|Kelvin }               // float32
interface{ ~chan int }                    // chan int
interface{ ~chan int|~chan<- int }        // chan<- int
interface{ ~[]*data; String() string }    // []*data

Examples of interfaces without core types:

interface{}                               // no single underlying type
interface{ Celsius|float64 }              // no single underlying type
interface{ chan int | chan<- string }     // channels have different element types
interface{ <-chan int | chan<- int }      // directional channels have different directions

Some operations ([slice expressions](https://go.dev/ref/spec#Slice_expressions), [`append` and `copy`](https://go.dev/ref/spec#Appending_and_copying_slices)) rely on a slightly more loose form of core types which accept byte slices and strings. Specifically, if there are exactly two types, `[]byte` and `string`, which are the underlying types of all types in the type set of interface `T`, the core type of `T` is called `bytestring`.

Examples of interfaces with `bytestring` core types:

interface{ int }                          // int (same as ordinary core type)
interface{ []byte | string }              // bytestring
interface{ ~[]byte | myString }           // bytestring

Note that `bytestring` is not a real type; it cannot be used to declare variables or compose other types. It exists solely to describe the behavior of some operations that read from a sequence of bytes, which may be a byte slice or a string.  
请注意， `bytestring` 不是真实类型；它不能用于声明变量或组成其他类型。它的存在只是为了描述从字节序列（可能是字节片或字符串）读取的某些操作的行为。

### Type identity 类型标识

Two types are either *identical* or *different*.  
两种类型要么相同，要么不同。

A [named type](https://go.dev/ref/spec#Types) is always different from any other type. Otherwise, two types are identical if their [underlying](https://go.dev/ref/spec#Types) type literals are structurally equivalent; that is, they have the same literal structure and corresponding components have identical types. In detail:  
命名类型始终不同于任何其他类型。否则，如果两个类型的底层类型文字在结构上等效，则它们是相同的；也就是说，它们具有相同的字面结构，并且相应的组件具有相同的类型。详细地：

- Two array types are identical if they have identical element types and the same array length.  
  如果两个数组类型具有相同的元素类型和相同的数组长度，则它们是相同的。
- Two slice types are identical if they have identical element types.  
  如果两个切片类型具有相同的元素类型，则它们是相同的。
- Two struct types are identical if they have the same sequence of fields, and if corresponding fields have the same names, and identical types, and identical tags. [Non-exported](https://go.dev/ref/spec#Exported_identifiers) field names from different packages are always different.  
  如果两个结构类型具有相同的字段序列，并且相应的字段具有相同的名称、相同的类型和相同的标记，则它们是相同的。不同包中的非导出字段名称总是不同的。
- Two pointer types are identical if they have identical base types.  
  如果两个指针类型具有相同的基类型，则它们是相同的。
- Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match.  
  如果两个函数类型具有相同数量的参数和结果值，相应的参数和结果类型相同，并且两个函数都是可变参数或都不是，则它们是相同的。参数和结果名称不需要匹配。
- Two interface types are identical if they define the same type set.  
  如果两个接口类型定义相同的类型集，则它们是相同的。
- Two map types are identical if they have identical key and element types.  
  如果两个映射类型具有相同的键和元素类型，则它们是相同的。
- Two channel types are identical if they have identical element types and the same direction.  
  如果两个通道类型具有相同的元素类型和相同的方向，则它们是相同的。
- Two [instantiated](https://go.dev/ref/spec#Instantiations) types are identical if their defined types and all type arguments are identical.  
  如果两个实例化类型的定义类型和所有类型参数都相同，则它们是相同的。

Given the declarations  鉴于声明

type (
    A0 = []string
    A1 = A0
    A2 = struct{ a, b int }
    A3 = int
    A4 = func(A3, float64) *A0
    A5 = func(x int, _ float64) *[]string
    B0 A0
    B1 []string
    B2 struct{ a, b int }
    B3 struct{ a, c int }
    B4 func(int, float64) *B0
    B5 func(x int, y float64) *A1
    C0 = B0
    D0[P1, P2 any] struct{ x P1; y P2 }
    E0 = D0[int, string]
)

these types are identical:  
这些类型是相同的：

A0, A1, and []string
A2 and struct{ a, b int }
A3 and int
A4, func(int, float64) *[]string, and A5
B0 and C0
D0[int, string] and E0
[]int and []int
struct{ a, b *B5 } and struct{ a, b *B5 }
func(x int, y float64) *[]string, func(int, float64) (result *[]string), and A5

`B0` and `B1` are different because they are new types created by distinct [type definitions](https://go.dev/ref/spec#Type_definitions); `func(int, float64) *B0` and `func(x int, y float64) *[]string` are different because `B0` is different from `[]string`; and `P1` and `P2` are different because they are different type parameters. `D0[int, string]` and `struct{ x int; y string }` are different because the former is an [instantiated](https://go.dev/ref/spec#Instantiations) defined type while the latter is a type literal (but they are still [assignable](https://go.dev/ref/spec#Assignability)).  
`B0` 和 `B1` 是不同的，因为它们是由不同类型定义创建的新​​类型； `func(int, float64) *B0` 和 `func(x int, y float64) *[]string` 不同，因为 `B0` 与 `[]string` 不同； `P1` 和 `P2` 是不同的，因为它们是不同的类型参数。 `D0[int, string]` 和 `struct{ x int; y string }` 是不同的，因为前者是实例化的定义类型，而后者是类型文字（但它们仍然是可赋值的）。

### Assignability 可分配性

A value `x` of type `V` is *assignable* to a [variable](https://go.dev/ref/spec#Variables) of type `T` ("`x` is assignable to `T`") if one of the following conditions applies:  
`V` 类型的值 `x` 可分配给 `T` 类型的变量（“ `x` 可分配给 `T` ") 如果满足以下条件之一：

- `V` and `T` are identical.  
  `V` 和 `T` 相同。
- `V` and `T` have identical [underlying types](https://go.dev/ref/spec#Underlying_types) but are not type parameters and at least one of `V` or `T` is not a [named type](https://go.dev/ref/spec#Types).  
  `V` 和 `T` 具有相同的基础类型，但不是类型参数，并且 `V` 或 `T` 至少其中之一不是命名类型。
- `V` and `T` are channel types with identical element types, `V` is a bidirectional channel, and at least one of `V` or `T` is not a [named type](https://go.dev/ref/spec#Types).  
  `V` 和 `T` 是具有相同元素类型的通道类型， `V` 是双向通道， `V` 或 不是命名类型。
- `T` is an interface type, but not a type parameter, and `x` [implements](https://go.dev/ref/spec#Implementing_an_interface) `T`.  
  `T` 是接口类型，但不是类型参数，并且 `x` 实现 `T` 。
- `x` is the predeclared identifier `nil` and `T` is a pointer, function, slice, map, channel, or interface type, but not a type parameter.  
  `x` 是预声明的标识符 `nil` ， `T` 是指针、函数、切片、映射、通道或接口类型，但不是类型参数。
- `x` is an untyped [constant](https://go.dev/ref/spec#Constants) [representable](https://go.dev/ref/spec#Representability) by a value of type `T`.  
  `x` 是一个无类型常量，由 `T` 类型的值表示。

Additionally, if `x`'s type `V` or `T` are type parameters, `x` is assignable to a variable of type `T` if one of the following conditions applies:

- `x` is the predeclared identifier `nil`, `T` is a type parameter, and `x` is assignable to each type in `T`'s type set.
- `V` is not a [named type](https://go.dev/ref/spec#Types), `T` is a type parameter, and `x` is assignable to each type in `T`'s type set.
- `V` is a type parameter and `T` is not a named type, and values of each type in `V`'s type set are assignable to `T`.

### Representability

A [constant](https://go.dev/ref/spec#Constants) `x` is *representable* by a value of type `T`, where `T` is not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), if one of the following conditions applies:

- `x` is in the set of values [determined](https://go.dev/ref/spec#Types) by `T`.
- `T` is a [floating-point type](https://go.dev/ref/spec#Numeric_types) and `x` can be rounded to `T`'s precision without overflow. Rounding uses IEEE 754 round-to-even rules but with an IEEE negative zero further simplified to an unsigned zero. Note that constant values never result in an IEEE negative zero, NaN, or infinity.
- `T` is a complex type, and `x`'s [components](https://go.dev/ref/spec#Complex_numbers) `real(x)` and `imag(x)` are representable by values of `T`'s component type (`float32` or `float64`).

If `T` is a type parameter, `x` is representable by a value of type `T` if `x` is representable by a value of each type in `T`'s type set.

x                   T           x is representable by a value of T because
'a'                 byte        97 is in the set of byte values
97                  rune        rune is an alias for int32, and 97 is in the set of 32-bit integers
"foo"               string      "foo" is in the set of string values
1024                int16       1024 is in the set of 16-bit integers
42.0                byte        42 is in the set of unsigned 8-bit integers
1e10                uint64      10000000000 is in the set of unsigned 64-bit integers
2.718281828459045   float32     2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
-1e-1000            float64     -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
0i                  int         0 is an integer value
(42 + 0i)           float32     42.0 (with zero imaginary part) is in the set of float32 values

x                   T           x is not representable by a value of T because
0                   bool        0 is not in the set of boolean values
'a'                 string      'a' is a rune, it is not in the set of string values
1024                byte        1024 is not in the set of unsigned 8-bit integers
-1                  uint16      -1 is not in the set of unsigned 16-bit integers
1.1                 int         1.1 is not an integer value
42i                 float32     (0 + 42i) is not in the set of float32 values
1e1000              float64     1e1000 overflows to IEEE +Inf after rounding

### Method sets

The *method set* of a type determines the methods that can be [called](https://go.dev/ref/spec#Calls) on an [operand](https://go.dev/ref/spec#Operands) of that type. Every type has a (possibly empty) method set associated with it:

- The method set of a [defined type](https://go.dev/ref/spec#Type_definitions) `T` consists of all [methods](https://go.dev/ref/spec#Method_declarations) declared with receiver type `T`.
- The method set of a pointer to a defined type `T` (where `T` is neither a pointer nor an interface) is the set of all methods declared with receiver `*T` or `T`.
- The method set of an [interface type](https://go.dev/ref/spec#Interface_types) is the intersection of the method sets of each type in the interface's [type set](https://go.dev/ref/spec#Interface_types) (the resulting method set is usually just the set of declared methods in the interface).

Further rules apply to structs (and pointer to structs) containing embedded fields, as described in the section on [struct types](https://go.dev/ref/spec#Struct_types). Any other type has an empty method set.

In a method set, each method must have a [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers) non-[blank](https://go.dev/ref/spec#Blank_identifier) [method name](https://go.dev/ref/spec#MethodName).

## Blocks

A *block* is a possibly empty sequence of declarations and statements within matching brace brackets.

Block = "{" [StatementList](https://go.dev/ref/spec#StatementList) "}" .
StatementList = { [Statement](https://go.dev/ref/spec#Statement) ";" } .

In addition to explicit blocks in the source code, there are implicit blocks:

1. The *universe block* encompasses all Go source text.
2. Each [package](https://go.dev/ref/spec#Packages) has a *package block* containing all Go source text for that package.
3. Each file has a *file block* containing all Go source text in that file.
4. Each ["if"](https://go.dev/ref/spec#If_statements), ["for"](https://go.dev/ref/spec#For_statements), and ["switch"](https://go.dev/ref/spec#Switch_statements) statement is considered to be in its own implicit block.
5. Each clause in a ["switch"](https://go.dev/ref/spec#Switch_statements) or ["select"](https://go.dev/ref/spec#Select_statements) statement acts as an implicit block.

Blocks nest and influence [scoping](https://go.dev/ref/spec#Declarations_and_scope).

## Declarations and scope

A *declaration* binds a non-[blank](https://go.dev/ref/spec#Blank_identifier) identifier to a [constant](https://go.dev/ref/spec#Constant_declarations), [type](https://go.dev/ref/spec#Type_declarations), [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), [variable](https://go.dev/ref/spec#Variable_declarations), [function](https://go.dev/ref/spec#Function_declarations), [label](https://go.dev/ref/spec#Labeled_statements), or [package](https://go.dev/ref/spec#Import_declarations). Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

The [blank identifier](https://go.dev/ref/spec#Blank_identifier) may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier `init` may only be used for [`init` function](https://go.dev/ref/spec#Package_initialization) declarations, and like the blank identifier it does not introduce a new binding.

Declaration   = [ConstDecl](https://go.dev/ref/spec#ConstDecl) | [TypeDecl](https://go.dev/ref/spec#TypeDecl) | [VarDecl](https://go.dev/ref/spec#VarDecl) .
TopLevelDecl  = [Declaration](https://go.dev/ref/spec#Declaration) | [FunctionDecl](https://go.dev/ref/spec#FunctionDecl) | [MethodDecl](https://go.dev/ref/spec#MethodDecl) .

The *scope* of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

Go is lexically scoped using [blocks](https://go.dev/ref/spec#Blocks):

1. The scope of a [predeclared identifier](https://go.dev/ref/spec#Predeclared_identifiers) is the universe block.
2. The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.  
   表示在顶层（任何函数外部）声明的常量、类型、变量或函数（但不是方法）的标识符的范围是包块。
3. The scope of the package name of an imported package is the file block of the file containing the import declaration.  
   导入包的包名范围是包含导入声明的文件的文件块。
4. The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.  
   表示方法接收者、函数参数或结果变量的标识符的范围是函数体。
5. The scope of an identifier denoting a type parameter of a function or declared by a method receiver begins after the name of the function and ends at the end of the function body.  
   表示函数类型参数或由方法接收者声明的标识符的范围从函数名称之后开始，到函数体末尾结束。
6. The scope of an identifier denoting a type parameter of a type begins after the name of the type and ends at the end of the TypeSpec.  
   表示类型的类型参数的标识符的范围从类型名称之后开始，到 TypeSpec 末尾结束。
7. The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.  
   函数内部声明的常量或变量标识符的范围从 ConstSpec 或 VarSpec（用于短变量声明的 ShortVarDecl）的末尾开始，到最内层包含块的末尾结束。
8. The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.  
   函数内部声明的类型标识符的范围从 TypeSpec 中的标识符开始，到最内层包含块的末尾结束。

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.  
在块中声明的标识符可以在内部块中重新声明。当内部声明的标识符在范围内时，它表示内部声明所声明的实体。

The [package clause](https://go.dev/ref/spec#Package_clause) is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same [package](https://go.dev/ref/spec#Packages) and to specify the default package name for import declarations.  
包条款不是声明；包名称不会出现在任何范围内。其目的是识别属于同一包的文件并指定导入声明的默认包名称。

### Label scopes 标签范围

Labels are declared by [labeled statements](https://go.dev/ref/spec#Labeled_statements) and are used in the ["break"](https://go.dev/ref/spec#Break_statements), ["continue"](https://go.dev/ref/spec#Continue_statements), and ["goto"](https://go.dev/ref/spec#Goto_statements) statements. It is illegal to define a label that is never used. In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.  
标签由带标签的语句声明，并用在“break”、“continue”和“goto”语句中。定义从未使用过的标签是非法的。与其他标识符相比，标签不是块作用域的，并且不会与非标签的标识符发生冲突。标签的作用域是声明它的函数体，不包括任何嵌套函数的体。

### Blank identifier 空白标识符

The *blank identifier* is represented by the underscore character `_`. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in [declarations](https://go.dev/ref/spec#Declarations_and_scope), as an [operand](https://go.dev/ref/spec#Operands), and in [assignment statements](https://go.dev/ref/spec#Assignment_statements).  
空白标识符由下划线字符 `_` 表示。它用作匿名占位符，而不是常规（非空白）标识符，并且在声明、操作数和赋值语句中具有特殊含义。

### Predeclared identifiers 预声明标识符

The following identifiers are implicitly declared in the [universe block](https://go.dev/ref/spec#Blocks) [[Go 1.18](https://go.dev/ref/spec#Go_1.18)] [[Go 1.21](https://go.dev/ref/spec#Go_1.21)]:  
以下标识符在 Universe 块中隐式声明 [Go 1.18] [Go 1.21]：

Types:
    any bool byte comparable
    complex64 complex128 error float32 float64
    int int8 int16 int32 int64 rune string
    uint uint8 uint16 uint32 uint64 uintptr
Constants:
    true false iota
Zero value:
    nil
Functions:
    append cap clear close complex copy delete imag len
    make max min new panic print println real recover

### Exported identifiers 导出的标识符

An identifier may be *exported* to permit access to it from another package. An identifier is exported if both:  
可以导出标识符以允许从另一个包访问它。如果满足以下条件，则导出标识符：

1. the first character of the identifier's name is a Unicode uppercase letter (Unicode character category Lu); and
2. the identifier is declared in the [package block](https://go.dev/ref/spec#Blocks) or it is a [field name](https://go.dev/ref/spec#Struct_types) or [method name](https://go.dev/ref/spec#MethodName).

All other identifiers are not exported.

### Uniqueness of identifiers

Given a set of identifiers, an identifier is called *unique* if it is *different* from every other in the set. Two identifiers are different if they are spelled differently, or if they appear in different [packages](https://go.dev/ref/spec#Packages) and are not [exported](https://go.dev/ref/spec#Exported_identifiers). Otherwise, they are the same.

### Constant declarations

A constant declaration binds a list of identifiers (the names of the constants) to the values of a list of [constant expressions](https://go.dev/ref/spec#Constant_expressions). The number of identifiers must be equal to the number of expressions, and the *n*th identifier on the left is bound to the value of the *n*th expression on the right.

ConstDecl      = "const" ( [ConstSpec](https://go.dev/ref/spec#ConstSpec) | "(" { [ConstSpec](https://go.dev/ref/spec#ConstSpec) ";" } ")" ) .
ConstSpec      = [IdentifierList](https://go.dev/ref/spec#IdentifierList) [ [ [Type](https://go.dev/ref/spec#Type) ] "=" [ExpressionList](https://go.dev/ref/spec#ExpressionList) ] .

IdentifierList = [identifier](https://go.dev/ref/spec#identifier) { "," [identifier](https://go.dev/ref/spec#identifier) } .
ExpressionList = [Expression](https://go.dev/ref/spec#Expression) { "," [Expression](https://go.dev/ref/spec#Expression) } .

If the type is present, all constants take the type specified, and the expressions must be [assignable](https://go.dev/ref/spec#Assignability) to that type, which must not be a type parameter. If the type is omitted, the constants take the individual types of the corresponding expressions. If the expression values are untyped [constants](https://go.dev/ref/spec#Constants), the declared constants remain untyped and the constant identifiers denote the constant values. For instance, if the expression is a floating-point literal, the constant identifier denotes a floating-point constant, even if the literal's fractional part is zero.

const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
    size int64 = 1024
    eof        = -1  // untyped integer constant
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0

Within a parenthesized `const` declaration list the expression list may be omitted from any but the first ConstSpec. Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any. Omitting the list of expressions is therefore equivalent to repeating the previous list. The number of identifiers must be equal to the number of expressions in the previous list. Together with the [`iota` constant generator](https://go.dev/ref/spec#Iota) this mechanism permits light-weight declaration of sequential values:

const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Partyday
    numberOfDays  // this constant is not exported
)

### Iota

Within a [constant declaration](https://go.dev/ref/spec#Constant_declarations), the predeclared identifier `iota` represents successive untyped integer [constants](https://go.dev/ref/spec#Constants). Its value is the index of the respective [ConstSpec](https://go.dev/ref/spec#ConstSpec) in that constant declaration, starting at zero. It can be used to construct a set of related constants:

const (
    c0 = iota  // c0 == 0
    c1 = iota  // c1 == 1
    c2 = iota  // c2 == 2
)
const (
    a = 1 << iota  // a == 1  (iota == 0)
    b = 1 << iota  // b == 2  (iota == 1)
    c = 3          // c == 3  (iota == 2, unused)
    d = 1 << iota  // d == 8  (iota == 3)
)
const (
    u         = iota * 42  // u == 0     (untyped integer constant)
    v float64 = iota * 42  // v == 42.0  (float64 constant)
    w         = iota * 42  // w == 84    (untyped integer constant)
)
const x = iota  // x == 0
const y = iota  // y == 0

By definition, multiple uses of `iota` in the same ConstSpec all have the same value:  
根据定义，在同一个 ConstSpec 中多次使用 `iota` 都具有相同的值：

const (
    bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0  (iota == 0)
    bit1, mask1                           // bit1 == 2, mask1 == 1  (iota == 1)
    _, _                                  //                        (iota == 2, unused)
    bit3, mask3                           // bit3 == 8, mask3 == 7  (iota == 3)
)

This last example exploits the [implicit repetition](https://go.dev/ref/spec#Constant_declarations) of the last non-empty expression list.  
最后一个示例利用了最后一个非空表达式列表的隐式重复。

### Type declarations 类型声明

A type declaration binds an identifier, the *type name*, to a [type](https://go.dev/ref/spec#Types). Type declarations come in two forms: alias declarations and type definitions.  
类型声明将标识符（类型名称）绑定到类型。类型声明有两种形式：别名声明和类型定义。

TypeDecl = "type" ( [TypeSpec](https://go.dev/ref/spec#TypeSpec) | "(" { [TypeSpec](https://go.dev/ref/spec#TypeSpec) ";" } ")" ) .
TypeSpec = [AliasDecl](https://go.dev/ref/spec#AliasDecl) | [TypeDef](https://go.dev/ref/spec#TypeDef) .

#### Alias declarations 别名声明

An alias declaration binds an identifier to the given type [[Go 1.9](https://go.dev/ref/spec#Go_1.9)].  
别名声明将标识符绑定到给定类型 [Go 1.9]。

AliasDecl = [identifier](https://go.dev/ref/spec#identifier) "=" [Type](https://go.dev/ref/spec#Type) .

Within the [scope](https://go.dev/ref/spec#Declarations_and_scope) of the identifier, it serves as an *alias* for the type.  
在标识符的范围内，它充当类型的别名。

type (
    nodeList = []*Node  // nodeList and []*Node are identical types
    Polar    = polar    // Polar and polar denote identical types
)

#### Type definitions 类型定义

A type definition creates a new, distinct type with the same [underlying type](https://go.dev/ref/spec#Underlying_types) and operations as the given type and binds an identifier, the *type name*, to it.  
类型定义创建一个新的、不同的类型，其具有与给定类型相同的基础类型和操作，并将标识符（类型名称）绑定到它。

TypeDef = [identifier](https://go.dev/ref/spec#identifier) [ [TypeParameters](https://go.dev/ref/spec#TypeParameters) ] [Type](https://go.dev/ref/spec#Type) .

The new type is called a *defined type*. It is [different](https://go.dev/ref/spec#Type_identity) from any other type, including the type it is created from.  
新类型称为定义类型。它不同于任何其他类型，包括创建它的类型。

type (
    Point struct{ x, y float64 }  // Point and struct{ x, y float64 } are different types
    polar Point                   // polar and Point denote different types
)
type TreeNode struct {
    left, right *TreeNode
    value any
}
type Block interface {
    BlockSize() int
    Encrypt(src, dst []byte)
    Decrypt(src, dst []byte)
}

A defined type may have [methods](https://go.dev/ref/spec#Method_declarations) associated with it. It does not inherit any methods bound to the given type, but the [method set](https://go.dev/ref/spec#Method_sets) of an interface type or of elements of a composite type remains unchanged:

// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }
// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex
// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex
// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
    Mutex
}
// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block

Type definitions may be used to define different boolean, numeric, or string types and associate methods with them:

type TimeZone int
const (
    EST TimeZone = -(5 + iota)
    CST
    MST
    PST
)
func (tz TimeZone) String() string {
    return fmt.Sprintf("GMT%+dh", tz)
}

If the type definition specifies [type parameters](https://go.dev/ref/spec#Type_parameter_declarations), the type name denotes a *generic type*. Generic types must be [instantiated](https://go.dev/ref/spec#Instantiations) when they are used.

type List[T any] struct {
    next  *List[T]
    value T
}

In a type definition the given type cannot be a type parameter.

type T[P any] P    // illegal: P is a type parameter
func f[T any]() {
    type L T   // illegal: T is a type parameter declared by the enclosing function
}

A generic type may also have [methods](https://go.dev/ref/spec#Method_declarations) associated with it. In this case, the method receivers must declare the same number of type parameters as present in the generic type definition.  
泛型类型也可能具有与其关联的方法。在这种情况下，方法接收者必须声明与泛型类型定义中存在的相同数量的类型参数。

// The method Len returns the number of elements in the linked list l.
func (l *List[T]) Len() int  { … }

### Type parameter declarations

类型参数声明

A type parameter list declares the *type parameters* of a generic function or type declaration. The type parameter list looks like an ordinary [function parameter list](https://go.dev/ref/spec#Function_types) except that the type parameter names must all be present and the list is enclosed in square brackets rather than parentheses [[Go 1.18](https://go.dev/ref/spec#Go_1.18)].  
类型参数列表声明泛型函数或类型声明的类型参数。类型参数列表看起来像普通的函数参数列表，只是类型参数名称必须全部存在并且列表用方括号而不是圆括号括起来 [Go 1.18]。

TypeParameters  = "[" [TypeParamList](https://go.dev/ref/spec#TypeParamList) [ "," ] "]" .
TypeParamList   = [TypeParamDecl](https://go.dev/ref/spec#TypeParamDecl) { "," [TypeParamDecl](https://go.dev/ref/spec#TypeParamDecl) } .
TypeParamDecl   = [IdentifierList](https://go.dev/ref/spec#IdentifierList) [TypeConstraint](https://go.dev/ref/spec#TypeConstraint) .

All non-blank names in the list must be unique. Each name declares a type parameter, which is a new and different [named type](https://go.dev/ref/spec#Types) that acts as a placeholder for an (as of yet) unknown type in the declaration. The type parameter is replaced with a *type argument* upon [instantiation](https://go.dev/ref/spec#Instantiations) of the generic function or type.  
列表中的所有非空白名称必须是唯一的。每个名称都声明一个类型参数，它是一个新的、不同的命名类型，充当声明中（迄今为止）未知类型的占位符。在泛型函数或类型实例化时，类型形参将替换为类型实参。

[P any]
[S interface{ ~[]byte|string }]
[S ~[]E, E any]
[P Constraint[int]]
[_ any]

Just as each ordinary function parameter has a parameter type, each type parameter has a corresponding (meta-)type which is called its [*type constraint*](https://go.dev/ref/spec#Type_constraints).  
正如每个普通函数参数都有一个参数类型一样，每个类型参数都有一个相应的（元）类型，称为类型约束。

A parsing ambiguity arises when the type parameter list for a generic type declares a single type parameter `P` with a constraint `C` such that the text `P C` forms a valid expression:  
当泛型类型的类型参数列表声明带有约束 `C` 的单个类型参数 `P` 以使文本 `P C` 形成有效表达式时，会出现解析歧义:

type T[P *C] …
type T[P (C)] …
type T[P *C|Q] …
…

In these rare cases, the type parameter list is indistinguishable from an expression and the type declaration is parsed as an array type declaration. To resolve the ambiguity, embed the constraint in an [interface](https://go.dev/ref/spec#Interface_types) or use a trailing comma:  
在这些罕见的情况下，类型参数列表与表达式无法区分，并且类型声明被解析为数组类型声明。要解决歧义，请将约束嵌入接口中或使用尾随逗号：

type T[P interface{*C}] …
type T[P *C,] …

Type parameters may also be declared by the receiver specification of a [method declaration](https://go.dev/ref/spec#Method_declarations) associated with a generic type.  
类型参数也可以由与泛型类型相关联的方法声明的接收者规范来声明。

Within a type parameter list of a generic type `T`, a type constraint may not (directly, or indirectly through the type parameter list of another generic type) refer to `T`.  
在泛型类型 `T` 的类型参数列表中，类型约束不能（直接或通过另一个泛型类型的类型参数列表间接）引用 `T` 。

type T1[P T1[P]] …                    // illegal: T1 refers to itself
type T2[P interface{ T2[int] }] …     // illegal: T2 refers to itself
type T3[P interface{ m(T3[int])}] …   // illegal: T3 refers to itself
type T4[P T5[P]] …                    // illegal: T4 refers to T5 and
type T5[P T4[P]] …                    //          T5 refers to T4
type T6[P int] struct{ f *T6[P] }     // ok: reference to T6 is not in type parameter list

#### Type constraints 类型限制

A *type constraint* is an [interface](https://go.dev/ref/spec#Interface_types) that defines the set of permissible type arguments for the respective type parameter and controls the operations supported by values of that type parameter [[Go 1.18](https://go.dev/ref/spec#Go_1.18)].  
类型约束是一个接口，它定义了相应类型参数的允许类型参数集，并控制该类型参数的值支持的操作[Go 1.18]。

TypeConstraint = [TypeElem](https://go.dev/ref/spec#TypeElem) .

If the constraint is an interface literal of the form `interface{E}` where `E` is an embedded [type element](https://go.dev/ref/spec#Interface_types) (not a method), in a type parameter list the enclosing `interface{ … }` may be omitted for convenience:  
如果约束是 `interface{E}` 形式的接口文字，其中 `E` 是嵌入类型元素（不是方法），则在类型参数列表中包含 `interface{ … }` 为了方便起见可以省略：

[T []P]                      // = [T interface{[]P}]
[T ~int]                     // = [T interface{~int}]
[T int|string]               // = [T interface{int|string}]
type Constraint ~int         // illegal: ~int is not in a type parameter list

The [predeclared](https://go.dev/ref/spec#Predeclared_identifiers) [interface type](https://go.dev/ref/spec#Interface_types) `comparable` denotes the set of all non-interface types that are [strictly comparable](https://go.dev/ref/spec#Comparison_operators) [[Go 1.18](https://go.dev/ref/spec#Go_1.18)].  
预先声明的接口类型 `comparable` 表示严格可比较的所有非接口类型的集合[Go 1.18]。

Even though interfaces that are not type parameters are [comparable](https://go.dev/ref/spec#Comparison_operators), they are not strictly comparable and therefore they do not implement `comparable`. However, they [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) `comparable`.  
尽管不是类型参数的接口是可比较的，但它们并不严格可比，因此它们不实现 `comparable` 。然而，它们满足 `comparable` 。

int                          // implements comparable (int is strictly comparable)
[]byte                       // does not implement comparable (slices cannot be compared)
interface{}                  // does not implement comparable (see above)
interface{ ~int | ~string }  // type parameter only: implements comparable (int, string types are strictly comparable)
interface{ comparable }      // type parameter only: implements comparable (comparable implements itself)
interface{ ~int | ~[]byte }  // type parameter only: does not implement comparable (slices are not comparable)
interface{ ~struct{ any } }  // type parameter only: does not implement comparable (field any is not strictly comparable)

The `comparable` interface and interfaces that (directly or indirectly) embed `comparable` may only be used as type constraints. They cannot be the types of values or variables, or components of other, non-interface types.  
`comparable` 接口和（直接或间接）嵌入 `comparable` 的接口只能用作类型约束。它们不能是值或变量的类型，也不能是其他非接口类型的组件。

#### Satisfying a type constraint

满足类型约束

A type argument `T` *satisfies* a type constraint `C` if `T` is an element of the type set defined by `C`; i.e., if `T` [implements](https://go.dev/ref/spec#Implementing_an_interface) `C`. As an exception, a [strictly comparable](https://go.dev/ref/spec#Comparison_operators) type constraint may also be satisfied by a [comparable](https://go.dev/ref/spec#Comparison_operators) (not necessarily strictly comparable) type argument [[Go 1.20](https://go.dev/ref/spec#Go_1.20)]. More precisely:  
如果 `T` 是 `C` 定义的类型集的元素，则类型参数 `T` 满足类型约束 `C` ；即，如果 `T` 实现 `C` 。作为例外，严格可比较类型约束也可以通过可比较（不一定严格可比较）类型参数来满足[Go 1.20]。更确切地说：

A type T *satisfies* a constraint `C` if

- `T` [implements](https://go.dev/ref/spec#Implementing_an_interface) `C`; or
- `C` can be written in the form `interface{ comparable; E }`, where `E` is a [basic interface](https://go.dev/ref/spec#Basic_interfaces) and `T` is [comparable](https://go.dev/ref/spec#Comparison_operators) and implements `E`.

type argument      type constraint                // constraint satisfaction
int                interface{ ~int }              // satisfied: int implements interface{ ~int }
string             comparable                     // satisfied: string implements comparable (string is strictly comparable)
[]byte             comparable                     // not satisfied: slices are not comparable
any                interface{ comparable; int }   // not satisfied: any does not implement interface{ int }
any                comparable                     // satisfied: any is comparable and implements the basic interface any
struct{f any}      comparable                     // satisfied: struct{f any} is comparable and implements the basic interface any
any                interface{ comparable; m() }   // not satisfied: any does not implement the basic interface interface{ m() }
interface{ m() }   interface{ comparable; m() }   // satisfied: interface{ m() } is comparable and implements the basic interface interface{ m() }

Because of the exception in the constraint satisfaction rule, comparing operands of type parameter type may panic at run-time (even though comparable type parameters are always strictly comparable).

### Variable declarations

A variable declaration creates one or more [variables](https://go.dev/ref/spec#Variables), binds corresponding identifiers to them, and gives each a type and an initial value.

VarDecl     = "var" ( [VarSpec](https://go.dev/ref/spec#VarSpec) | "(" { [VarSpec](https://go.dev/ref/spec#VarSpec) ";" } ")" ) .
VarSpec     = [IdentifierList](https://go.dev/ref/spec#IdentifierList) ( [Type](https://go.dev/ref/spec#Type) [ "=" [ExpressionList](https://go.dev/ref/spec#ExpressionList) ] | "=" [ExpressionList](https://go.dev/ref/spec#ExpressionList) ) .

var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
    i       int
    u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"

If a list of expressions is given, the variables are initialized with the expressions following the rules for [assignment statements](https://go.dev/ref/spec#Assignment_statements). Otherwise, each variable is initialized to its [zero value](https://go.dev/ref/spec#The_zero_value).

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to its [default type](https://go.dev/ref/spec#Constants); if it is an untyped boolean value, it is first implicitly converted to type `bool`. The predeclared value `nil` cannot be used to initialize a variable with no explicit type.

var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal

Implementation restriction: A compiler may make it illegal to declare a variable inside a [function body](https://go.dev/ref/spec#Function_declarations) if the variable is never used.

### Short variable declarations

A *short variable declaration* uses the syntax:

ShortVarDecl = [IdentifierList](https://go.dev/ref/spec#IdentifierList) ":=" [ExpressionList](https://go.dev/ref/spec#ExpressionList) .

It is shorthand for a regular [variable declaration](https://go.dev/ref/spec#Variable_declarations) with initializer expressions but no types:

"var" IdentifierList "=" ExpressionList .

i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
_, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate

Unlike regular variable declarations, a short variable declaration may *redeclare* variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-[blank](https://go.dev/ref/spec#Blank_identifier) variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original. The non-blank variable names on the left side of `:=` must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers).

field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
x, y, x := 1, 2, 3                        // illegal: x repeated on left side of :=

Short variable declarations may appear only inside functions. In some contexts such as the initializers for ["if"](https://go.dev/ref/spec#If_statements), ["for"](https://go.dev/ref/spec#For_statements), or ["switch"](https://go.dev/ref/spec#Switch_statements) statements, they can be used to declare local temporary variables.

### Function declarations

A function declaration binds an identifier, the *function name*, to a function.

FunctionDecl = "func" [FunctionName](https://go.dev/ref/spec#FunctionName) [ [TypeParameters](https://go.dev/ref/spec#TypeParameters) ] [Signature](https://go.dev/ref/spec#Signature) [ [FunctionBody](https://go.dev/ref/spec#FunctionBody) ] .
FunctionName = [identifier](https://go.dev/ref/spec#identifier) .
FunctionBody = [Block](https://go.dev/ref/spec#Block) .

If the function's [signature](https://go.dev/ref/spec#Function_types) declares result parameters, the function body's statement list must end in a [terminating statement](https://go.dev/ref/spec#Terminating_statements).

func IndexRune(s string, r rune) int {
    for i, c := range s {
        if c == r {
            return i
        }
    }
    // invalid: missing return statement
}

If the function declaration specifies [type parameters](https://go.dev/ref/spec#Type_parameter_declarations), the function name denotes a *generic function*. A generic function must be [instantiated](https://go.dev/ref/spec#Instantiations) before it can be called or used as a value.

func min[T ~int|~float64](x, y T) T {
    if x < y {
        return x
    }
    return y
}

A function declaration without type parameters may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine.  
没有类型参数的函数声明可以省略函数体。这样的声明提供了在 Go 外部实现的函数的签名，例如汇编例程。

func flushICache(begin, end uintptr)  // implemented externally

### Method declarations 方法声明

A method is a [function](https://go.dev/ref/spec#Function_declarations) with a *receiver*. A method declaration binds an identifier, the *method name*, to a method, and associates the method with the receiver's *base type*.  
方法是带有接收器的函数。方法声明将标识符（方法名称）绑定到方法，并将该方法与接收者的基类型相关联。

MethodDecl = "func" [Receiver](https://go.dev/ref/spec#Receiver) [MethodName](https://go.dev/ref/spec#MethodName) [Signature](https://go.dev/ref/spec#Signature) [ [FunctionBody](https://go.dev/ref/spec#FunctionBody) ] .
Receiver   = [Parameters](https://go.dev/ref/spec#Parameters) .

The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be a [defined](https://go.dev/ref/spec#Type_definitions) type `T` or a pointer to a defined type `T`, possibly followed by a list of type parameter names `[P1, P2, …]` enclosed in square brackets. `T` is called the receiver *base type*. A receiver base type cannot be a pointer or interface type and it must be defined in the same package as the method. The method is said to be *bound* to its receiver base type and the method name is visible only within [selectors](https://go.dev/ref/spec#Selectors) for type `T` or `*T`.  
接收者是通过方法名称前面的额外参数部分指定的。该参数部分必须声明一个非可变参数，即接收者。其类型必须是已定义类型 `T` 或指向已定义类型 `T` 的指针，后面可能是方括号中括起来的类型参数名称 `[P1, P2, …]` 列表。 `T` 称为接收者基本类型。接收者基类型不能是指针或接口类型，并且必须在与方法相同的包中定义。据说该方法绑定到其接收者基类型，并且方法名称仅在类型 `T` 或 `*T` 的选择器中可见。

A non-[blank](https://go.dev/ref/spec#Blank_identifier) receiver identifier must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers) in the method signature. If the receiver's value is not referenced inside the body of the method, its identifier may be omitted in the declaration. The same applies in general to parameters of functions and methods.  
非空白接收者标识符在方法签名中必须是唯一的。如果方法体内未引用接收者的值，则声明中可以省略其标识符。这同样适用于函数和方法的参数。

For a base type, the non-blank names of methods bound to it must be unique. If the base type is a [struct type](https://go.dev/ref/spec#Struct_types), the non-blank method and field names must be distinct.  
对于基类型，绑定到它的方法的非空名称必须是唯一的。如果基类型是结构类型，则非空方法名称和字段名称必须不同。

Given defined type `Point` the declarations  
给定定义的类型 `Point` 声明

func (p *Point) Length() float64 {
    return math.Sqrt(p.x * p.x + p.y * p.y)
}
func (p *Point) Scale(factor float64) {
    p.x *= factor
    p.y *= factor
}

bind the methods `Length` and `Scale`, with receiver type `*Point`, to the base type `Point`.  
将方法 `Length` 和 `Scale` 与接收器类型 `*Point` 绑定到基本类型 `Point` 。

If the receiver base type is a [generic type](https://go.dev/ref/spec#Type_declarations), the receiver specification must declare corresponding type parameters for the method to use. This makes the receiver type parameters available to the method. Syntactically, this type parameter declaration looks like an [instantiation](https://go.dev/ref/spec#Instantiations) of the receiver base type: the type arguments must be identifiers denoting the type parameters being declared, one for each type parameter of the receiver base type. The type parameter names do not need to match their corresponding parameter names in the receiver base type definition, and all non-blank parameter names must be unique in the receiver parameter section and the method signature. The receiver type parameter constraints are implied by the receiver base type definition: corresponding type parameters have corresponding constraints.  
如果接收者基类型是泛型类型，则接收者规范必须为要使用的方法声明相应的类型参数。这使得接收者类型参数可供该方法使用。从语法上讲，此类型参数声明看起来像接收者基类型的实例化：类型参数必须是表示所声明的类型参数的标识符，一个对应于接收者基类型的每个类型参数。类型参数名称不需要与接收者基本类型定义中相应的参数名称相匹配，并且所有非空白参数名称在接收者参数部分和方法签名中必须是唯一的。接收者类型参数约束由接收者基本类型定义隐含：相应的类型参数具有相应的约束。

type Pair[A, B any] struct {
    a A
    b B
}
func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // receiver declares A, B
func (p Pair[First, _]) First() First  { … }  // receiver declares First, corresponds to A in Pair

## Expressions 表达式

An expression specifies the computation of a value by applying operators and functions to operands.  
表达式通过将运算符和函数应用于操作数来指定值的计算。

### Operands 操作数

Operands denote the elementary values in an expression. An operand may be a literal, a (possibly [qualified](https://go.dev/ref/spec#Qualified_identifiers)) non-[blank](https://go.dev/ref/spec#Blank_identifier) identifier denoting a [constant](https://go.dev/ref/spec#Constant_declarations), [variable](https://go.dev/ref/spec#Variable_declarations), or [function](https://go.dev/ref/spec#Function_declarations), or a parenthesized expression.  
操作数表示表达式中的基本值。操作数可以是文字、表示常量、变量或函数的（可能是限定的）非空白标识符，或者是带括号的表达式。

Operand     = [Literal](https://go.dev/ref/spec#Literal) | [OperandName](https://go.dev/ref/spec#OperandName) [ [TypeArgs](https://go.dev/ref/spec#TypeArgs) ] | "(" [Expression](https://go.dev/ref/spec#Expression) ")" .
Literal     = [BasicLit](https://go.dev/ref/spec#BasicLit) | [CompositeLit](https://go.dev/ref/spec#CompositeLit) | [FunctionLit](https://go.dev/ref/spec#FunctionLit) .
BasicLit    = [int_lit](https://go.dev/ref/spec#int_lit) | [float_lit](https://go.dev/ref/spec#float_lit) | [imaginary_lit](https://go.dev/ref/spec#imaginary_lit) | [rune_lit](https://go.dev/ref/spec#rune_lit) | [string_lit](https://go.dev/ref/spec#string_lit) .
OperandName = [identifier](https://go.dev/ref/spec#identifier) | [QualifiedIdent](https://go.dev/ref/spec#QualifiedIdent) .

An operand name denoting a [generic function](https://go.dev/ref/spec#Function_declarations) may be followed by a list of [type arguments](https://go.dev/ref/spec#Instantiations); the resulting operand is an [instantiated](https://go.dev/ref/spec#Instantiations) function.

The [blank identifier](https://go.dev/ref/spec#Blank_identifier) may appear as an operand only on the left-hand side of an [assignment statement](https://go.dev/ref/spec#Assignment_statements).

Implementation restriction: A compiler need not report an error if an operand's type is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) with an empty [type set](https://go.dev/ref/spec#Interface_types). Functions with such type parameters cannot be [instantiated](https://go.dev/ref/spec#Instantiations); any attempt will lead to an error at the instantiation site.

### Qualified identifiers

A *qualified identifier* is an identifier qualified with a package name prefix. Both the package name and the identifier must not be [blank](https://go.dev/ref/spec#Blank_identifier).

QualifiedIdent = [PackageName](https://go.dev/ref/spec#PackageName) "." [identifier](https://go.dev/ref/spec#identifier) .

A qualified identifier accesses an identifier in a different package, which must be [imported](https://go.dev/ref/spec#Import_declarations). The identifier must be [exported](https://go.dev/ref/spec#Exported_identifiers) and declared in the [package block](https://go.dev/ref/spec#Blocks) of that package.

math.Sin // denotes the Sin function in package math

### Composite literals

Composite literals construct new composite values each time they are evaluated. They consist of the type of the literal followed by a brace-bound list of elements. Each element may optionally be preceded by a corresponding key.

CompositeLit  = [LiteralType](https://go.dev/ref/spec#LiteralType) [LiteralValue](https://go.dev/ref/spec#LiteralValue) .
LiteralType   = [StructType](https://go.dev/ref/spec#StructType) | [ArrayType](https://go.dev/ref/spec#ArrayType) | "[" "..." "]" [ElementType](https://go.dev/ref/spec#ElementType) |
                [SliceType](https://go.dev/ref/spec#SliceType) | [MapType](https://go.dev/ref/spec#MapType) | [TypeName](https://go.dev/ref/spec#TypeName) [ [TypeArgs](https://go.dev/ref/spec#TypeArgs) ] .
LiteralValue  = "{" [ [ElementList](https://go.dev/ref/spec#ElementList) [ "," ] ] "}" .
ElementList   = [KeyedElement](https://go.dev/ref/spec#KeyedElement) { "," [KeyedElement](https://go.dev/ref/spec#KeyedElement) } .
KeyedElement  = [ [Key](https://go.dev/ref/spec#Key) ":" ] [Element](https://go.dev/ref/spec#Element) .
Key           = [FieldName](https://go.dev/ref/spec#FieldName) | [Expression](https://go.dev/ref/spec#Expression) | [LiteralValue](https://go.dev/ref/spec#LiteralValue) .
FieldName     = [identifier](https://go.dev/ref/spec#identifier) .
Element       = [Expression](https://go.dev/ref/spec#Expression) | [LiteralValue](https://go.dev/ref/spec#LiteralValue) .

The LiteralType's [core type](https://go.dev/ref/spec#Core_types) `T` must be a struct, array, slice, or map type (the syntax enforces this constraint except when the type is given as a TypeName). The types of the elements and keys must be [assignable](https://go.dev/ref/spec#Assignability) to the respective field, element, and key types of type `T`; there is no additional conversion. The key is interpreted as a field name for struct literals, an index for array and slice literals, and a key for map literals. For map literals, all elements must have a key. It is an error to specify multiple elements with the same field name or constant key value. For non-constant map keys, see the section on [evaluation order](https://go.dev/ref/spec#Order_of_evaluation).

For struct literals the following rules apply:

- A key must be a field name declared in the struct type.
- An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
- If any element has a key, every element must have a key.
- An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
- A literal may omit the element list; such a literal evaluates to the zero value for its type.
- It is an error to specify an element for a non-exported field of a struct belonging to a different package.

Given the declarations

type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }

one may write

origin := Point3D{}                            // zero value for Point3D
line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x

For array and slice literals the following rules apply:

- Each element has an associated integer index marking its position in the array.
- An element with a key uses the key as its index. The key must be a non-negative constant [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; and if it is typed it must be of [integer type](https://go.dev/ref/spec#Numeric_types).
- An element without a key uses the previous element's index plus one. If the first element has no key, its index is zero.

[Taking the address](https://go.dev/ref/spec#Address_operators) of a composite literal generates a pointer to a unique [variable](https://go.dev/ref/spec#Variables) initialized with the literal's value.

var pointer *Point3D = &Point3D{y: 1000}

Note that the [zero value](https://go.dev/ref/spec#The_zero_value) for a slice or map type is not the same as an initialized but empty value of the same type. Consequently, taking the address of an empty slice or map composite literal does not have the same effect as allocating a new slice or map value with [new](https://go.dev/ref/spec#Allocation).

p1 := &[]int{}    // p1 points to an initialized, empty slice with value []int{} and length 0
p2 := new([]int)  // p2 points to an uninitialized slice with value nil and length 0

The length of an array literal is the length specified in the literal type. If fewer elements than the length are provided in the literal, the missing elements are set to the zero value for the array element type. It is an error to provide elements with index values outside the index range of the array. The notation `...` specifies an array length equal to the maximum element index plus one.

buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2

A slice literal describes the entire underlying array literal. Thus the length and capacity of a slice literal are the maximum element index plus one. A slice literal has the form

[]T{x1, x2, … xn}

and is shorthand for a slice operation applied to an array:

tmp := [n]T{x1, x2, … xn}
tmp[0 : n]

Within a composite literal of array, slice, or map type `T`, elements or map keys that are themselves composite literals may elide the respective literal type if it is identical to the element or key type of `T`. Similarly, elements or keys that are addresses of composite literals may elide the `&T` when the element or key type is `*T`.

[...]Point{{1.5, -3.5}, {0, 0}}     // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[][]int{{1, 2, 3}, {4, 5}}          // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
[][]Point{{{0, 1}, {1, 2}}}         // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}
map[string]Point{"orig": {0, 0}}    // same as map[string]Point{"orig": Point{0, 0}}
map[Point]string{{0, 0}: "orig"}    // same as map[Point]string{Point{0, 0}: "orig"}
type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}          // same as [2]*Point{&Point{1.5, -3.5}, &Point{}}
[2]PPoint{{1.5, -3.5}, {}}          // same as [2]PPoint{PPoint(&Point{1.5, -3.5}), PPoint(&Point{})}

A parsing ambiguity arises when a composite literal using the TypeName form of the LiteralType appears as an operand between the [keyword](https://go.dev/ref/spec#Keywords) and the opening brace of the block of an "if", "for", or "switch" statement, and the composite literal is not enclosed in parentheses, square brackets, or curly braces. In this rare case, the opening brace of the literal is erroneously parsed as the one introducing the block of statements. To resolve the ambiguity, the composite literal must appear within parentheses.  
当使用 LiteralType 的 TypeName 形式的复合文字作为“if”、“for”或“switch”语句块的关键字和左大括号之间的操作数出现时，会出现解析歧义，并且复合文字为不包含在圆括号、方括号或花括号中。在这种罕见的情况下，文字的左大括号被错误地解析为引入语句块的大括号。为了解决歧义，复合文字必须出现在括号内。

if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }

Examples of valid array, slice, and map literals:  
有效数组、切片和映射文字的示例：

// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}
// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}
// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}
// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
    "C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
    "G0": 24.50, "A0": 27.50, "B0": 30.87,
}

### Function literals 函数字面量

A function literal represents an anonymous [function](https://go.dev/ref/spec#Function_declarations). Function literals cannot declare type parameters.  
函数字面量表示匿名函数。函数文字不能声明类型参数。

FunctionLit = "func" [Signature](https://go.dev/ref/spec#Signature) [FunctionBody](https://go.dev/ref/spec#FunctionBody) .

func(a, b int, z float64) bool { return a*b < int(z) }

A function literal can be assigned to a variable or invoked directly.  
函数文字可以分配给变量或直接调用。

f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)

Function literals are *closures*: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.  
函数字面量是闭包：它们可以引用周围函数中定义的变量。然后，这些变量在周围的函数和函数文字之间共享，并且只要可访问，它们就会一直存在。

### Primary expressions 主要表达方式

Primary expressions are the operands for unary and binary expressions.  
主表达式是一元和二元表达式的操作数。

PrimaryExpr =
    [Operand](https://go.dev/ref/spec#Operand) |
    [Conversion](https://go.dev/ref/spec#Conversion) |
    [MethodExpr](https://go.dev/ref/spec#MethodExpr) |
    [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) [Selector](https://go.dev/ref/spec#Selector) |
    [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) [Index](https://go.dev/ref/spec#Index) |
    [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) [Slice](https://go.dev/ref/spec#Slice) |
    [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) [TypeAssertion](https://go.dev/ref/spec#TypeAssertion) |
    [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) [Arguments](https://go.dev/ref/spec#Arguments) .

Selector       = "." [identifier](https://go.dev/ref/spec#identifier) .
Index          = "[" [Expression](https://go.dev/ref/spec#Expression) [ "," ] "]" .
Slice          = "[" [ [Expression](https://go.dev/ref/spec#Expression) ] ":" [ [Expression](https://go.dev/ref/spec#Expression) ] "]" |
                 "[" [ [Expression](https://go.dev/ref/spec#Expression) ] ":" [Expression](https://go.dev/ref/spec#Expression) ":" [Expression](https://go.dev/ref/spec#Expression) "]" .
TypeAssertion  = "." "(" [Type](https://go.dev/ref/spec#Type) ")" .
Arguments      = "(" [ ( [ExpressionList](https://go.dev/ref/spec#ExpressionList) | [Type](https://go.dev/ref/spec#Type) [ "," [ExpressionList](https://go.dev/ref/spec#ExpressionList) ] ) [ "..." ] [ "," ] ] ")" .

x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()

### Selectors

For a [primary expression](https://go.dev/ref/spec#Primary_expressions) `x` that is not a [package name](https://go.dev/ref/spec#Package_clause), the *selector expression*

x.f

denotes the field or method `f` of the value `x` (or sometimes `*x`; see below). The identifier `f` is called the (field or method) *selector*; it must not be the [blank identifier](https://go.dev/ref/spec#Blank_identifier). The type of the selector expression is the type of `f`. If `x` is a package name, see the section on [qualified identifiers](https://go.dev/ref/spec#Qualified_identifiers).

A selector `f` may denote a field or method `f` of a type `T`, or it may refer to a field or method `f` of a nested [embedded field](https://go.dev/ref/spec#Struct_types) of `T`. The number of embedded fields traversed to reach `f` is called its *depth* in `T`. The depth of a field or method `f` declared in `T` is zero. The depth of a field or method `f` declared in an embedded field `A` in `T` is the depth of `f` in `A` plus one.

The following rules apply to selectors:

1. For a value `x` of type `T` or `*T` where `T` is not a pointer or interface type, `x.f` denotes the field or method at the shallowest depth in `T` where there is such an `f`. If there is not exactly [one `f`](https://go.dev/ref/spec#Uniqueness_of_identifiers) with shallowest depth, the selector expression is illegal.
2. For a value `x` of type `I` where `I` is an interface type, `x.f` denotes the actual method with name `f` of the dynamic value of `x`. If there is no method with name `f` in the [method set](https://go.dev/ref/spec#Method_sets) of `I`, the selector expression is illegal.  
   对于类型 `I` 的值 `x` （其中 `I` 是接口类型）， `x.f` 表示名称为 `f` 的方法集中不存在名称为 `f` 的方法，则选择器表达式非法。
3. As an exception, if the type of `x` is a [defined](https://go.dev/ref/spec#Type_definitions) pointer type and `(*x).f` is a valid selector expression denoting a field (but not a method), `x.f` is shorthand for `(*x).f`.  
   作为例外，如果 `x` 的类型是定义的指针类型，并且 `(*x).f` 是表示字段（但不是方法）的有效选择器表达式，则 `x.f` 是 `(*x).f` 的简写。
4. In all other cases, `x.f` is illegal.  
   在所有其他情况下， `x.f` 都是非法的。
5. If `x` is of pointer type and has the value `nil` and `x.f` denotes a struct field, assigning to or evaluating `x.f` causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).  
   如果 `x` 是指针类型并且值为 `nil` 并且 `x.f` 表示结构体字段，则分配或计算 `x.f` 会导致运行-时间恐慌。
6. If `x` is of interface type and has the value `nil`, [calling](https://go.dev/ref/spec#Calls) or [evaluating](https://go.dev/ref/spec#Method_values) the method `x.f` causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).  
   如果 `x` 是接口类型并且具有值 `nil` ，则调用或评估方法 `x.f` 会导致运行时恐慌。

For example, given the declarations:

type T0 struct {
    x int
}
func (*T0) M0()
type T1 struct {
    y int
}
func (T1) M1()
type T2 struct {
    z int
    T1
    *T0
}
func (*T2) M2()
type Q *T2
var t T2     // with t.T0 != nil
var p *T2    // with p != nil and (*p).T0 != nil
var q Q = p

one may write:

t.z          // t.z
t.y          // t.T1.y
t.x          // (*t.T0).x
p.z          // (*p).z
p.y          // (*p).T1.y
p.x          // (*(*p).T0).x
q.x          // (*(*q).T0).x        (*q).x is a valid field selector
p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
p.M2()       // p.M2()              M2 expects *T2 receiver
t.M2()       // (&t).M2()           M2 expects *T2 receiver, see section on Calls

but the following is invalid:

q.M0()       // (*q).M0 is valid but not a field selector

### Method expressions

If `M` is in the [method set](https://go.dev/ref/spec#Method_sets) of type `T`, `T.M` is a function that is callable as a regular function with the same arguments as `M` prefixed by an additional argument that is the receiver of the method.

MethodExpr    = [ReceiverType](https://go.dev/ref/spec#ReceiverType) "." [MethodName](https://go.dev/ref/spec#MethodName) .
ReceiverType  = [Type](https://go.dev/ref/spec#Type) .

Consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.  
考虑一个具有两个方法的结构体类型 `T` ， `Mv` ，其接收者的类型为 `T` ，以及 `Mp` ，其接收者的类型 `*T` 。

type T struct {
    a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver
var t T

The expression  表达方式

T.Mv

yields a function equivalent to `Mv` but with an explicit receiver as its first argument; it has signature  
产生一个与 `Mv` 等价的函数，但具有显式接收者作为其第一个参数；它有签名

func(tv T, a int) int

That function may be called normally with an explicit receiver, so these five invocations are equivalent:  
该函数可以通过显式接收器正常调用，因此这五个调用是等效的：

t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)

Similarly, the expression  
类似地，表达式

(*T).Mp

yields a function value representing `Mp` with signature  
产生一个代表带有签名的 `Mp` 的函数值

func(tp *T, f float32) float32

For a method with a value receiver, one can derive a function with an explicit pointer receiver, so  
对于具有值接收器的方法，可以派生出具有显式指针接收器的函数，因此

(*T).Mv

yields a function value representing `Mv` with signature  
产生一个代表带有签名的 `Mv` 的函数值

func(tv *T, a int) int

Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.  
这样的函数间接通过接收者创建一个值作为接收者传递给底层方法；该方法不会覆盖在函数调用中传递的地址的值。

The final case, a value-receiver function for a pointer-receiver method, is illegal because pointer-receiver methods are not in the method set of the value type.  
最后一种情况，即指针接收器方法的值接收器函数是非法的，因为指针接收器方法不在值类型的方法集中。

Function values derived from methods are called with function call syntax; the receiver is provided as the first argument to the call. That is, given `f := T.Mv`, `f` is invoked as `f(t, 7)` not `t.f(7)`. To construct a function that binds the receiver, use a [function literal](https://go.dev/ref/spec#Function_literals) or [method value](https://go.dev/ref/spec#Method_values).

It is legal to derive a function value from a method of an interface type. The resulting function takes an explicit receiver of that interface type.

### Method values

If the expression `x` has static type `T` and `M` is in the [method set](https://go.dev/ref/spec#Method_sets) of type `T`, `x.M` is called a *method value*. The method value `x.M` is a function value that is callable with the same arguments as a method call of `x.M`. The expression `x` is evaluated and saved during the evaluation of the method value; the saved copy is then used as the receiver in any calls, which may be executed later.

type S struct { *T }
type T int
func (t T) M() { print(t) }
t := new(T)
s := S{T: t}
f := t.M                    // receiver *t is evaluated and stored in f
g := s.M                    // receiver *(s.T) is evaluated and stored in g
*t = 42                     // does not affect stored receivers in f and g

The type `T` may be an interface or non-interface type.

As in the discussion of [method expressions](https://go.dev/ref/spec#Method_expressions) above, consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

type T struct {
    a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver
var t T
var pt *T
func makeT() T

The expression

t.Mv

yields a function value of type

func(int) int

These two invocations are equivalent:

t.Mv(7)
f := t.Mv; f(7)

Similarly, the expression

pt.Mp

yields a function value of type

func(float32) float32

As with [selectors](https://go.dev/ref/spec#Selectors), a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`.

As with [method calls](https://go.dev/ref/spec#Calls), a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

f := t.Mv; f(7)   // like t.Mv(7)
f := pt.Mp; f(7)  // like pt.Mp(7)
f := pt.Mv; f(7)  // like (*pt).Mv(7)
f := t.Mp; f(7)   // like (&t).Mp(7)
f := makeT().Mp   // invalid: result of makeT() is not addressable

Although the examples above use non-interface types, it is also legal to create a method value from a value of interface type.

var i interface { M(int) } = myVal
f := i.M; f(7)  // like i.M(7)

### Index expressions

A primary expression of the form

a[x]

denotes the element of the array, pointer to array, slice, string or map `a` indexed by `x`. The value `x` is called the *index* or *map key*, respectively. The following rules apply:

If `a` is neither a map nor a type parameter:  
如果 `a` 既不是映射也不是类型参数：

- the index `x` must be an untyped constant or its [core type](https://go.dev/ref/spec#Core_types) must be an [integer](https://go.dev/ref/spec#Numeric_types)  
  索引 `x` 必须是无类型常量或其核心类型必须是整数
- a constant index must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`  
  常量索引必须是非负的并且可以由 `int` 类型的值表示
- a constant index that is untyped is given type `int`  
  无类型常量索引的类型为 `int`
- the index `x` is *in range* if `0 <= x < len(a)`, otherwise it is *out of range*  
  如果 `0 <= x < len(a)` 则索引 `x` 在范围内，否则超出范围

For `a` of [array type](https://go.dev/ref/spec#Array_types) `A`:  
对于数组类型 `A` 的 `a` ：

- a [constant](https://go.dev/ref/spec#Constants) index must be in range  
  常量索引必须在范围内
- if `x` is out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs  
  如果 `x` 在运行时超出范围，则会发生运行时恐慌
- `a[x]` is the array element at index `x` and the type of `a[x]` is the element type of `A`  
  `a[x]` 是索引 `x` 处的数组元素， `a[x]` 的类型是 `A` 的元素类型

For `a` of [pointer](https://go.dev/ref/spec#Pointer_types) to array type:  
对于数组类型指针的 `a` ：

- `a[x]` is shorthand for `(*a)[x]`  
  `a[x]` 是 `(*a)[x]` 的简写

For `a` of [slice type](https://go.dev/ref/spec#Slice_types) `S`:  
对于切片类型 `S` 的 `a` ：

- if `x` is out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs  
  如果 `x` 在运行时超出范围，则会发生运行时恐慌
- `a[x]` is the slice element at index `x` and the type of `a[x]` is the element type of `S`  
  `a[x]` 是索引 `x` 处的切片元素， `a[x]` 的类型是 `S` 的元素类型

For `a` of [string type](https://go.dev/ref/spec#String_types):  
对于字符串类型的 `a` ：

- a [constant](https://go.dev/ref/spec#Constants) index must be in range if the string `a` is also constant  
  如果字符串 `a` 也是常量，则常量索引必须在范围内
- if `x` is out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs  
  如果 `x` 在运行时超出范围，则会发生运行时恐慌
- `a[x]` is the non-constant byte value at index `x` and the type of `a[x]` is `byte`  
  `a[x]` 是索引 `x` 处的非常量字节值， `a[x]` 的类型为 `byte`
- `a[x]` may not be assigned to  
  `a[x]` 不能分配给

For `a` of [map type](https://go.dev/ref/spec#Map_types) `M`:  
对于地图类型 `M` 的 `a` ：

- `x`'s type must be [assignable](https://go.dev/ref/spec#Assignability) to the key type of `M`  
  `x` 的类型必须可分配给 `M` 的键类型
- if the map contains an entry with key `x`, `a[x]` is the map element with key `x` and the type of `a[x]` is the element type of `M`  
  如果地图包含带有键 `x` 的条目，则 `a[x]` 是带有键 `x` 的地图元素，而 `a[x]` 的类型是该元素 `M` 的类型
- if the map is `nil` or does not contain such an entry, `a[x]` is the [zero value](https://go.dev/ref/spec#The_zero_value) for the element type of `M`  
  如果地图是 `nil` 或不包含此类条目，则 `a[x]` 是 `M` 元素类型的零值

For `a` of [type parameter type](https://go.dev/ref/spec#Type_parameter_declarations) `P`:  
对于类型参数类型 `P` 的 `a` ：

- The index expression `a[x]` must be valid for values of all types in `P`'s type set.  
  索引表达式 `a[x]` 必须对 `P` 类型集中所有类型的值都有效。
- The element types of all types in `P`'s type set must be identical. In this context, the element type of a string type is `byte`.  
  `P` 类型集中所有类型的元素类型必须相同。在这种情况下，字符串类型的元素类型是 `byte` 。
- If there is a map type in the type set of `P`, all types in that type set must be map types, and the respective key types must be all identical.  
  如果 `P` 类型集中存在映射类型，则该类型集中的所有类型都必须是映射类型，并且各自的键类型必须全部相同。
- `a[x]` is the array, slice, or string element at index `x`, or the map element with key `x` of the type argument that `P` is instantiated with, and the type of `a[x]` is the type of the (identical) element types.  
  `a[x]` 是索引 `x` 处的数组、切片或字符串元素，或者是 `P` 的映射元素> 被实例化， `a[x]` 的类型是（相同）元素类型的类型。
- `a[x]` may not be assigned to if `P`'s type set includes string types.  
  如果 `P` 的类型集包含字符串类型，则不能分配 `a[x]` 。

Otherwise `a[x]` is illegal.  
否则 `a[x]` 是非法的。

An index expression on a map `a` of type `map[K]V` used in an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or initialization of the special form  
`map[K]V` 类型的映射 `a` 上的索引表达式，用于赋值语句或特殊形式的初始化

v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]

yields an additional untyped boolean value. The value of `ok` is `true` if the key `x` is present in the map, and `false` otherwise.

Assigning to an element of a `nil` map causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).

### Slice expressions

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice. There are two variants: a simple form that specifies a low and high bound, and a full form that also specifies a bound on the capacity.

#### Simple slice expressions

The primary expression

a[low : high]

constructs a substring or slice. The [core type](https://go.dev/ref/spec#Core_types) of `a` must be a string, array, pointer to array, slice, or a [`bytestring`](https://go.dev/ref/spec#Core_types). The *indices* `low` and `high` select which elements of operand `a` appear in the result. The result has indices starting at 0 and length equal to `high` - `low`. After slicing the array `a`

a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]

the slice `s` has type `[]int`, length 3, capacity 4, and elements

s[0] == 2
s[1] == 3
s[2] == 4

For convenience, any of the indices may be omitted. A missing `low` index defaults to zero; a missing `high` index defaults to the length of the sliced operand:

a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]

If `a` is a pointer to an array, `a[low : high]` is shorthand for `(*a)[low : high]`.

For arrays or strings, the indices are *in range* if `0` <= `low` <= `high` <= `len(a)`, otherwise they are *out of range*. For slices, the upper index bound is the slice capacity `cap(a)` rather than the length. A [constant](https://go.dev/ref/spec#Constants) index must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; for arrays or constant strings, constant indices must also be in range. If both indices are constant, they must satisfy `low <= high`. If the indices are out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.

Except for [untyped strings](https://go.dev/ref/spec#Constants), if the sliced operand is a string or slice, the result of the slice operation is a non-constant value of the same type as the operand. For untyped string operands the result is a non-constant value of type `string`. If the sliced operand is an array, it must be [addressable](https://go.dev/ref/spec#Address_operators) and the result of the slice operation is a slice with the same element type as the array.

If the sliced operand of a valid slice expression is a `nil` slice, the result is a `nil` slice. Otherwise, if the result is a slice, it shares its underlying array with the operand.

var a [10]int
s1 := a[3:7]   // underlying array of s1 is array a; &s1[2] == &a[5]
s2 := s1[1:4]  // underlying array of s2 is underlying array of s1 which is array a; &s2[1] == &a[5]
s2[1] = 42     // s2[1] == s1[2] == a[5] == 42; they all refer to the same underlying array element
var s []int
s3 := s[:0]    // s3 == nil

#### Full slice expressions

The primary expression

a[low : high : max]

constructs a slice of the same type, and with the same length and elements as the simple slice expression `a[low : high]`. Additionally, it controls the resulting slice's capacity by setting it to `max - low`. Only the first index may be omitted; it defaults to 0. The [core type](https://go.dev/ref/spec#Core_types) of `a` must be an array, pointer to array, or slice (but not a string). After slicing the array `a`

a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]

the slice `t` has type `[]int`, length 2, capacity 4, and elements

t[0] == 2
t[1] == 3

As for simple slice expressions, if `a` is a pointer to an array, `a[low : high : max]` is shorthand for `(*a)[low : high : max]`. If the sliced operand is an array, it must be [addressable](https://go.dev/ref/spec#Address_operators).

The indices are *in range* if `0 <= low <= high <= max <= cap(a)`, otherwise they are *out of range*. A [constant](https://go.dev/ref/spec#Constants) index must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; for arrays, constant indices must also be in range. If multiple indices are constant, the constants that are present must be in range relative to each other. If the indices are out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.

### Type assertions

For an expression `x` of [interface type](https://go.dev/ref/spec#Interface_types), but not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), and a type `T`, the primary expression

x.(T)

asserts that `x` is not `nil` and that the value stored in `x` is of type `T`. The notation `x.(T)` is called a *type assertion*.

More precisely, if `T` is not an interface type, `x.(T)` asserts that the dynamic type of `x` is [identical](https://go.dev/ref/spec#Type_identity) to the type `T`. In this case, `T` must [implement](https://go.dev/ref/spec#Method_sets) the (interface) type of `x`; otherwise the type assertion is invalid since it is not possible for `x` to store a value of type `T`. If `T` is an interface type, `x.(T)` asserts that the dynamic type of `x` [implements](https://go.dev/ref/spec#Implementing_an_interface) the interface `T`.

If the type assertion holds, the value of the expression is the value stored in `x` and its type is `T`. If the type assertion is false, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. In other words, even though the dynamic type of `x` is known only at run time, the type of `x.(T)` is known to be `T` in a correct program.

var x interface{} = 7          // x has dynamic type int and value 7
i := x.(int)                   // i has type int and value 7
type I interface { m() }
func f(y I) {
    s := y.(string)        // illegal: string does not implement I (missing method m)
    r := y.(io.Reader)     // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
    …
}

A type assertion used in an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or initialization of the special form  
在赋值语句或特殊形式的初始化中使用的类型断言

v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok interface{} = x.(T) // dynamic types of v and ok are T and bool

yields an additional untyped boolean value. The value of `ok` is `true` if the assertion holds. Otherwise it is `false` and the value of `v` is the [zero value](https://go.dev/ref/spec#The_zero_value) for type `T`. No [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs in this case.  
产生一个额外的无类型布尔值。如果断言成立，则 `ok` 的值为 `true` 。否则它是 `false` 并且 `v` 的值是类型 `T` 的零值。在这种情况下不会发生运行时恐慌。

### Calls 通话

Given an expression `f` with a [core type](https://go.dev/ref/spec#Core_types) `F` of [function type](https://go.dev/ref/spec#Function_types),  
给定一个具有函数类型核心类型 `F` 的表达式 `f` ，

f(a1, a2, … an)

calls `f` with arguments `a1, a2, … an`. Except for one special case, arguments must be single-valued expressions [assignable](https://go.dev/ref/spec#Assignability) to the parameter types of `F` and are evaluated before the function is called. The type of the expression is the result type of `F`. A method invocation is similar but the method itself is specified as a selector upon a value of the receiver type for the method.  
使用参数 `a1, a2, … an` 调用 `f` 。除一种特殊情况外，参数必须是可分配给 `F` 参数类型的单值表达式，并在调用函数之前进行计算。表达式的类型是 `F` 的结果类型。方法调用类似，但方法本身被指定为该方法的接收者类型值的选择器。

math.Atan2(x, y)  // function call
var pt *Point
pt.Scale(3.5)     // method call with receiver pt

If `f` denotes a generic function, it must be [instantiated](https://go.dev/ref/spec#Instantiations) before it can be called or used as a function value.  
如果 `f` 表示泛型函数，则必须先实例化它，然后才能调用它或将其用作函数值。

In a function call, the function value and arguments are evaluated in [the usual order](https://go.dev/ref/spec#Order_of_evaluation). After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution. The return parameters of the function are passed by value back to the caller when the function returns.  
在函数调用中，函数值和参数按通常的顺序求值。在评估它们之后，调用的参数将按值传递给函数，并且被调用的函数开始执行。当函数返回时，函数的返回参数按值传递回调用者。

Calling a `nil` function value causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).  
调用 `nil` 函数值会导致运行时恐慌。

As a special case, if the return values of a function or method `g` are equal in number and individually assignable to the parameters of another function or method `f`, then the call `f(g(*parameters_of_g*))` will invoke `f` after binding the return values of `g` to the parameters of `f` in order. The call of `f` must contain no parameters other than the call of `g`, and `g` must have at least one return value. If `f` has a final `...` parameter, it is assigned the return values of `g` that remain after assignment of regular parameters.  
作为一种特殊情况，如果函数或方法 `g` 的返回值数量相等并且可单独分配给另一个函数或方法 `f` 的参数，则调用 `f(g(*parameters_of_g*))` 将 `g` 的返回值按顺序绑定到 `f` 的参数后调用 `f` 。 `f` 的调用除了 `g` 的调用之外不得包含任何参数，并且 `g` 必须至少有一个返回值。如果 `f` 有一个最终的 `...` 参数，则会为其分配常规参数分配后保留的 `g` 返回值。

func Split(s string, pos int) (string, string) {
    return s[0:pos], s[pos:]
}
func Join(s, t string) string {
    return s + t
}
if Join(Split(value, len(value)/2)) != value {
    log.Panic("test fails")
}

A method call `x.m()` is valid if the [method set](https://go.dev/ref/spec#Method_sets) of (the type of) `x` contains `m` and the argument list can be assigned to the parameter list of `m`. If `x` is [addressable](https://go.dev/ref/spec#Address_operators) and `&x`'s method set contains `m`, `x.m()` is shorthand for `(&x).m()`:

var p Point
p.Scale(3.5)

There is no distinct method type and there are no method literals.

### Passing arguments to `...` parameters

If `f` is [variadic](https://go.dev/ref/spec#Function_types) with a final parameter `p` of type `...T`, then within `f` the type of `p` is equivalent to type `[]T`. If `f` is invoked with no actual arguments for `p`, the value passed to `p` is `nil`. Otherwise, the value passed is a new slice of type `[]T` with a new underlying array whose successive elements are the actual arguments, which all must be [assignable](https://go.dev/ref/spec#Assignability) to `T`. The length and capacity of the slice is therefore the number of arguments bound to `p` and may differ for each call site.

Given the function and calls

func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")

within `Greeting`, `who` will have the value `nil` in the first call, and `[]string{"Joe", "Anna", "Eileen"}` in the second.

If the final argument is assignable to a slice type `[]T` and is followed by `...`, it is passed unchanged as the value for a `...T` parameter. In this case no new slice is created.

Given the slice `s` and call

s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)

within `Greeting`, `who` will have the same value as `s` with the same underlying array.

### Instantiations

A generic function or type is *instantiated* by substituting *type arguments* for the type parameters [[Go 1.18](https://go.dev/ref/spec#Go_1.18)]. Instantiation proceeds in two steps:

1. Each type argument is substituted for its corresponding type parameter in the generic declaration. This substitution happens across the entire function or type declaration, including the type parameter list itself and any types in that list.
2. After substitution, each type argument must [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) the [constraint](https://go.dev/ref/spec#Type_parameter_declarations) (instantiated, if necessary) of the corresponding type parameter. Otherwise instantiation fails.

Instantiating a type results in a new non-generic [named type](https://go.dev/ref/spec#Types); instantiating a function produces a new non-generic function.

type parameter list    type arguments    after substitution
[P any]                int               int satisfies any
[S ~[]E, E any]        []int, int        []int satisfies ~[]int, int satisfies any
[P io.Writer]          string            illegal: string doesn't satisfy io.Writer
[P comparable]         any               any satisfies (but does not implement) comparable

When using a generic function, type arguments may be provided explicitly, or they may be partially or completely [inferred](https://go.dev/ref/spec#Type_inference) from the context in which the function is used. Provided that they can be inferred, type argument lists may be omitted entirely if the function is:

- [called](https://go.dev/ref/spec#Calls) with ordinary arguments,
- [assigned](https://go.dev/ref/spec#Assignment_statements) to a variable with a known type
- [passed as an argument](https://go.dev/ref/spec#Calls) to another function, or
- [returned as a result](https://go.dev/ref/spec#Return_statements).

In all other cases, a (possibly partial) type argument list must be present. If a type argument list is absent or partial, all missing type arguments must be inferrable from the context in which the function is used.

// sum returns the sum (concatenation, for strings) of its arguments.
func sum[T ~int | ~float64 | ~string](x... T) T { … }
x := sum                       // illegal: the type of x is unknown
intSum := sum[int]             // intSum has type func(x... int) int
a := intSum(2, 3)              // a has value 5 of type int
b := sum[float64](2.0, 3)      // b has value 5.0 of type float64
c := sum(b, -1)                // c has value 4.0 of type float64
type sumFunc func(x... string) string
var f sumFunc = sum            // same as var f sumFunc = sum[string]
f = sum                        // same as f = sum[string]

A partial type argument list cannot be empty; at least the first argument must be present. The list is a prefix of the full list of type arguments, leaving the remaining arguments to be inferred. Loosely speaking, type arguments may be omitted from "right to left".  
部分类型参数列表不能为空；至少必须存在第一个参数。该列表是类型参数完整列表的前缀，剩下的参数将被推断。宽松地说，类型参数可以从“从右到左”省略。

func apply[S ~[]E, E any](s S, f func(E) E) S { … }
f0 := apply[]                  // illegal: type argument list cannot be empty
f1 := apply[[]int]             // type argument for S explicitly provided, type argument for E inferred
f2 := apply[[]string, string]  // both type arguments explicitly provided
var bytes []byte
r := apply(bytes, func(byte) byte { … })  // both type arguments inferred from the function arguments

For a generic type, all type arguments must always be provided explicitly.  
对于泛型类型，必须始终显式提供所有类型参数。

### Type inference 类型推断

A use of a generic function may omit some or all type arguments if they can be *inferred* from the context within which the function is used, including the constraints of the function's type parameters. Type inference succeeds if it can infer the missing type arguments and [instantiation](https://go.dev/ref/spec#Instantiations) succeeds with the inferred type arguments. Otherwise, type inference fails and the program is invalid.  
如果可以从使用函数的上下文（包括函数类型参数的约束）推断出某些或全部类型参数，则使用泛型函数可以省略这些类型参数。如果类型推断可以推断出缺失的类型参数，并且推断出的类型参数的实例化成功，则类型推断成功。否则，类型推断失败，程序无效。

Type inference uses the type relationships between pairs of types for inference: For instance, a function argument must be [assignable](https://go.dev/ref/spec#Assignability) to its respective function parameter; this establishes a relationship between the type of the argument and the type of the parameter. If either of these two types contains type parameters, type inference looks for the type arguments to substitute the type parameters with such that the assignability relationship is satisfied. Similarly, type inference uses the fact that a type argument must [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) the constraint of its respective type parameter.  
类型推断使用类型对之间的类型关系进行推断：例如，函数参数必须可分配给其各自的函数参数；这在参数类型和形参类型之间建立了关系。如果这两种类型中的任何一个包含类型参数，则类型推断会查找类型参数来替换类型参数，以满足可赋值关系。类似地，类型推断使用类型参数必须满足其各自类型参数的约束这一事实。

Each such pair of matched types corresponds to a *type equation* containing one or multiple type parameters, from one or possibly multiple generic functions. Inferring the missing type arguments means solving the resulting set of type equations for the respective type parameters.  
每对这样的匹配类型对应于包含来自一个或可能多个泛型函数的一个或多个类型参数的类型方程。推断缺失的类型参数意味着求解相应类型参数的类型方程结果集。

For example, given  例如，给定

// dedup returns a copy of the argument slice with any duplicate entries removed.
func dedup[S ~[]E, E comparable](S) S { … }
type Slice []int
var s Slice
s = dedup(s)   // same as s = dedup[Slice, int](s)

the variable `s` of type `Slice` must be assignable to the function parameter type `S` for the program to be valid. To reduce complexity, type inference ignores the directionality of assignments, so the type relationship between `Slice` and `S` can be expressed via the (symmetric) type equation `Slice ≡A S` (or `S ≡A Slice` for that matter), where the `A` in `≡A` indicates that the LHS and RHS types must match per assignability rules (see the section on [type unification](https://go.dev/ref/spec#Type_unification) for details). Similarly, the type parameter `S` must satisfy its constraint `~[]E`. This can be expressed as `S ≡C ~[]E` where `X ≡C Y` stands for "`X` satisfies constraint `Y`". These observations lead to a set of two equations  
`Slice` 类型的变量 `s` 必须可分配给函数参数类型 `S` ，程序才有效。为了降低复杂性，类型推断忽略了赋值的方向性，因此 `Slice` 和 `S` 之间的类型关系可以通过（对称）类型方程 `Slice ≡A S` 来表达（或 `S ≡A Slice` ），其中 `≡A` 中的 `A` 指示 LHS 和 RHS 类型必须根据可分配性规则匹配（请参阅有关类型统一的部分细节）。同样，类型参数 `S` 必须满足其约束 `~[]E` 。这可以表示为 `S ≡C ~[]E` ，其中 `X ≡C Y` 代表“ `X` 满足约束 `Y` ”。这些观察结果得出一组两个方程

    Slice ≡A S      (1)
    S     ≡C ~[]E   (2)

which now can be solved for the type parameters `S` and `E`. From (1) a compiler can infer that the type argument for `S` is `Slice`. Similarly, because the underlying type of `Slice` is `[]int` and `[]int` must match `[]E` of the constraint, a compiler can infer that `E` must be `int`. Thus, for these two equations, type inference infers  
现在可以解决类型参数 `S` 和 `E` 的问题。从 (1) 编译器可以推断 `S` 的类型参数是 `Slice` 。同样，由于 `Slice` 的基础类型是 `[]int` 并且 `[]int` 必须与约束的 `[]E` 匹配，因此编译器可以推断出 `E` 必须是 `int` 。因此，对于这两个方程，类型推断推断出

    S ➞ Slice
    E ➞ int

Given a set of type equations, the type parameters to solve for are the type parameters of the functions that need to be instantiated and for which no explicit type arguments is provided. These type parameters are called *bound* type parameters. For instance, in the `dedup` example above, the type parameters `S` and `E` are bound to `dedup`. An argument to a generic function call may be a generic function itself. The type parameters of that function are included in the set of bound type parameters. The types of function arguments may contain type parameters from other functions (such as a generic function enclosing a function call). Those type parameters may also appear in type equations but they are not bound in that context. Type equations are always solved for the bound type parameters only.

Type inference supports calls of generic functions and assignments of generic functions to (explicitly function-typed) variables. This includes passing generic functions as arguments to other (possibly also generic) functions, and returning generic functions as results. Type inference operates on a set of equations specific to each of these cases. The equations are as follows (type argument lists are omitted for clarity):

- For a function call `f(a0, a1, …)` where `f` or a function argument `ai` is a generic function:  
  Each pair `(ai, pi)` of corresponding function arguments and parameters where `ai` is not an [untyped constant](https://go.dev/ref/spec#Constants) yields an equation `typeof(pi) ≡A typeof(ai)`.  
  If `ai` is an untyped constant `cj`, and `typeof(pi)` is a bound type parameter `Pk`, the pair `(cj, Pk)` is collected separately from the type equations.

- For an assignment `v = f` of a generic function `f` to a (non-generic) variable `v` of function type:  
  `typeof(v) ≡A typeof(f)`.

- For a return statement `return …, f, …` where `f` is a generic function returned as a result to a (non-generic) result variable `r` of function type:  
  `typeof(r) ≡A typeof(f)`.

Additionally, each type parameter `Pk` and corresponding type constraint `Ck` yields the type equation `Pk ≡C Ck`.

Type inference gives precedence to type information obtained from typed operands before considering untyped constants. Therefore, inference proceeds in two phases:

1. The type equations are solved for the bound type parameters using [type unification](https://go.dev/ref/spec#Type_unification). If unification fails, type inference fails.

2. For each bound type parameter `Pk` for which no type argument has been inferred yet and for which one or more pairs `(cj, Pk)` with that same type parameter were collected, determine the [constant kind](https://go.dev/ref/spec#Constant_expressions) of the constants `cj` in all those pairs the same way as for [constant expressions](https://go.dev/ref/spec#Constant_expressions). The type argument for `Pk` is the [default type](https://go.dev/ref/spec#Constants) for the determined constant kind. If a constant kind cannot be determined due to conflicting constant kinds, type inference fails.

If not all type arguments have been found after these two phases, type inference fails.

If the two phases are successful, type inference determined a type argument for each bound type parameter:

    Pk ➞ Ak

A type argument `Ak` may be a composite type, containing other bound type parameters `Pk` as element types (or even be just another bound type parameter). In a process of repeated simplification, the bound type parameters in each type argument are substituted with the respective type arguments for those type parameters until each type argument is free of bound type parameters.

If type arguments contain cyclic references to themselves through bound type parameters, simplification and thus type inference fails. Otherwise, type inference succeeds.

#### Type unification

Type inference solves type equations through *type unification*. Type unification recursively compares the LHS and RHS types of an equation, where either or both types may be or contain bound type parameters, and looks for type arguments for those type parameters such that the LHS and RHS match (become identical or assignment-compatible, depending on context). To that effect, type inference maintains a map of bound type parameters to inferred type arguments; this map is consulted and updated during type unification. Initially, the bound type parameters are known but the map is empty. During type unification, if a new type argument `A` is inferred, the respective mapping `P ➞ A` from type parameter to argument is added to the map. Conversely, when comparing types, a known type argument (a type argument for which a map entry already exists) takes the place of its corresponding type parameter. As type inference progresses, the map is populated more and more until all equations have been considered, or until unification fails. Type inference succeeds if no unification step fails and the map has an entry for each type parameter.

For example, given the type equation with the bound type parameter `P`

    [10]struct{ elem P, list []P } ≡A [10]struct{ elem string; list []string }

type inference starts with an empty map. Unification first compares the top-level structure of the LHS and RHS types. Both are arrays of the same length; they unify if the element types unify. Both element types are structs; they unify if they have the same number of fields with the same names and if the field types unify. The type argument for `P` is not known yet (there is no map entry), so unifying `P` with `string` adds the mapping `P ➞ string` to the map. Unifying the types of the `list` field requires unifying `[]P` and `[]string` and thus `P` and `string`. Since the type argument for `P` is known at this point (there is a map entry for `P`), its type argument `string` takes the place of `P`. And since `string` is identical to `string`, this unification step succeeds as well. Unification of the LHS and RHS of the equation is now finished. Type inference succeeds because there is only one type equation, no unification step failed, and the map is fully populated.  
类型推断从空映射开始。统一首先比较LHS和RHS类型的顶层结构。两者都是相同长度的数组；如果元素类型统一，它们就会统一。两种元素类型都是结构体；如果它们具有相同数量、相同名称的字段并且字段类型统一，则它们是统一的。 `P` 的类型参数尚不清楚（没有映射条目），因此将 `P` 与 `string` 统一会添加映射 `P ➞ string` 到地图。统一 `list` 字段的类型需要统一 `[]P` 和 `[]string` ，从而统一 `P` 和 `string` 。由于此时 `P` 的类型参数已知（有 `P` 的映射条目），因此其类型参数 `string` 取代 `P` 。由于 `string` 与 `string` 相同，因此此统一步骤也成功。方程的 LHS 和 RHS 的统一现已完成。类型推断成功，因为只有一个类型方程，没有统一步骤失败，并且映射已完全填充。

Unification uses a combination of *exact* and *loose* unification depending on whether two types have to be [identical](https://go.dev/ref/spec#Type_identity), [assignment-compatible](https://go.dev/ref/spec#Assignability), or only structurally equal. The respective [type unification rules](https://go.dev/ref/spec#Type_unification_rules) are spelled out in detail in the [Appendix](https://go.dev/ref/spec#Appendix).  
统一使用精确统一和松散统一的组合，具体取决于两种类型是否必须相同、分配兼容或仅在结构上相等。各个类型统一规则在附录中详细说明。

For an equation of the form `X ≡A Y`, where `X` and `Y` are types involved in an assignment (including parameter passing and return statements), the top-level type structures may unify loosely but element types must unify exactly, matching the rules for assignments.  
对于 `X ≡A Y` 形式的方程，其中 `X` 和 `Y` 是赋值所涉及的类型（包括参数传递和返回语句），顶级类型结构可以松散地统一，但元素类型必须精确统一，符合赋值规则。

For an equation of the form `P ≡C C`, where `P` is a type parameter and `C` its corresponding constraint, the unification rules are bit more complicated:  
对于 `P ≡C C` 形式的方程，其中 `P` 是类型参数， `C` 是其相应的约束，统一规则有点复杂：

- If `C` has a [core type](https://go.dev/ref/spec#Core_types) `core(C)` and `P` has a known type argument `A`, `core(C)` and `A` must unify loosely. If `P` does not have a known type argument and `C` contains exactly one type term `T` that is not an underlying (tilde) type, unification adds the mapping `P ➞ T` to the map.  
  如果 `C` 具有核心类型 `core(C)` 且 `P` 具有已知类型参数 `A` 、 `core(C)` 和 `A` 必须松散地统一。如果 `P` 没有已知的类型参数，并且 `C` 恰好包含一个不是基础（波形符）类型的类型术语 `T` ，则统一会添加映射 < b9> 到地图。
- If `C` does not have a core type and `P` has a known type argument `A`, `A` must have all methods of `C`, if any, and corresponding method types must unify exactly.  
  如果 `C` 没有核心类型并且 `P` 具有已知类型参数 `A` ，则 `A` 必须具有 `C` ，如果有的话，以及相应的方法类型必须完全统一。

When solving type equations from type constraints, solving one equation may infer additional type arguments, which in turn may enable solving other equations that depend on those type arguments. Type inference repeats type unification as long as new type arguments are inferred.  
当根据类型约束求解类型方程时，求解一个方程可能会推断出其他类型参数，这反过来又可以求解依赖于这些类型参数的其他方程。只要推断出新类型参数，类型推断就会重复类型统一。

### Operators 运营商

Operators combine operands into expressions.  
运算符将操作数组合成表达式。

Expression = [UnaryExpr](https://go.dev/ref/spec#UnaryExpr) | [Expression](https://go.dev/ref/spec#Expression) [binary_op](https://go.dev/ref/spec#binary_op) [Expression](https://go.dev/ref/spec#Expression) .
UnaryExpr  = [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) | [unary_op](https://go.dev/ref/spec#unary_op) [UnaryExpr](https://go.dev/ref/spec#UnaryExpr) .

binary_op  = "||" | "&&" | [rel_op](https://go.dev/ref/spec#rel_op) | [add_op](https://go.dev/ref/spec#add_op) | [mul_op](https://go.dev/ref/spec#mul_op) .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .

Comparisons are discussed [elsewhere](https://go.dev/ref/spec#Comparison_operators). For other binary operators, the operand types must be [identical](https://go.dev/ref/spec#Type_identity) unless the operation involves shifts or untyped [constants](https://go.dev/ref/spec#Constants). For operations involving constants only, see the section on [constant expressions](https://go.dev/ref/spec#Constant_expressions).  
比较将在别处讨论。对于其他二元运算符，操作数类型必须相同，除非运算涉及移位或无类型常量。对于仅涉及常量的运算，请参阅有关常量表达式的部分。

Except for shift operations, if one operand is an untyped [constant](https://go.dev/ref/spec#Constants) and the other operand is not, the constant is implicitly [converted](https://go.dev/ref/spec#Conversions) to the type of the other operand.  
除移位操作外，如果一个操作数是无类型常量而另一个操作数不是，则该常量将隐式转换为另一个操作数的类型。

The right operand in a shift expression must have [integer type](https://go.dev/ref/spec#Numeric_types) [[Go 1.13](https://go.dev/ref/spec#Go_1.13)] or be an untyped constant [representable](https://go.dev/ref/spec#Representability) by a value of type `uint`. If the left operand of a non-constant shift expression is an untyped constant, it is first implicitly converted to the type it would assume if the shift expression were replaced by its left operand alone.  
移位表达式中的右操作数必须具有整数类型 [Go 1.13] 或者是可由 `uint` 类型的值表示的无类型常量。如果非常量移位表达式的左操作数是无类型常量，则它首先会隐式转换为移位表达式仅由其左操作数替换时所假定的类型。

var a [1024]byte
var s uint = 33
// The results of the following examples are given for 64-bit ints.
var i = 1<<s                   // 1 has type int
var j int32 = 1<<s             // 1 has type int32; j == 0
var k = uint64(1<<s)           // 1 has type uint64; k == 1<<33
var m int = 1.0<<s             // 1.0 has type int; m == 1<<33
var n = 1.0<<s == j            // 1.0 has type int32; n == true
var o = 1<<s == 2<<s           // 1 and 2 have type int; o == false
var p = 1<<s == 1<<33          // 1 has type int; p == true
var u = 1.0<<s                 // illegal: 1.0 has type float64, cannot shift
var u1 = 1.0<<s != 0           // illegal: 1.0 has type float64, cannot shift
var u2 = 1<<s != 1.0           // illegal: 1 has type float64, cannot shift
var v1 float32 = 1<<s          // illegal: 1 has type float32, cannot shift
var v2 = string(1<<s)          // illegal: 1 is converted to a string, cannot shift
var w int64 = 1.0<<33          // 1.0<<33 is a constant shift expression; w == 1<<33
var x = a[1.0<<s]              // panics: 1.0 has type int, but 1<<33 overflows array bounds
var b = make([]byte, 1.0<<s)   // 1.0 has type int; len(b) == 1<<33
// The results of the following examples are given for 32-bit ints,
// which means the shifts will overflow.
var mm int = 1.0<<s            // 1.0 has type int; mm == 0
var oo = 1<<s == 2<<s          // 1 and 2 have type int; oo == true
var pp = 1<<s == 1<<33         // illegal: 1 has type int, but 1<<33 overflows int
var xx = a[1.0<<s]             // 1.0 has type int; xx == a[0]
var bb = make([]byte, 1.0<<s)  // 1.0 has type int; len(bb) == 0

#### Operator precedence

Unary operators have the highest precedence. As the `++` and `--` operators form statements, not expressions, they fall outside the operator hierarchy. As a consequence, statement `*p++` is the same as `(*p)++`.

There are five precedence levels for binary operators. Multiplication operators bind strongest, followed by addition operators, comparison operators, `&&` (logical AND), and finally `||` (logical OR):

Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||

Binary operators of the same precedence associate from left to right. For instance, `x / y * z` is the same as `(x / y) * z`.

+x                         // x
42 + a - b                 // (42 + a) - b
23 + 3*x[i]                // 23 + (3 * x[i])
x <= f()                   // x <= f()
^a >> b                    // (^a) >> b
f() || g()                 // f() || g()
x == y+1 && <-chanInt > 0  // (x == (y+1)) && ((<-chanInt) > 0)

### Arithmetic operators

Arithmetic operators apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (`+`, `-`, `*`, `/`) apply to [integer](https://go.dev/ref/spec#Numeric_types), [floating-point](https://go.dev/ref/spec#Numeric_types), and [complex](https://go.dev/ref/spec#Numeric_types) types; `+` also applies to [strings](https://go.dev/ref/spec#String_types). The bitwise logical and shift operators apply to integers only.

+ sum                    integers, floats, complex values, strings
- difference             integers, floats, complex values
* product                integers, floats, complex values
  /    quotient               integers, floats, complex values
  %    remainder              integers
  &    bitwise AND            integers
  |    bitwise OR             integers
  ^    bitwise XOR            integers
  &^   bit clear (AND NOT)    integers
  <<   left shift             integer << integer >= 0
  
  > >   right shift            integer >> integer >= 0

If the operand type is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), the operator must apply to each type in that type set. The operands are represented as values of the type argument that the type parameter is [instantiated](https://go.dev/ref/spec#Instantiations) with, and the operation is computed with the precision of that type argument. For example, given the function:

func dotProduct[F ~float32|~float64](v1, v2 []F) F {
    var s F
    for i, x := range v1 {
        y := v2[i]
        s += x * y
    }
    return s
}

the product `x * y` and the addition `s += x * y` are computed with `float32` or `float64` precision, respectively, depending on the type argument for `F`.

#### Integer operators

For two integer values `x` and `y`, the integer quotient `q = x / y` and remainder `r = x % y` satisfy the following relationships:

x = q*y + r  and  |r| < |y|

with `x / y` truncated towards zero (["truncated division"](https://en.wikipedia.org/wiki/Modulo_operation)).  
其中 `x / y` 被截断为零（“截断除法”）。

 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2

The one exception to this rule is that if the dividend `x` is the most negative value for the int type of `x`, the quotient `q = x / -1` is equal to `x` (and `r = 0`) due to two's-complement [integer overflow](https://go.dev/ref/spec#Integer_overflow):  
此规则的一个例外是，如果被除数 `x` 是 int 类型 `x` 的最大负值，则商 `q = x / -1` 等于 `x` （和 `r = 0` ）由于补码整数溢出：

                         x, q

int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808

If the divisor is a [constant](https://go.dev/ref/spec#Constants), it must not be zero. If the divisor is zero at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. If the dividend is non-negative and the divisor is a constant power of 2, the division may be replaced by a right shift, and computing the remainder may be replaced by a bitwise AND operation:  
如果除数是常数，则它不能为零。如果除数在运行时为零，则会发生运行时恐慌。如果被除数为非负且除数为 2 的常数幂，则除法可以用右移代替，余数的计算可以用按位 AND 运算代替：

 x     x / 4     x % 4     x >> 2     x & 3
 11      2         3         2          3
-11     -2        -3        -3          1

The shift operators shift the left operand by the shift count specified by the right operand, which must be non-negative. If the shift count is negative at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. The shift operators implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer. There is no upper limit on the shift count. Shifts behave as if the left operand is shifted `n` times by 1 for a shift count of `n`. As a result, `x << 1` is the same as `x*2` and `x >> 1` is the same as `x/2` but truncated towards negative infinity.  
移位运算符将左操作数移位右操作数指定的移位计数，该计数必须为非负数。如果运行时班次计数为负，则会发生运行时恐慌。如果左操作数是有符号整数，则移位运算符实现算术移位；如果左操作数是无符号整数，则移位运算符实现逻辑移位。轮班数没有上限。移位的行为就像左操作数移位 `n` 次 1，移位计数为 `n` 。因此， `x << 1` 与 `x*2` 相同， `x >> 1` 与 `x/2` 相同，但被截断为负无穷大。

For integer operands, the unary operators `+`, `-`, and `^` are defined as follows:  
对于整数操作数，一元运算符 `+` 、 `-` 和 `^` 定义如下：

+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x

#### Integer overflow 整数溢出

For [unsigned integer](https://go.dev/ref/spec#Numeric_types) values, the operations `+`, `-`, `*`, and `<<` are computed modulo 2*n*, where *n* is the bit width of the unsigned integer's type. Loosely speaking, these unsigned integer operations discard high bits upon overflow, and programs may rely on "wrap around".  
对于无符号整数值，运算 `+` 、 `-` 、 `*` 和 `<<` 计算模 2 *n* ，其中 n 是无符号整数类型的位宽度。宽松地说，这些无符号整数运算在溢出时会丢弃高位，并且程序可能依赖于“环绕”。

For signed integers, the operations `+`, `-`, `*`, `/`, and `<<` may legally overflow and the resulting value exists and is deterministically defined by the signed integer representation, the operation, and its operands. Overflow does not cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics). A compiler may not optimize code under the assumption that overflow does not occur. For instance, it may not assume that `x < x + 1` is always true.  
对于有符号整数，操作 `+` 、 `-` 、 `*` 、 `/` 和 `<<` 可能会合法溢出，并且结果值存在，并且由有符号整数表示、运算及其操作数确定性地定义。溢出不会导致运行时恐慌。编译器可能不会在不发生溢出的假设下优化代码。例如，它可能不会假设 `x < x + 1` 始终为真。

#### Floating-point operators

For floating-point and complex numbers, `+x` is the same as `x`, while `-x` is the negation of `x`. The result of a floating-point or complex division by zero is not specified beyond the IEEE-754 standard; whether a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs is implementation-specific.

An implementation may combine multiple floating-point operations into a single fused operation, possibly across statements, and produce a result that differs from the value obtained by executing and rounding the instructions individually. An explicit [floating-point type](https://go.dev/ref/spec#Numeric_types) [conversion](https://go.dev/ref/spec#Conversions) rounds to the precision of the target type, preventing fusion that would discard that rounding.

For instance, some architectures provide a "fused multiply and add" (FMA) instruction that computes `x*y + z` without rounding the intermediate result `x*y`. These examples show when a Go implementation can use that instruction:

// FMA allowed for computing r, because x*y is not explicitly rounded:
r  = x*y + z
r  = z;   r += x*y
t  = x*y; r = t + z
*p = x*y; r = *p + z
r  = x*y + float64(z)
// FMA disallowed for computing r, because it would omit rounding of x*y:
r  = float64(x*y) + z
r  = z; r += float64(x*y)
t  = float64(x*y); r = t + z

#### String concatenation

Strings can be concatenated using the `+` operator or the `+=` assignment operator:

s := "hi" + string(c)
s += " and good bye"

String addition creates a new string by concatenating the operands.

### Comparison operators

Comparison operators compare two operands and yield an untyped boolean value.

==    equal
!=    not equal
<     less
<=    less or equal

>     greater
> 
> =    greater or equal

In any comparison, the first operand must be [assignable](https://go.dev/ref/spec#Assignability) to the type of the second operand, or vice versa.

The equality operators `==` and `!=` apply to operands of *comparable* types. The ordering operators `<`, `<=`, `>`, and `>=` apply to operands of *ordered* types. These terms and the result of the comparisons are defined as follows:

- Boolean types are comparable. Two boolean values are equal if they are either both `true` or both `false`.
- Integer types are comparable and ordered. Two integer values are compared in the usual way.
- Floating-point types are comparable and ordered. Two floating-point values are compared as defined by the IEEE-754 standard.
- Complex types are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
- String types are comparable and ordered. Two string values are compared lexically byte-wise.
- Pointer types are comparable. Two pointer values are equal if they point to the same variable or if both have value `nil`. Pointers to distinct [zero-size](https://go.dev/ref/spec#Size_and_alignment_guarantees) variables may or may not be equal.
- Channel types are comparable. Two channel values are equal if they were created by the same call to [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels) or if both have value `nil`.
- Interface types that are not type parameters are comparable. Two interface values are equal if they have [identical](https://go.dev/ref/spec#Type_identity) dynamic types and equal dynamic values or if both have value `nil`.
- A value `x` of non-interface type `X` and a value `t` of interface type `T` can be compared if type `X` is comparable and `X` [implements](https://go.dev/ref/spec#Implementing_an_interface) `T`. They are equal if `t`'s dynamic type is identical to `X` and `t`'s dynamic value is equal to `x`.
- Struct types are comparable if all their field types are comparable. Two struct values are equal if their corresponding non-[blank](https://go.dev/ref/spec#Blank_identifier) field values are equal. The fields are compared in source order, and comparison stops as soon as two field values differ (or all fields have been compared).
- Array types are comparable if their array element types are comparable. Two array values are equal if their corresponding element values are equal. The elements are compared in ascending index order, and comparison stops as soon as two element values differ (or all elements have been compared).
- Type parameters are comparable if they are strictly comparable (see below).

A comparison of two interface values with identical dynamic types causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics) if that type is not comparable. This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.

Slice, map, and function types are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`. Comparison of pointer, channel, and interface values to `nil` is also allowed and follows from the general rules above.

const c = 3 < 4            // c is the untyped boolean constant true
type MyBool bool
var x, y int
var (
    // The result of a comparison is an untyped boolean.
    // The usual assignment rules apply.
    b3        = x == y // b3 has type bool
    b4 bool   = x == y // b4 has type bool
    b5 MyBool = x == y // b5 has type MyBool
)

A type is *strictly comparable* if it is comparable and not an interface type nor composed of interface types. Specifically:

- Boolean, numeric, string, pointer, and channel types are strictly comparable.
- Struct types are strictly comparable if all their field types are strictly comparable.  
  如果结构体类型的所有字段类型都严格可比较，则结构类型也严格可比较。
- Array types are strictly comparable if their array element types are strictly comparable.  
  如果数组元素类型是严格可比较的，则数组类型是严格可比较的。
- Type parameters are strictly comparable if all types in their type set are strictly comparable.  
  如果类型集中的所有类型都是严格可比较的，则类型参数是严格可比较的。

### Logical operators 逻辑运算符

Logical operators apply to [boolean](https://go.dev/ref/spec#Boolean_types) values and yield a result of the same type as the operands. The left operand is evaluated, and then the right if the condition requires it.  
逻辑运算符适用于布尔值并产生与操作数相同类型的结果。计算左操作数，然后计算右操作数（如果条件需要）。

&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"

### Address operators 地址运算符

For an operand `x` of type `T`, the address operation `&x` generates a pointer of type `*T` to `x`. The operand must be *addressable*, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, `x` may also be a (possibly parenthesized) [composite literal](https://go.dev/ref/spec#Composite_literals). If the evaluation of `x` would cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics), then the evaluation of `&x` does too.  
对于 `T` 类型的操作数 `x` ，地址运算 `&x` 会生成一个 `*T` 类型的指向 `x` 的指针。操作数必须是可寻址的，即变量、指针间接或切片索引操作；或可寻址结构操作数的字段选择器；或可寻址数组的数组索引操作。作为可寻址性要求的一个例外， `x` 也可以是（可能带括号的）复合文字。如果 `x` 的计算会导致运行时恐慌，那么 `&x` 的计算也会导致运行时恐慌。

For an operand `x` of pointer type `*T`, the pointer indirection `*x` denotes the [variable](https://go.dev/ref/spec#Variables) of type `T` pointed to by `x`. If `x` is `nil`, an attempt to evaluate `*x` will cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics).  
对于指针类型 `*T` 的操作数 `x` ，指针间接寻址 `*x` 表示 `x` 类型的变量/b4> .如果 `x` 是 `nil` ，则尝试计算 `*x` 将导致运行时恐慌。

&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)
var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic

### Receive operator 接收操作员

For an operand `ch` whose [core type](https://go.dev/ref/spec#Core_types) is a [channel](https://go.dev/ref/spec#Channel_types), the value of the receive operation `<-ch` is the value received from the channel `ch`. The channel direction must permit receive operations, and the type of the receive operation is the element type of the channel. The expression blocks until a value is available. Receiving from a `nil` channel blocks forever. A receive operation on a [closed](https://go.dev/ref/spec#Close) channel can always proceed immediately, yielding the element type's [zero value](https://go.dev/ref/spec#The_zero_value) after any previously sent values have been received.  
对于核心类型为通道的操作数 `ch` 来说，接收操作 `<-ch` 的值是从通道 `ch` 接收到的值。通道方向必须允许接收操作，并且接收操作的类型是通道的元素类型。该表达式会阻塞，直到有值可用为止。从 `nil` 通道接收永远阻塞。关闭通道上的接收操作始终可以立即进行，在收到任何先前发送的值后生成元素类型的零值。

v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // wait until clock pulse and discard received value

A receive expression used in an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or initialization of the special form  
用于赋值语句或特殊形式初始化的接收表达式

x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch

yields an additional untyped boolean result reporting whether the communication succeeded. The value of `ok` is `true` if the value received was delivered by a successful send operation to the channel, or `false` if it is a zero value generated because the channel is closed and empty.  
产生一个额外的无类型布尔结果，报告通信是否成功。如果接收到的值是通过成功的发送操作传递到通道的，则 `ok` 的值为 `true` ；如果由于以下原因生成零值，则 `false` ：通道已关闭并且是空的。

### Conversions 转换

A conversion changes the [type](https://go.dev/ref/spec#Types) of an expression to the type specified by the conversion. A conversion may appear literally in the source, or it may be *implied* by the context in which an expression appears.  
转换将表达式的类型更改为转换指定的类型。转换可能按字面意思出现在源中，也可能由表达式出现的上下文暗示。

An *explicit* conversion is an expression of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type `T`.  
显式转换是 `T(x)` 形式的表达式，其中 `T` 是类型， `x` 是可以转换为类型 `T` .

Conversion = [Type](https://go.dev/ref/spec#Type) "(" [Expression](https://go.dev/ref/spec#Expression) [ "," ] ")" .

If the type starts with the operator `*` or `<-`, or if the type starts with the keyword `func` and has no result list, it must be parenthesized when necessary to avoid ambiguity:  
如果类型以运算符 `*` 或 `<-` 开头，或者如果类型以关键字 `func` 开头并且没有结果列表，则必须在必要时使用括号为避免歧义：

*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
<-chan int(c)    // same as <-(chan int(c))
(<-chan int)(c)  // c is converted to <-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)

A [constant](https://go.dev/ref/spec#Constants) value `x` can be converted to type `T` if `x` is [representable](https://go.dev/ref/spec#Representability) by a value of `T`. As a special case, an integer constant `x` can be explicitly converted to a [string type](https://go.dev/ref/spec#String_types) using the [same rule](https://go.dev/ref/spec#Conversions_to_and_from_a_string_type) as for non-constant `x`.  
如果 `x` 可由值 `T` 表示，则常量值 `x` 可以转换为类型 `T` 。作为一种特殊情况，可以使用与非常量 `x` 相同的规则将整数常量 `x` 显式转换为字符串类型。

Converting a constant to a type that is not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) yields a typed constant.  
将常量转换为非类型参数的类型会生成类型化常量。

uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
myString("foo" + "bar")  // "foobar" of type myString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant

Converting a constant to a type parameter yields a *non-constant* value of that type, with the value represented as a value of the type argument that the type parameter is [instantiated](https://go.dev/ref/spec#Instantiations) with. For example, given the function:

func f[P ~float32|~float64]() {
    … P(1.1) …
}

the conversion `P(1.1)` results in a non-constant value of type `P` and the value `1.1` is represented as a `float32` or a `float64` depending on the type argument for `f`. Accordingly, if `f` is instantiated with a `float32` type, the numeric value of the expression `P(1.1) + 1.2` will be computed with the same precision as the corresponding non-constant `float32` addition.

A non-constant value `x` can be converted to type `T` in any of these cases:

- `x` is [assignable](https://go.dev/ref/spec#Assignability) to `T`.
- ignoring struct tags (see below), `x`'s type and `T` are not [type parameters](https://go.dev/ref/spec#Type_parameter_declarations) but have [identical](https://go.dev/ref/spec#Type_identity) [underlying types](https://go.dev/ref/spec#Underlying_types).
- ignoring struct tags (see below), `x`'s type and `T` are pointer types that are not [named types](https://go.dev/ref/spec#Types), and their pointer base types are not type parameters but have identical underlying types.
- `x`'s type and `T` are both integer or floating point types.
- `x`'s type and `T` are both complex types.
- `x` is an integer or a slice of bytes or runes and `T` is a string type.
- `x` is a string and `T` is a slice of bytes or runes.
- `x` is a slice, `T` is an array [[Go 1.20](https://go.dev/ref/spec#Go_1.20)] or a pointer to an array [[Go 1.17](https://go.dev/ref/spec#Go_1.17)], and the slice and array types have [identical](https://go.dev/ref/spec#Type_identity) element types.

Additionally, if `T` or `x`'s type `V` are type parameters, `x` can also be converted to type `T` if one of the following conditions applies:

- Both `V` and `T` are type parameters and a value of each type in `V`'s type set can be converted to each type in `T`'s type set.
- Only `V` is a type parameter and a value of each type in `V`'s type set can be converted to `T`.
- Only `T` is a type parameter and `x` can be converted to each type in `T`'s type set.

[Struct tags](https://go.dev/ref/spec#Struct_types) are ignored when comparing struct types for identity for the purpose of conversion:

type Person struct {
    Name    string
    Address *struct {
        Street string
        City   string
    }
}
var data *struct {
    Name    string `json:"name"`
    Address *struct {
        Street string `json:"street"`
        City   string `json:"city"`
    } `json:"address"`
}
var person = (*Person)(data)  // ignoring tags, the underlying types are identical

Specific rules apply to (non-constant) conversions between numeric types or to and from a string type. These conversions may change the representation of `x` and incur a run-time cost. All other conversions only change the type but not the representation of `x`.

There is no linguistic mechanism to convert between pointers and integers. The package [`unsafe`](https://go.dev/ref/spec#Package_unsafe) implements this functionality under restricted circumstances.  
没有语言机制可以在指针和整数之间进行转换。包 `unsafe` 在受限情况下实现此功能。

#### Conversions between numeric types

数值类型之间的转换

For the conversion of non-constant numeric values, the following rules apply:  
对于非常量数值的转换，适用以下规则：

1. When converting between [integer types](https://go.dev/ref/spec#Numeric_types), if the value is a signed integer, it is sign extended to implicit infinite precision; otherwise it is zero extended. It is then truncated to fit in the result type's size. For example, if `v := uint16(0x10F0)`, then `uint32(int8(v)) == 0xFFFFFFF0`. The conversion always yields a valid value; there is no indication of overflow.  
   整数类型之间转换时，如果值为有符号整数，则符号扩展为隐式无限精度；否则为零扩展。然后它被截断以适合结果类型的大小。例如，如果 `v := uint16(0x10F0)` ，则 `uint32(int8(v)) == 0xFFFFFFF0` 。转换始终产生有效值；没有溢出的迹象。
2. When converting a [floating-point number](https://go.dev/ref/spec#Numeric_types) to an integer, the fraction is discarded (truncation towards zero).  
   将浮点数转换为整数时，小数部分将被丢弃（向零截断）。
3. When converting an integer or floating-point number to a floating-point type, or a [complex number](https://go.dev/ref/spec#Numeric_types) to another complex type, the result value is rounded to the precision specified by the destination type. For instance, the value of a variable `x` of type `float32` may be stored using additional precision beyond that of an IEEE-754 32-bit number, but float32(x) represents the result of rounding `x`'s value to 32-bit precision. Similarly, `x + 0.1` may use more than 32 bits of precision, but `float32(x + 0.1)` does not.  
   将整数或浮点数转换为浮点类型，或将复数转换为另一种复数类型时，结果值将四舍五入为目标类型指定的精度。例如，类型 `float32` 的变量 `x` 的值可以使用超出 IEEE-754 32 位数字的附加精度来存储，但 float32(x) 表示结果将 `x` 的值舍入为 32 位精度。同样， `x + 0.1` 可以使用超过 32 位的精度，但 `float32(x + 0.1)` 则不会。

In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.  
在涉及浮点或复数值的所有非常量转换中，如果结果类型不能表示转换成功的值，但结果值取决于实现。

#### Conversions to and from a string type

与字符串类型之间的转换

1. Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.  
   将字节切片转换为字符串类型会生成一个字符串，其连续字节是该切片的元素。
   
   string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
   string([]byte{})                                     // ""
   string([]byte(nil))                                  // ""
   type bytes []byte
   string(bytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})    // "hellø"
   type myByte byte
   string([]myByte{'w', 'o', 'r', 'l', 'd', '!'})       // "world!"
   myString([]myByte{'\xf0', '\x9f', '\x8c', '\x8d'})   // "🌍"

2. Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.  
   将符文切片转换为字符串类型会生成一个字符串，该字符串是转换为字符串的各个符文值的串联。
   
   string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
   string([]rune{})                         // ""
   string([]rune(nil))                      // ""
   type runes []rune
   string(runes{0x767d, 0x9d6c, 0x7fd4})    // "\u767d\u9d6c\u7fd4" == "白鵬翔"
   type myRune rune
   string([]myRune{0x266b, 0x266c})         // "\u266b\u266c" == "♫♬"
   myString([]myRune{0x1f30e})              // "\U0001f30e" == "🌎"

3. Converting a value of a string type to a slice of bytes type yields a non-nil slice whose successive elements are the bytes of the string.  
   将字符串类型的值转换为字节类型切片会生成一个非零切片，其连续元素是字符串的字节。
   
   []byte("hellø")             // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   []byte("")                  // []byte{}
   bytes("hellø")              // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   []myByte("world!")          // []myByte{'w', 'o', 'r', 'l', 'd', '!'}
   []myByte(myString("🌏"))    // []myByte{'\xf0', '\x9f', '\x8c', '\x8f'}

4. Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.  
   将字符串类型的值转换为 runes 类型的切片会生成一个包含字符串的各个 Unicode 代码点的切片。
   
   []rune(myString("白鵬翔"))   // []rune{0x767d, 0x9d6c, 0x7fd4}
   []rune("")                  // []rune{}
   runes("白鵬翔")              // []rune{0x767d, 0x9d6c, 0x7fd4}
   []myRune("♫♬")              // []myRune{0x266b, 0x266c}
   []myRune(myString("🌐"))    // []myRune{0x1f310}

5. Finally, for historical reasons, an integer value may be converted to a string type. This form of conversion yields a string containing the (possibly multi-byte) UTF-8 representation of the Unicode code point with the given integer value. Values outside the range of valid Unicode code points are converted to `"\uFFFD"`.
   
   string('a')          // "a"
   string(65)           // "A"
   string('\xf8')       // "\u00f8" == "ø" == "\xc3\xb8"
   string(-1)           // "\ufffd" == "\xef\xbf\xbd"
   type myString string
   myString('\u65e5')   // "\u65e5" == "日" == "\xe6\x97\xa5"
   
   Note: This form of conversion may eventually be removed from the language. The [`go vet`](https://go.dev/pkg/cmd/vet) tool flags certain integer-to-string conversions as potential errors. Library functions such as [`utf8.AppendRune`](https://go.dev/pkg/unicode/utf8#AppendRune) or [`utf8.EncodeRune`](https://go.dev/pkg/unicode/utf8#EncodeRune) should be used instead.  
   最后，由于历史原因，整数值可能会转换为字符串类型。这种形式的转换生成一个字符串，其中包含具有给定整数值的 Unicode 代码点的（可能是多字节）UTF-8 表示形式。有效 Unicode 代码点范围之外的值将转换为 `"\uFFFD"` 。注意：这种形式的转换最终可能会从语言中删除。 `go vet` 工具将某些整数到字符串的转换标记为潜在错误。应改用 `utf8.AppendRune` 或 `utf8.EncodeRune` 等库函数。

#### Conversions from slice to array or array pointer

从切片到数组或数组指针的转换

Converting a slice to an array yields an array containing the elements of the underlying array of the slice. Similarly, converting a slice to an array pointer yields a pointer to the underlying array of the slice. In both cases, if the [length](https://go.dev/ref/spec#Length_and_capacity) of the slice is less than the length of the array, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.  
将切片转换为数组会生成一个包含切片基础数组元素的数组。类似地，将切片转换为数组指针会生成指向切片底层数组的指针。在这两种情况下，如果切片的长度小于数组的长度，则会发生运行时恐慌。

s := make([]byte, 2, 4)
a0 := [0]byte(s)
a1 := [1]byte(s[1:])     // a1[0] == s[1]
a2 := [2]byte(s)         // a2[0] == s[0]
a4 := [4]byte(s)         // panics: len([4]byte) > len(s)
s0 := (*[0]byte)(s)      // s0 != nil
s1 := (*[1]byte)(s[1:])  // &s1[0] == &s[1]
s2 := (*[2]byte)(s)      // &s2[0] == &s[0]
s4 := (*[4]byte)(s)      // panics: len([4]byte) > len(s)
var t []string
t0 := [0]string(t)       // ok for nil slice t
t1 := (*[0]string)(t)    // t1 == nil
t2 := (*[1]string)(t)    // panics: len([1]string) > len(t)
u := make([]byte, 0)
u0 := (*[0]byte)(u)      // u0 != nil

### Constant expressions

Constant expressions may contain only [constant](https://go.dev/ref/spec#Constants) operands and are evaluated at compile time.

Untyped boolean, numeric, and string constants may be used as operands wherever it is legal to use an operand of boolean, numeric, or string type, respectively.

A constant [comparison](https://go.dev/ref/spec#Comparison_operators) always yields an untyped boolean constant. If the left operand of a constant [shift expression](https://go.dev/ref/spec#Operators) is an untyped constant, the result is an integer constant; otherwise it is a constant of the same type as the left operand, which must be of [integer type](https://go.dev/ref/spec#Numeric_types).

Any other operation on untyped constants results in an untyped constant of the same kind; that is, a boolean, integer, floating-point, complex, or string constant. If the untyped operands of a binary operation (other than a shift) are of different kinds, the result is of the operand's kind that appears later in this list: integer, rune, floating-point, complex. For example, an untyped integer constant divided by an untyped complex constant yields an untyped complex constant.  
对无类型常量的任何其他操作都会产生同类的无类型常量；即布尔、整数、浮点、复数或字符串常量。如果二元运算（移位除外）的无类型操作数属于不同类型，则结果是此列表后面出现的操作数类型：整数、符文、浮点、复数。例如，无类型整型常量除以无类型复数常量会产生无类型复数常量。

const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)

Applying the built-in function `complex` to untyped integer, rune, or floating-point constants yields an untyped complex constant.  
将内置函数 `complex` 应用于无类型整数、符文或浮点常量会生成无类型复数常量。

const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)

Constant expressions are always evaluated exactly; intermediate values and the constants themselves may require precision significantly larger than supported by any predeclared type in the language. The following are legal declarations:  
常量表达式总是被精确计算；中间值和常量本身可能需要比语言中任何预声明类型支持的精度大得多的精度。以下为法律声明：

const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge >> 98  // Four == 4                                (type int8)

The divisor of a constant division or remainder operation must not be zero:  
常量除法或余数运算的除数不得为零：

3.14 / 0.0   // illegal: division by zero

The values of *typed* constants must always be accurately [representable](https://go.dev/ref/spec#Representability) by values of the constant type. The following constant expressions are illegal:  
类型常量的值必须始终可以由常量类型的值准确表示。以下常量表达式是非法的：

uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)

The mask used by the unary bitwise complement operator `^` matches the rule for non-constants: the mask is all 1s for unsigned constants and -1 for signed and untyped constants.  
一元按位求补运算符 `^` 使用的掩码与非常量的规则匹配：对于无符号常量，掩码为全 1；对于有符号和无类型常量，掩码为 -1。

^1         // untyped integer constant, equal to -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2

Implementation restriction: A compiler may use rounding while computing untyped floating-point or complex constant expressions; see the implementation restriction in the section on [constants](https://go.dev/ref/spec#Constants). This rounding may cause a floating-point constant expression to be invalid in an integer context, even if it would be integral when calculated using infinite precision, and vice versa.

### Order of evaluation

At package level, [initialization dependencies](https://go.dev/ref/spec#Package_initialization) determine the evaluation order of individual initialization expressions in [variable declarations](https://go.dev/ref/spec#Variable_declarations). Otherwise, when evaluating the [operands](https://go.dev/ref/spec#Operands) of an expression, assignment, or [return statement](https://go.dev/ref/spec#Return_statements), all function calls, method calls, [receive operations](https://go.dev/ref/spec#Receive%20operator), and [binary logical operations](https://go.dev/ref/spec#Logical_operators) are evaluated in lexical left-to-right order.

For example, in the (function-local) assignment

y[f()], ok = g(z || h(), i()+x[j()], <-c), k()

the function calls and communication happen in the order `f()`, `h()` (if `z` evaluates to false), `i()`, `j()`, `<-c`, `g()`, and `k()`. However, the order of those events compared to the evaluation and indexing of `x` and the evaluation of `y` and `z` is not specified, except as required lexically. For instance, `g` cannot be called before its arguments are evaluated.

a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified

At package level, initialization dependencies override the left-to-right rule for individual initialization expressions, but not for operands within each expression:

var a, b, c = f() + v(), g(), sqr(u()) + v()
func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }
// functions u and v are independent of all other variables and functions

The function calls happen in the order `u()`, `sqr()`, `v()`, `f()`, `v()`, and `g()`.

Floating-point operations within a single expression are evaluated according to the associativity of the operators. Explicit parentheses affect the evaluation by overriding the default associativity. In the expression `x + (y + z)` the addition `y + z` is performed before adding `x`.

## Statements

Statements control execution.

Statement =
    [Declaration](https://go.dev/ref/spec#Declaration) | [LabeledStmt](https://go.dev/ref/spec#LabeledStmt) | [SimpleStmt](https://go.dev/ref/spec#SimpleStmt) |
    [GoStmt](https://go.dev/ref/spec#GoStmt) | [ReturnStmt](https://go.dev/ref/spec#ReturnStmt) | [BreakStmt](https://go.dev/ref/spec#BreakStmt) | [ContinueStmt](https://go.dev/ref/spec#ContinueStmt) | [GotoStmt](https://go.dev/ref/spec#GotoStmt) |
    [FallthroughStmt](https://go.dev/ref/spec#FallthroughStmt) | [Block](https://go.dev/ref/spec#Block) | [IfStmt](https://go.dev/ref/spec#IfStmt) | [SwitchStmt](https://go.dev/ref/spec#SwitchStmt) | [SelectStmt](https://go.dev/ref/spec#SelectStmt) | [ForStmt](https://go.dev/ref/spec#ForStmt) |
    [DeferStmt](https://go.dev/ref/spec#DeferStmt) .

SimpleStmt = [EmptyStmt](https://go.dev/ref/spec#EmptyStmt) | [ExpressionStmt](https://go.dev/ref/spec#ExpressionStmt) | [SendStmt](https://go.dev/ref/spec#SendStmt) | [IncDecStmt](https://go.dev/ref/spec#IncDecStmt) | [Assignment](https://go.dev/ref/spec#Assignment) | [ShortVarDecl](https://go.dev/ref/spec#ShortVarDecl) .

### Terminating statements

A *terminating statement* interrupts the regular flow of control in a [block](https://go.dev/ref/spec#Blocks). The following statements are terminating:

1. A ["return"](https://go.dev/ref/spec#Return_statements) or ["goto"](https://go.dev/ref/spec#Goto_statements) statement.

2. A call to the built-in function [`panic`](https://go.dev/ref/spec#Handling_panics).

3. A [block](https://go.dev/ref/spec#Blocks) in which the statement list ends in a terminating statement.

4. An ["if" statement](https://go.dev/ref/spec#If_statements) in which:
   
   - the "else" branch is present, and
   - both branches are terminating statements.

5. A ["for" statement](https://go.dev/ref/spec#For_statements) in which:
   
   - there are no "break" statements referring to the "for" statement, and
   - the loop condition is absent, and
   - the "for" statement does not use a range clause.

6. A ["switch" statement](https://go.dev/ref/spec#Switch_statements) in which:
   
   - there are no "break" statements referring to the "switch" statement,
   - there is a default case, and
   - the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled ["fallthrough" statement](https://go.dev/ref/spec#Fallthrough_statements).

7. A ["select" statement](https://go.dev/ref/spec#Select_statements) in which:
   
   - there are no "break" statements referring to the "select" statement, and
   - the statement lists in each case, including the default if present, end in a terminating statement.

8. A [labeled statement](https://go.dev/ref/spec#Labeled_statements) labeling a terminating statement.

All other statements are not terminating.

A [statement list](https://go.dev/ref/spec#Blocks) ends in a terminating statement if the list is not empty and its final non-empty statement is terminating.

### Empty statements

The empty statement does nothing.

EmptyStmt = .

### Labeled statements

A labeled statement may be the target of a `goto`, `break` or `continue` statement.

LabeledStmt = [Label](https://go.dev/ref/spec#Label) ":" [Statement](https://go.dev/ref/spec#Statement) .
Label       = [identifier](https://go.dev/ref/spec#identifier) .

Error: log.Panic("error encountered")

### Expression statements 表达式语句

With the exception of specific built-in functions, function and method [calls](https://go.dev/ref/spec#Calls) and [receive operations](https://go.dev/ref/spec#Receive_operator) can appear in statement context. Such statements may be parenthesized.  
除了特定的内置函数之外，函数和方法调用以及接收操作可以出现在语句上下文中。此类陈述可以加括号。

ExpressionStmt = [Expression](https://go.dev/ref/spec#Expression) .

The following built-in functions are not permitted in statement context:  
语句上下文中不允许使用以下内置函数：

append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice unsafe.SliceData unsafe.String unsafe.StringData

h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function

### Send statements 发送报表

A send statement sends a value on a channel. The channel expression's [core type](https://go.dev/ref/spec#Core_types) must be a [channel](https://go.dev/ref/spec#Channel_types), the channel direction must permit send operations, and the type of the value to be sent must be [assignable](https://go.dev/ref/spec#Assignability) to the channel's element type.  
发送语句在通道上发送一个值。通道表达式的核心类型必须是通道，通道方向必须允许发送操作，并且要发送的值的类型必须可分配给通道的元素类型。

SendStmt = [Channel](https://go.dev/ref/spec#Channel) "<-" [Expression](https://go.dev/ref/spec#Expression) .
Channel  = [Expression](https://go.dev/ref/spec#Expression) .

Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a [run-time panic](https://go.dev/ref/spec#Run_time_panics). A send on a `nil` channel blocks forever.  
在通信开始之前，将对通道和值表达式进行评估。通信将阻塞，直到发送可以继续。如果接收器准备就绪，则可以在无缓冲通道上进行发送。如果缓冲区中有空间，则可以继续在缓冲通道上发送。关闭通道上的发送会导致运行时恐慌。 `nil` 通道上的发送将永远阻塞。

ch <- 3  // send value 3 to channel ch

### IncDec statements IncDec 报表

The "++" and "--" statements increment or decrement their operands by the untyped [constant](https://go.dev/ref/spec#Constants) `1`. As with an assignment, the operand must be [addressable](https://go.dev/ref/spec#Address_operators) or a map index expression.  
“++”和“--”语句通过无类型常量 `1` 增加或减少其操作数。与赋值一样，操作数必须是可寻址的或映射索引表达式。

IncDecStmt = [Expression](https://go.dev/ref/spec#Expression) ( "++" | "--" ) .

The following [assignment statements](https://go.dev/ref/spec#Assignment_statements) are semantically equivalent:  
以下赋值语句在语义上是等效的：

IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1

### Assignment statements 赋值语句

An *assignment* replaces the current value stored in a [variable](https://go.dev/ref/spec#Variables) with a new value specified by an [expression](https://go.dev/ref/spec#Expressions). An assignment statement may assign a single value to a single variable, or multiple values to a matching number of variables.  
赋值将变量中存储的当前值替换为表达式指定的新值。赋值语句可以将单个值分配给单个变量，或者将多个值分配给匹配数量的变量。

Assignment = [ExpressionList](https://go.dev/ref/spec#ExpressionList) [assign_op](https://go.dev/ref/spec#assign_op) [ExpressionList](https://go.dev/ref/spec#ExpressionList) .

assign_op = [ [add_op](https://go.dev/ref/spec#add_op) | [mul_op](https://go.dev/ref/spec#mul_op) ] "=" .

Each left-hand side operand must be [addressable](https://go.dev/ref/spec#Address_operators), a map index expression, or (for `=` assignments only) the [blank identifier](https://go.dev/ref/spec#Blank_identifier). Operands may be parenthesized.  
每个左侧操作数必须是可寻址的、映射索引表达式或（仅对于 `=` 赋值）空白标识符。操作数可以用括号括起来。

x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch

An *assignment operation* `x` *op*`=` `y` where *op* is a binary [arithmetic operator](https://go.dev/ref/spec#Arithmetic_operators) is equivalent to `x` `=` `x` *op* `(y)` but evaluates `x` only once. The *op*`=` construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier.  
赋值操作 `x` op `=` `y` 其中 op 是二元算术运算符，相当于 `x` `=` < b5> op `(y)` 但仅计算 `x` 一次。 op `=` 构造是单个标记。在赋值运算中，左侧表达式列表和右侧表达式列表都必须恰好包含一个单值表达式，并且左侧表达式不能是空白标识符。

a[i] <<= 2
i &^= 1<<n

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a [channel](https://go.dev/ref/spec#Channel_types) or [map](https://go.dev/ref/spec#Map_types) operation, or a [type assertion](https://go.dev/ref/spec#Type_assertions). The number of operands on the left hand side must match the number of values. For instance, if `f` is a function returning two values,

x, y = f()

assigns the first value to `x` and the second to `y`. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the *n*th expression on the right is assigned to the *n*th operand on the left:

one, two, three = '一', '二', '三'

The [blank identifier](https://go.dev/ref/spec#Blank_identifier) provides a way to ignore right-hand side values in an assignment:  
空白标识符提供了一种忽略赋值中右侧值的方法：

_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value

The assignment proceeds in two phases. First, the operands of [index expressions](https://go.dev/ref/spec#Index_expressions) and [pointer indirections](https://go.dev/ref/spec#Address_operators) (including implicit pointer indirections in [selectors](https://go.dev/ref/spec#Selectors)) on the left and the expressions on the right are all [evaluated in the usual order](https://go.dev/ref/spec#Order_of_evaluation). Second, the assignments are carried out in left-to-right order.  
任务分两个阶段进行。首先，左边的索引表达式和指针间接（包括选择器中的隐式指针间接）的操作数和右边的表达式都按通常的顺序求值。其次，作业按从左到右的顺序进行。

a, b = b, a  // exchange a and b
x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2
i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1
x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)
x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.
type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7
i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
    break
}
// after this loop, i == 0 and x is []int{3, 5, 3}

In assignments, each value must be [assignable](https://go.dev/ref/spec#Assignability) to the type of the operand to which it is assigned, with the following special cases:

1. Any typed value may be assigned to the blank identifier.
2. If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first implicitly [converted](https://go.dev/ref/spec#Conversions) to its [default type](https://go.dev/ref/spec#Constants).
3. If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first implicitly converted to type `bool`.  
   如果将无类型布尔值分配给接口类型或空白标识符的变量，则它首先会隐式转换为类型 `bool` 。

### If statements 如果语句

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.  
“If”语句根据布尔表达式的值指定两个分支的条件执行。如果表达式计算结果为真，则执行“if”分支，否则，如果存在，则执行“else”分支。

IfStmt = "if" [ [SimpleStmt](https://go.dev/ref/spec#SimpleStmt) ";" ] [Expression](https://go.dev/ref/spec#Expression) [Block](https://go.dev/ref/spec#Block) [ "else" ( [IfStmt](https://go.dev/ref/spec#IfStmt) | [Block](https://go.dev/ref/spec#Block) ) ] .

if x > max {
    x = max
}

The expression may be preceded by a simple statement, which executes before the expression is evaluated.  
表达式前面可以有一个简单的语句，该语句在计算表达式之前执行。

if x := f(); x < y {
    return x
} else if x > z {
    return z
} else {
    return y
}

### Switch statements Switch 语句

"Switch" statements provide multi-way execution. An expression or type is compared to the "cases" inside the "switch" to determine which branch to execute.  
“Switch”语句提供多路执行。将表达式或类型与“switch”内的“cases”进行比较，以确定要执行哪个分支。

SwitchStmt = [ExprSwitchStmt](https://go.dev/ref/spec#ExprSwitchStmt) | [TypeSwitchStmt](https://go.dev/ref/spec#TypeSwitchStmt) .

There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.  
有两种形式：表达式开关和类型开关。在表达式 switch 中，case 包含与 switch 表达式的值进行比较的表达式。在类型开关中，cases 包含与特殊注释的开关表达式的类型进行比较的类型。 switch 表达式在 switch 语句中只计算一次。

#### Expression switches 表情开关

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value `true`.  
在表达式 switch 中，对 switch 表达式进行求值，而 case 表达式（不必是常量）按从左到右、从上到下的顺序求值；第一个等于 switch 表达式的触发执行关联 case 的语句；其他情况被跳过。如果没有 case 匹配并且存在“default”case，则执行其语句。最多可以有一种默认情况，它可以出现在“switch”语句中的任何位置。缺少的 switch 表达式相当于布尔值 `true` 。

ExprSwitchStmt = "switch" [ [SimpleStmt](https://go.dev/ref/spec#SimpleStmt) ";" ] [ [Expression](https://go.dev/ref/spec#Expression) ] "{" { [ExprCaseClause](https://go.dev/ref/spec#ExprCaseClause) } "}" .
ExprCaseClause = [ExprSwitchCase](https://go.dev/ref/spec#ExprSwitchCase) ":" [StatementList](https://go.dev/ref/spec#StatementList) .
ExprSwitchCase = "case" [ExpressionList](https://go.dev/ref/spec#ExpressionList) | "default" .

If the switch expression evaluates to an untyped constant, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to its [default type](https://go.dev/ref/spec#Constants). The predeclared untyped value `nil` cannot be used as a switch expression. The switch expression type must be [comparable](https://go.dev/ref/spec#Comparison_operators).  
如果 switch 表达式的计算结果为无类型常量，则它首先会隐式转换为其默认类型。预声明的无类型值 `nil` 不能用作 switch 表达式。 switch 表达式类型必须是可比较的。

If a case expression is untyped, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to the type of the switch expression. For each (possibly converted) case expression `x` and the value `t` of the switch expression, `x == t` must be a valid [comparison](https://go.dev/ref/spec#Comparison_operators).  
如果 case 表达式是无类型的，则它首先会隐式转换为 switch 表达式的类型。对于每个（可能已转换的）case 表达式 `x` 和 switch 表达式的值 `t` ， `x == t` 必须是有效的比较。

In other words, the switch expression is treated as if it were used to declare and initialize a temporary variable `t` without explicit type; it is that value of `t` against which each case expression `x` is tested for equality.  
换句话说，switch 表达式被视为用于声明和初始化临时变量 `t` ，而无需显式类型；它是 `t` 的值，每个 case 表达式 `x` 都根据该值进行相等性测试。

In a case or default clause, the last non-empty statement may be a (possibly [labeled](https://go.dev/ref/spec#Labeled_statements)) ["fallthrough" statement](https://go.dev/ref/spec#Fallthrough_statements) to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.  
在 case 或 default 子句中，最后一个非空语句可能是（可能标记为）“fallthrough”语句，以指示控制应从本子句的末尾流到下一个子句的第一个语句。否则控制流到“switch”语句的末尾。 “fallthrough”语句可能会作为表达式 switch 的最后一个子句之外的所有语句的最后一个语句出现。

The switch expression may be preceded by a simple statement, which executes before the expression is evaluated.  
switch 表达式前面可以有一个简单的语句，该语句在表达式求值之前执行。

switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}
switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}
switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}

Implementation restriction: A compiler may disallow multiple case expressions evaluating to the same constant. For instance, the current compilers disallow duplicate integer, floating point, or string constants in case expressions.

#### Type switches

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a [type assertion](https://go.dev/ref/spec#Type_assertions) using the keyword `type` rather than an actual type:

switch x.(type) {
// cases
}

Cases then match actual types `T` against the dynamic type of the expression `x`. As with type assertions, `x` must be of [interface type](https://go.dev/ref/spec#Interface_types), but not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), and each non-interface type `T` listed in a case must implement the type of `x`. The types listed in the cases of a type switch must all be [different](https://go.dev/ref/spec#Type_identity).

TypeSwitchStmt  = "switch" [ [SimpleStmt](https://go.dev/ref/spec#SimpleStmt) ";" ] [TypeSwitchGuard](https://go.dev/ref/spec#TypeSwitchGuard) "{" { [TypeCaseClause](https://go.dev/ref/spec#TypeCaseClause) } "}" .
TypeSwitchGuard = [ [identifier](https://go.dev/ref/spec#identifier) ":=" ] [PrimaryExpr](https://go.dev/ref/spec#PrimaryExpr) "." "(" "type" ")" .
TypeCaseClause  = [TypeSwitchCase](https://go.dev/ref/spec#TypeSwitchCase) ":" [StatementList](https://go.dev/ref/spec#StatementList) .
TypeSwitchCase  = "case" [TypeList](https://go.dev/ref/spec#TypeList) | "default" .

The TypeSwitchGuard may include a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations). When that form is used, the variable is declared at the end of the TypeSwitchCase in the [implicit block](https://go.dev/ref/spec#Blocks) of each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.

Instead of a type, a case may use the predeclared identifier [`nil`](https://go.dev/ref/spec#Predeclared_identifiers); that case is selected when the expression in the TypeSwitchGuard is a `nil` interface value. There may be at most one `nil` case.

Given an expression `x` of type `interface{}`, the following type switch:

switch i := x.(type) {
case nil:
    printString("x is nil")                // type of i is type of x (interface{})
case int:
    printInt(i)                            // type of i is int
case float64:
    printFloat64(i)                        // type of i is float64
case func(int) float64:
    printFunction(i)                       // type of i is func(int) float64
case bool, string:
    printString("type is bool or string")  // type of i is type of x (interface{})
default:
    printString("don't know the type")     // type of i is type of x (interface{})
}

could be rewritten:

v := x  // x is evaluated exactly once
if v == nil {
    i := v                                 // type of i is type of x (interface{})
    printString("x is nil")
} else if i, isInt := v.(int); isInt {
    printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
    printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
    printFunction(i)                       // type of i is func(int) float64
} else {
    _, isBool := v.(bool)
    _, isString := v.(string)
    if isBool || isString {
        i := v                         // type of i is type of x (interface{})
        printString("type is bool or string")
    } else {
        i := v                         // type of i is type of x (interface{})
        printString("don't know the type")
    }
}

A [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) or a [generic type](https://go.dev/ref/spec#Type_declarations) may be used as a type in a case. If upon [instantiation](https://go.dev/ref/spec#Instantiations) that type turns out to duplicate another entry in the switch, the first matching case is chosen.  
类型参数或泛型类型可以用作情况中的类型。如果在实例化时该类型与交换机中的另一个条目重复，则选择第一个匹配的情况。

func f[P any](x any) int {
    switch x.(type) {
    case P:
        return 0
    case string:
        return 1
    case []P:
        return 2
    case []byte:
        return 3
    default:
        return 4
    }
}
var v1 = f[string]("foo")   // v1 == 0
var v2 = f[byte]([]byte{})  // v2 == 2

The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.  
类型开关防护前面可以有一个简单的语句，该语句在评估防护之前执行。

The "fallthrough" statement is not permitted in a type switch.  
类型开关中不允许使用“fallthrough”语句。

### For statements 对于报表

A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.  
“for”语句指定重复执行块。有三种形式： 迭代可以由单个条件、“for”子句或“range”子句控制。

ForStmt = "for" [ [Condition](https://go.dev/ref/spec#Condition) | [ForClause](https://go.dev/ref/spec#ForClause) | [RangeClause](https://go.dev/ref/spec#RangeClause) ] [Block](https://go.dev/ref/spec#Block) .
Condition = [Expression](https://go.dev/ref/spec#Expression) .

#### For statements with single condition

对于具有单一条件的语句

In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value `true`.  
在最简单的形式中，“for”语句指定只要布尔条件计算结果为真，就重复执行块。在每次迭代之前都会评估条件。如果条件不存在，则相当于布尔值 `true` 。

for a < b {
    a *= 2
}

#### For statements with `for` clause

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an *init* and a *post* statement, such as an assignment, an increment or decrement statement. The init statement may be a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations), but the post statement must not.

ForClause = [ [InitStmt](https://go.dev/ref/spec#InitStmt) ] ";" [ [Condition](https://go.dev/ref/spec#Condition) ] ";" [ [PostStmt](https://go.dev/ref/spec#PostStmt) ] .
InitStmt = [SimpleStmt](https://go.dev/ref/spec#SimpleStmt) .
PostStmt = [SimpleStmt](https://go.dev/ref/spec#SimpleStmt) .

for i := 0; i < 10; i++ {
    f(i)
}

If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the [semicolons](https://go.dev/ref/spec#Semicolons) are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value `true`.

for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }

Each iteration has its own separate declared variable (or variables) [[Go 1.22](https://go.dev/ref/spec#Go_1.22)]. The variable used by the first iteration is declared by the init statement. The variable used by each subsequent iteration is declared implicitly before executing the post statement and initialized to the value of the previous iteration's variable at that moment.

var prints []func()
for i := 0; i < 5; i++ {
    prints = append(prints, func() { println(i) })
    i++
}
for _, p := range prints {
    p()
}

prints

1
3
5

Prior to [[Go 1.22](https://go.dev/ref/spec#Go_1.22)], iterations share one set of variables instead of having their own separate variables. In that case, the example above prints

6
6
6

#### For statements with `range` clause

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, values received on a channel, or integer values from zero to an upper limit [[Go 1.22](https://go.dev/ref/spec#Go_1.22)]. For each entry it assigns *iteration values* to corresponding *iteration variables* if present and then executes the block.

RangeClause = [ [ExpressionList](https://go.dev/ref/spec#ExpressionList) "=" | [IdentifierList](https://go.dev/ref/spec#IdentifierList) ":=" ] "range" [Expression](https://go.dev/ref/spec#Expression) .

The expression on the right in the "range" clause is called the *range expression*, its [core type](https://go.dev/ref/spec#Core_types) must be an array, pointer to an array, slice, string, map, channel permitting [receive operations](https://go.dev/ref/spec#Receive_operator), or an integer. As with an assignment, if present the operands on the left must be [addressable](https://go.dev/ref/spec#Address_operators) or map index expressions; they denote the iteration variables. If the range expression is a channel or integer, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the [blank identifier](https://go.dev/ref/spec#Blank_identifier), the range clause is equivalent to the same clause without that identifier.

The range expression `x` is evaluated once before beginning the loop, with one exception: if at most one iteration variable is present and `len(x)` is [constant](https://go.dev/ref/spec#Length_and_capacity), the range expression is not evaluated.

Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

Range expression                          1st value          2nd value
array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
integer         n  integer type           value    i  see below

1. For an array, pointer to array, or slice value `a`, the index iteration values are produced in increasing order, starting at element index 0. If at most one iteration variable is present, the range loop produces iteration values from 0 up to `len(a)-1` and does not index into the array or slice itself. For a `nil` slice, the number of iterations is 0.
2. For a string value, the "range" clause iterates over the Unicode code points in the string starting at byte index 0. On successive iterations, the index value will be the index of the first byte of successive UTF-8-encoded code points in the string, and the second value, of type `rune`, will be the value of the corresponding code point. If the iteration encounters an invalid UTF-8 sequence, the second value will be `0xFFFD`, the Unicode replacement character, and the next iteration will advance a single byte in the string.  
   对于字符串值，“range”子句从字节索引 0 开始迭代字符串中的 Unicode 代码点。在连续迭代中，索引值将是连续 UTF-8 编码代码点的第一个字节的索引字符串和 `rune` 类型的第二个值将是相应代码点的值。如果迭代遇到无效的 UTF-8 序列，则第二个值将为 `0xFFFD` （Unicode 替换字符），并且下一次迭代将在字符串中前进一个字节。
3. The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next. If a map entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced. If a map entry is created during iteration, that entry may be produced during the iteration or may be skipped. The choice may vary for each entry created and from one iteration to the next. If the map is `nil`, the number of iterations is 0.  
   映射的迭代顺序未指定，并且不保证从一次迭代到下一次迭代的顺序相同。如果在迭代过程中删除了尚未到达的映射条目，则不会产生相应的迭代值。如果在迭代期间创建映射条目，则该条目可以在迭代期间产生或者可以被跳过。对于创建的每个条目以及从一次迭代到下一次迭代，选择可能会有所不同。如果映射为 `nil` ，则迭代次数为 0。
4. For channels, the iteration values produced are the successive values sent on the channel until the channel is [closed](https://go.dev/ref/spec#Close). If the channel is `nil`, the range expression blocks forever.  
   对于通道，生成的迭代值是在通道关闭之前在通道上发送的连续值。如果通道是 `nil` ，则范围表达式将永远阻塞。
5. For an integer value `n`, the iteration values 0 through `n-1` are produced in increasing order. If `n` <= 0, the loop does not run any iterations.  
   对于整数值 `n` ，迭代值 0 到 `n-1` 按升序生成。如果 `n` <= 0，则循环不运行任何迭代。

The iteration variables may be declared by the "range" clause using a form of [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations) (`:=`). In this case their [scope](https://go.dev/ref/spec#Declarations_and_scope) is the block of the "for" statement and each iteration has its own new variables [[Go 1.22](https://go.dev/ref/spec#Go_1.22)] (see also ["for" statements with a ForClause](https://go.dev/ref/spec#For_clause)). If the range expression is a (possibly untyped) integer expression `n`, the variable has the same type as if it was [declared](https://go.dev/ref/spec#Variable_declarations) with initialization expression `n`. Otherwise, the variables have the types of their respective iteration values.  
迭代变量可以通过“range”子句使用短变量声明（ `:=` ）的形式来声明。在这种情况下，它们的范围是“for”语句的块，并且每次迭代都有自己的新变量[Go 1.22]（另请参见带有 ForClause 的“for”语句）。如果范围表达式是（可能是无类型的）整数表达式 `n` ，则该变量具有与使用初始化表达式 `n` 声明的类型相同的类型。否则，变量具有其各自迭代值的类型。

If the iteration variables are not explicitly declared by the "range" clause, they must be preexisting. In this case, the iteration values are assigned to the respective variables as in an [assignment statement](https://go.dev/ref/spec#Assignment_statements). If the range expression is a (possibly untyped) integer expression `n`, `n` too must be [assignable](https://go.dev/ref/spec#Assignability) to the iteration variable; if there is no iteration variable, `n` must be assignable to `int`.  
如果迭代变量没有通过“range”子句显式声明，则它们必须是预先存在的。在这种情况下，迭代值将像赋值语句一样分配给各个变量。如果范围表达式是（可能是无类型的）整数表达式 `n` ，则 `n` 也必须可分配给迭代变量；如果没有迭代变量，则 `n` 必须可分配给 `int` 。

var testdata *struct {
    a *[7]int
}
for i, _ := range testdata.a {
    // testdata.a is never evaluated; len(testdata.a) is constant
    // i ranges from 0 to 6
    f(i)
}
var a [10]string
for i, s := range a {
    // type of i is int
    // type of s is string
    // s == a[i]
    g(i, s)
}
var key string
var val interface{}  // element type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
    h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]
var ch chan Work = producer()
for w := range ch {
    doWork(w)
}
// empty a channel
for range ch {}
// call f(0), f(1), ... f(9)
for i := range 10 {
    // type of i is int (default type for untyped constant 10)
    f(i)
}
// invalid: 256 cannot be assigned to uint8
var u uint8
for u = range 256 {
}

### Go statements

A "go" statement starts the execution of a function call as an independent concurrent thread of control, or *goroutine*, within the same address space.

GoStmt = "go" [Expression](https://go.dev/ref/spec#Expression) .

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](https://go.dev/ref/spec#Expression_statements).

The function value and parameters are [evaluated as usual](https://go.dev/ref/spec#Calls) in the calling goroutine, but unlike with a regular call, program execution does not wait for the invoked function to complete. Instead, the function begins executing independently in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values, they are discarded when the function completes.

go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)

### Select statements 选择语句

A "select" statement chooses which of a set of possible [send](https://go.dev/ref/spec#Send_statements) or [receive](https://go.dev/ref/spec#Receive_operator) operations will proceed. It looks similar to a ["switch"](https://go.dev/ref/spec#Switch_statements) statement but with the cases all referring to communication operations.  
“选择”语句选择将进行一组可能的发送或接收操作中的哪一个。它看起来类似于“switch”语句，但所有情况都指的是通信操作。

SelectStmt = "select" "{" { [CommClause](https://go.dev/ref/spec#CommClause) } "}" .
CommClause = [CommCase](https://go.dev/ref/spec#CommCase) ":" [StatementList](https://go.dev/ref/spec#StatementList) .
CommCase   = "case" ( [SendStmt](https://go.dev/ref/spec#SendStmt) | [RecvStmt](https://go.dev/ref/spec#RecvStmt) ) | "default" .
RecvStmt   = [ [ExpressionList](https://go.dev/ref/spec#ExpressionList) "=" | [IdentifierList](https://go.dev/ref/spec#IdentifierList) ":=" ] [RecvExpr](https://go.dev/ref/spec#RecvExpr) .
RecvExpr   = [Expression](https://go.dev/ref/spec#Expression) .

A case with a RecvStmt may assign the result of a RecvExpr to one or two variables, which may be declared using a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations). The RecvExpr must be a (possibly parenthesized) receive operation. There can be at most one default case and it may appear anywhere in the list of cases.

Execution of a "select" statement proceeds in several steps:

1. For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
2. If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
3. Unless the selected case is the default case, the respective communication operation is executed.
4. If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
5. The statement list of the selected case is executed.

Since communication on `nil` channels can never proceed, a select with only `nil` channels and no default case blocks forever.

var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
    print("received ", i1, " from c1\n")
case c2 <- i2:
    print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
    if ok {
        print("received ", i3, " from c3\n")
    } else {
        print("c3 is closed\n")
    }
case a[f()] = <-c4:
    // same as:
    // case t := <-c4
    //    a[f()] = t
default:
    print("no communication\n")
}
for {  // send random sequence of bits to c
    select {
    case c <- 0:  // note: no statement, no fallthrough, no folding of cases
    case c <- 1:
    }
}
select {}  // block forever

### Return statements 退货声明

A "return" statement in a function `F` terminates the execution of `F`, and optionally provides one or more result values. Any functions [deferred](https://go.dev/ref/spec#Defer_statements) by `F` are executed before `F` returns to its caller.  
函数 `F` 中的“return”语句终止 `F` 的执行，并可选择提供一个或多个结果值。任何由 `F` 推迟的函数都会在 `F` 返回其调用者之前执行。

ReturnStmt = "return" [ [ExpressionList](https://go.dev/ref/spec#ExpressionList) ] .

In a function without a result type, a "return" statement must not specify any result values.  
在没有结果类型的函数中，“return”语句不得指定任何结果值。

func noResult() {
    return
}

There are three ways to return values from a function with a result type:  
有三种方法可以从具有结果类型的函数返回值：

1. The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and [assignable](https://go.dev/ref/spec#Assignability) to the corresponding element of the function's result type.  
   返回值可以在“return”语句中明确列出。每个表达式必须是单值并且可分配给函数结果类型的相应元素。
   
   func simpleF() int {
      return 2
   }
   func complexF1() (re float64, im float64) {
      return -7.0, -4.0
   }

2. The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply.  
   “return”语句中的表达式列表可以是对多值函数的单个调用。效果就好像从该函数返回的每个值都被分配给具有相应值类型的临时变量，后跟列出这些变量的“return”语句，此时应用前一种情况的规则。
   
   func complexF2() (re float64, im float64) {
      return complexF1()
   }

3. The expression list may be empty if the function's result type specifies names for its [result parameters](https://go.dev/ref/spec#Function_types). The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables.  
   如果函数的结果类型指定了其结果参数的名称，则表达式列表可能为空。结果参数充当普通的局部变量，函数可以根据需要为其赋值。 “return”语句返回这些变量的值。
   
   func complexF3() (re float64, im float64) {
      re = 7.0
      im = 4.0
      return
   }
   func (devnull) Write(p []byte) (n int, _ error) {
      n = len(p)
      return
   }

Regardless of how they are declared, all the result values are initialized to the [zero values](https://go.dev/ref/spec#The_zero_value) for their type upon entry to the function. A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

Implementation restriction: A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) with the same name as a result parameter is in [scope](https://go.dev/ref/spec#Declarations_and_scope) at the place of the return.

func f(n int) (res int, err error) {
    if _, err := f(n-1); err != nil {
        return  // invalid return statement: err is shadowed
    }
    return
}

### Break statements

A "break" statement terminates execution of the innermost ["for"](https://go.dev/ref/spec#For_statements), ["switch"](https://go.dev/ref/spec#Switch_statements), or ["select"](https://go.dev/ref/spec#Select_statements) statement within the same function.

BreakStmt = "break" [ [Label](https://go.dev/ref/spec#Label) ] .

If there is a label, it must be that of an enclosing "for", "switch", or "select" statement, and that is the one whose execution terminates.

OuterLoop:
    for i = 0; i < n; i++ {
        for j = 0; j < m; j++ {
            switch a[i][j] {
            case nil:
                state = Error
                break OuterLoop
            case item:
                state = Found
                break OuterLoop
            }
        }
    }

### Continue statements

A "continue" statement begins the next iteration of the innermost enclosing ["for" loop](https://go.dev/ref/spec#For_statements) by advancing control to the end of the loop block. The "for" loop must be within the same function.

ContinueStmt = "continue" [ [Label](https://go.dev/ref/spec#Label) ] .

If there is a label, it must be that of an enclosing "for" statement, and that is the one whose execution advances.

RowLoop:
    for y, row := range rows {
        for x, data := range row {
            if data == endOfRow {
                continue RowLoop
            }
            row[x] = data + bias(x, y)
        }
    }

### Goto statements

A "goto" statement transfers control to the statement with the corresponding label within the same function.

GotoStmt = "goto" [Label](https://go.dev/ref/spec#Label) .

goto Error

Executing the "goto" statement must not cause any variables to come into [scope](https://go.dev/ref/spec#Declarations_and_scope) that were not already in scope at the point of the goto. For instance, this example:

    goto L  // BAD
    v := 3

L:

is erroneous because the jump to label `L` skips the creation of `v`.

A "goto" statement outside a [block](https://go.dev/ref/spec#Blocks) cannot jump to a label inside that block. For instance, this example:

if n%2 == 1 {
    goto L1
}
for n > 0 {
    f()
    n--
L1:
    f()
    n--
}

is erroneous because the label `L1` is inside the "for" statement's block but the `goto` is not.

### Fallthrough statements

A "fallthrough" statement transfers control to the first statement of the next case clause in an [expression "switch" statement](https://go.dev/ref/spec#Expression_switches). It may be used only as the final non-empty statement in such a clause.

FallthroughStmt = "fallthrough" .

### Defer statements

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a [return statement](https://go.dev/ref/spec#Return_statements), reached the end of its [function body](https://go.dev/ref/spec#Function_declarations), or because the corresponding goroutine is [panicking](https://go.dev/ref/spec#Handling_panics).  
“defer”语句调用一个函数，该函数的执行被推迟到周围函数返回的那一刻，要么是因为周围函数执行了 return 语句，到达了其函数体的末尾，要么是因为相应的 goroutine 正在恐慌。

DeferStmt = "defer" [Expression](https://go.dev/ref/spec#Expression) .

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](https://go.dev/ref/spec#Expression_statements).  
表达式必须是函数或方法调用；它不能被括号括起来。对于表达式语句，内置函数的调用受到限制。

Each time a "defer" statement executes, the function value and parameters to the call are [evaluated as usual](https://go.dev/ref/spec#Calls) and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. That is, if the surrounding function returns through an explicit [return statement](https://go.dev/ref/spec#Return_statements), deferred functions are executed *after* any result parameters are set by that return statement but *before* the function returns to its caller. If a deferred function value evaluates to `nil`, execution [panics](https://go.dev/ref/spec#Handling_panics) when the function is invoked, not when the "defer" statement is executed.  
每次执行“defer”语句时，调用的函数值和参数都会照常评估并重新保存，但不会调用实际函数。相反，延迟函数会在周围函数返回之前立即调用，调用顺序与延迟函数相反。也就是说，如果周围函数通过显式 return 语句返回，则延迟函数将在该 return 语句设置任何结果参数之后但在函数返回到其调用者之前执行。如果延迟函数值的计算结果为 `nil` ，则在调用该函数时（而不是执行“defer”语句时）执行会发生恐慌。

For instance, if the deferred function is a [function literal](https://go.dev/ref/spec#Function_literals) and the surrounding function has [named result parameters](https://go.dev/ref/spec#Function_types) that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. (See also the section on [handling panics](https://go.dev/ref/spec#Handling_panics).)  
例如，如果延迟函数是函数文字，并且周围函数具有在文字范围内的命名结果参数，则延迟函数可以在返回结果参数之前访问和修改结果参数。如果延迟函数有任何返回值，它们将在函数完成时被丢弃。 （另请参阅有关处理恐慌的部分。）

lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns
// prints 3 2 1 0 before surrounding function returns
for i := 0; i <= 3; i++ {
    defer fmt.Print(i)
}
// f returns 42
func f() (result int) {
    defer func() {
        // result is accessed after it was set to 6 by the return statement
        result *= 7
    }()
    return 6
}

## Built-in functions 内置功能

Built-in functions are [predeclared](https://go.dev/ref/spec#Predeclared_identifiers). They are called like any other function but some of them accept a type instead of an expression as the first argument.  
内置函数是预先声明的。它们的调用方式与任何其他函数一样，但其中一些函数接受类型而不是表达式作为第一个参数。

The built-in functions do not have standard Go types, so they can only appear in [call expressions](https://go.dev/ref/spec#Calls); they cannot be used as function values.  
内置函数没有标准的 Go 类型，因此只能出现在调用表达式中；它们不能用作函数值。

### Appending to and copying slices

追加和复制切片

The built-in functions `append` and `copy` assist in common slice operations. For both functions, the result is independent of whether the memory referenced by the arguments overlaps.  
内置函数 `append` 和 `copy` 有助于常见的切片操作。对于这两个函数，结果与参数引用的内存是否重叠无关。

The [variadic](https://go.dev/ref/spec#Function_types) function `append` appends zero or more values `x` to a slice `s` and returns the resulting slice of the same type as `s`. The [core type](https://go.dev/ref/spec#Core_types) of `s` must be a slice of type `[]E`. The values `x` are passed to a parameter of type `...E` and the respective [parameter passing rules](https://go.dev/ref/spec#Passing_arguments_to_..._parameters) apply. As a special case, if the core type of `s` is `[]byte`, `append` also accepts a second argument with core type [`bytestring`](https://go.dev/ref/spec#Core_types) followed by `...`. This form appends the bytes of the byte slice or string.  
可变参数函数 `append` 将零个或多个值 `x` 附加到切片 `s` 并返回与 `s` 类型相同的结果切片。 `s` 的核心类型必须是 `[]E` 类型的切片。值 `x` 将传递给类型 `...E` 的参数，并应用相应的参数传递规则。作为一种特殊情况，如果 `s` 的核心类型是 `[]byte` ，则 `append` 还接受第二个参数，其核心类型为 `bytestring` ，后跟 `...` 。此形式附加字节片或字符串的字节。

append(s S, x ...E) S  // core type of S is []E

If the capacity of `s` is not large enough to fit the additional values, `append` [allocates](https://go.dev/ref/spec#Allocation) a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, `append` re-uses the underlying array.  
如果 `s` 的容量不足以容纳附加值， `append` 会分配一个新的、足够大的底层数组，以容纳现有切片元素和附加值。否则， `append` 会重新使用底层数组。

s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 is []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 is []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 is []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 is []int{3, 5, 7, 2, 3, 5, 7, 0, 0}
var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                             t is []interface{}{42, 3.1415, "foo"}
var b []byte
b = append(b, "bar"...)            // append string contents      b is []byte{'b', 'a', 'r' }

The function `copy` copies slice elements from a source `src` to a destination `dst` and returns the number of elements copied. The [core types](https://go.dev/ref/spec#Core_types) of both arguments must be slices with [identical](https://go.dev/ref/spec#Type_identity) element type. The number of elements copied is the minimum of `len(src)` and `len(dst)`. As a special case, if the destination's core type is `[]byte`, `copy` also accepts a source argument with core type [`bytestring`](https://go.dev/ref/spec#Core_types). This form copies the bytes from the byte slice or string into the byte slice.

copy(dst, src []T) int
copy(dst []byte, src string) int

Examples:

var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s is []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s is []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b is []byte("Hello")

### Clear

The built-in function `clear` takes an argument of [map](https://go.dev/ref/spec#Map_types), [slice](https://go.dev/ref/spec#Slice_types), or [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) type, and deletes or zeroes out all elements [[Go 1.21](https://go.dev/ref/spec#Go_1.21)].

Call        Argument type     Result
clear(m)    map[K]T           deletes all entries, resulting in an
                              empty map (len(m) == 0)
clear(s)    []T               sets all elements up to the length of
                              `s` to the zero value of T
clear(t)    type parameter    see below

If the type of the argument to `clear` is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), all types in its type set must be maps or slices, and `clear` performs the operation corresponding to the actual type argument.

If the map or slice is `nil`, `clear` is a no-op.

### Close

For an argument `ch` with a [core type](https://go.dev/ref/spec#Core_types) that is a [channel](https://go.dev/ref/spec#Channel_types), the built-in function `close` records that no more values will be sent on the channel. It is an error if `ch` is a receive-only channel. Sending to or closing a closed channel causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics). Closing the nil channel also causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics). After calling `close`, and after any previously sent values have been received, receive operations will return the zero value for the channel's type without blocking. The multi-valued [receive operation](https://go.dev/ref/spec#Receive_operator) returns a received value along with an indication of whether the channel is closed.

### Manipulating complex numbers

Three functions assemble and disassemble complex numbers. The built-in function `complex` constructs a complex value from a floating-point real and imaginary part, while `real` and `imag` extract the real and imaginary parts of a complex value.

complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT

The type of the arguments and return value correspond. For `complex`, the two arguments must be of the same [floating-point type](https://go.dev/ref/spec#Numeric_types) and the return type is the [complex type](https://go.dev/ref/spec#Numeric_types) with the corresponding floating-point constituents: `complex64` for `float32` arguments, and `complex128` for `float64` arguments. If one of the arguments evaluates to an untyped constant, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to the type of the other argument. If both arguments evaluate to untyped constants, they must be non-complex numbers or their imaginary parts must be zero, and the return value of the function is an untyped complex constant.

For `real` and `imag`, the argument must be of complex type, and the return type is the corresponding floating-point type: `float32` for a `complex64` argument, and `float64` for a `complex128` argument. If the argument evaluates to an untyped constant, it must be a number, and the return value of the function is an untyped floating-point constant.

The `real` and `imag` functions together form the inverse of `complex`, so for a value `z` of a complex type `Z`, `z == Z(complex(real(z), imag(z)))`.

If the operands of these functions are all constants, the return value is a constant.

var a = complex(2, -2)             // complex128
const b = complex(1.0, -1.4)       // untyped complex constant 1 - 1.4i
x := float32(math.Cos(math.Pi/2))  // float32
var c64 = complex(5, -x)           // complex64
var s int = complex(1, 0)          // untyped complex constant 1 + 0i can be converted to int
_ = complex(1, 2<<s)               // illegal: 2 assumes floating-point type, cannot shift
var rl = real(c64)                 // float32
var im = imag(a)                   // float64
const c = imag(b)                  // untyped constant -1.4
_ = imag(3 << s)                   // illegal: 3 assumes complex type, cannot shift

Arguments of type parameter type are not permitted.

### Deletion of map elements

The built-in function `delete` removes the element with key `k` from a [map](https://go.dev/ref/spec#Map_types) `m`. The value `k` must be [assignable](https://go.dev/ref/spec#Assignability) to the key type of `m`.  
内置函数 `delete` 从映射 `m` 中删除带有键 `k` 的元素。值 `k` 必须可分配给 `m` 的键类型。

delete(m, k)  // remove element m[k] from map m

If the type of `m` is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), all types in that type set must be maps, and they must all have identical key types.  
如果 `m` 的类型是类型参数，则该类型集中的所有类型都必须是映射，并且它们都必须具有相同的键类型。

If the map `m` is `nil` or the element `m[k]` does not exist, `delete` is a no-op.  
如果映射 `m` 为 `nil` 或元素 `m[k]` 不存在，则 `delete` 为空操作。

### Length and capacity 长度和容量

The built-in functions `len` and `cap` take arguments of various types and return a result of type `int`. The implementation guarantees that the result always fits into an `int`.  
内置函数 `len` 和 `cap` 接受各种类型的参数并返回 `int` 类型的结果。该实现保证结果始终适合 `int` 。

Call      Argument type    Result
len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer
          type parameter   see below
cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
          type parameter   see below

If the argument type is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) `P`, the call `len(e)` (or `cap(e)` respectively) must be valid for each type in `P`'s type set. The result is the length (or capacity, respectively) of the argument whose type corresponds to the type argument with which `P` was [instantiated](https://go.dev/ref/spec#Instantiations).  
如果参数类型是类型参数 `P` ，则调用 `len(e)` （或分别为 `cap(e)` ）对于 `P` 中的每种类型都必须有效的类型集。结果是参数的长度（或容量），该参数的类型对应于实例化 `P` 的类型参数。

The capacity of a slice is the number of elements for which there is space allocated in the underlying array. At any time the following relationship holds:  
切片的容量是底层数组中为其分配空间的元素数量。在任何时候，以下关系都成立：

0 <= len(s) <= cap(s)

The length of a `nil` slice, map or channel is 0. The capacity of a `nil` slice or channel is 0.  
`nil` 切片、贴图或通道的长度为 0。 `nil` 切片或通道的容量为 0。

The expression `len(s)` is [constant](https://go.dev/ref/spec#Constants) if `s` is a string constant. The expressions `len(s)` and `cap(s)` are constants if the type of `s` is an array or pointer to an array and the expression `s` does not contain [channel receives](https://go.dev/ref/spec#Receive_operator) or (non-constant) [function calls](https://go.dev/ref/spec#Calls); in this case `s` is not evaluated. Otherwise, invocations of `len` and `cap` are not constant and `s` is evaluated.  
如果 `s` 是字符串常量，则表达式 `len(s)` 是常量。如果 `s` 的类型是数组或指向数组的指针，而表达式 `s` 不是常量，则表达式 `len(s)` 和 `cap(s)` 是常量包含通道接收或（非常量）函数调用；在这种情况下 `s` 不会被评估。否则， `len` 和 `cap` 的调用不是恒定的，并且 `s` 被评估。

const (
    c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
    c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
    c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
    c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
    c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
)
var z complex128

### Making slices, maps and channels

The built-in function `make` takes a type `T`, optionally followed by a type-specific list of expressions. The [core type](https://go.dev/ref/spec#Core_types) of `T` must be a slice, map or channel. It returns a value of type `T` (not `*T`). The memory is initialized as described in the section on [initial values](https://go.dev/ref/spec#The_zero_value).

Call             Core type    Result
make(T, n)       slice        slice of type T with length n and capacity n
make(T, n, m)    slice        slice of type T with length n and capacity m
make(T)          map          map of type T
make(T, n)       map          map of type T with initial space for approximately n elements
make(T)          channel      unbuffered channel of type T
make(T, n)       channel      buffered channel of type T, buffer size n

Each of the size arguments `n` and `m` must be of [integer type](https://go.dev/ref/spec#Numeric_types), have a [type set](https://go.dev/ref/spec#Interface_types) containing only integer types, or be an untyped [constant](https://go.dev/ref/spec#Constants). A constant size argument must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. If both `n` and `m` are provided and are constant, then `n` must be no larger than `m`. For slices and channels, if `n` is negative or larger than `m` at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.

s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements

Calling `make` with a map type and size hint `n` will create a map with initial space to hold `n` map elements. The precise behavior is implementation-dependent.

### Min and max

The built-in functions `min` and `max` compute the smallest—or largest, respectively—value of a fixed number of arguments of [ordered types](https://go.dev/ref/spec#Comparison_operators). There must be at least one argument [[Go 1.21](https://go.dev/ref/spec#Go_1.21)].

The same type rules as for [operators](https://go.dev/ref/spec#Operators) apply: for [ordered](https://go.dev/ref/spec#Comparison_operators) arguments `x` and `y`, `min(x, y)` is valid if `x + y` is valid, and the type of `min(x, y)` is the type of `x + y` (and similarly for `max`). If all arguments are constant, the result is constant.

var x, y int
m := min(x)                 // m == x
m := min(x, y)              // m is the smaller of x and y
m := max(x, y, 10)          // m is the larger of x and y but at least 10
c := max(1, 2.0, 10)        // c == 10.0 (floating-point kind)
f := max(0, float32(x))     // type of f is float32
var s []string
_ = min(s...)               // invalid: slice arguments are not permitted
t := max("", "foo", "bar")  // t == "foo" (string kind)

For numeric arguments, assuming all NaNs are equal, `min` and `max` are commutative and associative:

min(x, y)    == min(y, x)
min(x, y, z) == min(min(x, y), z) == min(x, min(y, z))

For floating-point arguments negative zero, NaN, and infinity the following rules apply:

   x        y    min(x, y)    max(x, y)
  -0.0    0.0         -0.0          0.0    // negative zero is smaller than (non-negative) zero
  -Inf      y         -Inf            y    // negative infinity is smaller than any other number
  +Inf      y            y         +Inf    // positive infinity is larger than any other number
   NaN      y          NaN          NaN    // if any argument is a NaN, the result is a NaN

For string arguments the result for `min` is the first argument with the smallest (or for `max`, largest) value, compared lexically byte-wise:

min(x, y)    == if x <= y then x else y
min(x, y, z) == min(min(x, y), z)

### Allocation

The built-in function `new` takes a type `T`, allocates storage for a [variable](https://go.dev/ref/spec#Variables) of that type at run time, and returns a value of type `*T` [pointing](https://go.dev/ref/spec#Pointer_types) to it. The variable is initialized as described in the section on [initial values](https://go.dev/ref/spec#The_zero_value).

new(T)

For instance

type S struct { a int; b float64 }
new(S)

allocates storage for a variable of type `S`, initializes it (`a=0`, `b=0.0`), and returns a value of type `*S` containing the address of the location.

### Handling panics

Two built-in functions, `panic` and `recover`, assist in reporting and handling [run-time panics](https://go.dev/ref/spec#Run_time_panics) and program-defined error conditions.

func panic(interface{})
func recover() interface{}

While executing a function `F`, an explicit call to `panic` or a [run-time panic](https://go.dev/ref/spec#Run_time_panics) terminates the execution of `F`. Any functions [deferred](https://go.dev/ref/spec#Defer_statements) by `F` are then executed as usual. Next, any deferred functions run by `F`'s caller are run, and so on up to any deferred by the top-level function in the executing goroutine. At that point, the program is terminated and the error condition is reported, including the value of the argument to `panic`. This termination sequence is called *panicking*.

panic(42)
panic("unreachable")
panic(Error("cannot parse"))

The `recover` function allows a program to manage behavior of a panicking goroutine. Suppose a function `G` defers a function `D` that calls `recover` and a panic occurs in a function on the same goroutine in which `G` is executing. When the running of deferred functions reaches `D`, the return value of `D`'s call to `recover` will be the value passed to the call of `panic`. If `D` returns normally, without starting a new `panic`, the panicking sequence stops. In that case, the state of functions called between `G` and the call to `panic` is discarded, and normal execution resumes. Any functions deferred by `G` before `D` are then run and `G`'s execution terminates by returning to its caller.  
`recover` 函数允许程序管理恐慌 goroutine 的行为。假设函数 `G` 推迟了调用 `recover` 的函数 `D` ，并且在 `G` 所在的同一 goroutine 上的函数中发生了恐慌。执行。当延迟函数运行到 `D` 时， `D` 调用 `recover` 的返回值将是传递给 `panic` 正常返回，而不启动新的 `panic` ，则恐慌序列将停止。在这种情况下， `G` 和对 `panic` 的调用之间调用的函数的状态将被丢弃，并恢复正常执行。然后，在 `D` 之前由 `G` 推迟的任何函数都会运行，并且 `G` 的执行将通过返回到其调用者来终止。

The return value of `recover` is `nil` when the goroutine is not panicking or `recover` was not called directly by a deferred function. Conversely, if a goroutine is panicking and `recover` was called directly by a deferred function, the return value of `recover` is guaranteed not to be `nil`. To ensure this, calling `panic` with a `nil` interface value (or an untyped `nil`) causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).  
当 goroutine 没有恐慌或者 `recover` 没有被延迟函数直接调用时， `recover` 的返回值为 `nil` 。相反，如果 goroutine 发生恐慌并且 `recover` 被延迟函数直接调用，则 `recover` 的返回值保证不是 `nil` 。为了确保这一点，使用 `nil` 接口值（或无类型的 `nil` ）调用 `panic` 会导致运行时恐慌。

The `protect` function in the example below invokes the function argument `g` and protects callers from run-time panics raised by `g`.  
下面示例中的 `protect` 函数调用函数参数 `g` 并保护调用者免受 `g` 引发的运行时恐慌。

func protect(g func()) {
    defer func() {
        log.Println("done")  // Println executes normally even if there is a panic
        if x := recover(); x != nil {
            log.Printf("run time panic: %v", x)
        }
    }()
    log.Println("start")
    g()
}

### Bootstrapping 自举

Current implementations provide several built-in functions useful during bootstrapping. These functions are documented for completeness but are not guaranteed to stay in the language. They do not return a result.  
当前的实现提供了几个在引导期间有用的内置函数。这些函数的记录是为了完整性，但不保证保留在该语言中。他们不返回结果。

Function   Behavior
print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end

Implementation restriction: `print` and `println` need not accept arbitrary argument types, but printing of boolean, numeric, and string [types](https://go.dev/ref/spec#Types) must be supported.  
实现限制： `print` 和 `println` 不需要接受任意参数类型，但必须支持布尔、数字和字符串类型的打印。

## Packages 套餐

Go programs are constructed by linking together *packages*. A package in turn is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package. Those elements may be [exported](https://go.dev/ref/spec#Exported_identifiers) and used in another package.  
Go 程序是通过将包链接在一起来构建的。包又由一个或多个源文件构造而成，这些源文件一起声明属于该包的常量、类型、变量和函数，并且可以在同一包的所有文件中访问它们。这些元素可以导出并在另一个包中使用。

### Source file organization

源文件组织

Each source file consists of a package clause defining the package to which it belongs, followed by a possibly empty set of import declarations that declare packages whose contents it wishes to use, followed by a possibly empty set of declarations of functions, types, variables, and constants.  
每个源文件都包含一个包子句，定义它所属的包，后跟一组可能为空的导入声明，声明其希望使用其内容的包，后跟一组可能为空的函数、类型、变量声明，和常数。

SourceFile       = [PackageClause](https://go.dev/ref/spec#PackageClause) ";" { [ImportDecl](https://go.dev/ref/spec#ImportDecl) ";" } { [TopLevelDecl](https://go.dev/ref/spec#TopLevelDecl) ";" } .

### Package clause 套餐条款

A package clause begins each source file and defines the package to which the file belongs.  
package 子句开始每个源文件并定义该文件所属的包。

PackageClause  = "package" [PackageName](https://go.dev/ref/spec#PackageName) .
PackageName    = [identifier](https://go.dev/ref/spec#identifier) .

The PackageName must not be the [blank identifier](https://go.dev/ref/spec#Blank_identifier).

package math

A set of files sharing the same PackageName form the implementation of a package. An implementation may require that all source files for a package inhabit the same directory.

### Import declarations

An import declaration states that the source file containing the declaration depends on functionality of the *imported* package ([§Program initialization and execution](https://go.dev/ref/spec#Program_initialization_and_execution)) and enables access to [exported](https://go.dev/ref/spec#Exported_identifiers) identifiers of that package. The import names an identifier (PackageName) to be used for access and an ImportPath that specifies the package to be imported.

ImportDecl       = "import" ( [ImportSpec](https://go.dev/ref/spec#ImportSpec) | "(" { [ImportSpec](https://go.dev/ref/spec#ImportSpec) ";" } ")" ) .
ImportSpec       = [ "." | [PackageName](https://go.dev/ref/spec#PackageName) ] [ImportPath](https://go.dev/ref/spec#ImportPath) .
ImportPath       = [string_lit](https://go.dev/ref/spec#string_lit) .

The PackageName is used in [qualified identifiers](https://go.dev/ref/spec#Qualified_identifiers) to access exported identifiers of the package within the importing source file. It is declared in the [file block](https://go.dev/ref/spec#Blocks). If the PackageName is omitted, it defaults to the identifier specified in the [package clause](https://go.dev/ref/spec#Package_clause) of the imported package. If an explicit period (`.`) appears instead of a name, all the package's exported identifiers declared in that package's [package block](https://go.dev/ref/spec#Blocks) will be declared in the importing source file's file block and must be accessed without a qualifier.

The interpretation of the ImportPath is implementation-dependent but it is typically a substring of the full file name of the compiled package and may be relative to a repository of installed packages.

Implementation restriction: A compiler may restrict ImportPaths to non-empty strings using only characters belonging to [Unicode's](https://www.unicode.org/versions/Unicode6.3.0/) L, M, N, P, and S general categories (the Graphic characters without spaces) and may also exclude the characters ``!"#$%&'()*,:;<=>?[\]^`{|}`` and the Unicode replacement character U+FFFD.

Consider a compiled a package containing the package clause `package math`, which exports function `Sin`, and installed the compiled package in the file identified by `"lib/math"`. This table illustrates how `Sin` is accessed in files that import the package after the various types of import declaration.

Import declaration          Local name of Sin
import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin

An import declaration declares a dependency relation between the importing and imported package. It is illegal for a package to import itself, directly or indirectly, or to directly import a package without referring to any of its exported identifiers. To import a package solely for its side-effects (initialization), use the [blank](https://go.dev/ref/spec#Blank_identifier) identifier as explicit package name:

import _ "lib/math"

### An example package

Here is a complete Go package that implements a concurrent prime sieve.

package main
import "fmt"
// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
    for i := 2; ; i++ {
        ch <- i  // Send 'i' to channel 'ch'.
    }
}
// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
    for i := range src {  // Loop over values received from 'src'.
        if i%prime != 0 {
            dst <- i  // Send 'i' to channel 'dst'.
        }
    }
}
// The prime sieve: Daisy-chain filter processes together.
func sieve() {
    ch := make(chan int)  // Create a new channel.
    go generate(ch)       // Start generate() as a subprocess.
    for {
        prime := <-ch
        fmt.Print(prime, "\n")
        ch1 := make(chan int)
        go filter(ch, ch1, prime)
        ch = ch1
    }
}
func main() {
    sieve()
}

## Program initialization and execution

程序初始化和执行

### The zero value 零值

When storage is allocated for a [variable](https://go.dev/ref/spec#Variables), either through a declaration or a call of `new`, or when a new value is created, either through a composite literal or a call of `make`, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the *zero value* for its type: `false` for booleans, `0` for numeric types, `""` for strings, and `nil` for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.  
当通过声明或调用 `new` 为变量分配存储空间时，或者通过复合文字或调用 `make` 创建新值时，以及没有提供显式初始化，变量或值被赋予默认值。此类变量或值的每个元素都设置为其类型的零值：对于布尔值， `false` ，对于数字类型， `0` ，对于字符串， `""` ，以及 `nil` 用于指针、函数、接口、切片、通道和映射。此初始化是递归完成的，因此，例如，如果未指定值，则结构数组的每个元素都将其字段归零。

These two simple declarations are equivalent:  
这两个简单的声明是等效的：

var i int
var i int = 0

After  后

type T struct { i int; f float64; next *T }
t := new(T)

the following holds:  以下成立：

t.i == 0
t.f == 0.0
t.next == nil

The same would also be true after  
之后也是如此

var t T

### Package initialization 包初始化

Within a package, package-level variable initialization proceeds stepwise, with each step selecting the variable earliest in *declaration order* which has no dependencies on uninitialized variables.  
在包内，包级变量初始化逐步进行，每一步都选择声明顺序中最早的变量，该变量不依赖于未初始化的变量。

More precisely, a package-level variable is considered *ready for initialization* if it is not yet initialized and either has no [initialization expression](https://go.dev/ref/spec#Variable_declarations) or its initialization expression has no *dependencies* on uninitialized variables. Initialization proceeds by repeatedly initializing the next package-level variable that is earliest in declaration order and ready for initialization, until there are no variables ready for initialization.  
更准确地说，如果包级变量尚未初始化并且没有初始化表达式或其初始化表达式不依赖于未初始化的变量，则认为该变量已准备好初始化。通过重复初始化声明顺序中最早并准备初始化的下一个包级变量来进行初始化，直到没有准备好初始化的变量为止。

If any variables are still uninitialized when this process ends, those variables are part of one or more initialization cycles, and the program is not valid.  
如果在此过程结束时任何变量仍未初始化，则这些变量是一个或多个初始化周期的一部分，并且程序无效。

Multiple variables on the left-hand side of a variable declaration initialized by single (multi-valued) expression on the right-hand side are initialized together: If any of the variables on the left-hand side is initialized, all those variables are initialized in the same step.  
由右侧单个（多值）表达式初始化的变量声明左侧的多个变量会一起初始化：如果左侧的任何变量被初始化，则所有这些变量都会被初始化在同一步骤中。

var x = a
var a, b = f() // a and b are initialized together, before x is initialized

For the purpose of package initialization, [blank](https://go.dev/ref/spec#Blank_identifier) variables are treated like any other variables in declarations.

The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler: Variables declared in the first file are declared before any of the variables declared in the second file, and so on. To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.

Dependency analysis does not rely on the actual values of the variables, only on lexical *references* to them in the source, analyzed transitively. For instance, if a variable `x`'s initialization expression refers to a function whose body refers to variable `y` then `x` depends on `y`. Specifically:

- A reference to a variable or function is an identifier denoting that variable or function.
- A reference to a method `m` is a [method value](https://go.dev/ref/spec#Method_values) or [method expression](https://go.dev/ref/spec#Method_expressions) of the form `t.m`, where the (static) type of `t` is not an interface type, and the method `m` is in the [method set](https://go.dev/ref/spec#Method_sets) of `t`. It is immaterial whether the resulting function value `t.m` is invoked.
- A variable, function, or method `x` depends on a variable `y` if `x`'s initialization expression or body (for functions and methods) contains a reference to `y` or to a function or method that depends on `y`.

For example, given the declarations

var (
    a = c + b  // == 9
    b = f()    // == 4
    c = f()    // == 5
    d = 3      // == 5 after initialization has finished
)
func f() int {
    d++
    return d
}

the initialization order is `d`, `b`, `c`, `a`. Note that the order of subexpressions in initialization expressions is irrelevant: `a = c + b` and `a = b + c` result in the same initialization order in this example.

Dependency analysis is performed per package; only references referring to variables, functions, and (non-interface) methods declared in the current package are considered. If other, hidden, data dependencies exists between variables, the initialization order between those variables is unspecified.

For instance, given the declarations

var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b
var _ = sideEffect()  // unrelated to x, a, or b
var a = b
var b = 42
type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }

the variable `a` will be initialized after `b` but whether `x` is initialized before `b`, between `b` and `a`, or after `a`, and thus also the moment at which `sideEffect()` is called (before or after `x` is initialized) is not specified.

Variables may also be initialized using functions named `init` declared in the package block, with no arguments and no result parameters.

func init() { … }

Multiple such functions may be defined per package, even within a single source file. In the package block, the `init` identifier can be used only to declare `init` functions, yet the identifier itself is not [declared](https://go.dev/ref/spec#Declarations_and_scope). Thus `init` functions cannot be referred to from anywhere in a program.

The entire package is initialized by assigning initial values to all its package-level variables followed by calling all `init` functions in the order they appear in the source, possibly in multiple files, as presented to the compiler.

### Program initialization

The packages of a complete program are initialized stepwise, one package at a time. If a package has imports, the imported packages are initialized before initializing the package itself. If multiple packages import a package, the imported package will be initialized only once. The importing of packages, by construction, guarantees that there can be no cyclic initialization dependencies. More precisely:

Given the list of all packages, sorted by import path, in each step the first uninitialized package in the list for which all imported packages (if any) are already initialized is [initialized](https://go.dev/ref/spec#Package_initialization). This step is repeated until all packages are initialized.

Package initialization—variable initialization and the invocation of `init` functions—happens in a single goroutine, sequentially, one package at a time. An `init` function may launch other goroutines, which can run concurrently with the initialization code. However, initialization always sequences the `init` functions: it will not invoke the next one until the previous one has returned.

### Program execution

A complete program is created by linking a single, unimported package called the *main package* with all the packages it imports, transitively. The main package must have package name `main` and declare a function `main` that takes no arguments and returns no value.

func main() { … }

Program execution begins by [initializing the program](https://go.dev/ref/spec#Program_initialization) and then invoking the function `main` in package `main`. When that function invocation returns, the program exits. It does not wait for other (non-`main`) goroutines to complete.

## Errors

The predeclared type `error` is defined as

type error interface {
    Error() string
}

It is the conventional interface for representing an error condition, with the nil value representing no error. For instance, a function to read data from a file might be defined:

func Read(f *File, b []byte) (n int, err error)

## Run-time panics 运行时恐慌

Execution errors such as attempting to index an array out of bounds trigger a *run-time panic* equivalent to a call of the built-in function [`panic`](https://go.dev/ref/spec#Handling_panics) with a value of the implementation-defined interface type `runtime.Error`. That type satisfies the predeclared interface type [`error`](https://go.dev/ref/spec#Errors). The exact error values that represent distinct run-time error conditions are unspecified.  
执行错误（例如尝试对数组进行索引越界）会触发运行时恐慌，相当于使用实现定义的接口类型 `runtime.Error` 的值调用内置函数 `panic` b1> .该类型满足预先声明的接口类型 `error` 。表示不同运行时错误条件的确切错误值未指定。

package runtime
type Error interface {
    error
    // and perhaps other methods
}

## System considerations 系统注意事项

### Package `unsafe` 套餐 `unsafe`

The built-in package `unsafe`, known to the compiler and accessible through the [import path](https://go.dev/ref/spec#Import_declarations) `"unsafe"`, provides facilities for low-level programming including operations that violate the type system. A package using `unsafe` must be vetted manually for type safety and may not be portable. The package provides the following interface:  
编译器已知并可通过导入路径 `"unsafe"` 访问的内置包 `unsafe` 提供了用于低级编程的设施，包括违反类型系统的操作。使用 `unsafe` 的包必须经过手动审查以确保类型安全，并且可能不可移植。该包提供了以下接口：

package unsafe
type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
type Pointer *ArbitraryType
func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr
type IntegerType int  // shorthand for an integer type; it is not a real type
func Add(ptr Pointer, len IntegerType) Pointer
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
func SliceData(slice []ArbitraryType) *ArbitraryType
func String(ptr *byte, len IntegerType) string
func StringData(str string) *byte

A `Pointer` is a [pointer type](https://go.dev/ref/spec#Pointer_types) but a `Pointer` value may not be [dereferenced](https://go.dev/ref/spec#Address_operators). Any pointer or value of [core type](https://go.dev/ref/spec#Core_types) `uintptr` can be [converted](https://go.dev/ref/spec#Conversions) to a type of core type `Pointer` and vice versa. The effect of converting between `Pointer` and `uintptr` is implementation-defined.  
`Pointer` 是指针类型，但 `Pointer` 值可能无法取消引用。核心类型 `uintptr` 的任何指针或值都可以转换为核心类型 `Pointer` 的类型，反之亦然。 `Pointer` 和 `uintptr` 之间转换的效果是实现定义的。

var f float64
bits = *(*uint64)(unsafe.Pointer(&f))
type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&f))
func f[P ~*B, B any](p P) uintptr {
    return uintptr(unsafe.Pointer(p))
}
var p ptr = nil

The functions `Alignof` and `Sizeof` take an expression `x` of any type and return the alignment or size, respectively, of a hypothetical variable `v` as if `v` was declared via `var v = x`.  
函数 `Alignof` 和 `Sizeof` 采用任意类型的表达式 `x` 并分别返回假设变量 `v` 的对齐方式或大小就好像 `v` 是通过 `var v = x` 声明的。

The function `Offsetof` takes a (possibly parenthesized) [selector](https://go.dev/ref/spec#Selectors) `s.f`, denoting a field `f` of the struct denoted by `s` or `*s`, and returns the field offset in bytes relative to the struct's address. If `f` is an [embedded field](https://go.dev/ref/spec#Struct_types), it must be reachable without pointer indirections through fields of the struct. For a struct `s` with field `f`:

uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))

Computer architectures may require memory addresses to be *aligned*; that is, for addresses of a variable to be a multiple of a factor, the variable's type's *alignment*. The function `Alignof` takes an expression denoting a variable of any type and returns the alignment of the (type of the) variable in bytes. For a variable `x`:

uintptr(unsafe.Pointer(&x)) % unsafe.Alignof(x) == 0

A (variable of) type `T` has *variable size* if `T` is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), or if it is an array or struct type containing elements or fields of variable size. Otherwise the size is *constant*. Calls to `Alignof`, `Offsetof`, and `Sizeof` are compile-time [constant expressions](https://go.dev/ref/spec#Constant_expressions) of type `uintptr` if their arguments (or the struct `s` in the selector expression `s.f` for `Offsetof`) are types of constant size.

The function `Add` adds `len` to `ptr` and returns the updated pointer `unsafe.Pointer(uintptr(ptr) + uintptr(len))` [[Go 1.17](https://go.dev/ref/spec#Go_1.17)]. The `len` argument must be of [integer type](https://go.dev/ref/spec#Numeric_types) or an untyped [constant](https://go.dev/ref/spec#Constants). A constant `len` argument must be [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. The rules for [valid uses](https://go.dev/pkg/unsafe#Pointer) of `Pointer` still apply.

The function `Slice` returns a slice whose underlying array starts at `ptr` and whose length and capacity are `len`. `Slice(ptr, len)` is equivalent to

(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]

except that, as a special case, if `ptr` is `nil` and `len` is zero, `Slice` returns `nil` [[Go 1.17](https://go.dev/ref/spec#Go_1.17)].  
但作为特殊情况，如果 `ptr` 为 `nil` 并且 `len` 为零，则 `Slice` 返回 `nil` [转到1.17]。

The `len` argument must be of [integer type](https://go.dev/ref/spec#Numeric_types) or an untyped [constant](https://go.dev/ref/spec#Constants). A constant `len` argument must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. At run time, if `len` is negative, or if `ptr` is `nil` and `len` is not zero, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs [[Go 1.17](https://go.dev/ref/spec#Go_1.17)].  
`len` 参数必须是整数类型或无类型常量。常量 `len` 参数必须是非负数并且可以由 `int` 类型的值表示；如果它是无类型常量，则指定类型 `int` 。在运行时，如果 `len` 为负数，或者 `ptr` 为 `nil` 并且 `len` 不为零，则会发生运行时恐慌[转到1.17]。

The function `SliceData` returns a pointer to the underlying array of the `slice` argument. If the slice's capacity `cap(slice)` is not zero, that pointer is `&slice[:1][0]`. If `slice` is `nil`, the result is `nil`. Otherwise it is a non-`nil` pointer to an unspecified memory address [[Go 1.20](https://go.dev/ref/spec#Go_1.20)].  
函数 `SliceData` 返回一个指向 `slice` 参数的底层数组的指针。如果切片的容量 `cap(slice)` 不为零，则该指针为 `&slice[:1][0]` 。如果 `slice` 为 `nil` ，则结果为 `nil` 。否则它是一个非 `nil` 指针，指向未指定的内存地址 [Go 1.20]。

The function `String` returns a `string` value whose underlying bytes start at `ptr` and whose length is `len`. The same requirements apply to the `ptr` and `len` argument as in the function `Slice`. If `len` is zero, the result is the empty string `""`. Since Go strings are immutable, the bytes passed to `String` must not be modified afterwards. [[Go 1.20](https://go.dev/ref/spec#Go_1.20)]  
函数 `String` 返回一个 `string` 值，其底层字节从 `ptr` 开始，长度为 `len` 。与函数 `Slice` 中一样，同样的要求也适用于 `ptr` 和 `len` 参数。如果 `len` 为零，则结果为空字符串 `""` 。由于 Go 字符串是不可变的，因此传递给 `String` 的字节之后不得修改。 [转到1.20]

The function `StringData` returns a pointer to the underlying bytes of the `str` argument. For an empty string the return value is unspecified, and may be `nil`. Since Go strings are immutable, the bytes returned by `StringData` must not be modified [[Go 1.20](https://go.dev/ref/spec#Go_1.20)].  
函数 `StringData` 返回指向 `str` 参数的底层字节的指针。对于空字符串，返回值未指定，可能是 `nil` 。由于 Go 字符串是不可变的，因此 `StringData` 返回的字节不得修改 [Go 1.20]。

### Size and alignment guarantees

尺寸和对齐保证

For the [numeric types](https://go.dev/ref/spec#Numeric_types), the following sizes are guaranteed:  
对于数字类型，保证以下大小：

type                                 size in bytes
byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16

The following minimal alignment properties are guaranteed:  
保证以下最小对齐属性：

1. For a variable `x` of any type: `unsafe.Alignof(x)` is at least 1.  
   对于任何类型的变量 `x` ： `unsafe.Alignof(x)` 至少为 1。
2. For a variable `x` of struct type: `unsafe.Alignof(x)` is the largest of all the values `unsafe.Alignof(x.f)` for each field `f` of `x`, but at least 1.  
   对于结构体类型的变量 `x` ： `unsafe.Alignof(x)` 是 `x` 的所有值 `unsafe.Alignof(x.f)` 中最大的一个/b4> ，但至少 1。
3. For a variable `x` of array type: `unsafe.Alignof(x)` is the same as the alignment of a variable of the array's element type.  
   对于数组类型的变量 `x` ： `unsafe.Alignof(x)` 与数组元素类型的变量的对齐方式相同。

A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.  
如果结构体或数组类型不包含大小大于零的字段（或元素），则其大小为零。两个不同的零大小变量在内存中可能具有相同的地址。

## Appendix 附录

### Language versions 语言版本

The [Go 1 compatibility guarantee](https://go.dev/doc/go1compat) ensures that programs written to the Go 1 specification will continue to compile and run correctly, unchanged, over the lifetime of that specification. More generally, as adjustments are made and features added to the language, the compatibility guarantee ensures that a Go program that works with a specific Go language version will continue to work with any subsequent version.  
Go 1 兼容性保证确保按照 Go 1 规范编写的程序将在该规范的生命周期内继续正确编译和运行，不会发生任何变化。更一般地说，随着语言的调整和功能的添加，兼容性保证可确保适用于特定 Go 语言版本的 Go 程序将继续适用于任何后续版本。

For instance, the ability to use the prefix `0b` for binary integer literals was introduced with Go 1.13, indicated by [[Go 1.13](https://go.dev/ref/spec#Go_1.13)] in the section on [integer literals](https://go.dev/ref/spec#Integer_literals). Source code containing an integer literal such as `0b1011` will be rejected if the implied or required language version used by the compiler is older than Go 1.13.  
例如，Go 1.13 引入了对二进制整数文字使用前缀 `0b` 的功能，如整数文字部分中的 [ Go 1.13] 所示。如果编译器使用的隐含或必需的语言版本早于 Go 1.13，则包含整数文字（例如 `0b1011` ）的源代码将被拒绝。

The following table describes the minimum language version required for features introduced after Go 1.

#### Go 1.9

- An [alias declaration](https://go.dev/ref/spec#Alias_declarations) may be used to declare an alias name for a type.

#### Go 1.13

- [Integer literals](https://go.dev/ref/spec#Integer_literals) may use the prefixes `0b`, `0B`, `0o`, and `0O` for binary, and octal literals, respectively.
- Hexadecimal [floating-point literals](https://go.dev/ref/spec#Floating-point_literals) may be written using the prefixes `0x` and `0X`.
- The [imaginary suffix](https://go.dev/ref/spec#Imaginary_literals) `i` may be used with any (binary, decimal, hexadecimal) integer or floating-point literal, not just decimal literals.
- The digits of any number literal may be [separated](https://go.dev/ref/spec#Integer_literals) (grouped) using underscores `_`.
- The shift count in a [shift operation](https://go.dev/ref/spec#Operators) may be a signed integer type.

#### Go 1.14

- Emdedding a method more than once through different [embedded interfaces](https://go.dev/ref/spec#Embedded_interfaces) is not an error.

#### Go 1.17

- A slice may be [converted](https://go.dev/ref/spec#Conversions) to an array pointer if the slice and array element types match, and the array is not longer than the slice.
- The built-in [package `unsafe`](https://go.dev/ref/spec#Package_unsafe) includes the new functions `Add` and `Slice`.

#### Go 1.18

The 1.18 release adds polymorphic functions and types ("generics") to the language. Specifically:

- The set of [operators and punctuation](https://go.dev/ref/spec#Operators_and_punctuation) includes the new token `~`.
- Function and type declarations may declare [type parameters](https://go.dev/ref/spec#Type_parameter_declarations).
- Interface types may [embed arbitrary types](https://go.dev/ref/spec#General_interfaces) (not just type names of interfaces) as well as union and `~T` type elements.
- The set of [predeclared](https://go.dev/ref/spec#Predeclared_identifiers) types includes the new types `any` and `comparable`.

#### Go 1.20

- A slice may be [converted](https://go.dev/ref/spec#Conversions) to an array if the slice and array element types match and the array is not longer than the slice.
- The built-in [package `unsafe`](https://go.dev/ref/spec#Package_unsafe) includes the new functions `SliceData`, `String`, and `StringData`.
- [Comparable types](https://go.dev/ref/spec#Comparison_operators) (such as ordinary interfaces) may satisfy `comparable` constraints, even if the type arguments are not strictly comparable.

#### Go 1.21

- The set of [predeclared](https://go.dev/ref/spec#Predeclared_identifiers) functions includes the new functions `min`, `max`, and `clear`.
- [Type inference](https://go.dev/ref/spec#Type_inference) uses the types of interface methods for inference. It also infers type arguments for generic functions assigned to variables or passed as arguments to other (possibly generic) functions.

#### Go 1.22

- In a ["for" statement](https://go.dev/ref/spec#For_statements), each iteration has its own set of iteration variables rather than sharing the same variables in each iteration.
- A "for" statement with ["range" clause](https://go.dev/ref/spec#For_range) may iterate over integer values from zero to an upper limit.

### Type unification rules

The type unification rules describe if and how two types unify. The precise details are relevant for Go implementations, affect the specifics of error messages (such as whether a compiler reports a type inference or other error), and may explain why type inference fails in unusual code situations. But by and large these rules can be ignored when writing Go code: type inference is designed to mostly "work as expected", and the unification rules are fine-tuned accordingly.

Type unification is controlled by a *matching mode*, which may be *exact* or *loose*. As unification recursively descends a composite type structure, the matching mode used for elements of the type, the *element matching mode*, remains the same as the matching mode except when two types are unified for [assignability](https://go.dev/ref/spec#Assignability) (`≡A`): in this case, the matching mode is *loose* at the top level but then changes to *exact* for element types, reflecting the fact that types don't have to be identical to be assignable.

Two types that are not bound type parameters unify exactly if any of following conditions is true:

- Both types are [identical](https://go.dev/ref/spec#Type_identity).
- Both types have identical structure and their element types unify exactly.
- Exactly one type is an [unbound](https://go.dev/ref/spec#Type_inference) type parameter with a [core type](https://go.dev/ref/spec#Core_types), and that core type unifies with the other type per the unification rules for `≡A` (loose unification at the top level and exact unification for element types).

If both types are bound type parameters, they unify per the given matching modes if:  
如果两种类型都是绑定类型参数，则它们根据给定的匹配模式进行统一，如果：

- Both type parameters are identical.  
  两个类型参数是相同的。
- At most one of the type parameters has a known type argument. In this case, the type parameters are *joined*: they both stand for the same type argument. If neither type parameter has a known type argument yet, a future type argument inferred for one the type parameters is simultaneously inferred for both of them.  
  最多一个类型形参具有已知的类型实参。在这种情况下，类型参数被连接：它们都代表相同的类型参数。如果两个类型形参都没有已知的类型实参，则同时为这两个类型形参推断出为其中一个类型形参推断的未来类型实参。
- Both type parameters have a known type argument and the type arguments unify per the given matching modes.  
  两个类型参数都有一个已知的类型参数，并且类型参数根据给定的匹配模式进行统一。

A single bound type parameter `P` and another type `T` unify per the given matching modes if:  
单个绑定类型参数 `P` 和另一个类型 `T` 根据给定的匹配模式进行统一，如果：

- `P` doesn't have a known type argument. In this case, `T` is inferred as the type argument for `P`.  
  `P` 没有已知的类型参数。在本例中， `T` 被推断为 `P` 的类型参数。
- `P` does have a known type argument `A`, `A` and `T` unify per the given matching modes, and one of the following conditions is true:  
  `P` 确实有一个已知的类型参数 `A` 、 `A` 和 `T` 根据给定的匹配模式进行统一，并且以下条件之一是真的：
  - Both `A` and `T` are interface types: In this case, if both `A` and `T` are also [defined](https://go.dev/ref/spec#Type_definitions) types, they must be [identical](https://go.dev/ref/spec#Type_identity). Otherwise, if neither of them is a defined type, they must have the same number of methods (unification of `A` and `T` already established that the methods match).  
    `A` 和 `T` 都是接口类型：在这种情况下，如果 `A` 和 `T` 也是定义类型，则它们必须相同。否则，如果它们都不是定义的类型，则它们必须具有相同数量的方法（ `A` 和 `T` 的统一已经确定方法匹配）。
  - Neither `A` nor `T` are interface types: In this case, if `T` is a defined type, `T` replaces `A` as the inferred type argument for `P`.  
    `A` 和 `T` 都不是接口类型：在这种情况下，如果 `T` 是定义的类型，则 `T` 替换 `A` 的推断类型参数。

Finally, two types that are not bound type parameters unify loosely (and per the element matching mode) if:  
最后，如果满足以下条件，则两种未绑定类型参数的类型会松散地统一（并且根据元素匹配模式）：

- Both types unify exactly.  
  两种类型完全统一。
- One type is a [defined type](https://go.dev/ref/spec#Type_definitions), the other type is a type literal, but not an interface, and their underlying types unify per the element matching mode.  
  一种类型是定义类型，另一种类型是类型文字，但不是接口，并且它们的底层类型根据元素匹配模式进行统一。
- Both types are interfaces (but not type parameters) with identical [type terms](https://go.dev/ref/spec#Interface_types), both or neither embed the predeclared type [comparable](https://go.dev/ref/spec#Predeclared_identifiers), corresponding method types unify exactly, and the method set of one of the interfaces is a subset of the method set of the other interface.  
  两种类型都是具有相同类型术语的接口（但不是类型参数），都嵌入或都不嵌入预声明的可比较类型，相应的方法类型完全统一，并且其中一个接口的方法集是另一个接口的方法集的子集。
- Only one type is an interface (but not a type parameter), corresponding methods of the two types unify per the element matching mode, and the method set of the interface is a subset of the method set of the other type.  
  只有一种类型是接口（但不是类型参数），两种类型对应的方法按照元素匹配方式统一，并且接口的方法集是另一种类型的方法集的子集。
- Both types have the same structure and their element types unify per the element matching mode.  
  两种类型具有相同的结构，并且它们的元素类型根据元素匹配模式进行统一。

[Why Go](https://go.dev/solutions/)[Use Cases](https://go.dev/solutions/use-cases)[Case Studies](https://go.dev/solutions/case-studies)

[Get Started](https://go.dev/learn/)[Playground](https://go.dev/play)[Tour](https://go.dev/tour/)[Stack Overflow](https://stackoverflow.com/questions/tagged/go?tab=Newest)[Help](https://go.dev/help/)

[Packages](https://pkg.go.dev/)[Standar](https://go.dev/pkg/)
