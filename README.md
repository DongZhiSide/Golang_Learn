# testify

[testify](https://github.com/stretchr/testify) 模块提供了一套完整的测试框架, 包括断言, mock, suite 等功能, 适用于单元测试与集成测试. 例如不用手动再判断是否`nil`了.

Testify 包含以下包:

- `github.com/stretchr/testify/assert`包提供了一套全面的断言函数, 与 Go 测试系统相关联.
- `github.com/stretchr/testify/require`提供了相同的断言, 但作为致命的检查.
- `github.com/stretchr/testify/mock`包提供了一个系统, 可以模拟你的对象并验证调用是否按预期进行.
- `github.com/stretchr/testify/suite`包提供了一个基本结构, 用于将结构体用作测试套件, 并将这些结构体上的方法作为测试. 它包含了接口上的设置/拆卸功能.

## testify 的参考资料

1. [testify go doc.](https://pkg.go.dev/github.com/stretchr/testify)
