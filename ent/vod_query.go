// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/vod"
)

// VodQuery is the builder for querying Vod entities.
type VodQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Vod
	// eager-loading edges.
	withChannel *ChannelQuery
	withQueue   *QueueQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VodQuery builder.
func (vq *VodQuery) Where(ps ...predicate.Vod) *VodQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VodQuery) Limit(limit int) *VodQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VodQuery) Offset(offset int) *VodQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VodQuery) Unique(unique bool) *VodQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *VodQuery) Order(o ...OrderFunc) *VodQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryChannel chains the current query on the "channel" edge.
func (vq *VodQuery) QueryChannel() *ChannelQuery {
	query := &ChannelQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vod.ChannelTable, vod.ChannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQueue chains the current query on the "queue" edge.
func (vq *VodQuery) QueryQueue() *QueueQuery {
	query := &QueueQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(queue.Table, queue.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, vod.QueueTable, vod.QueueColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vod entity from the query.
// Returns a *NotFoundError when no Vod was found.
func (vq *VodQuery) First(ctx context.Context) (*Vod, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vod.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VodQuery) FirstX(ctx context.Context) *Vod {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vod ID from the query.
// Returns a *NotFoundError when no Vod ID was found.
func (vq *VodQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vod.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VodQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vod entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vod entity is found.
// Returns a *NotFoundError when no Vod entities are found.
func (vq *VodQuery) Only(ctx context.Context) (*Vod, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vod.Label}
	default:
		return nil, &NotSingularError{vod.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VodQuery) OnlyX(ctx context.Context) *Vod {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vod ID in the query.
// Returns a *NotSingularError when more than one Vod ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VodQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = &NotSingularError{vod.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VodQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vods.
func (vq *VodQuery) All(ctx context.Context) ([]*Vod, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VodQuery) AllX(ctx context.Context) []*Vod {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vod IDs.
func (vq *VodQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := vq.Select(vod.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VodQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VodQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VodQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VodQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VodQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VodQuery) Clone() *VodQuery {
	if vq == nil {
		return nil
	}
	return &VodQuery{
		config:      vq.config,
		limit:       vq.limit,
		offset:      vq.offset,
		order:       append([]OrderFunc{}, vq.order...),
		predicates:  append([]predicate.Vod{}, vq.predicates...),
		withChannel: vq.withChannel.Clone(),
		withQueue:   vq.withQueue.Clone(),
		// clone intermediate query.
		sql:    vq.sql.Clone(),
		path:   vq.path,
		unique: vq.unique,
	}
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithChannel(opts ...func(*ChannelQuery)) *VodQuery {
	query := &ChannelQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withChannel = query
	return vq
}

// WithQueue tells the query-builder to eager-load the nodes that are connected to
// the "queue" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithQueue(opts ...func(*QueueQuery)) *VodQuery {
	query := &QueueQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withQueue = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExtID string `json:"ext_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vod.Query().
//		GroupBy(vod.FieldExtID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VodQuery) GroupBy(field string, fields ...string) *VodGroupBy {
	group := &VodGroupBy{config: vq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExtID string `json:"ext_id,omitempty"`
//	}
//
//	client.Vod.Query().
//		Select(vod.FieldExtID).
//		Scan(ctx, &v)
//
func (vq *VodQuery) Select(fields ...string) *VodSelect {
	vq.fields = append(vq.fields, fields...)
	return &VodSelect{VodQuery: vq}
}

func (vq *VodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !vod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VodQuery) sqlAll(ctx context.Context) ([]*Vod, error) {
	var (
		nodes       = []*Vod{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withChannel != nil,
			vq.withQueue != nil,
		}
	)
	if vq.withChannel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, vod.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Vod{config: vq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vq.withChannel; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Vod)
		for i := range nodes {
			if nodes[i].channel_vods == nil {
				continue
			}
			fk := *nodes[i].channel_vods
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(channel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "channel_vods" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Channel = n
			}
		}
	}

	if query := vq.withQueue; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Vod)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Queue(func(s *sql.Selector) {
			s.Where(sql.InValues(vod.QueueColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.vod_queue
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "vod_queue" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "vod_queue" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Queue = n
		}
	}

	return nodes, nil
}

func (vq *VodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.fields
	if len(vq.fields) > 0 {
		_spec.Unique = vq.unique != nil && *vq.unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vq *VodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vod.Table,
			Columns: vod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vod.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vod.FieldID)
		for i := range fields {
			if fields[i] != vod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vod.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = vod.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.unique != nil && *vq.unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VodGroupBy is the group-by builder for Vod entities.
type VodGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VodGroupBy) Aggregate(fns ...AggregateFunc) *VodGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vgb *VodGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VodGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vgb *VodGroupBy) StringsX(ctx context.Context) []string {
	v, err := vgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vgb *VodGroupBy) StringX(ctx context.Context) string {
	v, err := vgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VodGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vgb *VodGroupBy) IntsX(ctx context.Context) []int {
	v, err := vgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vgb *VodGroupBy) IntX(ctx context.Context) int {
	v, err := vgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VodGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vgb *VodGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vgb *VodGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VodGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vgb *VodGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VodGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vgb *VodGroupBy) BoolX(ctx context.Context) bool {
	v, err := vgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vgb *VodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vgb.fields {
		if !vod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VodGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}

// VodSelect is the builder for selecting fields of Vod entities.
type VodSelect struct {
	*VodQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VodQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vs *VodSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VodSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vs *VodSelect) StringsX(ctx context.Context) []string {
	v, err := vs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vs *VodSelect) StringX(ctx context.Context) string {
	v, err := vs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VodSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vs *VodSelect) IntsX(ctx context.Context) []int {
	v, err := vs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vs *VodSelect) IntX(ctx context.Context) int {
	v, err := vs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VodSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vs *VodSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vs *VodSelect) Float64X(ctx context.Context) float64 {
	v, err := vs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VodSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vs *VodSelect) BoolsX(ctx context.Context) []bool {
	v, err := vs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vs *VodSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = fmt.Errorf("ent: VodSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vs *VodSelect) BoolX(ctx context.Context) bool {
	v, err := vs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vs *VodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
