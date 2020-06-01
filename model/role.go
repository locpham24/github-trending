package model

type Role int

const (
	Admin Role = iota
	Moderator
	Member
	Banned
)

func (r Role) String() string {
	return []string{"Admin", "Moderator", "Member", "Banned"}[r]
}
