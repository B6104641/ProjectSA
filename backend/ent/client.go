// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/B6104641/app/ent/migrate"

	"github.com/B6104641/app/ent/company"
	"github.com/B6104641/app/ent/equipment"
	"github.com/B6104641/app/ent/order"
	"github.com/B6104641/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Company is the client for interacting with the Company builders.
	Company *CompanyClient
	// Equipment is the client for interacting with the Equipment builders.
	Equipment *EquipmentClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Company = NewCompanyClient(c.config)
	c.Equipment = NewEquipmentClient(c.config)
	c.Order = NewOrderClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Company:   NewCompanyClient(cfg),
		Equipment: NewEquipmentClient(cfg),
		Order:     NewOrderClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:    cfg,
		Company:   NewCompanyClient(cfg),
		Equipment: NewEquipmentClient(cfg),
		Order:     NewOrderClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Company.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Company.Use(hooks...)
	c.Equipment.Use(hooks...)
	c.Order.Use(hooks...)
	c.User.Use(hooks...)
}

// CompanyClient is a client for the Company schema.
type CompanyClient struct {
	config
}

// NewCompanyClient returns a client for the Company from the given config.
func NewCompanyClient(c config) *CompanyClient {
	return &CompanyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `company.Hooks(f(g(h())))`.
func (c *CompanyClient) Use(hooks ...Hook) {
	c.hooks.Company = append(c.hooks.Company, hooks...)
}

// Create returns a create builder for Company.
func (c *CompanyClient) Create() *CompanyCreate {
	mutation := newCompanyMutation(c.config, OpCreate)
	return &CompanyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Company.
func (c *CompanyClient) Update() *CompanyUpdate {
	mutation := newCompanyMutation(c.config, OpUpdate)
	return &CompanyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CompanyClient) UpdateOne(co *Company) *CompanyUpdateOne {
	mutation := newCompanyMutation(c.config, OpUpdateOne, withCompany(co))
	return &CompanyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CompanyClient) UpdateOneID(id int) *CompanyUpdateOne {
	mutation := newCompanyMutation(c.config, OpUpdateOne, withCompanyID(id))
	return &CompanyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Company.
func (c *CompanyClient) Delete() *CompanyDelete {
	mutation := newCompanyMutation(c.config, OpDelete)
	return &CompanyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CompanyClient) DeleteOne(co *Company) *CompanyDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CompanyClient) DeleteOneID(id int) *CompanyDeleteOne {
	builder := c.Delete().Where(company.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CompanyDeleteOne{builder}
}

// Create returns a query builder for Company.
func (c *CompanyClient) Query() *CompanyQuery {
	return &CompanyQuery{config: c.config}
}

// Get returns a Company entity by its id.
func (c *CompanyClient) Get(ctx context.Context, id int) (*Company, error) {
	return c.Query().Where(company.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CompanyClient) GetX(ctx context.Context, id int) *Company {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// QueryOrderCompanyy queries the order_companyy edge of a Company.
func (c *CompanyClient) QueryOrderCompanyy(co *Company) *OrderQuery {
	query := &OrderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.OrderCompanyyTable, company.OrderCompanyyColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CompanyClient) Hooks() []Hook {
	return c.hooks.Company
}

// EquipmentClient is a client for the Equipment schema.
type EquipmentClient struct {
	config
}

// NewEquipmentClient returns a client for the Equipment from the given config.
func NewEquipmentClient(c config) *EquipmentClient {
	return &EquipmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `equipment.Hooks(f(g(h())))`.
func (c *EquipmentClient) Use(hooks ...Hook) {
	c.hooks.Equipment = append(c.hooks.Equipment, hooks...)
}

// Create returns a create builder for Equipment.
func (c *EquipmentClient) Create() *EquipmentCreate {
	mutation := newEquipmentMutation(c.config, OpCreate)
	return &EquipmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Equipment.
func (c *EquipmentClient) Update() *EquipmentUpdate {
	mutation := newEquipmentMutation(c.config, OpUpdate)
	return &EquipmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EquipmentClient) UpdateOne(e *Equipment) *EquipmentUpdateOne {
	mutation := newEquipmentMutation(c.config, OpUpdateOne, withEquipment(e))
	return &EquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EquipmentClient) UpdateOneID(id int) *EquipmentUpdateOne {
	mutation := newEquipmentMutation(c.config, OpUpdateOne, withEquipmentID(id))
	return &EquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Equipment.
func (c *EquipmentClient) Delete() *EquipmentDelete {
	mutation := newEquipmentMutation(c.config, OpDelete)
	return &EquipmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EquipmentClient) DeleteOne(e *Equipment) *EquipmentDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EquipmentClient) DeleteOneID(id int) *EquipmentDeleteOne {
	builder := c.Delete().Where(equipment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EquipmentDeleteOne{builder}
}

// Create returns a query builder for Equipment.
func (c *EquipmentClient) Query() *EquipmentQuery {
	return &EquipmentQuery{config: c.config}
}

// Get returns a Equipment entity by its id.
func (c *EquipmentClient) Get(ctx context.Context, id int) (*Equipment, error) {
	return c.Query().Where(equipment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EquipmentClient) GetX(ctx context.Context, id int) *Equipment {
	e, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return e
}

// QueryOrderEquipmentt queries the order_equipmentt edge of a Equipment.
func (c *EquipmentClient) QueryOrderEquipmentt(e *Equipment) *OrderQuery {
	query := &OrderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, equipment.OrderEquipmenttTable, equipment.OrderEquipmenttColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EquipmentClient) Hooks() []Hook {
	return c.hooks.Equipment
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Create returns a create builder for Order.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id int) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OrderClient) DeleteOneID(id int) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Create returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{config: c.config}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id int) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id int) *Order {
	o, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return o
}

// QueryEquipment queries the equipment edge of a Order.
func (c *OrderClient) QueryEquipment(o *Order) *EquipmentQuery {
	query := &EquipmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.EquipmentTable, order.EquipmentColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCompany queries the company edge of a Order.
func (c *OrderClient) QueryCompany(o *Order) *CompanyQuery {
	query := &CompanyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.CompanyTable, order.CompanyColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Order.
func (c *OrderClient) QueryUser(o *Order) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.UserTable, order.UserColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryOrderUserr queries the order_userr edge of a User.
func (c *UserClient) QueryOrderUserr(u *User) *OrderQuery {
	query := &OrderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.OrderUserrTable, user.OrderUserrColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
