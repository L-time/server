// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/dal/dao"
)

func newPersonCollect(db *gorm.DB, opts ...gen.DOOption) personCollect {
	_personCollect := personCollect{}

	_personCollect.personCollectDo.UseDB(db, opts...)
	_personCollect.personCollectDo.UseModel(&dao.PersonCollect{})

	tableName := _personCollect.personCollectDo.TableName()
	_personCollect.ALL = field.NewAsterisk(tableName)
	_personCollect.ID = field.NewUint32(tableName, "prsn_clt_id")
	_personCollect.Category = field.NewString(tableName, "prsn_clt_cat")
	_personCollect.TargetID = field.NewUint32(tableName, "prsn_clt_mid")
	_personCollect.UserID = field.NewUint32(tableName, "prsn_clt_uid")
	_personCollect.CreatedTime = field.NewUint32(tableName, "prsn_clt_dateline")

	_personCollect.fillFieldMap()

	return _personCollect
}

// personCollect 人物收藏
type personCollect struct {
	personCollectDo personCollectDo

	ALL         field.Asterisk
	ID          field.Uint32
	Category    field.String
	TargetID    field.Uint32
	UserID      field.Uint32
	CreatedTime field.Uint32

	fieldMap map[string]field.Expr
}

func (p personCollect) Table(newTableName string) *personCollect {
	p.personCollectDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personCollect) As(alias string) *personCollect {
	p.personCollectDo.DO = *(p.personCollectDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personCollect) updateTableName(table string) *personCollect {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint32(table, "prsn_clt_id")
	p.Category = field.NewString(table, "prsn_clt_cat")
	p.TargetID = field.NewUint32(table, "prsn_clt_mid")
	p.UserID = field.NewUint32(table, "prsn_clt_uid")
	p.CreatedTime = field.NewUint32(table, "prsn_clt_dateline")

	p.fillFieldMap()

	return p
}

func (p *personCollect) WithContext(ctx context.Context) *personCollectDo {
	return p.personCollectDo.WithContext(ctx)
}

func (p personCollect) TableName() string { return p.personCollectDo.TableName() }

func (p personCollect) Alias() string { return p.personCollectDo.Alias() }

func (p personCollect) Columns(cols ...field.Expr) gen.Columns {
	return p.personCollectDo.Columns(cols...)
}

func (p *personCollect) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personCollect) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["prsn_clt_id"] = p.ID
	p.fieldMap["prsn_clt_cat"] = p.Category
	p.fieldMap["prsn_clt_mid"] = p.TargetID
	p.fieldMap["prsn_clt_uid"] = p.UserID
	p.fieldMap["prsn_clt_dateline"] = p.CreatedTime
}

func (p personCollect) clone(db *gorm.DB) personCollect {
	p.personCollectDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p personCollect) replaceDB(db *gorm.DB) personCollect {
	p.personCollectDo.ReplaceDB(db)
	return p
}

type personCollectDo struct{ gen.DO }

func (p personCollectDo) Debug() *personCollectDo {
	return p.withDO(p.DO.Debug())
}

func (p personCollectDo) WithContext(ctx context.Context) *personCollectDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personCollectDo) ReadDB() *personCollectDo {
	return p.Clauses(dbresolver.Read)
}

func (p personCollectDo) WriteDB() *personCollectDo {
	return p.Clauses(dbresolver.Write)
}

func (p personCollectDo) Session(config *gorm.Session) *personCollectDo {
	return p.withDO(p.DO.Session(config))
}

func (p personCollectDo) Clauses(conds ...clause.Expression) *personCollectDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personCollectDo) Returning(value interface{}, columns ...string) *personCollectDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personCollectDo) Not(conds ...gen.Condition) *personCollectDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personCollectDo) Or(conds ...gen.Condition) *personCollectDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personCollectDo) Select(conds ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personCollectDo) Where(conds ...gen.Condition) *personCollectDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personCollectDo) Order(conds ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personCollectDo) Distinct(cols ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personCollectDo) Omit(cols ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personCollectDo) Join(table schema.Tabler, on ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personCollectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personCollectDo) RightJoin(table schema.Tabler, on ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personCollectDo) Group(cols ...field.Expr) *personCollectDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personCollectDo) Having(conds ...gen.Condition) *personCollectDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personCollectDo) Limit(limit int) *personCollectDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personCollectDo) Offset(offset int) *personCollectDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personCollectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personCollectDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personCollectDo) Unscoped() *personCollectDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personCollectDo) Create(values ...*dao.PersonCollect) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personCollectDo) CreateInBatches(values []*dao.PersonCollect, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personCollectDo) Save(values ...*dao.PersonCollect) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personCollectDo) First() (*dao.PersonCollect, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonCollect), nil
	}
}

func (p personCollectDo) Take() (*dao.PersonCollect, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonCollect), nil
	}
}

func (p personCollectDo) Last() (*dao.PersonCollect, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonCollect), nil
	}
}

func (p personCollectDo) Find() ([]*dao.PersonCollect, error) {
	result, err := p.DO.Find()
	return result.([]*dao.PersonCollect), err
}

func (p personCollectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.PersonCollect, err error) {
	buf := make([]*dao.PersonCollect, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personCollectDo) FindInBatches(result *[]*dao.PersonCollect, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personCollectDo) Attrs(attrs ...field.AssignExpr) *personCollectDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personCollectDo) Assign(attrs ...field.AssignExpr) *personCollectDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personCollectDo) Joins(fields ...field.RelationField) *personCollectDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personCollectDo) Preload(fields ...field.RelationField) *personCollectDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personCollectDo) FirstOrInit() (*dao.PersonCollect, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonCollect), nil
	}
}

func (p personCollectDo) FirstOrCreate() (*dao.PersonCollect, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonCollect), nil
	}
}

func (p personCollectDo) FindByPage(offset int, limit int) (result []*dao.PersonCollect, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p personCollectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personCollectDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personCollectDo) Delete(models ...*dao.PersonCollect) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personCollectDo) withDO(do gen.Dao) *personCollectDo {
	p.DO = *do.(*gen.DO)
	return p
}
