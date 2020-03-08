package main

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_findBiggerSquare(t *testing.T) {
	veryBig, _ := (&big.Int{}).SetString("313230894596513941163065516500542159481861849753982064716706926040955753912601", 10)
	veryBigRes, _ := (&big.Int{}).SetString("559670344574834464368752451777605378546", 10)
	veryVeryBig, _ := (&big.Int{}).SetString("693532252693819955253847249043757529180324494073823494547591843582136805148466227647256127350410778190683007539432807696698056465273137755794986981751541604145575912570331718107924355720576988204493545005230906033105693432028755478062000534784352638264990268794303011750110045088373921373348481833734839258579645479000424548585936336555233860123720335766312774151565346595485784249448792904040188292802408882242972670662804668689397345518513975596694683183631", 10)
	veryVeryBigRes, _ := (&big.Int{}).SetString("26335000525798740816205266852795805541326880463277554675470589224513328126393817529421496624598433838006421267113419955818686164886768506921682314635059747745407488630302947521135203753155306998611178523254880676592688270926934049", 10)
	tests := []struct {
		name string
		num  *big.Int
		want *big.Int
	}{
		{
			name: "one",
			num:  big.NewInt(1),
		},
		{
			name: "5",
			num:  big.NewInt(5),
			want: big.NewInt(3),
		},
		{
			name: "simple",
			num:  big.NewInt(1488),
			want: big.NewInt(39),
		},
		{
			name: "already square",
			num:  big.NewInt(169),
			want: big.NewInt(14),
		},
		{
			name: "big 2^",
			num:  big.NewInt(1 << 60),
			want: big.NewInt(1073741825),
		},
		{
			name: "very big",
			num:  veryBig,
			want: veryBigRes,
		},
		{
			name: "very very big",
			num:  veryVeryBig,
			want: veryVeryBigRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBiggerSquare(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBiggerSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleSearch(t *testing.T) {
	tests := []struct {
		name  string
		num   *big.Int
		want1 *big.Int
		want2 *big.Int
	}{
		{
			name:  "six",
			num:   big.NewInt(6),
			want1: big.NewInt(3),
			want2: big.NewInt(2),
		},
		{
			name:  "one",
			num:   big.NewInt(1),
			want1: big.NewInt(1),
			want2: big.NewInt(1),
		},
		{
			name:  "five",
			num:   big.NewInt(5),
			want1: big.NewInt(5),
			want2: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := simpleSearch(tt.num); !(reflect.DeepEqual(got1, tt.want1) && reflect.DeepEqual(got2, tt.want2)) {
				t.Errorf("\ngot1: %v, want1: %v\ngot2: %v, want2: %v\n", got1, tt.want1, got2, tt.want2)
			}
		})
	}
}
