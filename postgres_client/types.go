package postgres_client

import "fmt"

type AccountEmailAddressObject struct {
	Id, UserId        int
	Email             string
	Verified, Primary bool
}

func (emailAddr AccountEmailAddressObject) String() string {
	return fmt.Sprintf(
		"AccountEmailAddressObject[id=%d, email=%s, verified=%t, primary=%t, user_id=%d]",
		emailAddr.Id, emailAddr.Email, emailAddr.Verified, emailAddr.Primary, emailAddr.UserId,
	)
}
