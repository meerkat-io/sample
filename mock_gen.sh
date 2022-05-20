mockgen -source ./services/user.go -package services -destination ./services/user_mock.go
mockgen -source ./repositories/user.go -package repositories -destination ./repositories/user_mock.go