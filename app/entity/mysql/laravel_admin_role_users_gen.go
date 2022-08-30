package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminRoleUsers struct {
	RoleId    int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_users_role_id_user_id_index,class:BTREE"` //
	UserId    int32          `gorm:"column:user_id;type:int(11);index:laravel_admin_role_users_role_id_user_id_index,class:BTREE"` //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                                //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                                //
}

func (receiver *LaravelAdminRoleUsers) TableName() string {
	return "laravel_admin_role_users"
}

type OrmLaravelAdminRoleUsers struct {
	db *gorm.DB
}

func NewOrmLaravelAdminRoleUsers() *OrmLaravelAdminRoleUsers {
	orm := &OrmLaravelAdminRoleUsers{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminRoleUsers{})
	return orm
}

func (orm *OrmLaravelAdminRoleUsers) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminRoleUsers) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminRoleUsers) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminRoleUsers) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminRoleUsers) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminRoleUsers) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminRoleUsers) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminRoleUsers) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminRoleUsers) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminRoleUsers) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminRoleUsers) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminRoleUsers) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminRoleUsers) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminRoleUsers) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminRoleUsers) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminRoleUsers) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminRoleUsers) Insert(row *LaravelAdminRoleUsers) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminRoleUsers) Inserts(rows []*LaravelAdminRoleUsers) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminRoleUsers) Order(value interface{}) *OrmLaravelAdminRoleUsers {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminRoleUsers) Limit(limit int) *OrmLaravelAdminRoleUsers {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminRoleUsers) Offset(offset int) *OrmLaravelAdminRoleUsers {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminRoleUsers) Get() LaravelAdminRoleUsersList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminRoleUsers) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminRoleUsers) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminRoleUsers{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminRoleUsers) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_role_users")
}

func (orm *OrmLaravelAdminRoleUsers) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminRoleUsers) First(conds ...interface{}) (*LaravelAdminRoleUsers, bool) {
	dest := &LaravelAdminRoleUsers{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminRoleUsers) Take(conds ...interface{}) (*LaravelAdminRoleUsers, int64) {
	dest := &LaravelAdminRoleUsers{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminRoleUsers) Last(conds ...interface{}) (*LaravelAdminRoleUsers, int64) {
	dest := &LaravelAdminRoleUsers{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminRoleUsers) Find(conds ...interface{}) (LaravelAdminRoleUsersList, int64) {
	list := make([]*LaravelAdminRoleUsers, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminRoleUsers) Paginate(page int, limit int) (LaravelAdminRoleUsersList, int64) {
	var total int64
	list := make([]*LaravelAdminRoleUsers, 0)
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
func (orm *OrmLaravelAdminRoleUsers) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminRoleUsers) FirstOrInit(dest *LaravelAdminRoleUsers, conds ...interface{}) (*LaravelAdminRoleUsers, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminRoleUsers) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRoleUsers) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRoleUsers) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminRoleUsers) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminRoleUsers) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminRoleUsers) Where(query interface{}, args ...interface{}) *OrmLaravelAdminRoleUsers {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminRoleUsers) And(fuc func(orm *OrmLaravelAdminRoleUsers)) *OrmLaravelAdminRoleUsers {
	ormAnd := NewOrmLaravelAdminRoleUsers()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminRoleUsers) Or(fuc func(orm *OrmLaravelAdminRoleUsers)) *OrmLaravelAdminRoleUsers {
	ormOr := NewOrmLaravelAdminRoleUsers()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminRoleUsers) WhereRoleId(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) InsertGetRoleId(row *LaravelAdminRoleUsers) int32 {
	orm.db.Create(row)
	return row.RoleId
}
func (orm *OrmLaravelAdminRoleUsers) WhereRoleIdIn(val []int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereRoleIdGt(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereRoleIdGte(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereRoleIdLt(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereRoleIdLte(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUserId(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`user_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUserIdIn(val []int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`user_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUserIdNe(val int32) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`user_id` <> ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereCreatedAt(val database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUpdatedAt(val database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleUsers) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminRoleUsers {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminRoleUsersList []*LaravelAdminRoleUsers

func (l LaravelAdminRoleUsersList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l LaravelAdminRoleUsersList) GetRoleIdMap() map[int32]*LaravelAdminRoleUsers {
	got := make(map[int32]*LaravelAdminRoleUsers)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l LaravelAdminRoleUsersList) GetUserIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.UserId)
	}
	return got
}
func (l LaravelAdminRoleUsersList) GetUserIdMap() map[int32]*LaravelAdminRoleUsers {
	got := make(map[int32]*LaravelAdminRoleUsers)
	for _, val := range l {
		got[val.UserId] = val
	}
	return got
}
