rm -r docs/swagger 2> /dev/null
cp -r -f ../swagger docs/swagger
pipenv shell mkdocs serve
