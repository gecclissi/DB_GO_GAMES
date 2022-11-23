

run_db:
	docker run --name postgres_go_db -e "POSTGRES_PASSWORD=Postgres2022!" -p 5432:5432 -v ${pwd}/db:/var/lib/postgresql/data -d postgres