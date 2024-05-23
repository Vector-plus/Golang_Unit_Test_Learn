# Golang单元测试

- [Golang单元测试](#golang单元测试)
	- [单元测试的概念](#单元测试的概念)
		- [单元测试的优势](#单元测试的优势)
	- [Golang项目中单元测试实现](#golang项目中单元测试实现)
		- [单元测试编写流程](#单元测试编写流程)
		- [单元测试中的mock](#单元测试中的mock)
		- [测试代码管理框架---\>GoConvey](#测试代码管理框架---goconvey)
		- [常见依赖函数的替换---\>goMock \& goMonkey](#常见依赖函数的替换---gomock--gomonkey)
			- [go-Mock原理及其使用](#go-mock原理及其使用)
			- [go-Monkey原理及其使用](#go-monkey原理及其使用)
			- [其他mock的实现](#其他mock的实现)
		- [http相关操作---\>go-httptest](#http相关操作---go-httptest)
		- [Mysql外部依赖---\>go-sqlmock](#mysql外部依赖---go-sqlmock)
		- [Reids外部依赖---\>go-miniredis](#reids外部依赖---go-miniredis)
	- [学习资料总结](#学习资料总结)
---

**<font face="HEI">本文主要讲述了什么是单元测试、单元测试的优缺点，实际项目中单元测试的必要性和Golang项目中单元测试的具体编写方法</font>**

## 单元测试的概念

>测试：对代码正确性的验证工作。一般分为单元测试、集成测试和系统测试。

**单元测试**：它是在软件开发过程中进行的最小级别的测试。单元测试通常针对软件应用中的一个小模块（例如，一个函数或方法）进行，将其与系统的其他部分隔离开来。目标是确保每个代码模块都达到预期效果，从而在初期阶段就挖掘出能够提升代码质量的bug。

**集成测试**：集成测试即将多个模块或单元组合起来测试，以确保这些单元组合在一起后可以正常工作。这在处理边界条件，检查不同模块间交互问题具有重要作用。集成测试能够发现在模块间交互过程中可能存在的问题。

**系统测试**：系统测试作为测试流程中的最后一个环节，它考虑了所有的软件组件以及完整的系统架构。系统测试将软件看作一个整体，测试所有代码、软件架构、用户界面、数据库等等，从而确保整个系统的行为都符合预期。系统测试的目的是模拟实际的应用场景确保软件产品作为一个整体达到或超出预期的标准。

### 单元测试的优势


![测试金字塔](/Golang_Unit_Test_Learn/blob/main/static/单元测试的优势.jpg)


单元测试在测试金字塔的最下层，其测试粒度最小，能更快地发现项目代码中的bug，研究表明在单元测试中发现bug能极大地减小系统的维护成本，相较于上面两层的测试方案其效率最高，能最快地解决代码中的问题，保证代码的健壮性。同时，单元测试还能提高代码的质量（开发人员编写测试）、促进代码的重用、支持项目的重构（根据已经存在的测试代码确保原功能没有被破坏）。
1. 验证代码逻辑

对于一个代码模块，编写单元测试的过程，就是对代码逻辑进行重新Review的过程；而执行单元测试的过程，就是验证代码是否按既定逻辑运行的过程。

2. 减少代码缺陷

我们的工程都是分层分模块的，每个模块都是独立的逻辑部分。通过单元测试保障工程各个“零件”按“规格”（需求）执行，就能保证整个“机器”（项目）运行正确，最大限度地减少bug。

3. 促进代码设计

在编写单测的过程中，如果发现单测代码非常难写，一般表明被测试的代码包含了太多的依赖或职责，需要反思代码的合理性，进而促进代码设计的优化。

4. 便于多人协作

在多人协助的项目中，所依赖的服务接口不一定已经开发完毕，导致服务进行联调工作。此时，单元测试有效地解决了这个问题——只需Mock 服务接口数据，便可以完成自己代码的测试。

5. 便于缺陷定位

由于单元规模较小，复杂性较低，因而发现错误后容易隔离和定位，有利于代码调试工作。

6. 利于代码重构

重构过程中复用测试代码。如今持续型的项目越来越多，代码不断的在变化和重构，通过单元测试，开发可以放心的修改重构代码，减少改代码时的心理负担，提高重构的成功率。

## Golang项目中单元测试实现

### 单元测试编写流程

>单元测试编写流程图

![单元测试编写流程](/Golang_Unit_Test_Learn/static/单元测试步骤.png)

1. 定义对象阶段

定义对象阶段主要包括：定义被测对象、需要模拟依赖的对象（成员变量）。

2. 模拟方法阶段

模拟方法阶段主要包括（主要针对需要mock替换的情况）：模拟依赖对象（其参数、返回值和err）、模拟依赖方法。

3. 调用方法阶段

调用方法阶段主要包括：模拟依赖对象、调用被测方法、验证参数对象（返回值和err）

4. 验证方法阶段

验证方法阶段主要包括：验证依赖方法、验证数据对象、验证依赖对象。

### 单元测试中的mock

**mock的必要性：**

1. Mock可以用来解除外部服务依赖，从而保证了测试用例的独立性。

现在的互联网软件系统，通常采用了分布式部署的微服务，为了单元测试某一服务而准备其它服务，存在极大的依赖性和不可行性。

2. Mock可以减少全链路测试数据准备，从而提高了编写测试用例的速度。

传统的集成测试，需要准备全链路的测试数据，可能某些环节并不是你所熟悉的。最后，耗费了大量的时间和精力，并不一定得到你想要的结果。现在的单元测试，只需要模拟上游的输入数据，并验证给下游的输出数据，编写测试用例并进行测试的速度可以提高很多倍。

3. Mock可以模拟一些非正常的流程，从而保证了测试用例的代码覆盖率。

根据单元测试的BCDE原则，需要进行边界值测试（Border）和强制错误信息输入（Error），这样有助于覆盖整个代码逻辑。在实际系统中，很难去构造这些边界值，也能难去触发这些错误信息。而Mock从根本上解决了这个问题：想要什么样的边界值，只需要进行Mock；想要什么样的错误信息，也只需要进行Mock。

4. Mock可以不用加载项目环境配置，从而保证了测试用例的执行速度。

在进行集成测试时，我们需要加载项目的所有环境配置，启动项目依赖的所有服务接口。往往执行一个测试用例，需要几分钟乃至几十分钟。采用Mock实现的测试用例，不用加载项目环境配置，也不依赖其它服务接口，执行速度往往在几秒之内，大大地提高了单元测试的执行速度。

**常见的mock情景**

1. 全局变量、对象、函数、方法。
2. http/rpc网络服务
3. MySQL/redis/mongoDB等数据库交互场景


### 测试代码管理框架--->GoConvey


**学习资料：**

[GoConvey官网](https://smartystreets.github.io/goconvey/)

[GoConvey框架使用指南博文](https://bbs.huaweicloud.com/blogs/363181)

一款针对Golang的测试框架，可以管理和运行测试用例，同时提供了丰富的断言函数，并支持很多 Web 界面特性。

**特点：**
- 直接与 go test 集成
- 巨大的回归测试套件
- 可读性强的色彩控制台输出
- 完全自动化的 Web UI
- 测试代码生成器
- 桌面提醒（可选）
- 自动在终端中运行自动测试脚本
- 可立即在 Sublime Text 中打开测试问题对应的代码行 (some assembly required)

**实现代码链接**：[go-Convey使用](/gotest/conveytest/cul_test.go)

**some待测函数**
```go
package conveytest

import "errors"

// 两数相加函数
func Add(a, b int) int {
	return a + b
}

// 两数差值函数
func Subt(a, b int) int {
	return a - b
}

// 两数相乘函数
func Mult(a, b int) int {
	return a * b
}

// 两数相除函数
func Div(a, b float64) (re float64, err error) {
	if b == 0 {
		re = -1
		err = errors.New("除数为0")
	} else {
		re = a / b
	}
	return
}

```
**Convey测试代码**
```go
package conveytest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 命令行查看测试结果（查看覆盖率） go test -v -cover
// web-UI查看测试结果（设置端口） goconvey -port 8099
func TestCulculate(t *testing.T) {
	//新增测试
	Convey("测试Add函数", t, func() {
		//断言判断结果是否正确
		So(Add(1, 2), ShouldEqual, 3)
	})

	Convey("测试Subt函数", t, func() {
		//断言函数判断
		So(Subt(1, 3), ShouldEqual, -2)
		So(Subt(10, 5), ShouldEqual, 5)
	})

	Convey("测试Mult函数", t, func() {
		//断言判断
		So(Mult(1, 2), ShouldEqual, 2)
		So(Mult(4, 6), ShouldEqual, 24)
	})

	Convey("测试Div函数", t, func() {
		//嵌套使用convey，只有最外层convey需要传入*testing.T参数
		Convey("除数为0", func() {
			_, err := Div(1, 0)
			So(err, ShouldNotBeNil)
			_, err = Div(7, 0)
			So(err, ShouldNotBeNil)
			_, err = Div(1999999, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("除数不为0", func() {
			re, err := Div(4, 2)
			So(err, ShouldBeNil)
			So(re, ShouldEqual, 2)

			re, err = Div(5, 2)
			So(err, ShouldBeNil)
			So(re, ShouldEqual, 2.5)

			re, err = Div(20, 2)
			So(err, ShouldBeNil)
			So(re, ShouldEqual, 10)
		})
	})
}
```
**命令行查看测试结果（查看覆盖率）效果图**
>指令： go test -v -cover

![命令窗口指令](/Golang_Unit_Test_Learn/static/convey-ctl.png)

**WEB UI查看**
>指令：goconvey -port 9999

![web-ui测试页面](/Golang_Unit_Test_Learn/static/convey-webUI.png)

### 常见依赖函数的替换--->goMock & goMonkey

>当在单元测试中有需要替换的依赖时（例如rpc远程函数、复杂函数），可以考虑goMock或者goMonkey进行依赖替换

**goMock和goMonkey如何选择？**

**最好使用goMock，没有条件就用goMonkey。goMock对interface具有强依赖但其对源代码的影响最小。**

#### go-Mock原理及其使用

**学习资料：**

[官网](https://github.com/golang/mock?tab=readme-ov-file)

[学习博文](https://juejin.cn/post/7103069443399352350)

**原理：**

Mock是在测试过程中，对于一些不容易构造/获取的对象，创建一个Mock对象来模拟对象的行为。Mock最大的功能是帮你把单元测试进行解耦通过mock模拟的机制，生成一个模拟方法，然后替换调用原有代码中的方法，它其实是做一个真实的环境替换掉业务本需要的环境。只能模拟 interface 方法，这就要求我们业务编写代码的时候具有非常好的接口设计，这样才能顺利生成 mock 代码。

**实现功能：**
+ 验证这个对象的某些方法的调用情况，调用了多少次，参数是什么，返回值是什么等等
+ 指定这个对象的某些方法的行为，返回特定的值，或者是执行特定的动作等等

**使用步骤**
1. 使用mockgen 为你想要 mock 的接口生成一个mock。

2. 在你的测试代码中,创建一个gomock.Controller实例并把它作为参数传递给 mock对象的构造函数来创建一个mock 对象。

3. 调用 EXPECT()为你的mock对象设置各种期望和返回值。

4. 调用 mock控制器的Finish()以验证 mock 的期望行为。

>具体实现代码

**代码链接：**[go-mock实现依赖替换](/gotest/mockytest/mock_test.go)
```go
package mockytest

import (
	"fmt"
	"gotest/dv1/example"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

// 测试go-Mock打桩，利用interface多态的特点生成测试对象，实现替换interface中的某些方法
// 生成对象指令（mockgen插件） mockgen -source=./mail.go -destination=../mocktest/mock_mockdemo.go -package=mocktest Mail
func TestMock(t *testing.T) {

	Convey("测试example.WiteAndSend(string,int)方法", t, func() {
		//获取mockctl，它代表mock生态系统中的顶级控件。定义了mock对象的范围、生命周期和期待值。另外它在多个goroutine中是安全的
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()

		//生成需要mock的实例
		mockIA := NewMockIA(mockCtl)

		//声明给定的调用应按顺序进行
		gomock.InOrder(

			//三个步骤：EXPECT()返回一个允许调用者设置期望和返回值的对象。FA("test")是设置入参并调用 mock 实例中的方法。Return是设置先前调用的方法出参。简单来说，就是设置入参并调用，最后设置返回值
			mockIA.EXPECT().FA("test").Return("test-----2024-05-20", nil),
			mockIA.EXPECT().FB(2).Return(1, nil),
			mockIA.EXPECT().FB(1).Return(1, nil),
		)

		//生成实例对象
		unitT := example.NewIA(mockIA)

		//调用测试方法
		ans, err := unitT.WiteAndSend("test", 2)
		So(err, ShouldBeNil)
		So(ans, ShouldEqual, "test-----2024-05-20 === 1 -----> 1\n")
		fmt.Println(ans)
	})

}
```

#### go-Monkey原理及其使用

**学习资料：**

[官网](https://github.com/bouk/monkey)

[Monkey框架使用](https://studygolang.com/articles/11296)

**原理：**

利用猴子补丁(monkey patching)，替换改写方法，成员方法、全局变量

**实现功能：**

通过 gomonkey(stub) 打桩来替换掉我们原本的执行逻辑,不需要一个接口就能打桩

**注意事项：**

+ 不支持内联函数，在测试的时候需要通过go test -gcflags=all=-l 关闭内联优化。
+ 不是线程安全的，所以不要把它用到并发的单元测试中。
+ 不支持异包未导出函数的打桩、不支持同包未导出方法的打桩

**使用步骤：**

1. 导入Gomonkey包
2. 使用 GoMonkey 的 Patch 函数来替换掉你想要的函数或者方法。
3. 编写替换函数。在测试中，当被劫持的函数被调用时，将会执行你在步骤2中定义的新函数，而不是原有的函数。

**具体实现**

实现代码：[go-monkey测试代码](/gotest/monkeytest/monkey_test.go)

```go
package monkeytest

import (
	"gotest/dv1/example"
	"reflect"
	"testing"

	"bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

// 使用monkey方式，直接替换对象中的方法，其原理为采用热补丁的方式替换二进制文件中的代码，在执行时直接跳转到桩实现
// 使用时需要使用 -gcflags=-l 关闭内联优化（跳转到正确的执行文件，该方式影响了源文件）
// 该方式不能用于多线程的测试
// 运行指令 go test -run=TestUser -v -gcflags=-l
func TestWiteAndSend2(t *testing.T) {
	Convey("monkey测试WithAndSend2方法", t, func() {

		unitT2 := &example.UnitT2{}

		//为对象方法进行打桩,需要传入相应的对象，函数名以及修改函数
		monkey.PatchInstanceMethod(reflect.TypeOf(unitT2), "FA2", func(*example.UnitT2, string) (string, error) {
			return "test----2024-05-20", nil
		})

		monkey.PatchInstanceMethod(reflect.TypeOf(unitT2), "FB2", func(e *example.UnitT2, n int) (int, error) {
			return n + 1, nil
		})

		//直接调用测试方法，此时该对象的FA2和FB2已经被替换
		ans, err := unitT2.WiteAndSend2("test", 1)

		So(err, ShouldBeNil)
		So(ans, ShouldEqual, "test----2024-05-20 === 2 -----> 3\n")

	})
}
```

#### 其他mock的实现

**Tencent开源框架：**[goom单测Mock框架](https://github.com/Tencent/goom)

**bytedance开源框架：**[Mockey单测框架](https://github.com/bytedance/mockey/blob/main/README_cn.md)

### http相关操作--->go-httptest

**学习资料：**

[官网](https://pkg.go.dev/net/http/httptest)

[使用方法博文](https://blog.csdn.net/lavorange/article/details/73369153)

**原理：**

httptest是Go官方提供的专门用于进行http Web开发测试的包。我们在单测过程中，也不要直连真正的web server，httptest.NewServer这个函数创建一个新的HTTP服务器进行测试，它返回一个*httptest.Server对象，该对象包含一个URL属性，代表该服务器的网络地址。

**实现功能：**

编写基于HTTP的单元测试，无需依赖实际的网络环境。

**使用步骤：**

1. 通过httptest.NewServer创建了一个测试的http server
2. 写自己的HandlerFunc函数，处理请求。设置返回参数等。

**具体实现：**
**测试实现：**[httptest实现代码](/gotest/httptest/httpserver_test.go)

```go
package httptest

import (
	"encoding/json"
	"fmt"
	"gotest/dv1/example"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Resp struct {
	Code string
	Data interface{}
	Msg  string
}

var msgData = &Resp{
	Code: "200",
	Data: &example.User{UserName: "fht", Password: "123456", Age: 12},
	Msg:  "ok",
}

func HttpSrv(w http.ResponseWriter, r *http.Request) {
	//在该方法中还能处理请求路径
	// r.URL.EscapedPath() != "/userInfo"
	// 获取请求参数：r.ParseForm，r.Form.Get("addr")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data, _ := json.Marshal(msgData)
		w.Write(data)
	}
}

func TestHttpSrv(t *testing.T) {
	Convey("测试httpSrv", t, func() {
		// 通过httptest.NewServer创建了一个测试的http server,写自己的HandlerFunc函数，处理请求。
		ts := httptest.NewServer(http.HandlerFunc(HttpSrv))
		defer ts.Close()

		resp, err := example.GetUserInfo(ts.URL)
		fmt.Println(resp)
		So(err, ShouldBeNil)
		param, _ := json.Marshal(msgData)
		So(resp, ShouldEqual, string(param))
	})
}

```



### Mysql外部依赖--->go-sqlmock

**学习资料：**

[官网](https://github.com/DATA-DOG/go-sqlmock)

[使用博文](https://www.wuguozhang.com/archives/150/)


**原理：**

go-sqlmock 本质是一个实现了 sql/driver 接口的 mock 库，它的设计目标是支持在测试中，模拟任何 sql driver 的行为，而不需要一个真正的数据库连接。

**实现功能：**

1. 模拟数据库查询结果：你可以使用 SQLMock 对数据库查询的结果进行模拟，来测试你的代码如何处理从数据库查询得到的结果。

2. 模拟执行 SQL 命令：SQLMock 允许你模拟数据库的各种操作，比如 INSERT、UPDATE、DELETE 等。你可以为这些操作定义期望的行为和结果。
模拟数据库事务：可以用 SQLMock 来创建和验证事务的行为，包括 commit 和 rollback。

3. 检查 SQL 语句：SQLMock 可以检查是否执行了预期的 SQL 语句，并验证其执行的顺序、次数以及参数是否正确。

4. 模拟数据库错误：你可以配置 SQLMock 使数据库操作产生错误，来测试你的代码如何处理各种数据库错误。

**使用步骤：**

1. 创建Mock数据库连接，sqlmock.New()
2. 根据项目中使用的orm框架生成对应的sql.Client
3. 模拟数据库中的行为调用Expectxxx()函数

**具体实现：**
**实现代码链接**[go-sqlmock实现](/gotest/mysqltest/sql_test.go)
```go
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
		sql := "SELECT * FROM `users` WHERE ID = ?"
		re := sqlmock.NewRows([]string{"UserName", "password", "age"}).AddRow("fht", "123456ddd", 12)
		//设置sql_mock期待执行的sql语句。可以根据不同的sql使用情景选择合适的Expectxxx方法
		mock.ExpectQuery(sql).WithArgs(uid).WillReturnRows(re)

		user, err := sqlDB.FindByUserId(uid)
		fmt.Println(user)
		So(err, ShouldBeNil)

	})

}
```

### Reids外部依赖--->go-miniredis

**学习资料：**

[官网](https://github.com/alicebob/miniredis)

[使用博文](https://www.bytezonex.com/archives/5rw0tYqE.html)



**原理：**

Miniredis实现了Redis服务器的一部分，用于单元测试。它通过真正的 TCP接口实现了简单、廉价、内存中的Redis替代。将其视为Redis版本net/http/httptest。

**实现功能：**

它使您无需使用模拟代码，而且由于redis服务器位于测试过程中，您可以直接查询值，而无需通过服务器堆栈。不依赖于外部二进制文件，可以使你的单元测试更加独立和可控，让你可以更是心地测试和验证你的代码。

**使用步骤：**

1. 导入miniredis库，启动miniredis服务，得到其访问地址。
2. 构建一个Redis客户端：使用miniredis的地址创建一个Redis客户端。
3. 调用Redis命令：通过客户端调用Redis命令。预置需要的数据到服务中。

**具体实现：**
**代码实现链接：**[go-miniredis](/gotest/redistest/redis_test.go)

```go
package redistest

import (
	"gotest/dv1/example"
	"testing"

	"github.com/redis/go-redis/v9"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/alicebob/miniredis/v2"
)

// miniredis是一个纯go实现的用于单元测试的redis server。它是一个简单易用的、基于内存的redis替代品，
// 它具有真正的TCP接口，你可以把它当成是redis版本的net/http/httptest。
// 当我们为一些包含Redis操作的代码编写单元测试时就可以使用它来mock Redis操作。
func TestMiniRedis(t *testing.T) {

	Convey("测试miniRedis", t, func() {
		//使用minireids新建一个redis
		rs, err := miniredis.Run()
		So(err, ShouldBeNil)
		defer rs.Close()

		key := "2024-05-20"

		//预置数据到redis中
		_, err = rs.SAdd(key, "test---miniredis")
		So(err, ShouldBeNil)

		reclient := redis.NewClient(&redis.Options{Addr: rs.Addr()})
		rdb := &example.Rdb{Rdb: reclient}

		ans, err := rdb.GetAllUser(key)

		So(err, ShouldBeNil)
		So(ans[0], ShouldEqual, "test---miniredis")

	})
}
```

## 学习资料总结

**Google软件测试书籍：**[google软件测试之道](/static/google软件测试之道.pdf)

**阿里单元测试书籍：**[Java单元测试实战](https://developer.aliyun.com/ebook/7895/102694?spm=a2c6h.26392470.ebook-read.11.5186121c0gF9wa)

**Golang单测实战经验：**[一文了解一线互联网大厂的 Golang 单测最佳实战经验](https://developer.aliyun.com/article/1295307?spm=5176.26934562.main.1.433c321djYvhex)

**Golang工程化测试：**[业务项目中的Go单元测试心得](https://cloud.tencent.com/developer/article/2219521)






