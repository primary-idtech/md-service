package fix

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLogonMsg_true(t *testing.T) {
	fixTxt := "8=FIXT.1.1|9=70|35=A|34=1|49=CLIENT12|52=20100225-19:41:57.316|56=B|98=0|108=30|10=128|"
	fixTxt = strings.Replace(fixTxt, "|", "\x01", -1)

	assert.True(t, isLogonMsg([]byte(fixTxt)))
}

func TestIsLogonMsg_false(t *testing.T) {
	fixTxt := "8=FIXT.1.1|9=70|35=D|34=1|49=CLIENT12|52=20100225-19:41:57.316|56=B|98=0|108=30|10=128|"
	fixTxt = strings.Replace(fixTxt, "|", "\x01", -1)

	assert.False(t, isLogonMsg([]byte(fixTxt)))
}
