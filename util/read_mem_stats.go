package util

import (
	"log"
	"runtime"
)

func ReadMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)\n", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
