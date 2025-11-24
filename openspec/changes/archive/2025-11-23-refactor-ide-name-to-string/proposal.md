# Change: 重构 IDEName 为直接 string 类型

## Why
当前 `IDEName` 字段使用枚举类型 `IDENameEnum`，限制了 IDE 名称的灵活性。当需要支持新的 IDE 或自定义 IDE 名称时，必须修改枚举定义和验证逻辑。这种设计不够灵活，增加了维护成本。

将 `IDEName` 直接定义为 `string` 类型可以提供更大的灵活性，允许任意有效的 IDE 名称，而不需要修改代码。

## What Changes
- 将 `types.IDENameEnum` 类型重构为直接使用 `string`
- 移除 `IDENameEnum` 枚举常量定义
- 移除 `VibeHistoryContent.Validate()` 方法中的枚举验证逻辑
- 更新相关结构体和使用该字段的代码

## Impact
- Affected specs: types capability (VibeHistoryContent)
- Affected code: `types/vibe_history_content.go`, 使用该类型的其他文件
- **BREAKING**: API 行为变更 - 不再限制 IDE 名称为预定义枚举值
- 向后兼容: 现有有效枚举值仍然有效，但不再强制限制
