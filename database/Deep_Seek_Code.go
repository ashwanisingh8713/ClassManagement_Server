package database

// ====================== Main Function ======================

/*func main() {
	ConnectDB()

	// Example: Run CRUD operations in the background
	go func() {
		// Create a user
		user := User{
			Username:     "john_doe",
			PasswordHash: "hashed_password",
			Role:         "Student",
			Email:        "john.doe@example.com",
		}
		if err := CreateUser(user); err != nil {
			log.Printf("Error creating user: %v", err)
		} else {
			log.Println("User created successfully!")
		}

		// Get the user
		retrievedUser, err := GetUser(user.UserID)
		if err != nil {
			log.Printf("Error retrieving user: %v", err)
		} else {
			log.Printf("Retrieved user: %+v\n", retrievedUser)
		}
	}()

	// Keep the main function running to allow background goroutines to complete
	time.Sleep(2 * time.Second)
}*/
