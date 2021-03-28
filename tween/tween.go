// animations
package tween

import "time"

//import "log"

var animSleep time.Duration = 100 //100 FPS is the best we are going to get

//Starts a clock running.  The thread will update the variable pointed to by v, moving the value from *start* to *end*.
//The animation with start at startTime.  The animation loops forever.
func Clock(v *float32, start, end, duration float32, startTime time.Time) {
	time.Sleep(animSleep * time.Millisecond)
	t := float32(time.Since(startTime).Seconds()) / duration
	*v = start + t*(end-start)
	if t > 1.0 {
		t = 0.0
	}
	Clock(v, start, end, duration, startTime)
}

//Similar to clock, but does not repeat
func Linear(v *float32, start, end, duration float32, startTime time.Time) {
	time.Sleep(animSleep * time.Millisecond)
	t := float32(time.Since(startTime).Seconds()) / duration
	*v = start + t*(end-start)
	if t > 1.0 {
		return
	} else {
		Linear(v, start, end, duration, startTime)
	}
}

func StartLinear(v *float32, start, end, duration float32) {
	go Linear(v, start, end, duration, time.Now())
}

//Starts an immediate clock, animating v from start to end
func StartClock(v *float32, start, end, duration float32) {
	go Clock(v, start, end, duration, time.Now())
}
