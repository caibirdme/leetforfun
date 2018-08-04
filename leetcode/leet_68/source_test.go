package leet_68

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFullJustify(t *testing.T) {
	var testCase = []struct {
		words    []string
		maxWidth int
		expect   []string
	}{
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expect: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expect: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},

		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
				"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expect: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := fullJustify(tc.words, tc.maxWidth)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
