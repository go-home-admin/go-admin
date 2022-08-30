package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminUserPermissions struct {
	UserId       int32          `gorm:"column:user_id;type:int(11);index:laravel_admin_user_permissions_user_id_permission_id_index,class:BTREE"`       //
	PermissionId int32          `gorm:"column:permission_id;type:int(11);index:laravel_admin_user_permissions_user_id_permission_id_index,class:BTREE"` //
	CreatedAt    *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                                                  //
	UpdatedAt    *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                                                  //
}

func (receiver *LaravelAdminUserPermissions) TableName() string {
	return "laravel_admin_user_permissions"
}

type OrmLaravelAdminUserPermissions struct {
	db *gorm.DB
}

func NewOrmLaravelAdminUserPermissions() *OrmLaravelAdminUserPermissions {
	orm := &OrmLaravelAdminUserPermissions{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminUserPermissions{})
	return orm
}

func (orm *OrmLaravelAdminUserPermissions) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminUserPermissions) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminUserPermissions) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminUserPermissions) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminUserPermissions) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminUserPermissions) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminUserPermissions) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminUserPermissions) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminUserPermissions) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminUserPermissions) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminUserPermissions) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminUserPermissions) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminUserPermissions) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminUserPermissions) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminUserPermissions) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminUserPermissions) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminUserPermissions) Insert(row *LaravelAdminUserPermissions) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminUserPermissions) Inserts(rows []*LaravelAdminUserPermissions) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminUserPermissions) Order(value interface{}) *OrmLaravelAdminUserPermissions {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminUserPermissions) Limit(limit int) *OrmLaravelAdminUserPermissions {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminUserPermissions) Offset(offset int) *OrmLaravelAdminUserPermissions {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminUserPermissions) Get() LaravelAdminUserPermissionsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminUserPermissions) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminUserPermissions) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminUserPermissions{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminUserPermissions) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_user_permissions")
}

func (orm *OrmLaravelAdminUserPermissions) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminUserPermissions) First(conds ...interface{}) (*LaravelAdminUserPermissions, bool) {
	dest := &LaravelAdminUserPermissions{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminUserPermissions) Take(conds ...interface{}) (*LaravelAdminUserPermissions, int64) {
	dest := &LaravelAdminUserPermissions{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminUserPermissions) Last(conds ...interface{}) (*LaravelAdminUserPermissions, int64) {
	dest := &LaravelAdminUserPermissions{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminUserPermissions) Find(conds ...interface{}) (LaravelAdminUserPermissionsList, int64) {
	list := make([]*LaravelAdminUserPermissions, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminUserPermissions) Paginate(page int, limit int) (LaravelAdminUserPermissionsList, int64) {
	var total int64
	list := make([]*LaravelAdminUserPermissions, 0)
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
func (orm *OrmLaravelAdminUserPermissions) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminUserPermissions) FirstOrInit(dest *LaravelAdminUserPermissions, conds ...interface{}) (*LaravelAdminUserPermissions, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminUserPermissions) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminUserPermissions) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminUserPermissions) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminUserPermissions) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminUserPermissions) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminUserPermissions) Where(query interface{}, args ...interface{}) *OrmLaravelAdminUserPermissions {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminUserPermissions) And(fuc func(orm *OrmLaravelAdminUserPermissions)) *OrmLaravelAdminUserPermissions {
	ormAnd := NewOrmLaravelAdminUserPermissions()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminUserPermissions) Or(fuc func(orm *OrmLaravelAdminUserPermissions)) *OrmLaravelAdminUserPermissions {
	ormOr := NewOrmLaravelAdminUserPermissions()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminUserPermissions) WhereUserId(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`user_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) InsertGetUserId(row *LaravelAdminUserPermissions) int32 {
	orm.db.Create(row)
	return row.UserId
}
func (orm *OrmLaravelAdminUserPermissions) WhereUserIdIn(val []int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`user_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUserIdGt(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`user_id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUserIdGte(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`user_id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUserIdLt(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`user_id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUserIdLte(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`user_id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WherePermissionId(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`permission_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WherePermissionIdIn(val []int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`permission_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WherePermissionIdNe(val int32) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`permission_id` <> ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereCreatedAt(val database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUpdatedAt(val database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminUserPermissions) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminUserPermissions {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminUserPermissionsList []*LaravelAdminUserPermissions

func (l LaravelAdminUserPermissionsList) GetUserIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.UserId)
	}
	return got
}
func (l LaravelAdminUserPermissionsList) GetUserIdMap() map[int32]*LaravelAdminUserPermissions {
	got := make(map[int32]*LaravelAdminUserPermissions)
	for _, val := range l {
		got[val.UserId] = val
	}
	return got
}
func (l LaravelAdminUserPermissionsList) GetPermissionIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.PermissionId)
	}
	return got
}
func (l LaravelAdminUserPermissionsList) GetPermissionIdMap() map[int32]*LaravelAdminUserPermissions {
	got := make(map[int32]*LaravelAdminUserPermissions)
	for _, val := range l {
		got[val.PermissionId] = val
	}
	return got
}
