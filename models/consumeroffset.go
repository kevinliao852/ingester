// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Consumeroffset is an object representing the database table.
type Consumeroffset struct {
	ConsumerID     int      `boil:"consumer_id" json:"consumer_id" toml:"consumer_id" yaml:"consumer_id"`
	ConsumerOffset null.Int `boil:"consumer_offset" json:"consumer_offset,omitempty" toml:"consumer_offset" yaml:"consumer_offset,omitempty"`

	R *consumeroffsetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L consumeroffsetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ConsumeroffsetColumns = struct {
	ConsumerID     string
	ConsumerOffset string
}{
	ConsumerID:     "consumer_id",
	ConsumerOffset: "consumer_offset",
}

var ConsumeroffsetTableColumns = struct {
	ConsumerID     string
	ConsumerOffset string
}{
	ConsumerID:     "consumeroffset.consumer_id",
	ConsumerOffset: "consumeroffset.consumer_offset",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ConsumeroffsetWhere = struct {
	ConsumerID     whereHelperint
	ConsumerOffset whereHelpernull_Int
}{
	ConsumerID:     whereHelperint{field: "\"consumeroffset\".\"consumer_id\""},
	ConsumerOffset: whereHelpernull_Int{field: "\"consumeroffset\".\"consumer_offset\""},
}

// ConsumeroffsetRels is where relationship names are stored.
var ConsumeroffsetRels = struct {
}{}

// consumeroffsetR is where relationships are stored.
type consumeroffsetR struct {
}

// NewStruct creates a new relationship struct
func (*consumeroffsetR) NewStruct() *consumeroffsetR {
	return &consumeroffsetR{}
}

// consumeroffsetL is where Load methods for each relationship are stored.
type consumeroffsetL struct{}

var (
	consumeroffsetAllColumns            = []string{"consumer_id", "consumer_offset"}
	consumeroffsetColumnsWithoutDefault = []string{}
	consumeroffsetColumnsWithDefault    = []string{"consumer_id", "consumer_offset"}
	consumeroffsetPrimaryKeyColumns     = []string{"consumer_id"}
	consumeroffsetGeneratedColumns      = []string{}
)

type (
	// ConsumeroffsetSlice is an alias for a slice of pointers to Consumeroffset.
	// This should almost always be used instead of []Consumeroffset.
	ConsumeroffsetSlice []*Consumeroffset
	// ConsumeroffsetHook is the signature for custom Consumeroffset hook methods
	ConsumeroffsetHook func(context.Context, boil.ContextExecutor, *Consumeroffset) error

	consumeroffsetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	consumeroffsetType                 = reflect.TypeOf(&Consumeroffset{})
	consumeroffsetMapping              = queries.MakeStructMapping(consumeroffsetType)
	consumeroffsetPrimaryKeyMapping, _ = queries.BindMapping(consumeroffsetType, consumeroffsetMapping, consumeroffsetPrimaryKeyColumns)
	consumeroffsetInsertCacheMut       sync.RWMutex
	consumeroffsetInsertCache          = make(map[string]insertCache)
	consumeroffsetUpdateCacheMut       sync.RWMutex
	consumeroffsetUpdateCache          = make(map[string]updateCache)
	consumeroffsetUpsertCacheMut       sync.RWMutex
	consumeroffsetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var consumeroffsetAfterSelectMu sync.Mutex
var consumeroffsetAfterSelectHooks []ConsumeroffsetHook

var consumeroffsetBeforeInsertMu sync.Mutex
var consumeroffsetBeforeInsertHooks []ConsumeroffsetHook
var consumeroffsetAfterInsertMu sync.Mutex
var consumeroffsetAfterInsertHooks []ConsumeroffsetHook

var consumeroffsetBeforeUpdateMu sync.Mutex
var consumeroffsetBeforeUpdateHooks []ConsumeroffsetHook
var consumeroffsetAfterUpdateMu sync.Mutex
var consumeroffsetAfterUpdateHooks []ConsumeroffsetHook

var consumeroffsetBeforeDeleteMu sync.Mutex
var consumeroffsetBeforeDeleteHooks []ConsumeroffsetHook
var consumeroffsetAfterDeleteMu sync.Mutex
var consumeroffsetAfterDeleteHooks []ConsumeroffsetHook

var consumeroffsetBeforeUpsertMu sync.Mutex
var consumeroffsetBeforeUpsertHooks []ConsumeroffsetHook
var consumeroffsetAfterUpsertMu sync.Mutex
var consumeroffsetAfterUpsertHooks []ConsumeroffsetHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Consumeroffset) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Consumeroffset) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Consumeroffset) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Consumeroffset) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Consumeroffset) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Consumeroffset) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Consumeroffset) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Consumeroffset) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Consumeroffset) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range consumeroffsetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddConsumeroffsetHook registers your hook function for all future operations.
func AddConsumeroffsetHook(hookPoint boil.HookPoint, consumeroffsetHook ConsumeroffsetHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		consumeroffsetAfterSelectMu.Lock()
		consumeroffsetAfterSelectHooks = append(consumeroffsetAfterSelectHooks, consumeroffsetHook)
		consumeroffsetAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		consumeroffsetBeforeInsertMu.Lock()
		consumeroffsetBeforeInsertHooks = append(consumeroffsetBeforeInsertHooks, consumeroffsetHook)
		consumeroffsetBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		consumeroffsetAfterInsertMu.Lock()
		consumeroffsetAfterInsertHooks = append(consumeroffsetAfterInsertHooks, consumeroffsetHook)
		consumeroffsetAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		consumeroffsetBeforeUpdateMu.Lock()
		consumeroffsetBeforeUpdateHooks = append(consumeroffsetBeforeUpdateHooks, consumeroffsetHook)
		consumeroffsetBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		consumeroffsetAfterUpdateMu.Lock()
		consumeroffsetAfterUpdateHooks = append(consumeroffsetAfterUpdateHooks, consumeroffsetHook)
		consumeroffsetAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		consumeroffsetBeforeDeleteMu.Lock()
		consumeroffsetBeforeDeleteHooks = append(consumeroffsetBeforeDeleteHooks, consumeroffsetHook)
		consumeroffsetBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		consumeroffsetAfterDeleteMu.Lock()
		consumeroffsetAfterDeleteHooks = append(consumeroffsetAfterDeleteHooks, consumeroffsetHook)
		consumeroffsetAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		consumeroffsetBeforeUpsertMu.Lock()
		consumeroffsetBeforeUpsertHooks = append(consumeroffsetBeforeUpsertHooks, consumeroffsetHook)
		consumeroffsetBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		consumeroffsetAfterUpsertMu.Lock()
		consumeroffsetAfterUpsertHooks = append(consumeroffsetAfterUpsertHooks, consumeroffsetHook)
		consumeroffsetAfterUpsertMu.Unlock()
	}
}

// One returns a single consumeroffset record from the query.
func (q consumeroffsetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Consumeroffset, error) {
	o := &Consumeroffset{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for consumeroffset")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Consumeroffset records from the query.
func (q consumeroffsetQuery) All(ctx context.Context, exec boil.ContextExecutor) (ConsumeroffsetSlice, error) {
	var o []*Consumeroffset

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Consumeroffset slice")
	}

	if len(consumeroffsetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Consumeroffset records in the query.
func (q consumeroffsetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count consumeroffset rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q consumeroffsetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if consumeroffset exists")
	}

	return count > 0, nil
}

// Consumeroffsets retrieves all the records using an executor.
func Consumeroffsets(mods ...qm.QueryMod) consumeroffsetQuery {
	mods = append(mods, qm.From("\"consumeroffset\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"consumeroffset\".*"})
	}

	return consumeroffsetQuery{q}
}

// FindConsumeroffset retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindConsumeroffset(ctx context.Context, exec boil.ContextExecutor, consumerID int, selectCols ...string) (*Consumeroffset, error) {
	consumeroffsetObj := &Consumeroffset{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"consumeroffset\" where \"consumer_id\"=$1", sel,
	)

	q := queries.Raw(query, consumerID)

	err := q.Bind(ctx, exec, consumeroffsetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from consumeroffset")
	}

	if err = consumeroffsetObj.doAfterSelectHooks(ctx, exec); err != nil {
		return consumeroffsetObj, err
	}

	return consumeroffsetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Consumeroffset) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no consumeroffset provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(consumeroffsetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	consumeroffsetInsertCacheMut.RLock()
	cache, cached := consumeroffsetInsertCache[key]
	consumeroffsetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			consumeroffsetAllColumns,
			consumeroffsetColumnsWithDefault,
			consumeroffsetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(consumeroffsetType, consumeroffsetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(consumeroffsetType, consumeroffsetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"consumeroffset\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"consumeroffset\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into consumeroffset")
	}

	if !cached {
		consumeroffsetInsertCacheMut.Lock()
		consumeroffsetInsertCache[key] = cache
		consumeroffsetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Consumeroffset.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Consumeroffset) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	consumeroffsetUpdateCacheMut.RLock()
	cache, cached := consumeroffsetUpdateCache[key]
	consumeroffsetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			consumeroffsetAllColumns,
			consumeroffsetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update consumeroffset, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"consumeroffset\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, consumeroffsetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(consumeroffsetType, consumeroffsetMapping, append(wl, consumeroffsetPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update consumeroffset row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for consumeroffset")
	}

	if !cached {
		consumeroffsetUpdateCacheMut.Lock()
		consumeroffsetUpdateCache[key] = cache
		consumeroffsetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q consumeroffsetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for consumeroffset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for consumeroffset")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ConsumeroffsetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), consumeroffsetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"consumeroffset\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, consumeroffsetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in consumeroffset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all consumeroffset")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Consumeroffset) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no consumeroffset provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(consumeroffsetColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	consumeroffsetUpsertCacheMut.RLock()
	cache, cached := consumeroffsetUpsertCache[key]
	consumeroffsetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			consumeroffsetAllColumns,
			consumeroffsetColumnsWithDefault,
			consumeroffsetColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			consumeroffsetAllColumns,
			consumeroffsetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert consumeroffset, could not build update column list")
		}

		ret := strmangle.SetComplement(consumeroffsetAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(consumeroffsetPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert consumeroffset, could not build conflict column list")
			}

			conflict = make([]string, len(consumeroffsetPrimaryKeyColumns))
			copy(conflict, consumeroffsetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"consumeroffset\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(consumeroffsetType, consumeroffsetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(consumeroffsetType, consumeroffsetMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert consumeroffset")
	}

	if !cached {
		consumeroffsetUpsertCacheMut.Lock()
		consumeroffsetUpsertCache[key] = cache
		consumeroffsetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Consumeroffset record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Consumeroffset) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Consumeroffset provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), consumeroffsetPrimaryKeyMapping)
	sql := "DELETE FROM \"consumeroffset\" WHERE \"consumer_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from consumeroffset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for consumeroffset")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q consumeroffsetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no consumeroffsetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from consumeroffset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for consumeroffset")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ConsumeroffsetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(consumeroffsetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), consumeroffsetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"consumeroffset\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, consumeroffsetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from consumeroffset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for consumeroffset")
	}

	if len(consumeroffsetAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Consumeroffset) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindConsumeroffset(ctx, exec, o.ConsumerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ConsumeroffsetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ConsumeroffsetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), consumeroffsetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"consumeroffset\".* FROM \"consumeroffset\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, consumeroffsetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ConsumeroffsetSlice")
	}

	*o = slice

	return nil
}

// ConsumeroffsetExists checks if the Consumeroffset row exists.
func ConsumeroffsetExists(ctx context.Context, exec boil.ContextExecutor, consumerID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"consumeroffset\" where \"consumer_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, consumerID)
	}
	row := exec.QueryRowContext(ctx, sql, consumerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if consumeroffset exists")
	}

	return exists, nil
}

// Exists checks if the Consumeroffset row exists.
func (o *Consumeroffset) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ConsumeroffsetExists(ctx, exec, o.ConsumerID)
}