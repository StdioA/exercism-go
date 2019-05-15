package scale

var pitchMap = map[string]int{
	"C": 0,
	"D": 2,
	"E": 4,
	"F": 5,
	"G": 7,
	"A": 9,
	"B": 11,
}
var stringMap = map[int][]string{
	0:  []string{"C"},
	1:  []string{"C#", "Db"},
	2:  []string{"D"},
	3:  []string{"D#", "Eb"},
	4:  []string{"E"},
	5:  []string{"F"},
	6:  []string{"F#", "Gb"},
	7:  []string{"G"},
	8:  []string{"G#", "Ab"},
	9:  []string{"A"},
	10: []string{"A#", "Bb"},
	11: []string{"B"},
}
var notationMap = map[int]int{
	0:  0, // C
	7:  0, // G
	2:  0, // D
	9:  0, // A
	4:  0, // E
	11: 0, // B
	6:  0, // F#
	5:  1, // F
	10: 1, // Bb
	3:  1, // Eb
	8:  1, // Ab
	1:  1, // Db
}
