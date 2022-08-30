package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type LaravelAdminOperationLog struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey"`                                //
	UserId    int32          `gorm:"column:user_id;type:int(11);index:laravel_admin_operation_log_user_id_index,class:BTREE"` //
	Path      string         `gorm:"column:path;type:varchar(255)"`                                                           //
	Method    string         `gorm:"column:method;type:varchar(10)"`                                                          //
	Ip        string         `gorm:"column:ip;type:varchar(255)"`                                                             //
	Input     string         `gorm:"column:input;type:text"`                                                                  //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                           //
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                                           //
}

func (receiver *LaravelAdminOperationLog) TableName() string {
	return "laravel_admin_operation_log"
}

type OrmLaravelAdminOperationLog struct {
	db *gorm.DB
}

func NewOrmLaravelAdminOperationLog() *OrmLaravelAdminOperationLog {
	orm := &OrmLaravelAdminOperationLog{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&LaravelAdminOperationLog{})
	return orm
}

func (orm *OrmLaravelAdminOperationLog) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmLaravelAdminOperationLog) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmLaravelAdminOperationLog) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmLaravelAdminOperationLog) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmLaravelAdminOperationLog) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmLaravelAdminOperationLog) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmLaravelAdminOperationLog) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmLaravelAdminOperationLog) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmLaravelAdminOperationLog) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmLaravelAdminOperationLog) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmLaravelAdminOperationLog) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmLaravelAdminOperationLog) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmLaravelAdminOperationLog) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmLaravelAdminOperationLog) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmLaravelAdminOperationLog) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmLaravelAdminOperationLog) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmLaravelAdminOperationLog) Insert(row *LaravelAdminOperationLog) error {
	return orm.db.Create(row).Error
}

func (orm *OrmLaravelAdminOperationLog) Inserts(rows []*LaravelAdminOperationLog) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmLaravelAdminOperationLog) Order(value interface{}) *OrmLaravelAdminOperationLog {
	orm.db.Order(value)
	return orm
}

func (orm *OrmLaravelAdminOperationLog) Limit(limit int) *OrmLaravelAdminOperationLog {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmLaravelAdminOperationLog) Offset(offset int) *OrmLaravelAdminOperationLog {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmLaravelAdminOperationLog) Get() LaravelAdminOperationLogList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmLaravelAdminOperationLog) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmLaravelAdminOperationLog) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&LaravelAdminOperationLog{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmLaravelAdminOperationLog) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM laravel_admin_operation_log")
}

func (orm *OrmLaravelAdminOperationLog) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmLaravelAdminOperationLog) First(conds ...interface{}) (*LaravelAdminOperationLog, bool) {
	dest := &LaravelAdminOperationLog{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmLaravelAdminOperationLog) Take(conds ...interface{}) (*LaravelAdminOperationLog, int64) {
	dest := &LaravelAdminOperationLog{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmLaravelAdminOperationLog) Last(conds ...interface{}) (*LaravelAdminOperationLog, int64) {
	dest := &LaravelAdminOperationLog{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmLaravelAdminOperationLog) Find(conds ...interface{}) (LaravelAdminOperationLogList, int64) {
	list := make([]*LaravelAdminOperationLog, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmLaravelAdminOperationLog) Paginate(page int, limit int) (LaravelAdminOperationLogList, int64) {
	var total int64
	list := make([]*LaravelAdminOperationLog, 0)
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
func (orm *OrmLaravelAdminOperationLog) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmLaravelAdminOperationLog) FirstOrInit(dest *LaravelAdminOperationLog, conds ...interface{}) (*LaravelAdminOperationLog, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmLaravelAdminOperationLog) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminOperationLog) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmLaravelAdminOperationLog) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmLaravelAdminOperationLog) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmLaravelAdminOperationLog) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmLaravelAdminOperationLog) Where(query interface{}, args ...interface{}) *OrmLaravelAdminOperationLog {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmLaravelAdminOperationLog) And(fuc func(orm *OrmLaravelAdminOperationLog)) *OrmLaravelAdminOperationLog {
	ormAnd := NewOrmLaravelAdminOperationLog()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmLaravelAdminOperationLog) Or(fuc func(orm *OrmLaravelAdminOperationLog)) *OrmLaravelAdminOperationLog {
	ormOr := NewOrmLaravelAdminOperationLog()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmLaravelAdminOperationLog) WhereId(val uint32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) InsertGetId(row *LaravelAdminOperationLog) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmLaravelAdminOperationLog) WhereIdIn(val []uint32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereIdGt(val uint32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereIdGte(val uint32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereIdLt(val uint32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereIdLte(val uint32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUserId(val int32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`user_id` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) InsertGetUserId(row *LaravelAdminOperationLog) int32 {
	orm.db.Create(row)
	return row.UserId
}
func (orm *OrmLaravelAdminOperationLog) WhereUserIdIn(val []int32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`user_id` IN ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUserIdGt(val int32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`user_id` > ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUserIdGte(val int32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`user_id` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUserIdLt(val int32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`user_id` < ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUserIdLte(val int32) *OrmLaravelAdminOperationLog {
	orm.db.Where("`user_id` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WherePath(val string) *OrmLaravelAdminOperationLog {
	orm.db.Where("`path` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereMethod(val string) *OrmLaravelAdminOperationLog {
	orm.db.Where("`method` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereIp(val string) *OrmLaravelAdminOperationLog {
	orm.db.Where("`ip` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereInput(val string) *OrmLaravelAdminOperationLog {
	orm.db.Where("`input` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereCreatedAt(val database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereCreatedAtLte(val database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereCreatedAtGte(val database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUpdatedAt(val database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUpdatedAtLte(val database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmLaravelAdminOperationLog) WhereUpdatedAtGte(val database.Time) *OrmLaravelAdminOperationLog {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type LaravelAdminOperationLogList []*LaravelAdminOperationLog

func (l LaravelAdminOperationLogList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l LaravelAdminOperationLogList) GetIdMap() map[uint32]*LaravelAdminOperationLog {
	got := make(map[uint32]*LaravelAdminOperationLog)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l LaravelAdminOperationLogList) GetUserIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.UserId)
	}
	return got
}
func (l LaravelAdminOperationLogList) GetUserIdMap() map[int32]*LaravelAdminOperationLog {
	got := make(map[int32]*LaravelAdminOperationLog)
	for _, val := range l {
		got[val.UserId] = val
	}
	return got
}
