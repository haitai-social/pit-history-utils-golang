// Package types 定义了 pit-history-utils 库使用的所有数据类型和错误类型
package types

import "fmt"

// ValidationError 验证错误
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// IndexError 索引错误
type IndexError struct {
	Index   int
	Min     int
	Max     int
	Message string
}

func (e *IndexError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("index %d is out of range (%d ~ %d)", e.Index, e.Min, e.Max)
}

// EmptyStringError 空字符串错误
type EmptyStringError struct {
	FieldName string
}

func (e *EmptyStringError) Error() string {
	return fmt.Sprintf("%s must be a non-empty string", e.FieldName)
}
