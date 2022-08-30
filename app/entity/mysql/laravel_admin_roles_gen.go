package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminRoles struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey"`                              //
	Name      string         `gorm:"column:name;type:varchar(50);index:laravel_admin_roles_name_unique,class:BTREE,unique"` //
	Slug      string         `gorm:"column:slug;type:varchar(50);index:laravel_admin_roles_slug_unique,class:BTREE,unique"` //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                         //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                         //
}

func (receiver *LaravelAdminRoles) TableName() string {
	return "laravel_admin_roles"
}

type OrmLaravelAdminRoles struct {
	db *gorm.DB
}

func NewOrmLaravelAdminRoles() *OrmLaravelAdminRoles {
	orm := &OrmLaravelAdminRoles{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminRoles{})
	return orm
}

func (orm *OrmLaravelAdminRoles) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminRoles) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminRoles) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminRoles) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminRoles) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminRoles) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminRoles) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminRoles) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminRoles) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminRoles) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminRoles) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminRoles) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminRoles) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminRoles) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminRoles) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminRoles) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminRoles) Insert(row *LaravelAdminRoles) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminRoles) Inserts(rows []*LaravelAdminRoles) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminRoles) Order(value interface{}) *OrmLaravelAdminRoles {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminRoles) Limit(limit int) *OrmLaravelAdminRoles {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminRoles) Offset(offset int) *OrmLaravelAdminRoles {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminRoles) Get() LaravelAdminRolesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminRoles) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminRoles) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminRoles{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminRoles) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_roles")
}

func (orm *OrmLaravelAdminRoles) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminRoles) First(conds ...interface{}) (*LaravelAdminRoles, bool) {
	dest := &LaravelAdminRoles{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminRoles) Take(conds ...interface{}) (*LaravelAdminRoles, int64) {
	dest := &LaravelAdminRoles{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminRoles) Last(conds ...interface{}) (*LaravelAdminRoles, int64) {
	dest := &LaravelAdminRoles{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminRoles) Find(conds ...interface{}) (LaravelAdminRolesList, int64) {
	list := make([]*LaravelAdminRoles, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminRoles) Paginate(page int, limit int) (LaravelAdminRolesList, int64) {
	var total int64
	list := make([]*LaravelAdminRoles, 0)
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
func (orm *OrmLaravelAdminRoles) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminRoles) FirstOrInit(dest *LaravelAdminRoles, conds ...interface{}) (*LaravelAdminRoles, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminRoles) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRoles) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRoles) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminRoles) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminRoles) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminRoles) Where(query interface{}, args ...interface{}) *OrmLaravelAdminRoles {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminRoles) And(fuc func(orm *OrmLaravelAdminRoles)) *OrmLaravelAdminRoles {
	ormAnd := NewOrmLaravelAdminRoles()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminRoles) Or(fuc func(orm *OrmLaravelAdminRoles)) *OrmLaravelAdminRoles {
	ormOr := NewOrmLaravelAdminRoles()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminRoles) WhereId(val uint32) *OrmLaravelAdminRoles {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) InsertGetId(row *LaravelAdminRoles) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmLaravelAdminRoles) WhereIdIn(val []uint32) *OrmLaravelAdminRoles {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereIdGt(val uint32) *OrmLaravelAdminRoles {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereIdGte(val uint32) *OrmLaravelAdminRoles {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereIdLt(val uint32) *OrmLaravelAdminRoles {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereIdLte(val uint32) *OrmLaravelAdminRoles {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereName(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) InsertGetName(row *LaravelAdminRoles) string {
	orm.db.Create(row)
	return row.Name
}
func (orm *OrmLaravelAdminRoles) WhereNameIn(val []string) *OrmLaravelAdminRoles {
	orm.db.Where("`name` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereNameGt(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`name` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereNameGte(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`name` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereNameLt(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`name` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereNameLte(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`name` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereSlug(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`slug` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) InsertGetSlug(row *LaravelAdminRoles) string {
	orm.db.Create(row)
	return row.Slug
}
func (orm *OrmLaravelAdminRoles) WhereSlugIn(val []string) *OrmLaravelAdminRoles {
	orm.db.Where("`slug` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereSlugGt(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`slug` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereSlugGte(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`slug` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereSlugLt(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`slug` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereSlugLte(val string) *OrmLaravelAdminRoles {
	orm.db.Where("`slug` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereCreatedAt(val database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereUpdatedAt(val database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoles) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminRoles {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminRolesList []*LaravelAdminRoles

func (l LaravelAdminRolesList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l LaravelAdminRolesList) GetIdMap() map[uint32]*LaravelAdminRoles {
	got := make(map[uint32]*LaravelAdminRoles)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
