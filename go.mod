module main

go 1.23.0

require github.com/lib/pq v1.10.9 // indirect

require internal/db v1.0.0

replace internal/db => ./internal/db

require internal/middlewares v1.0.0

replace internal/middlewares => ./internal/middlewares

require internal/routes v1.0.0

replace internal/routes => ./internal/routes

require internal/server v1.0.0

replace internal/server => ./internal/server

require internal/services v1.0.0

replace internal/services => ./internal/services
