package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type AdminRoleMenu struct {
	RoleId    int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_menu_role_id_menu_id_index,class:BTREE"` //
	MenuId    int32          `gorm:"column:menu_id;type:int(11);index:laravel_admin_role_menu_role_id_menu_id_index,class:BTREE"` //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                               //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                               //
}

func (receiver *AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}

type OrmAdminRoleMenu struct {
	db *gorm.DB
}

func NewOrmAdminRoleMenu() *OrmAdminRoleMenu {
	orm := &OrmAdminRoleMenu{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&AdminRoleMenu{})
	return orm
}

func (orm *OrmAdminRoleMenu) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmAdminRoleMenu) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminRoleMenu) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminRoleMenu) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminRoleMenu) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminRoleMenu) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminRoleMenu) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminRoleMenu) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminRoleMenu) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminRoleMenu) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminRoleMenu) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminRoleMenu) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminRoleMenu) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminRoleMenu) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminRoleMenu) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminRoleMenu) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminRoleMenu) Insert(row *AdminRoleMenu) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminRoleMenu) Inserts(rows []*AdminRoleMenu) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminRoleMenu) Order(value interface{}) *OrmAdminRoleMenu {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminRoleMenu) Limit(limit int) *OrmAdminRoleMenu {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminRoleMenu) Offset(offset int) *OrmAdminRoleMenu {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminRoleMenu) Get() AdminRoleMenuList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminRoleMenu) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminRoleMenu) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminRoleMenu{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminRoleMenu) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_role_menu")
}

func (orm *OrmAdminRoleMenu) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminRoleMenu) First(conds ...interface{}) (*AdminRoleMenu, bool) {
	dest := &AdminRoleMenu{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminRoleMenu) Take(conds ...interface{}) (*AdminRoleMenu, int64) {
	dest := &AdminRoleMenu{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminRoleMenu) Last(conds ...interface{}) (*AdminRoleMenu, int64) {
	dest := &AdminRoleMenu{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminRoleMenu) Find(conds ...interface{}) (AdminRoleMenuList, int64) {
	list := make([]*AdminRoleMenu, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminRoleMenu) Paginate(page int, limit int) (AdminRoleMenuList, int64) {
	var total int64
	list := make([]*AdminRoleMenu, 0)
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
func (orm *OrmAdminRoleMenu) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminRoleMenu) FirstOrInit(dest *AdminRoleMenu, conds ...interface{}) (*AdminRoleMenu, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminRoleMenu) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleMenu) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleMenu) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminRoleMenu) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminRoleMenu) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminRoleMenu) Where(query interface{}, args ...interface{}) *OrmAdminRoleMenu {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminRoleMenu) And(fuc func(orm *OrmAdminRoleMenu)) *OrmAdminRoleMenu {
	ormAnd := NewOrmAdminRoleMenu()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminRoleMenu) Or(fuc func(orm *OrmAdminRoleMenu)) *OrmAdminRoleMenu {
	ormOr := NewOrmAdminRoleMenu()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmAdminRoleMenu) WhereRoleId(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) InsertGetRoleId(row *AdminRoleMenu) int32 {
	orm.db.Create(row)
	return row.RoleId
}
func (orm *OrmAdminRoleMenu) WhereRoleIdIn(val []int32) *OrmAdminRoleMenu {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereRoleIdGt(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereRoleIdGte(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereRoleIdLt(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereRoleIdLte(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereMenuId(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`menu_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereMenuIdIn(val []int32) *OrmAdminRoleMenu {
	orm.db.Where("`menu_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereMenuIdNe(val int32) *OrmAdminRoleMenu {
	orm.db.Where("`menu_id` <> ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereCreatedAt(val database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereCreatedAtLte(val database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereCreatedAtGte(val database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereUpdatedAt(val database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereUpdatedAtLte(val database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenu) WhereUpdatedAtGte(val database.Time) *OrmAdminRoleMenu {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminRoleMenuList []*AdminRoleMenu

func (l AdminRoleMenuList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l AdminRoleMenuList) GetRoleIdMap() map[int32]*AdminRoleMenu {
	got := make(map[int32]*AdminRoleMenu)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l AdminRoleMenuList) GetMenuIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.MenuId)
	}
	return got
}
func (l AdminRoleMenuList) GetMenuIdMap() map[int32]*AdminRoleMenu {
	got := make(map[int32]*AdminRoleMenu)
	for _, val := range l {
		got[val.MenuId] = val
	}
	return got
}
