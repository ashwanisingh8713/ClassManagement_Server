package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

// Constants for table names
const (
	TableUsers                          = "users"
	TableTeacherProfiles                = "teacher_profiles"
	TableAdminProfiles                  = "admin_profiles"
	TableStudentProfiles                = "student_profiles"
	TableClasses                        = "classes"
	TableSections                       = "sections"
	TableSubjects                       = "subjects"
	TableClassSubjects                  = "class_subjects"
	TableTuitionPeriods                 = "tuition_periods"
	TableTeacherClassSubjectAssignments = "teacher_class_subject_assignments"
)

// Global database connection pool
var db *pgxpool.Pool

func GetDB() *pgxpool.Pool {
	return db
}

// ConnectDB initializes the database connection
func ConnectDB() {
	dsn := fmt.Sprintf("postgres://ashwani:@localhost:5432/ashwani")
	var err error
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	fmt.Println("Connected to PostgreSQL!")
}

// CreateTables creates all tables programmatically
func CreateTables() {
	queries := []string{
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				user_id SERIAL PRIMARY KEY,
				username VARCHAR(50) UNIQUE NOT NULL,
				password_hash VARCHAR(255) NOT NULL,
				role VARCHAR(20) NOT NULL CHECK (role IN ('Teacher', 'Admin', 'Student')),
				email VARCHAR(100) UNIQUE NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableUsers),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				teacher_id SERIAL PRIMARY KEY,
				user_id INT UNIQUE NOT NULL REFERENCES %s(user_id) ON DELETE CASCADE,
				first_name VARCHAR(50) NOT NULL,
				last_name VARCHAR(50) NOT NULL,
				qualifications TEXT,
				past_experiences TEXT,
				achievements TEXT,
				interests TEXT,
				specialization VARCHAR(100),
				experience_years INT,
				skills TEXT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableTeacherProfiles, TableUsers),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				admin_id SERIAL PRIMARY KEY,
				user_id INT UNIQUE NOT NULL REFERENCES %s(user_id) ON DELETE CASCADE,
				first_name VARCHAR(50) NOT NULL,
				last_name VARCHAR(50) NOT NULL,
				qualifications TEXT,
				past_experiences TEXT,
				achievements TEXT,
				interests TEXT,
				specialization VARCHAR(100),
				experience_years INT,
				skills TEXT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableAdminProfiles, TableUsers),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				student_id SERIAL PRIMARY KEY,
				user_id INT UNIQUE NOT NULL REFERENCES %s(user_id) ON DELETE CASCADE,
				first_name VARCHAR(50) NOT NULL,
				last_name VARCHAR(50) NOT NULL,
				achievements TEXT,
				parent_name VARCHAR(100),
				parent_contact VARCHAR(20),
				parent_email VARCHAR(100),
				parent_address TEXT,
				parent_occupation VARCHAR(100),
				grade VARCHAR(20),
				section VARCHAR(10),
				roll_no INT,
				dob DATE,
				student_address TEXT,
				student_contact VARCHAR(20),
				student_email VARCHAR(100),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableStudentProfiles, TableUsers),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				class_id SERIAL PRIMARY KEY,
				class_name VARCHAR(50) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableClasses),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				section_id SERIAL PRIMARY KEY,
				class_id INT NOT NULL REFERENCES %s(class_id) ON DELETE CASCADE,
				section_name VARCHAR(10) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				UNIQUE (class_id, section_name)
			);`, TableSections, TableClasses),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				subject_id SERIAL PRIMARY KEY,
				subject_name VARCHAR(100) NOT NULL,
				writer VARCHAR(100),
				description TEXT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableSubjects),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				class_subject_id SERIAL PRIMARY KEY,
				class_id INT NOT NULL REFERENCES %s(class_id) ON DELETE CASCADE,
				subject_id INT NOT NULL REFERENCES %s(subject_id) ON DELETE CASCADE,
				difficulty_level VARCHAR(50),
				description TEXT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				UNIQUE (class_id, subject_id)
			);`, TableClassSubjects, TableClasses, TableSubjects),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				tuition_period_id SERIAL PRIMARY KEY,
				period_name VARCHAR(100) NOT NULL,
				start_date DATE NOT NULL,
				end_date DATE NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`, TableTuitionPeriods),
		fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				assignment_id SERIAL PRIMARY KEY,
				teacher_id INT NOT NULL REFERENCES %s(teacher_id) ON DELETE CASCADE,
				class_subject_id INT NOT NULL REFERENCES %s(class_subject_id) ON DELETE CASCADE,
				tuition_period_id INT NOT NULL REFERENCES %s(tuition_period_id) ON DELETE CASCADE,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				UNIQUE (teacher_id, class_subject_id, tuition_period_id)
			);`, TableTeacherClassSubjectAssignments, TableTeacherProfiles, TableClassSubjects, TableTuitionPeriods),
	}

	var wg sync.WaitGroup
	for _, query := range queries {
		wg.Add(1)
		go func(q string) {
			defer wg.Done()
			_, err := db.Exec(context.Background(), q)
			if err != nil {
				log.Printf("Error executing query: %v\nQuery: %s", err, q)
			}
		}(query)
	}
	wg.Wait()
	fmt.Println("All tables created successfully!")
}
