## About this project

For the sake of simplicity, to avoid creating project structure, router, database migrations, assets, models, sessions, e.t.c. A so called "go web development" ecosystem called [Buffalo](https://gobuffalo.io/en) was used to get up to speed as quickly as possible. I wouldn't call it a full fledged framework as it's merely a carefully wrapped collection of all the popular and commonly used go libraries for doing all sort of different things.

### What works

- Customer listing / searching with pagination, creating, updating
- Warning if someone saves the customer while someone else is editing it
- Localization on two languages, russian and english
- Some basic tests
- Validation for
	- first_name presence, max length 100
	- last_name presence, max length 100
	- email presence, valid email, unique email
	- gender presence, valid gender
	- birth date, no younger 18, no older 60
	- address, max length 200

### What could be better

- Styling for pagination
- Tests, integration tests
- Various seed data with different cases
- Some refactoring with genderOptions helper
- Clean up console output, for example moment complaining on missing locales
- Localize error messages

Unfortunately, since I am leaving on 15.04, I had to skip these things but if you are interested in seeing how do I implement this, I can proceed once I am back.

## Database Setup

Database used in this application is postgres as per requirements.
Prior to starting application, it's neccessary to configure correct database credentials.

The first thing you need to do is open up the "database.yml" file and edit it to use the correct database connection credentials.

### Create Your Databases

There are couple of things you need to do prior to running the application
First of all we need to create both development and test databases.

	$ buffalo db create -a

### Migrate Your Database

Cool, given that you supplied correct credentials in database.yml, you should now have two new databases
`gowatest_development` and `gowatest_test`. Now let's create some tables in our database by running

	$ buffalo db migrate
	$ buffalo db migrate -e test

### Seed Your Database

As a last step, let's add some test data there

	$ buffalo task db:seed

Variety of test data could be better but it's basically there to demonstrate pagination and search functionality

## Starting the Application

That's about it, now in order to run the application you need to run

	$ buffalo dev

And point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000). Bear in mind that since webpack is used here, it might take couple of seconds for assets to compile and before that, styles might look broken. Just refresh a page in couple of seconds and it will work again.

## Testing the Application

Since our tests are located in actions directory, you first need to navigate to actions directory

	$ cd actions

And then execute

	$ go test

to run all tests