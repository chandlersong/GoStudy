package pattern

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {

	t.Log("test singleton pattern")
	s := GetInstance()
	s.value = 100
	t.Log(s.value)
	s.value = 200

	s2 := GetInstance()
	t.Log(s2.value)
	s2.value = 300

	assert.Equal(t, s.value, s2.value)
	t.Log("test singleton finished")
}

type singleton struct {
	value int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
