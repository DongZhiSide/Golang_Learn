# otp

一次性密码(OTP)是一种在纯密码基础上提升安全性的机制. 将基于时间的一次性密码(TOTP)存储在用户手机中, 再结合用户所知信息(密码), 即可轻松实现多因素认证, 且无需依赖短信服务商.

这种密码与 [TOTP](https://zhuanlan.zhihu.com/p/484991482) 相结合的方式, 已被谷歌, GitHub, 脸书, 赛富时等众多主流网站采用. 由于 TOTP 技术标准化程度高且部署范围广, 市面上存在大量对应的移动客户端与软件实现方案.

[otp](https://github.com/pquerna/otp) 模块可助你在自有应用中便捷集成 TOTP 功能, 提升用户账户抵御大规模密码泄露与恶意软件攻击的安全性. 支持:

- 生成二维码图片, 方便用户快速绑定.
- 基于时间的一次性密码算法(TOTP)(RFC 6238 标准), 主流的一次性密码生成方案.
- 基于哈希消息认证码的一次性密码算法(HOTP)(RFC 4226 标准), TOTP 算法的底层基础, 采用计数器机制生成密码.
- 支持上述两种算法的一次性密码生成与校验功能.

## otp 的参考资料

1. [otp go doc.](https://pkg.go.dev/github.com/pquerna/otp)
2. [otp repo example.](https://github.com/pquerna/otp/blob/master/example/main.go)
