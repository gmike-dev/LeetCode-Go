package n1912_design_movie_rental_system

import (
	"LeetCode-Go/tst"
	"strings"
	"testing"
)

func TestMovieRentingSystem(t *testing.T) {
	s := Constructor(3, tst.Str2Matrix("[[0,1,5],[0,2,6],[0,3,7],[1,1,4],[1,2,7],[2,1,5]]"))
	tst.ListEqual(t, "[1,0,2]", s.Search(1))
	s.Rent(0, 1)
	s.Rent(1, 2)
	tst.MatrixEqual(t, "[[0,1],[1,2]]", s.Report())
	s.Drop(1, 2)
	tst.ListEqual(t, "[0,1]", s.Search(2))
}

func TestMovieRentingSystem4(t *testing.T) {
	s := Constructor(69, tst.Str2Matrix(`[[16,4156,1511],[20,8501,8417],[34,7901,7776],[54,6691,9511],[44,8931,8434],[42,9640,5251],[22,4534,9161],[32,6506,6831],[13,8501,731],[4,7610,8474],[33,820,2341],[17,6490,1161],[29,7120,2703],[8,8723,7613],[38,9544,1804],[30,8723,1047],[1,5015,7763],[60,1625,2383],[29,3336,3542],[39,7535,6066],[1,9074,9400],[39,1625,7944],[26,9160,6874],[55,2465,888],[35,8530,6025]]`))

	commands := `"rent","search","search","report","rent","rent","report","report","search","search","rent","rent",
"search","drop","drop","drop","drop","rent","report","report","rent","drop","search","report","drop",
"report","drop","rent","report","search","search","rent","rent","report","report","drop","report","report",
"drop","report","drop","rent","drop","search","rent","search","drop","rent","drop","report","rent","drop",
"rent","rent","drop","report","report","report","report","rent","drop","report","drop","rent","search","drop",
"report","rent","search","search","report","rent","report","report","rent","report","report","search","rent",
"rent","search"`
	data := tst.Str2Matrix(`[[32,6506],[8501],[6275],[],[30,8723],[8,8723],[],[],[6699],[115],[20,8501],[16,4156],[9447],[30,8723],[8,8723],[32,6506],[16,4156],[42,9640],[],[],[17,6490],[20,8501],[8175],[],[17,6490],[],[42,9640],[54,6691],[],[1625],[3291],[60,1625],[39,1625],[],[],[60,1625],[],[],[39,1625],[],[54,6691],[8,8723],[8,8723],[2260],[29,7120],[746],[29,7120],[38,9544],[38,9544],[],[1,9074],[1,9074],[54,6691],[39,1625],[54,6691],[],[],[],[],[26,9160],[26,9160],[],[39,1625],[42,9640],[9640],[42,9640],[],[29,7120],[5630],[1842],[],[16,4156],[],[],[1,9074],[],[],[7992],[4,7610],[29,3336],[1333]]`)
	for i, c := range strings.Split(commands, ",") {
		args := data[i]
		switch c {
		case `"rent"`:
			s.Rent(args[0], args[1])
			break
		case `"search"`:
			s.Search(args[0])
			break
		case `"report"`:
			s.Report()
			break
		case `"drop"`:
			s.Drop(args[0], args[1])
			break
		}
	}

}
