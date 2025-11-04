package types

// RoleEnum 定义聊天角色枚举
type RoleEnum string

const (
	RoleUser      RoleEnum = "user"
	RoleAssistant RoleEnum = "assistant"
	RoleTool      RoleEnum = "tool"
)

// SingleChat 单条聊天记录
type SingleChat struct {
	Role     RoleEnum `json:"role"`
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	IsSelect bool     `json:"is_select"`
}

// NewSingleChat 创建一个新的单条聊天记录，使用默认值
func NewSingleChat() *SingleChat {
	return &SingleChat{
		Role:     RoleUser,
		Name:     "",
		Content:  "",
		IsSelect: true,
	}
}

// Validate 验证 SingleChat 的有效性
func (s *SingleChat) Validate() error {
	// 验证角色是否有效
	switch s.Role {
	case RoleUser, RoleAssistant, RoleTool:
		// 有效角色
	default:
		return &ValidationError{Field: "role", Message: "invalid role value"}
	}
	return nil
}
