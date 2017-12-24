package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Jamaah struct {
	Id         int       `orm:"column(jamaah_id);pk"`
	Firstname  string    `orm:"column(firstname);size(100)"`
	Lastname   string    `orm:"column(lastname);size(100)"`
	Sex        string    `orm:"column(sex);size(20)"`
	Birthdate  time.Time `orm:"column(birthdate);type(date)"`
	Passportno string    `orm:"column(passportno);size(30)"`
	Status     string    `orm:"column(status);size(50)"`
	Address    string    `orm:"column(address);size(200)"`
	City       int       `orm:"column(city)"`
	Province   string    `orm:"column(province);size(30)"`
	Phone      string    `orm:"column(phone);size(20)"`
	Mobile     string    `orm:"column(mobile);size(20)"`
	Email      string    `orm:"column(email);size(100)"`
	ScheduleId int64     `orm:"column(schedule_id);null"`
}

func (t *Jamaah) TableName() string {
	return "jamaah"
}

func init() {
	orm.RegisterModel(new(Jamaah))
}

// AddJamaah insert a new Jamaah into database and returns
// last inserted Id on success.
func AddJamaah(m *Jamaah) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetJamaahById retrieves Jamaah by Id. Returns error if
// Id doesn't exist
func GetJamaahById(id int) (v *Jamaah, err error) {
	o := orm.NewOrm()
	v = &Jamaah{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllJamaah retrieves all Jamaah matches certain condition. Returns empty list if
// no records exist
func GetAllJamaah(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Jamaah))
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

	var l []Jamaah
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

// UpdateJamaah updates Jamaah by Id and returns error if
// the record to be updated doesn't exist
func UpdateJamaahById(m *Jamaah) (err error) {
	o := orm.NewOrm()
	v := Jamaah{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteJamaah deletes Jamaah by Id and returns error if
// the record to be deleted doesn't exist
func DeleteJamaah(id int) (err error) {
	o := orm.NewOrm()
	v := Jamaah{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Jamaah{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
