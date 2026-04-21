package commons

const (
	// 1xx — Positive Preliminary
	DirectoryListing     = "150 Here comes the directory listing.\r\n"
	FileTransferStarting = "150 Opening BINARY mode data connection for file transfer.\r\n"

	// 2xx — Positive Completion
	FeatStart            = "211-Features:\r\n"
	FeatEnd              = "211 End\r\n"
	LinuxSystem          = "215 UNIX Type: L8\r\n"
	Welcome              = "220 Welcome to MyFTP\r\n"
	Goodbye              = "221 Goodbye.\r\n"
	DirectoryOK          = "226 Directory send OK.\r\n"
	FileTransferComplete = "226 Transfer complete.\r\n"
	AlreadyLogged        = "230 Already logged in.\r\n"
	UserLoginSuccess     = "230 User logged in, proceed.\r\n"
	CommandOK            = "200 OK.\r\n"
	PortOK               = "200 PORT command successful.\r\n"
	// use fmt.Sprintf.
	TypeSet              = "200 Type set to %s.\r\n"
	DirectoryChanged     = "250 Directory successfully changed.\r\n"
	FileDeleted          = "250 File deleted successfully.\r\n"
	// (the path); use fmt.Sprintf.
	CurrentDirectory     = "257 \"%s\" is the current directory\r\n"
	DirectoryCreated     = "257 Directory created.\r\n"
	// (h1,h2,h3,h4,p1,p2 address); use fmt.Sprintf.
	EnteringPassiveMode  = "227 Entering Passive Mode (%s)\r\n"

	// 3xx — Positive Intermediate
	UsernameOK = "331 User name okay, password needed.\r\n"

	// 4xx — Transient Negative
	CantOpenDataConn = "425 Can't open data connection.\r\n"
	NoDataConnection = "425 Use PORT or PASV first.\r\n"

	// 5xx — Permanent Negative
	InvalidCommand      = "500 Invalid command.\r\n"
	InternalError       = "500 Internal error.\r\n"
	SyntaxError         = "501 Syntax error in parameters.\r\n"
	PortIPMismatch      = "501 PORT IP does not match client address.\r\n"
	UserFirst           = "503 Login with USER first.\r\n"
	TypeNotSupported    = "504 Type not supported.\r\n"
	LoginFirst          = "530 Please login with USER and PASS.\r\n"
	IncorrectLogin      = "530 Incorrect Login.\r\n"
	AuthenticationError = "530 Authentication error.\r\n"
	KOChangeUser        = "530 Can't change to another user\r\n"
	InvalidPath         = "550 Invalid path\r\n"
	DirectoryNotFound   = "550 Directory not found.\r\n"
	NoDirectorySelected = "550 No directory selected\r\n"
	FileNotFound        = "550 File not found.\r\n"
	FileDeleteError     = "550 Cannot delete file.\r\n"
	FileOpenError       = "550 Cannot open file.\r\n"
	FileTransferError   = "550 File transfer failed.\r\n"
	ListFailed          = "550 Failed to list directory\r\n"
)
