package huffman

import (
	"math"
	"sort"
)

type node struct {
	Score float64
	left *node
	right *node
}

type queue []node

func BuildTree(frequencies []float64) node {
	sort.Float64s(frequencies)
	q1 := initializeQueue(frequencies)
	q2 := queue{}

	for (len(q1) > 1 || len(q2) > 1) {
		if len(q1) > 0 {
			minNode1 := q1[0]
			minNode2 := q1[0]
			if len(q1) > 1 {
				minNode2 = q1[1]
			} else {
				minNode2 = q2[0]
			}

			if len(q2) > 0 {
				if q2[0].Score < minNode1.Score {
					minNode2 = minNode1
					minNode1 = q2[0]
				}
				if len(q2) > 1 {
					if q2[1].Score < q1[0].Score {
						minNode2 = q2[1]
					}
				}
			}

			if minNode1 == q1[0] {
				q1 = q1[1:]
				if len(q2) > 0 && minNode2 == q2[0] {
					q2 = q2[1:]
				} else {
					q1 = q1[1:]
				}
			} else if len(q2) > 0 && minNode1 == q2[0] {
				q2 = q2[1:]
				if len(q2) > 1 {
					if minNode2 == q2[1] {
						q2 = q2[1:]
					} else {
						q1 = q1[1:]
					}
				} else {
					q1 = q1[1:]
				}
			}

			// build new tree enqueue to q2
			newNode := node { (minNode1.Score + minNode2.Score), &minNode1, &minNode2 }
			q2 = append(q2, newNode)
		} else {
			minNode1 := q2[0]
			minNode2 := q2[1]
			q2 = q2[2:]

			// build new true enqueue to q1
			newNode := node { (minNode1.Score + minNode2.Score), &minNode1, &minNode2 }
			q2 = append(q2, newNode)
		}
	}

	if len(q1) > 0 {
		return q1[0]
	} else {
		return q2[0]
	}
}

func (tree *node) findMinAndMaxDepth() (int, int) {
	if tree == nil {
		return 0, 0
	}

	leftMin, leftMax := tree.left.findMinAndMaxDepth()
	rightMin, rightMax := tree.right.findMinAndMaxDepth()

	return 1 + int(math.Min(float64(leftMin), float64(rightMin))), 1 + int(math.Max(float64(leftMax), float64(rightMax)))
}

func (tree *node) FindMinAndMaxLengthCodeWord() (int, int) {
	minDepth, maxDepth := tree.findMinAndMaxDepth()

	return minDepth - 1, maxDepth - 1
}

func (tree node) isLeaf() bool {
	return tree.left == nil && tree.right == nil
}

func initializeQueue(values []float64) queue {
	q := make(queue, len(values))

	for i, v := range(values) {
		q[i] = node{ v, nil, nil }
	}

	return q
}