package plants

import (
	"time"
)

var plants map[string]Plant

func init() {
	plants = make(map[string]Plant)
}

type Plant struct {
	PlantId   string
	CreatedAt time.Time
}

func Add(p Plant) (err error, plantId string) {
	plants[p.PlantId] = p
	return nil, p.PlantId
}

func Update(p Plant) {
	plants[p.PlantId] = p
}
