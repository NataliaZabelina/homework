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
	cases := []struct {
		input string
		expected string
		err error
	}{
		{"", "", nil},
		{"ala0b", "alb", nil},
		{"aaa0b", "aab", nil},
		{"v2yhK3ujj", "vvyhKKKujj", nil},
		{"a4bc2д5e", "aaaabccдддддe", nil},
		{"O4", "OOOO", nil},
		{"Зеленогла4зоеТакси", "ЗеленоглаааазоеТакси", nil},
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abccd", "abccd", nil},
		{"45", "", ErrInvalidString},
		{"x", "x", nil},
		{"9", "", ErrInvalidString},
		{"0", "", ErrInvalidString},
		{"z0", "", nil},
		{"z00", "", ErrInvalidString},
		{" ", "", ErrInvalidString},
		{"aaa10b", "", ErrInvalidString},
		{"#abc", "", ErrInvalidString},
		{"Hello,", "", ErrInvalidString},
		{"abc kLi", "", ErrInvalidString},
		{"'low',", "", ErrInvalidString},
		{`res2t`, "resst", nil},
		{`re\n2t`, "", ErrInvalidString},
		{"d\n5abc", "d\n\n\n\n\nabc", nil},
		{"\n3ipo", "\n\n\nipo", nil},
		{"\n", "\n", nil},
		{"\n2", "\n\n", nil},
		{"\n0", "", nil},
	}

	for _, tc := range cases {
	tc := tc
	t.Run(tc.input, func(t *testing.T) {
		result, err := Unpack(tc.input)
		require.Equal(t, tc.err, err)
		require.Equal(t, tc.expected, result)
		})
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
