module submarineapi_v0.1

go 1.16

require (
	github.com/emicklei/go-restful v2.15.0+incompatible
	github.com/json-iterator/go v1.1.12 // indirect
	repositories v0.0.0-00010101000000-000000000000
)

replace usecases => ./../../usecases

replace domain => ./../../domain

replace repositories => ./../../gateway/repositories
