package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PaketTravel struct {
	Id          int       `orm:"column(id_packet);auto"`
	PacketName  string    `orm:"column(packet_name);size(30)"`
	Longtrip    int       `orm:"column(longtrip)"`
	Origin      string    `orm:"column(origin);size(3)"`
	Destination string    `orm:"column(destination);size(3)"`
	Category    string    `orm:"column(category);size(50)"`
	Description string    `orm:"column(description);size(255)"`
	IsPromo     int8      `orm:"column(isPromo)"`
	Promocode   string    `orm:"column(promocode);size(20)"`
	Createdby   string    `orm:"column(createdby);size(30)"`
	Createddate time.Time `orm:"column(createddate);type(datetime)"`
	Modifyby    string    `orm:"column(modifyby);size(30)"`
	Modifydate  time.Time `orm:"column(modifydate);type(datetime)"`
	Quota       int       `orm:"column(quota)"`
	IsPublish   int8      `orm:"column(isPublish)"`
	Id_RENAME   int64     `orm:"column(id)"`
}

func (t *PaketTravel) TableName() string {
	return "paket_travel"
}

func init() {
	orm.RegisterModel(new(PaketTravel))
}

// AddPaketTravel insert a new PaketTravel into database and returns
// last inserted Id on success.
func AddPaketTravel(m *PaketTravel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPaketTravelById retrieves PaketTravel by Id. Returns error if
// Id doesn't exist
func GetPaketTravelById(id int) (v *PaketTravel, err error) {
	o := orm.NewOrm()
	v = &PaketTravel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPaketTravel retrieves all PaketTravel matches certain condition. Returns empty list if
// no records exist
func GetAllPaketTravel(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PaketTravel))
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

	var l []PaketTravel
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

// UpdatePaketTravel updates PaketTravel by Id and returns error if
// the record to be updated doesn't exist
func UpdatePaketTravelById(m *PaketTravel) (err error) {
	o := orm.NewOrm()
	v := PaketTravel{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePaketTravel deletes PaketTravel by Id and returns error if
// the record to be deleted doesn't exist
func DeletePaketTravel(id int) (err error) {
	o := orm.NewOrm()
	v := PaketTravel{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PaketTravel{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
