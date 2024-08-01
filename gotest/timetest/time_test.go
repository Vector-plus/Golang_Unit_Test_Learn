package timetest

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTool(t *testing.T) {
	af := time.Now().Add(time.Minute * 10)
	fmt.Println(af.Unix() - time.Now().Unix())
}
