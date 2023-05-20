# PIGSTY PROJECT

This is our Software Engineering Project using mongodb to connect to a pig monitoring system. 

# Unit Test

To run Unit Test open terminal and type "go test". 
Terminal Should load up and should display "PASS" and "ok" if the test passed.


# Integration Testing

Getting Started

1. You will need an Atlas MongoDb Account if you would like to create your own account. Login to your account and create a new connection. Get your mongo-driver module and use it to connect to MongoDb from your go program.


![image](https://user-images.githubusercontent.com/123121590/229192181-5a199726-cf97-4738-9487-b66012e01e46.png)

2. Step 1. can be skipped if you are just testing my program. Mines alread has the driver module.


3. First, run the server: Open Terminal and type "go run main.go"

4. Install visual studio extension, Thunder Client.
 
5. Create a New Request and type in local houst rout, example:
 http://localhost:4000/api/
 
6. Navigate into the tab "body" and under the json field, type in the json string to insert new room and press "Send".
 ![image](https://user-images.githubusercontent.com/123121590/229192541-7afdb8df-063d-49d4-8756-9fcf69c18eaf.png)

 
7. Terminal should indicate that the new room was added into the database.
 ![image](https://user-images.githubusercontent.com/123121590/229192321-d39da85b-7540-438f-9b97-57b7368d32de.png)

 
8. Entry should be seen inside MongoDb. 
 ![image](https://user-images.githubusercontent.com/123121590/229192048-ec70f287-b1b9-496a-a1eb-4e062c5ba431.png)

 


