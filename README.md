# Hi Mama Technical Assessment

### Submitting the work

* Frontend Link: https://hi-mama-checkin.herokuapp.com/
* Backend Api Link: https://secure-dawn-73962.herokuapp.com/v1/healthz
* Google drive Link: https://drive.google.com/drive/folders/1zVTeYAElRcV9PiwM-3FIpKzfFnLlwarn?usp=sharing
* Github Link

    * https://github.com/nashmaniac/hi-mama-frontend
    * https://github.com/nashmaniac/hi-mama-backend

### Code challenge

* A teacher will need to clock in and out multiple times a day (e.g., for lunch) <br/>
    This feature is implemented.
* Time entries will need to be editable. <br/>
    This feature is implemented.

* The application should support many users either by adding user names to each
timesheet entry or by adding user authentication. <br/>
    The application supports multiple users by user authentication. An user can register themselves first then login and clock the time.

* Are there any validations the application should do? If so, how should it display errors to
users? <br/>
    We have done several validations. Those are as follows
    
    * An user can't clock in if they are already clock in
    * An user can't clock out if they are not clocked in
    * While editing the time entry, the end time can not be smaller than the start time
    * All the api errors are being shown in a notification toast which pops at the right top corner on the screen.

    <br/>


* Testing is an important part of our workflow at HiMama, consider including tests for at
least one part of the application. <br/>
    I have written the unit test cases for the usecases only as writing all the test cases would be time consuming and fairly repititive. My core focus was to finish the task first. All the test cases written have more than 95% of coverage.



### Written Response
* Please describe your process for approaching the code challenge. What kind of planning
did you do? Did your plans change as you began coding? <br/>
    First I draw the general system diagram on a pen and paper first and try to find out the dependencies. As the time was short, I first started developing the backend application and followed TDD development first but later switched to code first as goal is to finish the challenge on time. But later on, I came and finished the test cases.

    After that, I started code the frontend part of it. I have a boilerplate code that I used from earlier as login/logout feature is pretty straightforward.


* Describe the schema design you chose. Why did you choose this design? What other
alternatives did you consider?
    For the schema design, I chose postgresql as database. The schema design was fairly straightforward. I thought of two models. 
    * User
        * Username - unique field
        * ID - primary key
        * Password
        * Created time
        * Updated time
        * Deleted At
    * Entry
        * UserID - foreign key to the user model
        * ID - primary key
        * Clock in time
        * Clock out time
        * Created time
        * Updated time
        * Deleted At

    I thought of nosql. But at this case, it would be a overkill.

* If you were given another day to work on this, how would you spend it? What if you were
given a month? <br/>

    If I would be given another day, I would 

    * Enable pagination on the time entries
    * Enable filtering time entries by date

    If I would be given another month, I would

    * Build a dashboard so that it enables the admin to see the timing for all the users in the system
    * Enables consolidated reporting so that it enables generating total time reports per user
    
