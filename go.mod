module submarine_v0.1

go 1.16

require (
	github.com/emicklei/go-restful v2.15.0+incompatible
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	repositories v0.0.0-00010101000000-000000000000
	submarineapi v0.0.0-00010101000000-000000000000
	usecases v0.0.0-00010101000000-000000000000 // indirect
)

replace submarineapi => ./gateway/api

replace repositories => ./gateway/repositories

replace usecases => ./usecases

replace domain => ./domain
