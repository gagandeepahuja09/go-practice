package inmemorydb

import "fmt"

var (
	errTableAlreadyExists                                   = "table: '%s' already exists"
	errTableNotFound                                        = "table: '%s' not found"
	errColumnValuesCountProvidedNotMatchingWithColumnsCount = "number of column values provided: '%d' doesn't match with the number of columns: '%d"
	errRowIdNotFound                                        = "row_id: '%d' not found"
)

var database map[string]table

type table struct {
	columnNames []string
	rows        map[int][]string
	nextRowId   int
}

type columnIdentifier struct {
	tableName  string
	columnName string
}

var tableNameColumnNameVsRowsMap map[columnIdentifier][]int

func CreateTable(tableName string, columnNames []string) error {
	if _, ok := database[tableName]; ok {
		return fmt.Errorf(errTableAlreadyExists, tableName)
	}
	database[tableName] = table{
		columnNames: columnNames,
	}
	return nil
}

func appendRowInAllIndexes(table table, tableName string, insertedRowId int) {
	for _, columnName := range table.columnNames {
		columnIdentifier := columnIdentifier{
			tableName:  tableName,
			columnName: columnName,
		}

		if rowIds, ok := tableNameColumnNameVsRowsMap[columnIdentifier]; ok {
			tableNameColumnNameVsRowsMap[columnIdentifier] = append(rowIds, insertedRowId)
		}
	}
}

func removeRowInAllIndexes(table table, tableName string, deletedRowId int) {
	for _, columnName := range table.columnNames {
		columnIdentifier := columnIdentifier{
			tableName:  tableName,
			columnName: columnName,
		}

		if rowIds, ok := tableNameColumnNameVsRowsMap[columnIdentifier]; ok {
			updatedRowIds := []int{}
			for _, currentRowId := range rowIds {
				if currentRowId == deletedRowId {
					continue
				}
				// this is a slow: O(n) operation to reconstruct entire array
				// : based on tradeoffs in search, can be optimised
				updatedRowIds = append(updatedRowIds, currentRowId)
			}
			tableNameColumnNameVsRowsMap[columnIdentifier] = updatedRowIds
		}
	}
}

func InsertRow(tableName string, columnValues []string) error {
	table, ok := database[tableName]
	if !ok {
		return fmt.Errorf(errTableNotFound, tableName)
	}
	if len(columnValues) != len(table.columnNames) {
		return fmt.Errorf(errColumnValuesCountProvidedNotMatchingWithColumnsCount, len(columnValues), len(table.columnNames))
	}
	table.rows[table.nextRowId] = columnValues
	appendRowInAllIndexes(table, tableName, table.nextRowId)
	table.nextRowId++
	return nil
}

func DeleteRow(tableName string, rowId int) error {
	table, ok := database[tableName]
	if !ok {
		return fmt.Errorf(errTableNotFound, tableName)
	}
	if rowId < 0 || rowId >= table.nextRowId {
		return fmt.Errorf(errRowIdNotFound, rowId)
	}
	delete(table.rows, rowId)
	appendRowInAllIndexes(table, tableName, table.nextRowId)
	return nil
}
