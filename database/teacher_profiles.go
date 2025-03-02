package database

import (
	"context"
	"fmt"
	"time"
)

// TeacherProfile represents the TeacherProfiles table
type TeacherProfile struct {
	TeacherID       int       `json:"teacher_id"`
	UserID          int       `json:"user_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Qualifications  string    `json:"qualifications"`
	PastExperiences string    `json:"past_experiences"`
	Achievements    string    `json:"achievements"`
	Interests       string    `json:"interests"`
	Specialization  string    `json:"specialization"`
	ExperienceYears int       `json:"experience_years"`
	Skills          string    `json:"skills"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ====================== TeacherProfiles Table CRUD Operations ======================

// CreateTeacherProfile inserts a new teacher profile into the database
func CreateTeacherProfile(profile TeacherProfile) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, first_name, last_name, qualifications, past_experiences, achievements, interests, specialization, experience_years, skills)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING teacher_id, created_at, updated_at`, TableTeacherProfiles)
	return db.QueryRow(context.Background(), query, profile.UserID, profile.FirstName, profile.LastName, profile.Qualifications, profile.PastExperiences, profile.Achievements, profile.Interests, profile.Specialization, profile.ExperienceYears, profile.Skills).
		Scan(&profile.TeacherID, &profile.CreatedAt, &profile.UpdatedAt)
}

// GetTeacherProfile retrieves a teacher profile by ID
func GetTeacherProfile(teacherID int) (TeacherProfile, error) {
	var profile TeacherProfile
	query := fmt.Sprintf(`
		SELECT teacher_id, user_id, first_name, last_name, qualifications, past_experiences, achievements, interests, specialization, experience_years, skills, created_at, updated_at
		FROM %s
		WHERE teacher_id = $1`, TableTeacherProfiles)
	err := db.QueryRow(context.Background(), query, teacherID).
		Scan(&profile.TeacherID, &profile.UserID, &profile.FirstName, &profile.LastName, &profile.Qualifications, &profile.PastExperiences, &profile.Achievements, &profile.Interests, &profile.Specialization, &profile.ExperienceYears, &profile.Skills, &profile.CreatedAt, &profile.UpdatedAt)
	return profile, err
}

// UpdateTeacherProfile updates an existing teacher profile
func UpdateTeacherProfile(profile TeacherProfile) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET first_name = $1, last_name = $2, qualifications = $3, past_experiences = $4, achievements = $5, interests = $6, specialization = $7, experience_years = $8, skills = $9, updated_at = CURRENT_TIMESTAMP
		WHERE teacher_id = $10`, TableTeacherProfiles)
	_, err := db.Exec(context.Background(), query, profile.FirstName, profile.LastName, profile.Qualifications, profile.PastExperiences, profile.Achievements, profile.Interests, profile.Specialization, profile.ExperienceYears, profile.Skills, profile.TeacherID)
	return err
}

// DeleteTeacherProfile deletes a teacher profile by ID
func DeleteTeacherProfile(teacherID int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE teacher_id = $1`, TableTeacherProfiles)
	_, err := db.Exec(context.Background(), query, teacherID)
	return err
}
