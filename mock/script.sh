#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen


# repository
mockgen -package=mock -destination=mock/mock_AuthenticationRepository.go github.com/AndiGanesha/authentication/repository IAuthenticationRepository