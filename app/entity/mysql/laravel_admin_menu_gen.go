package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminMenu struct {
	Id         uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey"` //
	ParentId   int32          `gorm:"column:parent_id;type:int(11);default:0"`                  //
	Order      int32          `gorm:"column:order;type:int(11);default:0"`                      //
	Title      string         `gorm:"column:title;type:varchar(50)"`                            //
	Icon       string         `gorm:"column:icon;type:varchar(50)"`                             //
	Uri        *string        `gorm:"column:uri;type:varchar(255);default:NULL"`                //
	Permission *string        `gorm:"column:permission;type:varchar(255);default:NULL"`         //
	CreatedAt  *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`            //
	UpdatedAt  *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`            //
}

func (receiver *LaravelAdminMenu) TableName() string {
	return "laravel_admin_menu"
}

type OrmLaravelAdminMenu struct {
	db *gorm.DB
}

func NewOrmLaravelAdminMenu() *OrmLaravelAdminMenu {
	orm := &OrmLaravelAdminMenu{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminMenu{})
	return orm
}

func (orm *OrmLaravelAdminMenu) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminMenu) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminMenu) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminMenu) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminMenu) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminMenu) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminMenu) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminMenu) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminMenu) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminMenu) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminMenu) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminMenu) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminMenu) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminMenu) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminMenu) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminMenu) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminMenu) Insert(row *LaravelAdminMenu) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminMenu) Inserts(rows []*LaravelAdminMenu) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminMenu) Order(value interface{}) *OrmLaravelAdminMenu {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminMenu) Limit(limit int) *OrmLaravelAdminMenu {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminMenu) Offset(offset int) *OrmLaravelAdminMenu {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminMenu) Get() LaravelAdminMenuList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminMenu) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminMenu) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminMenu{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminMenu) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_menu")
}

func (orm *OrmLaravelAdminMenu) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminMenu) First(conds ...interface{}) (*LaravelAdminMenu, bool) {
	dest := &LaravelAdminMenu{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminMenu) Take(conds ...interface{}) (*LaravelAdminMenu, int64) {
	dest := &LaravelAdminMenu{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminMenu) Last(conds ...interface{}) (*LaravelAdminMenu, int64) {
	dest := &LaravelAdminMenu{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminMenu) Find(conds ...interface{}) (LaravelAdminMenuList, int64) {
	list := make([]*LaravelAdminMenu, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminMenu) Paginate(page int, limit int) (LaravelAdminMenuList, int64) {
	var total int64
	list := make([]*LaravelAdminMenu, 0)
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
func (orm *OrmLaravelAdminMenu) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminMenu) FirstOrInit(dest *LaravelAdminMenu, conds ...interface{}) (*LaravelAdminMenu, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminMenu) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminMenu) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminMenu) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminMenu) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminMenu) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminMenu) Where(query interface{}, args ...interface{}) *OrmLaravelAdminMenu {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminMenu) And(fuc func(orm *OrmLaravelAdminMenu)) *OrmLaravelAdminMenu {
	ormAnd := NewOrmLaravelAdminMenu()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminMenu) Or(fuc func(orm *OrmLaravelAdminMenu)) *OrmLaravelAdminMenu {
	ormOr := NewOrmLaravelAdminMenu()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminMenu) WhereId(val uint32) *OrmLaravelAdminMenu {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) InsertGetId(row *LaravelAdminMenu) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmLaravelAdminMenu) WhereIdIn(val []uint32) *OrmLaravelAdminMenu {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereIdGt(val uint32) *OrmLaravelAdminMenu {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereIdGte(val uint32) *OrmLaravelAdminMenu {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereIdLt(val uint32) *OrmLaravelAdminMenu {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereIdLte(val uint32) *OrmLaravelAdminMenu {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereParentId(val int32) *OrmLaravelAdminMenu {
	orm.db.Where("`parent_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereParentIdIn(val []int32) *OrmLaravelAdminMenu {
	orm.db.Where("`parent_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereParentIdNe(val int32) *OrmLaravelAdminMenu {
	orm.db.Where("`parent_id` <> ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereOrder(val int32) *OrmLaravelAdminMenu {
	orm.db.Where("`order` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereTitle(val string) *OrmLaravelAdminMenu {
	orm.db.Where("`title` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereIcon(val string) *OrmLaravelAdminMenu {
	orm.db.Where("`icon` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereUri(val string) *OrmLaravelAdminMenu {
	orm.db.Where("`uri` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WherePermission(val string) *OrmLaravelAdminMenu {
	orm.db.Where("`permission` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereCreatedAt(val database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereUpdatedAt(val database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminMenu) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminMenu {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminMenuList []*LaravelAdminMenu

func (l LaravelAdminMenuList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l LaravelAdminMenuList) GetIdMap() map[uint32]*LaravelAdminMenu {
	got := make(map[uint32]*LaravelAdminMenu)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l LaravelAdminMenuList) GetParentIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.ParentId)
	}
	return got
}
func (l LaravelAdminMenuList) GetParentIdMap() map[int32]*LaravelAdminMenu {
	got := make(map[int32]*LaravelAdminMenu)
	for _, val := range l {
		got[val.ParentId] = val
	}
	return got
}
