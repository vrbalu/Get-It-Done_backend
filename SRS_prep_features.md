## Features
An overview of desired features in Get it done project.
### Technical stack used
- Backend - **Golang**
  - libraries - core, gin, http
- Frontend - Angular
  - libraries - Bootstrap, rxjs
  
### Basic functionality
In *Get it done* a user will be able to create projects and in these projects create, see, update, assign and delete
**issues** and track their process. The style will be in a Kanban board. 
The project have **backlog** of issues, these issues can be pulled into **sprints**. 
When sprint is active, there will be displayed active issues assigned to the sprint on a
board. Issues can be assigned to members of a **team**. Teams can be managed. Persons can be added and removed.
Issues are then **filtered** based on teams and projects. **Columns** of the Kanban board will be manageable in Frontend

### Architecture
- REST API on backend with microservices
- Frontend - calling API
### Components

#### Projects
 - CRUD on Project
 - Creator is owner
 - Assign issues to Project
 - Assign Team to Project
 - Have its own unique ID (Mongo collection)
#### Teams
 - CRUD on Teams
 - Persons can be assigned via ID on **profile**?
 - Have its own unique ID (Mongo collection)
#### Issues
 - CRUD on Teams
 - Have various statuses
 - Assigned to users
 - Have its own unique ID - together with project abbreviation in prefix
#### Sprints
 - Have time period
 - Have Name and id
 - MUST be manually closed + offered transfer of issues + warning of being expired
#### Backlog
 - All incomplete issues related to one project
####

### Security
API is secured with API Key. Login in Frontend handled with OAuth.
