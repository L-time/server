// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/bangumi/server/internal/dal/dao"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newOAuthAccessToken(db *gorm.DB) oAuthAccessToken {
	_oAuthAccessToken := oAuthAccessToken{}

	_oAuthAccessToken.oAuthAccessTokenDo.UseDB(db)
	_oAuthAccessToken.oAuthAccessTokenDo.UseModel(&dao.OAuthAccessToken{})

	tableName := _oAuthAccessToken.oAuthAccessTokenDo.TableName()
	_oAuthAccessToken.ALL = field.NewField(tableName, "*")
	_oAuthAccessToken.AccessToken = field.NewString(tableName, "access_token")
	_oAuthAccessToken.ClientID = field.NewString(tableName, "client_id")
	_oAuthAccessToken.UserID = field.NewString(tableName, "user_id")
	_oAuthAccessToken.Expires = field.NewTime(tableName, "expires")
	_oAuthAccessToken.Scope = field.NewString(tableName, "scope")

	_oAuthAccessToken.fillFieldMap()

	return _oAuthAccessToken
}

type oAuthAccessToken struct {
	oAuthAccessTokenDo oAuthAccessTokenDo

	ALL         field.Field
	AccessToken field.String
	ClientID    field.String
	UserID      field.String
	Expires     field.Time
	Scope       field.String

	fieldMap map[string]field.Expr
}

func (o oAuthAccessToken) Table(newTableName string) *oAuthAccessToken {
	o.oAuthAccessTokenDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o oAuthAccessToken) As(alias string) *oAuthAccessToken {
	o.oAuthAccessTokenDo.DO = *(o.oAuthAccessTokenDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *oAuthAccessToken) updateTableName(table string) *oAuthAccessToken {
	o.ALL = field.NewField(table, "*")
	o.AccessToken = field.NewString(table, "access_token")
	o.ClientID = field.NewString(table, "client_id")
	o.UserID = field.NewString(table, "user_id")
	o.Expires = field.NewTime(table, "expires")
	o.Scope = field.NewString(table, "scope")

	o.fillFieldMap()

	return o
}

func (o *oAuthAccessToken) WithContext(ctx context.Context) *oAuthAccessTokenDo {
	return o.oAuthAccessTokenDo.WithContext(ctx)
}

func (o oAuthAccessToken) TableName() string { return o.oAuthAccessTokenDo.TableName() }

func (o *oAuthAccessToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *oAuthAccessToken) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 5)
	o.fieldMap["access_token"] = o.AccessToken
	o.fieldMap["client_id"] = o.ClientID
	o.fieldMap["user_id"] = o.UserID
	o.fieldMap["expires"] = o.Expires
	o.fieldMap["scope"] = o.Scope
}

func (o oAuthAccessToken) clone(db *gorm.DB) oAuthAccessToken {
	o.oAuthAccessTokenDo.ReplaceDB(db)
	return o
}

type oAuthAccessTokenDo struct{ gen.DO }

func (o oAuthAccessTokenDo) Debug() *oAuthAccessTokenDo {
	return o.withDO(o.DO.Debug())
}

func (o oAuthAccessTokenDo) WithContext(ctx context.Context) *oAuthAccessTokenDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o oAuthAccessTokenDo) Clauses(conds ...clause.Expression) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o oAuthAccessTokenDo) Not(conds ...gen.Condition) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o oAuthAccessTokenDo) Or(conds ...gen.Condition) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o oAuthAccessTokenDo) Select(conds ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o oAuthAccessTokenDo) Where(conds ...gen.Condition) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o oAuthAccessTokenDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oAuthAccessTokenDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o oAuthAccessTokenDo) Order(conds ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o oAuthAccessTokenDo) Distinct(cols ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o oAuthAccessTokenDo) Omit(cols ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o oAuthAccessTokenDo) Join(table schema.Tabler, on ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o oAuthAccessTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o oAuthAccessTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o oAuthAccessTokenDo) Group(cols ...field.Expr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o oAuthAccessTokenDo) Having(conds ...gen.Condition) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o oAuthAccessTokenDo) Limit(limit int) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o oAuthAccessTokenDo) Offset(offset int) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o oAuthAccessTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o oAuthAccessTokenDo) Unscoped() *oAuthAccessTokenDo {
	return o.withDO(o.DO.Unscoped())
}

func (o oAuthAccessTokenDo) Create(values ...*dao.OAuthAccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o oAuthAccessTokenDo) CreateInBatches(values []*dao.OAuthAccessToken, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o oAuthAccessTokenDo) Save(values ...*dao.OAuthAccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o oAuthAccessTokenDo) First() (*dao.OAuthAccessToken, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthAccessToken), nil
	}
}

func (o oAuthAccessTokenDo) Take() (*dao.OAuthAccessToken, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthAccessToken), nil
	}
}

func (o oAuthAccessTokenDo) Last() (*dao.OAuthAccessToken, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthAccessToken), nil
	}
}

func (o oAuthAccessTokenDo) Find() ([]*dao.OAuthAccessToken, error) {
	result, err := o.DO.Find()
	return result.([]*dao.OAuthAccessToken), err
}

func (o oAuthAccessTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.OAuthAccessToken, err error) {
	buf := make([]*dao.OAuthAccessToken, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o oAuthAccessTokenDo) FindInBatches(result *[]*dao.OAuthAccessToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o oAuthAccessTokenDo) Attrs(attrs ...field.AssignExpr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o oAuthAccessTokenDo) Assign(attrs ...field.AssignExpr) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o oAuthAccessTokenDo) Joins(field field.RelationField) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Joins(field))
}

func (o oAuthAccessTokenDo) Preload(field field.RelationField) *oAuthAccessTokenDo {
	return o.withDO(o.DO.Preload(field))
}

func (o oAuthAccessTokenDo) FirstOrInit() (*dao.OAuthAccessToken, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthAccessToken), nil
	}
}

func (o oAuthAccessTokenDo) FirstOrCreate() (*dao.OAuthAccessToken, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthAccessToken), nil
	}
}

func (o oAuthAccessTokenDo) FindByPage(offset int, limit int) (result []*dao.OAuthAccessToken, count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = o.Offset(offset).Limit(limit).Find()
	return
}

func (o oAuthAccessTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o *oAuthAccessTokenDo) withDO(do gen.Dao) *oAuthAccessTokenDo {
	o.DO = *do.(*gen.DO)
	return o
}