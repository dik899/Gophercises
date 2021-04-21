package primitive


import (
	"testing"

	

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	_, err := Transform("../sample.jpeg", 3, 120, "jpeg")
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equalf(t,nil, err, "they should be equal")
}

func TestBatchTransformMode(t *testing.T) {
	testOutput := BatchTransform("../sample.jpeg", []int{0, 1, 2, 3}, []int{50}, "jpeg")
	assert.Equalf(t, len(testOutput), 4, "they should be equal")
}

func TestBatchTransformNumber(t *testing.T) {
	testOutput := BatchTransform("../sample.jpeg", []int{0}, []int{50, 100, 150, 200}, "jpeg")
	assert.Equalf(t, len(testOutput), 4, "they should be equal")
}

// func TestMain(m *testing.M) {
// 	dashtest.ControlCoverage(m)
// 	m.Run()
// }
