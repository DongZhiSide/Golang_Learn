# casbin

[Casbin](https://github.com/casbin/casbin) 可以：

- 支持自定义请求的格式，默认的请求格式为{subject, object, action}。
- 具有访问控制模型model和策略policy两个核心概念。
- 支持RBAC中的多层角色继承，不止主体可以有角色，资源也可以具有角色。
- 支持内置的超级用户 例如：root 或 administrator。超级用户可以执行任何操作而无需显式的权限声明。
- 支持多种内置的操作符，如 keyMatch，方便对路径式的资源进行管理，如 /foo/bar 可以映射到 /foo*

Casbin 不能：

- 身份认证 authentication（即验证用户的用户名和密码），Casbin 只负责访问控制。应该有其他专门的组件负责身份认证，然后由 Casbin 进行访问控制，二者是相互配合的关系。
- 管理用户列表或角色列表。 Casbin 认为由项目自身来管理用户、角色列表更为合适， 用户通常有他们的密码，但是 Casbin 的设计思想并不是把它作为一个存储密码的容器。 而是存储RBAC方案中用户和角色之间的映射关系。

> 这个casbin的理论建立在[一种基于元模型的访问控制策略描述语言](https://www.jos.org.cn/html/2020/2/5624.htm#outline_anchor_7)这篇论文之上的。这是中国人写的，因此这个仓库的最初开发也是中国人写的。罗杨作者，北大的助理教授。
>
> 太强了，我等只能仰望。

casbin 相关概念并不难懂，觉得难以理解的可以看看参考1，作者写的非常详细。

## casbin 参考资料

1. [Go 每日一库之 casbin.](https://darjun.github.io/2020/06/12/godailylib/casbin/)
2. [casbin 官方 工作原理(老版本).](https://v1.casbin.org/docs/zh-CN/how-it-works)
3. [casbin 官方.](https://casbin.org/zh/docs/get-started)
4. [casbin 编辑器.](https://casbin.org/editor/)
