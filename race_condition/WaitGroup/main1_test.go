package main1

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello World!"

	wg.Add(1)

	go updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}

//When running from the cli
//use go test -race . (To check the race condition in colsole)
