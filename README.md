mockgen -destination=mocks/mock_doer.go -package=mocks integral_test/repository DatabaseInterface


go test -v -run TestUsecase_Addition