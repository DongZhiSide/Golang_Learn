# wire

[wire](https://github.com/google/wire): wire 是一个由 Google 开发的自动依赖注入框架, 专门用于 Go 语言. wire 通过代码生成而非运行时反射来实现依赖注入, 这与许多其他语言中的依赖注入框架不同.

这种方法使得注入的代码在编译时就已经确定, 从而提高了性能并保证了代码的可维护性.

## wire 的概念与术语

1. Provider: 提供依赖的函数, 依赖注入器会调用 Provider 函数来生成依赖.
2. Injector: 依赖注入器, 依赖注入器会调用 Provider 函数来生成依赖, 并将依赖注入到需要依赖的地方.

## wire 的参考资料

1. [万字长文: 在 Go 中如何优雅的使用 wire 依赖注入工具提高开发效率? 上篇.](https://segmentfault.com/a/1190000044962140)
2. [万字长文: 在 Go 中如何优雅的使用 wire 依赖注入工具提高开发效率? 下篇.](https://segmentfault.com/a/1190000044962144)
3. [Wire Tutorial github.](https://github.com/google/wire/blob/main/_tutorial/README.md)
4. [Wire User Guide github.](https://github.com/google/wire/blob/main/docs/guide.md)