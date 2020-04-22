package optDefault

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/optDefault"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_optDefault_Opt(t *testing.T) {
	opt := optDefault.Opt()
	assert.NotEmpty(t, opt)
	_, err := opt.Ping()
	assert.Empty(t, err)
}
