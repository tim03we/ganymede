// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/migrate"

	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/user"
	"github.com/zibbp/ganymede/ent/vod"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Channel is the client for interacting with the Channel builders.
	Channel *ChannelClient
	// Queue is the client for interacting with the Queue builders.
	Queue *QueueClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Vod is the client for interacting with the Vod builders.
	Vod *VodClient
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
	c.Channel = NewChannelClient(c.config)
	c.Queue = NewQueueClient(c.config)
	c.User = NewUserClient(c.config)
	c.Vod = NewVodClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		Channel: NewChannelClient(cfg),
		Queue:   NewQueueClient(cfg),
		User:    NewUserClient(cfg),
		Vod:     NewVodClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Channel: NewChannelClient(cfg),
		Queue:   NewQueueClient(cfg),
		User:    NewUserClient(cfg),
		Vod:     NewVodClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Channel.
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
	c.Channel.Use(hooks...)
	c.Queue.Use(hooks...)
	c.User.Use(hooks...)
	c.Vod.Use(hooks...)
}

// ChannelClient is a client for the Channel schema.
type ChannelClient struct {
	config
}

// NewChannelClient returns a client for the Channel from the given config.
func NewChannelClient(c config) *ChannelClient {
	return &ChannelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `channel.Hooks(f(g(h())))`.
func (c *ChannelClient) Use(hooks ...Hook) {
	c.hooks.Channel = append(c.hooks.Channel, hooks...)
}

// Create returns a create builder for Channel.
func (c *ChannelClient) Create() *ChannelCreate {
	mutation := newChannelMutation(c.config, OpCreate)
	return &ChannelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Channel entities.
func (c *ChannelClient) CreateBulk(builders ...*ChannelCreate) *ChannelCreateBulk {
	return &ChannelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Channel.
func (c *ChannelClient) Update() *ChannelUpdate {
	mutation := newChannelMutation(c.config, OpUpdate)
	return &ChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChannelClient) UpdateOne(ch *Channel) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannel(ch))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChannelClient) UpdateOneID(id uuid.UUID) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannelID(id))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Channel.
func (c *ChannelClient) Delete() *ChannelDelete {
	mutation := newChannelMutation(c.config, OpDelete)
	return &ChannelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChannelClient) DeleteOne(ch *Channel) *ChannelDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChannelClient) DeleteOneID(id uuid.UUID) *ChannelDeleteOne {
	builder := c.Delete().Where(channel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChannelDeleteOne{builder}
}

// Query returns a query builder for Channel.
func (c *ChannelClient) Query() *ChannelQuery {
	return &ChannelQuery{
		config: c.config,
	}
}

// Get returns a Channel entity by its id.
func (c *ChannelClient) Get(ctx context.Context, id uuid.UUID) (*Channel, error) {
	return c.Query().Where(channel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChannelClient) GetX(ctx context.Context, id uuid.UUID) *Channel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVods queries the vods edge of a Channel.
func (c *ChannelClient) QueryVods(ch *Channel) *VodQuery {
	query := &VodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(channel.Table, channel.FieldID, id),
			sqlgraph.To(vod.Table, vod.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, channel.VodsTable, channel.VodsColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChannelClient) Hooks() []Hook {
	return c.hooks.Channel
}

// QueueClient is a client for the Queue schema.
type QueueClient struct {
	config
}

// NewQueueClient returns a client for the Queue from the given config.
func NewQueueClient(c config) *QueueClient {
	return &QueueClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `queue.Hooks(f(g(h())))`.
func (c *QueueClient) Use(hooks ...Hook) {
	c.hooks.Queue = append(c.hooks.Queue, hooks...)
}

// Create returns a create builder for Queue.
func (c *QueueClient) Create() *QueueCreate {
	mutation := newQueueMutation(c.config, OpCreate)
	return &QueueCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Queue entities.
func (c *QueueClient) CreateBulk(builders ...*QueueCreate) *QueueCreateBulk {
	return &QueueCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Queue.
func (c *QueueClient) Update() *QueueUpdate {
	mutation := newQueueMutation(c.config, OpUpdate)
	return &QueueUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QueueClient) UpdateOne(q *Queue) *QueueUpdateOne {
	mutation := newQueueMutation(c.config, OpUpdateOne, withQueue(q))
	return &QueueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QueueClient) UpdateOneID(id uuid.UUID) *QueueUpdateOne {
	mutation := newQueueMutation(c.config, OpUpdateOne, withQueueID(id))
	return &QueueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Queue.
func (c *QueueClient) Delete() *QueueDelete {
	mutation := newQueueMutation(c.config, OpDelete)
	return &QueueDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *QueueClient) DeleteOne(q *Queue) *QueueDeleteOne {
	return c.DeleteOneID(q.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *QueueClient) DeleteOneID(id uuid.UUID) *QueueDeleteOne {
	builder := c.Delete().Where(queue.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QueueDeleteOne{builder}
}

// Query returns a query builder for Queue.
func (c *QueueClient) Query() *QueueQuery {
	return &QueueQuery{
		config: c.config,
	}
}

// Get returns a Queue entity by its id.
func (c *QueueClient) Get(ctx context.Context, id uuid.UUID) (*Queue, error) {
	return c.Query().Where(queue.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QueueClient) GetX(ctx context.Context, id uuid.UUID) *Queue {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVod queries the vod edge of a Queue.
func (c *QueueClient) QueryVod(q *Queue) *VodQuery {
	query := &VodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(queue.Table, queue.FieldID, id),
			sqlgraph.To(vod.Table, vod.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, queue.VodTable, queue.VodColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *QueueClient) Hooks() []Hook {
	return c.hooks.Queue
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

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VodClient is a client for the Vod schema.
type VodClient struct {
	config
}

// NewVodClient returns a client for the Vod from the given config.
func NewVodClient(c config) *VodClient {
	return &VodClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vod.Hooks(f(g(h())))`.
func (c *VodClient) Use(hooks ...Hook) {
	c.hooks.Vod = append(c.hooks.Vod, hooks...)
}

// Create returns a create builder for Vod.
func (c *VodClient) Create() *VodCreate {
	mutation := newVodMutation(c.config, OpCreate)
	return &VodCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vod entities.
func (c *VodClient) CreateBulk(builders ...*VodCreate) *VodCreateBulk {
	return &VodCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vod.
func (c *VodClient) Update() *VodUpdate {
	mutation := newVodMutation(c.config, OpUpdate)
	return &VodUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VodClient) UpdateOne(v *Vod) *VodUpdateOne {
	mutation := newVodMutation(c.config, OpUpdateOne, withVod(v))
	return &VodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VodClient) UpdateOneID(id uuid.UUID) *VodUpdateOne {
	mutation := newVodMutation(c.config, OpUpdateOne, withVodID(id))
	return &VodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vod.
func (c *VodClient) Delete() *VodDelete {
	mutation := newVodMutation(c.config, OpDelete)
	return &VodDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VodClient) DeleteOne(v *Vod) *VodDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VodClient) DeleteOneID(id uuid.UUID) *VodDeleteOne {
	builder := c.Delete().Where(vod.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VodDeleteOne{builder}
}

// Query returns a query builder for Vod.
func (c *VodClient) Query() *VodQuery {
	return &VodQuery{
		config: c.config,
	}
}

// Get returns a Vod entity by its id.
func (c *VodClient) Get(ctx context.Context, id uuid.UUID) (*Vod, error) {
	return c.Query().Where(vod.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VodClient) GetX(ctx context.Context, id uuid.UUID) *Vod {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChannel queries the channel edge of a Vod.
func (c *VodClient) QueryChannel(v *Vod) *ChannelQuery {
	query := &ChannelQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, id),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vod.ChannelTable, vod.ChannelColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryQueue queries the queue edge of a Vod.
func (c *VodClient) QueryQueue(v *Vod) *QueueQuery {
	query := &QueueQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, id),
			sqlgraph.To(queue.Table, queue.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, vod.QueueTable, vod.QueueColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VodClient) Hooks() []Hook {
	return c.hooks.Vod
}
