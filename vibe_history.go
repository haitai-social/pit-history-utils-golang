package pithistoryutils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/haitai-social/pit-history-utils-golang/common"
	"github.com/haitai-social/pit-history-utils-golang/types"
)

// VibeHistoryModel 历史记录模型
type VibeHistoryModel struct {
	Content *types.VibeHistoryContent `json:"content"`
}

// versionedData 带版本信息的数据结构
type versionedData struct {
	Version string                    `json:"version"`
	Content *types.VibeHistoryContent `json:"content"`
}

// exportedSingleChat 导出的单条聊天记录（不包含 is_select 字段）
type exportedSingleChat struct {
	Role    types.RoleEnum `json:"role"`
	Name    string         `json:"name"`
	Content string         `json:"content"`
}

// exportedVibeHistoryV1 导出的 v1 版本历史数据
type exportedVibeHistoryV1 struct {
	Version string `json:"version"`
	Content struct {
		IDEName  types.IDENameEnum     `json:"ide_name"`
		ChatList []*exportedSingleChat `json:"chat_list"`
	} `json:"content"`
}

// FromJSON 从 JSON 字符串创建 VibeHistoryModel
func FromJSON(input string) (*VibeHistoryModel, error) {
	var rawData map[string]interface{}

	// 解析 JSON
	if err := json.Unmarshal([]byte(input), &rawData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// 检查是否包含版本信息
	if version, ok := rawData["version"].(string); ok && version == common.JSONVersion {
		// 处理带版本的数据
		var vData versionedData
		if err := json.Unmarshal([]byte(input), &vData); err != nil {
			return nil, fmt.Errorf("failed to parse versioned data: %w", err)
		}
		if vData.Content == nil {
			return nil, fmt.Errorf("missing 'content' property in v1 history data")
		}
		// 验证数据
		if err := vData.Content.Validate(); err != nil {
			return nil, err
		}
		return &VibeHistoryModel{Content: vData.Content}, nil
	}

	// 处理不带版本的数据（向后兼容）
	var content types.VibeHistoryContent
	if err := json.Unmarshal([]byte(input), &content); err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}
	// 验证数据
	if err := content.Validate(); err != nil {
		return nil, err
	}
	return &VibeHistoryModel{Content: &content}, nil
}

// NewVibeHistoryModel 创建一个新的 VibeHistoryModel
func NewVibeHistoryModel(content *types.VibeHistoryContent) *VibeHistoryModel {
	return &VibeHistoryModel{Content: content}
}

// validateIndex 验证索引是否有效
func (v *VibeHistoryModel) validateIndex(index int) error {
	if index < 0 || index >= len(v.Content.ChatList) {
		maxIndex := len(v.Content.ChatList) - 1
		return &types.IndexError{
			Index: index,
			Min:   0,
			Max:   maxIndex,
		}
	}
	return nil
}

// validateNonEmptyString 验证字符串是否非空
func (v *VibeHistoryModel) validateNonEmptyString(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return &types.EmptyStringError{FieldName: fieldName}
	}
	return nil
}

// UnselectChatAtIndex 取消选择指定索引的聊天记录
func (v *VibeHistoryModel) UnselectChatAtIndex(index int) error {
	if err := v.validateIndex(index); err != nil {
		return err
	}
	v.Content.ChatList[index].IsSelect = false
	return nil
}

// SelectChatAtIndex 选择指定索引的聊天记录
func (v *VibeHistoryModel) SelectChatAtIndex(index int) error {
	if err := v.validateIndex(index); err != nil {
		return err
	}
	v.Content.ChatList[index].IsSelect = true
	return nil
}

// EditNameAtIndex 编辑指定索引的聊天名称
func (v *VibeHistoryModel) EditNameAtIndex(index int, newName string) error {
	if err := v.validateNonEmptyString(newName, "newName"); err != nil {
		return err
	}
	if err := v.validateIndex(index); err != nil {
		return err
	}
	v.Content.ChatList[index].Name = newName
	return nil
}

// EditIDEName 编辑 IDE 名称
func (v *VibeHistoryModel) EditIDEName(newName types.IDENameEnum) {
	v.Content.IDEName = newName
}

// AppendChatHistory 添加新的聊天记录
func (v *VibeHistoryModel) AppendChatHistory(chat *types.SingleChat) error {
	// 验证聊天记录
	if err := chat.Validate(); err != nil {
		return err
	}
	v.Content.ChatList = append(v.Content.ChatList, chat)
	return nil
}

// ToJSON 导出为 JSON 字符串（仅导出选中的聊天记录）
func (v *VibeHistoryModel) ToJSON() (string, error) {
	// 构建导出对象
	exported := exportedVibeHistoryV1{
		Version: common.JSONVersion,
	}
	exported.Content.IDEName = v.Content.IDEName

	// 筛选选中的聊天记录并转换格式
	selectedChats := make([]*exportedSingleChat, 0)
	for _, chat := range v.Content.ChatList {
		if chat.IsSelect {
			selectedChats = append(selectedChats, &exportedSingleChat{
				Role:    chat.Role,
				Name:    chat.Name,
				Content: chat.Content,
			})
		}
	}
	exported.Content.ChatList = selectedChats

	// 转换为 JSON
	jsonBytes, err := json.MarshalIndent(exported, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	return string(jsonBytes), nil
}
