package postgres

import (
	"time"

	"github.com/google/uuid"
)

// UserModel represents the users table
type UserModel struct {
	ID              int64      `gorm:"primaryKey;autoIncrement"`
	Role            string     //`gorm:"type:varchar(32);not null;index:idx_users_role"`
	Email           string     //`gorm:"type:varchar(64);uniqueIndex:uniq_users_email;not null"`
	Password        string     //`gorm:"type:varchar(256);not null"`
	CreatedAt       time.Time  //`gorm:"not null;index:idx_users_created_at"`
	UpdatedAt       *time.Time
	DeletedAt       *time.Time //`gorm:"index:idx_users_deleted_at"`
	CreatedByUserID int64      //`gorm:"not null"`
}

func (UserModel) TableName() string {
	return "users"
}

// MasterModel represents the masters table
type MasterModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    int64     //`gorm:"not null;index:idx_7493b149a76ed395"`
	IsActive  bool      //`gorm:"not null;default:true"`
	IsBanned  bool      //`gorm:"not null;default:false"`
	CreatedAt time.Time //`gorm:"not null"`
	UpdatedAt time.Time //`gorm:"not null"`
}

func (MasterModel) TableName() string {
	return "masters"
}

// MasterProfileModel represents the master_profile table
type MasterProfileModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	MasterID  int64     //`gorm:"not null;index:idx_5d70476713b3db11"`
	Nickname  string    //`gorm:"type:varchar(64);not null"`
	Avatar    *string   //`gorm:"type:text"`
	CreatedAt time.Time //`gorm:"not null"`
	UpdatedAt time.Time //`gorm:"not null"`
}

func (MasterProfileModel) TableName() string {
	return "master_profile"
}

// MasterAccessLogModel represents the master_access_logs table
type MasterAccessLogModel struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	MasterID      int64     //`gorm:"index;not null"`
	UserID        int64     //`gorm:"index;not null;index:idx_c5e10dfa76ed395"`
	Token         string    //`gorm:"type:text;not null;index:idx_master_access_logstoken"`
	IsTerminated  bool      //`gorm:"not null;default:false"`
	IsExpired     bool      //`gorm:"not null;default:false"`
	ExpiresAt     time.Time //`gorm:"not null;index:idx_master_access_logsexpires_at"`
	RefreshCount  int       //`gorm:"not null;default:0"`
	CreatedAt     time.Time //`gorm:"not null"`
	UpdatedAt     time.Time //`gorm:"not null"`
	IPAddress     *string   //`gorm:"type:varchar(45)"`
	UserAgent     *string   //`gorm:"type:text"`
	RequestsCount int       //`gorm:"not null;default:0"`
	Payload       *string   //`gorm:"type:json"`
}

func (MasterAccessLogModel) TableName() string {
	return "master_access_logs"
}

// AdminModel represents the admins table
type AdminModel struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	UserID       int64     //`gorm:"not null;index:idx_a2e0150fa76ed395"`
	IsActive     bool      //`gorm:"not null;default:true"`
	IsBanned     bool      //`gorm:"not null;default:false"`
	IsSuperadmin bool      //`gorm:"not null;default:false;index:idx_admins_is_superadmin"`
	EmployeeID   *int64
	CreatedAt    time.Time //`gorm:"not null"`
	UpdatedAt    time.Time //`gorm:"not null"`
}

func (AdminModel) TableName() string {
	return "admins"
}

// AdminProfileModel represents the admin_profile table
type AdminProfileModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	AdminID   int64     //`gorm:"not null;index:idx_456b2886642b8210"`
	Nickname  string    //`gorm:"type:varchar(64);not null"`
	Avatar    *string   //`gorm:"type:text"`
	CreatedAt time.Time //`gorm:"not null"`
	UpdatedAt time.Time //`gorm:"not null"`
}

func (AdminProfileModel) TableName() string {
	return "admin_profile"
}

// AdminAccessLogModel represents the admin_access_logs table
type AdminAccessLogModel struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	UserID        int64     //`gorm:"not null;index:idx_9b4ce953a76ed395"`
	Token         string    //`gorm:"type:text;not null;index:idx_admin_access_logstoken"`
	IsTerminated  bool      //`gorm:"not null;default:false"`
	IsExpired     bool      //`gorm:"not null;default:false"`
	ExpiresAt     time.Time //`gorm:"not null;index:idx_admin_access_logsexpires_at"`
	RefreshCount  int       //`gorm:"not null;default:0"`
	CreatedAt     time.Time //`gorm:"not null"`
	UpdatedAt     time.Time //`gorm:"not null"`
	IPAddress     *string   //`gorm:"type:varchar(45)"`
	UserAgent     *string   //`gorm:"type:text"`
	RequestsCount int       //`gorm:"not null;default:0"`
	Payload       *string   //`gorm:"type:json"`
}

func (AdminAccessLogModel) TableName() string {
	return "admin_access_logs"
}

// EmployeeModel represents the employees table
type EmployeeModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    int64     //`gorm:"not null;index:idx_ba82c300a76ed395"`
	UUID      uuid.UUID //`gorm:"type:uuid;not null;uniqueIndex:uniq_employees_uuid"`
	IsActive  bool      //`gorm:"not null;default:true"`
	IsBanned  bool      //`gorm:"not null;default:false"`
	CreatedAt time.Time //`gorm:"not null"`
	UpdatedAt time.Time //`gorm:"not null"`
}

func (EmployeeModel) TableName() string {
	return "employees"
}

// EmployeeProfileModel represents the employee_profile table
type EmployeeProfileModel struct {
	ID         int64      `gorm:"primaryKey;autoIncrement"`
	EmployeeID int64      //`gorm:"not null;index:idx_11bfc008c03f15c"`
	Name       string     //`gorm:"type:varchar(64);not null;uniqueIndex:uniq_employee_profile_name_surname,priority:1"`
	Surname    string     //`gorm:"type:varchar(64);not null;uniqueIndex:uniq_employee_profile_name_surname,priority:2"`
	Birthdate  *time.Time //`gorm:"type:date"`
}

func (EmployeeProfileModel) TableName() string {
	return "employee_profile"
}

// EmployeeAccessLogModel represents the employee_access_logs table
type EmployeeAccessLogModel struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	UserID        int64     //`gorm:"not null;index:idx_87706127a76ed395"`
	Token         string    //`gorm:"type:text;not null;index:idx_employee_access_logstoken"`
	IsTerminated  bool      //`gorm:"not null;default:false"`
	IsExpired     bool      //`gorm:"not null;default:false"`
	ExpiresAt     time.Time //`gorm:"not null;index:idx_employee_access_logsexpires_at"`
	RefreshCount  int       //`gorm:"not null;default:0"`
	CreatedAt     time.Time //`gorm:"not null"`
	UpdatedAt     time.Time //`gorm:"not null"`
	IPAddress     *string   //`gorm:"type:varchar(45)"`
	UserAgent     *string   //`gorm:"type:text"`
	RequestsCount int       //`gorm:"not null;default:0"`
	Payload       *string   //`gorm:"type:json"`
}

func (EmployeeAccessLogModel) TableName() string {
	return "employee_access_logs"
}

// OfficeDepartmentModel represents the office_departments table
type OfficeDepartmentModel struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	Name        string    //`gorm:"type:varchar(64);not null;uniqueIndex:uniq_office_departments_name"`
	Description string    //`gorm:"type:varchar(256);not null"`
	CreatedAt   time.Time //`gorm:"not null"`
	UpdatedAt   time.Time //`gorm:"not null"`
}

func (OfficeDepartmentModel) TableName() string {
	return "office_departments"
}

// OfficeJobModel represents the office_jobs table
type OfficeJobModel struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	DepartmentID int64     //`gorm:"not null;index:idx_457cc6c9ae80f5df"`
	Title        string    //`gorm:"type:varchar(64);not null;uniqueIndex:uniq_office_jobs_title"`
	Description  string    //`gorm:"type:varchar(256);not null"`
	CreatedAt    time.Time //`gorm:"not null"`
	UpdatedAt    time.Time //`gorm:"not null"`
}

func (OfficeJobModel) TableName() string {
	return "office_jobs"
}
