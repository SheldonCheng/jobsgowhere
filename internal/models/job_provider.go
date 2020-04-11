// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// JobProvider is an object representing the database table.
type JobProvider struct {
	PersonID       string    `boil:"person_id" json:"person_id" toml:"person_id" yaml:"person_id"`
	Title          string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	CurrentCompany string    `boil:"current_company" json:"current_company" toml:"current_company" yaml:"current_company"`
	HuntingMode    null.Int  `boil:"hunting_mode" json:"hunting_mode,omitempty" toml:"hunting_mode" yaml:"hunting_mode,omitempty"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *jobProviderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jobProviderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JobProviderColumns = struct {
	PersonID       string
	Title          string
	CurrentCompany string
	HuntingMode    string
	CreatedAt      string
}{
	PersonID:       "person_id",
	Title:          "title",
	CurrentCompany: "current_company",
	HuntingMode:    "hunting_mode",
	CreatedAt:      "created_at",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
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

var JobProviderWhere = struct {
	PersonID       whereHelperstring
	Title          whereHelperstring
	CurrentCompany whereHelperstring
	HuntingMode    whereHelpernull_Int
	CreatedAt      whereHelpertime_Time
}{
	PersonID:       whereHelperstring{field: "\"job_provider\".\"person_id\""},
	Title:          whereHelperstring{field: "\"job_provider\".\"title\""},
	CurrentCompany: whereHelperstring{field: "\"job_provider\".\"current_company\""},
	HuntingMode:    whereHelpernull_Int{field: "\"job_provider\".\"hunting_mode\""},
	CreatedAt:      whereHelpertime_Time{field: "\"job_provider\".\"created_at\""},
}

// JobProviderRels is where relationship names are stored.
var JobProviderRels = struct {
}{}

// jobProviderR is where relationships are stored.
type jobProviderR struct {
}

// NewStruct creates a new relationship struct
func (*jobProviderR) NewStruct() *jobProviderR {
	return &jobProviderR{}
}

// jobProviderL is where Load methods for each relationship are stored.
type jobProviderL struct{}

var (
	jobProviderAllColumns            = []string{"person_id", "title", "current_company", "hunting_mode", "created_at"}
	jobProviderColumnsWithoutDefault = []string{"person_id", "title", "current_company", "hunting_mode", "created_at"}
	jobProviderColumnsWithDefault    = []string{}
	jobProviderPrimaryKeyColumns     = []string{"person_id"}
)

type (
	// JobProviderSlice is an alias for a slice of pointers to JobProvider.
	// This should generally be used opposed to []JobProvider.
	JobProviderSlice []*JobProvider
	// JobProviderHook is the signature for custom JobProvider hook methods
	JobProviderHook func(context.Context, boil.ContextExecutor, *JobProvider) error

	jobProviderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jobProviderType                 = reflect.TypeOf(&JobProvider{})
	jobProviderMapping              = queries.MakeStructMapping(jobProviderType)
	jobProviderPrimaryKeyMapping, _ = queries.BindMapping(jobProviderType, jobProviderMapping, jobProviderPrimaryKeyColumns)
	jobProviderInsertCacheMut       sync.RWMutex
	jobProviderInsertCache          = make(map[string]insertCache)
	jobProviderUpdateCacheMut       sync.RWMutex
	jobProviderUpdateCache          = make(map[string]updateCache)
	jobProviderUpsertCacheMut       sync.RWMutex
	jobProviderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var jobProviderBeforeInsertHooks []JobProviderHook
var jobProviderBeforeUpdateHooks []JobProviderHook
var jobProviderBeforeDeleteHooks []JobProviderHook
var jobProviderBeforeUpsertHooks []JobProviderHook

var jobProviderAfterInsertHooks []JobProviderHook
var jobProviderAfterSelectHooks []JobProviderHook
var jobProviderAfterUpdateHooks []JobProviderHook
var jobProviderAfterDeleteHooks []JobProviderHook
var jobProviderAfterUpsertHooks []JobProviderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *JobProvider) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *JobProvider) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *JobProvider) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *JobProvider) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *JobProvider) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *JobProvider) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *JobProvider) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *JobProvider) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *JobProvider) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobProviderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJobProviderHook registers your hook function for all future operations.
func AddJobProviderHook(hookPoint boil.HookPoint, jobProviderHook JobProviderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		jobProviderBeforeInsertHooks = append(jobProviderBeforeInsertHooks, jobProviderHook)
	case boil.BeforeUpdateHook:
		jobProviderBeforeUpdateHooks = append(jobProviderBeforeUpdateHooks, jobProviderHook)
	case boil.BeforeDeleteHook:
		jobProviderBeforeDeleteHooks = append(jobProviderBeforeDeleteHooks, jobProviderHook)
	case boil.BeforeUpsertHook:
		jobProviderBeforeUpsertHooks = append(jobProviderBeforeUpsertHooks, jobProviderHook)
	case boil.AfterInsertHook:
		jobProviderAfterInsertHooks = append(jobProviderAfterInsertHooks, jobProviderHook)
	case boil.AfterSelectHook:
		jobProviderAfterSelectHooks = append(jobProviderAfterSelectHooks, jobProviderHook)
	case boil.AfterUpdateHook:
		jobProviderAfterUpdateHooks = append(jobProviderAfterUpdateHooks, jobProviderHook)
	case boil.AfterDeleteHook:
		jobProviderAfterDeleteHooks = append(jobProviderAfterDeleteHooks, jobProviderHook)
	case boil.AfterUpsertHook:
		jobProviderAfterUpsertHooks = append(jobProviderAfterUpsertHooks, jobProviderHook)
	}
}

// One returns a single jobProvider record from the query.
func (q jobProviderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*JobProvider, error) {
	o := &JobProvider{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for job_provider")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all JobProvider records from the query.
func (q jobProviderQuery) All(ctx context.Context, exec boil.ContextExecutor) (JobProviderSlice, error) {
	var o []*JobProvider

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to JobProvider slice")
	}

	if len(jobProviderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all JobProvider records in the query.
func (q jobProviderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count job_provider rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q jobProviderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if job_provider exists")
	}

	return count > 0, nil
}

// JobProviders retrieves all the records using an executor.
func JobProviders(mods ...qm.QueryMod) jobProviderQuery {
	mods = append(mods, qm.From("\"job_provider\""))
	return jobProviderQuery{NewQuery(mods...)}
}

// FindJobProvider retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJobProvider(ctx context.Context, exec boil.ContextExecutor, personID string, selectCols ...string) (*JobProvider, error) {
	jobProviderObj := &JobProvider{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"job_provider\" where \"person_id\"=$1", sel,
	)

	q := queries.Raw(query, personID)

	err := q.Bind(ctx, exec, jobProviderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from job_provider")
	}

	return jobProviderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *JobProvider) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no job_provider provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jobProviderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	jobProviderInsertCacheMut.RLock()
	cache, cached := jobProviderInsertCache[key]
	jobProviderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			jobProviderAllColumns,
			jobProviderColumnsWithDefault,
			jobProviderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(jobProviderType, jobProviderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jobProviderType, jobProviderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"job_provider\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"job_provider\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into job_provider")
	}

	if !cached {
		jobProviderInsertCacheMut.Lock()
		jobProviderInsertCache[key] = cache
		jobProviderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the JobProvider.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *JobProvider) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	jobProviderUpdateCacheMut.RLock()
	cache, cached := jobProviderUpdateCache[key]
	jobProviderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			jobProviderAllColumns,
			jobProviderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update job_provider, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"job_provider\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jobProviderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jobProviderType, jobProviderMapping, append(wl, jobProviderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update job_provider row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for job_provider")
	}

	if !cached {
		jobProviderUpdateCacheMut.Lock()
		jobProviderUpdateCache[key] = cache
		jobProviderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q jobProviderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for job_provider")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for job_provider")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JobProviderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"job_provider\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, jobProviderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in jobProvider slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all jobProvider")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *JobProvider) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no job_provider provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jobProviderColumnsWithDefault, o)

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

	jobProviderUpsertCacheMut.RLock()
	cache, cached := jobProviderUpsertCache[key]
	jobProviderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			jobProviderAllColumns,
			jobProviderColumnsWithDefault,
			jobProviderColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			jobProviderAllColumns,
			jobProviderPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert job_provider, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jobProviderPrimaryKeyColumns))
			copy(conflict, jobProviderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"job_provider\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(jobProviderType, jobProviderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jobProviderType, jobProviderMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert job_provider")
	}

	if !cached {
		jobProviderUpsertCacheMut.Lock()
		jobProviderUpsertCache[key] = cache
		jobProviderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single JobProvider record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JobProvider) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no JobProvider provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jobProviderPrimaryKeyMapping)
	sql := "DELETE FROM \"job_provider\" WHERE \"person_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from job_provider")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for job_provider")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q jobProviderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no jobProviderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from job_provider")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for job_provider")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JobProviderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(jobProviderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"job_provider\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jobProviderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from jobProvider slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for job_provider")
	}

	if len(jobProviderAfterDeleteHooks) != 0 {
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
func (o *JobProvider) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJobProvider(ctx, exec, o.PersonID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JobProviderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JobProviderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"job_provider\".* FROM \"job_provider\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jobProviderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JobProviderSlice")
	}

	*o = slice

	return nil
}

// JobProviderExists checks if the JobProvider row exists.
func JobProviderExists(ctx context.Context, exec boil.ContextExecutor, personID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"job_provider\" where \"person_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, personID)
	}
	row := exec.QueryRowContext(ctx, sql, personID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if job_provider exists")
	}

	return exists, nil
}