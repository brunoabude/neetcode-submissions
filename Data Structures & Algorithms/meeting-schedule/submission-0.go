/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func quicksort(intervals []Interval, low, high int) {
	if low >= high {
		return
	}

	p := partition(intervals, low, high)
	quicksort(intervals, low, p-1)
	quicksort(intervals, p+1, high)
}

func partition(intervals []Interval, low, high int) int {
	pivot := high

	i := low - 1

	for j := low; j < high; j++ {
		if intervals[j].start < intervals[pivot].start {
			i++
			intervals[i], intervals[j] = intervals[j], intervals[i]
		}
	}

	i++
	intervals[i], intervals[pivot] = intervals[pivot], intervals[i]

	return i
}

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) <= 1 {
		return true
	}

	quicksort(intervals, 0, len(intervals)-1)

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1].start < intervals[i].end {
			return false
		}
	}

	return true
}
