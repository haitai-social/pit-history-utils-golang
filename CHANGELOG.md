# Changelog

本文档记录了项目的所有重要变更。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
本项目遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [1.0.0] - 2025-11-04

### 新增
- 初始版本发布
- `VibeHistoryModel` 核心类，用于管理历史记录
- `FromJSON` 静态方法，支持从 JSON 字符串创建模型
- `UnselectChatAtIndex` 方法，取消选择指定索引的聊天记录
- `SelectChatAtIndex` 方法，选择指定索引的聊天记录
- `EditNameAtIndex` 方法，编辑聊天记录名称
- `EditIDEName` 方法，编辑 IDE 名称
- `AppendChatHistory` 方法，添加新的聊天记录
- `ToJSON` 方法，导出选中的聊天历史为 JSON
- 完整的类型定义：
  - `SingleChat` - 单条聊天记录
  - `VibeHistoryContent` - 历史内容结构
  - `RoleEnum` - 角色枚举（user, assistant, tool）
  - IDE 名称字段现在使用 string 类型，支持任意 IDE 名称
- 错误类型：
  - `ValidationError` - 验证错误
  - `IndexError` - 索引错误
  - `EmptyStringError` - 空字符串错误
- 版本化数据格式支持（v1）
- 完整的数据验证机制
- 使用示例（example/main.go）
- 单元测试覆盖核心功能
- MIT 许可证
- 完整的 README 文档

### 特性
- 支持从带版本和不带版本的 JSON 数据创建模型（向后兼容）
- 严格的数据验证，确保数据完整性
- 导出时自动过滤未选中的聊天记录
- 清晰的错误处理和错误类型

[1.0.0]: https://github.com/haitai-social/pit-history-utils-golang/releases/tag/v1.0.0

