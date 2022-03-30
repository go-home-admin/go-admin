package demo

import (
	"database/sql"
	"github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MysqlTableName struct {
	Id       int64         `gorm:"primaryKey;column:id"`        //
	Date     database.Time `gorm:"column:date"`                 // 时间
	Guard    string        `gorm:"column:guard"`                // guard
	Ip       string        `gorm:"column:ip"`                   // ip
	Uri      string        `gorm:"primaryKey;column:uri"`       // 请求地址
	Params   string        `gorm:"column:params"`               // 参数
	UserId   int64         `gorm:"column:user_id"`              // uid
	User     string        `gorm:"column:user"`                 // uid
	HttpType string        `gorm:"primaryKey;column:http_type"` // http_type
}

func (receiver *MysqlTableName) TableName() string {
	return "action_logs"
}

type OrmMysqlTableName struct {
	db *gorm.DB
}

func NewOrmTableName() *OrmMysqlTableName {
	orm := &OrmMysqlTableName{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB)
	return orm
}

func (orm *OrmMysqlTableName) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmMysqlTableName) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmMysqlTableName) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmMysqlTableName) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmMysqlTableName) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmMysqlTableName) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmMysqlTableName) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmMysqlTableName) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmMysqlTableName) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmMysqlTableName) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmMysqlTableName) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmMysqlTableName) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmMysqlTableName) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmMysqlTableName) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmMysqlTableName) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmMysqlTableName) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

type TableNameList []*MysqlTableName

func (orm *OrmMysqlTableName) Insert(row *MysqlTableName) *gorm.DB {
	return orm.db.Create(row)
}

func (orm *OrmMysqlTableName) Inserts(rows []*MysqlTableName) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmMysqlTableName) Limit(limit int) *OrmMysqlTableName {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmMysqlTableName) Offset(offset int) *OrmMysqlTableName {
	orm.db.Offset(offset)
	return orm
}

func (orm *OrmMysqlTableName) Get() (TableNameList, int64) {
	return orm.Find()
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmMysqlTableName) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Model(&MysqlTableName{}).Pluck(column, dest)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (orm *OrmMysqlTableName) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.Model(&MysqlTableName{}).Delete(value, conds...)
}

func (orm *OrmMysqlTableName) Count() int64 {
	var count int64
	orm.db.Model(&MysqlTableName{}).Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmMysqlTableName) First(conds ...interface{}) (*MysqlTableName, int64) {
	dest := &MysqlTableName{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmMysqlTableName) Take(conds ...interface{}) (*MysqlTableName, int64) {
	dest := &MysqlTableName{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmMysqlTableName) Last(conds ...interface{}) (*MysqlTableName, int64) {
	dest := &MysqlTableName{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmMysqlTableName) Find(conds ...interface{}) (TableNameList, int64) {
	list := make([]*MysqlTableName, 0)
	tx := orm.db.Model(&MysqlTableName{}).Find(list, conds...)
	if tx.Error != nil {
		log.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// FindInBatches find records in batches
func (orm *OrmMysqlTableName) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmMysqlTableName) FirstOrInit(dest *MysqlTableName, conds ...interface{}) (*MysqlTableName, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmMysqlTableName) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMysqlTableName) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Model(&MysqlTableName{}).Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMysqlTableName) Updates(values interface{}) *gorm.DB {
	return orm.db.Model(&MysqlTableName{}).Updates(values)
}

func (orm *OrmMysqlTableName) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.Model(&MysqlTableName{}).UpdateColumn(column, value)
}

func (orm *OrmMysqlTableName) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.Model(&MysqlTableName{}).UpdateColumns(values)
}

func (orm *OrmMysqlTableName) Where(query interface{}, args ...interface{}) *OrmMysqlTableName {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmMysqlTableName) And(fuc func(orm *OrmMysqlTableName)) *OrmMysqlTableName {
	fuc(orm)
	orm.db.Where(orm.db)
	return orm
}

func (orm *OrmMysqlTableName) Or(fuc func(orm *OrmMysqlTableName)) *OrmMysqlTableName {
	fuc(orm)
	orm.db.Or(orm.db)
	return orm
}
