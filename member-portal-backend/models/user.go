package models

type Users struct {
	UserID   string `gorm:"primaryKey;column:user_id;type:varchar(20)"`
	Nickname string `gorm:"column:nickname;type:varchar(20)"`
}

type Auth struct {
	UserID         string `gorm:"primaryKey;column:user_id;type:varchar(20)"`
	HashedPassword string `gorm:"column:hashed_password;type:varchar(100)"`
	RoleID         int    `gorm:"column:role_id"`
	Role           Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	RoleID   int    `gorm:"primaryKey;column:role_id"`
	RoleName string `gorm:"column:role_name;type:varchar(20)"`
}

type Affiliation struct {
	UserID    string  `gorm:"primaryKey;column:user_id;type:varchar(20)"`
	FacultyID int     `gorm:"column:faculty_id"`
	Grade     int     `gorm:"column:grade"`
	Faculty   Faculty `gorm:"foreignKey:FacultyID"`
}

type Faculty struct {
	FacultyID      int    `gorm:"primaryKey;column:faculty_id"`
	FacultyName    string `gorm:"column:faculty_name;type:varchar(20)"`
	DepartmentName string `gorm:"column:department_name;type:varchar(20)"`
}

type Contact struct {
	UserID      string `gorm:"primaryKey;column:user_id;type:varchar(20)"`
	SchoolEmail string `gorm:"column:school_email"`
	SubEmail    string `gorm:"column:sub_email"`
	DiscordID   string `gorm:"column:discord_id"`
	GithubID    string `gorm:"column:github_id"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20)"`
}

type Name struct {
	UserID     string `gorm:"primaryKey;column:user_id;type:varchar(20)"`
	FirstName  string `gorm:"column:first_name;type:varchar(20)"`
	LastName   string `gorm:"column:last_name;type:varchar(20)"`
	MiddleName string `gorm:"column:middle_name;type:varchar(40)"`
}

type Profile struct {
	UserID       string `gorm:"primaryKey;column:user_id;type:varchar(20)"`
	ProfileImage string `gorm:"column:profile_image"`
	Bio          string `gorm:"column:bio;size:200"`
}

type ActivityLog struct {
	ActivityID          int    `gorm:"primaryKey;column:activity_id"`
	UserID              string `gorm:"column:user_id;type:varchar(20)"`
	ActivityDate        string `gorm:"column:activity_date"`
	ActivityDescription string `gorm:"column:activity_description"`
}
