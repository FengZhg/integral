# integral
- 一种积分模块
- 计划进行不同实现方式的性能比较





## 实现计划
| Branch Name | 实现方式 | 其他 |
| ----- | ---- | ---- |
| normal | 基本普通实现 |  |
| eventLoop | 为redis、mysql、pulsar建立eventLoop | pulsar-client-go的consumer内部就是eventloop实现的 |
| language | 语言层面(协程绑定线程之类的)优化 | 火焰图、函数调用图优化 |
| kernel | 内核层面优化 | |

**目前只列了自己一下子想到的，等我再研究研究，想到啥优化再补充**

