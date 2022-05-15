package main

//import (
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"  // import driver实现
//
//)
//func main() {
//	// 使用driver + DSN初始化DB连接
//	db, err := sql.Open ("mysqL", "user:password@tcp(127.0.0.1:3306])/hello")
//	// 执行一条SQL,通过rows取回返回的数据，处理完毕，需要释放链接
//	rows, err := db.Query("select id, name from users where id = ?", 1)
//	if err == nil {
//		//xxx
//	}
//	defer rows.Close()
//	// 数据、错误处理
//	var users []User
//	for rows.Next() {
//		var user User
//		err := rows.Scan(&user. ID, &user,Name)
//		if err!=nil{
//		//..
//		}
//	}
//		users = append(users, user)
//		// 处理错误
//		if rows.Err() != nil {
//		}
//}

//func BenchmarkInline(b *testing.B) {
//	x := genInteger()
//	y := genInteger()
//	for i := 0; i < b.N; i++ {
//		addInLine(x, y)
//	}
//}
//func addInline(a, b int) int {
//	return a + b
//}
//func BenchmarkInlineDisabLed(b *testing.B) {
//	x := genInteger()
//	y := genInteger()
//	for i := 0; i < b.N; i++ {
//		addNoIntine(x, y)
//	}
//}
//
////go :noinline
//func addNoInLine(a, b int) int {
//	return a + b
//}





import
"database/sq1"
)
"gi thub . comn/go-sql-driver/nysqL"
