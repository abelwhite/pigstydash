# Software Enigineer Final Project
Pigstydash is a dashboard for a pig monitoring system. It was built for my Software Engineer class.

# |----> Report 1

This direcotry contains the PDF for report 1

# |----> Report 2

This direcotry contains the PDF for report 2

# |----> Report 3

This direcotry contains the PDF for report 3

# |----> Brochure

This directory contains the PDF for our project Brochure 


# |----> unit_integration_file

 This directory contains files for unit test and integration test. 
 A readme.md file is found inside with instructions on how to run both tests.

# |----> Pigstydash 

This directory contains files for the final demo. 

1. What you need to run?

    a. Open the pigstydash directory with visual studio. 


    b. Setup a psql database on your local machine called PIGSTYDB_DB_DSN (create as an environmental variable)

    c. install citext

    d. Open the terminal and run the the migrations folder with this code: migrate -path=./migrations -database=$PIGSTYDB_DB_DSN up

    e. Inside the terminal run the command: 
    go run ./cmd/web

    f. This should boot up the server so you can access on your local host on:
    http://localhost:4000/login

    g. You can now signup new user credentials to be authenticated into the app.

2. Report 1 and 2 are submitted as pdf within the folder directory.

3. PowerPoint Slides are also submitted within the direcotry.

4. Report 3 is also submitted as pdf within the folder directory. 

5. Complete source code is located in directory called "pigstydash"

6. Images and icons are stored inside the diretory called "img".
-
|
+---+---pigstydash
    |
    +---+---ui 
        |
	    +---+---static 
            |
		    +---images 
		   
7. HTML Templates are submitted inside the direcorty "html".
-
|
+---+---pigstydash
    |
    +---+---ui 
        |
	    +---html

8. Database tables are submitted as migrations under the direcorty "migrations". 

9. Program only requires psql database with the migrations as stated in step 1.

10. Unit test can be found inside the direcotory unit_integration_tests. A README file is also inside the folder with instruction on how to run it.

11. Integration test can be found inside the direcotory unit_integration_tests.A README file is also inside the folder with instruction on how to run it.


