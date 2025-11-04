# 项目结构

## 目录结构

```
pit-history-utils-golang/
├── common/                    # 通用常量和工具
│   └── version.go            # 版本常量定义
├── types/                     # 类型定义
│   ├── errors.go             # 错误类型定义
│   ├── single_chat.go        # 单条聊天记录类型
│   └── vibe_history_content.go # 历史内容结构类型
├── example/                   # 使用示例
│   └── main.go               # 示例程序
├── vibe_history.go           # 核心功能实现
├── vibe_history_test.go      # 单元测试
├── go.mod                     # Go 模块定义
├── README.md                  # 项目文档
├── CHANGELOG.md               # 变更日志
├── LICENSE                    # MIT 许可证
└── .gitignore                # Git 忽略文件配置
```

## 核心组件

### 1. `types` 包

定义了所有数据类型和错误类型：

- **single_chat.go**: 单条聊天记录的结构和验证
  - `RoleEnum`: 角色枚举（user, assistant, tool）
  - `SingleChat`: 单条聊天记录结构
  - `NewSingleChat()`: 创建默认聊天记录
  - `Validate()`: 验证数据有效性

- **vibe_history_content.go**: 历史内容的结构和验证
  - `IDENameEnum`: IDE 名称枚举
  - `VibeHistoryContent`: 历史内容结构
  - `NewVibeHistoryContent()`: 创建默认历史内容
  - `Validate()`: 验证数据有效性

- **errors.go**: 自定义错误类型
  - `ValidationError`: 验证错误
  - `IndexError`: 索引错误
  - `EmptyStringError`: 空字符串错误

### 2. `common` 包

提供通用常量：

- **version.go**: JSON 数据格式版本常量（v1）

### 3. 主包功能

**vibe_history.go** 实现核心功能：

- `VibeHistoryModel`: 历史记录模型类
  - `FromJSON(input string)`: 从 JSON 创建模型（静态方法）
  - `NewVibeHistoryModel(content)`: 创建新模型
  - `UnselectChatAtIndex(index)`: 取消选择聊天记录
  - `SelectChatAtIndex(index)`: 选择聊天记录
  - `EditNameAtIndex(index, name)`: 编辑聊天名称
  - `EditIDEName(name)`: 编辑 IDE 名称
  - `AppendChatHistory(chat)`: 添加聊天记录
  - `ToJSON()`: 导出为 JSON

## 数据流

1. **输入**: JSON 字符串
2. **解析**: `FromJSON()` 解析并验证数据
3. **操作**: 使用各种方法操作数据
4. **验证**: 自动验证所有输入
5. **输出**: `ToJSON()` 导出选中的数据

## 版本管理

- 当前版本: **v1.0.0**
- 数据格式版本: **v1**
- 向后兼容：支持带版本和不带版本的 JSON 数据

## 测试

**vibe_history_test.go** 包含完整的单元测试：

- 基本数据解析测试
- 版本化数据解析测试
- 所有操作方法的测试
- 错误处理测试
- 边界条件测试

运行测试：
```bash
go test -v ./...
```

## 示例

**example/main.go** 包含三个完整示例：

1. 基本操作示例
2. 处理带版本数据示例
3. 创建新历史记录示例

运行示例：
```bash
go run example/main.go
```

## 命名约定

项目严格遵循 Go 语言命名规范：

- 文件名：小写字母 + 下划线（如 `single_chat.go`）
- 包名：小写单数词（如 `types`, `common`）
- 导出类型/函数：大驼峰命名（如 `VibeHistoryModel`）
- 未导出函数：小驼峰命名（如 `validateIndex`）
- 常量：大驼峰命名（如 `JSONVersion`）

## 依赖

本项目无外部依赖，仅使用 Go 标准库：

- `encoding/json`: JSON 序列化/反序列化
- `fmt`: 格式化输出
- `strings`: 字符串处理

## Git 仓库

- 远程仓库: https://github.com/haitai-social/pit-history-utils-golang.git
- 初始版本: v1.0.0
- 分支: main

