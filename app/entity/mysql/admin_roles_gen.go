package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type AdminRoles struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey"`                              //
	Name      string         `gorm:"column:name;type:varchar(50);index:laravel_admin_roles_name_unique,class:BTREE,unique"` //
	Slug      string         `gorm:"column:slug;type:varchar(50);index:laravel_admin_roles_slug_unique,class:BTREE,unique"` //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                         //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                         //
}

func (receiver *AdminRoles) TableName() string {
	return "admin_roles"
}

type OrmAdminRoles struct {
	db *gorm.DB
}

func NewOrmAdminRoles() *OrmAdminRoles {
	orm := &OrmAdminRoles{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&AdminRoles{})
	return orm
}

func (orm *OrmAdminRoles) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmAdminRoles) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminRoles) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminRoles) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminRoles) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminRoles) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminRoles) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminRoles) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminRoles) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminRoles) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminRoles) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminRoles) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminRoles) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminRoles) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminRoles) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminRoles) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminRoles) Insert(row *AdminRoles) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminRoles) Inserts(rows []*AdminRoles) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminRoles) Order(value interface{}) *OrmAdminRoles {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminRoles) Limit(limit int) *OrmAdminRoles {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminRoles) Offset(offset int) *OrmAdminRoles {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminRoles) Get() AdminRolesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminRoles) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminRoles) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminRoles{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminRoles) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_roles")
}

func (orm *OrmAdminRoles) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminRoles) First(conds ...interface{}) (*AdminRoles, bool) {
	dest := &AdminRoles{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminRoles) Take(conds ...interface{}) (*AdminRoles, int64) {
	dest := &AdminRoles{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminRoles) Last(conds ...interface{}) (*AdminRoles, int64) {
	dest := &AdminRoles{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminRoles) Find(conds ...interface{}) (AdminRolesList, int64) {
	list := make([]*AdminRoles, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminRoles) Paginate(page int, limit int) (AdminRolesList, int64) {
	var total int64
	list := make([]*AdminRoles, 0)
	orm.db.Count(&total)
	if total > 0 {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * limit
		tx := orm.db.Offset(offset).Limit(limit).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}

	return list, total
}

// FindInBatches find records in batches
func (orm *OrmAdminRoles) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminRoles) FirstOrInit(dest *AdminRoles, conds ...interface{}) (*AdminRoles, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminRoles) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoles) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoles) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminRoles) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminRoles) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminRoles) Where(query interface{}, args ...interface{}) *OrmAdminRoles {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminRoles) And(fuc func(orm *OrmAdminRoles)) *OrmAdminRoles {
	ormAnd := NewOrmAdminRoles()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminRoles) Or(fuc func(orm *OrmAdminRoles)) *OrmAdminRoles {
	ormOr := NewOrmAdminRoles()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmAdminRoles) WhereId(val uint32) *OrmAdminRoles {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminRoles) InsertGetId(row *AdminRoles) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAdminRoles) WhereIdIn(val []uint32) *OrmAdminRoles {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereIdGt(val uint32) *OrmAdminRoles {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereIdGte(val uint32) *OrmAdminRoles {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereIdLt(val uint32) *OrmAdminRoles {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereIdLte(val uint32) *OrmAdminRoles {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereName(val string) *OrmAdminRoles {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmAdminRoles) InsertGetName(row *AdminRoles) string {
	orm.db.Create(row)
	return row.Name
}
func (orm *OrmAdminRoles) WhereNameIn(val []string) *OrmAdminRoles {
	orm.db.Where("`name` IN ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereNameGt(val string) *OrmAdminRoles {
	orm.db.Where("`name` > ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereNameGte(val string) *OrmAdminRoles {
	orm.db.Where("`name` >= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereNameLt(val string) *OrmAdminRoles {
	orm.db.Where("`name` < ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereNameLte(val string) *OrmAdminRoles {
	orm.db.Where("`name` <= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereSlug(val string) *OrmAdminRoles {
	orm.db.Where("`slug` = ?", val)
	return orm
}
func (orm *OrmAdminRoles) InsertGetSlug(row *AdminRoles) string {
	orm.db.Create(row)
	return row.Slug
}
func (orm *OrmAdminRoles) WhereSlugIn(val []string) *OrmAdminRoles {
	orm.db.Where("`slug` IN ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereSlugGt(val string) *OrmAdminRoles {
	orm.db.Where("`slug` > ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereSlugGte(val string) *OrmAdminRoles {
	orm.db.Where("`slug` >= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereSlugLt(val string) *OrmAdminRoles {
	orm.db.Where("`slug` < ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereSlugLte(val string) *OrmAdminRoles {
	orm.db.Where("`slug` <= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereCreatedAt(val database.Time) *OrmAdminRoles {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoles {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoles) WhereCreatedAtLte(val database.Time) *OrmAdminRoles {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereCreatedAtGte(val database.Time) *OrmAdminRoles {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereUpdatedAt(val database.Time) *OrmAdminRoles {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoles {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoles) WhereUpdatedAtLte(val database.Time) *OrmAdminRoles {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoles) WhereUpdatedAtGte(val database.Time) *OrmAdminRoles {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminRolesList []*AdminRoles

func (l AdminRolesList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminRolesList) GetIdMap() map[uint32]*AdminRoles {
	got := make(map[uint32]*AdminRoles)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
