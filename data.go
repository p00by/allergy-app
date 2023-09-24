package main

import "time"

var data Diagnosis

type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

type Diagnosis struct {
	Allergies    string
	Observations []Observation
}

type Observation struct {
	Date     time.Time
	Allergic bool
}

func init() {
	data = Diagnosis{
		Allergies:    "Not known, please provide more data",
		Observations: nil,
	}
}

func updateAllergies(newDiagnosedDate string, allergic bool) {
	var time, timeParseError = time.Parse("2006-01-02", newDiagnosedDate)
	if timeParseError == nil {
		var newObservations = append(data.Observations, Observation{
			Date:     time,
			Allergic: allergic,
		})
		data = Diagnosis{
			Allergies:    getPotentialAllergies(newObservations),
			Observations: newObservations,
		}
	} else {
		panic(timeParseError)
	}
}

func getPotentialAllergies(newObservations []Observation) string {
	return "I don't know"
}
