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

	"godemo/internal/gostorage/gormgendemo/model"
)

func newMyModel(db *gorm.DB, opts ...gen.DOOption) myModel {
	_myModel := myModel{}

	_myModel.myModelDo.UseDB(db, opts...)
	_myModel.myModelDo.UseModel(&model.MyModel{})

	tableName := _myModel.myModelDo.TableName()
	_myModel.ALL = field.NewAsterisk(tableName)
	_myModel.ID = field.NewString(tableName, "id")
	_myModel.Name = field.NewString(tableName, "name")
	_myModel.ProjectID = field.NewString(tableName, "project_id")
	_myModel.ParentModelID = field.NewString(tableName, "parent_model_id")
	_myModel.ModelTypeID = field.NewString(tableName, "model_type_id")
	_myModel.Description = field.NewString(tableName, "description")
	_myModel.Position = field.NewString(tableName, "position")
	_myModel.Icon = field.NewString(tableName, "icon")
	_myModel.TableName_ = field.NewString(tableName, "table_name")
	_myModel.Sort = field.NewInt32(tableName, "sort")
	_myModel.IsHidden = field.NewBool(tableName, "is_hidden")
	_myModel.IsDefault = field.NewBool(tableName, "is_default")
	_myModel.IsNet = field.NewBool(tableName, "is_net")
	_myModel.FrontendTemplate = field.NewString(tableName, "frontend_template")
	_myModel.ListModelIds = field.NewString(tableName, "list_model_ids")
	_myModel.Script = field.NewString(tableName, "script")
	_myModel.ViewRelations = field.NewString(tableName, "view_relations")
	_myModel.IsAutoFlush = field.NewBool(tableName, "is_auto_flush")
	_myModel.SSHConnectConf = field.NewInt32(tableName, "ssh_connect_conf")
	_myModel.TemplateIds = field.NewString(tableName, "template_ids")

	_myModel.fillFieldMap()

	return _myModel
}

type myModel struct {
	myModelDo

	ALL              field.Asterisk
	ID               field.String
	Name             field.String
	ProjectID        field.String
	ParentModelID    field.String
	ModelTypeID      field.String
	Description      field.String
	Position         field.String
	Icon             field.String
	TableName_       field.String
	Sort             field.Int32
	IsHidden         field.Bool
	IsDefault        field.Bool
	IsNet            field.Bool
	FrontendTemplate field.String
	ListModelIds     field.String
	Script           field.String
	ViewRelations    field.String // 视图关联关系
	IsAutoFlush      field.Bool   // 是否自动刷新
	SSHConnectConf   field.Int32  // ssh连接配置项
	TemplateIds      field.String

	fieldMap map[string]field.Expr
}

func (m myModel) Table(newTableName string) *myModel {
	m.myModelDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m myModel) As(alias string) *myModel {
	m.myModelDo.DO = *(m.myModelDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *myModel) updateTableName(table string) *myModel {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewString(table, "id")
	m.Name = field.NewString(table, "name")
	m.ProjectID = field.NewString(table, "project_id")
	m.ParentModelID = field.NewString(table, "parent_model_id")
	m.ModelTypeID = field.NewString(table, "model_type_id")
	m.Description = field.NewString(table, "description")
	m.Position = field.NewString(table, "position")
	m.Icon = field.NewString(table, "icon")
	m.TableName_ = field.NewString(table, "table_name")
	m.Sort = field.NewInt32(table, "sort")
	m.IsHidden = field.NewBool(table, "is_hidden")
	m.IsDefault = field.NewBool(table, "is_default")
	m.IsNet = field.NewBool(table, "is_net")
	m.FrontendTemplate = field.NewString(table, "frontend_template")
	m.ListModelIds = field.NewString(table, "list_model_ids")
	m.Script = field.NewString(table, "script")
	m.ViewRelations = field.NewString(table, "view_relations")
	m.IsAutoFlush = field.NewBool(table, "is_auto_flush")
	m.SSHConnectConf = field.NewInt32(table, "ssh_connect_conf")
	m.TemplateIds = field.NewString(table, "template_ids")

	m.fillFieldMap()

	return m
}

func (m *myModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *myModel) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 20)
	m.fieldMap["id"] = m.ID
	m.fieldMap["name"] = m.Name
	m.fieldMap["project_id"] = m.ProjectID
	m.fieldMap["parent_model_id"] = m.ParentModelID
	m.fieldMap["model_type_id"] = m.ModelTypeID
	m.fieldMap["description"] = m.Description
	m.fieldMap["position"] = m.Position
	m.fieldMap["icon"] = m.Icon
	m.fieldMap["table_name"] = m.TableName_
	m.fieldMap["sort"] = m.Sort
	m.fieldMap["is_hidden"] = m.IsHidden
	m.fieldMap["is_default"] = m.IsDefault
	m.fieldMap["is_net"] = m.IsNet
	m.fieldMap["frontend_template"] = m.FrontendTemplate
	m.fieldMap["list_model_ids"] = m.ListModelIds
	m.fieldMap["script"] = m.Script
	m.fieldMap["view_relations"] = m.ViewRelations
	m.fieldMap["is_auto_flush"] = m.IsAutoFlush
	m.fieldMap["ssh_connect_conf"] = m.SSHConnectConf
	m.fieldMap["template_ids"] = m.TemplateIds
}

func (m myModel) clone(db *gorm.DB) myModel {
	m.myModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m myModel) replaceDB(db *gorm.DB) myModel {
	m.myModelDo.ReplaceDB(db)
	return m
}

type myModelDo struct{ gen.DO }

func (m myModelDo) Debug() *myModelDo {
	return m.withDO(m.DO.Debug())
}

func (m myModelDo) WithContext(ctx context.Context) *myModelDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m myModelDo) ReadDB() *myModelDo {
	return m.Clauses(dbresolver.Read)
}

func (m myModelDo) WriteDB() *myModelDo {
	return m.Clauses(dbresolver.Write)
}

func (m myModelDo) Session(config *gorm.Session) *myModelDo {
	return m.withDO(m.DO.Session(config))
}

func (m myModelDo) Clauses(conds ...clause.Expression) *myModelDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m myModelDo) Returning(value interface{}, columns ...string) *myModelDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m myModelDo) Not(conds ...gen.Condition) *myModelDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m myModelDo) Or(conds ...gen.Condition) *myModelDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m myModelDo) Select(conds ...field.Expr) *myModelDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m myModelDo) Where(conds ...gen.Condition) *myModelDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m myModelDo) Order(conds ...field.Expr) *myModelDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m myModelDo) Distinct(cols ...field.Expr) *myModelDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m myModelDo) Omit(cols ...field.Expr) *myModelDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m myModelDo) Join(table schema.Tabler, on ...field.Expr) *myModelDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m myModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *myModelDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m myModelDo) RightJoin(table schema.Tabler, on ...field.Expr) *myModelDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m myModelDo) Group(cols ...field.Expr) *myModelDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m myModelDo) Having(conds ...gen.Condition) *myModelDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m myModelDo) Limit(limit int) *myModelDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m myModelDo) Offset(offset int) *myModelDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m myModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *myModelDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m myModelDo) Unscoped() *myModelDo {
	return m.withDO(m.DO.Unscoped())
}

func (m myModelDo) Create(values ...*model.MyModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m myModelDo) CreateInBatches(values []*model.MyModel, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m myModelDo) Save(values ...*model.MyModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m myModelDo) First() (*model.MyModel, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyModel), nil
	}
}

func (m myModelDo) Take() (*model.MyModel, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyModel), nil
	}
}

func (m myModelDo) Last() (*model.MyModel, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyModel), nil
	}
}

func (m myModelDo) Find() ([]*model.MyModel, error) {
	result, err := m.DO.Find()
	return result.([]*model.MyModel), err
}

func (m myModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MyModel, err error) {
	buf := make([]*model.MyModel, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m myModelDo) FindInBatches(result *[]*model.MyModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m myModelDo) Attrs(attrs ...field.AssignExpr) *myModelDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m myModelDo) Assign(attrs ...field.AssignExpr) *myModelDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m myModelDo) Joins(fields ...field.RelationField) *myModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m myModelDo) Preload(fields ...field.RelationField) *myModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m myModelDo) FirstOrInit() (*model.MyModel, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyModel), nil
	}
}

func (m myModelDo) FirstOrCreate() (*model.MyModel, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyModel), nil
	}
}

func (m myModelDo) FindByPage(offset int, limit int) (result []*model.MyModel, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m myModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m myModelDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m myModelDo) Delete(models ...*model.MyModel) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *myModelDo) withDO(do gen.Dao) *myModelDo {
	m.DO = *do.(*gen.DO)
	return m
}