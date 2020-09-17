# mahjim
a tool to generate mahjong image

这是一个用来快速生成麻将图案的工具.

支持日麻/国标两种图案

![国标暗杠白](https://mahjim.black-desk.cn/+2白+?country=cn)

![日麻暗杠白](https://mahjim.black-desk.cn/+2白+?country=jp)

支持吃,碰,杠的表示

![_123s](https://mahjim.black-desk.cn/_123s)

![1_11s](https://mahjim.black-desk.cn/1_11s?country=cn)

![国标杠白](https://mahjim.black-desk.cn/3%E7%99%BD_%E7%99%BD?country=cn)

在我书写这个文档的时候, [这里](https://mahjim.black-desk.cn)有一个demo, 如果需要将这个工具应用在您的文章中, 请保存生成的图片, 而不要只留有一个链接在你的文档中. 

## 语法说明

### 使用说明

你可以直接使用上面提到的demo, 它运行在 https://mahjim.black-desk.cn, 你可以这样来使用它:

下面我们会说到 `123m` 代表 一万二万三万, 那么当你访问 https://mahjim.black-desk.cn/123m 就可以得到一张如下所示的图片:

![123m](https://mahjim.black-desk.cn/123m)

本文档中的图片也都是这么生成的.

配合markdown编辑器使用本工具会更方便文档的书写

### 基础语法

日麻中有用 p s m 来表示饼子, 条 (索) 子, 万字的习惯, 并且字牌按照 东南西北白发中 的顺序编号后用 z 表示, 比如 5z 表示白.

当参数 country=cn 时 z 的顺序为 东南西北中发白

`1234567z`

![1234567z](https://mahjim.black-desk.cn/1234567z|)


`1234567z?country=cn`:

![1234567z?country=cn](https://mahjim.black-desk.cn/1234567z?country=cn)

所以我们可以用 `1s2s3s` 表示一条二条三条:

![1s2s3s](https://mahjim.black-desk.cn/1s2s3s)

当然支持简写, 比如 `123s` 和 `1s2s3s`是等价的:

`123s`:

![123s](https://mahjim.black-desk.cn/123s)

`1s2s3s`:

![123s](https://mahjim.black-desk.cn/1s2s3s)

为了输入方便, 我们也可以直接用汉字来代表字牌, 例如 `白白白` 代表三张白:

`白白白`:

![白白白](https://mahjim.black-desk.cn/白白白?country=jp)

当然我们可以输入 `666z` :

`666z`:

![666z](https://mahjim.black-desk.cn/666z?country=jp)

特别地,当使用汉字表示牌时可以使用简写, 例如`3中` 和  `中中中` 也是等价的:

`3中`:

![3中](https://mahjim.black-desk.cn/3中?country=jp)

特别的, `+` 表示牌背:

`+`:

![+](https://mahjim.black-desk.cn/+?country=jp)

也支持春夏秋冬和梅兰竹菊, 但是不能以z的形式输入,只能使用汉字:

`春夏秋冬梅兰竹菊`

![春夏秋冬梅兰竹菊](https://mahjim.black-desk.cn/春夏秋冬梅兰竹菊)



### 吃 碰 杠与加杠

当发生吃碰杠的时候, 有一些牌会被横置, 在以上规则的基础上, 用 `_1m` 表示横置的一万:

![_1m](https://mahjim.black-desk.cn/_1m)

那么我吃上家的 1m 形成的面子 123m, 可以如此表示

 `_123m`:

![_123m](https://mahjim.black-desk.cn/_123m)

碰的例子也类似:

`中_中中`:

![中_中中](https://mahjim.black-desk.cn/中_中中)

而加杠形成的双横置以 `^` 表示, 例如 `2中^中`, 表示 碰下家的中之后加杠:

![2中^中](https://mahjim.black-desk.cn/2中^中)

当然, 使用 `77^7z`, 也可以表示相同的事情:

![77^7z](https://mahjim.black-desk.cn/77^7z)

你也可以用另一种方式表示暗杠:

`+^中+`:

![+^中+](https://mahjim.black-desk.cn/+^中+)

### 日麻与国标

可以在整个字符串的后面加上参数 `country=cn` 或者 `country=jp` 来说明希望生成 国标/日麻 图案, 如果不做说明默认是日麻图案, 例如以下分别是 `3白?country=cn` 和 `3白?country=jp`, 不添加时默认为日麻图案.

![3白?country=cn](https://mahjim.black-desk.cn/3白?country=cn)

![3白?country=jp](https://mahjim.black-desk.cn/3白?country=jp)

### 间隔

可以在牌与牌之间插入 `|` 来制造空隙, 每一个 `|` 是 1/10 个牌的宽度, 可以通过形如 `3|` 的方式来增加间隙, 例如以下分别是 `123s|456p` 和 `123s5|456p`

![123s|456p](https://mahjim.black-desk.cn/123s|456p)

![123s5|456p](https://mahjim.black-desk.cn/123s5|456p)

### 缩小和放大

默认生成的牌的大小为 70px x 100px, 如果需要更大或者更小的图片请使用缩放功能, 例如在整个字符串的最后加上参数 `scale=1.2` 可以将图片放大为原来的 1.2 倍, 该参数不能大于 10, 以下分别是`123m?scale=1.2`和`123m?scale=0.5`

![123m?scale=1.2](https://mahjim.black-desk.cn/123m?scale=1.2)

![123m?scale=0.5](https://mahjim.black-desk.cn/123m?scale=0.5)

### 牌河

参数`river=true`可以开启牌河生成模式, 牌河模式下 `^` 的意义是模切, `_` 一样是横置, 一排放置 6 张牌,以下是`1^233_4^5^6^7^8^9s?scale=0.7&river=true`

![1^233_4^5^6^7^8^9s](https://mahjim.black-desk.cn/1^233_4^5^6^7^8^9s?scale=0.7&river=true)

## 鸣谢

感谢图片来源 Neerdge qq820812450

## TODO

- [ ] 性能测试
- [ ] 解析器复用
- [ ] 图片缓存