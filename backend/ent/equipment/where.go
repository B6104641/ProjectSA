// Code generated by entc, DO NOT EDIT.

package equipment

import (
	"github.com/B6104641/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EQUIPMENTNAME applies equality check predicate on the "EQUIPMENT_NAME" field. It's identical to EQUIPMENTNAMEEQ.
func EQUIPMENTNAME(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTPRICE applies equality check predicate on the "EQUIPMENT_PRICE" field. It's identical to EQUIPMENTPRICEEQ.
func EQUIPMENTPRICE(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// EQUIPMENTNAMEEQ applies the EQ predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEEQ(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMENEQ applies the NEQ predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMENEQ(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEIn applies the In predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEIn(vs ...string) predicate.Equipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEQUIPMENTNAME), v...))
	})
}

// EQUIPMENTNAMENotIn applies the NotIn predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMENotIn(vs ...string) predicate.Equipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEQUIPMENTNAME), v...))
	})
}

// EQUIPMENTNAMEGT applies the GT predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEGT(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEGTE applies the GTE predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEGTE(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMELT applies the LT predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMELT(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMELTE applies the LTE predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMELTE(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEContains applies the Contains predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEContains(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEHasPrefix applies the HasPrefix predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEHasSuffix applies the HasSuffix predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEEqualFold applies the EqualFold predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTNAMEContainsFold applies the ContainsFold predicate on the "EQUIPMENT_NAME" field.
func EQUIPMENTNAMEContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEQUIPMENTNAME), v))
	})
}

// EQUIPMENTPRICEEQ applies the EQ predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICEEQ(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// EQUIPMENTPRICENEQ applies the NEQ predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICENEQ(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// EQUIPMENTPRICEIn applies the In predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICEIn(vs ...int) predicate.Equipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEQUIPMENTPRICE), v...))
	})
}

// EQUIPMENTPRICENotIn applies the NotIn predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICENotIn(vs ...int) predicate.Equipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEQUIPMENTPRICE), v...))
	})
}

// EQUIPMENTPRICEGT applies the GT predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICEGT(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// EQUIPMENTPRICEGTE applies the GTE predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICEGTE(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// EQUIPMENTPRICELT applies the LT predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICELT(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// EQUIPMENTPRICELTE applies the LTE predicate on the "EQUIPMENT_PRICE" field.
func EQUIPMENTPRICELTE(v int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEQUIPMENTPRICE), v))
	})
}

// HasOrderEquipmentt applies the HasEdge predicate on the "order_equipmentt" edge.
func HasOrderEquipmentt() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderEquipmenttTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderEquipmenttTable, OrderEquipmenttColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderEquipmenttWith applies the HasEdge predicate on the "order_equipmentt" edge with a given conditions (other predicates).
func HasOrderEquipmenttWith(preds ...predicate.Order) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderEquipmenttInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderEquipmenttTable, OrderEquipmenttColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		p(s.Not())
	})
}
