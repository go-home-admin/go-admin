package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminRoleMenu struct {
	RoleId    int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_menu_role_id_menu_id_index,class:BTREE"` //
	MenuId    int32          `gorm:"column:menu_id;type:int(11);index:laravel_admin_role_menu_role_id_menu_id_index,class:BTREE"` //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                               //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                               //
}

func (receiver *LaravelAdminRoleMenu) TableName() string {
	return "laravel_admin_role_menu"
}

type OrmLaravelAdminRoleMenu struct {
	db *gorm.DB
}

func NewOrmLaravelAdminRoleMenu() *OrmLaravelAdminRoleMenu {
	orm := &OrmLaravelAdminRoleMenu{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminRoleMenu{})
	return orm
}

func (orm *OrmLaravelAdminRoleMenu) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminRoleMenu) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminRoleMenu) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminRoleMenu) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminRoleMenu) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminRoleMenu) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminRoleMenu) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminRoleMenu) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminRoleMenu) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminRoleMenu) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminRoleMenu) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminRoleMenu) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminRoleMenu) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminRoleMenu) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminRoleMenu) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminRoleMenu) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminRoleMenu) Insert(row *LaravelAdminRoleMenu) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminRoleMenu) Inserts(rows []*LaravelAdminRoleMenu) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminRoleMenu) Order(value interface{}) *OrmLaravelAdminRoleMenu {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminRoleMenu) Limit(limit int) *OrmLaravelAdminRoleMenu {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminRoleMenu) Offset(offset int) *OrmLaravelAdminRoleMenu {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminRoleMenu) Get() LaravelAdminRoleMenuList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminRoleMenu) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminRoleMenu) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminRoleMenu{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminRoleMenu) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_role_menu")
}

func (orm *OrmLaravelAdminRoleMenu) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminRoleMenu) First(conds ...interface{}) (*LaravelAdminRoleMenu, bool) {
	dest := &LaravelAdminRoleMenu{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminRoleMenu) Take(conds ...interface{}) (*LaravelAdminRoleMenu, int64) {
	dest := &LaravelAdminRoleMenu{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminRoleMenu) Last(conds ...interface{}) (*LaravelAdminRoleMenu, int64) {
	dest := &LaravelAdminRoleMenu{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminRoleMenu) Find(conds ...interface{}) (LaravelAdminRoleMenuList, int64) {
	list := make([]*LaravelAdminRoleMenu, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminRoleMenu) Paginate(page int, limit int) (LaravelAdminRoleMenuList, int64) {
	var total int64
	list := make([]*LaravelAdminRoleMenu, 0)
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
func (orm *OrmLaravelAdminRoleMenu) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminRoleMenu) FirstOrInit(dest *LaravelAdminRoleMenu, conds ...interface{}) (*LaravelAdminRoleMenu, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminRoleMenu) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRoleMenu) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminRoleMenu) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminRoleMenu) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminRoleMenu) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminRoleMenu) Where(query interface{}, args ...interface{}) *OrmLaravelAdminRoleMenu {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminRoleMenu) And(fuc func(orm *OrmLaravelAdminRoleMenu)) *OrmLaravelAdminRoleMenu {
	ormAnd := NewOrmLaravelAdminRoleMenu()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminRoleMenu) Or(fuc func(orm *OrmLaravelAdminRoleMenu)) *OrmLaravelAdminRoleMenu {
	ormOr := NewOrmLaravelAdminRoleMenu()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminRoleMenu) WhereRoleId(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) InsertGetRoleId(row *LaravelAdminRoleMenu) int32 {
	orm.db.Create(row)
	return row.RoleId
}
func (orm *OrmLaravelAdminRoleMenu) WhereRoleIdIn(val []int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereRoleIdGt(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereRoleIdGte(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereRoleIdLt(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereRoleIdLte(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereMenuId(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`menu_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereMenuIdIn(val []int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`menu_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereMenuIdNe(val int32) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`menu_id` <> ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereCreatedAt(val database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereUpdatedAt(val database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminRoleMenu) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminRoleMenu {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminRoleMenuList []*LaravelAdminRoleMenu

func (l LaravelAdminRoleMenuList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l LaravelAdminRoleMenuList) GetRoleIdMap() map[int32]*LaravelAdminRoleMenu {
	got := make(map[int32]*LaravelAdminRoleMenu)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l LaravelAdminRoleMenuList) GetMenuIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.MenuId)
	}
	return got
}
func (l LaravelAdminRoleMenuList) GetMenuIdMap() map[int32]*LaravelAdminRoleMenu {
	got := make(map[int32]*LaravelAdminRoleMenu)
	for _, val := range l {
		got[val.MenuId] = val
	}
	return got
}
