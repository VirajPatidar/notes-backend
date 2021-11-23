# notes-backend (Backend API for Note Making Application)

_This project serves as backend for [Note Making Application (notes-frontend)](https://github.com/VirajPatidar/notes-frontend). The API was developed using GO Fiber._

<br/>

<!-- **Link to the website:** [https://oems.netlify.app/](https://oems.netlify.app/) -->

**Link to frontend repo:** [https://github.com/VirajPatidar/notes-frontend](https://github.com/VirajPatidar/notes-frontend)


### Tech Stack ###
* GO
* Fiber
* SQL
* HTTP Only cookies


### API Endpoints ###
**Authentication:**
| Method | API Endpoint | Description |
| :---         | :---         | :---         
| `POST`   | `/register`     | To register or sign-up user    |
| `GET`     | `/user`       |  To get details of signed in user      |
| `POST`     | `/login`       |  To log into Note Making Application     |
| `POST`     | `/auth/logout`       |    To log out from Note Making Application   |

<br/>

**Notes:**
| Method | API Endpoint | Description |
| :---         | :---         | :--- 
| `GET` | `/note/{note_id}` | To get a particular note by id |
| `PUT` | `/update-note/{note_id}` | To update/edit a particular note by id |
| `GET` | `/notes` | To get all notes of the user |
| `POST` | `/create-note` | To create a note |
| `DELETE` | `/delete-note/{note_id}` | To delete a particular note by id |
