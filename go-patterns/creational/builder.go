package creational

import (
	"errors"
	"time"
)

// DBPool stands for database connection pool
type DBPool struct {
	dsn             string
	maxOpenConn     int
	maxIdleConn     int
	maxConnLifeTime time.Duration
}

// DBPoolBuilder builds a DBPool
type DBPoolBuilder struct {
	DBPool
	err error
}

// Builder create a new DBPoolBuilder instance
func Builder() *DBPoolBuilder {
	return &DBPoolBuilder{
		DBPool: DBPool{
			dsn:             "127.0.0.1:3306",
			maxOpenConn:     30,
			maxIdleConn:     5,
			maxConnLifeTime: time.Hour,
		},
	}
}

// Build create a DBPool through DBPoolBuilder
func (b *DBPoolBuilder) Build() (*DBPool, error) {
	if b.DBPool.maxOpenConn < b.DBPool.maxIdleConn {
		return nil, errors.New("maxOpenConn must be greater than maxIdleConn")
	}
	return &b.DBPool, b.err
}

// DSN setup DBPool dns connection string
func (b *DBPoolBuilder) DSN(dsn string) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if dsn == "" {
		b.err = errors.New("invalid dsn string")
	}
	b.DBPool.dsn = dsn
	return b
}

// MaxOpenConn setup DBPool max open connection
func (b *DBPoolBuilder) MaxOpenConn(maxOpenConn int) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if maxOpenConn <= 0 {
		b.err = errors.New("invalid maxOpenConn")
	}
	b.DBPool.maxOpenConn = maxOpenConn
	return b
}

// MaxIdleConn setup DBPool max idle connection
func (b *DBPoolBuilder) MaxIdleConn(maxIdleConn int) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if maxIdleConn <= 0 {
		b.err = errors.New("invalid maxIdleConn")
	}
	b.DBPool.maxIdleConn = maxIdleConn
	return b
}

// MaxConnLifeTime setup DBPool max connection life time
func (b *DBPoolBuilder) MaxConnLifeTime(maxConnLifeTime time.Duration) *DBPoolBuilder {
	if b.err != nil {
		return b
	}
	if maxConnLifeTime <= 0 {
		b.err = errors.New("invalid maxConnLifeTime")
	}
	b.DBPool.maxConnLifeTime = maxConnLifeTime
	return b
}
