* Database
* Multiple tables
* Table multiple rows and columns

* Table  
    columnNames []string
    nextRowId int (auto-increment)
    Rows map[int]Row (key: rowId, value: rowData)

* Row []string

* Create Index
    * Efficiency in finding rowIds with column value.
    * map[columnValue][]Row 
    