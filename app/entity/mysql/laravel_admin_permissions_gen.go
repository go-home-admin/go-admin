package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminPermissions struct {
	Id         uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey"`                                    //
	Name       string         `gorm:"column:name;type:varchar(50);index:laravel_admin_permissions_name_unique,class:BTREE,unique"` //
	Slug       string         `gorm:"column:slug;type:varchar(50);index:laravel_admin_permissions_slug_unique,class:BTREE,unique"` //
	HttpMethod *string        `gorm:"column:http_method;type:varchar(255);default:NULL"`                                           //
	HttpPath   *string        `gorm:"column:http_path;type:text;default:NULL"`                                                     //
	CreatedAt  *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                               //
	UpdatedAt  *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                               //
}

func (receiver *LaravelAdminPermissions) TableName() string {
	return "laravel_admin_permissions"
}

type OrmLaravelAdminPermissions struct {
	db *gorm.DB
}

func NewOrmLaravelAdminPermissions() *OrmLaravelAdminPermissions {
	orm := &OrmLaravelAdminPermissions{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminPermissions{})
	return orm
}

func (orm *OrmLaravelAdminPermissions) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminPermissions) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminPermissions) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminPermissions) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminPermissions) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminPermissions) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminPermissions) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminPermissions) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminPermissions) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminPermissions) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminPermissions) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminPermissions) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminPermissions) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminPermissions) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminPermissions) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminPermissions) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminPermissions) Insert(row *LaravelAdminPermissions) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminPermissions) Inserts(rows []*LaravelAdminPermissions) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminPermissions) Order(value interface{}) *OrmLaravelAdminPermissions {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminPermissions) Limit(limit int) *OrmLaravelAdminPermissions {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminPermissions) Offset(offset int) *OrmLaravelAdminPermissions {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminPermissions) Get() LaravelAdminPermissionsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminPermissions) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminPermissions) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminPermissions{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminPermissions) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_permissions")
}

func (orm *OrmLaravelAdminPermissions) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminPermissions) First(conds ...interface{}) (*LaravelAdminPermissions, bool) {
	dest := &LaravelAdminPermissions{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminPermissions) Take(conds ...interface{}) (*LaravelAdminPermissions, int64) {
	dest := &LaravelAdminPermissions{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminPermissions) Last(conds ...interface{}) (*LaravelAdminPermissions, int64) {
	dest := &LaravelAdminPermissions{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminPermissions) Find(conds ...interface{}) (LaravelAdminPermissionsList, int64) {
	list := make([]*LaravelAdminPermissions, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminPermissions) Paginate(page int, limit int) (LaravelAdminPermissionsList, int64) {
	var total int64
	list := make([]*LaravelAdminPermissions, 0)
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
func (orm *OrmLaravelAdminPermissions) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminPermissions) FirstOrInit(dest *LaravelAdminPermissions, conds ...interface{}) (*LaravelAdminPermissions, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminPermissions) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminPermissions) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminPermissions) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminPermissions) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminPermissions) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminPermissions) Where(query interface{}, args ...interface{}) *OrmLaravelAdminPermissions {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminPermissions) And(fuc func(orm *OrmLaravelAdminPermissions)) *OrmLaravelAdminPermissions {
	ormAnd := NewOrmLaravelAdminPermissions()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminPermissions) Or(fuc func(orm *OrmLaravelAdminPermissions)) *OrmLaravelAdminPermissions {
	ormOr := NewOrmLaravelAdminPermissions()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminPermissions) WhereId(val uint32) *OrmLaravelAdminPermissions {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) InsertGetId(row *LaravelAdminPermissions) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmLaravelAdminPermissions) WhereIdIn(val []uint32) *OrmLaravelAdminPermissions {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereIdGt(val uint32) *OrmLaravelAdminPermissions {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereIdGte(val uint32) *OrmLaravelAdminPermissions {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereIdLt(val uint32) *OrmLaravelAdminPermissions {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereIdLte(val uint32) *OrmLaravelAdminPermissions {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereName(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) InsertGetName(row *LaravelAdminPermissions) string {
	orm.db.Create(row)
	return row.Name
}
func (orm *OrmLaravelAdminPermissions) WhereNameIn(val []string) *OrmLaravelAdminPermissions {
	orm.db.Where("`name` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereNameGt(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`name` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereNameGte(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`name` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereNameLt(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`name` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereNameLte(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`name` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereSlug(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`slug` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) InsertGetSlug(row *LaravelAdminPermissions) string {
	orm.db.Create(row)
	return row.Slug
}
func (orm *OrmLaravelAdminPermissions) WhereSlugIn(val []string) *OrmLaravelAdminPermissions {
	orm.db.Where("`slug` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereSlugGt(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`slug` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereSlugGte(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`slug` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereSlugLt(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`slug` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereSlugLte(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`slug` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereHttpMethod(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`http_method` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereHttpPath(val string) *OrmLaravelAdminPermissions {
	orm.db.Where("`http_path` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereCreatedAt(val database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereUpdatedAt(val database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminPermissions) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminPermissions {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminPermissionsList []*LaravelAdminPermissions

func (l LaravelAdminPermissionsList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l LaravelAdminPermissionsList) GetIdMap() map[uint32]*LaravelAdminPermissions {
	got := make(map[uint32]*LaravelAdminPermissions)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
