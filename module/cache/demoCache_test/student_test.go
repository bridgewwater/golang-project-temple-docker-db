package demoCache_test

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/demoCache"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/demo"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testKey = "000123"

func Test_demoCache_Set(t *testing.T) {
	f := fStudent{}
	err := faker.FakeData(&f)
	if err != nil {
		t.Errorf("FakeData Err: %v", err)
	}
	err = demoCache.Set(testKey,
		&demo.Student{
			StuName: f.StuName,
		},
	)
	if err != nil {
		t.Errorf("demoCache.Set Err: %v", err)
	}
}

func Test_demoCache_Exists(t *testing.T) {
	hasStudent, err := demoCache.ExistsStudent(testKey)
	if err != nil {
		t.Errorf("demoCache.ExistsStudent Err: %v", err)
	}

	assert.True(t, hasStudent)
}

func Test_demoCache_Get(t *testing.T) {
	var findStudent demo.Student
	err := demoCache.Get(testKey, &findStudent)
	if err != nil {
		t.Errorf("demoCache.Get Err: %v", err)
	}
	t.Logf("demoCache.Get findStudent.StuName %v", findStudent.StuName)
	assert.NotEmpty(t, findStudent)
}
