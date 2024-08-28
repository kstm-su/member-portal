package models

type User struct {
	UserID       string `gorm:"primaryKey;column:user_id"`
	Nickname     string `gorm:"column:nickname"`
	Auth         Auth
	Affiliation  Affiliation
	Contact      Contact
	Name         Name
	Profile      Profile
	ActivityLogs []ActivityLog
}

type Auth struct {
	UserID         string `gorm:"primaryKey;column:user_id"`
	HashedPassword string `gorm:"column:hashed_password"`
	RoleID         int    `gorm:"column:role_id"`
	Role           Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	RoleID   int    `gorm:"primaryKey;column:role_id"`
	RoleName string `gorm:"column:role_name"`
}

type Affiliation struct {
	UserID    string  `gorm:"primaryKey;column:user_id"`
	FacultyID int     `gorm:"column:faculty_id"`
	Grade     int     `gorm:"column:grade"`
	Faculty   Faculty `gorm:"foreignKey:FacultyID"`
}

type Faculty struct {
	FacultyID      int    `gorm:"primaryKey;column:faculty_id"`
	FacultyName    string `gorm:"column:faculty_name"`
	DepartmentName string `gorm:"column:department_name"`
}

type Contact struct {
	UserID      string `gorm:"primaryKey;column:user_id"`
	SchoolEmail string `gorm:"column:school_email"`
	SubEmail    string `gorm:"column:sub_email"`
	DiscordID   string `gorm:"column:discord_id"`
	GithubID    string `gorm:"column:github_id"`
	PhoneNumber string `gorm:"column:phone_number"`
}

type Name struct {
	UserID     string `gorm:"primaryKey;column:user_id"`
	FirstName  string `gorm:"column:first_name"`
	LastName   string `gorm:"column:last_name"`
	MiddleName string `gorm:"column:middle_name"`
}

type Profile struct {
	UserID       string `gorm:"primaryKey;column:user_id"`
	ProfileImage string `gorm:"column:profile_image"`
	Bio          string `gorm:"column:bio"`
}

type ActivityLog struct {
	ActivityID          int    `gorm:"primaryKey;column:activity_id"`
	UserID              string `gorm:"column:user_id"`
	ActivityDate        string `gorm:"column:activity_date"`
	ActivityDescription string `gorm:"column:activity_description"`
}
