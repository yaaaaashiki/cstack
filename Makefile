DBNAME := cstack_dev
TESTDBNAME := cstack_test
DBCONF := dbconfig.yml
ENV    := development
TESTENV    := test 

SQL_MIGRATE_TASKS := migrate/status migrate/up

.PHONY: help sql-migrate
.DEFAULT_GOAL: help

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "\033[36m%-25s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST)

run: ## Run app
	go run ./cmd/cstack/cstack.go

lint:
	[ -z `gofmt -l . | grep -v vendor -v '.glide/cache'` ]

test: lint ## Run tests
	go test -v $$(go list ./... | grep -v /vendor/ )

integration-test: ## Run integration tests
	go test -tags=integration -v $$(go list ./... | grep -v /vendor/ )

gen:
	which scaneo || go get github.com/variadico/scaneo
	cd model && go generate

sql-migrate:
	which $@ || go get github.com/rubenv/sql-migrate/...

################################################################################################################

drop: PROTOCOL=tcp
drop:
	mysql -u root -h localhost --protocol $(PROTOCOL) -p$(dbpass) \
	-e "drop DATABASE \`$(DBNAME)\`"

migrate/init: PROTOCOL=tcp
migrate/init:
	mysql -u root -h localhost --protocol $(PROTOCOL) -p$(dbpass) \
	-e "CREATE DATABASE IF NOT EXISTS \`$(DBNAME)\` /*!40100 DEFAULT CHARACTER SET utf8 */"

seed: PROTOCOL=tcp
seed:
	mysql -u root -h localhost --protocol $(PROTOCOL) -p$(dbpass) $(DBNAME) < ./db/seed/seed.sql

developenv:
	make drop PROTOCOL=socket
	make migrate/init PROTOCOL=socket
	make migrate/up
	make seed PROTOCOL=socket 

################################################################################################################

drop/test: PROTOCOL=tcp
drop/test:
	mysql -u root -h localhost --protocol $(PROTOCOL) -p$(dbpass) \
	-e "drop DATABASE \`$(TESTDBNAME)\`"

migrate/init/test: PROTOCOL=tcp
migrate/init/test:
	mysql -u root -h localhost --protocol $(PROTOCOL) -p$(dbpass) \
	-e "CREATE DATABASE IF NOT EXISTS \`$(TESTDBNAME)\` /*!40100 DEFAULT CHARACTER SET utf8 */"

seed/test: PROTOCOL=tcp
seed/test: dbpass = password
seed/test:
	mysql -u root -h localhost --protocol $(PROTOCOL) $(TESTDBNAME) < ./db/seed/seed.sql

testenv:
	make drop/test PROTOCOL=socket
	make migrate/init/test PROTOCOL=socket 
	make migrate/up ENV=test 
	make seed/test PROTOCOL=socket

################################################################################################################

$(SQL_MIGRATE_TASKS): sql-migrate
	sql-migrate $(@F) -env=$(ENV) -config=$(DBCONF)

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: deps
deps: glide
	glide install
