

## Accounting Notebook

This project is a sample accounting system web application.
We emulate debit and credit operations for a single user.
The application includes:
- Go + Echo for backend
- React + TypeScript for frontend
- Bootstrap theme

### Running the app


Binary files have been inclued for modern platforms, to start the app, simply go to the directory and run the approppiate version, for example:<br />
- cd accounting_notebook
- ./server.go-linux-amd64

Open [http://localhost:8080](http://localhost:8080) to view it in the browser.


The list of versions included:
- MacOS - "server.go-darwin-386"
- MacOS - "server.go-darwin-amd64"
- Linux - "server.go-linux-386"
- Linux - "server.go-linux-amd64"
- Windows - "server.go-windows-386"
- Windows - "server.go-windows-amd64"

** You can also build this app by yourself if those deliverables didn't work for you, you are going to need installed in your computer Go and yarn, you can follow these operations to do so:

- cd accounting_notebook
- yarn install
- yarn build
- go build server.go
- ./server

### `yarn build`

Builds the fronend libraries for production to the `public` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

### `go build server.go`

Builds the Go module into a binary executable

## API Endpoints

Included in this repository, you can find an example Postman collection: **Accounting Notebook.postman_collection.json**, simply import it to start using<br />

- **GET /**: displays transactions list
- **GET /api/transactions**: Obtains the list of transactions
- **GET /api/transactions/:id**: Find transaction by ID
- **GET /api/default**:  Obtains the current balance of the user
