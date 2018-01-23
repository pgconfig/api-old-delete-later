package categories

import "github.com/pgconfig/api/params"

var (
	// MemoryCategory contains all parameters related
	MemoryCategory = Category{
		Name:        "memory_related",
		Description: "Memory Configuration",
		Parameters: []params.Parameter{
			params.CheckPointSegments,
		},
	}
)
