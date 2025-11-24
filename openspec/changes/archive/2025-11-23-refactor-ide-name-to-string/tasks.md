## 1. 类型定义重构
- [x] 1.1 修改 `types/vibe_history_content.go` 中的 `VibeHistoryContent` 结构体，将 `IDEName IDENameEnum` 改为 `IDEName string`
- [x] 1.2 移除 `IDENameEnum` 类型定义和所有枚举常量
- [x] 1.3 移除 `NewVibeHistoryContent()` 函数中对枚举常量的引用，使用空字符串作为默认值

## 2. 验证逻辑更新
- [x] 2.1 移除 `VibeHistoryContent.Validate()` 方法中的 IDE 名称枚举验证逻辑
- [x] 2.2 更新验证方法，只保留聊天列表验证逻辑

## 3. 使用处更新
- [x] 3.1 检查并更新 `vibe_history.go` 中对 `IDENameEnum` 的引用
- [x] 3.2 检查并更新测试文件中的枚举常量引用
- [x] 3.3 检查并更新示例代码中的枚举常量引用

## 4. 清理和文档
- [x] 4.1 移除不再需要的枚举相关代码
- [x] 4.2 更新 README.md 和其他文档反映变更
- [x] 4.3 验证所有变更后的代码编译通过
