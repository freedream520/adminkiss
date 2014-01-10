package models

//”√ªß±Ì
type User struct {
        Id            int64
        Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
        Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
        Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
        Nickname      string    `orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
        Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
        Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
        Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
        Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
        Createtime    time.Time `orm:"type(datetime);auto_now_add" `
        Role          []*Role   `orm:"rel(m2m)"`
}
