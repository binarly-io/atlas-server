module github.com/binarly-io/atlas-server

go 1.13

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/binarly-io/atlas v0.0.0-20200621094407-4a8319aa73e5
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.5.1 // indirect
	google.golang.org/grpc v1.28.1
)

replace golang.org/x/oauth2 => github.com/sunsingerus/oauth2 v0.0.0-20200410181841-d7afaacd4cbe
