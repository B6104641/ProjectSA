// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/B6104641/app/ent/company"
	"github.com/B6104641/app/ent/equipment"
	"github.com/B6104641/app/ent/order"
	"github.com/B6104641/app/ent/predicate"
	"github.com/B6104641/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks      []Hook
	mutation   *OrderMutation
	predicates []predicate.Order
}

// Where adds a new predicate for the builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.predicates = append(ou.predicates, ps...)
	return ou
}

// SetAMOUNT sets the AMOUNT field.
func (ou *OrderUpdate) SetAMOUNT(i int) *OrderUpdate {
	ou.mutation.ResetAMOUNT()
	ou.mutation.SetAMOUNT(i)
	return ou
}

// AddAMOUNT adds i to AMOUNT.
func (ou *OrderUpdate) AddAMOUNT(i int) *OrderUpdate {
	ou.mutation.AddAMOUNT(i)
	return ou
}

// SetPRICE sets the PRICE field.
func (ou *OrderUpdate) SetPRICE(i int) *OrderUpdate {
	ou.mutation.ResetPRICE()
	ou.mutation.SetPRICE(i)
	return ou
}

// AddPRICE adds i to PRICE.
func (ou *OrderUpdate) AddPRICE(i int) *OrderUpdate {
	ou.mutation.AddPRICE(i)
	return ou
}

// SetADDEDTIME sets the ADDED_TIME field.
func (ou *OrderUpdate) SetADDEDTIME(t time.Time) *OrderUpdate {
	ou.mutation.SetADDEDTIME(t)
	return ou
}

// SetEquipmentID sets the equipment edge to Equipment by id.
func (ou *OrderUpdate) SetEquipmentID(id int) *OrderUpdate {
	ou.mutation.SetEquipmentID(id)
	return ou
}

// SetNillableEquipmentID sets the equipment edge to Equipment by id if the given value is not nil.
func (ou *OrderUpdate) SetNillableEquipmentID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetEquipmentID(*id)
	}
	return ou
}

// SetEquipment sets the equipment edge to Equipment.
func (ou *OrderUpdate) SetEquipment(e *Equipment) *OrderUpdate {
	return ou.SetEquipmentID(e.ID)
}

// SetCompanyID sets the company edge to Company by id.
func (ou *OrderUpdate) SetCompanyID(id int) *OrderUpdate {
	ou.mutation.SetCompanyID(id)
	return ou
}

// SetNillableCompanyID sets the company edge to Company by id if the given value is not nil.
func (ou *OrderUpdate) SetNillableCompanyID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetCompanyID(*id)
	}
	return ou
}

// SetCompany sets the company edge to Company.
func (ou *OrderUpdate) SetCompany(c *Company) *OrderUpdate {
	return ou.SetCompanyID(c.ID)
}

// SetUserID sets the user edge to User by id.
func (ou *OrderUpdate) SetUserID(id int) *OrderUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetUserID(*id)
	}
	return ou
}

// SetUser sets the user edge to User.
func (ou *OrderUpdate) SetUser(u *User) *OrderUpdate {
	return ou.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearEquipment clears the equipment edge to Equipment.
func (ou *OrderUpdate) ClearEquipment() *OrderUpdate {
	ou.mutation.ClearEquipment()
	return ou
}

// ClearCompany clears the company edge to Company.
func (ou *OrderUpdate) ClearCompany() *OrderUpdate {
	ou.mutation.ClearCompany()
	return ou
}

// ClearUser clears the user edge to User.
func (ou *OrderUpdate) ClearUser() *OrderUpdate {
	ou.mutation.ClearUser()
	return ou
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.AMOUNT(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAMOUNT,
		})
	}
	if value, ok := ou.mutation.AddedAMOUNT(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAMOUNT,
		})
	}
	if value, ok := ou.mutation.PRICE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldPRICE,
		})
	}
	if value, ok := ou.mutation.AddedPRICE(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldPRICE,
		})
	}
	if value, ok := ou.mutation.ADDEDTIME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldADDEDTIME,
		})
	}
	if ou.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.EquipmentTable,
			Columns: []string{order.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.EquipmentTable,
			Columns: []string{order.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CompanyTable,
			Columns: []string{order.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CompanyTable,
			Columns: []string{order.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// SetAMOUNT sets the AMOUNT field.
func (ouo *OrderUpdateOne) SetAMOUNT(i int) *OrderUpdateOne {
	ouo.mutation.ResetAMOUNT()
	ouo.mutation.SetAMOUNT(i)
	return ouo
}

// AddAMOUNT adds i to AMOUNT.
func (ouo *OrderUpdateOne) AddAMOUNT(i int) *OrderUpdateOne {
	ouo.mutation.AddAMOUNT(i)
	return ouo
}

// SetPRICE sets the PRICE field.
func (ouo *OrderUpdateOne) SetPRICE(i int) *OrderUpdateOne {
	ouo.mutation.ResetPRICE()
	ouo.mutation.SetPRICE(i)
	return ouo
}

// AddPRICE adds i to PRICE.
func (ouo *OrderUpdateOne) AddPRICE(i int) *OrderUpdateOne {
	ouo.mutation.AddPRICE(i)
	return ouo
}

// SetADDEDTIME sets the ADDED_TIME field.
func (ouo *OrderUpdateOne) SetADDEDTIME(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetADDEDTIME(t)
	return ouo
}

// SetEquipmentID sets the equipment edge to Equipment by id.
func (ouo *OrderUpdateOne) SetEquipmentID(id int) *OrderUpdateOne {
	ouo.mutation.SetEquipmentID(id)
	return ouo
}

// SetNillableEquipmentID sets the equipment edge to Equipment by id if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableEquipmentID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetEquipmentID(*id)
	}
	return ouo
}

// SetEquipment sets the equipment edge to Equipment.
func (ouo *OrderUpdateOne) SetEquipment(e *Equipment) *OrderUpdateOne {
	return ouo.SetEquipmentID(e.ID)
}

// SetCompanyID sets the company edge to Company by id.
func (ouo *OrderUpdateOne) SetCompanyID(id int) *OrderUpdateOne {
	ouo.mutation.SetCompanyID(id)
	return ouo
}

// SetNillableCompanyID sets the company edge to Company by id if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCompanyID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetCompanyID(*id)
	}
	return ouo
}

// SetCompany sets the company edge to Company.
func (ouo *OrderUpdateOne) SetCompany(c *Company) *OrderUpdateOne {
	return ouo.SetCompanyID(c.ID)
}

// SetUserID sets the user edge to User by id.
func (ouo *OrderUpdateOne) SetUserID(id int) *OrderUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetUserID(*id)
	}
	return ouo
}

// SetUser sets the user edge to User.
func (ouo *OrderUpdateOne) SetUser(u *User) *OrderUpdateOne {
	return ouo.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearEquipment clears the equipment edge to Equipment.
func (ouo *OrderUpdateOne) ClearEquipment() *OrderUpdateOne {
	ouo.mutation.ClearEquipment()
	return ouo
}

// ClearCompany clears the company edge to Company.
func (ouo *OrderUpdateOne) ClearCompany() *OrderUpdateOne {
	ouo.mutation.ClearCompany()
	return ouo
}

// ClearUser clears the user edge to User.
func (ouo *OrderUpdateOne) ClearUser() *OrderUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// Save executes the query and returns the updated entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {

	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	o, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (o *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Order.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ouo.mutation.AMOUNT(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAMOUNT,
		})
	}
	if value, ok := ouo.mutation.AddedAMOUNT(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAMOUNT,
		})
	}
	if value, ok := ouo.mutation.PRICE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldPRICE,
		})
	}
	if value, ok := ouo.mutation.AddedPRICE(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldPRICE,
		})
	}
	if value, ok := ouo.mutation.ADDEDTIME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldADDEDTIME,
		})
	}
	if ouo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.EquipmentTable,
			Columns: []string{order.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.EquipmentTable,
			Columns: []string{order.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CompanyTable,
			Columns: []string{order.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CompanyTable,
			Columns: []string{order.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	o = &Order{config: ouo.config}
	_spec.Assign = o.assignValues
	_spec.ScanValues = o.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return o, nil
}
