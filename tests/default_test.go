package test

import (
	_ "BeegoDemo/routers"
	//	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDataBase("default", "sqlite3", "eland.sqlite")
	orm.Debug = true
	beego.BConfig.RouterCaseSensitive = false
	//方面测试的内容
	testValue := url.Values{}
	strings.NewReader(testValue.Encode())
	var testInt int = 11
	strconv.Itoa(testInt)
}

// TestMain is a sample to run an endpoint test
//func TestMain(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)

//	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

//	Convey("Subject: Test Station Endpoint\n", t, func() {
//		Convey("Status Code Should Be 200", func() {
//			So(w.Code, ShouldEqual, 200)
//		})
//		Convey("The Result Should Not Be Empty", func() {
//			So(w.Body.Len(), ShouldBeGreaterThan, 0)
//		})
//	})
//}

//func TestLogin(t *testing.T) {
//	//		v := url.Values{}
//	//		v.Set("userid", "liche1")
//	//		v.Add("username", "liche1_name")
//	//	r, _ := http.NewRequest("POST", "/user/login", strings.NewReader(v.Encode()))
//	//	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

//	v := `{"username":"liche","userid":"4444"}`
//	r, _ := http.NewRequest("POST", "/user/login", strings.NewReader(v))
//	r.Header.Set("Content-Type", "application/json")

//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//	beego.Trace("testing", "TestSetUser", w.Code, w.Body.String())
//	Convey("Login", t, func() {

//		So(w.Code, ShouldEqual, 200)
//	})
//}

//func TestCreateUser(t *testing.T) {
//	UserId := "liche2"
//	v := url.Values{}
//	v.Set("UserId", UserId)
//	v.Set("UserName", "李澈2")
//	v.Set("Password", "2222")
//	v.Set("UseChk", "true")
//	v.Set("AuthChk", "true")
//	r, _ := http.NewRequest("POST", "/user/createuser", strings.NewReader(v.Encode()))
//	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

func TestCreateBookmark(t *testing.T) {
	var userId string = "liche1"
	var bookmarkName = "TestName"
	var desc string = "TestDesc"
	var fixChk string = "true"
	var numChk string = "true"
	var placeholder string = "{{}}"
	var useChk string = "true"
	/*this.Desc, this.FixChk, this.NumChk, this.Placeholder, this.InUserId*/
	url := "/bookmark/CreateBookmark?BookmarkName=" + bookmarkName + "&Desc=" + desc + "&FixChk=" + fixChk + "&NumChk=" + numChk + "&Placeholder=" + placeholder + "&UseChk=" + useChk + "&InUserId=" + userId

	beego.Trace("testing", "TestCreateBookmark", url)
	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("TestCreateBookmark", t, func() {
		So(w.Code, ShouldEqual, 200)
	})
}

func TestGetBookmarkByName(t *testing.T) {
	var name string = "TestName"
	url := "/bookmark/GetBookmarkByName?Name=" + name

	beego.Trace("testing", "TestGetBookmarkByName", url)

	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("TestGetBookmarkByName", t, func() {
		So(w.Code, ShouldEqual, 200)
	})
}

func TestDeleteBookmarkById(t *testing.T) {
	var bookmarkId string = "10"
	url := "/bookmark/DeleteBookmarkById?BookmarkId=" + bookmarkId

	beego.Trace("testing", "TestDeleteBookmarkById", url)

	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("TestDeleteBookmarkById", t, func() {
		So(w.Code, ShouldEqual, 200)
	})
}

func TestUpdateBookmarkById(t *testing.T) {
	var userId string = "chenwenqiang"
	var bookmarkId string = "11"
	var bookmarkName = "TestNameHaHa"
	var desc string = "TestDesc"
	var fixChk string = "true"
	var numChk string = "true"
	var placeholder string = "{{}}"
	var useChk string = "true"
	/*this.Desc, this.FixChk, this.NumChk, this.Placeholder, this.InUserId*/
	url := "/bookmark/UpdateBookmarkById?BookmarkId=" + bookmarkId + "&BookmarkName=" + bookmarkName + "&Desc=" + desc + "&FixChk=" + fixChk + "&NumChk=" + numChk + "&Placeholder=" + placeholder + "&UseChk=" + useChk + "&InUserId=" + userId

	beego.Trace("testing", "TestUpdateBookmarkById", url)
	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("TestUpdateBookmarkById", t, func() {
		So(w.Code, ShouldEqual, 200)
	})
}

//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)

//	//	beego.Trace("testing", "TestCreateUser", "Code[%d]\n%s", w.Code, w.Body.String())
//	Convey("TestCreateUser", t, func() {
//		So(w.Code, ShouldEqual, 200)
//	})
//}

//func TestUserGetById(t *testing.T) {
//	var userId string = "liche1"
//	url := "/user/UserGetById?UserId=" + userId

//	beego.Trace("testing", "TestUserGetById", url)

//	r, _ := http.NewRequest("GET", url, nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)

//	Convey("TestUserGetById", t, func() {
//		So(w.Code, ShouldEqual, 200)
//	})

//}
