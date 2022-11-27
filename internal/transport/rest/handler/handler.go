package handler

type Service interface {

}

type Handler struct {
	AuthHandler AuthHandler
	AccountHandler AccountHandler
}


// auth/singin
// auth/singup
	// Проверка на пустату 
	// проверка на дубль 
	// проверка пароля на пустату 
	// сформировать структуру и передать сервису 
// auth/logout

// account/changePassword
// account/getInfo
// account/setInfo