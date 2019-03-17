package doodle

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

const apiBase = "https://doodle.com/api/v2.0/polls/"
const typeText = "TEXT"
const prefTypeYN = "YESNO"
const prefTypeYNI = "YESNOIFNEEDBE"

type Doodle struct {
	doodleResponse
}

type doodleResponse struct {
	Options []struct {
		Text string
	}

	Participants []struct {
		Name        string
		Preferences []int
	}

	Type            string
	PreferencesType string
}

type Option string
type Participant string

func ParseDoodle(url string) (*Doodle, error) {
	id := regexp.MustCompile(`\w{16,}`).FindString(url)
	resp, err := http.Get(apiBase + id)
	if err != nil {
		return nil, err
	}

	doodle := &Doodle{}
	err = json.NewDecoder(resp.Body).Decode(&(doodle.doodleResponse))
	if err != nil {
		return nil, err
	}

	if doodle.Type != typeText {
		return nil, errors.New("only polls of type '" + typeText + "' are supported")
	}

	return doodle, nil
}

func (d *Doodle) Options() []Option {
	var opts []Option

	for _, opt := range d.doodleResponse.Options {
		opts = append(opts, Option(opt.Text))
	}

	return opts
}

func (d *Doodle) Participants() []Participant {
	var parts []Participant

	for _, part := range d.doodleResponse.Participants {
		parts = append(parts, Participant(part.Name))
	}

	return parts
}

func (d *Doodle) Results() map[Participant]map[Option]int {
	results := map[Participant]map[Option]int{}

	for _, part := range d.doodleResponse.Participants {
		results[Participant(part.Name)] = map[Option]int{}
		for i, opt := range d.doodleResponse.Options {
			var value int

			switch d.PreferencesType {
			case prefTypeYN:
				value = part.Preferences[i] * 2
			case prefTypeYNI:
				value = part.Preferences[i] * 2
			}

			results[Participant(part.Name)][Option(opt.Text)] = value
		}
	}

	return results
}
