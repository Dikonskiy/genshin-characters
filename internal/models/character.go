package models

import "time"

type Character struct {
	Id                  int       `db:"id"`
	CharacterName       string    `db:"CharacterName"`
	Rarity              string    `db:"Rarity"`
	Region              string    `db:"Region"`
	Vision              string    `db:"Vision"`
	Arkhe               string    `db:"Arkhe"`
	WeaponType          string    `db:"WeaponType"`
	ReleaseDate         time.Time `db:"ReleaseDate"`
	Model               string    `db:"Model"`
	Constellation       string    `db:"Constellation"`
	Birthday            string    `db:"Birthday"`
	SpecialDish         string    `db:"SpecialDish"`
	Affiliation         string    `db:"Affiliation"`
	Limited             string    `db:"Limited"`
	VoiceEng            string    `db:"VoiceEng"`
	VoiceCN             string    `db:"VoiceCN"`
	VoiceJP             string    `db:"VoiceJP"`
	VoiceKR             string    `db:"VoiceKR"`
	Ascension           string    `db:"Ascension"`
	AscensionSpecialty  string    `db:"AscensionSpecialty"`
	AscensionBoss       string    `db:"AscensionBoss"`
	AscensionMaterial02 string    `db:"AscensionMaterial02"`
	AscensionMaterial24 string    `db:"AscensionMaterial24"`
	AscensionMaterial46 string    `db:"AscensionMaterial46"`
	AscensionGem01      string    `db:"AscensionGem01"`
	AscensionGem13      string    `db:"AscensionGem13"`
	AscensionGem35      string    `db:"AscensionGem35"`
	AscensionGem56      string    `db:"AscensionGem56"`
	TalentMaterial      string    `db:"TalentMaterial"`
	TalentBook12        string    `db:"TalentBook12"`
	TalentBook23        string    `db:"TalentBook23"`
	TalentBook34        string    `db:"TalentBook34"`
	TalentBook45        string    `db:"TalentBook45"`
	TalentBook56        string    `db:"TalentBook56"`
	TalentBook67        string    `db:"TalentBook67"`
	TalentBook78        string    `db:"TalentBook78"`
	TalentBook89        string    `db:"TalentBook89"`
	TalentBook910       string    `db:"TalentBook910"`
	TalentWeekly        string    `db:"TalentWeekly"`
}
