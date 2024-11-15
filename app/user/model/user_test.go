package model_test

import (
	"gin-seed/app/user/model"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestNewGuest(t *testing.T) {
	fakeIp := gofakeit.IPv4Address()

	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{"Input IP", fakeIp, fakeIp},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			guest := model.NewGuest(tc.in)
			assert.Equal(t, guest.Ip, tc.out, "should be equal")
		})
	}
}

func TestRegister_Success(t *testing.T) {
	guest := model.NewGuest(gofakeit.IPv4Address())
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, true, false, 12)
	credential, _ := model.NewCredential(username, password)

	t.Run("Valid credential", func(t *testing.T) {
		user := guest.Register(credential)
		assert.Equal(t, &user.Credential, credential)
	})
}

func TestRegister_Fail(t *testing.T) {
	guest := model.NewGuest(gofakeit.IPv4Address())

	t.Run("Valid credential", func(t *testing.T) {
		assert.PanicsWithError(
			t,
			"Empty credential",
			func() { guest.Register(nil) },
			"should panic",
		)
	})
}

func TestLogin_Success(t *testing.T) {
	guest := model.NewGuest(gofakeit.IPv4Address())
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, true, false, 12)
	credential, _ := model.NewCredential(username, password)
	user := guest.Register(credential)

	t.Run("", func(t *testing.T) {
		session := user.Login()
		assert.Equal(t, session.UserId, user.Id)
		assert.Equal(t, session.Metadata.Ip, user.Ip)
    assert.NotEqual(t, session.RefreshToken, session.Id, "should have different Id and RereshToken")
	})
}
