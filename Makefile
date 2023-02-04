run:
	go run .
## DATABASE SEEDING
# Probably do this in Go later.
seed-initial-menu-products:
	mysql -u root crud_demo < ./database/scripts/001-initial-menu-products-seed.sql