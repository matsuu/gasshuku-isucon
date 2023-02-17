package model

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/icrowley/fake"
	"github.com/mattn/go-gimei"
)

type Member struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Banned      bool      `json:"banned"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewMemberName() string {
	if rand.Intn(10) == 0 {
		return fake.FullName()
	}
	return gimei.NewName().Kanji()
}

func NewMemberAddress() string {
	if rand.Intn(10) == 0 {
		return fake.StreetAddress() + ", " + fake.City() + ", " + fake.State() + ", " + fake.Zip()
	}
	return fmt.Sprintf("%s%d-%d-%d", gimei.NewAddress().Kanji(),
		rand.Intn(10), rand.Intn(100), rand.Intn(1000))
}

func NewMemberPhoneNumber() string {
	if rand.Intn(10) == 0 {
		return fake.Phone()
	}
	return fmt.Sprintf("0%d-%d-%d", rand.Intn(10), rand.Intn(10000), rand.Intn(10000))
}
