package leet_863

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

const null = 999999

func createTree(nums []int) *TreeNode {
	length := len(nums)
	if 0 == length {
		return nil
	}
	q := new(list)
	root := &TreeNode{Val: nums[0]}
	q.Push(root)
	for i := 1; !q.Empty() && i < length; i++ {
		node := q.Pop()
		if nums[i] != null {
			node.Left = &TreeNode{
				Val: nums[i],
			}
			q.Push(node.Left)
		}
		i++
		if i < length && nums[i] != null {
			node.Right = &TreeNode{
				Val: nums[i],
			}
			q.Push(node.Right)
		}
	}
	return root
}

type list struct {
	buf []*TreeNode
}

func (l *list) Pop() *TreeNode {
	if len(l.buf) == 0 {
		return nil
	}
	t := l.buf[0]
	l.buf = l.buf[1:]
	return t
}

func (l *list) Push(r *TreeNode) {
	l.buf = append(l.buf, r)
}

func (l *list) Empty() bool {
	return len(l.buf) == 0
}

func TestDistanceK(t *testing.T) {
	var testCase = []struct {
		nums   []int
		target int
		k      int
		expect []int
	}{
		{
			nums:   []int{0, 216, 1, 411, 251, 2, 9, null, null, 264, null, 14, 3, 18, 35, 476, null, 21, 362, 6, 4, 84, 39, 54, 48, null, null, 32, 27, null, null, 7, 19, 11, 5, null, 112, 56, 49, 77, 180, 74, null, 37, null, 59, 100, 20, 8, 266, 26, 16, 24, 15, 12, 138, 260, null, 170, 51, 302, null, 85, null, 233, 101, 454, 122, 61, 81, 345, 245, 154, 31, 66, 13, 10, null, 395, null, 36, 33, 17, 44, 38, 91, 30, 34, 40, 333, 227, 446, 343, 218, null, 92, 70, null, null, null, 99, 497, null, 468, 102, null, null, 327, 147, 117, 62, 104, 87, null, 372, 257, null, null, 294, 96, 57, 75, 413, 25, 63, 42, 90, null, null, 52, 64, 45, 68, 22, 23, 76, 116, null, 60, null, 151, 238, 228, 78, 131, null, 134, null, 451, null, 382, null, null, 469, null, 348, null, 201, 111, 492, 298, 267, 172, null, null, null, null, 182, null, null, null, 250, 186, 305, 220, 80, 86, 175, null, 385, 97, 416, 429, null, null, null, 383, null, 160, null, 120, 225, 262, null, null, 124, 28, 118, 110, 71, null, 103, 98, 82, 162, 241, 73, 121, null, null, 265, 46, 69, 41, 55, null, null, 129, 288, 126, 105, null, 152, 428, 408, null, null, 83, 143, 244, 312, null, 214, null, null, null, 488, null, null, 452, null, 208, 311, 287, 141, null, null, 314, null, 270, null, 291, 253, null, null, 309, 338, null, 423, null, null, 277, 299, 135, 156, 114, 443, null, 356, 387, null, 200, 153, null, 461, null, 433, null, null, 290, null, 276, 352, 306, 240, null, null, 188, 169, 29, 53, 165, 178, null, 132, 194, 232, 316, 205, 158, 296, 211, 93, null, 179, null, null, null, 419, 146, null, 272, 279, 89, 50, 88, 246, 168, 119, 72, 58, 231, 140, 450, 417, 230, 176, 315, 149, 207, null, null, null, null, 409, 95, 148, 167, 489, null, 328, 368, 359, null, null, null, null, null, 467, 217, null, null, 355, null, 330, null, null, null, null, null, 426, 394, null, 334, null, null, null, null, null, null, null, 278, 357, null, null, 185, null, null, null, 133, null, null, null, null, null, null, null, 331, null, 174, 366, null, null, null, 439, 459, null, null, null, null, null, 407, 367, null, 318, 301, null, null, 171, 43, 47, 199, 144, 379, 249, null, null, 166, 444, null, null, null, 310, null, null, null, null, 203, 226, null, null, null, 256, 195, null, null, null, null, null, 213, 285, null, 404, null, 472, 145, 177, 123, null, 127, 106, 397, null, 424, 346, null, 173, 109, 79, 65, 137, 393, null, null, 283, null, null, null, null, 386, null, 204, 193, null, 353, 187, null, 415, 300, null, null, 161, null, 189, 181, 190, null, null, null, 449, null, null, 398, null, null, null, null, null, 274, null, null, null, null, null, 442, null, null, null, null, 284, null, 402, null, 215, null, null, 336, null, null, 470, 317, null, null, null, null, null, null, null, null, null, null, 323, null, null, 341, null, null, 67, 155, 197, 198, null, 297, null, 435, null, null, 354, 255, null, 375, 480, null, null, null, 360, 252, 374, null, 406, 482, null, null, 286, null, null, null, null, null, null, null, null, 191, 271, null, 282, null, 235, 498, 273, 107, 457, null, null, null, null, 400, 361, 380, 125, 222, 128, 108, 94, 329, 434, 295, null, 437, null, null, null, null, null, 219, null, 462, null, 445, 236, 319, null, null, 339, null, 163, 326, 209, null, 392, 258, 196, null, null, null, null, 412, 471, 499, null, 448, 324, null, null, null, 391, 242, null, 475, null, null, null, null, null, null, 390, null, 261, 130, 237, 184, 313, null, 281, 487, null, null, null, null, null, 425, null, null, null, null, null, null, 414, 466, null, 453, null, null, null, null, null, null, 292, 478, 192, null, 293, 303, null, 370, 248, null, null, null, 430, null, 139, 304, 465, 477, null, null, null, 364, null, null, null, 229, 320, null, 183, null, 212, 113, 496, null, null, 340, null, null, 381, null, null, null, 221, null, null, null, null, null, 335, 388, null, null, null, 363, null, 440, null, null, 325, 247, null, null, null, null, 289, null, null, 432, null, null, null, null, null, null, null, 484, null, null, 321, 269, null, null, null, null, null, 373, 485, 403, 344, 243, 254, 202, null, 473, 347, 455, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, 308, null, null, null, null, null, null, null, null, null, 142, null, null, 369, null, null, null, null, 436, null, 376, 420, null, null, null, null, null, null, 115, 263, null, null, null, 399, null, null, 371, 234, 351, null, 474, null, 441, null, null, null, 483, 418, null, null, 358, 427, null, null, null, null, 349, null, 365, null, 389, null, null, null, null, null, 384, 438, null, 410, null, null, 378, 350, null, 479, null, null, null, 464, null, 332, 405, null, null, null, null, null, null, null, null, 447, 136, 164, 280, null, 401, null, 456, null, null, 422, null, null, null, null, null, null, null, null, 486, null, null, null, null, 493, null, null, null, null, null, null, 458, null, null, null, null, null, null, null, null, 490, null, null, null, 494, null, null, null, null, null, null, 150, 157, 206, 396, null, 337, null, null, null, 495, null, 481, null, null, null, null, null, null, null, null, null, null, 159, 210, 307, 377, 259, null, null, 421, null, null, null, null, null, null, 268, null, 431, 223, null, 342, null, 491, null, null, null, null, null, null, null, 460, 275, 224, null, null, null, null, null, null, null, null, 322, 239, 463},
			target: 461,
			k:      16,
			expect: []int{7, 4, 1},
		},
		{
			nums:   []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
			target: 5,
			k:      2,
			expect: []int{7, 4, 1},
		},
		{
			nums:   []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
			target: 5,
			k:      1,
			expect: []int{3, 6, 2},
		},
		{
			nums:   []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
			target: 5,
			k:      3,
			expect: []int{0, 8},
		},
		{
			nums:   []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
			target: 5,
			k:      0,
			expect: []int{5},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		root := createTree(tc.nums)
		tar := &TreeNode{Val: tc.target}
		actual := distanceK(root, tar, tc.k)
		sort.Ints(actual)
		sort.Ints(tc.expect)
		should.Equal(tc.expect, actual, "case$%d failed", idx)
	}
}
