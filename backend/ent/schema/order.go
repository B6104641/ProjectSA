package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{

		field.Int("AMOUNT"),
		field.Int("PRICE"),
		field.Time("ADDED_TIME"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("order_equipmentt").Unique(),
		edge.From("company", Company.Type).Ref("order_companyy").Unique(),
		edge.From("user", User.Type).Ref("order_userr").Unique(),
	}
}
