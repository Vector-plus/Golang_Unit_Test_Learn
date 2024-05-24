package mysqltest

import (
	"fmt"
	"gotest/dv1/example"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	. "github.com/smartystreets/goconvey/convey"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	//虚拟建立数据库，返回*sql.DB对象，用来新建mysql连接
	//sqlmock.New()方法创建数据库
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		return nil, nil, err
	}

	//gorm创建mysql连接，用于后续的使用
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

// 在使用gorm等orm框架时，由于需要和数据库进行交互，使得单元测试难于编写，
// 使用go-sqlmock库可以很好的缓解这些问题，其不需要建立真正的数据库连接，
// 可以在单元测试中模拟任何 sql 驱动程序的行为，有助于测试数据库交互。

func TestSqlMock(t *testing.T) {

	Convey("测试mysql数据库操作", t, func() {
		//获取虚拟数据库连接，mock替换
		db, mock, err := getDBMock()
		So(err, ShouldBeNil)

		//获取数据库对象
		sqlDB := example.NewUserDb(db)

		//定义的传入参数、期望的sql语句和期待返回的值
		uid := 1
		//sqlite
		//gorm.DryRun
		//设计模式依赖注入

		// sql := "SELECT * FROM `users` WHERE ID = ?"

		u := &example.User{}
		stmt := db.Session(&gorm.Session{DryRun: true}).Where("ID = ?", uid).Find(&u).Statement
		sql := stmt.SQL.String()
		// fmt.Println("sqlll", sql)

		re := sqlmock.NewRows([]string{"UserName", "password", "age"}).AddRow("fht", "123456ddd", 12)
		//设置sql_mock期待执行的sql语句。可以根据不同的sql使用情景选择合适的Expectxxx方法
		mock.ExpectQuery(sql).WithArgs(uid).WillReturnRows(re)

		user, err := sqlDB.FindByUserId(uid)
		fmt.Println(user)
		So(err, ShouldBeNil)

	})

}
