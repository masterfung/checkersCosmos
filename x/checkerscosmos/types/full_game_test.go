package types_test

import (
	// "strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/masterfung/checkersCosmos/testutil"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/rules"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
		Black: alice,
		Red:   bob,
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
		BeforeIndex: types.NoFifoIndex,
		AfterIndex: types.NoFifoIndex,
		Deadline: types.DeadlineLayout,
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	aliceAddr, err1 := sdk.AccAddressFromBech32(alice)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, aliceAddr, black)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4" // Bad last digit
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t,
		err,
		"Black address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h" // Bad last digit
	red, err := storedGame.GetRedAddress()
	require.Nil(t, red)
	require.EqualError(t,
		err,
		"Red address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h: decoding bech32 failed: invalid checksum (expected xqhc8g got xqhc8h)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseDeadlineCorrect(t *testing.T) {
	deadline, err := GetStoredGame1().GetDeadlineAsTime()
	require.Nil(t, err)
	require.Equal(t, time.Time(time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)), deadline)
}

func TestParseDeadlineMissingMonth(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Deadline = "2006-02 15:04:05.999999999 +0000 UTC"
	_, err := storedGame.GetDeadlineAsTime()
	require.EqualError(t,
		err,
		"deadline cannot be parsed: 2006-02 15:04:05.999999999 +0000 UTC: parsing time \"2006-02 15:04:05.999999999 +0000 UTC\" as \"2006-01-02 15:04:05.999999999 +0000 UTC\": cannot parse \" 15:04:05.999999999 +0000 UTC\" as \"-\"")
	require.EqualError(t, storedGame.Validate(), err.Error())
}
