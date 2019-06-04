package framerr_test

import (
	"testing"

	"github.com/jbowes/framerr"
	"golang.org/x/xerrors"
)

type dummyErr struct {
	msg string
}

func (d *dummyErr) Error() string {
	return d.msg
}

func TestExtract(t *testing.T) {
	tcs := []struct {
		name      string
		in        error
		out       []framerr.Frame
		hasFrames bool
	}{
		{"nil err", nil, []framerr.Frame{}, false},
		{"single error, no frame", &dummyErr{"test"}, []framerr.Frame{
			{Text: "test"},
		}, false},
		{"single error, with frame", xerrors.New("test"), []framerr.Frame{
			{Text: "test"},
		}, true},
		{"multi error", xerrors.Errorf("other: %w", xerrors.New("test")), []framerr.Frame{
			{Text: "other"},
			{Text: "test"},
		}, true},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			out := framerr.Extract(tc.in)
			if len(tc.out) != len(out) {
				t.Fatal("bad frame. wanted:", tc.out, "got:", out)
			}

			for i, o := range out {
				if tc.out[i].Text != o.Text {
					t.Fatal("bad frame. wanted:", tc.out, "got:", out)
				}

				if tc.hasFrames {
					if o.Source == nil {
						t.Fatalf("expected stack frames")
					}

					if o.Source.Func != "TestExtract" || o.Source.Package != "github.com/jbowes/framerr_test" {
						t.Fatal("bad stack frame detail. saw:", o.Source)
					}
				}
			}
		})
	}
}
