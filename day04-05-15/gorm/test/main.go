package main

func main() {
	//for i := 0; i < maxBadConnRetries; i++ {
	//	// 从连接池获取连接或通过driver新建连接
	//	dc, err := db.conn(ctx, strategy)
	//	//有空闲连接-> reuse -> max life time
	//	//新建连接-> max open.. .
	//	//将连接放回连接池
	//	defer dc.db.putConn(dc, err, true)
	//	/// validateConnection有无错误
	//	// max life time, max idle conns检查
	//	//连接实现driver.Queryer; driver.Execer等interface
	//	if err == nil{
	//		err = dc.ci.Query(sql, args...)
	//	}
	//	isBadConn = errors.Is(err, driver. ErrBadConn)
	//	if !isBadConn {
	//		break
	//	}
}

//func (db *DB) SetConnMaxIdleTime(d time.Duration)
//func (db *DB) SetConnMaxLifetime(d time.Duration)
//func (db *DB) SetMaxIdleConns(n int)
//func (db *DB) SetMaxOpenConns(n int)
