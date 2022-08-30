package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type FailedJobs struct {
	Id         uint64        `gorm:"column:id;autoIncrement;type:bigint(20) unsigned;primaryKey"`                   //
	Uuid       string        `gorm:"column:uuid;type:varchar(255);index:failed_jobs_uuid_unique,class:HASH,unique"` //
	Connection string        `gorm:"column:connection;type:text"`                                                   //
	Queue      string        `gorm:"column:queue;type:text"`                                                        //
	Payload    string        `gorm:"column:payload;type:longtext"`                                                  //
	Exception  string        `gorm:"column:exception;type:longtext"`                                                //
	FailedAt   database.Time `gorm:"column:failed_at;type:timestamp;default:current_timestamp()"`                   //
}

func (receiver *FailedJobs) TableName() string {
	return "failed_jobs"
}

type OrmFailedJobs struct {
	db *gorm.DB
}

func NewOrmFailedJobs() *OrmFailedJobs {
	orm := &OrmFailedJobs{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&FailedJobs{})
	return orm
}

func (orm *OrmFailedJobs) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmFailedJobs) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmFailedJobs) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmFailedJobs) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmFailedJobs) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmFailedJobs) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmFailedJobs) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmFailedJobs) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmFailedJobs) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmFailedJobs) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmFailedJobs) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmFailedJobs) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmFailedJobs) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmFailedJobs) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmFailedJobs) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmFailedJobs) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmFailedJobs) Insert(row *FailedJobs) error {
	return orm.db.Create(row).Error
}

func (orm *OrmFailedJobs) Inserts(rows []*FailedJobs) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmFailedJobs) Order(value interface{}) *OrmFailedJobs {
	orm.db.Order(value)
	return orm
}

func (orm *OrmFailedJobs) Limit(limit int) *OrmFailedJobs {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmFailedJobs) Offset(offset int) *OrmFailedJobs {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmFailedJobs) Get() FailedJobsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmFailedJobs) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmFailedJobs) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&FailedJobs{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmFailedJobs) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM failed_jobs")
}

func (orm *OrmFailedJobs) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmFailedJobs) First(conds ...interface{}) (*FailedJobs, bool) {
	dest := &FailedJobs{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmFailedJobs) Take(conds ...interface{}) (*FailedJobs, int64) {
	dest := &FailedJobs{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmFailedJobs) Last(conds ...interface{}) (*FailedJobs, int64) {
	dest := &FailedJobs{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmFailedJobs) Find(conds ...interface{}) (FailedJobsList, int64) {
	list := make([]*FailedJobs, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmFailedJobs) Paginate(page int, limit int) (FailedJobsList, int64) {
	var total int64
	list := make([]*FailedJobs, 0)
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
func (orm *OrmFailedJobs) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmFailedJobs) FirstOrInit(dest *FailedJobs, conds ...interface{}) (*FailedJobs, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmFailedJobs) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmFailedJobs) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmFailedJobs) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmFailedJobs) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmFailedJobs) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmFailedJobs) Where(query interface{}, args ...interface{}) *OrmFailedJobs {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmFailedJobs) And(fuc func(orm *OrmFailedJobs)) *OrmFailedJobs {
	ormAnd := NewOrmFailedJobs()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmFailedJobs) Or(fuc func(orm *OrmFailedJobs)) *OrmFailedJobs {
	ormOr := NewOrmFailedJobs()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmFailedJobs) WhereId(val uint64) *OrmFailedJobs {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) InsertGetId(row *FailedJobs) uint64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmFailedJobs) WhereIdIn(val []uint64) *OrmFailedJobs {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereIdGt(val uint64) *OrmFailedJobs {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereIdGte(val uint64) *OrmFailedJobs {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereIdLt(val uint64) *OrmFailedJobs {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereIdLte(val uint64) *OrmFailedJobs {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereUuid(val string) *OrmFailedJobs {
	orm.db.Where("`uuid` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) InsertGetUuid(row *FailedJobs) string {
	orm.db.Create(row)
	return row.Uuid
}
func (orm *OrmFailedJobs) WhereUuidIn(val []string) *OrmFailedJobs {
	orm.db.Where("`uuid` IN ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereUuidGt(val string) *OrmFailedJobs {
	orm.db.Where("`uuid` > ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereUuidGte(val string) *OrmFailedJobs {
	orm.db.Where("`uuid` >= ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereUuidLt(val string) *OrmFailedJobs {
	orm.db.Where("`uuid` < ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereUuidLte(val string) *OrmFailedJobs {
	orm.db.Where("`uuid` <= ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereConnection(val string) *OrmFailedJobs {
	orm.db.Where("`connection` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereQueue(val string) *OrmFailedJobs {
	orm.db.Where("`queue` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) WherePayload(val string) *OrmFailedJobs {
	orm.db.Where("`payload` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereException(val string) *OrmFailedJobs {
	orm.db.Where("`exception` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereFailedAt(val database.Time) *OrmFailedJobs {
	orm.db.Where("`failed_at` = ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereFailedAtBetween(begin database.Time, end database.Time) *OrmFailedJobs {
	orm.db.Where("`failed_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmFailedJobs) WhereFailedAtLte(val database.Time) *OrmFailedJobs {
	orm.db.Where("`failed_at` <= ?", val)
	return orm
}
func (orm *OrmFailedJobs) WhereFailedAtGte(val database.Time) *OrmFailedJobs {
	orm.db.Where("`failed_at` >= ?", val)
	return orm
}

type FailedJobsList []*FailedJobs

func (l FailedJobsList) GetIdList() []uint64 {
	got := make([]uint64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l FailedJobsList) GetIdMap() map[uint64]*FailedJobs {
	got := make(map[uint64]*FailedJobs)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l FailedJobsList) GetUuidList() []string {
	got := make([]string, 0)
	for _, val := range l {
		got = append(got, val.Uuid)
	}
	return got
}
func (l FailedJobsList) GetUuidMap() map[string]*FailedJobs {
	got := make(map[string]*FailedJobs)
	for _, val := range l {
		got[val.Uuid] = val
	}
	return got
}
