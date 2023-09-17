package entities

import (
	"sort"
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

type showTable struct {
	TableID      int       `json:"table_id"`
	SeaterAmount int       `json:"seater_amount"`
	IsAvailable  bool      `json:"available"`
	CreatedAt    time.Time `json:"create_time"`
}

func ShowTable() []showTable {
	var showT []showTable
	for tID, t := range tables {
		showT = append(showT, showTable{
			TableID:      tID,
			SeaterAmount: t.SeaterAmount,
			IsAvailable:  t.IsAvailable,
			CreatedAt:    t.CreatedAt,
		})
	}

	sort.Slice(showT, func(i, j int) bool {
		return showT[i].TableID < showT[j].TableID
	})

	return showT
}
