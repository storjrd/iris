default:
	docker-compose up --force-recreate --build

build:
	docker-compose build --no-cache

dev:
	docker-compose -f docker-compose-dev.yml up --force-recreate --build

test:
	mkdir -p ui-tests-output
	docker-compose -f docker-compose-test.yml up --force-recreate --build --abort-on-container-exit

prettier:
	prettier --write . !ui/browser
	git add .
	git commit -m "prettierize"
	git push

prettier-check:
	prettier --check .

update-browser:
	git pull
	cd ui/browser
	git fetch
	git checkout main
	git pull
	cd ../
	git add .
	git commit -m "bump browser version"
