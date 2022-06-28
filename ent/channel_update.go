// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/vod"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ChannelUpdate) SetName(s string) *ChannelUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *ChannelUpdate) SetDisplayName(s string) *ChannelUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetImagePath sets the "image_path" field.
func (cu *ChannelUpdate) SetImagePath(s string) *ChannelUpdate {
	cu.mutation.SetImagePath(s)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ChannelUpdate) SetUpdatedAt(t time.Time) *ChannelUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (cu *ChannelUpdate) AddVodIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddVodIDs(ids...)
	return cu
}

// AddVods adds the "vods" edges to the Vod entity.
func (cu *ChannelUpdate) AddVods(v ...*Vod) *ChannelUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVodIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearVods clears all "vods" edges to the Vod entity.
func (cu *ChannelUpdate) ClearVods() *ChannelUpdate {
	cu.mutation.ClearVods()
	return cu
}

// RemoveVodIDs removes the "vods" edge to Vod entities by IDs.
func (cu *ChannelUpdate) RemoveVodIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveVodIDs(ids...)
	return cu
}

// RemoveVods removes "vods" edges to Vod entities.
func (cu *ChannelUpdate) RemoveVods(v ...*Vod) *ChannelUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVodIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChannelUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldDisplayName,
		})
	}
	if value, ok := cu.mutation.ImagePath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldImagePath,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldUpdatedAt,
		})
	}
	if cu.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVodsIDs(); len(nodes) > 0 && !cu.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetName sets the "name" field.
func (cuo *ChannelUpdateOne) SetName(s string) *ChannelUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *ChannelUpdateOne) SetDisplayName(s string) *ChannelUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetImagePath sets the "image_path" field.
func (cuo *ChannelUpdateOne) SetImagePath(s string) *ChannelUpdateOne {
	cuo.mutation.SetImagePath(s)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ChannelUpdateOne) SetUpdatedAt(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (cuo *ChannelUpdateOne) AddVodIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddVodIDs(ids...)
	return cuo
}

// AddVods adds the "vods" edges to the Vod entity.
func (cuo *ChannelUpdateOne) AddVods(v ...*Vod) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVodIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearVods clears all "vods" edges to the Vod entity.
func (cuo *ChannelUpdateOne) ClearVods() *ChannelUpdateOne {
	cuo.mutation.ClearVods()
	return cuo
}

// RemoveVodIDs removes the "vods" edge to Vod entities by IDs.
func (cuo *ChannelUpdateOne) RemoveVodIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveVodIDs(ids...)
	return cuo
}

// RemoveVods removes "vods" edges to Vod entities.
func (cuo *ChannelUpdateOne) RemoveVods(v ...*Vod) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVodIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	var (
		err  error
		node *Channel
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChannelUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Channel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldDisplayName,
		})
	}
	if value, ok := cuo.mutation.ImagePath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldImagePath,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldUpdatedAt,
		})
	}
	if cuo.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVodsIDs(); len(nodes) > 0 && !cuo.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
