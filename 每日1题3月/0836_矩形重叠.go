package 每日1题3月

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	ret := rec2[1] >= rec1[3] || rec2[3] <= rec1[1] || rec2[2] <= rec1[0] || rec2[0] >= rec1[2]
	return !ret
}
