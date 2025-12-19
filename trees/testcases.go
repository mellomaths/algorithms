package trees

type traversalTreeTestAsserts struct {
	key           int
	expected      []int
	expectedError error
	name          string
}

type traversalTreeTest struct {
	root    *TreeNode[int]
	asserts traversalTreeTestAsserts
}

type propertyTreeTestAsserts struct {
	key           int
	expected      int
	expectedError error
	name          string
}

type propertyTreeTest struct {
	root    *TreeNode[int]
	asserts propertyTreeTestAsserts
}

var dfsInOrderTests = []traversalTreeTest{
	{
		&TreeNode[int]{
			Value: 1,
			Left: &TreeNode[int]{
				Value: 2,
				Left: &TreeNode[int]{
					Value: 4,
				},
				Right: &TreeNode[int]{
					Value: 5,
				},
			},
			Right: &TreeNode[int]{
				Value: 3,
				Left: &TreeNode[int]{
					Value: 6,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
		traversalTreeTestAsserts{
			key:           1,
			expected:      []int{4, 2, 5, 1, 6, 3, 7},
			expectedError: nil,
			name:          "InOrder",
		},
	},
}

var dfsPreOrderTests = []traversalTreeTest{
	{
		&TreeNode[int]{
			Value: 1,
			Left: &TreeNode[int]{
				Value: 2,
				Left: &TreeNode[int]{
					Value: 4,
				},
				Right: &TreeNode[int]{
					Value: 5,
				},
			},
			Right: &TreeNode[int]{
				Value: 3,
				Left: &TreeNode[int]{
					Value: 6,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
		traversalTreeTestAsserts{
			key:           1,
			expected:      []int{1, 2, 4, 5, 3, 6, 7},
			expectedError: nil,
			name:          "PreOrder",
		},
	},
}

var dfsPostOrderTests = []traversalTreeTest{
	{
		&TreeNode[int]{
			Value: 1,
			Left: &TreeNode[int]{
				Value: 2,
				Left: &TreeNode[int]{
					Value: 4,
				},
				Right: &TreeNode[int]{
					Value: 5,
				},
			},
			Right: &TreeNode[int]{
				Value: 3,
				Left: &TreeNode[int]{
					Value: 6,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
		traversalTreeTestAsserts{
			key:           1,
			expected:      []int{4, 5, 2, 6, 7, 3, 1},
			expectedError: nil,
			name:          "PostOrder",
		},
	},
}

var bfsLevelOrderTests = []traversalTreeTest{
	{
		&TreeNode[int]{
			Value: 1,
			Left: &TreeNode[int]{
				Value: 2,
				Left: &TreeNode[int]{
					Value: 4,
				},
				Right: &TreeNode[int]{
					Value: 5,
				},
			},
			Right: &TreeNode[int]{
				Value: 3,
				Left: &TreeNode[int]{
					Value: 6,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
		traversalTreeTestAsserts{
			key:           1,
			expected:      []int{1, 2, 3, 4, 5, 6, 7},
			expectedError: nil,
			name:          "BfsLevelOrder",
		},
	},
	{
		&TreeNode[int]{
			Value: 5,
			Left: &TreeNode[int]{
				Value: 12,
				Left: &TreeNode[int]{
					Value: 7,
					Left: &TreeNode[int]{
						Value: 17,
					},
					Right: &TreeNode[int]{
						Value: 23,
					},
				},
			},
			Right: &TreeNode[int]{
				Value: 13,
				Left: &TreeNode[int]{
					Value: 14,
					Left: &TreeNode[int]{
						Value: 27,
					},
					Right: &TreeNode[int]{
						Value: 3,
					},
				},
				Right: &TreeNode[int]{
					Value: 2,
					Left: &TreeNode[int]{
						Value: 8,
					},
					Right: &TreeNode[int]{
						Value: 11,
					},
				},
			},
		},
		traversalTreeTestAsserts{
			key:           2,
			expected:      []int{5, 12, 13, 7, 14, 2, 17, 23, 27, 3, 8, 11},
			expectedError: nil,
			name:          "BfsLevelOrder",
		},
	},
}

var flipTreeClockwiseTests = []traversalTreeTest{
	{
		&TreeNode[int]{
			Value: 1,
			Left: &TreeNode[int]{
				Value: 2,
				Left: &TreeNode[int]{
					Value: 4,
				},
				Right: &TreeNode[int]{
					Value: 5,
				},
			},
			Right: &TreeNode[int]{
				Value: 3,
				Left: &TreeNode[int]{
					Value: 6,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
		traversalTreeTestAsserts{
			key:           1,
			expected:      []int{4, 5, 2, 3, 1, 6, 7},
			expectedError: nil,
			name:          "FlipTreeClockwise",
		},
	},
	{
		&TreeNode[int]{
			Value: 5,
			Left: &TreeNode[int]{
				Value: 12,
				Left: &TreeNode[int]{
					Value: 7,
					Left: &TreeNode[int]{
						Value: 17,
					},
					Right: &TreeNode[int]{
						Value: 23,
					},
				},
			},
			Right: &TreeNode[int]{
				Value: 13,
				Left: &TreeNode[int]{
					Value: 14,
					Left: &TreeNode[int]{
						Value: 27,
					},
					Right: &TreeNode[int]{
						Value: 3,
					},
				},
				Right: &TreeNode[int]{
					Value: 2,
					Left: &TreeNode[int]{
						Value: 8,
					},
					Right: &TreeNode[int]{
						Value: 11,
					},
				},
			},
		},
		traversalTreeTestAsserts{
			key:           2,
			expected:      []int{17, 23, 7, 12, 13, 5, 14, 2, 27, 3, 8, 11},
			expectedError: nil,
			name:          "BfsLevelOrder",
		},
	},
}

var maximumHeightOfTreeTests = []propertyTreeTest{
	{
		&TreeNode[int]{
			Value: 1,
			Left: &TreeNode[int]{
				Value: 2,
				Left: &TreeNode[int]{
					Value: 4,
				},
				Right: &TreeNode[int]{
					Value: 5,
				},
			},
			Right: &TreeNode[int]{
				Value: 3,
				Left: &TreeNode[int]{
					Value: 6,
				},
				Right: &TreeNode[int]{
					Value: 7,
				},
			},
		},
		propertyTreeTestAsserts{
			key:           1,
			expected:      3,
			expectedError: nil,
			name:          "TreeHeight",
		},
	},
	{
		&TreeNode[int]{
			Value: 5,
			Left: &TreeNode[int]{
				Value: 12,
				Left: &TreeNode[int]{
					Value: 7,
					Left: &TreeNode[int]{
						Value: 17,
					},
					Right: &TreeNode[int]{
						Value: 23,
					},
				},
			},
			Right: &TreeNode[int]{
				Value: 13,
				Left: &TreeNode[int]{
					Value: 14,
					Left: &TreeNode[int]{
						Value: 27,
					},
					Right: &TreeNode[int]{
						Value: 3,
					},
				},
				Right: &TreeNode[int]{
					Value: 2,
					Left: &TreeNode[int]{
						Value: 8,
					},
					Right: &TreeNode[int]{
						Value: 11,
					},
				},
			},
		},
		propertyTreeTestAsserts{
			key:           2,
			expected:      4,
			expectedError: nil,
			name:          "TreeHeight",
		},
	},
	{
		nil,
		propertyTreeTestAsserts{
			key:           3,
			expected:      0,
			expectedError: ErrEmptyTree,
			name:          "HeightOfEmptyTree",
		},
	},
}
