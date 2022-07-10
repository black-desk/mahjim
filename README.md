# mahjim
a tool to generate mahjong image

这是一个用来快速生成麻将图案的工具.

支持日麻/国标两种图案

![国标](https://mahjim.black-desk.cn/1s|1234567z?country=cn)

![日麻](https://mahjim.black-desk.cn/1s|1234567z?country=jp)

支持吃,碰,明/暗杠的表示

![_123s](https://mahjim.black-desk.cn/_123m)

![2_22s](https://mahjim.black-desk.cn/2_22s)

![33_3p](https://mahjim.black-desk.cn/33_3p)

![+2白+](https://mahjim.black-desk.cn/+2白+)

![+^中+](https://mahjim.black-desk.cn/+^中+)

![发^发发](https://mahjim.black-desk.cn/^发2发)

牌背颜色有两种可选:

![blue](https://mahjim.black-desk.cn/+?color=blue)

![orange](https://mahjim.black-desk.cn/+?color=orange)

在我书写这个文档的时候, [这里](https://mahjim.black-desk.cn)有一个demo, 如果需要将这个工具应用在您的文章中, 请保存生成的图片, 而不要只留有一个链接在你的文档中. 

你也可以将这个程序运行在你自己的电脑上, 来获得更稳定的体验.

## 语法说明



### 使用方法

以上面提到的 demo 为实例, 我们简单学习一下如何使用这个工具, 它运行在 https://mahjim.black-desk.cn 这个地址, 你可以像这样来使用它:

字符串 `123m` 代表 一万二万三万, 那么当你访问 https://mahjim.black-desk.cn/123m 就可以得到一张如下所示的图片:

![123m](https://mahjim.black-desk.cn/123m)

本文档中的图片也都是这么生成的.

推荐配合 markdown 编辑器使用本工具, 更方便书写



### 描述语法

日麻中有用 p s m 来表示饼子, 条 (索) 子, 万字的习惯, 并且字牌按照 东南西北白发中 的顺序编号后用 z 表示, 比如 5z 表示白. 但是中国模式下, 字牌的顺序并不一样, 顺序为 东南西北中发白

`1234567z`

![1234567z](https://mahjim.black-desk.cn/1234567z)

`1234567z?country=cn`:

![1234567z?country=cn](https://mahjim.black-desk.cn/1234567z?country=cn)

看几个简单的例子: 可以用 `1s2s3s` 表示一条二条三条:

![1s2s3s](https://mahjim.black-desk.cn/1s2s3s)

和上面字牌的情况一样, 数字牌也支持简写, 比如 `123p` 和 `1p2p3p`是等价的:

`123p`:

![123s](https://mahjim.black-desk.cn/123p)

`1p2p3p`:

![123p](https://mahjim.black-desk.cn/1p2p3p)

为了输入方便, 我们也可以直接用汉字来代表字牌, 例如 `白白白` 代表三张白:

`白白白`:

![白白白](https://mahjim.black-desk.cn/白白白?country=jp)

当然我们可以输入 `666z` 来表示三张发:

`666z`:

![666z](https://mahjim.black-desk.cn/666z?country=jp)

特别地,当使用汉字表示牌时可以简写, 例如 `3中` 和  `中中中` 也是等价的:

`3中`:

![3中](https://mahjim.black-desk.cn/3中?country=jp)

特别的, `+` 表示牌背:

`+`:

![+](https://mahjim.black-desk.cn/+?country=jp)

多个并排的牌背也可以简写, 比如 `4+`:

`4+`:

![4+](https://mahjim.black-desk.cn/4+)

也支持春夏秋冬和梅兰竹菊, 但是不能以z的形式输入,只能使用汉字:

`春夏秋冬梅兰竹菊`

![春夏秋冬梅兰竹菊](https://mahjim.black-desk.cn/春夏秋冬梅兰竹菊)

汉字输入**暂时不支持繁体**



### 牌的变化

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

本工具只是一个图片生成器, 并不检查牌是否合理, 所以实际上你也可以摆出一些很诡异的牌型, 比如:

`2中^发2白`:

![2中^发白](https://mahjim.black-desk.cn/2中^发2白)

`_6_6_6_6p`:

![_6_6_6_6p](https://mahjim.black-desk.cn/_6_6_6_6p)

虽然不知道有什么用, 但是牌背是可以被横置的:

`_+`:

![_+](https://mahjim.black-desk.cn/_+)



### 控制参数

#### 日麻与国标 country

可以在整个字符串的后面加上参数 `country=cn` 或者 `country=jp` 来说明希望生成 国标/日麻 图案, 如果不做说明默认是日麻图案, 例如以下分别是 `3白?country=cn` 和 `3白?country=jp`, 不添加此参数时默认为日麻图案.

![3发?country=cn](https://mahjim.black-desk.cn/3发s?country=cn)

![3发?country=jp](https://mahjim.black-desk.cn/3发?country=jp)

这里展示一下两种风格下所有的牌画:

![cn](https://mahjim.black-desk.cn/+123456789s123456789m123456789p1234567z春夏秋冬梅兰竹菊?country=cn&river=true&scale=0.7)

![jp](https://mahjim.black-desk.cn/+0123456789s0123456789m0123456789p1234567z?country=jp&river=true&scale=0.7)

#### 缩小和放大 scale

默认生成的每张牌的大小为 70px x 100px, 如果需要更大或者更小的图片请使用缩放功能, 例如在整个字符串的最后加上参数 `scale=1.2` 可以将图片放大为原来的 1.2 倍, 该参数不能大于 10, 如果输入了一个 >10 的参数, 那么返回结果就是一张 x10 的图, 以下分别是`123m?scale=1.2`和`123m?scale=0.5`

![123m?scale=1.2](https://mahjim.black-desk.cn/123m?scale=1.2)

![123m?scale=0.5](https://mahjim.black-desk.cn/123m?scale=0.5)

#### 牌背颜色 color

可以使用 `color=orange` 来获得橙色的牌背, 此参数不存在时默认为蓝色牌背.

![+?color=orange](https://mahjim.black-desk.cn/+?color=orange)



#### 牌河 river

参数`river=true`可以开启牌河生成模式, 牌河模式下 `^` 的意义是模切, `_` 一样是横置, 一排放置 6 张牌,以下是`1^233_4^5^6^7^8^9s?scale=0.7&river=true`

![1^233_4^5^6^7^8^9s^7^8^9m?scale=0.7&river=true](https://mahjim.black-desk.cn/1^233_4^5^6^7^8^9s^7^8^9m?scale=0.7&river=true)

出于实现方便, 立直牌靠下放置.

TODO: 这个功能目前并不完善

- 立直之后应该默认就全是模切.
- 无法表示模切立直.
- 打满 3 排之后不再换行.

#### 多参数

同时使用多个参数时, 使用 `&` 来连接不同的参数. 可以看一下上节的牌河生成, 很好懂的.



### 间隔

可以在牌与牌之间插入 `|` 来制造空隙, 每一个 `|` 是 1/10 个牌的宽度, 可以通过形如 `3|` 的方式来增加间隙, 例如以下分别是 `123s|456p` 和 `123s5|456p`

![123s|456p](https://mahjim.black-desk.cn/123s|456p)

![123s5|456p](https://mahjim.black-desk.cn/123s5|456p)

由于要使用`n|`的方式表示间隔的大小 ,所以 123m 456m 并不能写成 `123|456m`. 需要写成 `123m|456m`



## 本地安装

你可以在 Release 页面下载与系统对应的可执行文件, 通常来说, 作为一个 windows 用户, 你需要下载 mahjim_windows_amd64.exe, 双击运行之后它监听的地址是 `localhost:8080/`, 之后你只要在浏览器里打开这个地址(你可以点击 [这个链接](http://localhost:8080/) 快速打开), 之后就可以使用了. 如果你使用 markdown 就可以通过输入 `![](http://localhost:8080/123s)`  来获得一张图片了.

## Docker

使用命令 `docker run -d --name mahjim -p 8080:8080 blackddesk/mahjim` 即可启动容器

## 鸣谢

感谢图片来源 Neerdge qq820812450

## TODO

- [ ] 性能测试
- [ ] 牌河模式
