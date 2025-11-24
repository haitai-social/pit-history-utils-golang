# pit-history-utils-golang

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Golang utility library for managing and exporting Haitai Community IDE chat history fragments.

## Features

- 🔧 **Comprehensive Type Definitions**: Leverages Go's strong typing to ensure complete type safety
- 📦 **Data Validation**: Strict built-in data validation
- 🔄 **Versioning**: Supports compatible versions for history data
- 🎯 **Selection Management**: Select/unselect chat records
- ✏️ **Content Editing**: Edit chat names and IDE name
- ➕ **Add History**: Append new chat records
- 📤 **Export Data**: Export selected chat history into a standard format

## Installation

```bash
go get github.com/haitai-social/pit-history-utils-golang
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"

    pithistory "github.com/haitai-social/pit-history-utils-golang"
    "github.com/haitai-social/pit-history-utils-golang/types"
)

func main() {
    // Create a history model from a JSON string
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

    // Edit the IDE name
    history.EditIDEName("cursor")

    // Unselect the first chat record
    if err := history.UnselectChatAtIndex(0); err != nil {
        log.Fatal(err)
    }

    // Edit the name of the second chat
    if err := history.EditNameAtIndex(1, "AI Assistant"); err != nil {
        log.Fatal(err)
    }

    // Append a new chat record
    newChat := &types.SingleChat{
        Role:     types.RoleUser,
        Name:     "User",
        Content:  "Thanks for your help!",
        IsSelect: true,
    }
    if err := history.AppendChatHistory(newChat); err != nil {
        log.Fatal(err)
    }

    // Export the selected chat history
    exportedData, err := history.ToJSON()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(exportedData)
}
```

### Handling Versioned Data

```go
// Handling v1 version data
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

## API Documentation

### VibeHistoryModel

The main history management model.

#### Static Methods

##### `FromJSON(input string) (*VibeHistoryModel, error)`
Parse a JSON string and create an instance of the history model.

**Params:**
- `input string` - History data as JSON string

**Returns:** an instance of `*VibeHistoryModel` and any possible error

**Errors:**
- Returns an error if JSON parsing fails
- Returns an error if the data structure is invalid

#### Instance Methods

##### `UnselectChatAtIndex(index int) error`
Unselect the chat record at the specified index.

##### `SelectChatAtIndex(index int) error`
Select the chat record at the specified index.

##### `EditNameAtIndex(index int, newName string) error`
Edit the name of the chat record at the specified index.

##### `EditIDEName(newName string)`
Edit the IDE name.

##### `AppendChatHistory(chat *types.SingleChat) error`
Append a new chat record to the end of the list.

##### `ToJSON() (string, error)`
Export selected chat records as a v1 format JSON string.

**Return:** The exported history data (including version info and filtered chat list)

## Type Definitions

### SingleChat

Definition of a single chat record:

```go
type SingleChat struct {
    Role     RoleEnum `json:"role"`      // Role (e.g. "user", "assistant")
    Name     string   `json:"name"`      // Chat name
    Content  string   `json:"content"`   // Chat content
    IsSelect bool     `json:"is_select"` // Whether selected (used internally)
}
```

### VibeHistoryContent

Main structure for history content:

```go
type VibeHistoryContent struct {
    IDEName  string        `json:"ide_name"`  // IDE name
    ChatList []*SingleChat `json:"chat_list"` // List of chat records
}
```

### Enum Types

#### RoleEnum

```go
const (
    RoleUser      RoleEnum = "user"
    RoleAssistant RoleEnum = "assistant"
    RoleTool      RoleEnum = "tool"
)
```


## Data Validation

This library has strict built-in data validation to guarantee:

- All required fields are present
- Data types are correct
- String fields are not empty (where appropriate)
- Array structures are correct
- Enum values are valid

## Error Handling

This library returns the following error types:

- `ValidationError` - Parameter validation failed
- `IndexError` - Index out of range
- `EmptyStringError` - Empty string encountered
- Standard errors - e.g. JSON parsing failed or data structure mismatch

## Development

### Build the Project

```bash
go build ./...
```

### Run Tests

```bash
go test ./...
```

### Format Code

```bash
go fmt ./...
```

## Related Projects

- [pit-history-utils-typescript](https://github.com/haitai-social/pit-history-utils-typescript) - TypeScript version
- [Haitai Community IDE](https://github.com/haitai-social/community-ide)
- [Model Context Protocol](https://github.com/modelcontextprotocol)

## Contributing

Contributions via issue or pull request are welcome!

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to your branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**haitai-social** - [GitHub](https://github.com/haitai-social)

## Support

If you encounter any problems using this library:

1. Check the [Issues](https://github.com/haitai-social/pit-history-utils-golang/issues) page
2. Create a new issue describing your problem
3. Provide relevant code samples and error details

