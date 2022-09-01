package main

/**
1560.在既定时间做作业的学生人数(Easy)
*/
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0
	for i, v := range startTime {
		if queryTime >= v && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}

//func busyStudent(startTime []int, endTime []int, queryTime int) int {
//	res := 0
//	for i := 0; i < len(startTime); i++ {
//		if startTime[i] < queryTime && queryTime > endTime[i] {
//			res++
//		}
//	}
//	return res
//}
