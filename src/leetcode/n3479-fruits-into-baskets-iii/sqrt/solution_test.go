package sqrt

import (
	"LeetCode-Go/tst"
	"bufio"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_numOfUnplacedFruits1(t *testing.T) {
	assert.Equal(t, 1, numOfUnplacedFruits(tst.Str2List("[4,2,5]"), tst.Str2List("[3,5,4]")))
}

func Test_numOfUnplacedFruits2(t *testing.T) {
	assert.Equal(t, 0, numOfUnplacedFruits(tst.Str2List("[3,6,1]"), tst.Str2List("[6,4,7]")))
}

func Test_numOfUnplacedFruits(t *testing.T) {
	var testCases []struct {
		Fruits   string `json:"fruits"`
		Baskets  string `json:"baskets"`
		Expected int    `json:"expected"`
	}
	file, err := os.Open("test_input.json")
	require.NoError(t, err)
	defer file.Close()

	require.NoError(t, json.NewDecoder(bufio.NewReader(file)).Decode(&testCases))

	for _, ts := range testCases {
		done := make(chan int)
		go func() {
			defer close(done)
			assert.Equal(t, ts.Expected, numOfUnplacedFruits(tst.Str2List(ts.Fruits), tst.Str2List(ts.Baskets)))
		}()
		select {
		case <-done:
		case <-time.After(100 * time.Millisecond):
			t.Errorf("timeout")
		}
	}
}
