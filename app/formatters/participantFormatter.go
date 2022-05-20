package formatters

import (
	"base-api/app/entities"
	"time"

	"github.com/google/uuid"
)

type ParticipantFormatter struct {
	ID           uuid.UUID `json:"id"`
	FullName     string    `json:"full_name"`
	BusinessName string    `json:"business_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
}

func FormatParticipant(data entities.Participant) ParticipantFormatter {
	dataFormatter := ParticipantFormatter{}
	dataFormatter.ID = data.ID
	dataFormatter.FullName = data.FullName
	dataFormatter.BusinessName = data.BusinessName
	dataFormatter.Email = data.Email
	dataFormatter.PhoneNumber = data.PhoneNumber
	dataFormatter.CreatedAt = data.CreatedAt

	return dataFormatter
}

func FormatParticipants(datas []entities.Participant) []ParticipantFormatter {
	if len(datas) == 0 {
		return []ParticipantFormatter{}
	}

	datasFormattter := []ParticipantFormatter{}

	for _, data := range datas {
		dataFormatter := FormatParticipant(data)
		datasFormattter = append(datasFormattter, dataFormatter)
	}

	return datasFormattter
}
