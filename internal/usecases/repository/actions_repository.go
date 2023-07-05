package repository

type ActionsRepository interface {
	Create(userId string, data []byte) error
}
