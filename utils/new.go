package utils

import (
	"fmt"
	"os"
)

const suffix =".ez"

func CreateFile(title string)error{
	//检查文件是否存在
	fileName:=ezTitle(title)
	_,err:=os.Stat(fileName)
	if err==nil{
		fmt.Println("Failed to create. ")
		fmt.Println("Warning: "+fileName+" is already created,please check it!")
		return err
	}
	//创建文件
	fmt.Println("Creating a new file named: ",fileName)
	newFile,err:=os.Create(fileName)
	if err!=nil{
		fmt.Println("Failed to create. ")
		return err
	}
	err=newFile.Close()
	if err !=nil{
		fmt.Println("Failed to create file: "+fileName)
		return err
	}
	fmt.Println(fileName+" creat complete!")

	//再次检查是否创建成功
	_,err=os.Stat(fileName)
	if err != nil {
		fmt.Println("Wow! the new file created just now does not exit!")
		return err
	}
	return nil
}
func ezTitle(title string) string{
	if len(title) >= len(suffix) && title[len(title)-len(suffix):] != suffix{
		return title+suffix
	}
	return title
}