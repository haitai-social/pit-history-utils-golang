# pit-history-utils-golang

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个用于管理和导出 Haitai 社区 IDE 历史记录片段的 Go 工具库。

## 功能特性

- 🔧 **完整类型定义**: 基于 Go 的强类型系统，提供完整的类型定义
- 📦 **数据验证**: 内置严格的数据验证机制
- 🔄 **版本管理**: 支持历史数据的版本兼容性
- 🎯 **选择管理**: 支持选择/取消选择聊天记录
- ✏️ **内容编辑**: 支持编辑聊天名称和 IDE 名称
- ➕ **历史添加**: 支持添加新的聊天历史记录
- 📤 **数据导出**: 支持导出选中的聊天历史为标准格式

## 安装

```bash
go get github.com/haitai-social/pit-history-utils-golang
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "log"

    pithistory "github.com/haitai-social/pit-history-utils-golang"
    "github.com/haitai-social/pit-history-utils-golang/types"
)

func main() {
    // 从 JSON 字符串创建历史记录模型
    jsonData := `{
        "ide_name": "cursor",
        "chat_list": [
            {
                "role": "user",
                "name": "User",
                "content": "Hello",
                "is_select": true
            },
            {
                "role": "assistant",
                "name": "Assistant",
                "content": "Hello! How can I help you?",
                "is_select": true
            }
        ]
    }`

    history, err := pithistory.FromJSON(jsonData)
    if err != nil {
        log.Fatal(err)
    }

    // 编辑 IDE 名称
    history.EditIDEName(types.IDENameCursor)

    // 取消选择第一条聊天记录
    if err := history.UnselectChatAtIndex(0); err != nil {
        log.Fatal(err)
    }

    // 编辑聊天名称
    if err := history.EditNameAtIndex(1, "AI Assistant"); err != nil {
        log.Fatal(err)
    }

    // 添加新的聊天记录
    newChat := &types.SingleChat{
        Role:     types.RoleUser,
        Name:     "User",
        Content:  "Thanks for your help!",
        IsSelect: true,
    }
    if err := history.AppendChatHistory(newChat); err != nil {
        log.Fatal(err)
    }

    // 导出选中的聊天历史
    exportedData, err := history.ToJSON()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(exportedData)
}
```

### 处理带版本的数据

```go
// 处理 v1 版本数据
v1JsonData := `{
    "version": "v1",
    "content": {
        "ide_name": "cursor",
        "chat_list": [
            {
                "role": "user",
                "name": "User",
                "content": "Hello"
            }
        ]
    }
}`

history, err := pithistory.FromJSON(v1JsonData)
if err != nil {
    log.Fatal(err)
}
```

## API 文档

### VibeHistoryModel

主要的历史记录管理类。

#### 静态方法

##### `FromJSON(input string) (*VibeHistoryModel, error)`
解析 JSON 字符串并创建历史记录模型实例。

**参数:**
- `input string` - JSON 格式的历史数据字符串

**返回值:** `*VibeHistoryModel` 实例和可能的错误

**错误:**
- JSON 解析失败时返回错误
- 数据结构不正确时返回错误

#### 实例方法

##### `UnselectChatAtIndex(index int) error`
取消选择指定索引的聊天记录。

##### `SelectChatAtIndex(index int) error`
选择指定索引的聊天记录。

##### `EditNameAtIndex(index int, newName string) error`
编辑指定索引的聊天记录名称。

##### `EditIDEName(newName types.IDENameEnum)`
编辑 IDE 名称。

##### `AppendChatHistory(chat *types.SingleChat) error`
在历史列表末尾添加新的聊天记录。

##### `ToJSON() (string, error)`
导出选中的聊天历史为 v1 格式 JSON 数据。

**返回值:** 导出的历史数据（包含版本信息和过滤后的聊天列表）

## 类型定义

### SingleChat
单条聊天记录的类型定义：

```go
type SingleChat struct {
    Role     RoleEnum `json:"role"`       // 角色（如 "user", "assistant"）
    Name     string   `json:"name"`       // 聊天名称
    Content  string   `json:"content"`    // 聊天内容
    IsSelect bool     `json:"is_select"`  // 是否选中（仅供内部使用）
}
```

### VibeHistoryContent
历史内容的主要结构：

```go
type VibeHistoryContent struct {
    IDEName  IDENameEnum   `json:"ide_name"`   // IDE 名称
    ChatList []*SingleChat `json:"chat_list"`  // 聊天记录列表
}
```

### 枚举类型

#### RoleEnum
```go
const (
    RoleUser      RoleEnum = "user"
    RoleAssistant RoleEnum = "assistant"
    RoleTool      RoleEnum = "tool"
)
```

#### IDENameEnum
```go
const (
    IDENameCursor     IDENameEnum = "cursor"
    IDENameClaudeCode IDENameEnum = "claude code"
    IDENameTrea       IDENameEnum = "trea"
    IDENameWinsurf    IDENameEnum = "winsurf"
    IDENameCodex      IDENameEnum = "codex"
)
```

## 数据验证

本库内置了严格的数据验证机制，确保：

- 所有必需字段都存在
- 数据类型正确
- 字符串字段非空（适当时）
- 数组结构正确
- 枚举值有效

## 错误处理

本库返回以下类型的错误：

- `ValidationError` - 参数验证失败
- `IndexError` - 索引超出范围
- `EmptyStringError` - 字符串为空
- 标准错误 - JSON 解析失败或数据结构不匹配

## 开发

### 构建项目

```bash
go build ./...
```

### 运行测试

```bash
go test ./...
```

### 格式化代码

```bash
go fmt ./...
```

## 相关项目

- [pit-history-utils-typescript](https://github.com/haitai-social/pit-history-utils-typescript) - TypeScript 版本
- [Haitai Community IDE](https://github.com/haitai-social/community-ide)
- [Model Context Protocol](https://github.com/modelcontextprotocol)

## 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 作者

**haitai-social** - [GitHub](https://github.com/haitai-social)

## 支持

如果您在使用本库时遇到任何问题：

1. 查看 [Issues](https://github.com/haitai-social/pit-history-utils-golang/issues) 页面
2. 创建新 Issue 描述您的问题
3. 提供相关的代码示例和错误信息

