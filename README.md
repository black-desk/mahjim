# mahjim
a tool to generate mahjong image

这是一个用来快速生成麻将图案的工具.

支持日麻/国标两种图案

![国标杠白](https://mahjim.black-desk.cn/cn|+2白+)

![国标杠白](https://mahjim.black-desk.cn/jp|+2白+)

支持吃,碰,杠的表示

![国标杠白](https://mahjim.black-desk.cn/_123s)

![国标杠白](https://mahjim.black-desk.cn/1_11s)

![国标杠白](https://mahjim.black-desk.cn/3白_白)

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

所以我们可以用 `1s2s3s` 表示一条二条三条:

![1s2s3s](https://mahjim.black-desk.cn/1s2s3s)

当然支持简写, 比如 `123s` 和 `1s2s3s`是等价的:

![123s](https://mahjim.black-desk.cn/123s)

为了输入方便, 我们也可以直接用汉字来代表字牌, 例如 `白白白` 代表三张白:

![白白白](https://mahjim.black-desk.cn/白白白)

当然也有简写, `3白` 和  `白白白` 也是等价的:

![3白](https://mahjim.black-desk.cn/3白)

特别的, `+` 表示牌背:

![国标杠白](https://mahjim.black-desk.cn/+)

### 吃 碰 杠与加杠

当发生吃碰杠的时候, 有一些牌会被横置, 在以上规则的基础上, 用 `_1m` 表示横置的一万:

![_1m](https://mahjim.black-desk.cn/_1m)

那么我吃上家的 1m 形成的面子 123m, 可以如此表示 `_123m`:

![_123m](https://mahjim.black-desk.cn/_123m)

碰的例子也类似:

![中_中中](https://mahjim.black-desk.cn/中_中中)

而加杠形成的双横置以 `^` 表示, 例如 `2中^中`, 表示 碰下家的中之后加杠:

![2中^中](https://mahjim.black-desk.cn/2中^中)

当然, 使用 `77^7z`, 也可以表示相同的事情:

![77^7z](https://mahjim.black-desk.cn/77^7z)

### 日麻与国标

可以在整个字符串的前面加上 `cn|` 或者 `jp|` 来说明希望生成 国标/日麻 图案, 如果不做说明默认是日麻图案, 例如以下分别是 `cn|3白` 和 `jp|3白`

![cn|3白](https://mahjim.black-desk.cn/cn|3白)

![jp|3白](https://mahjim.black-desk.cn/jp|3白)

## 鸣谢

感谢图片来源 Neerdge qq820812450

## TODO

- [ ]  牌之间的空隙可调
- [ ]  牌背颜色可调
- [ ]  支持图包中的数字牌和字母牌
- [ ]  生成牌河
- [ ]  生成牌山
- [ ]  横置牌右侧的空隙
- [ ]  新的图片寻找机制

