## Documentation

* https://pkg.go.dev/fmt@go1.22.4

### Overview

- Printing
  打印

- Explicit argument indexes
  显式参数索引

- Format errors
  格式错误

- Scanning
  扫描

Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
fmt 包使用类似于 C 语言的 printf 和 scanf 的函数来实现格式化 I/O。 “动词”格式源自 C 语言，但更简单。

#### Printing

The verbs:
动词：

###### General
一般的：

```
%v    the value in a default format
      when printing structs, the plus flag (%+v) adds field names

%#v   a Go-syntax representation of the value

%T    a Go-syntax representation of the type of the value

%%    a literal percent sign; consumes no value
```

###### Boolean
布尔值：

```
%t    the word true or false
```

###### Integer
整数：

```
%b    base 2

%c    the character represented by the corresponding Unicode code point
      相应 Unicode 代码点所表示的字符

%d    base 10

%o    base 8

%O    base 8 with 0o prefix

%q    a single-quoted character literal safely escaped with Go syntax.
      使用 Go 语法可以安全地转义单引号字符文字。

%x    base 16, with lower-case letters for a-f

%X    base 16, with upper-case letters for A-F

%U    Unicode format: U+1234; same as "U+%04X"
```

###### Floating-point and complex constituents:
浮点和复数成分：

```
%b    decimalless scientific notation with exponent a power of two,
      in the manner of strconv.FormatFloat with the 'b' format,
      e.g. -123456p-78
      无小数的科学计数法，指数为 2 的幂
      类似于 strconv.FormatFloat 的 'b' 格式

%e    scientific notation, e.g. -1.234456e+78
      科学计数法

%E    scientific notation, e.g. -1.234456E+78
      科学计数法

%f    decimal point but no exponent, e.g. 123.456
      有小数点但没有指数

%F    synonym for %f
      与 %f 同义

%g    %e for large exponents, %f otherwise. Precision is discussed below.
      %e 表示大指数，否则为 %f

%G    %E for large exponents, %F otherwise
      %E 表示大指数，否则为 %F

%x    hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
       十六进制计数法（指数为 2 的十进制幂）

%X    upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
      大写十六进制计数法
```

###### String and slice of bytes (treated equivalently with these verbs):
字符串和字节切片（与这些动词等效处理）：

```
%s    the uninterpreted bytes of the string or slice
      字符串或切片的未解释字节

%q    a double-quoted string safely escaped with Go syntax
      使用 Go 语法安全转义的双引号字符串

%x    base 16, lower-case, two characters per byte
      16 进制，小写，每个字节两个字符

%X    base 16, upper-case, two characters per byte
      16 进制，大写，每个字节两个字符
```

###### Slice

```
%p    address of 0th element in base 16 notation, with leading 0x
      以 16 进制表示的第 0 个元素的地址，以 0x 开头

```

###### Pointer
指针：

```
%p    base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.
      以 16 为基数的表示法，以 0x 为前导
%b、%d、%o、%x 和 %X 动词也适用于指针，
将值格式化为整数。
```

The default format for %v is:
%v 的默认格式是：

```
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

For compound objects, the elements are printed using these rules, recursively, laid out like this:
对于复合对象，使用这些规则递归地打印元素，布局如下：

```
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]
```

Width is specified by an optional decimal number immediately preceding the verb. If absent, the width is whatever is necessary to represent the value. Precision is specified after the (optional) width by a period followed by a decimal number. If no period is present, a default precision is used. A period with no following number specifies a precision of zero. Examples:
宽度由紧接动词前面的可选十进制数指定。如果不存在，则宽度是表示该值所需的任何宽度。精度在（可选）宽度之后指定一个句点，后跟一个十进制数。如果不存在句点，则使用默认精度。没有后续数字的句点指定精度为零。例子：

```
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
```

Width and precision are measured in units of Unicode code points, that is, runes. (This differs from C's printf where the units are always measured in bytes.) Either or both of the flags may be replaced with the character '*', causing their values to be obtained from the next operand (preceding the one to format), which must be of type int.
宽度和精度以 Unicode 代码点（即符文）为单位进行测量。 （这与 C 的 printf 不同，后者的单位始终以字节为单位。）其中一个或两个标志都可以替换为字符“*”，从而导致从下一个操作数（在要格式化的操作数之前）获取它们的值，它必须是 int 类型。

For most values, width is the minimum number of runes to output, padding the formatted form with spaces if necessary.
对于大多数值，宽度是要输出的最小符文数量，如有必要，请用空格填充格式化形式。

For strings, byte slices and byte arrays, however, precision limits the length of the input to be formatted (not the size of the output), truncating if necessary. Normally it is measured in runes, but for these types when formatted with the %x or %X format it is measured in bytes.
然而，对于字符串、字节切片和字节数组，精度限制了要格式化的输入的长度（而不是输出的大小），并在必要时截断。通常，它以符文为单位进行测量，但对于这些类型，当使用 %x 或 %X 格式进行格式化时，它以字节为单位进行测量。

For floating-point values, width sets the minimum width of the field and precision sets the number of places after the decimal, if appropriate, except that for %g/%G precision sets the maximum number of significant digits (trailing zeros are removed). For example, given 12.345 the format %6.3f prints 12.345 while %.3g prints 12.3. The default precision for %e, %f and %#g is 6; for %g it is the smallest number of digits necessary to identify the value uniquely.
对于浮点值，宽度设置字段的最小宽度，精度设置小数点后的位数（如果适用），但对于 %g/%G 精度设置最大有效位数（删除尾随零） 。例如，给定 12.345，格式 %6.3f 打印 12.345，而 %.3g 打印 12.3。 %e、%f 和 %#g 的默认精度为 6；对于 %g，它是唯一标识该值所需的最小位数。

For complex numbers, the width and precision apply to the two components independently and the result is parenthesized, so %f applied to 1.2+3.4i produces (1.200000+3.400000i).
对于复数，宽度和精度独立应用于两个分量，并且结果带括号，因此应用于 1.2+3.4i 的 %f 会生成 (1.200000+3.400000i)。

When formatting a single integer code point or a rune string (type []rune) with %q, invalid Unicode code points are changed to the Unicode replacement character, U+FFFD, as in strconv.QuoteRune.
当使用 %q 格式化单个整数代码点或符文字符串（类型 []rune）时，无效的 Unicode 代码点将更改为 Unicode 替换字符 U+FFFD，如 strconv.QuoteRune 中所示。

###### Other flags
其他标志：

```
'+'    always print a sign for numeric values;
       guarantee ASCII-only output for %q (%+q)
       始终为数值打印一个符号；
       保证 %q (%+q) 仅输出 ASCII 格式

'-'    pad with spaces on the right rather than the left (left-justify the field)
       在右侧而不是左侧填充空格（左对齐字段）

'#'    alternate format:
       add leading 0b for binary (%#b),
       0 for octal (%#o),
       0x or 0X for hex (%#x or %#X);
       suppress 0x for %p (%#p);
       for %q, print a raw (backquoted) string if strconv.CanBackquote returns true;
       always print a decimal point for %e, %E, %f, %F, %g and %G;
       do not remove trailing zeros for %g and %G;
       write e.g. U+0078 'x' if the character is printable for %U (%#U).
       替代格式：
       为二进制 (%#b) 添加前导 0b，
       为八进制 (%#o) 添加 0，
       为十六进制 (%#x 或 %#X) 添加 0x 或 0X；
       为 %p (%#p) 隐藏 0x；
       对于 %q，如果 strconv.CanBackquote 返回 true，则打印原始（反引号）字符串；
       始终为 %e、%E、%f、%F、%g 和 %G 打印小数点；
       不要删除 %g 和 %G 的尾随零；
       如果字符可打印为 %U (%#U)，则写入例如 U+0078 'x'。

' '    (space) leave a space for elided sign in numbers (% d);
       put spaces between bytes printing strings or slices in hex (% x, % X)
       (空格) 在数字中为省略的符号留出空格 (% d);
       在打印十六进制字符串或切片的字节之间留出空格 (% x, % X)

'0'    pad with leading zeros rather than spaces;
       for numbers, this moves the padding after the sign;
       ignored for strings, byte slices and byte arrays
       用前导零填充，而不是空格；
       对于数字，这会将填充移动到符号后；
       对于字符串、字节切片和字节数组，忽略
```

Flags are ignored by verbs that do not expect them. For example there is no alternate decimal format, so %#d and %d behave identically.
不需要标记的动词会忽略标记。例如，没有替代的十进制格式，因此 %#d 和 %d 的行为相同。

For each Printf-like function, there is also a Print function that takes no format and is equivalent to saying %v for every operand. Another variant Println inserts blanks between operands and appends a newline.
对于每个类似 Printf 的函数，还有一个不带格式的 Print 函数，相当于对每个操作数说 %v。 Println 的另一个变体在操作数之间插入空格并附加换行符。

Regardless of the verb, if an operand is an interface value, the internal concrete value is used, not the interface itself. Thus:
无论动词如何，如果操作数是接口值，则使用内部具体值，而不是接口本身。因此：

```go
var i interface{} = 23
fmt.Printf("%v\n", i)
```

will print 23.

Except when printed using the verbs %T and %p, special formatting considerations apply for operands that implement certain interfaces. In order of application:
除非使用动词 %T 和 %p 打印，否则特殊格式注意事项适用于实现某些接口的操作数。按申请顺序：

1. If the operand is a reflect.Value, the operand is replaced by the concrete value that it holds, and printing continues with the next rule.
   如果操作数是reflect.Value，则操作数将被它所保存的具体值替换，并继续打印下一条规则。

2. If an operand implements the Formatter interface, it will be invoked. In this case the interpretation of verbs and flags is controlled by that implementation.
   如果一个操作数实现了Formatter接口，它将被调用。在这种情况下，动词和标志的解释由该实现控制。

3. If the %v verb is used with the # flag (%#v) and the operand implements the GoStringer interface, that will be invoked.
   如果 %v 动词与 # 标志 (%#v) 一起使用并且操作数实现 GoStringer 接口，则将调用该接口。

If the format (which is implicitly %v for Println etc.) is valid for a string (%s %q %x %X), or is %v but not %#v, the following two rules apply:
如果格式（对于 Println 等隐式为 %v）对于字符串 (%s %q %x %X) 有效，或者是 %v 但不是 %#v，则适用以下两个规则：

4. If an operand implements the error interface, the Error method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).
   如果操作数实现了 error 接口，则将调用 Error 方法将对象转换为字符串，然后根据动词（如果有）要求对其进行格式化。

5. If an operand implements method String() string, that method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).
   如果操作数实现 String() string 方法，则将调用该方法将对象转换为字符串，然后根据动词（如果有）要求对其进行格式化。

For compound operands such as slices and structs, the format applies to the elements of each operand, recursively, not to the operand as a whole. Thus %q will quote each element of a slice of strings, and %6.2f will control formatting for each element of a floating-point array.
对于复合操作数（例如切片和结构体），该格式递归地应用于每个操作数的元素，而不是整个操作数。因此 %q 将引用字符串切片的每个元素，而 %6.2f 将控制浮点数组的每个元素的格式。

However, when printing a byte slice with a string-like verb (%s %q %x %X), it is treated identically to a string, as a single item.
但是，当使用类似字符串的动词 (%s %q %x %X) 打印字节切片时，它会被视为与字符串相同的单个项目。

To avoid recursion in cases such as
为了避免在以下情况下递归

```go
type X string
func (x X) String() string { return Sprintf("<%s>", x) }
```

convert the value before recurring:
在重复之前转换值：

```go
func (x X) String() string { return Sprintf("<%s>", string(x)) }
```

Infinite recursion can also be triggered by self-referential data structures, such as a slice that contains itself as an element, if that type has a String method. Such pathologies are rare, however, and the package does not protect against them.
无限递归也可以由自引用数据结构触发，例如包含自身作为元素的切片（如果该类型具有 String 方法）。然而，这种病症很少见，并且该包装不能预防它们。

When printing a struct, fmt cannot and therefore does not invoke formatting methods such as Error or String on unexported fields.
打印结构时，fmt 不能，因此不会在未导出的字段上调用格式化方法，例如 Error 或 String。

#### Explicit argument indexes

In Printf, Sprintf, and Fprintf, the default behavior is for each formatting verb to format successive arguments passed in the call. However, the notation [n] immediately before the verb indicates that the nth one-indexed argument is to be formatted instead. The same notation before a '*' for a width or precision selects the argument index holding the value. After processing a bracketed expression [n], subsequent verbs will use arguments n+1, n+2, etc. unless otherwise directed.
在 Printf、Sprintf 和 Fprintf 中，每个格式化动词的默认行为是格式化调用中传递的连续参数。然而，动词前面的符号 [n] 表示要格式化第 n 个一索引参数。宽度或精度的“*”之前的相同符号选择保存该值的参数索引。处理完括号表达式 [n] 后，后续动词将使用参数 n+1、n+2 等，除非另有指示。

For example,

```go
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```

will yield "22 11", while

```go
fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
```

equivalent to

```go
fmt.Sprintf("%6.2f", 12.0)
```

will yield " 12.00". Because an explicit index affects subsequent verbs, this notation can be used to print the same values multiple times by resetting the index for the first argument to be repeated:
将产生“12.00”。由于显式索引会影响后续动词，因此可以通过重置要重复的第一个参数的索引，使用此表示法多次打印相同的值：

```go
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
```

will yield "16 17 0x10 0x11".

#### Format errors

If an invalid argument is given for a verb, such as providing a string to %d, the generated string will contain a description of the problem, as in these examples:
如果为动词提供了无效参数，例如向 %d 提供字符串，则生成的字符串将包含问题的描述，如以下示例所示：

```
Wrong type or unknown verb: %!verb(type=value)
    Printf("%d", "hi"):        %!d(string=hi)

Too many arguments: %!(EXTRA type=value)
    Printf("hi", "guys"):      hi%!(EXTRA string=guys)

Too few arguments: %!verb(MISSING)
    Printf("hi%d"):            hi%!d(MISSING)

Non-int for width or precision: %!(BADWIDTH) or %!(BADPREC)
    Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
    Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi

Invalid or invalid use of argument index: %!(BADINDEX)
    Printf("%*[2]d", 7):       %!d(BADINDEX)
    Printf("%.[2]d", 7):       %!d(BADINDEX)
```

All errors begin with the string "%!" followed sometimes by a single character (the verb) and end with a parenthesized description.
所有错误均以字符串“%!”开头有时后跟单个字符（动词）并以括号内的描述结尾。

If an Error or String method triggers a panic when called by a print routine, the fmt package reformats the error message from the panic, decorating it with an indication that it came through the fmt package. For example, if a String method calls panic("bad"), the resulting formatted message will look like
如果 Error 或 String 方法在被打印例程调用时触发恐慌，则 fmt 包会重新格式化来自恐慌的错误消息，并用它来自 fmt 包的指示来装饰它。例如，如果 String 方法调用panic("bad")，则生成的格式化消息将如下所示

```
%!s(PANIC=bad)
```

The %!s just shows the print verb in use when the failure occurred. If the panic is caused by a nil receiver to an Error or String method, however, the output is the undecorated string, `<nil>`.
%!s 仅显示发生故障时正在使用的打印动词。但是，如果恐慌是由 Error 或 String 方法的 nil 接收器引起的，则输出是未修饰的字符串`<nil>`。

#### Scanning

An analogous set of functions scans formatted text to yield values. **Scan, Scanf and Scanln** read from os.Stdin; **Fscan, Fscanf and Fscanln** read from a specified io.Reader; **Sscan, Sscanf and Sscanln** read from an argument string.
一组类似的函数扫描格式化文本以产生值。 Scan、Scanf 和 Scanln 从 os.Stdin 读取； Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 读取； Sscan、Sscanf 和 Sscanln 从参数字符串中读取。

Scan, Fscan, Sscan treat newlines in the input as spaces.
Scan、Fscan、Sscan 将输入中的换行符视为空格。

Scanln, Fscanln and Sscanln stop scanning at a newline and require that the items be followed by a newline or EOF.
Scanln、Fscanln 和 Sscanln 在换行符处停止扫描，并要求项目后跟换行符或 EOF。

Scanf, Fscanf, and Sscanf parse the arguments according to a format string, analogous to that of Printf. In the text that follows, 'space' means any Unicode whitespace character except newline.
Scanf、Fscanf 和 Sscanf 根据格式字符串解析参数，类似于 Printf。在接下来的文本中，“空格”表示除换行符之外的任何 Unicode 空白字符。

In the format string, a verb introduced by the % character consumes and parses input; these verbs are described in more detail below. A character other than %, space, or newline in the format consumes exactly that input character, which must be present. A newline with zero or more spaces before it in the format string consumes zero or more spaces in the input followed by a single newline or the end of the input. A space following a newline in the format string consumes zero or more spaces in the input. Otherwise, any run of one or more spaces in the format string consumes as many spaces as possible in the input. Unless the run of spaces in the format string appears adjacent to a newline, the run must consume at least one space from the input or find the end of the input.
在格式字符串中，%字符引入的动词消耗并解析输入；下面更详细地描述这些动词。格式中除 %、空格或换行符之外的字符完全消耗该输入字符，该字符必须存在。格式字符串中前面有零个或多个空格的换行符会占用输入中的零个或多个空格，后跟单个换行符或输入末尾。格式字符串中换行符后面的空格会占用输入中的零个或多个空格。否则，格式字符串中的任何一个或多个空格都会在输入中消耗尽可能多的空格。除非格式字符串中的连续空格出现在换行符附近，否则该连续空格必须至少占用输入中的一个空格或找到输入的末尾。

The handling of spaces and newlines differs from that of C's scanf family: in C, newlines are treated as any other space, and it is never an error when a run of spaces in the format string finds no spaces to consume in the input.
空格和换行符的处理与 C 的 scanf 系列不同：在 C 中，换行符被视为任何其他空格，并且当格式字符串中的一系列空格在输入中找不到可使用的空格时，它永远不会出现错误。

The verbs behave analogously to those of Printf. For example, %x will scan an integer as a hexadecimal number, and %v will scan the default representation format for the value. The Printf verbs %p and %T and the flags # and + are not implemented. For floating-point and complex values, all valid formatting verbs (%b %e %E %f %F %g %G %x %X and %v) are equivalent and accept both decimal and hexadecimal notation (for example: "2.3e+7", "0x4.5p-8") and digit-separating underscores (for example: "3.14159_26535_89793").
这些动词的行为与 Printf 的动词类似。例如，%x 将扫描整数作为十六进制数，%v 将扫描该值的默认表示格式。 Printf 动词 %p 和 %T 以及标志 # 和 + 未实现。对于浮点和复数值，所有有效的格式动词（%b %e %E %f %F %g %G %x %X 和 %v）都是等效的，并且接受十进制和十六进制表示法（例如：“2.3 e+7"、"0x4.5p-8"）和数字分隔下划线（例如："3.14159_26535_89793"）。

Input processed by verbs is implicitly space-delimited: the implementation of every verb except %c starts by discarding leading spaces from the remaining input, and the %s verb (and %v reading into a string) stops consuming input at the first space or newline character.
动词处理的输入是隐式空格分隔的：除 %c 之外的每个动词的实现都首先从剩余输入中丢弃前导空格，并且 %s 动词（和读入字符串的 %v）停止在第一个空格处消耗输入或换行符。

The familiar base-setting prefixes 0b (binary), 0o and 0 (octal), and 0x (hexadecimal) are accepted when scanning integers without a format or with the %v verb, as are digit-separating underscores.
当扫描没有格式或带有 %v 动词的整数时，会接受熟悉的基本设置前缀 0b（二进制）、0o 和 0（八进制）以及 0x（十六进制），数字分隔下划线也是如此。

Width is interpreted in the input text but there is no syntax for scanning with a precision (no %5.2f, just %5f). If width is provided, it applies after leading spaces are trimmed and specifies the maximum number of runes to read to satisfy the verb. For example,
宽度在输入文本中进行解释，但没有用于精确扫描的语法（没有 %5.2f，只有 %5f）。如果提供了宽度，则它会在修剪前导空格后应用，并指定要读取以满足动词的最大符文数。例如，

```go
Sscanf(" 1234567 ", "%5s%d", &s, &i)
```

will set s to "12345" and i to 67 while

```go
Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
```

will set s to "12" and i to 34.

In all the scanning functions, a carriage return followed immediately by a newline is treated as a plain newline (\r\n means the same as \n).
在所有扫描函数中，回车后紧跟换行符被视为普通换行符（\r\n 与 \n 含义相同）。

In all the scanning functions, if an operand implements method Scan (that is, it implements the Scanner interface) that method will be used to scan the text for that operand. Also, if the number of arguments scanned is less than the number of arguments provided, an error is returned.
在所有扫描函数中，如果操作数实现了 Scan 方法（即，它实现了 Scanner 接口），则该方法将用于扫描该操作数的文本。此外，如果扫描的参数数量少于提供的参数数量，则会返回错误。

All arguments to be scanned must be either pointers to basic types or implementations of the Scanner interface.
所有要扫描的参数必须是指向基本类型的指针或 Scanner 接口的实现。

Like Scanf and Fscanf, Sscanf need not consume its entire input. There is no way to recover how much of the input string Sscanf used.
与 Scanf 和 Fscanf 一样，Sscanf 不需要消耗其整个输入。无法恢复 Sscanf 使用了多少输入字符串。

Note: Fscan etc. can read one character (rune) past the input they return, which means that a loop calling a scan routine may skip some of the input. This is usually a problem only when there is no space between input values. If the reader provided to Fscan implements ReadRune, that method will be used to read characters. If the reader also implements UnreadRune, that method will be used to save the character and successive calls will not lose data. To attach ReadRune and UnreadRune methods to a reader without that capability, use bufio.NewReader.
注意：Fscan 等可以读取它们返回的输入之后的一个字符（符文），这意味着调用扫描例程的循环可能会跳过某些输入。仅当输入值之间没有空格时，这通常才会出现问题。如果提供给 Fscan 的读取器实现了 ReadRune，则该方法将用于读取字符。如果阅读器也实现了 UnreadRune，则该方法将用于保存字符，连续调用不会丢失数据。要将 ReadRune 和 UnreadRune 方法附加到没有该功能的读取器，请使用 bufio.NewReader。

### index

```go
func Print(a ...any) (n int, err error)
func Printf(format string, a ...any) (n int, err error)
func Println(a ...any) (n int, err error)
```

```go
func Fprint(w io.Writer, a ...any) (n int, err error)
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
func Fprintln(w io.Writer, a ...any) (n int, err error)
```

--------------------------------------

```go
func Scan(a ...any) (n int, err error)
func Scanf(format string, a ...any) (n int, err error)
func Scanln(a ...any) (n int, err error)
```

```go
func Sscan(str string, a ...any) (n int, err error)
func Sscanf(str string, format string, a ...any) (n int, err error)
func Sscanln(str string, a ...any) (n int, err error)
```

```go
func Fscan(r io.Reader, a ...any) (n int, err error)
func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
func Fscanln(r io.Reader, a ...any) (n int, err error)
```

--------------------------------------

```go
func Sprint(a ...any) string
func Sprintf(format string, a ...any) string
func Sprintln(a ...any) string
```

```go
func FormatString(state State, verb rune) string
```

--------------------------------------

```go
func Append(b []byte, a ...any) []byte
func Appendf(b []byte, format string, a ...any) []byte
func Appendln(b []byte, a ...any) []byte
```

--------------------------------------

```go
func Errorf(format string, a ...any) error
```
