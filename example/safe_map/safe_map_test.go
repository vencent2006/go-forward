package safe_map

import "testing"

func TestSafeMap_LoadOrStore(t *testing.T) {
	s := NewSafeMap[string, string]()
	key := "name"
	store, existed := s.LoadOrStore(key, "vincent")
	if existed {
		t.Logf("existed: %s %s", key, store)
	} else {
		t.Logf("not existed: %s %s", key, store)
	}

}
