package hbase

import (
	"errors"
	"fmt"
	"io"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"github.com/tsuna/gohbase/hrpc"
	"golang.org/x/net/context"
	"github.com/alecthomas/log4go"

)

var client gohbase.Client

// Hbase的连接
func ConnectHBase() {
	hbaseUrl := beego.AppConfig.String("hbaseHost")
	user := beego.AppConfig.String("userName")
	option := gohbase.EffectiveUser(user)
	client = gohbase.NewClient(hbaseUrl, option)

}

// 向表中添加数据
func PutsByRowkey(table, rowKey string, values map[string]map[string][]byte) error {
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values)
	if err != nil {
		log4go.Error("hrpc.NewPutStr: %s", err.Error())
		return err
	}
	_, err = client.Put(putRequest)
	if err != nil {
		log4go.Error("clients.Put: %s", err.Error())
		return err
	}
	return nil
}

func UpdataHbase(table, rowKey string, values map[string]map[string][]byte)  error {
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values)
	if err != nil {
		log4go.Error("hrpc.NewPutStr: %s", err.Error())
		return err
	}
	_, err = client.Put(putRequest)
	if err != nil {
		log4go.Error("hbase clients: %s", err.Error())
		return err
	}
	return nil
}

// 查询数据
func Gets(table, rowKey string) (*hrpc.Result, error) {

	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowKey)
	if err != nil {
		log4go.Error("hrpc.NewGetStr: %s", err.Error())
		return nil, err
	}
	res, err := client.Get(getRequest)
	if err != nil {
		log4go.Error("hbase clients: %s", err.Error())
		return nil, err
	}
	defer func() {
		if errs := recover(); errs != nil {
			switch fmt.Sprintf("%v", errs) {
			case "runtime error: index out of range":
				err = errors.New("no such rowKey or qualifier exception")
			case "runtime error: invalid memory address or nil pointer dereference":
				err = errors.New("no such colFamily exception")
			default:
				err = fmt.Errorf("%v", errs)
			}
		}
	}()
	return res, nil
}

// 查看rowkey是否存在
func IsExistRowkey(table, rowKey string) bool {
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowKey)
	if err != nil {
		log4go.Error("hrpc.NewGetStr: %s", err.Error())
		return false
	}
	res, err := client.Get(getRequest)
	if err != nil {
		log4go.Error("get from hbase: %s", err.Error())
		return false
	}
	if len(res.Cells) > 0 {
		return true
	} else {
		return false
	}
}

// 删除数据
func DeleteByRowkey(table, rowkey string, value map[string]map[string][]byte)  error {
	deleteRequest, err := hrpc.NewDelStr(context.Background(), table, rowkey, value)
	if err != nil {
		log4go.Error("hrpc.NewDelStrRef: %s", err.Error())
		return err
	}
	//fmt.Println("deleteRequest:", deleteRequest)
	_, err = client.Delete(deleteRequest)
	if err != nil {
		log4go.Error("hrpc.Scan: %s", err.Error())
		return err
	}
	return nil
}

// 分页：其中startRow为页码列表的开始索引，stopRow页码列表的结束索引，limit每页显示多少条
func PagedQuery(table, startRow, stopRow string, limit int64) (rsp []*hrpc.Result, err error) {
	var (
		scanRequest *hrpc.Scan
		res         *hrpc.Result
	)

	pFilter := filter.NewPageFilter(limit + 1)
	scanRequest, err = hrpc.NewScanRangeStr(context.Background(), table, startRow, stopRow, hrpc.Reversed(), hrpc.Filters(pFilter))
	if err != nil {
		log4go.Error("hrpc.NewScanStr: %s", err.Error())
	}

	scanner := client.Scan(scanRequest)

	for {
		res, err = scanner.Next()
		if err == io.EOF || res == nil {
			break
		}
		if err != nil {
			log4go.Error("hrpc.Scan: %s", err.Error())
		}
		rsp = append(rsp, res)
	}

	return rsp, err
}

func GetHbaseIndex(startRow,stopRow, table string, limit int64) (rsp []*hrpc.Result, str string) {
	var nextIndex  string
	temScanner, err := PagedQuery(table, startRow, stopRow, limit)
	if err != nil {
		fmt.Println("GetHbaseIndex with limit: ", err)
	}
	if int64(len(temScanner)) <= limit {
		nextIndex = ""
		return temScanner, nextIndex
	} else {
		nextIndex = string(temScanner[limit].Cells[0].Row)
		return temScanner[:len(temScanner)-1], nextIndex
	}

}
