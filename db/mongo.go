package db

import (
	"gopkg.in/mgo.v2"
)

//mongo
type DBConnection struct {
	session *mgo.Session
}

func NewConnection(host string) (conn *DBConnection) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}

	return conn
}

func (conn *DBConnection) use(dbName, tableName string) (session *mgo.Session, collection *mgo.Collection) {
	copySession := conn.session.Copy()
	return copySession, copySession.DB(dbName).C(tableName)
}

func (conn *DBConnection) Close() {
	conn.session.Close()
	return
}

func (conn *DBConnection) EnsureIndex(dbName string, tableName string, index mgo.Index) {
	ms, c := conn.use(dbName, tableName)
	defer ms.Close()
	c.EnsureIndex(index)
}

func (conn *DBConnection) EnsureIndexKey(dbName string, tableName string, key ...string) {
	ms, c := conn.use(dbName, tableName)
	defer ms.Close()
	c.EnsureIndexKey(key...)
}

func (conn *DBConnection) IsEmpty(db, collection string) bool {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	count, _ := c.Count()
	return count == 0
}

func (conn *DBConnection) Count(db, collection string, query interface{}) (int, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func (conn *DBConnection) Insert(db, collection string, docs ...interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func (conn *DBConnection) FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func (conn *DBConnection) FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func (conn *DBConnection) FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func (conn *DBConnection) FindIter(db, collection string, query, selector interface{}) *mgo.Iter {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Iter()
}

func (conn *DBConnection) PageCursor(db, collection string, page, limit int, query, selector interface{}, result []interface{}) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	iter := c.Find(query).Select(selector).Skip(page * limit).Limit(limit).Iter()
	var r interface{}
	if iter.Next(&r) {
		result = append(result, r)
	}
}

func (conn *DBConnection) Update(db, collection string, selector, update interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func (conn *DBConnection) Upsert(db, collection string, selector, update interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func (conn *DBConnection) UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func (conn *DBConnection) Remove(db, collection string, selector interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func (conn *DBConnection) RemoveAll(db, collection string, selector interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

//insert one or multi documents
func (conn *DBConnection) BulkInsert(db, collection string, docs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func (conn *DBConnection) BulkRemove(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func (conn *DBConnection) BulkRemoveAll(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func (conn *DBConnection) BulkUpdate(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func (conn *DBConnection) BulkUpdateAll(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func (conn *DBConnection) BulkUpsert(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func (conn *DBConnection) PipeAll(db, collection string, pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.All(result)
}

func (conn *DBConnection) PipeOne(db, collection string, pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}

func (conn *DBConnection) PipeIter(db, collection string, pipeline interface{}, allowDiskUse bool) *mgo.Iter {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}

	return pipe.Iter()

}

func (conn *DBConnection) Explain(db, collection string, pipeline, result interface{}) error {
	ms, c := conn.use(db, collection)
	defer ms.Close()
	pipe := c.Pipe(pipeline)
	return pipe.Explain(result)
}
