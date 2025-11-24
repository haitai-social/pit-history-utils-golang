package types

// VibeHistoryContent 历史记录内容结构
type VibeHistoryContent struct {
	IDEName  string        `json:"ide_name"`
	ChatList []*SingleChat `json:"chat_list"`
}

// NewVibeHistoryContent 创建一个新的历史记录内容，使用默认值
func NewVibeHistoryContent() *VibeHistoryContent {
	return &VibeHistoryContent{
		IDEName:  "cursor",
		ChatList: make([]*SingleChat, 0),
	}
}

// Validate 验证 VibeHistoryContent 的有效性
func (v *VibeHistoryContent) Validate() error {
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
