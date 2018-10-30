package api

const (
	_ UserRole = iota
	AdminRole
	NormalUserRole
)

// APIs about user management
type (

	// User represents a user account
	User struct {
		ID       UserID   `json:"Id"`
		Username string   `json:"Username"`
		Password string   `json:"Password,omitempty"`
		Role     UserRole `json:"Role"`
	}

	// UserID represents a user identifier
	UserID int

	// UserRole represents the role of a user. It can be either an admin or a regular user
	UserRole int

	// UserService is used for user manager
	UserService interface {
		User(UserID) (*User, error)
		UserByUsername(string) (*User, error)
		Users() ([]User, error)
		UsersByRole(UserRole) ([]User, error)
		UpdateUser(UserID, *User) error
		CreateUser(*User) error
		DeleteUser(UserID) error
	}
)

// APIs about authentication
type (

	// TokenData represents the data embedded in a JWT token
	TokenData struct {
		ID       UserID
		Username string
		Role     UserRole
	}

	// CryptoService represents a fileService for encrypting/hashing data
	CryptoService interface {
		Hash(data string) (string, error)
		Verify(hash string, data string) error
	}

	// JWTService represents a fileService for managing JWT tokens
	JWTService interface {
		// Generate token based on TokenData
		Generate(*TokenData) (string, error)
		// Verify and decrypt generated token
		Decrypt(string) (*TokenData, error)
	}
)

// APIs about files
type (

	// FileService represents a fileService for managing files
	FileService interface {
		FileExists(path string) (bool, error)
	}

	// DataStore represents a fileService for store information
	DataStore interface {
		Open() error
		Init() error
		Close() error
		GetUserService() UserService
	}
)
