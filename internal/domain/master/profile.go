package master

import (
	"time"
)

type MasterProfile struct {
	id        int64
	masterID  int64
	nickname  string
	avatar    *string
	createdAt time.Time
	updatedAt time.Time
}

func NewMasterProfile(
	masterID int64,
	nickname string,
) *MasterProfile {
	now := time.Now()

	return &MasterProfile{
		masterID:  masterID,
		nickname:  nickname,
		avatar:    nil,
		createdAt: now,
		updatedAt: now,
	}
}

func ReconstructMasterProfile(
	id int64,
	masterID int64,
	nickname string,
	avatar *string,
	createdAt, updatedAt time.Time,
) *MasterProfile {
	return &MasterProfile{
		id:        id,
		masterID:  masterID,
		nickname:  nickname,
		avatar:    avatar,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// Getters
func (p *MasterProfile) ID() int64             { return p.id }
func (p *MasterProfile) MasterID() int64       { return p.masterID }
func (p *MasterProfile) Nickname() string      { return p.nickname }
func (p *MasterProfile) Avatar() *string       { return p.avatar }
func (p *MasterProfile) CreatedAt() time.Time  { return p.createdAt }
func (p *MasterProfile) UpdatedAt() time.Time  { return p.updatedAt }

// Business methods
func (p *MasterProfile) UpdateNickname(nickname string) {
	p.nickname = nickname
	p.updatedAt = time.Now()
}

func (p *MasterProfile) UpdateAvatar(avatar *string) {
	p.avatar = avatar
	p.updatedAt = time.Now()
}

func (p *MasterProfile) RemoveAvatar() {
	p.avatar = nil
	p.updatedAt = time.Now()
}

/*
type MasterProfile struct {
	userID      int64 //int64
	nickname    string
	avatar      string
	birthDate   *valueobject.DateTime
	phoneNumber string
	department  string
}

func NewMasterProfile(
	userID int64,
	name, surname, phoneNumber, department string,
	birthDate *time.Time,
) *MasterProfile {
	var birthDateVO *valueobject.DateTime
	if birthDate != nil {
		dt := valueobject.NewDateTime(*birthDate)
		birthDateVO = &dt
	}

	return &MasterProfile{
		userID:      userID,
		name:        name,
		surname:     surname,
		birthDate:   birthDateVO,
		phoneNumber: phoneNumber,
		department:  department,
	}
}

func ReconstructMasterProfile(
	userID int64,
	name, surname, phoneNumber, department string,
	birthDate *time.Time,
) *MasterProfile {
	return NewMasterProfile(userID, name, surname, phoneNumber, department, birthDate)
}
*/
// Getters
//func (ap *MasterProfile) UserID() int64          { return ap.userID }

// Business methods
/* func (ap *MasterProfile) UpdateSurname(surname string) {
	ap.surname = surname
}

func (ap *MasterProfile) UpdatePhoneNumber(phoneNumber string) {
	ap.phoneNumber = phoneNumber
}

func (ap *MasterProfile) UpdateDepartment(department string) {
	ap.department = department
}

func (ap *MasterProfile) UpdateBirthDate(birthDate *time.Time) {
	if birthDate != nil {
		dt := valueobject.NewDateTime(*birthDate)
		ap.birthDate = &dt
	} else {
		ap.birthDate = nil
	}
}

func (ap *MasterProfile) FullName() string {
	return ap.name + " " + ap.surname
} */
