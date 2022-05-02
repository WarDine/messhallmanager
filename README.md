```
.
├── domain
│   │
│   └── postgress.go        <-- All interface definitions go here
│
├── gateway                 <-- Access to bussiness logic goes here 
│   │
│   └── repositories        <-- Database operations go here 
│       │
│       └── postgress.go
│   │
│   └── api                 <-- API logic goes here
│       │
│       └── api.go
│
├── go.mod
│
├── main.go
│
├── README.md         
│
└── usecases                <-- All entities logic goes here
    │
    └── recipe.go
```
    
