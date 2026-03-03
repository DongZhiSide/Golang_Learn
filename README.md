# swaggo

[swaggo](https://github.com/swaggo/swag/): swagger(又称 openapi)的 go 实现. 当前`v1.16.6`版本的 swaggo 不支持 swagger 3.0.

## swaggo 的参考资料

1. [swag.](https://golang.halfiisland.com/community/pkgs/web/swag.html)
2. [swagger 中文文档.](https://swagger.org.cn/docs/)
3. [swaggo github.](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
4. [swaggo 示例.](https://github.com/swaggo/swag/tree/master/example)

## swaggo 的概念与术语

swaggo 通过识别代码中的注释来生成 swagger 文档, 因此 swaggo 其实是一个外部程序, 它会扫描代码中的注释, 然后再生成 swagger 文档.
