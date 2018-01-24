package categories

import (
	"github.com/pgconfig/api/params"
)

// Category contain a group of Parameters
type Category struct {
	Name        string             `json:"category"`
	Description string             `json:"description"`
	Parameters  []params.Parameter `json:"parameters"`
}

// Compute calculates each parameter value
func (c *Category) Compute(args params.Input) (err error) {
	for i := 0; i < len(c.Parameters); i++ {
		err = c.Parameters[i].Compute(args)

		if err != nil {
			break
		}
	}

	return
}
