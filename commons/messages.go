package commons

const (
	DirectoryListing    = "150 Here comes the directory listing.\r\n"
	LinuxSystem         = "215 UNIX Type: L8\r\n"
	Goodbye             = "221 Goodbye.\r\n"
	DirectoryOK         = "226 Directory send OK.\r\n"
	AlreadyLogged       = "230 Already logged in.\r\n"
	UserLoginSuccess    = "230 User logged in, proceed.\r\n"
	DirectoryChanged    = "250 Directory successfully changed.\r\n"
	DirectoryCreated    = "257 Directory created.\r\n"
	UsernameOK          = "331 User name okay, password needed.\r\n"
	NoDataConnection    = "425 Use PORT or PASV first.\r\n"
	InvalidCommand      = "500 Invalid command.\r\n"
	InternalError       = "500 Internal error.\r\n"
	UserFirst           = "503 Login with USER first.\r\n"
	LoginFirst          = "530 Please login with USER and PASS.\r\n"
	IncorrectLogin      = "530 Incorrect Login.\r\n"
	KOChangeUser        = "530 Can't change to another user\r\n"
	NoDirectory         = "550 No directory selected\r\n"
	DirectoryNotFound   = "550 Directory not found.\r\n"
	NoDirectorySelected = "550 No directory selected\r\n"
)
