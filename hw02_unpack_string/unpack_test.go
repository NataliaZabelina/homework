package hw02_unpack_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "ala0b",
			expected: "alb",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "v2yhK3ujj",
			expected: "vvyhKKKujj",
		},
		{
			input:    "a4bc2д5e",
			expected: "aaaabccдддддe",
		},
		{
			input:    "O4",
			expected: "OOOO",
		},
		{
			input:    "Зеленогла4зоеТакси",
			expected: "ЗеленоглаааазоеТакси",
		},
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "x",
			expected: "x",
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "9",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "0",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "z0",
			expected: "",
		},
		{
			input:    "z00",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    " ",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "#abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "Hello,",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "abc kLi",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "'low'",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    `res2t`,
			expected: "resst",
		},
		{
			input:    `re\n2t`,
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
		{
			input:    "\n3ipo",
			expected: "\n\n\nipo",
		},
		{
			input:    "\n",
			expected: "\n",
		},
		{
			input:    "\n2",
			expected: "\n\n",
		},
		{
			input:    "\n0",
			expected: "",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackWithEscape(t *testing.T) {
	t.Skip() // NeedRemove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
