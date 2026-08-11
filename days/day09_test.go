package days

import "testing"

func Test_binarySearch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		segments []Segment
		val      func(s Segment) int
		target   int
		first    bool
		want     int
	}{
		{
			segments: []Segment{
				newSegment(1, 1, 1, 6),
				newSegment(2, 1, 2, 6),
				newSegment(2, 3, 2, 3),
				newSegment(3, 3, 3, 3),
				newSegment(4, 3, 4, 3),
			},
			val:    getStartX,
			target: 2,
			first:  true,
			want:   1,
		},
		{
			segments: []Segment{
				newSegment(2, 1, 1, 6),
				newSegment(2, 1, 2, 6),
				newSegment(2, 3, 2, 3),
				newSegment(3, 3, 3, 3),
				newSegment(4, 3, 4, 3),
			},
			val:    getStartX,
			target: 2,
			first:  true,
			want:   0,
		},
		{
			segments: []Segment{
				newSegment(1, 1, 1, 6),
				newSegment(2, 1, 2, 6),
				newSegment(2, 3, 2, 3),
				newSegment(4, 3, 4, 3),
				newSegment(5, 3, 5, 3),
			},
			val:    getStartX,
			target: 3,
			first:  true,
			want:   3,
		},
		{
			segments: []Segment{
				newSegment(1, 1, 1, 6),
				newSegment(2, 1, 2, 6),
				newSegment(2, 3, 2, 3),
				newSegment(4, 3, 4, 3),
				newSegment(5, 3, 5, 3),
			},
			val:    getStartX,
			target: 3,
			first:  false,
			want:   2,
		},
	}

	for ti, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.segments, tt.val, tt.target, tt.first)
			if got != tt.want {
				t.Errorf("test N°%v, binarySearch() = %v, want %v", ti, got, tt.want)
			}
		})
	}
}

func Test1Intersection(t *testing.T) {
	expectedRes := 3
	db := SegmentDB{}

	segments := []Segment{
		newSegment(50, 3, 70, 3),
		newSegment(0, 10, 40, 10),
		newSegment(60, 15, 100, 15),
		newSegment(30, 20, 50, 20),
		newSegment(25, 21, 50, 21),
	}

	for _, s := range segments {
		db.Insert(s)
	}

	res, err := db.CountIntersections(newSegment(35, 0, 35, 100))
	if err != nil {
		t.Errorf("%v", err)
	}

	if res != expectedRes {
		t.Errorf("res=%v, should be %v", res, expectedRes)
	}

}

func TestSegmentDB_CountIntersections(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       Segment
		segs    []Segment
		want    int
		wantErr bool
	}{
		{
			s: newSegment(7, 2, 7, 6),
			segs: []Segment{
				newSegment(7, 1, 11, 1),
				newSegment(11, 7, 9, 7),
				newSegment(9, 5, 2, 5),
				newSegment(2, 3, 7, 3),
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var db SegmentDB
			for _, s := range tt.segs {
				db.Insert(s)
			}

			got, gotErr := db.CountIntersections(tt.s)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CountIntersections() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CountIntersections() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("CountIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}
