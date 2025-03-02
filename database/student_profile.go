package database

import (
	"context"
	"fmt"
	"time"
)

// StudentProfile represents the StudentProfiles table
type StudentProfile struct {
	StudentID        int       `json:"student_id"`
	UserID           int       `json:"user_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Achievements     string    `json:"achievements"`
	ParentName       string    `json:"parent_name"`
	ParentContact    string    `json:"parent_contact"`
	ParentEmail      string    `json:"parent_email"`
	ParentAddress    string    `json:"parent_address"`
	ParentOccupation string    `json:"parent_occupation"`
	Grade            string    `json:"grade"`
	Section          string    `json:"section"`
	RollNo           int       `json:"roll_no"`
	DOB              time.Time `json:"dob"`
	StudentAddress   string    `json:"student_address"`
	StudentContact   string    `json:"student_contact"`
	StudentEmail     string    `json:"student_email"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ====================== StudentProfiles Table CRUD Operations ======================

// CreateStudentProfile inserts a new student profile into the database
func CreateStudentProfile(profile StudentProfile) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, first_name, last_name, achievements, parent_name, parent_contact, parent_email, parent_address, parent_occupation, grade, section, roll_no, dob, student_address, student_contact, student_email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING student_id, created_at, updated_at`, TableStudentProfiles)
	return db.QueryRow(context.Background(), query, profile.UserID, profile.FirstName, profile.LastName, profile.Achievements, profile.ParentName, profile.ParentContact, profile.ParentEmail, profile.ParentAddress, profile.ParentOccupation, profile.Grade, profile.Section, profile.RollNo, profile.DOB, profile.StudentAddress, profile.StudentContact, profile.StudentEmail).
		Scan(&profile.StudentID, &profile.CreatedAt, &profile.UpdatedAt)
}

// GetStudentProfile retrieves a student profile by ID
func GetStudentProfile(studentID int) (StudentProfile, error) {
	var profile StudentProfile
	query := fmt.Sprintf(`
		SELECT student_id, user_id, first_name, last_name, achievements, parent_name, parent_contact, parent_email, parent_address, parent_occupation, grade, section, roll_no, dob, student_address, student_contact, student_email, created_at, updated_at
		FROM %s
		WHERE student_id = $1`, TableStudentProfiles)
	err := db.QueryRow(context.Background(), query, studentID).
		Scan(&profile.StudentID, &profile.UserID, &profile.FirstName, &profile.LastName, &profile.Achievements, &profile.ParentName, &profile.ParentContact, &profile.ParentEmail, &profile.ParentAddress, &profile.ParentOccupation, &profile.Grade, &profile.Section, &profile.RollNo, &profile.DOB, &profile.StudentAddress, &profile.StudentContact, &profile.StudentEmail, &profile.CreatedAt, &profile.UpdatedAt)
	return profile, err
}

// UpdateStudentProfile updates an existing student profile
func UpdateStudentProfile(profile StudentProfile) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET first_name = $1, last_name = $2, achievements = $3, parent_name = $4, parent_contact = $5, parent_email = $6, parent_address = $7, parent_occupation = $8, grade = $9, section = $10, roll_no = $11, dob = $12, student_address = $13, student_contact = $14, student_email = $15, updated_at = CURRENT_TIMESTAMP
		WHERE student_id = $16`, TableStudentProfiles)
	_, err := db.Exec(context.Background(), query, profile.FirstName, profile.LastName, profile.Achievements, profile.ParentName, profile.ParentContact, profile.ParentEmail, profile.ParentAddress, profile.ParentOccupation, profile.Grade, profile.Section, profile.RollNo, profile.DOB, profile.StudentAddress, profile.StudentContact, profile.StudentEmail, profile.StudentID)
	return err
}

// DeleteStudentProfile deletes a student profile by ID
func DeleteStudentProfile(studentID int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE student_id = $1`, TableStudentProfiles)
	_, err := db.Exec(context.Background(), query, studentID)
	return err
}
