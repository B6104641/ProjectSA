package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{

		field.String("EQUIPMENT_NAME"),
		field.Int("EQUIPMENT_PRICE"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_equipmentt", Order.Type),
	}
}
