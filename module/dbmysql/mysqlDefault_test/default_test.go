package optDefault

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/module/dbmysql/mysqlDefault"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_opt_Default(t *testing.T) {
	opt := mysqlDefault.Opt()
	assert.NotEmpty(t, opt)
}
