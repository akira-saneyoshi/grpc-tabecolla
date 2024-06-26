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

// LoginUser is an object representing the database table.
type LoginUser struct { // ID
	ID int `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ユーザID文字列
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// メールアドレス
	UserMail string `boil:"user_mail" json:"user_mail" toml:"user_mail" yaml:"user_mail"`
	// パスワード
	UserPassword string `boil:"user_password" json:"user_password" toml:"user_password" yaml:"user_password"`
	// ユーザ名
	UserName string `boil:"user_name" json:"user_name" toml:"user_name" yaml:"user_name"`
	// 有効フラグ
	UserEnable bool `boil:"user_enable" json:"user_enable" toml:"user_enable" yaml:"user_enable"`

	R *loginUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L loginUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LoginUserColumns = struct {
	ID           string
	UserID       string
	UserMail     string
	UserPassword string
	UserName     string
	UserEnable   string
}{
	ID:           "id",
	UserID:       "user_id",
	UserMail:     "user_mail",
	UserPassword: "user_password",
	UserName:     "user_name",
	UserEnable:   "user_enable",
}

var LoginUserTableColumns = struct {
	ID           string
	UserID       string
	UserMail     string
	UserPassword string
	UserName     string
	UserEnable   string
}{
	ID:           "login_users.id",
	UserID:       "login_users.user_id",
	UserMail:     "login_users.user_mail",
	UserPassword: "login_users.user_password",
	UserName:     "login_users.user_name",
	UserEnable:   "login_users.user_enable",
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

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var LoginUserWhere = struct {
	ID           whereHelperint
	UserID       whereHelperstring
	UserMail     whereHelperstring
	UserPassword whereHelperstring
	UserName     whereHelperstring
	UserEnable   whereHelperbool
}{
	ID:           whereHelperint{field: "`login_users`.`id`"},
	UserID:       whereHelperstring{field: "`login_users`.`user_id`"},
	UserMail:     whereHelperstring{field: "`login_users`.`user_mail`"},
	UserPassword: whereHelperstring{field: "`login_users`.`user_password`"},
	UserName:     whereHelperstring{field: "`login_users`.`user_name`"},
	UserEnable:   whereHelperbool{field: "`login_users`.`user_enable`"},
}

// LoginUserRels is where relationship names are stored.
var LoginUserRels = struct {
}{}

// loginUserR is where relationships are stored.
type loginUserR struct {
}

// NewStruct creates a new relationship struct
func (*loginUserR) NewStruct() *loginUserR {
	return &loginUserR{}
}

// loginUserL is where Load methods for each relationship are stored.
type loginUserL struct{}

var (
	loginUserAllColumns            = []string{"id", "user_id", "user_mail", "user_password", "user_name", "user_enable"}
	loginUserColumnsWithoutDefault = []string{"user_id", "user_mail", "user_password", "user_name", "user_enable"}
	loginUserColumnsWithDefault    = []string{"id"}
	loginUserPrimaryKeyColumns     = []string{"id"}
	loginUserGeneratedColumns      = []string{}
)

type (
	// LoginUserSlice is an alias for a slice of pointers to LoginUser.
	// This should almost always be used instead of []LoginUser.
	LoginUserSlice []*LoginUser
	// LoginUserHook is the signature for custom LoginUser hook methods
	LoginUserHook func(context.Context, boil.ContextExecutor, *LoginUser) error

	loginUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	loginUserType                 = reflect.TypeOf(&LoginUser{})
	loginUserMapping              = queries.MakeStructMapping(loginUserType)
	loginUserPrimaryKeyMapping, _ = queries.BindMapping(loginUserType, loginUserMapping, loginUserPrimaryKeyColumns)
	loginUserInsertCacheMut       sync.RWMutex
	loginUserInsertCache          = make(map[string]insertCache)
	loginUserUpdateCacheMut       sync.RWMutex
	loginUserUpdateCache          = make(map[string]updateCache)
	loginUserUpsertCacheMut       sync.RWMutex
	loginUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var loginUserAfterSelectMu sync.Mutex
var loginUserAfterSelectHooks []LoginUserHook

var loginUserBeforeInsertMu sync.Mutex
var loginUserBeforeInsertHooks []LoginUserHook
var loginUserAfterInsertMu sync.Mutex
var loginUserAfterInsertHooks []LoginUserHook

var loginUserBeforeUpdateMu sync.Mutex
var loginUserBeforeUpdateHooks []LoginUserHook
var loginUserAfterUpdateMu sync.Mutex
var loginUserAfterUpdateHooks []LoginUserHook

var loginUserBeforeDeleteMu sync.Mutex
var loginUserBeforeDeleteHooks []LoginUserHook
var loginUserAfterDeleteMu sync.Mutex
var loginUserAfterDeleteHooks []LoginUserHook

var loginUserBeforeUpsertMu sync.Mutex
var loginUserBeforeUpsertHooks []LoginUserHook
var loginUserAfterUpsertMu sync.Mutex
var loginUserAfterUpsertHooks []LoginUserHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LoginUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LoginUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LoginUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LoginUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LoginUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LoginUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LoginUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LoginUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LoginUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loginUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLoginUserHook registers your hook function for all future operations.
func AddLoginUserHook(hookPoint boil.HookPoint, loginUserHook LoginUserHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		loginUserAfterSelectMu.Lock()
		loginUserAfterSelectHooks = append(loginUserAfterSelectHooks, loginUserHook)
		loginUserAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		loginUserBeforeInsertMu.Lock()
		loginUserBeforeInsertHooks = append(loginUserBeforeInsertHooks, loginUserHook)
		loginUserBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		loginUserAfterInsertMu.Lock()
		loginUserAfterInsertHooks = append(loginUserAfterInsertHooks, loginUserHook)
		loginUserAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		loginUserBeforeUpdateMu.Lock()
		loginUserBeforeUpdateHooks = append(loginUserBeforeUpdateHooks, loginUserHook)
		loginUserBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		loginUserAfterUpdateMu.Lock()
		loginUserAfterUpdateHooks = append(loginUserAfterUpdateHooks, loginUserHook)
		loginUserAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		loginUserBeforeDeleteMu.Lock()
		loginUserBeforeDeleteHooks = append(loginUserBeforeDeleteHooks, loginUserHook)
		loginUserBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		loginUserAfterDeleteMu.Lock()
		loginUserAfterDeleteHooks = append(loginUserAfterDeleteHooks, loginUserHook)
		loginUserAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		loginUserBeforeUpsertMu.Lock()
		loginUserBeforeUpsertHooks = append(loginUserBeforeUpsertHooks, loginUserHook)
		loginUserBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		loginUserAfterUpsertMu.Lock()
		loginUserAfterUpsertHooks = append(loginUserAfterUpsertHooks, loginUserHook)
		loginUserAfterUpsertMu.Unlock()
	}
}

// One returns a single loginUser record from the query.
func (q loginUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LoginUser, error) {
	o := &LoginUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for login_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LoginUser records from the query.
func (q loginUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (LoginUserSlice, error) {
	var o []*LoginUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LoginUser slice")
	}

	if len(loginUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LoginUser records in the query.
func (q loginUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count login_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q loginUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if login_users exists")
	}

	return count > 0, nil
}

// LoginUsers retrieves all the records using an executor.
func LoginUsers(mods ...qm.QueryMod) loginUserQuery {
	mods = append(mods, qm.From("`login_users`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`login_users`.*"})
	}

	return loginUserQuery{q}
}

// FindLoginUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLoginUser(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*LoginUser, error) {
	loginUserObj := &LoginUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `login_users` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, loginUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from login_users")
	}

	if err = loginUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return loginUserObj, err
	}

	return loginUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LoginUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no login_users provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loginUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	loginUserInsertCacheMut.RLock()
	cache, cached := loginUserInsertCache[key]
	loginUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			loginUserAllColumns,
			loginUserColumnsWithDefault,
			loginUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(loginUserType, loginUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(loginUserType, loginUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `login_users` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `login_users` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `login_users` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, loginUserPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into login_users")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == loginUserMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for login_users")
	}

CacheNoHooks:
	if !cached {
		loginUserInsertCacheMut.Lock()
		loginUserInsertCache[key] = cache
		loginUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LoginUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LoginUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	loginUserUpdateCacheMut.RLock()
	cache, cached := loginUserUpdateCache[key]
	loginUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			loginUserAllColumns,
			loginUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update login_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `login_users` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, loginUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(loginUserType, loginUserMapping, append(wl, loginUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update login_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for login_users")
	}

	if !cached {
		loginUserUpdateCacheMut.Lock()
		loginUserUpdateCache[key] = cache
		loginUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q loginUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for login_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for login_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LoginUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loginUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `login_users` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loginUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in loginUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all loginUser")
	}
	return rowsAff, nil
}

var mySQLLoginUserUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LoginUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no login_users provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loginUserColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLoginUserUniqueColumns, o)

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

	loginUserUpsertCacheMut.RLock()
	cache, cached := loginUserUpsertCache[key]
	loginUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			loginUserAllColumns,
			loginUserColumnsWithDefault,
			loginUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			loginUserAllColumns,
			loginUserPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert login_users, could not build update column list")
		}

		ret := strmangle.SetComplement(loginUserAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`login_users`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `login_users` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(loginUserType, loginUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(loginUserType, loginUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for login_users")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == loginUserMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(loginUserType, loginUserMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for login_users")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for login_users")
	}

CacheNoHooks:
	if !cached {
		loginUserUpsertCacheMut.Lock()
		loginUserUpsertCache[key] = cache
		loginUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single LoginUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LoginUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LoginUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), loginUserPrimaryKeyMapping)
	sql := "DELETE FROM `login_users` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from login_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for login_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q loginUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no loginUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from login_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for login_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LoginUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(loginUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loginUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `login_users` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loginUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from loginUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for login_users")
	}

	if len(loginUserAfterDeleteHooks) != 0 {
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
func (o *LoginUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLoginUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LoginUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LoginUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loginUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `login_users`.* FROM `login_users` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loginUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LoginUserSlice")
	}

	*o = slice

	return nil
}

// LoginUserExists checks if the LoginUser row exists.
func LoginUserExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `login_users` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if login_users exists")
	}

	return exists, nil
}

// Exists checks if the LoginUser row exists.
func (o *LoginUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LoginUserExists(ctx, exec, o.ID)
}
