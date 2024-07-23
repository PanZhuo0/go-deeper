package model

import "gorm.io/gen"

// custom by append annotation
// GEN 框架读取注释，自动帮你生成一个注释下的方法的实现
type Querier interface {
	// SELECT * FROM @@table where id=@id
	GetByID(id int) (gen.T, error) //return struct & error

	// SELECT * FROM @@table WHERE id=@id
	GetByIDReturnMap(id int) (gen.M, error) //return map & error

	// SELECT * FROM @@table WHERE author=@author
	GetBooksByAuthor(author string) ([]*gen.T, error) //return struct slice & error
}
