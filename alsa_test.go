package alsa

import (
	"testing"
	"fmt"
)

func TestParse(t *testing.T) {
	handle := New()
	err := handle.Open("default", StreamTypePlayback, ModeBlock)
	if err != nil {
		t.Fatalf("Open failed. %s", err)
	}

	handle.SampleFormat = SampleFormatS16LE
	handle.SampleRate = 44100
	handle.Channels = 2
	err = handle.ApplyHwParams()
	if err != nil {
		t.Fatalf("SetHwParams failed. %s", err)
	}

	buf := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	wrote, err := handle.Write(buf)
	if err != nil {
		t.Fatalf("Writei failed. %s", err)
	} else {
		fmt.Printf("Wrote %d\n", wrote)
	}

	handle.Close()
}
