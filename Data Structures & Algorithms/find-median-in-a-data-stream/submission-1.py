import heapq
class MedianFinder:

    def __init__(self):
        self._bottom = []
        self._top = []

    def addNum(self, num: int) -> None:
        if not self._top:
            heapq.heappush(self._top, num)
            return
        if num < self._top[0]:
            heapq.heappush(self._bottom, -num)
        else:
            heapq.heappush(self._top, num)
        
        # Keep it balanced
        if len(self._bottom) > len(self._top)+1:
            heapq.heappush(self._top, -heapq.heappop(self._bottom))
            return
        if len(self._top) > len(self._bottom)+1:
            heapq.heappush(self._bottom, -heapq.heappop(self._top))
            return
    def findMedian(self) -> float:
        if len(self._bottom) == len(self._top):
            return (-self._bottom[0] + self._top[0]) / 2.0
        elif len(self._bottom) > len(self._top):
            return -self._bottom[0]
        else:
            return self._top[0]   