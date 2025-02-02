尝试写一个log库

# 预期功能
1. 选择合适的日志级别
    - 日志库应支持常见的日志级别，如 DEBUG, INFO, WARN, ERROR, FATAL 等。
    - 用户可以根据需要设置全局或局部的日志级别，以控制输出哪些级别的日志。
2. 灵活的输出方式
    - 支持将日志输出到不同的目标，如：
    - 标准输出（stdout）
    - 文件
    - 远程日志服务器（如通过网络发送日志）
    - 数据库等
3. 日志格式化
    - 提供多种日志格式选项，例如：
    - 时间戳
    - 日志级别
    - 文件名和行号
    - 自定义前缀或后缀
    - 支持用户自定义日志格式。
4. 并发安全
    - 确保日志库在多线程环境中是线程安全的，特别是在高并发场景下。
    - 可以使用互斥锁（sync.Mutex）或其他同步机制来保证线程安全。
5. 性能优化
    - 尽量减少日志记录的性能开销，特别是在生产环境中。
    - 对于频繁的日志记录操作，可以考虑批量写入或异步写入。
6. 日志轮转
    - 实现日志文件的轮转功能，防止日志文件过大。
    - 支持按大小、时间等方式进行日志轮转，并自动归档旧日志。
7. 错误处理
    - 在日志记录过程中可能会遇到错误（如文件写入失败），确保这些错误能够被合理处理或报告。
8. 扩展性
    - 设计时考虑到未来的扩展需求，例如：
    - 支持插件化架构，允许用户添加自定义的日志处理器。
    - 提供钩子函数，允许用户在日志记录前后执行自定义逻辑。
9. 配置管理
    - 提供灵活的配置方式，如通过环境变量、配置文件或代码配置来调整日志库的行为。
10. 参考现有库
    Go语言社区已经有一些成熟的日志库，如 logrus, zap 等，可以参考它们的设计思路和实现细节，避免重复造轮子。
