package mockstore

//go:generate mockgen -destination=./store.go -package=mockstore github.com/opencars/bot/pkg/domain MessageRepository
