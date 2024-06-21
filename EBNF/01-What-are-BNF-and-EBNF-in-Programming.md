# What are BNF and EBNF in Programming?

* https://www.freecodecamp.org/news/what-are-bnf-and-ebnf

As programmers, we communicate with computers through many languages: Python, JavaScript, SQL, C... you name it. But do you know how the creators of these languages precisely describe their syntax to us, leaving no room for doubt?
作为程序员，我们通过多种语言与计算机进行通信：Python、JavaScript、SQL、C……凡是你能想到的语言。但你知道这些语言的创造者如何向我们精确地描述它们的语法，不留任何怀疑的余地吗？

They could've relied on plain English, but that would not be a good solution because of the potential verbosity and ambiguity. So they used specially designed languages for it.
他们可以依靠简单的英语，但这不是一个好的解决方案，因为可能会出现冗长和歧义。所以他们使用了专门设计的语言。

In this article, you'll learn about two of these widely used languages: **BNF** and **EBNF**.
在本文中，您将了解其中两种广泛使用的语言：BNF 和 EBNF。

Another fascinating aspect of these special languages or notations is that you can write the grammar of your own language using them and give it as input to some magical computer programs called "parser generators". These can output other programs capable of parsing any text according to the grammar you used. How amazing is that?
这些特殊语言或符号的另一个令人着迷的方面是，您可以使用它们编写自己语言的语法，并将其作为一些称为“解析器生成器”的神奇计算机程序的输入。这些可以输出能够根据您使用的语法解析任何文本的其他程序。那有多神奇？

This feature can save you a lot of time since manually writing such programs is challenging and time-consuming.
此功能可以为您节省大量时间，因为手动编写此类程序既具有挑战性又耗时。

Before learning (E)BNF, it's helpful to be able to distinguish between syntax and semantics. So let's start from there.
在学习 (E)BNF 之前，能够区分语法和语义会很有帮助。那么让我们从那里开始吧。

## Syntax vs Semantics in Programming Languages
编程语言中的语法与语义

Syntax refers to the structure of the elements of a language based on its type. On the other hand, semantics are all about the meaning.
语法是指基于语言类型的语言元素的结构。另一方面，语义都是关于意义的。

Something written syntactically correctly in a language can be completely meaningless. And no text can be meaningful if its syntax is incorrect.
用一种语言写出的语法正确的东西可能完全没有意义。如果语法不正确，任何文本都没有意义。

Two of the most famous sentences regarding syntax and semantics are [composed by Noam Chomsky](https://en.wikipedia.org/wiki/Colorless_green_ideas_sleep_furiously):
关于语法和语义的最著名的两个句子是由诺姆·乔姆斯基（Noam Chomsky）撰写的：

1. Colorless green ideas sleep furiously.
   无色的绿色思想疯狂地沉睡。

2. Furiously sleep ideas green colorless.
   疯狂地睡意绿色无色。

The first sentence's syntax is correct but it's meaningless. And since the second one is syntactically wrong, it is far from being meaningful.
第一句话的语法是正确的，但毫无意义。由于第二个在语法上是错误的，因此它根本没有意义。

The same is true for programming languages too. Let's look at the following two JavaScript code snippets to see what I mean.
对于编程语言来说也是如此。让我们看一下下面的两个 JavaScript 代码片段，看看我的意思。

The following code is syntactically correct but semantically wrong because it's not possible to reassign something to a constant variable:
以下代码在语法上是正确的，但在语义上是错误的，因为不可能将某些内容重新分配给常量变量：

```js
const name = "Palash";
name = "Akash";
```

The following is syntactically incorrect and thus does not even have any chance to be semantically correct.
以下内容在语法上是不正确的，因此甚至没有任何机会在语义上是正确的。

```js
"Palash" = const name;
"Akash" = name;
```

You check the syntax of your JavaScript code online with a tool like the [Esprima Syntax Validator](https://esprima.org/demo/validate.html).
您可以使用 Esprima Syntax Validator 等工具在线检查 JavaScript 代码的语法。

There are two more concepts you need to understand before learning to read BNF/EBNF.
在学习阅读 BNF/EBNF 之前，您还需要了解两个概念。

## Terminals and Non-Terminals
终端和非终端

BNF/EBNF is usually used to specify the grammar of a language. Grammar is a set of *rules* (also called **production rules**). Here language refers to nothing but a set of strings that are valid according to the rules of its grammar.
BNF/EBNF通常用于指定语言的语法。语法是一组规则（也称为产生式规则）。这里的语言只是指一组根据其语法规则有效的字符串。

A BNF/EBNF grammar description is an unordered list of rules. *Rules* are used to define *symbols* with the help of other symbols.
BNF/EBNF 语法描述是一个无序列表的规则。规则用于在其他符号的帮助下定义符号。

You can think of *symbols* as the building blocks of grammar. There are two kinds of symbols:
您可以将符号视为语法的构建块。符号有两种：

- **Terminal (or Terminal symbol)**: Terminals are strings written within quotes. They are meant to be used as they are. Nothing is hidden behind them. For example `"freeCodeCamp"` or `"firefly"`.
  终端（或终端符号）：终端是写在引号内的字符串。它们应该按原样使用。他们背后没有隐藏任何东西。例如 `"freeCodeCamp"` 或 `"firefly"` 。

- **Non-terminal (or Non-terminal symbol)**: Sometimes we need a name to refer to something else. These are called *non-terminals*. In BNF, *non-terminal* names are written within angle brackets (for example `<statement>`), while in EBNF they don't usually use brackets (for example `statement`).
  非终结符（或非终结符）：有时我们需要一个名称来引用其他东西。这些称为非终结符。在 BNF 中，非终结符名称写在尖括号内（例如 `<statement>` ），而在 EBNF 中，它们通常不使用括号（例如 `statement` ）。

The whole language is derived from a single *non-terminal* symbol. This is called the **start** or **root symbol** of the grammar. By convention, it is written as the first non-terminal in the BNF/EBNF grammar description.
整个语言源自单个非终结符。这称为语法的开始或根符号。按照惯例，它被写为 BNF/EBNF 语法描述中的第一个非终结符。

Finally, you are ready to learn BNF. It's easier than you might think it is.
最后，您已经准备好学习 BNF 了。这比您想象的要容易。

## What is BNF? 
什么是BNF？

BNF stands for **B**ackus–**N**aur **F**orm which resulted primarily from the contributions of [John Backus](https://en.wikipedia.org/wiki/John_Backus) and [Peter Naur](https://en.wikipedia.org/wiki/Peter_Naur).
BNF 代表巴科斯-诺尔范式，主要源自约翰·巴克斯和彼得·诺尔的贡献。

The syntax of BNF/EBNF is so simple that many people adopted their styles. So in different places, you will most likely see different styles. If the syntax is different from conventional ones, that's usually documented there. In this article I will use one particular style, just to keep things simple.
BNF/EBNF 的语法非常简单，很多人都采用了他们的风格。所以在不同的地方，你很可能会看到不同的风格。如果语法与传统语法不同，通常会在那里记录。在本文中，我将使用一种特定的样式，只是为了让事情变得简单。

Below is an example of a simple **production rule** in BNF:
下面是 BNF 中简单产生式规则的示例：

```bnf
<something> ::= "content"
```

Each rule in BNF (also in ENBF) has three parts:
BNF（也在 ENBF）中的每条规则都包含三个部分：

- **Left-hand side**: Here we write a non-terminal to define it. In the above example, it is `<something>`.
  左侧：这里我们编写一个非终结符来定义它。在上面的示例中，它是 `<something>` 。

- **`::=`**: This character group separates the **Left hand side** from **Right hand side**. Read this symbol as "is defined as".
  `::=` ：该字符组将左侧与右侧分开。将该符号读作“定义为”。

- **Right-hand side**: The definition of the non-terminal specified on the right-hand side. In the above example, it's `"content"`.
  右侧：右侧指定的非终结符的定义。在上面的示例中，它是 `"content"` 。

The above `<something>` is just one thing fixed thing. Let's now see all the ways you can compose a *non-terminal*.
上面的 `<something>` 只是一件事固定的事情。现在让我们看看组成非终结符的所有方法。

### How to compose a non-terminal
如何编写非终结符

BNF offers two methods to us:
BNF为我们提供了两种方法：

- Sequencing 测序
- Choice 选择

You can just write a combination of one or more terminals or non-terminals in a sequence and the result is their concatenation, with non-terminals being replaced by their content. For example, you can express your breakfast in the following ways:
您可以按顺序编写一个或多个终结符或非终结符的组合，结果是它们的串联，非终结符被它们的内容替换。例如，你可以用以下方式来表达你的早餐：

```bnf
<breakfast> ::= <drink> " and biscuit"
<drink> ::= "tea"
```

It means the only option for breakfast for you is `"tea and biscuit"`. Note that here, the order of symbols is important.
这意味着您早餐的唯一选择是 `"tea and biscuit"` 。请注意，此处符号的顺序很重要。

Let's say someday you want to drink coffee instead of tea. In this case, you can express your possible breakfast items like below:
假设有一天您想喝咖啡而不是茶。在这种情况下，您可以表达您可能的早餐项目，如下所示：

```bnf
<breakfast> ::= <drink> " and biscuit"
<drink> ::= "tea" | "coffee"
```

The `|` operator indicates that the parts separated by it are choices. Which means the non-terminal on the left can be any such part. Here the order is *unimportant*, that is there is no difference between `"tea" | "coffee` and `"coffee" | "tea"`.
`|` 运算符表示由它分隔的部分是选项。这意味着左侧的非终结符可以是任何这样的部分。这里的顺序并不重要，即 `"tea" | "coffee` 和 `"coffee" | "tea"` 之间没有区别。

That is really all you need to know about BNF to read and understand it and even express the syntax of your own language using it. Believe it or not, it's that simple. And yet it can be used to describe the syntax of many programming languages and other kinds of coding languages.
这确实是您需要了解的关于 BNF 的全部信息，以便阅读和理解它，甚至使用它来表达您自己语言的语法。不管你信不信，事情就是这么简单。然而它可以用来描述许多编程语言和其他类型的编码语言的语法。

The thing that makes it possible to break down complex syntax programming languages easily is the ability to define non-terminal symbols recursively.
能够轻松分解复杂语法编程语言的是递归定义非终结符号的能力。

As a simple example let's see how you express one or more digits in BNF:
作为一个简单的例子，让我们看看如何用 BNF 表达一个或多个数字：

```bnf
<digits> ::= <digit> | <digit> <digits>
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

If you want to see a simple real-world example of BNF grammar checkout: [Semver notation](https://semver.org/#backusnaur-form-grammar-for-valid-semver-versions).
如果您想查看 BNF 语法检查的简单现实示例：Semver notation。

## What is EBNF? 
什么是 EBNF？

BNF is fine, but sometimes it can become verbose and hard to interpret. EBNF (which stands for **E**xtended **B**ackus–**N**aur **F**orm) may help you in those cases. For example, the previous example can be written in EBNF like below:
BNF 很好，但有时它会变得冗长且难以解释。在这些情况下，EBNF（扩展巴科斯-诺尔范式）可能会为您提供帮助。例如，前面的示例可以用 EBNF 编写，如下所示：

```ebnf
digits = digit { digit }
digit = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

The braces above mean that its inner part may be repeated 0 or more times. It frees your mind from getting lost in recursion.
上面的大括号意味着它的内部部分可以重复0次或多次。它使您的思想不再迷失在递归中。

One interesting fact is that everything you can express in EBNF can also be expressed in BNF.
一个有趣的事实是，所有可以用 EBNF 表达的东西也可以用 BNF 表达。

EBNF usually uses a slightly different notation than BNF. For example:
EBNF 通常使用与 BNF 略有不同的表示法。例如：

- `::=` becomes just `=`.
  `::=` 变为 `=` 。

- There are no angle brackets around non-terminals.
  非终结符周围没有尖括号。

```ad-info
For concatenation, instead of juxtaposition, some prefer `,` to be more explicit. However, I will not use it here.
对于连接，有些人更喜欢使用 `,`，而不是并列，这样更明确。不过，在这里我不会使用它。
```

Don't assume that these styles to be universal. There are several variants of them and they are usually clear from the context. The more important thing to focus on is the new operations it offers like the braces we've seen above.
不要假设这些样式是通用的。它们有多种变体，通常可以从上下文中清楚地看出。更重要的是要关注它提供的新操作，例如我们上面看到的大括号。

EBNF extends BNF by adding the following 3 operations:
EBNF 通过添加以下 3 个操作来扩展 BNF：

- Option 选项
- Repetition 重复
- Grouping 分组

### Option
选项

Option uses square brackets to make the inner content optional. Example:
Option 使用方括号使内部内容可选。例子：

```ebnf
thing = "water" [ "melon" ]
```

So the above `thing` is either `water` or `watermelon`.
所以上面的 `thing` 要么是 `water` 要么是 `watermelon` 。

### Repetition
重复

Curly braces indicate the inner content may be repeated 0 or more times. You have already seen a good example of it above. Below is a very simple one just to make the idea solid in your mind:
大括号表示内部内容可以重复0次或多次。您已经在上面看到了一个很好的例子。下面是一个非常简单的例子，只是为了让这个想法在你的脑海中更加牢固：

```ebnf
long_google = "Goo" { "o" } "gle"
```

So `"Google"`, `"Gooogle"`, `"Gooooooogle"` are all valid `long_google` non-terminal.
因此 `"Google"` 、 `"Gooogle"` 、 `"Gooooooogle"` 都是有效的 `long_google` 非终结符。

### Grouping
分组

Parentheses can be used to indicate grouping. It means everything they wrap can be replaced with any of the valid strings that the contents of the group represent according to the rules of EBNF. For example:
括号可用于指示分组。这意味着它们包装的所有内容都可以根据 EBNF 规则替换为组内容表示的任何有效字符串。例如：

```ebnf
fly = ("fire" | "fruit") "fly"
```

Here  `fly` is either `"firefly"` or `"fruitfly"`.
这里 `fly` 是 `"firefly"` 或 `"fruitfly"` 。

With BNF we could not do that in one line. It would look like the following in BNF:
对于 BNF，我们无法在一行中完成这一任务。在 BNF 中它看起来像下面这样：

```ebnf
<fly> ::= <type> "fly"
<type> ::= "fire" | "fruit"
```

## The BNF Playground
BNF 游乐场

There is a very nice online playground for BNF and EBNF: [<BNF> Playground](https://bnfplayground.pauliankline.com/).
有一个非常好的 BNF 和 EBNF 在线游乐场：<BNF> Playground。

I recommend you check it out and play with it. It uses a slightly different notation so read the "Grammar Help" section beforehand.
我建议你检查一下并玩一下。它使用稍微不同的符号，因此请事先阅读“语法帮助”部分。

It can test if a string is valid according to the grammar you entered. It can also generate random strings based on your grammar!
它可以根据您输入的语法测试字符串是否有效。它还可以根据您的语法生成随机字符串！

For fun this is the syntax of a poem-like text (credit goes to chatGPT):
为了好玩，这是类似诗歌的文本的语法（归功于 chatGPT）：

```ebnf
<poem> ::= <line> | <line> "\n" <poem>
<line> ::= <noun_phrase> " " <verb_phrase> " " <adjective>
<noun_phrase> ::= "the " <adjective> " " <noun> | <noun>
<verb_phrase> ::= <verb> | <verb> " " <adverb>
<adjective> ::= "red" | "blue" | "green" | "yellow"
<noun> ::= "sky" | "sun" | "grass" | "flower"
<verb> ::= "shines" | "glows" | "grows" | "blooms"
<adverb> ::= "brightly" | "slowly" | "vividly" | "peacefully"
```

Go ahead and copy-paste it into the playground and press the "Generate Random" button to get some mostly meaningless lines of a grammatically correct poem.
继续将其复制粘贴到操场上，然后按“生成随机”按钮，以获得语法正确的诗中一些几乎毫无意义的行。

## Conclusion
结论

BNF and EBNF are simple and powerful notations to write what computer scientists call *context-free grammar*.
BNF 和 EBNF 是简单而强大的符号，用于编写计算机科学家所说的上下文无关语法。

In simple terms it means the expansion of a non-terminal is not dependent on the context (surrounding symbols), that is it's context-free. It is the most widely used grammar form to formalize the syntax of coding languages.
简单来说，这意味着非终结符的扩展不依赖于上下文（周围的符号），即它是上下文无关的。它是形式化编码语言语法的最广泛使用的语法形式。

Here are some resources you might find interesting:
以下是您可能会感兴趣的一些资源：

- [EBNF: How to Describe the Grammar of a Language](https://tomassetti.me/ebnf/)
  EBNF：如何描述语言的语法

- [The language of languages](https://matt.might.net/articles/grammars-bnf-ebnf/)
  语言中的语言

- Parser generators:
  解析器生成器：

  - [ANTLR](https://www.antlr.org/), a very powerful parser generator capable of writing parsers in many languages.
    ANTLR，一个非常强大的解析器生成器，能够用多种语言编写解析器。

  - If you are a JavaScript person like me and want to get started with a parser generator, take a look at [nearly.js](https://nearley.js.org/) for a gentle start.
    如果您是像我一样的 JavaScript 开发者，并且想要开始使用解析器生成器，请查看近乎.js 以轻松开始。

Below are some real-world grammars written using BNF/EBNF or similar notations that you might find interesting:
以下是一些使用 BNF/EBNF 或您可能感兴趣的类似符号编写的现实语法：

- [Lisp](https://iamwilhelm.github.io/bnf-examples/lisp)

- [Lua](https://www.lua.org/manual/5.4/manual.html#9)

- [Semver](https://semver.org/#backusnaur-form-grammar-for-valid-semver-versions)

- [JavaScript](https://tc39.es/ecma262/multipage/grammar-summary.html#sec-grammar-summary)

- [JSX](https://facebook.github.io/jsx/)

- [Python](https://docs.python.org/3/reference/grammar.html)

- [Value Definition Syntax in CSS](https://developer.mozilla.org/en-US/docs/Web/CSS/Value_definition_syntax)
