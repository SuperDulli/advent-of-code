package main

import (
	"testing"

	"github.com/SuperDulli/advent-of-code/util"
)

const example = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

const invalidPassports = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
`

const validPassports = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"example", util.SplitLines(example), 2},
		{"actual", util.ReadLines("input.txt"), 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"invalid examples", util.SplitLines(invalidPassports), 0},
		{"valid examples", util.SplitLines(validPassports), 4},
		{"actual", util.ReadLines("input.txt"), 131},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidYear(t *testing.T)  {
	tests := []struct {
		name string
		input string
		want  bool
	}{
		{"valid birth year", "2002", true},
		{"invalid birth year", "2003", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidYear(tt.input, 1920, 2002); got != tt.want {
				t.Errorf("isValidBirthYear(%v, 1920, 2002) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func Test_isValidHeight(t *testing.T)  {
	tests := []struct {
		name string
		input string
		want  bool
	}{
		{"valid height inch", "60in", true},
		{"valid height cm", "190cm", true},
		{"invalid height inch", "190in", false},
		{"invalid height", "190", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidHeight(tt.input); got != tt.want {
				t.Errorf("isValidHeight(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func Test_isValidHairColor(t *testing.T)  {
	tests := []struct {
		name string
		input string
		want  bool
	}{
		{"valid hair color", "#123abc", true},
		{"invalid hair color", "#123abz", false},
		{"invalid hair color", "123abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidHairColor(tt.input); got != tt.want {
				t.Errorf("isValidHairColor(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func Test_isValidEyeColor(t *testing.T)  {
	tests := []struct {
		name string
		input string
		want  bool
	}{
		{"valid eye color", "brn", true},
		{"invalid eye color", "wat", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEyeColor(tt.input); got != tt.want {
				t.Errorf("isValidEyeColor(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func Test_isValidPassportID(t *testing.T)  {
	tests := []struct {
		name string
		input string
		want  bool
	}{
		{"valid passport id", "000000001", true},
		{"invalid passport id", "0123456789", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPassportID(tt.input); got != tt.want {
				t.Errorf("isValidPassportID(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
