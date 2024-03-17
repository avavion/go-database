package mysql

import "time"

type Mysql struct {
	maxPoolSize        int
	connectionAttempts int
	connectionTimeout  time.Duration
}

func New(url string) {

}

func (mysql *Mysql) Close() {

}
