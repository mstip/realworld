package main

import "testing"

func TestCreateTag(t *testing.T) {
	ts := NewTagService(NewSeededInMemDB())

	einneuertag, err := ts.CreateTag("einneuertag")
	if err != nil {
		t.Fatalf("failed to create tag %#v", err)
	}

	if einneuertag.ID == 0 {
		t.Errorf("id is %d", einneuertag.ID)
	}

	if einneuertag.Tag != "einneuertag" {
		t.Errorf("Tag should %#v but is %#v", "einneuertag", einneuertag.Tag)
	}

	sametagagain, err := ts.CreateTag("einneuertag")
	if err != nil {
		t.Fatalf("failed to create tag %#v", err)
	}

	if einneuertag.ID != sametagagain.ID {
		t.Errorf("a new tag was created but shouldnt")
	}
}
