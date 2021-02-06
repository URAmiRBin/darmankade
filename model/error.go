package model

type NotFoundPage struct {
	Title     string
	HelpTitle string
	HelpLink  string
}

var DatabaseNotFound = NotFoundPage{
	Title:     "ارتباط با دیتابیس برقرار نشد",
	HelpTitle: "لطفا مجددا تلاش کنید",
	HelpLink:  "/",
}

var UnexpectedError = NotFoundPage{
	Title:     "در سیستم مشکلی پیش آمده",
	HelpTitle: "بازگشت به صفحه اصلی",
	HelpLink:  "/",
}

var UserAlreadyExists = NotFoundPage{
	Title:     "این نام کاربری قبلا ثبت شده است",
	HelpTitle: "وارد شوید",
	HelpLink:  "/login.html",
}
