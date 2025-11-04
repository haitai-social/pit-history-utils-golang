package pithistoryutils

import (
	"strings"
	"testing"

	"github.com/haitai-social/pit-history-utils-golang/types"
)

func TestFromJSON_BasicData(t *testing.T) {
	jsonData := `{
		"ide_name": "cursor",
		"chat_list": [
			{
				"role": "user",
				"name": "User",
				"content": "Hello",
				"is_select": true
			}
		]
	}`

	history, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	if history.Content.IDEName != types.IDENameCursor {
		t.Errorf("Expected IDE name 'cursor', got '%s'", history.Content.IDEName)
	}

	if len(history.Content.ChatList) != 1 {
		t.Errorf("Expected 1 chat, got %d", len(history.Content.ChatList))
	}
}

func TestFromJSON_VersionedData(t *testing.T) {
	jsonData := `{
		"version": "v1",
		"content": {
			"ide_name": "cursor",
			"chat_list": [
				{
					"role": "user",
					"name": "User",
					"content": "Test"
				}
			]
		}
	}`

	history, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	if history.Content.IDEName != types.IDENameCursor {
		t.Errorf("Expected IDE name 'cursor', got '%s'", history.Content.IDEName)
	}
}

func TestUnselectChatAtIndex(t *testing.T) {
	content := types.NewVibeHistoryContent()
	chat := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "Test",
		Content:  "Test content",
		IsSelect: true,
	}
	content.ChatList = append(content.ChatList, chat)

	history := NewVibeHistoryModel(content)

	err := history.UnselectChatAtIndex(0)
	if err != nil {
		t.Fatalf("UnselectChatAtIndex failed: %v", err)
	}

	if history.Content.ChatList[0].IsSelect {
		t.Error("Expected chat to be unselected")
	}
}

func TestSelectChatAtIndex(t *testing.T) {
	content := types.NewVibeHistoryContent()
	chat := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "Test",
		Content:  "Test content",
		IsSelect: false,
	}
	content.ChatList = append(content.ChatList, chat)

	history := NewVibeHistoryModel(content)

	err := history.SelectChatAtIndex(0)
	if err != nil {
		t.Fatalf("SelectChatAtIndex failed: %v", err)
	}

	if !history.Content.ChatList[0].IsSelect {
		t.Error("Expected chat to be selected")
	}
}

func TestEditNameAtIndex(t *testing.T) {
	content := types.NewVibeHistoryContent()
	chat := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "Old Name",
		Content:  "Test content",
		IsSelect: true,
	}
	content.ChatList = append(content.ChatList, chat)

	history := NewVibeHistoryModel(content)

	err := history.EditNameAtIndex(0, "New Name")
	if err != nil {
		t.Fatalf("EditNameAtIndex failed: %v", err)
	}

	if history.Content.ChatList[0].Name != "New Name" {
		t.Errorf("Expected name 'New Name', got '%s'", history.Content.ChatList[0].Name)
	}
}

func TestEditNameAtIndex_EmptyString(t *testing.T) {
	content := types.NewVibeHistoryContent()
	chat := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "Old Name",
		Content:  "Test content",
		IsSelect: true,
	}
	content.ChatList = append(content.ChatList, chat)

	history := NewVibeHistoryModel(content)

	err := history.EditNameAtIndex(0, "   ")
	if err == nil {
		t.Error("Expected error for empty string")
	}
}

func TestEditIDEName(t *testing.T) {
	content := types.NewVibeHistoryContent()
	history := NewVibeHistoryModel(content)

	history.EditIDEName(types.IDENameWinsurf)

	if history.Content.IDEName != types.IDENameWinsurf {
		t.Errorf("Expected IDE name 'winsurf', got '%s'", history.Content.IDEName)
	}
}

func TestAppendChatHistory(t *testing.T) {
	content := types.NewVibeHistoryContent()
	history := NewVibeHistoryModel(content)

	chat := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "User",
		Content:  "New chat",
		IsSelect: true,
	}

	err := history.AppendChatHistory(chat)
	if err != nil {
		t.Fatalf("AppendChatHistory failed: %v", err)
	}

	if len(history.Content.ChatList) != 1 {
		t.Errorf("Expected 1 chat, got %d", len(history.Content.ChatList))
	}
}

func TestToJSON(t *testing.T) {
	content := types.NewVibeHistoryContent()
	chat1 := &types.SingleChat{
		Role:     types.RoleUser,
		Name:     "User",
		Content:  "Hello",
		IsSelect: true,
	}
	chat2 := &types.SingleChat{
		Role:     types.RoleAssistant,
		Name:     "Assistant",
		Content:  "Hi",
		IsSelect: false,
	}
	content.ChatList = append(content.ChatList, chat1, chat2)

	history := NewVibeHistoryModel(content)

	jsonStr, err := history.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	if !strings.Contains(jsonStr, `"version": "v1"`) {
		t.Error("Expected version field in JSON")
	}

	if !strings.Contains(jsonStr, "Hello") {
		t.Error("Expected selected chat in JSON")
	}

	if strings.Contains(jsonStr, "Hi") {
		t.Error("Did not expect unselected chat in JSON")
	}
}

func TestValidateIndex_OutOfRange(t *testing.T) {
	content := types.NewVibeHistoryContent()
	history := NewVibeHistoryModel(content)

	err := history.SelectChatAtIndex(0)
	if err == nil {
		t.Error("Expected error for index out of range")
	}
}

func TestFromJSON_InvalidJSON(t *testing.T) {
	jsonData := `{invalid json}`

	_, err := FromJSON(jsonData)
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestFromJSON_MissingContent(t *testing.T) {
	jsonData := `{
		"version": "v1"
	}`

	_, err := FromJSON(jsonData)
	if err == nil {
		t.Error("Expected error for missing content")
	}
}
