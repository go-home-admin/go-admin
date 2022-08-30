package mysql

import (
	sql "database/sql"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

type PersonalAccessTokens struct {
	Id            uint64         `gorm:"column:id;autoIncrement;type:bigint(20) unsigned;primaryKey"` //
	TokenableType string         `gorm:"column:tokenable_type;type:varchar(255)"`                     //
	TokenableId   uint64         `gorm:"column:tokenable_id;type:bigint(20) unsigned"`                //
	Name          string         `gorm:"column:name;type:varchar(255)"`                               //
	Token         string         `gorm:"column:token;type:varchar(64)"`                               //
	Abilities     *string        `gorm:"column:abilities;type:text;default:NULL"`                     //
	LastUsedAt    *database.Time `gorm:"column:last_used_at;type:timestamp;default:NULL"`             //
	CreatedAt     *database.Time `gorm:"column:created_at;type:timestamp;default:NULL"`               //
	UpdatedAt     *database.Time `gorm:"column:updated_at;type:timestamp;default:NULL"`               //
}

func (receiver *PersonalAccessTokens) TableName() string {
	return "personal_access_tokens"
}

type OrmPersonalAccessTokens struct {
	db *gorm.DB
}

func NewOrmPersonalAccessTokens() *OrmPersonalAccessTokens {
	orm := &OrmPersonalAccessTokens{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB).Model(&PersonalAccessTokens{})
	return orm
}

func (orm *OrmPersonalAccessTokens) GetDB() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmPersonalAccessTokens) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmPersonalAccessTokens) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmPersonalAccessTokens) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmPersonalAccessTokens) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmPersonalAccessTokens) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmPersonalAccessTokens) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmPersonalAccessTokens) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmPersonalAccessTokens) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmPersonalAccessTokens) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmPersonalAccessTokens) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmPersonalAccessTokens) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmPersonalAccessTokens) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmPersonalAccessTokens) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmPersonalAccessTokens) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmPersonalAccessTokens) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmPersonalAccessTokens) Insert(row *PersonalAccessTokens) error {
	return orm.db.Create(row).Error
}

func (orm *OrmPersonalAccessTokens) Inserts(rows []*PersonalAccessTokens) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmPersonalAccessTokens) Order(value interface{}) *OrmPersonalAccessTokens {
	orm.db.Order(value)
	return orm
}

func (orm *OrmPersonalAccessTokens) Limit(limit int) *OrmPersonalAccessTokens {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmPersonalAccessTokens) Offset(offset int) *OrmPersonalAccessTokens {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmPersonalAccessTokens) Get() PersonalAccessTokensList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmPersonalAccessTokens) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmPersonalAccessTokens) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&PersonalAccessTokens{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmPersonalAccessTokens) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM personal_access_tokens")
}

func (orm *OrmPersonalAccessTokens) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmPersonalAccessTokens) First(conds ...interface{}) (*PersonalAccessTokens, bool) {
	dest := &PersonalAccessTokens{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmPersonalAccessTokens) Take(conds ...interface{}) (*PersonalAccessTokens, int64) {
	dest := &PersonalAccessTokens{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmPersonalAccessTokens) Last(conds ...interface{}) (*PersonalAccessTokens, int64) {
	dest := &PersonalAccessTokens{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmPersonalAccessTokens) Find(conds ...interface{}) (PersonalAccessTokensList, int64) {
	list := make([]*PersonalAccessTokens, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmPersonalAccessTokens) Paginate(page int, limit int) (PersonalAccessTokensList, int64) {
	var total int64
	list := make([]*PersonalAccessTokens, 0)
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
func (orm *OrmPersonalAccessTokens) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmPersonalAccessTokens) FirstOrInit(dest *PersonalAccessTokens, conds ...interface{}) (*PersonalAccessTokens, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmPersonalAccessTokens) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmPersonalAccessTokens) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmPersonalAccessTokens) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmPersonalAccessTokens) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmPersonalAccessTokens) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmPersonalAccessTokens) Where(query interface{}, args ...interface{}) *OrmPersonalAccessTokens {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmPersonalAccessTokens) And(fuc func(orm *OrmPersonalAccessTokens)) *OrmPersonalAccessTokens {
	ormAnd := NewOrmPersonalAccessTokens()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmPersonalAccessTokens) Or(fuc func(orm *OrmPersonalAccessTokens)) *OrmPersonalAccessTokens {
	ormOr := NewOrmPersonalAccessTokens()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

func (orm *OrmPersonalAccessTokens) WhereId(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) InsertGetId(row *PersonalAccessTokens) uint64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmPersonalAccessTokens) WhereIdIn(val []uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereIdGt(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereIdGte(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereIdLt(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereIdLte(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereTokenableType(val string) *OrmPersonalAccessTokens {
	orm.db.Where("`tokenable_type` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereTokenableId(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`tokenable_id` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereTokenableIdIn(val []uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`tokenable_id` IN ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereTokenableIdNe(val uint64) *OrmPersonalAccessTokens {
	orm.db.Where("`tokenable_id` <> ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereName(val string) *OrmPersonalAccessTokens {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereToken(val string) *OrmPersonalAccessTokens {
	orm.db.Where("`token` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereAbilities(val string) *OrmPersonalAccessTokens {
	orm.db.Where("`abilities` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereLastUsedAt(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`last_used_at` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereLastUsedAtBetween(begin database.Time, end database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`last_used_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereLastUsedAtLte(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`last_used_at` <= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereLastUsedAtGte(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`last_used_at` >= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereCreatedAt(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereCreatedAtLte(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereCreatedAtGte(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereUpdatedAt(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereUpdatedAtLte(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmPersonalAccessTokens) WhereUpdatedAtGte(val database.Time) *OrmPersonalAccessTokens {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type PersonalAccessTokensList []*PersonalAccessTokens

func (l PersonalAccessTokensList) GetIdList() []uint64 {
	got := make([]uint64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l PersonalAccessTokensList) GetIdMap() map[uint64]*PersonalAccessTokens {
	got := make(map[uint64]*PersonalAccessTokens)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l PersonalAccessTokensList) GetTokenableIdList() []uint64 {
	got := make([]uint64, 0)
	for _, val := range l {
		got = append(got, val.TokenableId)
	}
	return got
}
func (l PersonalAccessTokensList) GetTokenableIdMap() map[uint64]*PersonalAccessTokens {
	got := make(map[uint64]*PersonalAccessTokens)
	for _, val := range l {
		got[val.TokenableId] = val
	}
	return got
}
