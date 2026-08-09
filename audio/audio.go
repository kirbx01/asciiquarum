package audio

import (
	_ "embed"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var Music []byte

func AutoPlay() {
	tmp, err := os.CreateTemp("", "asciiquarum-*.mp3")
	if err != nil {
		return
	}
	defer os.Remove(tmp.Name())

	if _, err := io.Copy(tmp, bytes.NewReader(Music)); err != nil {
		return
	}
	if err := tmp.Close(); err != nil {
		return
	}

	cmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", tmp.Name())
	if err := cmd.Start(); err == nil {
		go func() {
			if err := cmd.Wait(); err != nil {
				fmt.Fprintln(os.Stderr, "audio ffplay error:", err)
			}
		}()
		return
	}

	cmd = exec.Command("mpv", "--no-video", "--quiet", tmp.Name())
	if err := cmd.Start(); err == nil {
		go func() {
			if err := cmd.Wait(); err != nil {
				fmt.Fprintln(os.Stderr, "audio mpv error:", err)
			}
		}()
		return
	}

	fmt.Fprintln(os.Stderr, "audio: no player could start")
}
