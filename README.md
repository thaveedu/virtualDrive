# Personal Media Manager

The goal of this application is to provide an interface to manage and sync up files especially (video,images,documents) as a backup in a home network . This app will sync up files from mobile,desktop and other devices in a common location so that it can be searched and retrieved on demand. This will also have a backup copy of the files archieved in Amazon Glacier with encryption so that it is safe to keep personal files backed up in AWS safely.

### Dependencies
```bash
go get github.com/gorilla/mux
```
### Installation & SetUp

- clone the repository.
- RUN ``` go build && ./virtualdrive```
- Access the app in http://localhost:8000

