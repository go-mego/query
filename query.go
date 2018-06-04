package query

import "github.com/go-mego/mego"

const (
	// 欄位標籤名稱。
	fieldTag = "query"
)

// New 能夠建立一個新的網址參數模組建構體並且分析請求中的網址參數。
func New() mego.HandlerFunc {
	return func(c *mego.Context) {
		c.Map(&Query{
			context: c,
		})
	}
}

// Query 呈現了網址參數模組建構體。
type Query struct {
	context *mego.Context
}

// Bind 能夠將網址參數映射至本地的變數。當無法映射至建構體的時候會離開請求，並且以 `text/plain` 回傳一個 HTTP 400 錯誤狀態碼。
func (q *Query) Bind(dest interface{}) error {

}

// ShouldBind 與 `Bind` 相同作用，但會在映射失敗的時候不做任何處理（亦即：不會離開請求、不會回傳錯誤狀態碼）。
func (q *Query) ShouldBind(dest interface{}) error {

}

// Has 能夠得知一個指定的參數是否存在於網址中。
//     GET /?name=Manu&lastname=
//     c.Has("name")     == true
//     c.Has("id")       == false
//     c.Has("lastname") == true
func (q *Query) Has(key string) bool {
	v, ok := q.context.Request.URL.Query()[key]
	return ok && len(v) > 0
}

// Get 能夠從網址參數取得指定的變數資料，當該參數不存在時則會回傳一個空白字串。
//     GET /path?id=1234&name=Manu&value=
// 	   q.Get("id")    == "1234"
// 	   q.Get("name")  == "Manu"
// 	   q.Get("value") == ""
// 	   q.Get("wtf")   == ""
func (q *Query) Get(key string) string {
	if !q.Has(key) {
		return ""
	}
	return v.GetMulti(key)[0]
}

// GetDefault 與 `Get` 相同，但在參數不存在時會以 `defaultValue` 作為回傳的預設值。
//     GET /?name=Manu&lastname=
//     q.GetDefault("name", "unknown")  == "Manu"
//     q.GetDefault("id", "none")       == "none"
//     q.GetDefault("lastname", "none") == "none"
func (q *Query) GetDefault(key string, defaultValue string) string {
	v := c.Get(key)
	if v == "" {
		return defaultValue
	}
	return v
}

// GetMulti 會以字串切片的方式匯集網址參數中相同的變數資料，
// 切片的長度基於請求有多少個相同欄位而定。
func (q *Query) GetMulti(key string) []string {
	if q.Has(key) {
		return c.Request.URL.Query()[key]
	}
	return []string{}
}