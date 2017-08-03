package searchFile

import "testing"

func TestRetrieveContents(t *testing.T) {
		var file = "_packageTest.json"
		_,err := RetrieveContents(file)
		if err != nil  {
			t.Errorf("RetrieveContents %q", err)
		}
}

func TestRetrieveDirectories(t *testing.T) {
	    var base = "C:/nginx/html"
		_,err := RetrieveDirectories(base)
		if err != nil {
			t.Errorf("RetrieveDirectories %q", err)
		}
}

