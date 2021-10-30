/**
 * @Author: vincent
 * @Description:
 * @File:  mappter
 * @Version: 1.0.0
 * @Date: 2021/10/29 16:05
 */

package demo

import demoService "go-examples/course/handwriting-web-inf/code_13/app/provider/demo"

func StudentsToUserDTOs(students []demoService.Student) []UserDTO {
	ret := []UserDTO{}
	for _, student := range students {
		t := UserDTO{
			ID:   student.ID,
			Name: student.Name,
		}
		ret = append(ret, t)
	}
	return ret
}
