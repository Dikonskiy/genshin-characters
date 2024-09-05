package repository

import (
	"fmt"
	"genshin/internal/models"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	InsertCharacter(character models.Character) error
	GetCharacters(id string) (characters []models.Character, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repo {
	return &repository{
		db: db,
	}
}

const query = `INSERT INTO Characters`

func (r *repository) InsertCharacter(character models.Character) error {
	insertQuery, args := r.getInsertQuery(character)

	_, err := r.db.Exec(insertQuery, args...)

	if err != nil {
		return fmt.Errorf("error inserting character: %v", err)
	}

	return nil
}

func (r *repository) getInsertQuery(c models.Character) (string, []interface{}) {
	var (
		sql             strings.Builder
		columns, values []string
		args            []interface{}
	)

	sql.WriteString(query)

	columns = append(columns, "CharacterName", "Rarity", "Region", "Vision", "Arkhe", "WeaponType", "ReleaseDate", "Model", "Constellation",
		"Birthday", "SpecialDish", "Affiliation", "Limited", "VoiceEng", "VoiceCN", "VoiceJP", "VoiceKR", "Ascension", "AscensionSpecialty",
		"AscensionBoss", "AscensionMaterial02", "AscensionMaterial24", "AscensionMaterial46", "AscensionGem01", "AscensionGem13", "AscensionGem35",
		"AscensionGem56", "TalentMaterial", "TalentBook12", "TalentBook23", "TalentBook34", "TalentBook45", "TalentBook56", "TalentBook67", "TalentBook78",
		"TalentBook89", "TalentBook910", "TalentWeekly")

	args = append(args, c.CharacterName, c.Rarity, c.Region, c.Vision, c.Arkhe, c.WeaponType, c.ReleaseDate, c.Model, c.Constellation,
		c.Birthday, c.SpecialDish, c.Affiliation, c.Limited, c.VoiceEng, c.VoiceCN, c.VoiceJP, c.VoiceKR, c.Ascension, c.AscensionSpecialty,
		c.AscensionBoss, c.AscensionMaterial02, c.AscensionMaterial24, c.AscensionMaterial46, c.AscensionGem01, c.AscensionGem13, c.AscensionGem35,
		c.AscensionGem56, c.TalentMaterial, c.TalentBook12, c.TalentBook23, c.TalentBook34, c.TalentBook45, c.TalentBook56, c.TalentBook67, c.TalentBook78,
		c.TalentBook89, c.TalentBook910, c.TalentWeekly)

	for i := 0; i < len(columns); i++ {
		values = append(values, "?")
	}
	sql.WriteString(fmt.Sprintf("%s%s%s", "(", strings.Join(columns, ", "), ")"))
	sql.WriteString(" VALUES ")
	sql.WriteString(fmt.Sprintf("%s%s%s", "(", strings.Join(values, ", "), ")"))

	return sql.String(), args
}

const getQuery = `SELECT id, CharacterName, Rarity, Region, Vision, Arkhe, WeaponType, ReleaseDate, Model, Constellation,
		Birthday, SpecialDish, Affiliation, Limited, VoiceEng, VoiceCN, VoiceJP, VoiceKR, Ascension, AscensionSpecialty,
		AscensionBoss, AscensionMaterial02, AscensionMaterial24, AscensionMaterial46, AscensionGem01, AscensionGem13, AscensionGem35,
		AscensionGem56, TalentMaterial, TalentBook12, TalentBook23, TalentBook34, TalentBook45, TalentBook56, TalentBook67, TalentBook78,
		TalentBook89, TalentBook910, TalentWeekly from Characters`

func (r *repository) GetCharacters(id string) (characters []models.Character, err error) {
	query, args := r.getCharacters(id)

	if err = r.db.Select(&characters, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return characters, nil
}

func (r *repository) getCharacters(id string) (string, []interface{}) {
	var (
		sql     strings.Builder
		clauses []string
		args    []interface{}
	)

	sql.WriteString(getQuery)

	if id != "" {
		clauses = append(clauses, "id = ?")
		args = append(args, id)
	}

	if len(clauses) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(clauses, ` AND `))
	}

	return sql.String(), args
}
