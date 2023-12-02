package main

const (
	MainFile           = "package main\n\nimport (\n\t\"fmt\"\n\t\"github.com/spf13/viper\"\n)\n\nfunc main() {\n\tvar app {{.moduleName}}App\n\tviper.SetConfigName(\"properties\")\n\tviper.SetConfigType(\"yaml\")\n\tviper.AddConfigPath(\"./\")\n\terr := viper.ReadInConfig()\n\tif err != nil { // Handle errors reading the config file\n\t\tpanic(fmt.Errorf(\"fatal error config file: %w\", err))\n\t}\n\terr = viper.Unmarshal(&app)\n\tif err != nil {\n\t\tpanic(fmt.Errorf(\"fatal error config file: %w\", err))\n\t}\n\n\tapp.Init()\n}"
	AppFile            = "package main\n\ntype {{.moduleName}}App struct {\n\tEnv        string `mapstructure:\"env\"`\n\tAppName    string `mapstructure:\"appName\"`\n\tHttpServer struct {\n\t\tPort int `mapstructure:\"port\"`\n\t} `mapstructure:\"httpServer\"`\n}\n\nfunc (t {{.moduleName}}App) Init() {\n\n}"
	PropertiesFile     = "env: local\nappName: {{.moduleName}}\nhttpServer:\n  port: 8080"
	GoModFile          = "module {{.moduleRepository}}\n\ngo 1.21.3"
	CoreServiceFile    = "package core\n\nimport \"context\"\n\ntype I{{.moduleName}}Service interface {\n\tHelloWorld(ctx context.Context)\n}"
	CoreRepositoryFile = "package core\n\ntype IRepository interface {\n}"
	GitIgnoreFile      = "*.exe\n*.exe~\n*.dll\n*.so\n*.dylib\n\n*.test\n\n*.out\n\ngo.work\n\n.DS_Store\nbin/\n.idea\n.vs_code"
	DockerFile         = "FROM golang:1.21 as build\nWORKDIR /app\nCOPY go.mod ./\nCOPY go.sum ./\nRUN go mod tidy\nCOPY . ./\nRUN CGO_ENABLED=0 go build -o main\n\nFROM gcr.io/distroless/static-debian11\nWORKDIR /app\nUSER nonroot:nonroot\nCOPY --from=build /app/main /app/main\nCOPY --from=build /app/properties.yml /app/properties.yml\nEXPOSE 8080\nENTRYPOINT [\"./main\"]"

	CoreDirectory      = "core"
	CoreModelDirectory = "core/model"
	AdaptorDirectory   = "adaptor"
	InterfaceDirectory = "interface"
	LibraryDirectory   = "library"
)
