mockery --quiet --dir internal/gateway --name IMessagePublisher
printf "Generated Mocks for internal/gateway/IMessagePublisher\n"

mockery --quiet --dir internal/validators --name IRequestValidator
printf "Generated Mocks for internal/validators/IRequestValidator\n"


mockery --quiet --dir internal/repository --name IEventRepository
printf "Generated Mocks for internal/repository/IEventRepository\n"


mockery --quiet --dir internal/repository --name IUserRepository
printf "Generated Mocks for internal/repository/IUserRepository\n"


mockery --quiet --dir internal/service --name IAuthService
printf "Generated Mocks for internal/service/IAuthService\n"

mockery --quiet --dir internal/service --name IGenerator
printf "Generated Mocks for internal/service/IGenerator\n"

printf "Done!!\n"