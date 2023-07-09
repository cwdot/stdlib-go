package wood

import "testing"


func Test_splitName(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		id      string
		display string
	}{
		{"Empty", "", "", ""},
		{"OriginalFormat", "Text", "", "Text"},
		{"Level1", "a.Text", "a", "Text"},
		{"Level2", "a.b.Text", "a.b", "Text"},
		{"Level3", "a.b.c.Text", "a.b.c", "Text"},
		{"Level4", "a.b.c.d.Text", "a.b.c.d", "Text"},
		{"Level5", "a.b.c.d.e.Text", "a.b.c.d.e", "Text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := splitName(tt.text)
			if ls.Id != tt.id {
				t.Errorf("splitName() id = %v, want %v", ls.Id, tt.id)
			}
			if ls.Display != tt.display {
				t.Errorf("splitName() display = %v, want %v", ls.Display, tt.display)
			}
		})
	}
}
func Test_ignored(t *testing.T) {
	PrefixLevel("a", ErrorLevel)
	PrefixLevel("a.b", WarnLevel)
	PrefixLevel("a.b.c", InfoLevel)
	PrefixLevel("a.b.c.d", DebugLevel)

	tests := []struct {
		name   string
		id     string
		action Level
		want   bool
	}{
		{"DebugLevel1", "a", DebugLevel, true},
		{"DebugLevel2", "a.b", DebugLevel, true},
		{"DebugLevel3", "a.b.c", DebugLevel, true},
		{"DebugLevel4", "a.b.c.d", DebugLevel, false},

		{"InfoLevel1", "a", InfoLevel, true},
		{"InfoLevel2", "a.b", InfoLevel, true},
		{"InfoLevel3", "a.b.c", InfoLevel, false},
		{"InfoLevel4", "a.b.c.d", InfoLevel, false},

		{"WarnLevel1", "a", WarnLevel, true},
		{"WarnLevel2", "a.b", WarnLevel, false},
		{"WarnLevel3", "a.b.c", WarnLevel, false},
		{"WarnLevel4", "a.b.c.d", WarnLevel, false},

		{"ErrorLevel1", "a", ErrorLevel, false},
		{"ErrorLevel2", "a.b", ErrorLevel, false},
		{"ErrorLevel3", "a.b.c", ErrorLevel, false},
		{"ErrorLevel4", "a.b.c.d", ErrorLevel, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentId = tt.id
			if got := ignored(tt.action); got != tt.want {
				t.Errorf("ignored() = %v, want %v", got, tt.want)
			}
			currentId = ""
		})
	}
}
