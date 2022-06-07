package db

import (
	"time"
)

type Testboard struct {
	Users []*User
}

type User struct {
	Username string
}

func (db *Database) SaveUser(user *User) error {
	now := time.Now()
	custom := now.Format("01/02 15:04")
	pipe := db.Client.TxPipeline()
	pipe.LPush(Ctx, "board", custom+" "+user.Username)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetAllUser() (*Testboard, error) {
	boardList := db.Client.LRange(Ctx, "board", 0, -1)
	if boardList == nil {
		return nil, ErrNil
	}
	count := len(boardList.Val())
	users := make([]*User, count)
	for idx, member := range boardList.Val() {
		users[idx] = &User{
			member,
		}
	}
	testboard := &Testboard{
		users,
	}
	return testboard, nil
}
