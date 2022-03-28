# integral
- 一种积分模块
- 计划进行不同实现方式的性能比较





## 实现计划
| Banch Name | 实现方式 | 其他 |
| ----- | ---- | ---- |
| normal | 基本普通实现 |  |
| eventLoop | 为redis、mysql、pulsar建立eventLoop | pulsar-client-go的consumer内部就是eventloop实现的 |
| optimization | 从内核和go语言层面进行优化 |  |

**想到啥优化再补充**

