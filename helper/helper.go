package helper

import (
	"time"
)

func ConvertiData(data string) time.Time {
	// Stringa di esempio
	t, err := time.Parse("2006-01-02", data)
	if err != nil {
		panic(err)
	}

	// Converti in orario italiano (GMT+1)
	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		panic(err)
	}
	t = t.In(loc)

	return t
}

// DateRange Ritorna uno slice di time.Time partendo dalla data di inizio e sommando 24h fino al parametro fine
func DateRange(inizio, fine time.Time) []time.Time {
	var dates []time.Time
	for d := inizio; d.Before(fine) || d.Equal(fine); d = d.Add(time.Hour * 24) {
		dates = append(dates, d)
	}
	return dates
}

// SlicePosti Ritorna uno slice di int partendo dal parametro start per X volte pari al parametro count
func SlicePosti(start, count int) []int {
	result := make([]int, count)
	for i := range result {
		result[i] = start + i
	}
	return result
}

// Non più utilizzata
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
