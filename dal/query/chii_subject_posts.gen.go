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

func newSubjectPost(db *gorm.DB, opts ...gen.DOOption) subjectPost {
	_subjectPost := subjectPost{}

	_subjectPost.subjectPostDo.UseDB(db, opts...)
	_subjectPost.subjectPostDo.UseModel(&dao.SubjectPost{})

	tableName := _subjectPost.subjectPostDo.TableName()
	_subjectPost.ALL = field.NewAsterisk(tableName)
	_subjectPost.PostID = field.NewUint32(tableName, "sbj_pst_id")
	_subjectPost.FieldID = field.NewUint32(tableName, "sbj_pst_mid")
	_subjectPost.UserID = field.NewUint32(tableName, "sbj_pst_uid")
	_subjectPost.RelatedMessageID = field.NewUint32(tableName, "sbj_pst_related")
	_subjectPost.Content = field.NewString(tableName, "sbj_pst_content")
	_subjectPost.PostState = field.NewUint8(tableName, "sbj_pst_state")
	_subjectPost.CreatedTime = field.NewInt32(tableName, "sbj_pst_dateline")

	_subjectPost.fillFieldMap()

	return _subjectPost
}

type subjectPost struct {
	subjectPostDo subjectPostDo

	ALL              field.Asterisk
	PostID           field.Uint32
	FieldID          field.Uint32
	UserID           field.Uint32
	RelatedMessageID field.Uint32
	Content          field.String
	PostState        field.Uint8
	CreatedTime      field.Int32

	fieldMap map[string]field.Expr
}

func (s subjectPost) Table(newTableName string) *subjectPost {
	s.subjectPostDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectPost) As(alias string) *subjectPost {
	s.subjectPostDo.DO = *(s.subjectPostDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectPost) updateTableName(table string) *subjectPost {
	s.ALL = field.NewAsterisk(table)
	s.PostID = field.NewUint32(table, "sbj_pst_id")
	s.FieldID = field.NewUint32(table, "sbj_pst_mid")
	s.UserID = field.NewUint32(table, "sbj_pst_uid")
	s.RelatedMessageID = field.NewUint32(table, "sbj_pst_related")
	s.Content = field.NewString(table, "sbj_pst_content")
	s.PostState = field.NewUint8(table, "sbj_pst_state")
	s.CreatedTime = field.NewInt32(table, "sbj_pst_dateline")

	s.fillFieldMap()

	return s
}

func (s *subjectPost) WithContext(ctx context.Context) *subjectPostDo {
	return s.subjectPostDo.WithContext(ctx)
}

func (s subjectPost) TableName() string { return s.subjectPostDo.TableName() }

func (s subjectPost) Alias() string { return s.subjectPostDo.Alias() }

func (s subjectPost) Columns(cols ...field.Expr) gen.Columns { return s.subjectPostDo.Columns(cols...) }

func (s *subjectPost) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectPost) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["sbj_pst_id"] = s.PostID
	s.fieldMap["sbj_pst_mid"] = s.FieldID
	s.fieldMap["sbj_pst_uid"] = s.UserID
	s.fieldMap["sbj_pst_related"] = s.RelatedMessageID
	s.fieldMap["sbj_pst_content"] = s.Content
	s.fieldMap["sbj_pst_state"] = s.PostState
	s.fieldMap["sbj_pst_dateline"] = s.CreatedTime
}

func (s subjectPost) clone(db *gorm.DB) subjectPost {
	s.subjectPostDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s subjectPost) replaceDB(db *gorm.DB) subjectPost {
	s.subjectPostDo.ReplaceDB(db)
	return s
}

type subjectPostDo struct{ gen.DO }

func (s subjectPostDo) Debug() *subjectPostDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectPostDo) WithContext(ctx context.Context) *subjectPostDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectPostDo) ReadDB() *subjectPostDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectPostDo) WriteDB() *subjectPostDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectPostDo) Session(config *gorm.Session) *subjectPostDo {
	return s.withDO(s.DO.Session(config))
}

func (s subjectPostDo) Clauses(conds ...clause.Expression) *subjectPostDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectPostDo) Returning(value interface{}, columns ...string) *subjectPostDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectPostDo) Not(conds ...gen.Condition) *subjectPostDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectPostDo) Or(conds ...gen.Condition) *subjectPostDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectPostDo) Select(conds ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectPostDo) Where(conds ...gen.Condition) *subjectPostDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectPostDo) Order(conds ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectPostDo) Distinct(cols ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectPostDo) Omit(cols ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectPostDo) Join(table schema.Tabler, on ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectPostDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectPostDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectPostDo) Group(cols ...field.Expr) *subjectPostDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectPostDo) Having(conds ...gen.Condition) *subjectPostDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectPostDo) Limit(limit int) *subjectPostDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectPostDo) Offset(offset int) *subjectPostDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectPostDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectPostDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectPostDo) Unscoped() *subjectPostDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectPostDo) Create(values ...*dao.SubjectPost) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectPostDo) CreateInBatches(values []*dao.SubjectPost, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectPostDo) Save(values ...*dao.SubjectPost) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectPostDo) First() (*dao.SubjectPost, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectPost), nil
	}
}

func (s subjectPostDo) Take() (*dao.SubjectPost, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectPost), nil
	}
}

func (s subjectPostDo) Last() (*dao.SubjectPost, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectPost), nil
	}
}

func (s subjectPostDo) Find() ([]*dao.SubjectPost, error) {
	result, err := s.DO.Find()
	return result.([]*dao.SubjectPost), err
}

func (s subjectPostDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.SubjectPost, err error) {
	buf := make([]*dao.SubjectPost, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectPostDo) FindInBatches(result *[]*dao.SubjectPost, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectPostDo) Attrs(attrs ...field.AssignExpr) *subjectPostDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectPostDo) Assign(attrs ...field.AssignExpr) *subjectPostDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectPostDo) Joins(fields ...field.RelationField) *subjectPostDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectPostDo) Preload(fields ...field.RelationField) *subjectPostDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectPostDo) FirstOrInit() (*dao.SubjectPost, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectPost), nil
	}
}

func (s subjectPostDo) FirstOrCreate() (*dao.SubjectPost, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectPost), nil
	}
}

func (s subjectPostDo) FindByPage(offset int, limit int) (result []*dao.SubjectPost, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subjectPostDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectPostDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectPostDo) Delete(models ...*dao.SubjectPost) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectPostDo) withDO(do gen.Dao) *subjectPostDo {
	s.DO = *do.(*gen.DO)
	return s
}
