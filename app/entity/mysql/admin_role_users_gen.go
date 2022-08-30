package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type AdminRoleUsers struct {
	RoleId    int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_users_role_id_user_id_index,class:BTREE"` //
	UserId    int32          `gorm:"column:user_id;type:int(11);index:laravel_admin_role_users_role_id_user_id_index,class:BTREE"` //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                                //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                                //
}

func (receiver *AdminRoleUsers) TableName() string {
	return "admin_role_users"
}

type OrmAdminRoleUsers struct {
	db *gorm.DB
}

func NewOrmAdminRoleUsers() *OrmAdminRoleUsers {
	orm := &OrmAdminRoleUsers{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&AdminRoleUsers{})
	return orm
}

func (orm *OrmAdminRoleUsers) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmAdminRoleUsers) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminRoleUsers) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminRoleUsers) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminRoleUsers) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminRoleUsers) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminRoleUsers) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminRoleUsers) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminRoleUsers) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminRoleUsers) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminRoleUsers) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminRoleUsers) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminRoleUsers) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminRoleUsers) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminRoleUsers) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminRoleUsers) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminRoleUsers) Insert(row *AdminRoleUsers) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminRoleUsers) Inserts(rows []*AdminRoleUsers) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminRoleUsers) Order(value interface{}) *OrmAdminRoleUsers {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminRoleUsers) Limit(limit int) *OrmAdminRoleUsers {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminRoleUsers) Offset(offset int) *OrmAdminRoleUsers {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminRoleUsers) Get() AdminRoleUsersList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminRoleUsers) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminRoleUsers) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminRoleUsers{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminRoleUsers) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_role_users")
}

func (orm *OrmAdminRoleUsers) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminRoleUsers) First(conds ...interface{}) (*AdminRoleUsers, bool) {
	dest := &AdminRoleUsers{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminRoleUsers) Take(conds ...interface{}) (*AdminRoleUsers, int64) {
	dest := &AdminRoleUsers{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminRoleUsers) Last(conds ...interface{}) (*AdminRoleUsers, int64) {
	dest := &AdminRoleUsers{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminRoleUsers) Find(conds ...interface{}) (AdminRoleUsersList, int64) {
	list := make([]*AdminRoleUsers, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminRoleUsers) Paginate(page int, limit int) (AdminRoleUsersList, int64) {
	var total int64
	list := make([]*AdminRoleUsers, 0)
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
func (orm *OrmAdminRoleUsers) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminRoleUsers) FirstOrInit(dest *AdminRoleUsers, conds ...interface{}) (*AdminRoleUsers, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminRoleUsers) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleUsers) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleUsers) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminRoleUsers) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminRoleUsers) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminRoleUsers) Where(query interface{}, args ...interface{}) *OrmAdminRoleUsers {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminRoleUsers) And(fuc func(orm *OrmAdminRoleUsers)) *OrmAdminRoleUsers {
	ormAnd := NewOrmAdminRoleUsers()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminRoleUsers) Or(fuc func(orm *OrmAdminRoleUsers)) *OrmAdminRoleUsers {
	ormOr := NewOrmAdminRoleUsers()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmAdminRoleUsers) WhereRoleId(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) InsertGetRoleId(row *AdminRoleUsers) int32 {
	orm.db.Create(row)
	return row.RoleId
}
func (orm *OrmAdminRoleUsers) WhereRoleIdIn(val []int32) *OrmAdminRoleUsers {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereRoleIdGt(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereRoleIdGte(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereRoleIdLt(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereRoleIdLte(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUserId(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`user_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUserIdIn(val []int32) *OrmAdminRoleUsers {
	orm.db.Where("`user_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUserIdNe(val int32) *OrmAdminRoleUsers {
	orm.db.Where("`user_id` <> ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereCreatedAt(val database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereCreatedAtLte(val database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereCreatedAtGte(val database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUpdatedAt(val database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUpdatedAtLte(val database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsers) WhereUpdatedAtGte(val database.Time) *OrmAdminRoleUsers {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminRoleUsersList []*AdminRoleUsers

func (l AdminRoleUsersList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l AdminRoleUsersList) GetRoleIdMap() map[int32]*AdminRoleUsers {
	got := make(map[int32]*AdminRoleUsers)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l AdminRoleUsersList) GetUserIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.UserId)
	}
	return got
}
func (l AdminRoleUsersList) GetUserIdMap() map[int32]*AdminRoleUsers {
	got := make(map[int32]*AdminRoleUsers)
	for _, val := range l {
		got[val.UserId] = val
	}
	return got
}
