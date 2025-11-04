package main

import (
	"fmt"
	"log"

	pithistory "github.com/haitai-social/pit-history-utils-golang"
	"github.com/haitai-social/pit-history-utils-golang/types"
)

func main() {
	// 示例 1: 从 JSON 字符串创建历史记录模型
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

	fmt.Println("=== 示例 1: 基本操作 ===")

	// 编辑 IDE 名称
	history.EditIDEName(types.IDENameCursor)
	fmt.Println("已设置 IDE 名称为: cursor")

	// 取消选择第一条聊天记录
	if err := history.UnselectChatAtIndex(0); err != nil {
		log.Fatal(err)
	}
	fmt.Println("已取消选择第一条聊天记录")

	// 编辑聊天名称
	if err := history.EditNameAtIndex(1, "AI Assistant"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("已修改第二条聊天记录的名称")

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
	fmt.Println("已添加新的聊天记录")

	// 导出选中的聊天历史
	exportedData, err := history.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n导出的 JSON 数据:")
	fmt.Println(exportedData)

	// 示例 2: 处理带版本的数据
	fmt.Println("\n=== 示例 2: 处理带版本的数据 ===")
	v1JsonData := `{
		"version": "v1",
		"content": {
			"ide_name": "cursor",
			"chat_list": [
				{
					"role": "user",
					"name": "User",
					"content": "Hello from v1 data",
					"is_select": true
				}
			]
		}
	}`

	historyV1, err := pithistory.FromJSON(v1JsonData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功加载 v1 版本数据")
	fmt.Printf("IDE 名称: %s\n", historyV1.Content.IDEName)
	fmt.Printf("聊天记录数量: %d\n", len(historyV1.Content.ChatList))

	// 示例 3: 创建新的历史记录
	fmt.Println("\n=== 示例 3: 创建新的历史记录 ===")
	content := types.NewVibeHistoryContent()
	content.IDEName = types.IDENameCursor

	// 添加聊天记录
	chat1 := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "开发者",
		Content:  "如何使用这个库？",
		IsSelect: true,
	}
	chat2 := &types.SingleChat{
		Role:     types.RoleAssistant,
		Name:     "助手",
		Content:  "这是一个用于管理聊天历史的库，您可以...",
		IsSelect: true,
	}

	newHistory := pithistory.NewVibeHistoryModel(content)
	newHistory.AppendChatHistory(chat1)
	newHistory.AppendChatHistory(chat2)

	fmt.Println("创建了新的历史记录")
	fmt.Printf("聊天记录数量: %d\n", len(newHistory.Content.ChatList))

	exportedNew, err := newHistory.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n新历史记录的 JSON 数据:")
	fmt.Println(exportedNew)
}

