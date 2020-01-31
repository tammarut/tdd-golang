package main

import "testing"

import "bytes"

import "reflect"

import "time"

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v but got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleep(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	config := ConfigurableSleeper{
		sleepTime,
		spyTime.Sleep,
	}
	config.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("shuold have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
