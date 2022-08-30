package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminRolePermissions struct {
	RoleId       int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_permissions_role_id_permission_id_index,class:BTREE"`       //
	PermissionId int32          `gorm:"column:permission_id;type:int(11);index:laravel_admin_role_permissions_role_id_permission_id_index,class:BTREE"` //
	CreatedAt    *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                                                  //
	UpdatedAt    *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                                                  //
}

func (receiver *LaravelAdminRolePermissions) TableName() string {
	return "laravel_admin_role_permissions"
}

type OrmLaravelAdminRolePermissions struct {
	db *gorm.DB
}

func NewOrmLaravelAdminRolePermissions() *OrmLaravelAdminRolePermissions {
	orm := &OrmLaravelAdminRolePermissions{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminRolePermissions{})
	return orm
}

func (orm *OrmLaravelAdminRolePermissions) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminRolePermissions) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminRolePermissions) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminRolePermissions) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminRolePermissions) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminRolePermissions) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminRolePermissions) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminRolePermissions) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminRolePermissions) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminRolePermissions) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminRolePermissions) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminRolePermissions) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminRolePermissions) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminRolePermissions) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminRolePermissions) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminRolePermissions) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminRolePermissions) Insert(row *LaravelAdminRolePermissions) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminRolePermissions) Inserts(rows []*LaravelAdminRolePermissions) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminRolePermissions) Order(value interface{}) *OrmLaravelAdminRolePermissions {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminRolePermissions) Limit(limit int) *OrmLaravelAdminRolePermissions {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminRolePermissions) Offset(offset int) *OrmLaravelAdminRolePermissions {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminRolePermissions) Get() LaravelAdminRolePermissionsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminRolePermissions) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminRolePermissions) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminRolePermissions{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminRolePermissions) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_role_permissions")
}

func (orm *OrmLaravelAdminRolePermissions) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminRolePermissions) First(conds ...interface{}) (*LaravelAdminRolePermissions, bool) {
	dest := &LaravelAdminRolePermissions{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminRolePermissions) Take(conds ...interface{}) (*LaravelAdminRolePermissions, int64) {
	dest := &LaravelAdminRolePermissions{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminRolePermissions) Last(conds ...interface{}) (*LaravelAdminRolePermissions, int64) {
	dest := &LaravelAdminRolePermissions{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminRolePermissions) Find(conds ...interface{}) (LaravelAdminRolePermissionsList, int64) {
	list := make([]*LaravelAdminRolePermissions, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminRolePermissions) Paginate(page int, limit int) (LaravelAdminRolePermissionsList, int64) {
	var total int64
	list := make([]*LaravelAdminRolePermissions, 0)
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
func (orm *OrmLaravelAdminRolePermissions) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminRolePermissions) FirstOrInit(dest *LaravelAdminRolePermissions, conds ...interface{}) (*LaravelAdminRolePermissions, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminRolePermissions) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRolePermissions) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRolePermissions) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminRolePermissions) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminRolePermissions) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminRolePermissions) Where(query interface{}, args ...interface{}) *OrmLaravelAdminRolePermissions {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminRolePermissions) And(fuc func(orm *OrmLaravelAdminRolePermissions)) *OrmLaravelAdminRolePermissions {
	ormAnd := NewOrmLaravelAdminRolePermissions()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminRolePermissions) Or(fuc func(orm *OrmLaravelAdminRolePermissions)) *OrmLaravelAdminRolePermissions {
	ormOr := NewOrmLaravelAdminRolePermissions()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminRolePermissions) WhereRoleId(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) InsertGetRoleId(row *LaravelAdminRolePermissions) int32 {
	orm.db.Create(row)
	return row.RoleId
}
func (orm *OrmLaravelAdminRolePermissions) WhereRoleIdIn(val []int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereRoleIdGt(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereRoleIdGte(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereRoleIdLt(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereRoleIdLte(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WherePermissionId(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`permission_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WherePermissionIdIn(val []int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`permission_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WherePermissionIdNe(val int32) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`permission_id` <> ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereCreatedAt(val database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereUpdatedAt(val database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRolePermissions) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminRolePermissions {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminRolePermissionsList []*LaravelAdminRolePermissions

func (l LaravelAdminRolePermissionsList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l LaravelAdminRolePermissionsList) GetRoleIdMap() map[int32]*LaravelAdminRolePermissions {
	got := make(map[int32]*LaravelAdminRolePermissions)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l LaravelAdminRolePermissionsList) GetPermissionIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.PermissionId)
	}
	return got
}
func (l LaravelAdminRolePermissionsList) GetPermissionIdMap() map[int32]*LaravelAdminRolePermissions {
	got := make(map[int32]*LaravelAdminRolePermissions)
	for _, val := range l {
		got[val.PermissionId] = val
	}
	return got
}
