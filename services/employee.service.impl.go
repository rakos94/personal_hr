package services

import (
	"errors"
	"personal_hr/dao"
	"personal_hr/helper"
	"personal_hr/models"
	"strconv"
)

var employeeDao dao.EmployeeDao = dao.EmployeeDaoImpl{}
var personService PersonService = PersonServiceImpl{}

// EmployeeServiceImpl ...
type EmployeeServiceImpl struct{}

// CreateEmployee ...
func (EmployeeServiceImpl) CreateEmployee(emp *models.Employee) (*models.Employee, error) {
	_, err := personService.GetPersonByID(emp.PersonID)
	if err != nil {
		return nil, err
	}
	_, err = employeeDao.GetEmployeeByPersonID(emp.PersonID)
	if err == nil {
		return nil, errors.New("Person Id already exist")
	}

	count, err := employeeDao.GetEmployeeCount()
	if err != nil {
		return nil, err
	}

	emp.Nip = GenerateNIP(count)
	return employeeDao.CreateEmployee(emp)
}

// GetEmployeeByID ...
func (EmployeeServiceImpl) GetEmployeeByID(id string) (models.Employee, error) {
	return employeeDao.GetEmployeeByID(id)
}

// GetEmployeeAll ...
func (EmployeeServiceImpl) GetEmployeeAll() ([]models.Employee, error) {
	return employeeDao.GetEmployeeAll()
}

// GenerateNIP ...
func GenerateNIP(count int64) string {
	nip := "EMP."
	count = count + 1
	newID := strconv.FormatInt(count, 10)
	nip += helper.GenerateId(newID, "0", 6)
	return nip
}
