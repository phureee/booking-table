package entities

import (
	"time"
)

type Table struct {
	SeaterAmount int
	IsAvailable  bool
	CreatedAt    time.Time
}

// struct map[TableID]Booking
var tables map[int]Table

func SetTable(tableID int, table Table) {
	if tables == nil {
		tables = make(map[int]Table)
	}
	tables[tableID] = table
}

func GetTable() map[int]Table {
	return tables
}

func AmountTable() int {
	return len(tables)
}

func UpdateTableAvailable(tableID int, isAvailable bool) (IsSuccess bool) {
	if table, ok := tables[tableID]; ok {
		table.IsAvailable = isAvailable
		tables[tableID] = table
		return true
	}
	return false
}
