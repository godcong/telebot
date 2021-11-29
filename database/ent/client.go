// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/motomototv/telebot/database/ent/migrate"

	"github.com/motomototv/telebot/database/ent/command"
	"github.com/motomototv/telebot/database/ent/message"
	"github.com/motomototv/telebot/database/ent/statistic"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Command is the client for interacting with the Command builders.
	Command *CommandClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// Statistic is the client for interacting with the Statistic builders.
	Statistic *StatisticClient
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
	c.Command = NewCommandClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.Statistic = NewStatisticClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Command:   NewCommandClient(cfg),
		Message:   NewMessageClient(cfg),
		Statistic: NewStatisticClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:    cfg,
		Command:   NewCommandClient(cfg),
		Message:   NewMessageClient(cfg),
		Statistic: NewStatisticClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Command.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Command.Use(hooks...)
	c.Message.Use(hooks...)
	c.Statistic.Use(hooks...)
}

// CommandClient is a client for the Command schema.
type CommandClient struct {
	config
}

// NewCommandClient returns a client for the Command from the given config.
func NewCommandClient(c config) *CommandClient {
	return &CommandClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `command.Hooks(f(g(h())))`.
func (c *CommandClient) Use(hooks ...Hook) {
	c.hooks.Command = append(c.hooks.Command, hooks...)
}

// Create returns a create builder for Command.
func (c *CommandClient) Create() *CommandCreate {
	mutation := newCommandMutation(c.config, OpCreate)
	return &CommandCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Command entities.
func (c *CommandClient) CreateBulk(builders ...*CommandCreate) *CommandCreateBulk {
	return &CommandCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Command.
func (c *CommandClient) Update() *CommandUpdate {
	mutation := newCommandMutation(c.config, OpUpdate)
	return &CommandUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommandClient) UpdateOne(co *Command) *CommandUpdateOne {
	mutation := newCommandMutation(c.config, OpUpdateOne, withCommand(co))
	return &CommandUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommandClient) UpdateOneID(id int) *CommandUpdateOne {
	mutation := newCommandMutation(c.config, OpUpdateOne, withCommandID(id))
	return &CommandUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Command.
func (c *CommandClient) Delete() *CommandDelete {
	mutation := newCommandMutation(c.config, OpDelete)
	return &CommandDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommandClient) DeleteOne(co *Command) *CommandDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommandClient) DeleteOneID(id int) *CommandDeleteOne {
	builder := c.Delete().Where(command.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommandDeleteOne{builder}
}

// Query returns a query builder for Command.
func (c *CommandClient) Query() *CommandQuery {
	return &CommandQuery{
		config: c.config,
	}
}

// Get returns a Command entity by its id.
func (c *CommandClient) Get(ctx context.Context, id int) (*Command, error) {
	return c.Query().Where(command.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommandClient) GetX(ctx context.Context, id int) *Command {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CommandClient) Hooks() []Hook {
	return c.hooks.Command
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Create returns a create builder for Message.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id int) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MessageClient) DeleteOneID(id int) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id int) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id int) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// StatisticClient is a client for the Statistic schema.
type StatisticClient struct {
	config
}

// NewStatisticClient returns a client for the Statistic from the given config.
func NewStatisticClient(c config) *StatisticClient {
	return &StatisticClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `statistic.Hooks(f(g(h())))`.
func (c *StatisticClient) Use(hooks ...Hook) {
	c.hooks.Statistic = append(c.hooks.Statistic, hooks...)
}

// Create returns a create builder for Statistic.
func (c *StatisticClient) Create() *StatisticCreate {
	mutation := newStatisticMutation(c.config, OpCreate)
	return &StatisticCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Statistic entities.
func (c *StatisticClient) CreateBulk(builders ...*StatisticCreate) *StatisticCreateBulk {
	return &StatisticCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Statistic.
func (c *StatisticClient) Update() *StatisticUpdate {
	mutation := newStatisticMutation(c.config, OpUpdate)
	return &StatisticUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatisticClient) UpdateOne(s *Statistic) *StatisticUpdateOne {
	mutation := newStatisticMutation(c.config, OpUpdateOne, withStatistic(s))
	return &StatisticUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatisticClient) UpdateOneID(id int) *StatisticUpdateOne {
	mutation := newStatisticMutation(c.config, OpUpdateOne, withStatisticID(id))
	return &StatisticUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Statistic.
func (c *StatisticClient) Delete() *StatisticDelete {
	mutation := newStatisticMutation(c.config, OpDelete)
	return &StatisticDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StatisticClient) DeleteOne(s *Statistic) *StatisticDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StatisticClient) DeleteOneID(id int) *StatisticDeleteOne {
	builder := c.Delete().Where(statistic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatisticDeleteOne{builder}
}

// Query returns a query builder for Statistic.
func (c *StatisticClient) Query() *StatisticQuery {
	return &StatisticQuery{
		config: c.config,
	}
}

// Get returns a Statistic entity by its id.
func (c *StatisticClient) Get(ctx context.Context, id int) (*Statistic, error) {
	return c.Query().Where(statistic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatisticClient) GetX(ctx context.Context, id int) *Statistic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StatisticClient) Hooks() []Hook {
	return c.hooks.Statistic
}
