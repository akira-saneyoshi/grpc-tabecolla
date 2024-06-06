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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// StoreCategory is an object representing the database table.
type StoreCategory struct { // ID
	ID int `boil:"id" json:"id" toml:"id" yaml:"id"`
	// é£²é£Ÿåº—ã‚«ãƒ†ã‚´ãƒªIDæ–‡å­—åˆ—
	ObjID string `boil:"obj_id" json:"obj_id" toml:"obj_id" yaml:"obj_id"`
	// é£²é£Ÿåº—ã‚«ãƒ†ã‚´ãƒªå
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *storeCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L storeCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StoreCategoryColumns = struct {
	ID    string
	ObjID string
	Name  string
}{
	ID:    "id",
	ObjID: "obj_id",
	Name:  "name",
}

var StoreCategoryTableColumns = struct {
	ID    string
	ObjID string
	Name  string
}{
	ID:    "store_categories.id",
	ObjID: "store_categories.obj_id",
	Name:  "store_categories.name",
}

// Generated where

var StoreCategoryWhere = struct {
	ID    whereHelperint
	ObjID whereHelperstring
	Name  whereHelperstring
}{
	ID:    whereHelperint{field: "`store_categories`.`id`"},
	ObjID: whereHelperstring{field: "`store_categories`.`obj_id`"},
	Name:  whereHelperstring{field: "`store_categories`.`name`"},
}

// StoreCategoryRels is where relationship names are stored.
var StoreCategoryRels = struct {
	Stores string
}{
	Stores: "Stores",
}

// storeCategoryR is where relationships are stored.
type storeCategoryR struct {
	Stores StoreSlice `boil:"Stores" json:"Stores" toml:"Stores" yaml:"Stores"`
}

// NewStruct creates a new relationship struct
func (*storeCategoryR) NewStruct() *storeCategoryR {
	return &storeCategoryR{}
}

func (r *storeCategoryR) GetStores() StoreSlice {
	if r == nil {
		return nil
	}
	return r.Stores
}

// storeCategoryL is where Load methods for each relationship are stored.
type storeCategoryL struct{}

var (
	storeCategoryAllColumns            = []string{"id", "obj_id", "name"}
	storeCategoryColumnsWithoutDefault = []string{"obj_id", "name"}
	storeCategoryColumnsWithDefault    = []string{"id"}
	storeCategoryPrimaryKeyColumns     = []string{"id"}
	storeCategoryGeneratedColumns      = []string{}
)

type (
	// StoreCategorySlice is an alias for a slice of pointers to StoreCategory.
	// This should almost always be used instead of []StoreCategory.
	StoreCategorySlice []*StoreCategory
	// StoreCategoryHook is the signature for custom StoreCategory hook methods
	StoreCategoryHook func(context.Context, boil.ContextExecutor, *StoreCategory) error

	storeCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	storeCategoryType                 = reflect.TypeOf(&StoreCategory{})
	storeCategoryMapping              = queries.MakeStructMapping(storeCategoryType)
	storeCategoryPrimaryKeyMapping, _ = queries.BindMapping(storeCategoryType, storeCategoryMapping, storeCategoryPrimaryKeyColumns)
	storeCategoryInsertCacheMut       sync.RWMutex
	storeCategoryInsertCache          = make(map[string]insertCache)
	storeCategoryUpdateCacheMut       sync.RWMutex
	storeCategoryUpdateCache          = make(map[string]updateCache)
	storeCategoryUpsertCacheMut       sync.RWMutex
	storeCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var storeCategoryAfterSelectMu sync.Mutex
var storeCategoryAfterSelectHooks []StoreCategoryHook

var storeCategoryBeforeInsertMu sync.Mutex
var storeCategoryBeforeInsertHooks []StoreCategoryHook
var storeCategoryAfterInsertMu sync.Mutex
var storeCategoryAfterInsertHooks []StoreCategoryHook

var storeCategoryBeforeUpdateMu sync.Mutex
var storeCategoryBeforeUpdateHooks []StoreCategoryHook
var storeCategoryAfterUpdateMu sync.Mutex
var storeCategoryAfterUpdateHooks []StoreCategoryHook

var storeCategoryBeforeDeleteMu sync.Mutex
var storeCategoryBeforeDeleteHooks []StoreCategoryHook
var storeCategoryAfterDeleteMu sync.Mutex
var storeCategoryAfterDeleteHooks []StoreCategoryHook

var storeCategoryBeforeUpsertMu sync.Mutex
var storeCategoryBeforeUpsertHooks []StoreCategoryHook
var storeCategoryAfterUpsertMu sync.Mutex
var storeCategoryAfterUpsertHooks []StoreCategoryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StoreCategory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StoreCategory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StoreCategory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StoreCategory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StoreCategory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StoreCategory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StoreCategory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StoreCategory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StoreCategory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeCategoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStoreCategoryHook registers your hook function for all future operations.
func AddStoreCategoryHook(hookPoint boil.HookPoint, storeCategoryHook StoreCategoryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		storeCategoryAfterSelectMu.Lock()
		storeCategoryAfterSelectHooks = append(storeCategoryAfterSelectHooks, storeCategoryHook)
		storeCategoryAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		storeCategoryBeforeInsertMu.Lock()
		storeCategoryBeforeInsertHooks = append(storeCategoryBeforeInsertHooks, storeCategoryHook)
		storeCategoryBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		storeCategoryAfterInsertMu.Lock()
		storeCategoryAfterInsertHooks = append(storeCategoryAfterInsertHooks, storeCategoryHook)
		storeCategoryAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		storeCategoryBeforeUpdateMu.Lock()
		storeCategoryBeforeUpdateHooks = append(storeCategoryBeforeUpdateHooks, storeCategoryHook)
		storeCategoryBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		storeCategoryAfterUpdateMu.Lock()
		storeCategoryAfterUpdateHooks = append(storeCategoryAfterUpdateHooks, storeCategoryHook)
		storeCategoryAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		storeCategoryBeforeDeleteMu.Lock()
		storeCategoryBeforeDeleteHooks = append(storeCategoryBeforeDeleteHooks, storeCategoryHook)
		storeCategoryBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		storeCategoryAfterDeleteMu.Lock()
		storeCategoryAfterDeleteHooks = append(storeCategoryAfterDeleteHooks, storeCategoryHook)
		storeCategoryAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		storeCategoryBeforeUpsertMu.Lock()
		storeCategoryBeforeUpsertHooks = append(storeCategoryBeforeUpsertHooks, storeCategoryHook)
		storeCategoryBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		storeCategoryAfterUpsertMu.Lock()
		storeCategoryAfterUpsertHooks = append(storeCategoryAfterUpsertHooks, storeCategoryHook)
		storeCategoryAfterUpsertMu.Unlock()
	}
}

// One returns a single storeCategory record from the query.
func (q storeCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StoreCategory, error) {
	o := &StoreCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for store_categories")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all StoreCategory records from the query.
func (q storeCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (StoreCategorySlice, error) {
	var o []*StoreCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StoreCategory slice")
	}

	if len(storeCategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all StoreCategory records in the query.
func (q storeCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count store_categories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q storeCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if store_categories exists")
	}

	return count > 0, nil
}

// Stores retrieves all the store's Stores with an executor.
func (o *StoreCategory) Stores(mods ...qm.QueryMod) storeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`stores`.`store_category_id`=?", o.ObjID),
	)

	return Stores(queryMods...)
}

// LoadStores allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (storeCategoryL) LoadStores(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStoreCategory interface{}, mods queries.Applicator) error {
	var slice []*StoreCategory
	var object *StoreCategory

	if singular {
		var ok bool
		object, ok = maybeStoreCategory.(*StoreCategory)
		if !ok {
			object = new(StoreCategory)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeStoreCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeStoreCategory))
			}
		}
	} else {
		s, ok := maybeStoreCategory.(*[]*StoreCategory)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeStoreCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeStoreCategory))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &storeCategoryR{}
		}
		args[object.ObjID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &storeCategoryR{}
			}
			args[obj.ObjID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`stores`),
		qm.WhereIn(`stores.store_category_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stores")
	}

	var resultSlice []*Store
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stores")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on stores")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for stores")
	}

	if len(storeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Stores = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &storeR{}
			}
			foreign.R.StoreCategory = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ObjID == foreign.StoreCategoryID {
				local.R.Stores = append(local.R.Stores, foreign)
				if foreign.R == nil {
					foreign.R = &storeR{}
				}
				foreign.R.StoreCategory = local
				break
			}
		}
	}

	return nil
}

// AddStores adds the given related objects to the existing relationships
// of the store_category, optionally inserting them as new records.
// Appends related to o.R.Stores.
// Sets related.R.StoreCategory appropriately.
func (o *StoreCategory) AddStores(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Store) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StoreCategoryID = o.ObjID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `stores` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"store_category_id"}),
				strmangle.WhereClause("`", "`", 0, storePrimaryKeyColumns),
			)
			values := []interface{}{o.ObjID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StoreCategoryID = o.ObjID
		}
	}

	if o.R == nil {
		o.R = &storeCategoryR{
			Stores: related,
		}
	} else {
		o.R.Stores = append(o.R.Stores, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &storeR{
				StoreCategory: o,
			}
		} else {
			rel.R.StoreCategory = o
		}
	}
	return nil
}

// StoreCategories retrieves all the records using an executor.
func StoreCategories(mods ...qm.QueryMod) storeCategoryQuery {
	mods = append(mods, qm.From("`store_categories`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`store_categories`.*"})
	}

	return storeCategoryQuery{q}
}

// FindStoreCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStoreCategory(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*StoreCategory, error) {
	storeCategoryObj := &StoreCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `store_categories` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, storeCategoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from store_categories")
	}

	if err = storeCategoryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return storeCategoryObj, err
	}

	return storeCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *StoreCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no store_categories provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(storeCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	storeCategoryInsertCacheMut.RLock()
	cache, cached := storeCategoryInsertCache[key]
	storeCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			storeCategoryAllColumns,
			storeCategoryColumnsWithDefault,
			storeCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(storeCategoryType, storeCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(storeCategoryType, storeCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `store_categories` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `store_categories` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `store_categories` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, storeCategoryPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into store_categories")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == storeCategoryMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for store_categories")
	}

CacheNoHooks:
	if !cached {
		storeCategoryInsertCacheMut.Lock()
		storeCategoryInsertCache[key] = cache
		storeCategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the StoreCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *StoreCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	storeCategoryUpdateCacheMut.RLock()
	cache, cached := storeCategoryUpdateCache[key]
	storeCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			storeCategoryAllColumns,
			storeCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update store_categories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `store_categories` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, storeCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(storeCategoryType, storeCategoryMapping, append(wl, storeCategoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update store_categories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for store_categories")
	}

	if !cached {
		storeCategoryUpdateCacheMut.Lock()
		storeCategoryUpdateCache[key] = cache
		storeCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q storeCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for store_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for store_categories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StoreCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `store_categories` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, storeCategoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in storeCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all storeCategory")
	}
	return rowsAff, nil
}

var mySQLStoreCategoryUniqueColumns = []string{
	"id",
	"obj_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StoreCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no store_categories provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(storeCategoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLStoreCategoryUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	storeCategoryUpsertCacheMut.RLock()
	cache, cached := storeCategoryUpsertCache[key]
	storeCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			storeCategoryAllColumns,
			storeCategoryColumnsWithDefault,
			storeCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			storeCategoryAllColumns,
			storeCategoryPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert store_categories, could not build update column list")
		}

		ret := strmangle.SetComplement(storeCategoryAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`store_categories`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `store_categories` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(storeCategoryType, storeCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(storeCategoryType, storeCategoryMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for store_categories")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == storeCategoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(storeCategoryType, storeCategoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for store_categories")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for store_categories")
	}

CacheNoHooks:
	if !cached {
		storeCategoryUpsertCacheMut.Lock()
		storeCategoryUpsertCache[key] = cache
		storeCategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single StoreCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StoreCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StoreCategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), storeCategoryPrimaryKeyMapping)
	sql := "DELETE FROM `store_categories` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from store_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for store_categories")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q storeCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no storeCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from store_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for store_categories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StoreCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(storeCategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `store_categories` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, storeCategoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from storeCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for store_categories")
	}

	if len(storeCategoryAfterDeleteHooks) != 0 {
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
func (o *StoreCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStoreCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StoreCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StoreCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `store_categories`.* FROM `store_categories` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, storeCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StoreCategorySlice")
	}

	*o = slice

	return nil
}

// StoreCategoryExists checks if the StoreCategory row exists.
func StoreCategoryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `store_categories` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if store_categories exists")
	}

	return exists, nil
}

// Exists checks if the StoreCategory row exists.
func (o *StoreCategory) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return StoreCategoryExists(ctx, exec, o.ID)
}
