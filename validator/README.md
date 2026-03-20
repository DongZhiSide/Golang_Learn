# validator

[validator](https://github.com/go-playground/validator) 模块基于标签实现结构体及独立字段的值校验功能.

该包具备以下独特特性:

- 支持通过校验标签或自定义校验器实现跨字段与跨结构体校验.
- 支持切片, 数组及映射的层级遍历, 可对多维字段的任意或全部层级执行校验.
- 支持深入映射的键与值进行校验.
- 在校验前先判断接口类型的底层类型, 以实现接口类型的校验处理.
- 支持处理自定义字段类型, 例如`sql/driver`包下的`Valuer`类型.
- 支持校验标签别名, 可将多个校验规则映射至单个标签, 便于在结构体上定义校验规则.
- 支持提取自定义字段名, 例如在校验时可指定提取 JSON 字段名, 并在返回的`FieldError`中使用.
- 支持可定制化且具备国际化适配能力的错误提示信息.

## validator 的参考资料

1. [validator go doc.](https://pkg.go.dev/github.com/go-playground/validator)
2. [validator example.](https://github.com/go-playground/validator/tree/master/_examples)
