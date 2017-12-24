package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ManasikSchedule struct {
	Id          int       `orm:"column(id_schedule_manasik);auto"`
	Materi      string    `orm:"column(materi);size(200)"`
	Location    string    `orm:"column(location);size(100)"`
	Status      string    `orm:"column(status);size(30)"`
	IsPublish   int8      `orm:"column(isPublish)"`
	ScheduleId  int64     `orm:"column(schedule_id)"`
	Createdby   string    `orm:"column(createdby);size(50)"`
	Createddate time.Time `orm:"column(createddate);type(datetime)"`
	Modifyby    string    `orm:"column(modifyby);size(50)"`
	Modifydate  time.Time `orm:"column(modifydate);type(datetime)"`
}

func (t *ManasikSchedule) TableName() string {
	return "manasik_schedule"
}

func init() {
	orm.RegisterModel(new(ManasikSchedule))
}

// AddManasikSchedule insert a new ManasikSchedule into database and returns
// last inserted Id on success.
func AddManasikSchedule(m *ManasikSchedule) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetManasikScheduleById retrieves ManasikSchedule by Id. Returns error if
// Id doesn't exist
func GetManasikScheduleById(id int) (v *ManasikSchedule, err error) {
	o := orm.NewOrm()
	v = &ManasikSchedule{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllManasikSchedule retrieves all ManasikSchedule matches certain condition. Returns empty list if
// no records exist
func GetAllManasikSchedule(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ManasikSchedule))
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

	var l []ManasikSchedule
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

// UpdateManasikSchedule updates ManasikSchedule by Id and returns error if
// the record to be updated doesn't exist
func UpdateManasikScheduleById(m *ManasikSchedule) (err error) {
	o := orm.NewOrm()
	v := ManasikSchedule{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteManasikSchedule deletes ManasikSchedule by Id and returns error if
// the record to be deleted doesn't exist
func DeleteManasikSchedule(id int) (err error) {
	o := orm.NewOrm()
	v := ManasikSchedule{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ManasikSchedule{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
