package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type Users struct {
	Id              uint64         `gorm:"column:id;autoIncrement;type:bigint(20) unsigned;primaryKey"`               //
	Name            string         `gorm:"column:name;type:varchar(255)"`                                             //
	Email           string         `gorm:"column:email;type:varchar(255);index:users_email_unique,class:HASH,unique"` //
	EmailVerifiedAt *database.Time `gorm:"column:email_verified_at;type:timestamp;default:NULL"`                      //
	Password        string         `gorm:"column:password;type:varchar(255)"`                                         //
	RememberToken   *string        `gorm:"column:remember_token;type:varchar(100);default:NULL"`                      //
	CreatedAt       *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`                             //
	UpdatedAt       *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`                             //
}

func (receiver *Users) TableName() string {
	return "users"
}

type OrmUsers struct {
	db *gorm.DB
}

func NewOrmUsers() *OrmUsers {
	orm := &OrmUsers{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&Users{})
	return orm
}

func (orm *OrmUsers) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmUsers) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmUsers) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmUsers) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmUsers) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmUsers) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmUsers) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmUsers) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmUsers) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmUsers) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmUsers) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmUsers) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmUsers) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmUsers) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmUsers) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmUsers) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmUsers) Insert(row *Users) error {
	return orm.db.Create(row).Error
}

func (orm *OrmUsers) Inserts(rows []*Users) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmUsers) Order(value interface{}) *OrmUsers {
	orm.db.Order(value)
	return orm
}

func (orm *OrmUsers) Limit(limit int) *OrmUsers {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmUsers) Offset(offset int) *OrmUsers {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmUsers) Get() UsersList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmUsers) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmUsers) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&Users{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmUsers) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM users")
}

func (orm *OrmUsers) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmUsers) First(conds ...interface{}) (*Users, bool) {
	dest := &Users{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmUsers) Take(conds ...interface{}) (*Users, int64) {
	dest := &Users{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmUsers) Last(conds ...interface{}) (*Users, int64) {
	dest := &Users{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmUsers) Find(conds ...interface{}) (UsersList, int64) {
	list := make([]*Users, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmUsers) Paginate(page int, limit int) (UsersList, int64) {
	var total int64
	list := make([]*Users, 0)
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
func (orm *OrmUsers) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmUsers) FirstOrInit(dest *Users, conds ...interface{}) (*Users, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmUsers) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmUsers) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmUsers) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmUsers) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmUsers) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmUsers) Where(query interface{}, args ...interface{}) *OrmUsers {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmUsers) And(fuc func(orm *OrmUsers)) *OrmUsers {
	ormAnd := NewOrmUsers()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmUsers) Or(fuc func(orm *OrmUsers)) *OrmUsers {
	ormOr := NewOrmUsers()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmUsers) WhereId(val uint64) *OrmUsers {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmUsers) InsertGetId(row *Users) uint64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmUsers) WhereIdIn(val []uint64) *OrmUsers {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmUsers) WhereIdGt(val uint64) *OrmUsers {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmUsers) WhereIdGte(val uint64) *OrmUsers {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmUsers) WhereIdLt(val uint64) *OrmUsers {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmUsers) WhereIdLte(val uint64) *OrmUsers {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmUsers) WhereName(val string) *OrmUsers {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmail(val string) *OrmUsers {
	orm.db.Where("`email` = ?", val)
	return orm
}
func (orm *OrmUsers) InsertGetEmail(row *Users) string {
	orm.db.Create(row)
	return row.Email
}
func (orm *OrmUsers) WhereEmailIn(val []string) *OrmUsers {
	orm.db.Where("`email` IN ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailGt(val string) *OrmUsers {
	orm.db.Where("`email` > ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailGte(val string) *OrmUsers {
	orm.db.Where("`email` >= ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailLt(val string) *OrmUsers {
	orm.db.Where("`email` < ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailLte(val string) *OrmUsers {
	orm.db.Where("`email` <= ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailVerifiedAt(val database.Time) *OrmUsers {
	orm.db.Where("`email_verified_at` = ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailVerifiedAtBetween(begin database.Time, end database.Time) *OrmUsers {
	orm.db.Where("`email_verified_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmUsers) WhereEmailVerifiedAtLte(val database.Time) *OrmUsers {
	orm.db.Where("`email_verified_at` <= ?", val)
	return orm
}
func (orm *OrmUsers) WhereEmailVerifiedAtGte(val database.Time) *OrmUsers {
	orm.db.Where("`email_verified_at` >= ?", val)
	return orm
}
func (orm *OrmUsers) WherePassword(val string) *OrmUsers {
	orm.db.Where("`password` = ?", val)
	return orm
}
func (orm *OrmUsers) WhereRememberToken(val string) *OrmUsers {
	orm.db.Where("`remember_token` = ?", val)
	return orm
}
func (orm *OrmUsers) WhereCreatedAt(val database.Time) *OrmUsers {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmUsers) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmUsers {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmUsers) WhereCreatedAtLte(val database.Time) *OrmUsers {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmUsers) WhereCreatedAtGte(val database.Time) *OrmUsers {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmUsers) WhereUpdatedAt(val database.Time) *OrmUsers {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmUsers) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmUsers {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmUsers) WhereUpdatedAtLte(val database.Time) *OrmUsers {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmUsers) WhereUpdatedAtGte(val database.Time) *OrmUsers {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type UsersList []*Users

func (l UsersList) GetIdList() []uint64 {
	got := make([]uint64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l UsersList) GetIdMap() map[uint64]*Users {
	got := make(map[uint64]*Users)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
