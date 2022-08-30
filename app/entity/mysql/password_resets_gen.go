package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type PasswordResets struct {
	Email     string         `gorm:"column:email;type:varchar(255);index:password_resets_email_index,class:BTREE"` //
	Token     string         `gorm:"column:token;type:varchar(255)"`                                               //
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                                //
}

func (receiver *PasswordResets) TableName() string {
	return "password_resets"
}

type OrmPasswordResets struct {
	db *gorm.DB
}

func NewOrmPasswordResets() *OrmPasswordResets {
	orm := &OrmPasswordResets{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&PasswordResets{})
	return orm
}

func (orm *OrmPasswordResets) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmPasswordResets) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmPasswordResets) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmPasswordResets) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmPasswordResets) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmPasswordResets) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmPasswordResets) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmPasswordResets) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmPasswordResets) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmPasswordResets) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmPasswordResets) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmPasswordResets) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmPasswordResets) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmPasswordResets) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmPasswordResets) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmPasswordResets) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmPasswordResets) Insert(row *PasswordResets) error {
	return orm.db.Create(row).Error
}

func (orm *OrmPasswordResets) Inserts(rows []*PasswordResets) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmPasswordResets) Order(value interface{}) *OrmPasswordResets {
	orm.db.Order(value)
	return orm
}

func (orm *OrmPasswordResets) Limit(limit int) *OrmPasswordResets {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmPasswordResets) Offset(offset int) *OrmPasswordResets {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmPasswordResets) Get() PasswordResetsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmPasswordResets) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmPasswordResets) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&PasswordResets{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmPasswordResets) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM password_resets")
}

func (orm *OrmPasswordResets) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmPasswordResets) First(conds ...interface{}) (*PasswordResets, bool) {
	dest := &PasswordResets{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmPasswordResets) Take(conds ...interface{}) (*PasswordResets, int64) {
	dest := &PasswordResets{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmPasswordResets) Last(conds ...interface{}) (*PasswordResets, int64) {
	dest := &PasswordResets{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmPasswordResets) Find(conds ...interface{}) (PasswordResetsList, int64) {
	list := make([]*PasswordResets, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmPasswordResets) Paginate(page int, limit int) (PasswordResetsList, int64) {
	var total int64
	list := make([]*PasswordResets, 0)
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
func (orm *OrmPasswordResets) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmPasswordResets) FirstOrInit(dest *PasswordResets, conds ...interface{}) (*PasswordResets, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmPasswordResets) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmPasswordResets) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmPasswordResets) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmPasswordResets) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmPasswordResets) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmPasswordResets) Where(query interface{}, args ...interface{}) *OrmPasswordResets {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmPasswordResets) And(fuc func(orm *OrmPasswordResets)) *OrmPasswordResets {
	ormAnd := NewOrmPasswordResets()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmPasswordResets) Or(fuc func(orm *OrmPasswordResets)) *OrmPasswordResets {
	ormOr := NewOrmPasswordResets()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmPasswordResets) WhereEmail(val string) *OrmPasswordResets {
	orm.db.Where("`email` = ?", val)
	return orm
}
func (orm *OrmPasswordResets) InsertGetEmail(row *PasswordResets) string {
	orm.db.Create(row)
	return row.Email
}
func (orm *OrmPasswordResets) WhereEmailIn(val []string) *OrmPasswordResets {
	orm.db.Where("`email` IN ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereEmailGt(val string) *OrmPasswordResets {
	orm.db.Where("`email` > ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereEmailGte(val string) *OrmPasswordResets {
	orm.db.Where("`email` >= ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereEmailLt(val string) *OrmPasswordResets {
	orm.db.Where("`email` < ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereEmailLte(val string) *OrmPasswordResets {
	orm.db.Where("`email` <= ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereToken(val string) *OrmPasswordResets {
	orm.db.Where("`token` = ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereCreatedAt(val database.Time) *OrmPasswordResets {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmPasswordResets {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmPasswordResets) WhereCreatedAtLte(val database.Time) *OrmPasswordResets {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmPasswordResets) WhereCreatedAtGte(val database.Time) *OrmPasswordResets {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}

type PasswordResetsList []*PasswordResets
