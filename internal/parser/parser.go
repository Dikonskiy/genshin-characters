package parser

import (
	"encoding/csv"
	"fmt"
	"genshin/internal/models"
	"io"
	"log"
	"os"
	"time"
)

type Parse struct {
}

func NewParser() *Parse {
	return &Parse{}
}

func (p *Parse) FetchData() ([]models.Character, error) {
	f, err := os.Open("data/genshin.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var characters []models.Character
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		character, err := p.parseCharacterData(rec)
		if err != nil {
			log.Printf("Error parsing line: %v", err)
			continue
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func (p *Parse) parseCharacterData(data []string) (models.Character, error) {
	if len(data) != 39 {
		return models.Character{}, fmt.Errorf("invalid data format: expected 39 fields, got %d", len(data))
	}

	releaseDate, err := time.Parse("2006-01-02", data[6])
	if err != nil {
		return models.Character{}, err
	}

	return models.Character{
		CharacterName:       data[0],
		Rarity:              data[1],
		Region:              data[2],
		Vision:              data[3],
		Arkhe:               data[4],
		WeaponType:          data[5],
		ReleaseDate:         releaseDate,
		Model:               data[7],
		Constellation:       data[8],
		Birthday:            data[9],
		SpecialDish:         data[10],
		Affiliation:         data[11],
		Limited:             data[12],
		VoiceEng:            data[13],
		VoiceCN:             data[14],
		VoiceJP:             data[15],
		VoiceKR:             data[16],
		Ascension:           data[17],
		AscensionSpecialty:  data[18],
		AscensionBoss:       data[19],
		AscensionMaterial02: data[20],
		AscensionMaterial24: data[21],
		AscensionMaterial46: data[22],
		AscensionGem01:      data[23],
		AscensionGem13:      data[24],
		AscensionGem35:      data[25],
		AscensionGem56:      data[26],
		TalentMaterial:      data[27],
		TalentBook12:        data[28],
		TalentBook23:        data[29],
		TalentBook34:        data[30],
		TalentBook45:        data[31],
		TalentBook56:        data[32],
		TalentBook67:        data[33],
		TalentBook78:        data[34],
		TalentBook89:        data[35],
		TalentBook910:       data[36],
		TalentWeekly:        data[37],
	}, nil
}
