package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type Migrations struct {
	Id        uint32 `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey"` //
	Migration string `gorm:"column:migration;type:varchar(255)"`                       //
	Batch     int32  `gorm:"column:batch;type:int(11)"`                                //
}

func (receiver *Migrations) TableName() string {
	return "migrations"
}

type OrmMigrations struct {
	db *gorm.DB
}

func NewOrmMigrations() *OrmMigrations {
	orm := &OrmMigrations{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&Migrations{})
	return orm
}

func (orm *OrmMigrations) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmMigrations) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmMigrations) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmMigrations) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmMigrations) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmMigrations) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmMigrations) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmMigrations) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmMigrations) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmMigrations) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmMigrations) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmMigrations) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmMigrations) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmMigrations) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmMigrations) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmMigrations) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmMigrations) Insert(row *Migrations) error {
	return orm.db.Create(row).Error
}

func (orm *OrmMigrations) Inserts(rows []*Migrations) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmMigrations) Order(value interface{}) *OrmMigrations {
	orm.db.Order(value)
	return orm
}

func (orm *OrmMigrations) Limit(limit int) *OrmMigrations {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmMigrations) Offset(offset int) *OrmMigrations {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmMigrations) Get() MigrationsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmMigrations) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmMigrations) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&Migrations{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmMigrations) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM migrations")
}

func (orm *OrmMigrations) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmMigrations) First(conds ...interface{}) (*Migrations, bool) {
	dest := &Migrations{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmMigrations) Take(conds ...interface{}) (*Migrations, int64) {
	dest := &Migrations{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmMigrations) Last(conds ...interface{}) (*Migrations, int64) {
	dest := &Migrations{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmMigrations) Find(conds ...interface{}) (MigrationsList, int64) {
	list := make([]*Migrations, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmMigrations) Paginate(page int, limit int) (MigrationsList, int64) {
	var total int64
	list := make([]*Migrations, 0)
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
func (orm *OrmMigrations) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmMigrations) FirstOrInit(dest *Migrations, conds ...interface{}) (*Migrations, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmMigrations) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMigrations) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMigrations) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmMigrations) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmMigrations) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmMigrations) Where(query interface{}, args ...interface{}) *OrmMigrations {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmMigrations) And(fuc func(orm *OrmMigrations)) *OrmMigrations {
	ormAnd := NewOrmMigrations()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmMigrations) Or(fuc func(orm *OrmMigrations)) *OrmMigrations {
	ormOr := NewOrmMigrations()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmMigrations) WhereId(val uint32) *OrmMigrations {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmMigrations) InsertGetId(row *Migrations) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmMigrations) WhereIdIn(val []uint32) *OrmMigrations {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmMigrations) WhereIdGt(val uint32) *OrmMigrations {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmMigrations) WhereIdGte(val uint32) *OrmMigrations {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmMigrations) WhereIdLt(val uint32) *OrmMigrations {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmMigrations) WhereIdLte(val uint32) *OrmMigrations {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmMigrations) WhereMigration(val string) *OrmMigrations {
	orm.db.Where("`migration` = ?", val)
	return orm
}
func (orm *OrmMigrations) WhereBatch(val int32) *OrmMigrations {
	orm.db.Where("`batch` = ?", val)
	return orm
}

type MigrationsList []*Migrations

func (l MigrationsList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l MigrationsList) GetIdMap() map[uint32]*Migrations {
	got := make(map[uint32]*Migrations)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
