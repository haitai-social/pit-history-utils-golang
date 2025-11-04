package types

// IDENameEnum 定义 IDE 名称枚举
type IDENameEnum string

const (
	IDENameCursor     IDENameEnum = "cursor"
	IDENameClaudeCode IDENameEnum = "claude code"
	IDENameTrea       IDENameEnum = "trea"
	IDENameWinsurf    IDENameEnum = "winsurf"
	IDENameCodex      IDENameEnum = "codex"
)

// VibeHistoryContent 历史记录内容结构
type VibeHistoryContent struct {
	IDEName  IDENameEnum   `json:"ide_name"`
	ChatList []*SingleChat `json:"chat_list"`
}

// NewVibeHistoryContent 创建一个新的历史记录内容，使用默认值
func NewVibeHistoryContent() *VibeHistoryContent {
	return &VibeHistoryContent{
		IDEName:  IDENameCursor,
		ChatList: make([]*SingleChat, 0),
	}
}

// Validate 验证 VibeHistoryContent 的有效性
func (v *VibeHistoryContent) Validate() error {
	// 验证 IDE 名称是否有效
	switch v.IDEName {
	case IDENameCursor, IDENameClaudeCode, IDENameTrea, IDENameWinsurf, IDENameCodex:
		// 有效 IDE 名称
	default:
		return &ValidationError{Field: "ide_name", Message: "invalid ide_name value"}
	}

	// 验证聊天列表中的每一项
	for i, chat := range v.ChatList {
		if err := chat.Validate(); err != nil {
			return &ValidationError{
				Field:   "chat_list",
				Message: "invalid chat at index " + string(rune(i)) + ": " + err.Error(),
			}
		}
	}

	return nil
}
