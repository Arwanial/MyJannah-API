package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego"
	 "database/sql"
)

type Travelagent struct {
	Id                int       `orm:"column(id);auto"`
	Email             sql.NullString    `orm:"unique;column(email);size(100);null"`
	NamaTravel        sql.NullString    `orm:"unique;column(nama_travel);size(200)"`
	Phone             string    `orm:"column(phone);size(30)"`
	Fax               string    `orm:"column(fax);size(30)"`
	Mobile            sql.NullString    `orm:"column(mobile);size(30);null"`
	Website           string    `orm:"column(website);size(150)"`
	AlamatKantor      string    `orm:"column(alamat_kantor);size(255)"`
	KotaKantor        string    `orm:"column(kota_kantor);size(100)"`
	Provinsi          string    `orm:"column(provinsi);size(100)"`
	NoKemenagUmroh    sql.NullString    `orm:"column(no_kemenag_umroh);size(20)"`
	KemenangUmrohPath sql.NullString    `orm:"column(kemenag_umroh_path);size(200)"`
	NoKemenangHaji    sql.NullString    `orm:"column(no_kemenag_haji);size(20)"`
	KemenagHajiPath   sql.NullString    `orm:"column(kemenag_haji_path);size(200)"`
	JoinDate          time.Time `orm:"auto_now_add;column(join_date);type(datetime)"`
	Status            string    `orm:"column(status);size(100)"`
	Password          string    `orm:"column(password);size(50)"`
	RegisterNumber		string    `orm:"column(registration_number)"`
	RoleId            int64     `orm:"column(role_id)"`
	Username					string		`orm:"column(username)"`
}


func (t *Travelagent) TableName() string {
	return "travelagent"
}

func init() {
	orm.RegisterModel(new(Travelagent))
}

// AddTravelagent insert a new Travelagent into database and returns
// last inserted Id on success.
func AddTravelagent(m *Travelagent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTravelagentById retrieves Travelagent by Id. Returns error if
// Id doesn't exist
func GetTravelagentById(id int) (v *Travelagent, err error) {
	o := orm.NewOrm()
	v = &Travelagent{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTravelagent retrieves all Travelagent matches certain condition. Returns empty list if
// no records exist
func GetAllTravelagent(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Travelagent))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Travelagent
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTravelagent updates Travelagent by Id and returns error if
// the record to be updated doesn't exist
func UpdateTravelagentById(m *Travelagent) (err error) {
	o := orm.NewOrm()
	v := Travelagent{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTravelagent deletes Travelagent by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTravelagent(id int) (err error) {
	o := orm.NewOrm()
	v := Travelagent{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Travelagent{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetTravelagentById retrieves Travelagent by Id. Returns error if
// Id doesn't exist
func CheckUserByEmailAndPassword (username string, password string) (v *Travelagent, err error) {
	o := orm.NewOrm()
	o.Using("default")

	// var count int
	err = o.Raw("select * from travelagent where username = ? and password= ?", username, password).QueryRow(&v)
  return
}

//Check Account Based On Email Or phone
func CheckAccountByEmailOrPhone (account string) (count int, err error) {
	o := orm.NewOrm()
	o.Using("default")

	err = o.Raw("Select count(*) From travelagent where email=? or mobile=?", account, account).QueryRow(&count)
	return
}
