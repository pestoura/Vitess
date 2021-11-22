/*
   Copyright 2016 GitHub Inc.
	 See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package vrepl

import (
	"fmt"
	"strings"
)

// expandedDataTypes maps some known and difficult-to-compute by INFORMATION_SCHEMA data types which expand other data types.
// For example, in "date:datetime", datetime expands date because it has more precision. In "timestamp:date" date expands timestamp
// because it can contain years not covered by timestamp.
var expandedDataTypes = map[string]bool{
	"time:datetime":      true,
	"date:datetime":      true,
	"timestamp:datetime": true,
	"time:timestamp":     true,
	"date:timestamp":     true,
	"timestamp:date":     true,
}

// GetSharedColumns returns the intersection of two lists of columns in same order as the first list
func GetSharedColumns(
	sourceColumns, targetColumns *ColumnList,
	sourceVirtualColumns, targetVirtualColumns *ColumnList,
	parser *AlterTableParser,
) (
	sourceSharedColumns *ColumnList,
	targetSharedColumns *ColumnList,
	droppedSourceNonGeneratedColumns *ColumnList,
	sharedColumnsMap map[string]string,
) {
	sharedColumnNames := []string{}
	droppedSourceNonGeneratedColumnsNames := []string{}
	for _, sourceColumn := range sourceColumns.Names() {
		isSharedColumn := false
		isVirtualColumnOnSource := false
		for _, targetColumn := range targetColumns.Names() {
			if strings.EqualFold(sourceColumn, targetColumn) {
				// both tables have this column. Good start.
				isSharedColumn = true
				break
			}
			if strings.EqualFold(parser.columnRenameMap[sourceColumn], targetColumn) {
				// column in source is renamed in target
				isSharedColumn = true
				break
			}
		}
		for droppedColumn := range parser.DroppedColumnsMap() {
			if strings.EqualFold(sourceColumn, droppedColumn) {
				isSharedColumn = false
				break
			}
		}
		for _, virtualColumn := range sourceVirtualColumns.Names() {
			// virtual/generated columns on source are silently skipped
			if strings.EqualFold(sourceColumn, virtualColumn) {
				isSharedColumn = false
				isVirtualColumnOnSource = true
			}
		}
		for _, virtualColumn := range targetVirtualColumns.Names() {
			// virtual/generated columns on target are silently skipped
			if strings.EqualFold(sourceColumn, virtualColumn) {
				isSharedColumn = false
			}
		}
		if isSharedColumn {
			sharedColumnNames = append(sharedColumnNames, sourceColumn)
		} else if !isVirtualColumnOnSource {
			droppedSourceNonGeneratedColumnsNames = append(droppedSourceNonGeneratedColumnsNames, sourceColumn)
		}
	}
	sharedColumnsMap = map[string]string{}
	for _, columnName := range sharedColumnNames {
		if mapped, ok := parser.columnRenameMap[columnName]; ok {
			sharedColumnsMap[columnName] = mapped
		} else {
			sharedColumnsMap[columnName] = columnName
		}
	}
	mappedSharedColumnNames := []string{}
	for _, columnName := range sharedColumnNames {
		mappedSharedColumnNames = append(mappedSharedColumnNames, sharedColumnsMap[columnName])
	}
	return NewColumnList(sharedColumnNames), NewColumnList(mappedSharedColumnNames), NewColumnList(droppedSourceNonGeneratedColumnsNames), sharedColumnsMap
}

// isExpandedColumn sees if target column has any value set/range that is impossible in source column. See GetExpandedColumns comment for examples
func isExpandedColumn(sourceColumn *Column, targetColumn *Column) (bool, string) {
	if targetColumn.IsNullable && !sourceColumn.IsNullable {
		return true, "target is NULL-able, source is not"
	}
	if targetColumn.CharacterMaximumLength > sourceColumn.CharacterMaximumLength {
		return true, "higher CHARACTER_MAXIMUM_LENGTH"
	}
	if targetColumn.NumericPrecision > sourceColumn.NumericPrecision {
		return true, "higher NUMERIC_PRECISION"
	}
	if targetColumn.NumericScale > sourceColumn.NumericScale {
		return true, "higher NUMERIC_SCALE"
	}
	if targetColumn.DateTimePrecision > sourceColumn.DateTimePrecision {
		return true, "higher DATETIME_PRECISION"
	}
	if sourceColumn.IsNumeric() && targetColumn.IsNumeric() {
		if sourceColumn.IsUnsigned && !targetColumn.IsUnsigned {
			return true, "source is unsigned, target is signed"
		}
		if sourceColumn.NumericPrecision <= targetColumn.NumericPrecision && !sourceColumn.IsUnsigned && targetColumn.IsUnsigned {
			// e.g. INT SIGNED => INT UNSIGNED, INT SIGNED = BIGINT UNSIGNED
			return true, "target unsgined value exeeds source unsigned value"
		}
		if targetColumn.IsFloatingPoint() && !sourceColumn.IsFloatingPoint() {
			return true, "target is floating point, source is not"
		}
	}
	if expandedDataTypes[fmt.Sprintf("%s:%s", sourceColumn.DataType, targetColumn.DataType)] {
		return true, "target is expanded data type of source"
	}
	if sourceColumn.Charset != targetColumn.Charset {
		if targetColumn.Charset == "utf8mb4" {
			return true, "expand character set to utf8mb4"
		}
		if targetColumn.Charset == "utf8" && sourceColumn.Charset != "utf8mb4" {
			return true, "expand character set to utf8"
		}
	}
	return false, ""
}

// GetExpandedColumnNames is given source and target shared columns, and returns the list of columns whose data type is expanded.
// An expanded data type is one where the target can have a value which the source does not. Examples:
// - any NOT NULL to NULLable (a NULL in the target cannot appear on source)
// - INT -> BIGINT (obvious)
// - BIGINT UNSIGNED -> INT SIGNED (negative values)
// - TIMESTAMP -> TIMESTAMP(3)
// etc.
func GetExpandedColumnNames(
	sourceSharedColumns *ColumnList,
	targetSharedColumns *ColumnList,
) (
	expandedColumnNames []string,
) {
	for i := range sourceSharedColumns.Columns() {
		// source and target columns assumed to be mapped 1:1, same length
		sourceColumn := sourceSharedColumns.Columns()[i]
		targetColumn := targetSharedColumns.Columns()[i]

		if isExpanded, _ := isExpandedColumn(&sourceColumn, &targetColumn); isExpanded {
			expandedColumnNames = append(expandedColumnNames, sourceColumn.Name)
		}
	}
	return expandedColumnNames
}

// GetNoDefaultColumnNames returns names of columns which have no default value, out of given list of columns
func GetNoDefaultColumnNames(columns *ColumnList) (names []string) {
	names = []string{}
	for _, col := range columns.Columns() {
		if !col.HasDefault() {
			names = append(names, col.Name)
		}
	}
	return names
}
